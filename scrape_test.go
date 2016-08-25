package scrape

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const testHTML = `
<html>
  <body>
    <div class="list">
      <div class="container">
        <ul class="list">
         <li class="item">
           <span id="super"></span>
         </li>
         <li class="item2"></li>
        </ul>
        <div class="description">div</div>
        <a href="">Link</a>
      </div>
    </div>
    <a href="">Link 2</a>
    <div class="text"><a href="">  Text  <p>Text1</p></a>
  </body>
</html>
`

var findtests = []struct {
	in  string
	out int
}{
	{".list", 2},
	{".list a", 1},
	{".list .container ul #super", 1},
	{"div", 4},
	{"div div", 2},
}

func TestFind(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(testHTML))

	for _, test := range findtests {
		find := Find(root, test.in)
		if len(find) != test.out {
			t.Errorf("Expected %d nodes for '%s' selector but found %d.", test.out, test.in, len(find))
		}
	}
}

var closesttests = []struct {
	find string
	in   string
	out  bool
}{
	{".item", ".list", true},
	{".item", ".notfound", false},
}

func TestClosest(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(testHTML))

	for _, test := range closesttests {
		item := Find(root, test.find)[0]
		_, ok := Closest(item, test.in)
		if ok != test.out {
			t.Errorf("Expected %t when searching closest '%s' starting from '%s' but found %t", test.out, test.in, test.find, ok)
		}
	}
}

var texttests = []struct {
	in  string
	out string
}{
	{".text a", "Text"},
	{".text p", "Text1"},
}

func TestText(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(testHTML))

	for _, test := range texttests {
		item := Find(root, test.in)[0]
		if Text(item) != test.out {
			t.Errorf("Expected '%s' as text in node '%s' but found '%s'.", test.out, test.in, Text(item))
		}
	}
}

func TestNilNodes(t *testing.T) {
	attr := Attr(nil, "something")
	if attr != "" {
		t.Errorf("Expected empty string but found %s", attr)
	}

	find := Find(nil, ".selector")
	if len(find) != 0 {
		t.Errorf("Expected empty node array but found %v", find)
	}

	_, ok := Closest(nil, ".selector")
	if ok {
		t.Errorf("Expected nothing to be found in Closest but something was found")
	}

	text := Text(nil)
	if text != "" {
		t.Errorf("Expected empty string for Text but found %s", text)
	}
}
