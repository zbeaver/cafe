package main

import (
	"context"
	"fmt"

	"github.com/zbeaver/cafe/pkg/decoder"
	"github.com/zbeaver/cafe/pkg/render"
	"github.com/zbeaver/cafe/pkg/vui"
)

func main() {
	tpl := []byte(`
		<html>
			<body style="background-color: #993415; padding: 3">
			<div name="div1">
				xxx
				<div>world</div>
			</div>
			<div name="div2">
				hehe
				<div> coucu</div>
			</div>
			</body>
		</html>
	`)
	doc := vui.NewDocument()
	de := decoder.NewDecoder()
	_, err := de.Decode(doc, tpl)
	if err != nil {
		panic(err)
	}
	engine := render.NewEngine(context.Background(), doc)
	fmt.Println(engine.Render())
}
