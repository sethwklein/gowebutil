package webutil

import (
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
	"sethwklein.net/go/errutil"
)

// GetReadCloser returns the response body as an io.ReadCloser. The caller is
// responsible for closing it. GetReadCloser returns an error on
// any non-200 response.
func GetReadCloser(url string) (io.ReadCloser, error) {
	var err error
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = resp.Body.Close()
		return nil, errutil.Append(err, NewHTTPError(resp))
	}
	return resp.Body, nil
}

// GetBytes returns a slice of bytes containing the response body. It returns
// an error on any non-200 response.
func GetBytes(url string) (buf []byte, err error) {
	body, err := GetReadCloser(url)
	if err != nil {
		return nil, err
	}
	defer errutil.AppendCall(&err, body.Close)
	return ioutil.ReadAll(body)
}

// GetHTML returns the parsed HTML response body as a (pointer to) a
// go.net/html.Node. It returns an error on any non-200 response.
func GetHTML(url string) (doc *html.Node, err error) {
	body, err := GetReadCloser(url)
	if err != nil {
		return nil, err
	}
	defer errutil.AppendCall(&err, body.Close)
	return html.Parse(body)
}
