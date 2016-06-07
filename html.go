package webutil

import "golang.org/x/net/html"

// Walk calls f once for the node and each of it's descendants.
// This function probably belongs in go.net/html and may go away if it is added
// there.
func Walk(n *html.Node, f func(*html.Node)) {
	f(n)
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		Walk(child, f)
	}
}

// Text returns the text in the nodes.
func Text(n *html.Node) string {
	var buf []byte
	Walk(n, func(n *html.Node) {
		if n.Type == html.TextNode && len(n.Data) > 0 {
			buf = append(buf, n.Data...)
		}
	})
	return string(buf)
}
