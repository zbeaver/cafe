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
			<head>
				<style text="text/css">
					body {
						background-color: #993415;
						padding: 3;
					}
					#div1 {
						background-color: #ffffff;
						text-color: #000000;
						padding: 2;
					}
					#div2 {
						background-color: #004422;
						text-color: #000000;
						padding: 2;
					}
					#div3 {
						background-color: #004422;
						text-color: #000000;
						padding: 2;
					}
					.container {
						width: 100%;
					}
				</style>
			</head>
			<body>
			  <div class="container">
					<div id="div1">
						<div>sdfa world node</div>
					</div>
					<div id="div2">
						<div>ready node coucu</div>
					</div>
					<div id="div3">
						<div>ready node coucu</div>
					</div>
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
