<div align="right">

[![goreleaser](https://github.com/sheepla/duckgo/actions/workflows/release.yml/badge.svg)](https://github.com/sheepla/duckgo/actions/workflows/release.yml)

</div>

<div align="center">

<h1>🦆 duckgo</h1>

**duckgo** is a command line tool to query DuckDuckGo and open items from search result with Web browser quickly.

This tool is inspired by [ddgr](https://github.com/jarun/ddgr). However, it has been developed with an even greater focus on agility and portability.

</div>

## Usage

```
Usage: duckgo [--json] [--timeout TIMEOUT] [--user-agent USER-AGENT] [--referrer REFERRER] [--browser BROWSER] [--version] [QUERY [QUERY ...]]

Positional arguments:
  QUERY                  keywords to search

Options:
  --json, -j             output results in JSON format
  --timeout TIMEOUT, -t TIMEOUT
                         timeout seconds
  --user-agent USER-AGENT, -u USER-AGENT
                         User-Agent value
  --referrer REFERRER, -r REFERRER
                         Referrer value
  --browser BROWSER, -b BROWSER
                         the command of Web browser to open URL
  --version              show version
  --help, -h             display this help and exit
```

## Installation

You can download executable binaries from the release page.

> [latest release](https://github.com/sheepla/duckgo/releases/latest)

To build from source, run `make build`.

## Similar projects

- [ddgr](https://github.com/jarun/ddgr) - DuckDuckGo client written in Python
- [googler](https://github.com/jarun/googler) - Google client written in Python
- [rearx](https://github.com/garak92/rearx) - TUI client for the [searx](https://github.com/searx/searx) meta search engine written in Rust

## Author

[sheepla](https://github.com/sheepla)

## License

MIT

