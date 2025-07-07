# Web Crawler

A concurrent web crawler built in Go that discovers and analyzes internal links on websites.

## Features

- **Concurrent crawling**: Configurable number of concurrent goroutines
- **Link discovery**: Extracts and follows internal links from HTML pages
- **URL normalization**: Standardizes URLs for consistent tracking
- **Duplicate detection**: Tracks visited pages to avoid infinite loops
- **Configurable limits**: Set maximum pages to crawl
- **Detailed reporting**: Sorted report showing link counts per page

## Usage

```bash
go build -o crawler
./crawler <URL> <maxConcurrency> <maxPages>
```

### Parameters

- `URL`: The starting URL to crawl (e.g., `https://example.com`)
- `maxConcurrency`: Maximum number of concurrent goroutines (e.g., `3`)
- `maxPages`: Maximum number of pages to crawl (e.g., `10`)

### Example

```bash
./crawler https://wagslane.dev 3 5
```

## Output

The crawler provides real-time progress updates and a final report:

```
starting crawl of: https://wagslane.dev
crawling https://wagslane.dev
crawling https://wagslane.dev/posts/example/
crawling https://wagslane.dev/about/
...
=============================
  REPORT for https://wagslane.dev
=============================
Found 3 internal links to wagslane.dev
Found 2 internal links to wagslane.dev/posts/example
Found 1 internal links to wagslane.dev/about
```

## Architecture

- **main.go**: Entry point and command-line argument parsing
- **crawl_page.go**: Core crawling logic with concurrency control
- **get_html.go**: HTTP client for fetching web pages
- **get_urls_from_html.go**: HTML parsing and link extraction
- **normalize_url.go**: URL normalization and standardization
- **configue.go**: Configuration management and shared state

## Testing

Run the test suite:

```bash
go test -v
```

## Requirements

- Go 1.24.4 or later
- Internet connection for crawling external websites
