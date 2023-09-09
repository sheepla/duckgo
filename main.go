package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sheepla/duckgo/client"

	"github.com/alexflint/go-arg"
	fzf "github.com/koki-develop/go-fzf"
	"github.com/mattn/go-runewidth"
	"github.com/skratchdot/open-golang/open"
)

const (
	appVersion  = "unknown"
	appRevision = "unknown"
)

type Options struct {
	Json bool `arg:"-j, --json" help:"output results in JSON format"`
	//Shell bool     `help:"start bash-like interactive mode instead of fuzzy-finder UI"`
	//Page    bool     `arg:"-p, --page" default:"1" help:"index of page"`
	Browser string   `arg:"-b, --browser" help:"the command of Web browser to open URL"`
	Query   []string `arg:"positional" help:"keywords to search"`
	Version bool     `help:"show version"`
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
			fmt.Fprintf(os.Stderr, "%s-%s\n", appVersion, appRevision)
		default:
			return &opts, err
		}
	}

	return &opts, nil
}

func run(opts *Options) error {
	param, err := client.NewSearchParam(strings.Join(opts.Query, " "))
	if err != nil {
		return err
	}

	result, err := client.Search(param)
	if err != nil {
		return err
	}

	if opts.Json {
		if err := json.NewEncoder(os.Stdout).Encode(&result); err != nil {
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
			content := fmt.Sprintf(
				"\n\n  %s\n\n──────────────────\n\n  %s\n\n  %s\n",
				result[idx].Title, result[idx].Snippet, result[idx].Link,
			)

			return runewidth.Wrap(content, width-2)
		}),
	)
}
