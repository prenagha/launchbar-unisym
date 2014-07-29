package main

import . "github.com/nbjahan/go-launchbar"

func addCurrencies() {
	var i *Item
	v := pb.GetView("main")

	chars := [][2]string{
		{"Austral Sign", "₳"},
		{"Thai Currency Symbol Baht", "฿"},
		{"Fullwidth Cent Sign", "￠"},
		{"Colon Sign", "₡"},
		{"Cent Sign", "¢"},
		{"Cruzeiro Sign", "₢"},
		{"Cedi Sign", "₵"},
		{"Dong Sign", "₫"},
		{"Euro Sign", "€"},
		{"Fullwidth Pound Sign", "￡"},
		{"Pound Sign", "£"},
		{"Lira Sign", "₤"},
		{"French Franc Sign", "₣"},
		{"Latin Small Letter F With Hook", "ƒ"},
		{"Guarani Sign", "₲"},
		{"Kip Sign", "₭"},
		{"Mill Sign", "₥"},
		{"Naira Sign", "₦"},
		{"Peso Sign", "₱"},
		{"Fullwidth Dollar Sign", "＄"},
		{"Dollar Sign", "$"},
		{"Tugrik Sign", "₮"},
		{"Won Sign", "₩"},
		{"Fullwidth Won Sign", "￦"},
		{"Yen Sign", "¥"},
		{"Fullwidth Yen Sign", "￥"},
		{"Hryvnia Sign", "₴"},
		{"Currency Sign", "¤"},
		{"German Penny Symbol", "₰"},
		{"Khmer Currency Symbol Riel", "៛"},
		{"New Sheqel Sign", "₪"},
		{"Drachma Sign", "₯"},
		{"Euro Currency Sign", "₠"},
		{"Peseta Sign", "₧"},
		{"Rupee Sign", "₨"},
		{"Tamil Rupee Sign", "௹"},
		{"Rial Sign", "﷼"},
		{"Square Yuan", "㍐"},
		{"Bengali Rupee Mark", "৲"},
		{"Bengali Rupee Sign", "৳"},
		{"Indian Rupee Sign", "₹"},
	}
	i = v.NewItem("Currencies")
	i.SetIcon(iconPath("currencies", "Dollar Sign"))
	i.SetActionRunsInBackground(false)
	i.SetActionReturnsItems(true)
	i.SetRun(func(c *Context) Items {
		items := NewItems()
		for _, row := range chars {
			i := NewItem(row[0])
			i.SetIcon(iconPath("currencies", row[0]))
			i.Run("showChar", row[0], row[1])
			i.SetAction("unisym")
			i.SetActionReturnsItems(true)
			i.SetActionRunsInBackground(false)
			items.Add(i)
		}
		return *items
	})
}
