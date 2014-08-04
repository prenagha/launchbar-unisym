package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/nbjahan/go-launchbar"
	"github.com/nbjahan/gofuzz"
)

var scores []float64

type byScore Items

func (items byScore) Len() int      { return len(items) }
func (items byScore) Swap(i, j int) { items[i], items[j] = items[j], items[i] }
func (items byScore) Less(i, j int) bool {
	if scores[i] == scores[j] {
		return items[i].Item().Title < items[j].Item().Title
	}
	return scores[i] < scores[j]
}

func search(q string, n int) *Items {
	items := NewItems()
	q = strings.TrimSpace(q)
	if len(q) < 3 {
		return items
	}
	names := make(map[string]struct{})
	fuzz := gofuzz.New()
	fuzz.Threshold = 0.2
	scores = make([]float64, 0)
LOOP:
	for cat, symbols := range chars {
		for _, row := range symbols {
			pos, score := fuzz.Search(row[0], q, 0)
			if pos != -1 {
				if _, found := names[row[0]]; found {
					continue
				}
				names[row[0]] = struct{}{}
				scores = append(scores, score)
				i := NewItem(row[0])
				i.SetSubtitle(fmt.Sprintf("%s %X", row[1], []rune(row[1])[0]))
				i.SetIcon(iconPath(cat, row[1]))
				i.Run("showChar", row[0], row[1])
				i.SetAction("unisym")
				i.SetActionReturnsItems(true)
				i.SetActionRunsInBackground(false)
				items.Add(i)
				n -= 1
				if n == 0 {
					break LOOP
				}
			}
		}
	}
	sort.Sort(byScore(*items))
	return items
}
