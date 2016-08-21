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
			t.Errorf("Expected %d nodes for (%s) selector but found %d.", test.out, test.in, len(find))
		}
	}
}

func TestClosest(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(testHTML))

	item := Find(root, ".item")[0]
	_, ok := Closest(item, ".list")
	if !ok {
		t.Error("Expected list but nothing was found")
	}

	_, ok = Closest(item, ".notfound")
	if ok {
		t.Error("Expected empty node but something was found")
	}
}

func TestText(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(testHTML))
	link := Find(root, ".text a")

	text := Text(link[0])
	if text != "Text" {
		t.Error("Expected `Text` text in node but found", text)
	}
}
