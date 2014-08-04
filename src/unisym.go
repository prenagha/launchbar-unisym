package main

import (
	"fmt"
	"os/exec"
	"path"

	sjson "github.com/bitly/go-simplejson"
	. "github.com/nbjahan/go-launchbar"
)

var InDev string

var pb *Action
var chars = map[string][][2]string{}

func iconPath(f, s string) string {
	r := []rune(s)[0]
	s = fmt.Sprintf("%X", r) + "-Template.pdf"
	return path.Join(pb.ActionPath(), "Contents", "Resources", f, s)
}

func init() {
	pb = NewAction("Unicode Symbols", ConfigValues{
		"actionDefaultScript": "unisym",
	})
}

func pasteChar(s string) {
	exec.Command("osascript", "-e", fmt.Sprintf(`tell application "LaunchBar"
       perform action "Paste in Frontmost Application" with string "%s"
       end tell`, s)).Start()
}

var funcs = FuncMap{
	"pasteChar": func(c *Context) {
		args := c.Input.FuncArgsMapString()
		pasteChar(args[0])
	},
	"showChar": func(c *Context) Items {
		args := c.Input.FuncArgsMapString()
		if c.Action.IsShiftKey() || c.Action.IsControlKey() || c.Action.IsOptionKey() || c.Action.IsCommandKey() {
			pasteChar(args[1])
			return Items(nil)
		}
		items := NewItems()
		utf8 := fmt.Sprintf("%X", args[1])
		unicode := fmt.Sprintf("%X", []rune(args[1])[0])
		items.Add(NewItem(args[1]).SetSubtitle(args[0]).SetIcon("at.obdev.launchbar:Words").Run("pasteChar", args[1]).SetAction("unisym"))
		items.Add(NewItem(unicode).SetSubtitle("unicode").SetIcon("at.obdev.launchbar:InfoTemplate").Run("pasteChar", unicode).SetAction("unisym"))
		items.Add(NewItem(utf8).SetSubtitle("utf8").SetIcon("at.obdev.launchbar:InfoTemplate").Run("pasteChar", utf8).SetAction("unisym"))
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
