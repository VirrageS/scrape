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
        <div class="description"></div>
        <a href="">Link</a>
      </div>
    </div>
    <a href="">Link 2</a>
    <div class="text"><a href="">Text<p>Text1</p></a>
  </body>
</html>
`

func TestFindNestedNodes(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	lists := Find(node, ".list")

	if len(lists) != 2 {
		t.Error("Expected 2 nodes returned but only found", len(lists))
	}
}

func TestClosest(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	items := Find(node, ".item")
	if len(items) != 1 {
		t.Error("Expected 1 node of `item` class but found", len(items))
	}

	item := items[0]
	_, ok := Closest(item, ".list")
	if !ok {
		t.Error("Expected list but nothing was found")
	}

	_, ok = Closest(item, ".notfound")
	if ok {
		t.Error("Expected empty node but something was found")
	}
}

func TestFindComplexSelector(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	listAchnors := Find(node, ".list a")

	if len(listAchnors) != 1 {
		t.Error("Expected 1 node returned but found", len(listAchnors))
	}

	items := Find(node, ".list .container ul #super")
	if len(items) != 1 {
		t.Error("Expected 1 node returned but found", len(items))
	}
}

func TestText(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	link := Find(node, ".text a")

	text := Text(link[0])
	if text != "Text" {
		t.Error("Expected `Text` text in node but found", text)
	}
}
