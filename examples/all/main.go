package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/zbeaver/cafe/pkg/decoder"
	"github.com/zbeaver/cafe/pkg/render"
	"github.com/zbeaver/cafe/pkg/vui"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "elm.html", "link to layout file")
	flag.Parse()
	_, fpath, _, _ := runtime.Caller(0)
	path := filepath.Dir(fpath)
	tpl, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", path, file))
	if err != nil {
		panic(err)
	}
	doc := vui.NewDocument()
	de := decoder.NewDecoder()
	_, err = de.Decode(doc, tpl)
	if err != nil {
		panic(err)
	}
	engine := render.NewEngine(context.Background(), doc)
	fmt.Println(engine.Render())
}
