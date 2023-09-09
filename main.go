package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	fzf "github.com/koki-develop/go-fzf"
	"github.com/sheepla/duckgo/client"
	"github.com/skratchdot/open-golang/open"
)

type Options struct {
	//Open  bool     `help:"open target URL in default browser"`
	Json bool `arg:"-j, --json" help:"output results in JSON format"`
	//Shell bool     `help:"start bash-like interactive mode instead of fuzzy-finder UI"`
	Page    bool     `arg:"-p, --page" default:"1" help:"index of page"`
	Browser string   `arg:"-b, --browser" help:"browser"`
	Query   []string `arg:"positional" help:"keywords to search"`
}

func main() {
	opts, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := run(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func parseArgs(args []string) (*Options, error) {
	var opts Options

	p, err := arg.NewParser(arg.Config{
		Program:   "duckgo",
		IgnoreEnv: false,
	}, &opts)
	if err != nil {
		return &opts, err
	}

	if err := p.Parse(args); err != nil {
		switch {
		case errors.Is(err, arg.ErrHelp):
			p.WriteHelp(os.Stderr)
			return &opts, nil
		case errors.Is(err, arg.ErrVersion):
			p.WriteUsage(os.Stderr)
		default:
			return &opts, err
		}
	}

	return &opts, nil
}

func run(opts *Options) error {
	result, err := client.Search(
		client.NewSearchParam(strings.Join(opts.Query, " ")),
	)
	if err != nil {
		return err
	}

	if opts.Json {
		if err := json.NewDecoder(os.Stdout).Decode(&result); err != nil {
			return err
		}

		return nil
	}

	selected, err := find(*result)
	if err != nil {
		return err
	}

	for _, idx := range selected {
		if opts.Browser == "" {
			if err := open.Run((*result)[idx].Link); err != nil {
				return err
			}

			return nil
		} else {
			if err := open.RunWith((*result)[idx].Link, opts.Browser); err != nil {
				return err
			}

			return nil
		}

	}

	return nil
}

func find(result []client.SearchResult) ([]int, error) {
	f, err := fzf.New(
		fzf.WithInputPlaceholder("Filter..."),
	)
	if err != nil {
		panic(err)
	}

	return f.Find(
		result,
		func(i int) string {
			return result[i].Title
		},
		fzf.WithPreviewWindow(func(idx int, width int, height int) string {
			return fmt.Sprintf(
				"\n\n%s\n\n──────────────────\n\n%s\n\n%s\n",
				result[idx].Title, result[idx].Snippet, result[idx].Link,
			)
		}),
	)
}
