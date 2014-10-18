package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"runtime"

	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
	"github.com/sqs/gojs"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "webkit-eval-js evaluates a JavaScript expression in the context of a web page\n")
		fmt.Fprintf(os.Stderr, "running in a headless instance of the WebKit browser.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\twebkit-eval-js url script-file\n\n")
		fmt.Fprintf(os.Stderr, "url is the web page to execute the script in, and script-file is a local file\n")
		fmt.Fprintf(os.Stderr, "with the JavaScript you want to evaluate. The result is printed to stdout as JSON.\n\n")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "Example usage:\n\n")
		fmt.Fprintf(os.Stderr, "\tTo return the value of `document.title` on https://google.com:\n")
		fmt.Fprintf(os.Stderr, "\t    $ echo document.title | webkit-eval-js https://google.com /dev/stdin\n")
		fmt.Fprintf(os.Stderr, "\tPrints:\n")
		fmt.Fprintf(os.Stderr, "\t    \"Google\"\n\n")
		fmt.Fprintf(os.Stderr, "Notes:\n\n")
		fmt.Fprintf(os.Stderr, "\tBecause a headless WebKit instance is used, your $DISPLAY must be set. Use\n")
		fmt.Fprintf(os.Stderr, "\tXvfb if you are running on a machine without an existing X server. See\n")
		fmt.Fprintf(os.Stderr, "\thttps://sourcegraph.com/github.com/sourcegraph/go-webkit2/readme for more info.\n")
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	log := log.New(os.Stderr, "", 0)

	pageURL := flag.Arg(0)
	scriptFile := flag.Arg(1)

	if _, err := url.Parse(pageURL); err != nil {
		log.Fatalf("Failed to parse URL %q: %s", pageURL, err)
	}

	script, err := ioutil.ReadFile(scriptFile)
	if err != nil {
		log.Fatalf("Failed to open script file %q: %s", scriptFile, err)
	}

	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent webkit2.LoadEvent) {
		switch loadEvent {
		case webkit2.LoadFinished:
			webView.RunJavaScript(string(script), func(val *gojs.Value, err error) {
				if err != nil {
					log.Fatalf("JavaScript error: %s", err)
				} else {
					json, err := val.JSON()
					if err != nil {
						log.Fatal("JavaScript serialization error: %s", err)
					}
					fmt.Println(string(json))
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI(pageURL)
		return false
	})

	gtk.Main()
}
