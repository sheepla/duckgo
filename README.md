<div align="right">

[![goreleaser](https://github.com/sheepla/duckgo/actions/workflows/release.yml/badge.svg)](https://github.com/sheepla/duckgo/actions/workflows/release.yml)

</div>

<div align="center">

<h1>🦆 duckgo</h1>

**duckgo** is a command line tool to query DuckDuckGo and open items from search result with Web browser quickly
</div>



## Usage

```
Usage: duckgo [--json] [--browser BROWSER] [--version] [QUERY [QUERY ...]]

Positional arguments:
  QUERY                  keywords to search

Options:
  --json, -j             output results in JSON format
  --browser BROWSER, -b BROWSER
                         the command of Web browser to open URL
  --version              show version
  --help, -h             display this help and exit
```

## Installation

```
go install github.com/sheepla/duckgo@latest
```

## Author

[sheepla](https://github.com/sheepla)

## License

MIT

