package main

import (
	"fmt"
	"strings"

	. "github.com/nbjahan/go-launchbar"
)

func newCategory(name, icon string) *Item {
	key := strings.ToLower(name)
	i := NewItem(name)
	i.SetIcon(iconPath(key, icon))
	i.SetActionRunsInBackground(false)
	i.SetActionReturnsItems(true)
	i.SetAction("unisym")
	i.SetRun(func(c *Context) Items {
		items := NewItems()
		for _, row := range chars[key] {
			i := NewItem(row[0])
			i.SetIcon(iconPath(key, row[0]))
			i.Run("showChar", row[0], row[1])
			i.SetAction("unisym")
			i.SetActionReturnsItems(true)
			i.SetActionRunsInBackground(false)
			items.Add(i)
		}
		return *items
	})
	return i
}

func setupViews() {
	var i *Item
	var items *Items
	v := pb.NewView("main")

	in := pb.Input.String()

	if !pb.Input.IsEmpty() {
		i = v.NewItem("Unicode Symbols: Search")
		i.SetIcon("/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/MagnifyingGlassIcon.icns")
		i.SetSubtitle(fmt.Sprintf("for: %q", in))
		i.SetActionRunsInBackground(false)
		i.SetActionReturnsItems(true)
		i.SetRun(func(c *Context) Items {
			items = search(pb.Input.String())
			return *items
		})
		return
	}

	v.AddItem(newCategory("Popular", "Atom-Symbol"))
	v.AddItem(newCategory("Classics", "Eighth Note"))
	v.AddItem(newCategory("Arrows", "Downwards-Zigzag-Arrow"))
	v.AddItem(newCategory("Math", "Surface Integral"))
	v.AddItem(newCategory("Numerals", "Vulgar Fraction Zero Thirds"))
	v.AddItem(newCategory("Shapes", "Square With Upper Right To Lower Left Fill"))
	v.AddItem(newCategory("Symbols", "Yin Yang"))
	v.AddItem(newCategory("Punctuations", "Heavy Double Comma Quotation Mark Ornament"))
	v.AddItem(newCategory("Currencies", "Dollar Sign"))
}
