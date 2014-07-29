package main

import . "github.com/nbjahan/go-launchbar"

func addShapes() {
	var i *Item
	v := pb.GetView("main")

	chars := [][2]string{
		{"Black Up Pointing Triangle", "▲"},
		{"Black Down Pointing Triangle", "▼"},
		{"Black Left Pointing Triangle", "◀"},
		{"Black Right Pointing Triangle", "▶"},
		{"Black Lower Right Triangle", "◢"},
		{"Black Lower Left Triangle", "◣"},
		{"Black Upper Right Triangle", "◥"},
		{"Black Upper Left Triangle", "◤"},
		{"White Up Pointing Triangle", "△"},
		{"White Down Pointing Triangle", "▽"},
		{"Lower Right Triangle", "◿"},
		{"Lower Left Triangle", "◺"},
		{"Upper Right Triangle", "◹"},
		{"Upper Left Triangle", "◸"},
		{"Black Up Pointing Small Triangle", "▴"},
		{"Black Down Pointing Small Triangle", "▾"},
		{"Black Left Pointing Small Triangle", "◂"},
		{"Black Right Pointing Small Triangle", "▸"},
		{"White Up Pointing Small Triangle", "▵"},
		{"White Down Pointing Small Triangle", "▿"},
		{"White Left Pointing Small Triangle", "◃"},
		{"White Right Pointing Small Triangle", "▹"},
		{"White Left Pointing Triangle", "◁"},
		{"White Right Pointing Triangle", "▷"},
		{"White Left Pointing Pointer", "◅"},
		{"White Right Pointing Pointer", "▻"},
		{"White Up Pointing Triangle With Dot", "◬"},
		{"White Triangle Containing Small White Triangle", "⟁"},
		{"Triangle With Underbar", "⧋"},
		{"Triangle With Dot Above", "⧊"},
		{"Right Triangle", "⊿"},
		{"Increment", "∆"},
		{"Nabla", "∇"},
		{"Up Pointing Triangle With Left Half Black", "◭"},
		{"Up Pointing Triangle With Right Half Black", "◮"},
		{"Down Pointing Triangle With Right Half Black", "⧩"},
		{"Down Pointing Triangle With Left Half Black", "⧨"},
		{"Sector", "⌔"},
		{"White Diamond With Centred Dot", "⟐"},
		{"White Diamond", "◇"},
		{"Black Diamond", "◆"},
		{"White Diamond Containing Black Small Diamond", "◈"},
		{"Diamond With Left Half Black", "⬖"},
		{"Diamond With Right Half Black", "⬗"},
		{"Diamond With Top Half Black", "⬘"},
		{"Diamond With Bottom Half Black", "⬙"},
		{"White Pentagon", "⬠"},
		{"White Hexagon", "⬡"},
		{"Software Function Symbol", "⎔"},
		{"Diamond Operator", "⋄"},
		{"Lozenge", "◊"},
		{"Black Lozenge", "⧫"},
		{"Black Hexagon", "⬢"},
		{"Horizontal Black Hexagon", "⬣"},
		{"Black Parallelogram", "▰"},
		{"Black Small Square", "▪"},
		{"Black Medium Square", "◼"},
		{"Black Vertical Rectangle", "▮"},
		{"Black Medium Small Square", "◾"},
		{"Quadrant Lower Right", "▗"},
		{"Quadrant Lower Left", "▖"},
		{"Black Square", "■"},
		{"End Of Proof", "∎"},
		{"Lower Three Eighths Block", "▃"},
		{"Lower Half Block", "▄"},
		{"Lower Five Eighths Block", "▅"},
		{"Lower Three Quarters Block", "▆"},
		{"Lower Seven Eighths Block", "▇"},
		{"Full Block", "█"},
		{"Left Half Block", "▌"},
		{"Right Half Block", "▐"},
		{"Left Three Eighths Block", "▍"},
		{"Left One Quarter Block", "▎"},
		{"Left Seven Eighths Block", "▉"},
		{"Left Three Quarters Block", "▊"},
		{"Left Five Eighths Block", "▋"},
		{"Light Vertical Bar", "❘"},
		{"Medium Vertical Bar", "❙"},
		{"Heavy Vertical Bar", "❚"},
		{"Upper Half Block", "▀"},
		{"Quadrant Upper Left", "▘"},
		{"Quadrant Upper Right", "▝"},
		{"Quadrant Upper Left And Lower Left And Lower Right", "▙"},
		{"Quadrant Upper Left And Lower Right", "▚"},
		{"Quadrant Upper Left And Upper Right And Lower Left", "▛"},
		{"Quadrant Upper Left And Upper Right And Lower Right", "▜"},
		{"Quadrant Upper Right And Lower Left And Lower Right", "▟"},
		{"Quadrant Upper Right And Lower Left", "▞"},
		{"Light Shade", "░"},
		{"Medium Shade", "▒"},
		{"Dark Shade", "▓"},
		{"Lower One Quarter Block", "▂"},
		{"Lower One Eighth Block", "▁"},
		{"Black Rectangle", "▬"},
		{"Upper One Eighth Block", "▔"},
		{"White Small Square", "▫"},
		{"White Vertical Rectangle", "▯"},
		{"White Rectangle", "▭"},
		{"White Parallelogram", "▱"},
		{"White Medium Small Square", "◽"},
		{"White Square", "□"},
		{"White Medium Square", "◻"},
		{"White Square With Rounded Corners", "▢"},
		{"Squared Plus", "⊞"},
		{"Squared Dot Operator", "⊡"},
		{"Squared Minus", "⊟"},
		{"Squared Times", "⊠"},
		{"White Square Containing Black Small Square", "▣"},
		{"Square With Horizontal Fill", "▤"},
		{"Square With Vertical Fill", "▥"},
		{"Square With Orthogonal Crosshatch Fill", "▦"},
		{"Dotted Square", "⬚"},
		{"Square With Upper Left To Lower Right Fill", "▧"},
		{"Square With Upper Right To Lower Left Fill", "▨"},
		{"Square With Diagonal Crosshatch Fill", "▩"},
		{"Square With Bottom Half Black", "⬓"},
		{"Square With Left Half Black", "◧"},
		{"Square With Top Half Black", "⬒"},
		{"Square With Right Half Black", "◨"},
		{"Square With Upper Left Diagonal Half Black", "◩"},
		{"Square With Lower Right Diagonal Half Black", "◪"},
		{"Square With Upper Right Diagonal Half Black", "⬔"},
		{"Square With Lower Left Diagonal Half Black", "⬕"},
		{"Lower Right Drop Shadowed White Square", "❏"},
		{"Upper Right Drop Shadowed White Square", "❐"},
		{"Lower Right Shadowed White Square", "❑"},
		{"Upper Right Shadowed White Square", "❒"},
		{"Squared Square", "⧈"},
		{"White Square With Upper Left Quadrant", "◰"},
		{"White Square With Lower Left Quadrant", "◱"},
		{"White Square With Upper Right Quadrant", "◳"},
		{"White Square With Lower Right Quadrant", "◲"},
		{"White Square With Vertical Bisecting Line", "◫"},
		{"Squared Small Circle", "⧇"},
		{"Squared Falling Diagonal Slash", "⧅"},
		{"Squared Rising Diagonal Slash", "⧄"},
		{"Apl Functional Symbol Quad Slash", "⍁"},
		{"Apl Functional Symbol Quad Backslash", "⍂"},
		{"White Concave Sided Diamond", "⟡"},
		{"Two Joined Squares", "⧉"},
		{"Medium Small White Circle", "⚬"},
		{"White Circle", "○"},
		{"Medium White Circle", "⚪"},
		{"Dotted Circle", "◌"},
		{"Circle With Vertical Fill", "◍"},
		{"Bullseye", "◎"},
		{"Large Circle", "◯"},
		{"Shadowed White Circle", "❍"},
		{"Fisheye", "◉"},
		{"Circled White Bullet", "⦾"},
		{"Circled Dot Operator", "⊙"},
		{"Circled Bullet", "⦿"},
		{"Circled Equals", "⊜"},
		{"Circled Minus", "⊖"},
		{"Circled Division Slash", "⊘"},
		{"Circled Ring Operator", "⊚"},
		{"Circled Asterisk Operator", "⊛"},
		{"Circled Dash", "⊝"},
		{"Black Circle", "●"},
		{"Medium Black Circle", "⚫"},
		{"Z Notation Spot", "⦁"},
		{"Circle With Left Half Black", "◐"},
		{"Circle With Right Half Black", "◑"},
		{"Circle With Lower Half Black", "◒"},
		{"Circle With Upper Half Black", "◓"},
		{"Circle With Upper Right Quadrant Black", "◔"},
		{"Circle With All But Upper Left Quadrant Black", "◕"},
		{"Circled Vertical Bar", "⦶"},
		{"Circled Reverse Solidus", "⦸"},
		{"White Circle With Lower Left Quadrant", "◵"},
		{"White Circle With Upper Left Quadrant", "◴"},
		{"White Circle With Lower Right Quadrant", "◶"},
		{"White Circle With Upper Right Quadrant", "◷"},
		{"Circled Plus", "⊕"},
		{"Circled Times", "⊗"},
		{"Z Notation Left Image Bracket", "⦇"},
		{"Z Notation Right Image Bracket", "⦈"},
		{"Z Notation Left Binding Bracket", "⦉"},
		{"Z Notation Right Binding Bracket", "⦊"},
		{"Medium Left Parenthesis Ornament", "❨"},
		{"Medium Right Parenthesis Ornament", "❩"},
		{"Left Double Parenthesis", "⸨"},
		{"Right Double Parenthesis", "⸩"},
		{"Left Half Black Circle", "◖"},
		{"Right Half Black Circle", "◗"},
		{"Medium Flattened Left Parenthesis Ornament", "❪"},
		{"Medium Flattened Right Parenthesis Ornament", "❫"},
		{"Heavy Left Pointing Angle Quotation Mark Ornament", "❮"},
		{"Heavy Right Pointing Angle Quotation Mark Ornament", "❯"},
		{"Medium Left Pointing Angle Bracket Ornament", "❬"},
		{"Medium Right Pointing Angle Bracket Ornament", "❭"},
		{"Heavy Left Pointing Angle Bracket Ornament", "❰"},
		{"Heavy Right Pointing Angle Bracket Ornament", "❱"},
		{"Square Image Of", "⊏"},
		{"Square Original Of", "⊐"},
		{"Square Image Of Or Equal To", "⊑"},
		{"Square Original Of Or Equal To", "⊒"},
		{"Inverse Bullet", "◘"},
		{"Inverse White Circle", "◙"},
		{"Upper Half Inverse White Circle", "◚"},
		{"Lower Half Inverse White Circle", "◛"},
		{"Upper Left Quadrant Circular Arc", "◜"},
		{"Upper Right Quadrant Circular Arc", "◝"},
		{"Lower Right Quadrant Circular Arc", "◞"},
		{"Lower Left Quadrant Circular Arc", "◟"},
		{"Upper Half Circle", "◠"},
		{"Lower Half Circle", "◡"},
		{"Double Intersection", "⋒"},
		{"Double Union", "⋓"},
		{"Double Subset", "⋐"},
		{"Double Superset", "⋑"},
		{"Right Double Arrow With Rounded Head", "⥰"},
		{"Box Drawings Light Arc Up And Right", "╰"},
		{"Box Drawings Light Arc Down And Left", "╮"},
		{"Box Drawings Light Arc Down And Right", "╭"},
		{"Box Drawings Light Arc Up And Left", "╯"},
		{"Arc", "⌒"},
		{"Down Fish Tail", "⥿"},
		{"Up Fish Tail", "⥾"},
		{"Right Fish Tail", "⥽"},
		{"Left Fish Tail", "⥼"},
		{"Left Barb Up Right Barb Down Harpoon", "⥊"},
		{"Left Barb Down Right Barb Up Harpoon", "⥋"},
		{"Up Barb Right Down Barb Left Harpoon", "⥌"},
		{"Up Barb Left Down Barb Right Harpoon", "⥍"},
		{"Left Barb Up Right Barb Up Harpoon", "⥎"},
		{"Left Barb Down Right Barb Down Harpoon", "⥐"},
		{"Up Barb Left Down Barb Left Harpoon", "⥑"},
		{"Up Barb Right Down Barb Right Harpoon", "⥏"},
		{"Box Drawings Light Diagonal Cross", "╳"},
		{"Multiplication X", "✕"},
		{"Rising Diagonal Crossing Falling Diagonal", "⤫"},
		{"Falling Diagonal Crossing Rising Diagonal", "⤬"},
		{"Box Drawings Light Diagonal Upper Right To Lower Left", "╱"},
		{"Box Drawings Light Diagonal Upper Left To Lower Right", "╲"},
		{"Big Solidus", "⧸"},
		{"Big Reverse Solidus", "⧹"},
		{"Segment", "⌓"},
		{"White Bullet", "◦"},
		{"Black Diamond Minus White X", "❖"},
		{"Heavy Multiplication X", "✖"},
		{"Heavy Greek Cross", "✚"},
		{"Heavy Open Centre Cross", "✜"},
		{"Black Bowtie", "⧓"},
		{"Black Hourglass", "⧗"},
		{"Bowtie With Left Half Black", "⧑"},
		{"Bowtie With Right Half Black", "⧒"},
		{"White Hourglass", "⧖"},
		{"Low Line", "_"},
		{"Monogram For Yang", "⚊"},
		{"Box Drawings Light Left", "╴"},
		{"Box Drawings Light Left And Heavy Right", "╼"},
		{"Box Drawings Heavy Left And Light Right", "╾"},
		{"Hyphen", "‐"},
		{"Hyphen Bullet", "⁃"},
		{"Non Breaking Hyphen", "‑"},
		{"Figure Dash", "‒"},
		{"Hyphen Minus", "-"},
		{"En Dash", "–"},
		{"Horizontal Line Extension", "⎯"},
		{"Em Dash", "—"},
		{"Horizontal Bar", "―"},
		{"Box Drawings Light Right", "╶"},
		{"Box Drawings Heavy Right", "╺"},
		{"Box Drawings Heavy Left", "╸"},
		{"Box Drawings Light Horizontal", "─"},
		{"Box Drawings Heavy Horizontal", "━"},
		{"Box Drawings Light Triple Dash Horizontal", "┄"},
		{"Box Drawings Heavy Triple Dash Horizontal", "┅"},
		{"Box Drawings Light Quadruple Dash Horizontal", "┈"},
		{"Box Drawings Heavy Quadruple Dash Horizontal", "┉"},
		{"Box Drawings Light Double Dash Horizontal", "╌"},
		{"Box Drawings Heavy Double Dash Horizontal", "╍"},
		{"Box Drawings Double Horizontal", "═"},
		{"Strictly Equivalent To", "≣"},
		{"Identical To", "≡"},
		{"Trigram For Heaven", "☰"},
		{"Trigram For Lake", "☱"},
		{"Trigram For Fire", "☲"},
		{"Trigram For Thunder", "☳"},
		{"Trigram For Wind", "☴"},
		{"Trigram For Water", "☵"},
		{"Trigram For Mountain", "☶"},
		{"Trigram For Earth", "☷"},
		{"Box Drawings Light Up", "╵"},
		{"Box Drawings Light Down", "╷"},
		{"Box Drawings Heavy Up", "╹"},
		{"Box Drawings Heavy Down", "╻"},
		{"Box Drawings Light Vertical", "│"},
		{"Right One Eighth Block", "▕"},
		{"Left One Eighth Block", "▏"},
		{"Box Drawings Heavy Vertical", "┃"},
		{"Box Drawings Light Triple Dash Vertical", "┆"},
		{"Box Drawings Heavy Triple Dash Vertical", "┇"},
		{"Box Drawings Light Quadruple Dash Vertical", "┊"},
		{"Box Drawings Light Double Dash Vertical", "╎"},
		{"Box Drawings Heavy Quadruple Dash Vertical", "┋"},
		{"Box Drawings Heavy Up And Light Down", "╿"},
		{"Box Drawings Light Up And Heavy Down", "╽"},
		{"Bottom Left Corner", "⌞"},
		{"Bottom Right Corner", "⌟"},
		{"Top Left Corner", "⌜"},
		{"Top Right Corner", "⌝"},
		{"Left Floor", "⌊"},
		{"Right Floor", "⌋"},
		{"Left Ceiling", "⌈"},
		{"Right Ceiling", "⌉"},
		{"Box Drawings Light Down And Right", "┌"},
		{"Box Drawings Down Light And Right Heavy", "┍"},
		{"Box Drawings Down Heavy And Right Light", "┎"},
		{"Box Drawings Heavy Down And Right", "┏"},
		{"Box Drawings Light Down And Left", "┐"},
		{"Box Drawings Down Light And Left Heavy", "┑"},
		{"Box Drawings Down Heavy And Left Light", "┒"},
		{"Box Drawings Heavy Down And Left", "┓"},
		{"Box Drawings Light Up And Right", "└"},
		{"Box Drawings Up Light And Right Heavy", "┕"},
		{"Box Drawings Up Heavy And Right Light", "┖"},
		{"Box Drawings Heavy Up And Right", "┗"},
		{"Box Drawings Light Up And Left", "┘"},
		{"Box Drawings Up Light And Left Heavy", "┙"},
		{"Box Drawings Up Heavy And Left Light", "┚"},
		{"Box Drawings Heavy Up And Left", "┛"},
		{"Box Drawings Light Vertical And Right", "├"},
		{"Box Drawings Vertical Light And Right Heavy", "┝"},
		{"Box Drawings Up Heavy And Right Down Light", "┞"},
		{"Box Drawings Down Heavy And Right Up Light", "┟"},
		{"Box Drawings Vertical Heavy And Right Light", "┠"},
		{"Box Drawings Down Light And Right Up Heavy", "┡"},
		{"Box Drawings Up Light And Right Down Heavy", "┢"},
		{"Box Drawings Heavy Vertical And Right", "┣"},
		{"Box Drawings Light Vertical And Left", "┤"},
		{"Box Drawings Vertical Light And Left Heavy", "┥"},
		{"Box Drawings Up Heavy And Left Down Light", "┦"},
		{"Box Drawings Down Heavy And Left Up Light", "┧"},
		{"Box Drawings Vertical Heavy And Left Light", "┨"},
		{"Box Drawings Down Light And Left Up Heavy", "┩"},
		{"Box Drawings Up Light And Left Down Heavy", "┪"},
		{"Box Drawings Heavy Vertical And Left", "┫"},
		{"Box Drawings Light Down And Horizontal", "┬"},
		{"Box Drawings Left Heavy And Right Down Light", "┭"},
		{"Box Drawings Right Heavy And Left Down Light", "┮"},
		{"Box Drawings Heavy Down And Horizontal", "┳"},
		{"Box Drawings Light Up And Horizontal", "┴"},
		{"Box Drawings Left Heavy And Right Up Light", "┵"},
		{"Box Drawings Right Heavy And Left Up Light", "┶"},
		{"Box Drawings Up Light And Horizontal Heavy", "┷"},
		{"Box Drawings Up Heavy And Horizontal Light", "┸"},
		{"Box Drawings Right Light And Left Up Heavy", "┹"},
		{"Box Drawings Left Light And Right Up Heavy", "┺"},
		{"Box Drawings Heavy Up And Horizontal", "┻"},
		{"Box Drawings Light Vertical And Horizontal", "┼"},
		{"Box Drawings Left Heavy And Right Vertical Light", "┽"},
		{"Box Drawings Right Heavy And Left Vertical Light", "┾"},
		{"Box Drawings Vertical Light And Horizontal Heavy", "┿"},
		{"Box Drawings Up Heavy And Down Horizontal Light", "╀"},
		{"Box Drawings Down Heavy And Up Horizontal Light", "╁"},
		{"Box Drawings Vertical Heavy And Horizontal Light", "╂"},
		{"Box Drawings Left Up Heavy And Right Down Light", "╃"},
		{"Box Drawings Right Up Heavy And Left Down Light", "╄"},
		{"Box Drawings Left Down Heavy And Right Up Light", "╅"},
		{"Box Drawings Right Down Heavy And Left Up Light", "╆"},
		{"Box Drawings Down Light And Up Horizontal Heavy", "╇"},
		{"Box Drawings Up Light And Down Horizontal Heavy", "╈"},
		{"Box Drawings Right Light And Left Vertical Heavy", "╉"},
		{"Box Drawings Left Light And Right Vertical Heavy", "╊"},
		{"Box Drawings Heavy Vertical And Horizontal", "╋"},
		{"Box Drawings Heavy Double Dash Vertical", "╏"},
		{"Box Drawings Double Vertical", "║"},
		{"Box Drawings Double Down And Right", "╔"},
		{"Box Drawings Down Single And Right Double", "╒"},
		{"Box Drawings Down Double And Right Single", "╓"},
		{"Box Drawings Down Single And Left Double", "╕"},
		{"Box Drawings Down Double And Left Single", "╖"},
		{"Box Drawings Double Down And Left", "╗"},
		{"Box Drawings Double Up And Right", "╚"},
		{"Box Drawings Up Single And Right Double", "╘"},
		{"Box Drawings Up Double And Right Single", "╙"},
		{"Box Drawings Up Single And Left Double", "╛"},
		{"Box Drawings Up Double And Left Single", "╜"},
		{"Box Drawings Double Up And Left", "╝"},
		{"Box Drawings Vertical Single And Right Double", "╞"},
		{"Box Drawings Vertical Double And Right Single", "╟"},
		{"Box Drawings Double Vertical And Right", "╠"},
		{"Box Drawings Vertical Single And Left Double", "╡"},
		{"Box Drawings Vertical Double And Left Single", "╢"},
		{"Box Drawings Double Vertical And Left", "╣"},
		{"Box Drawings Down Single And Horizontal Double", "╤"},
		{"Box Drawings Down Double And Horizontal Single", "╥"},
		{"Box Drawings Double Down And Horizontal", "╦"},
		{"Box Drawings Up Single And Horizontal Double", "╧"},
		{"Box Drawings Up Double And Horizontal Single", "╨"},
		{"Box Drawings Double Up And Horizontal", "╩"},
		{"Box Drawings Vertical Single And Horizontal Double", "╪"},
		{"Box Drawings Vertical Double And Horizontal Single", "╫"},
		{"Box Drawings Double Vertical And Horizontal", "╬"},
	}
	i = v.NewItem("Shapes")
	i.SetIcon(iconPath("shapes", "Square With Upper Right To Lower Left Fill"))
	i.SetActionRunsInBackground(false)
	i.SetActionReturnsItems(true)
	i.SetRun(func(c *Context) Items {
		items := NewItems()
		for _, row := range chars {
			i := NewItem(row[0])
			i.SetIcon(iconPath("shapes", row[0]))
			i.Run("showChar", row[0], row[1])
			i.SetAction("unisym")
			i.SetActionReturnsItems(true)
			i.SetActionRunsInBackground(false)
			items.Add(i)
		}
		return *items
	})
}
