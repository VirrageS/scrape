# scrape
A jquery like interface for Go website scrapping.

## Usage

```go
import (
    "fmt"
    "net/http"

    "golang.org/x/net/html"

    "github.com/VirrageS/scrape"
)

func main() {
    response, err := http.Get("https://github.com/trending")
    if err != nil {
        return
    }

    root, err := html.Parse(response.Body)
    if err != nil {
        return
    }

    repos := scrape.Find(root, ".repo-list-item")
    for _, repo := range repos {
        // get url
        link := scrape.Find(repo, ".repo-list-name a")[0]
        url := "https://github.com" + scrape.Attr(link, "href")

        // get name
        name := scrape.Text(link)

        fmt.Printf("[REPO] name: %s; url: %s\n", url, name)
    }
}
```
