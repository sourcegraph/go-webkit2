# go-webkit2

[go-webkit2](https://sourcegraph.com/github.com/sourcegraph/go-webkit2/readme)
provides [Go](http://golang.org) bindings for the [WebKitGTK+ 2
API](http://webkitgtk.org/reference/webkit2gtk/stable/index.html). It permits
headless operation of WebKit as well as embedding a WebView in a GTK+
application.

* [Documentation on Sourcegraph](https://sourcegraph.com/github.com/sourcegraph/go-webkit2/tree)

[![status](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-webkit2/badges/status.png)](https://sourcegraph.com/github.com/sourcegraph/go-webkit2)
[![xrefs](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-webkit2/badges/xrefs.png)](https://sourcegraph.com/github.com/sourcegraph/go-webkit2)
[![funcs](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-webkit2/badges/funcs.png)](https://sourcegraph.com/github.com/sourcegraph/go-webkit2)
[![top func](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-webkit2/badges/top-func.png)](https://sourcegraph.com/github.com/sourcegraph/go-webkit2)
[![library users](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-webkit2/badges/library-users.png)](https://sourcegraph.com/github.com/sourcegraph/go-webkit2)


## Requirements

* [Go](http://golang.org) >= 1.2rc1 (due to [#3250](https://code.google.com/p/go/issues/detail?id=3250))
* [GTK+](http://www.gtk.org) 3.10+
* [WebKitGTK+](http://webkitgtk.org/) >= 2.0.0

You can specify Go build tags to omit bindings in
[gotk3](https://github.com/conformal/gotk3) for later versions of GTK
(e.g., `go build -tags gtk_3_10`).

#### Ubuntu 14.04 (Trusty)
```bash
sudo apt-get install libwebkit2gtk-3.0-dev
```

Pass `-tags gtk_3_10` to the go tool if you have GTK 3.10 installed.

#### Ubuntu 13.10 (Saucy)
```bash
sudo add-apt-repository ppa:gnome3-team/gnome3-staging
sudo apt-get update
sudo apt-get install libwebkit2gtk-3.0-dev
```
#### Ubuntu 13.04 (Raring)
```bash
sudo add-apt-repository ppa:gnome3-team/gnome3
sudo apt-get update
sudo apt-get install libwebkit2gtk-3.0-dev
```
#### Arch Linux
```bash
sudo pacman -S webkitgtk
```

#### Other platforms

Make sure you install WebKitGTK+ 2, not version 1. After installation, you
should have an include file that satisfies `#include <webkit2/webkit2.h>`.


## Usage

### As a Go package

```go
package webkit2_test

import (
	"fmt"
	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
	"github.com/sqs/gojs"
	"runtime"
)

func Example() {
	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	webView.Connect("load-changed", func(ctx *glib.CallbackContext) {
		loadEvent := webkit2.LoadEvent(ctx.Arg(0).Int())
		switch loadEvent {
		case webkit2.LoadFinished:
			fmt.Println("Load finished.")
			fmt.Printf("Title: %q\n", webView.Title())
			fmt.Printf("URI: %s\n", webView.URI())
			webView.RunJavaScript("window.location.hostname", func(val *gojs.Value, err error) {
				if err != nil {
					fmt.Println("JavaScript error.")
				} else {
					fmt.Printf("Hostname (from JavaScript): %q\n", val)
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI("https://www.google.com/")
		return false
	})

	gtk.Main()

	// output:
	// Load finished.
	// Title: "Google"
	// URI: https://www.google.com/
	// Hostname (from JavaScript): "www.google.com"
}
```

See the
[documentation](https://sourcegraph.com/github.com/sourcegraph/go-webkit2) and
the test files for usage information and examples.

For more information about the underlying WebKitGTK+ 2 API, refer to the
[WebKit2 docs](http://webkitgtk.org/reference/webkit2gtk/stable/index.html).


### As a program for evaluating JavaScript in the context of a web page

The included `webkit-eval-js` program runs the contents of a JavaScript file in the context of
a web page. Run with:

```
$ go get -tags gtk_3_10 github.com/sourcegraph/go-webkit2/webkit-eval-js
$ webkit-eval-js https://example.com scriptfile.js
```

For example:

```
$ echo document.title | webkit-eval-js https://google.com /dev/stdin
"Google"
```


## Used in

The following projects use go-webkit2:

* [WebLoop](https://sourcegraph.com/github.com/sourcegraph/webloop) - headless WebKit with a Go API

See [go-webkit2
users](https://sourcegraph.com/github.com/sourcegraph/go-webkit2/.dependents)
for a full list of repositories and people using go-webkit2.


## Running tests

```
go test ./webkit2
```

Note: The tests require an X display. If you are not running in a graphical
environment, you can use [Xvfb](http://en.wikipedia.org/wiki/Xvfb):

```
Xvfb :1 &
export DISPLAY=:1
go test ./webkit2
```


## TODO

* Implement more of the WebKitGTK+ 2 API. Right now, only certain parts of it
  are implemented.
* [Set up CI testing.](https://github.com/sourcegraph/go-webkit2/issues/1) This
  is difficult because all of the popular CI services run older versions of
  Ubuntu that make it difficult to install WebKitGTK+ >= 2.0.0.
* Create example applications.
* Fix memory leaks where C strings are allocated and not freed.


## Contributors

See the AUTHORS file for a list of contributors.

Submit contributions via GitHub pull request. Patches should include tests and
should pass [golint](https://github.com/golang/lint).
