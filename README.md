# webutil

`webutil` provides [Go](https://golang.org/) functions that operate at a level
better matching the level that I'm working at when I'm using JSON API's and
scraping the web.

## Features

* `GET` a `[]byte`. 404, etc. are an error.
* Same, but return an io.ReadCloser.
* Same, but parse with [`net/html`](https://godoc.org/golang.org/x/net/html)
    and return the html.Node.
    Useful for using [Cascadia](https://github.com/andybalholm/cascadia) to
    extract parts.
* Get text from an html.Node.
* Walk an html.Node.

## Usage

`go get sethwklein.net/go/webutil`

`import "sethwklein.net/go/webutil"`

http://godoc.org/sethwklein.net/go/webutil

## Rationale

I know better than to publish this package. Someone has said (I don't have
a link handy :( ) that packages with util in the name are almost always
a bad idea. And it makes no sense to use a package that pulls in an HTML
parser when all you need is a `[]byte`. But I can't seem to stop using it,
so I'm publishing it even if it should be a bad idea.
