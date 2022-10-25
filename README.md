A static blog written in Go, packaged in one binary.

Run it with

	$ go run ./cmd/blog

Or see it in production

- en: https://metnews.co/blog/
- zh: https://metword.co/blog/

### Features

- Rendered from markdown with metadata.
- Packaged in one binary using Go [embed] package.
- Static file hosting using Go [fs] package.
- Best practices for base template layout using Go [html/template] package.
- RSS feed supported.

[embed]: https://pkg.go.dev/embed
[fs]: https://pkg.go.dev/io/fs
[html/template]: https://pkg.go.dev/html/template#URL