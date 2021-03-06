package scrape

import (
	"strings"

	"golang.org/x/net/html"
)

// Find returns all nodes which match selector.
func Find(node *html.Node, selector string) []*html.Node {
	selectors := parseSelector(selector)

	return findNodes(node, selectors)
}

// Closest searches up HTML tree from the current node until either a
// match is found or the top is hit.
func Closest(node *html.Node, selector string) (*html.Node, bool) {
	if node == nil {
		return nil, false
	}

	for p := node.Parent; p != nil; p = p.Parent {
		if matchSelector(p, selector) {
			return p, true
		}
	}

	return nil, false
}

// Text searches for text and concatenate all separated strings
func Text(node *html.Node) string {
	if node == nil {
		return ""
	}

	var texts []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			trimmed := strings.TrimSpace(string(c.Data))
			if trimmed != "" {
				texts = append(texts, trimmed)
			}
		}
	}

	return strings.Join(texts, " ")
}

// Attr returns the value of an HTML attribute.
func Attr(node *html.Node, key string) string {
	if node == nil {
		return ""
	}

	for _, a := range node.Attr {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}

func findNodes(node *html.Node, selectors []string) []*html.Node {
	matched := []*html.Node{}

	if node == nil {
		return matched
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		s, ok := matchSelectors(c, selectors)
		if ok && len(selectors) == 1 {
			matched = append(matched, c)
		}

		found := findNodes(c, s)
		if len(found) > 0 {
			matched = append(matched, found...)
		}
	}

	return matched
}

func checkTag(node *html.Node, tag string) bool {
	return (node.Data == tag) && (node.Type != html.TextNode)
}

func checkId(node *html.Node, id string) bool {
	return Attr(node, "id") == id
}

func checkClass(node *html.Node, class string) bool {
	classes := strings.Fields(Attr(node, "class"))
	for _, c := range classes {
		if c == class {
			return true
		}
	}

	return false
}

func matchSelectors(node *html.Node, selectors []string) ([]string, bool) {
	if len(selectors) == 0 {
		return nil, false
	}

	ok := matchSelector(node, selectors[0])
	if ok && len(selectors) > 1 {
		selectors = selectors[1:]
	}

	return selectors, ok
}

func matchSelector(node *html.Node, selector string) bool {
	// TODO: add handling complex selector like 'a.class#id'

	ok := false
	switch selector[0] { // check for first char
	case '.':
		ok = checkClass(node, selector[1:])
	case '#':
		ok = checkId(node, selector[1:])
	default:
		ok = checkTag(node, selector)
	}

	return ok
}

func parseSelector(selector string) []string {
	return strings.Fields(selector)
}
