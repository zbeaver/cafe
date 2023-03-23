package decoder

import (
	"log"
	"strings"

	"github.com/aymerick/douceur/inliner"
	"github.com/ryboe/q"
	"github.com/zbeaver/cafe/pkg/vui"
	"golang.org/x/net/html"
)

type Decoder interface {
	Decode(vui.Documentary, Template) (vui.Elementary, error)
}

type Template []byte

type decoder struct{}

func NewDecoder() Decoder {
	return &decoder{}
}

func (d *decoder) Decode(doc vui.Documentary, tpl Template) (docElm vui.Elementary, err error) {
	raw, err := inliner.Inline(string(tpl))
	dom, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		log.Fatal(err)
	}
	docElm, _ = doc.CreateElement(
		"html",
		vui.WithNodeName("html", vui.DocumentNode),
	)
	doc.AppendChild(docElm)
	err = d.decodeElement(docElm, dom)
	return
}

func (d *decoder) decodeElement(root vui.INode, n *html.Node) (err error) {
	if n.Type == html.DocumentNode {
		return d.decodeElement(root, n.FirstChild)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		// when dom is element
		// handle attribute nodes specific cases
		case html.ElementNode:
			if c.Data == "html" {
				return d.decodeElement(root, c)
			}

			opts := []interface{}{
				vui.WithNodeName(c.Data, vui.ElementNode),
			}

			for _, attr := range c.Attr {
				switch attr.Key {
				case "id":
					opts = append(opts, vui.WithId(attr.Val))
				case "class":
					classes := strings.Split(attr.Val, " ")
					opts = append(opts, vui.WithClass(classes...))
				case "style":
					opts = append(opts, vui.WithStyle(attr.Val))
				}
			}

			child, err := root.
				OwnerDocument().
				CreateElement(
					c.Data,
					opts...,
				)
			if err != nil {
				panic(err)
			}
			root.AppendChild(child)
			err = d.decodeElement(child, c)

		// when dom is text node
		// append text node to root
		case html.TextNode:
			child, err := root.OwnerDocument().CreateText(c.Data)
			if err != nil {
				panic(err)
			}
			root.AppendChild(child)
			err = d.decodeElement(child, c)

		// More
		default:
			q.Q("default", int(c.Type), c.Data)
		}
	}
	return
}
