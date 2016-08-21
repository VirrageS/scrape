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
         <li class="item"></li>
         <li class="item2"></li>
        </ul>
        <div class="description"></div>
        <a href="">Link</a>
      </div>
    </div>
    <a href="">Link 2</a>
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

	items := Find(node, ".list .container .list .item")
	if len(items) != 1 {
		t.Error("Expected 1 node returned but found", len(items))
	}
}
