package main

import (
	"fmt"
	"strings"

	. "github.com/nbjahan/go-launchbar"
)

func newCategory(name, icon string, keys ...string) *Item {
	if len(keys) == 0 {
		keys = []string{strings.ToLower(name)}
	}
	i := NewItem(name)
	i.SetActionRunsInBackground(false)
	i.SetActionReturnsItems(true)
	i.SetAction("unisym")
	rows := make([][]string, 0)
	for _, key := range keys {
		for _, char := range chars[key] {
			dup := false
			for _, row := range rows {
				if row[1] == char[1] {
					dup = true
					break
				}
			}
			if dup {
				continue
			}
			if char[1] == icon {
				i.SetIcon(iconPath(key, icon))
			}
			rows = append(rows, []string{char[0], char[1], key})
		}
	}

	i.SetRun(func(c *Context) Items {
		items := NewItems()
		for _, row := range rows {
			i := NewItem(row[0])
			i.SetSubtitle(fmt.Sprintf("%s %X", row[1], []rune(row[1])[0]))
			i.SetIcon(iconPath(row[2], row[1]))
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
	v := pb.NewView("main")

	in := pb.Input.String()

	if !pb.Input.IsEmpty() {
		i := v.NewItem("Unicode Symbols: Search for all")
		i.SetIcon("/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/MagnifyingGlassIcon.icns")
		i.SetSubtitle(fmt.Sprintf("for: %q", in))
		i.SetActionRunsInBackground(false)
		i.SetActionReturnsItems(true)
		i.SetRun(func(c *Context) Items {
			items := search(pb.Input.String(), -1)
			return *items
		})

		for _, item := range *search(in, 10) {
			v.AddItem(item)
		}
		return
	}

	// v.AddItem(newCategory("Common Indic Number Forms", "\U0000a830", "0a830"))
	// v.AddItem(newCategory("Rumi Numeral Symbols", "\U00010e60", "10e60"))

	v.AddItem(newCategory("Arrows", "↯", "arrows", "\U00002190", "02190", "02b00", "027f0", "02900", "1f800"))
	// v.AddItem(newCategory("Arrows", "\U00002190", "02190"))
	// v.AddItem(newCategory("Miscellaneous Symbols and Arrows", "\U00002b00", "02b00"))
	// v.AddItem(newCategory("Supplemental Arrows-A", "\U000027f0", "027f0"))
	// v.AddItem(newCategory("Supplemental Arrows-B", "\U00002900", "02900"))
	// v.AddItem(newCategory("Supplemental Arrows-C", "\U0001f800", "1f800"))

	v.AddItem(newCategory("Musical Symbols", "\U0001d11e", "1d100", "1d200", "1d000"))
	// v.AddItem(newCategory("Musical Symbols", "\U0001d11e", "1d100"))
	// v.AddItem(newCategory("Ancient Greek Musical Notation", "\U0001d200", "1d200"))
	// v.AddItem(newCategory("Byzantine Musical Symbols", "\U0001d000", "1d000"))

	// v.AddItem(newCategory("Math", "∯", "math", "1d400", "02200", "027c0", "02980", "02a00"))
	v.AddItem(newCategory("Mathematical Operators", "\U0000222F", "02200", "02a00"))
	// v.AddItem(newCategory("Mathematical Operators", "\U00002200", "02200"))
	// v.AddItem(newCategory("Supplemental Mathematical Operators", "\U00002a00", "02a00"))

	v.AddItem(newCategory("Mathematical Symbols", "\U000029B2", "027c0", "02980"))
	// v.AddItem(newCategory("Miscellaneous Mathematical Symbols-A", "\U000027c0", "027c0"))
	// v.AddItem(newCategory("Miscellaneous Mathematical Symbols-B", "\U00002980", "02980"))

	v.AddItem(newCategory("Mathematical Alphanumeric Symbols", "\U0001d400", "1d400"))

	v.AddItem(newCategory("Numerals", "↉", "numerals", "10100", "10140", "1d360", "1f100", "02460", "02150"))
	// v.AddItem(newCategory("Aegean Numbers", "\U00010100", "10100"))
	// v.AddItem(newCategory("Ancient Greek Numbers", "\U00010140", "10140"))
	// v.AddItem(newCategory("Counting Rod Numerals", "\U0001d360", "1d360"))
	// v.AddItem(newCategory("Enclosed Alphanumeric Supplement", "\U0001f198", "1f100"))
	// v.AddItem(newCategory("Enclosed Alphanumerics", "\U00002460", "02460"))
	// v.AddItem(newCategory("Number Forms", "\U00002150", "02150"))

	v.AddItem(newCategory("Currency Symbols", "\U0000ff04", "020a0"))

	v.AddItem(newCategory("Transport and Map Symbols", "\U0001F680", "1f680"))
	v.AddItem(newCategory("Emoticons", "\U0001F604", "1f600"))

	// v.AddItem(newCategory("Symbols", "☯","symbols"))

	v.AddItem(newCategory("Dingbats", "\U00002700", "02700", "1f650"))
	// v.AddItem(newCategory("Dingbats", "\U00002700", "02700"))
	// v.AddItem(newCategory("Ornamental Dingbats", "\U0001F650", "1f650"))

	v.AddItem(newCategory("Miscellaneous Technical", "\U000023F0", "02300"))
	v.AddItem(newCategory("Letterlike Symbols", "\U00002100", "02100"))
	v.AddItem(newCategory("Miscellaneous Symbols and Pictographs", "\U0001F326", "1f300"))
	v.AddItem(newCategory("Miscellaneous Symbols", "\u2600", "02600"))

	v.AddItem(newCategory("Block Elements", "\U00002592", "02580"))
	v.AddItem(newCategory("Box Drawing", "\U0000255D", "02500"))

	v.AddItem(newCategory("Geometric Shapes", "\U0001F79C", "025a0", "1f780"))
	// v.AddItem(newCategory("Geometric Shapes", "\U000025a0", "025a0"))
	// v.AddItem(newCategory("Geometric Shapes Extended", "\U0001F78b", "1f780"))

	v.AddItem(newCategory("Playing Cards", "\U0001f0a1", "1f0a0"))
	v.AddItem(newCategory("Domino Tiles", "\U0001f032", "1f030"))
	v.AddItem(newCategory("Mahjong Tiles", "\U0001f000", "1f000"))

	v.AddItem(newCategory("Braille Patterns", "\U00002800", "02800"))
	v.AddItem(newCategory("Tags", "\U000e0031", "e0000"))
	v.AddItem(newCategory("Alchemical Symbols", "\U0001F74a", "1f700"))
	v.AddItem(newCategory("Ancient Symbols", "\U00010190", "10190"))
	v.AddItem(newCategory("Combining Diacritical Marks for Symbols", "\U000020d0", "020d0"))
	v.AddItem(newCategory("Optical Character Recognition", "\U00002440", "02440"))
	v.AddItem(newCategory("Yijing Hexagram Symbols", "\U00004dc0", "04dc0"))
	v.AddItem(newCategory("Tai Xuan Jing Symbols", "\U0001d300", "1d300"))

	// v.AddItem(newCategory("Cuneiform Numbers and Punctuation", "\U00012400", "12400"))
	v.AddItem(newCategory("Punctuations", "❞", "punctuations", "02000", "02e00", "03000"))
	// v.AddItem(newCategory("General Punctuation", "\U00002049", "02000"))
	// v.AddItem(newCategory("Supplemental Punctuation", "\U00002e00", "02e00"))
	// v.AddItem(newCategory("CJK Symbols and Punctuation", "\U00003004", "03000"))

	v.AddItem(newCategory("Superscripts and Subscripts", "\U00002070", "02070"))

	// v.AddItem(newCategory("Enclosed CJK Letters and Months", "\U00003200", "03200"))
	// v.AddItem(newCategory("Small Form Variants", "\U0000fe50", "0fe50"))
	// v.AddItem(newCategory("CJK Compatibility Forms", "\U0000fe45", "0fe30"))
	// v.AddItem(newCategory("Halfwidth and Fullwidth Forms", "\U0000ff01", "0ff00"))
	// v.AddItem(newCategory("Vertical Forms", "\U0000fe10", "0fe10"))
	// v.AddItem(newCategory("Popular", "⚛","popular"))
	// v.AddItem(newCategory("Classics", "♪","classics"))
	// v.AddItem(newCategory("Shapes", "▨","shapes"))
	// v.AddItem(newCategory("Currencies", "$"),"currencies")
}
