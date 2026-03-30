# whiskeyfyi

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/whiskeyfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/whiskeyfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Whiskey encyclopedia with expressions, distilleries, and regions — API client for [whiskeyfyi.com](https://whiskeyfyi.com).

> **Try the interactive tools at [whiskeyfyi.com](https://whiskeyfyi.com)**

## Install

`go get github.com/fyipedia/whiskeyfyi-go`

## Quick Start

```go
package main

import (
    "fmt"
    "fyipedia/whiskeyfyi-go"
)

func main() {
    client := whiskeyfyi.NewClient()
    result, err := client.Search("glenfiddich-12")
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Total, "results")
}
```

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install whiskeyfyi` | [PyPI](https://pypi.org/project/whiskeyfyi/) |
| **npm** | `npm install whiskeyfyi` | [npm](https://www.npmjs.com/package/whiskeyfyi) |
| **Go** | `go get github.com/fyipedia/whiskeyfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/whiskeyfyi-go) |
| **Rust** | `cargo add whiskeyfyi` | [crates.io](https://crates.io/crates/whiskeyfyi) |
| **Ruby** | `gem install whiskeyfyi` | [rubygems](https://rubygems.org/gems/whiskeyfyi) |

## Embed Widget

Embed [WhiskeyFYI](https://whiskeyfyi.com) widgets on any website with [whiskeyfyi-embed](https://widget.whiskeyfyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/whiskeyfyi-embed@1/dist/embed.min.js"></script>
<div data-whiskeyfyi="entity" data-slug="glenfiddich-12"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.whiskeyfyi.com)

## Links

- [WhiskeyFYI](https://whiskeyfyi.com) — Main site
- [API Documentation](https://whiskeyfyi.com/developers/)
- [OpenAPI Spec](https://whiskeyfyi.com/api/openapi.json)
- [Glossary](https://whiskeyfyi.com/glossary/)

## License

MIT
