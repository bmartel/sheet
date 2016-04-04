
## Sheet - A thin wrapper on html/template
###  Supports nested file globbing and a simple template inheritance

Example Usage (main.go of nethttp example reduced for brevity):

```go
  import (
  	"net/http"

  	"github.com/bmartel/sheet"
  )

  func main() {
  	sheet.TemplateExtension = ".html"

  	views := sheet.New()

  	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
      if r.URL.Path != "/" {
        h.notfound(w, r)
        return
      }

      data := map[string]interface{}{
        "title": "Example Index",
      }

      views.HTML(w, "index.html", data)
  	})

  	http.ListenAndServe(":8080", nil)
  }
```

By default, templates will be loaded using the following structure:

```bash
  |-- templates
    |-- includes
    |-- layouts
      |-- base.tmpl
    |-- views
```

The names of these directories are configurable (defaults shown commented to the right)

```go
  sheet.TemplatePath = "views" // templates

  sheet.TemplateExtension = ".html" // .tmpl

  sheet.DefaultLayout = "default" // base

  tmpl := sheet.NewFromDir(
    "pages", // views
    "templates", // layouts
    "blocks", // includes
  )
```

Root level views by default will look to load a `base` layout and any root level `includes`.

Nested views will take the parent most directory name under the views directory and try and
load that same name from `layouts` if it exists, defaulting back to `base` if not. It will also
attempt to load all root level `includes` for that same name.

Inheritance loosely follows the simple convention of a single directory nest under `views`.

Example:

```bash
|-- templates
  |-- includes
    |-- product
      |-- card.tmpl
    |-- sidebar.tmpl
  |-- layouts
    |-- base.tmpl
    |-- product.tmpl
  |-- views
    |-- product
      |-- list.tmpl
    |-- about.tmpl
    |-- index.tmpl
```

Compiling this file structure, the following will be cached for use


`key` => `value`

- `'about.tmpl' ` => template parsed using `['layouts/base.tmpl', 'views/about.tmpl', 'includes/sidebar.tmpl']`

- `'index.tmpl' ` => template parsed using `['layouts/base.tmpl', 'views/index.tmpl', 'includes/sidebar.tmpl']`

- `'product/list.tmpl' ` => template parsed using `['layouts/product.tmpl', 'views/product/list.tmpl', 'includes/sidebar.tmpl',
 'includes/product/card.tmpl']`


As you can see it just takes and loads the appropriate files in such an order that allows for a rudimentary inheritance.
