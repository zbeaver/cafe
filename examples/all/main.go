package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/zbeaver/cafe/pkg/decoder"
	"github.com/zbeaver/cafe/pkg/render"
	"github.com/zbeaver/cafe/pkg/vui"
)

func main() {
	tpl, err := ioutil.ReadFile("./layout.html")
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
