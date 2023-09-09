<div align="right">

[![goreleaser](https://github.com/sheepla/duckgo/actions/workflows/release.yml/badge.svg)](https://github.com/sheepla/duckgo/actions/workflows/release.yml)

</div>

<div align="center">

<h1>🦆 duckgo</h1>

**duckgo** is a command line tool to query DuckDuckGo and open items from search result with Web browser quickly
</div>

## Demo

<div align="center">

<img src="https://private-user-images.githubusercontent.com/62412884/266784515-4d5b3951-1e03-44f5-a79f-32d5cbc90ea2.gif?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTEiLCJleHAiOjE2OTQyNjc4NDMsIm5iZiI6MTY5NDI2NzU0MywicGF0aCI6Ii82MjQxMjg4NC8yNjY3ODQ1MTUtNGQ1YjM5NTEtMWUwMy00NGY1LWE3OWYtMzJkNWNiYzkwZWEyLmdpZj9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFJV05KWUFYNENTVkVINTNBJTJGMjAyMzA5MDklMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjMwOTA5VDEzNTIyM1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTA3MTk5MmU3YzBjYmQ4Y2QxMTY1YWU2NzcxZWYyZjhjODFlNjQ3MjlhNDI2MThmZTAzMmVmMmIyMTdlNzk1ZDImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.SPDMKpjYQHf2cbZqNgpvpzHiFm5N0cZ7t7j7aFWHiMw" width="70%" >

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

