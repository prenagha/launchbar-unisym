package main

import (
	"sort"

	. "github.com/nbjahan/go-launchbar"
	"github.com/nbjahan/gofuzz"
)

var scores []float64

type byScore Items

func (items byScore) Len() int           { return len(items) }
func (items byScore) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items byScore) Less(i, j int) bool { return scores[i] < scores[j] }

func search(q string) *Items {

	fuzz := gofuzz.New()
	fuzz.Threshold = 0.2
	items := NewItems()
	scores = make([]float64, 0)
	for cat, symbols := range chars {
		for _, row := range symbols {
			pos, score := fuzz.Search(row[0], q, 0)
			if pos != -1 {
				scores = append(scores, score)
				i := NewItem(row[0])
				i.SetIcon(iconPath(cat, row[0]))
				i.Run("showChar", row[0], row[1])
				i.SetAction("unisym")
				i.SetActionReturnsItems(true)
				i.SetActionRunsInBackground(false)
				items.Add(i)
			}
		}
	}
	sort.Sort(byScore(*items))
	return items
}
