package main

import (
	"fmt"
	"path"
	"strings"

	sjson "github.com/bitly/go-simplejson"
	. "github.com/nbjahan/go-launchbar"
)

var InDev string

var pb *Action
var chars = map[string][][2]string{}

func iconPath(f, s string) string {
	s = strings.Replace(s, " ", "-", -1) + "-Template.pdf"
	return path.Join(pb.ActionPath(), "Contents", "Resources", f, s)
}

func init() {
	pb = NewAction("Unicode Symbols", ConfigValues{
		"actionDefaultScript": "unisym",
	})
}

var funcs = FuncMap{
	"showChar": func(c *Context) Items {
		args := c.Input.FuncArgsMapString()
		items := NewItems()
		utf8 := fmt.Sprintf("%X", args[1])
		unicode := fmt.Sprintf("%X", []rune(args[1])[0])
		items.Add(NewItem(args[1]).SetSubtitle(args[0]).SetIcon("at.obdev.launchbar:Words"))
		items.Add(NewItem(unicode).SetSubtitle("unicode").SetIcon("at.obdev.launchbar:InfoTemplate"))
		items.Add(NewItem(utf8).SetSubtitle("utf8").SetIcon("at.obdev.launchbar:InfoTemplate"))
		return *items
	},
}

func main() {

	pb.Init(funcs)

	setupViews()

	if InDev != "" {
		pb.Logger.Printf("in:\n%s\n", pb.Input.Raw())
	}

	out := pb.Run()

	if InDev != "" {
		nice := out
		js, err := sjson.NewJson([]byte(out))
		if err == nil {
			b, err := js.EncodePretty()
			if err == nil {
				nice = string(b)
			}
		}
		pb.Logger.Println("out:", string(nice))
	}

	fmt.Println(out)
}
