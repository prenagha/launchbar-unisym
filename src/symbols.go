package main

import . "github.com/nbjahan/go-launchbar"

func addSymbols() {
	var i *Item
	v := pb.GetView("main")

	chars := [][2]string{
		{"Umbrella", "☂"},
		{"Umbrella With Rain Drops", "☔"},
		{"Airplane", "✈"},
		{"Black Sun With Rays", "☀"},
		{"White Sun With Rays", "☼"},
		{"Cloud", "☁"},
		{"High Voltage Sign", "⚡"},
		{"Electrical Arrow", "⌁"},
		{"Lightning", "☇"},
		{"Thunderstorm", "☈"},
		{"Snowflake", "❄"},
		{"Tight Trifoliate Snowflake", "❅"},
		{"Heavy Chevron Snowflake", "❆"},
		{"Snowman", "☃"},
		{"Sun", "☉"},
		{"Comet", "☄"},
		{"Black Star", "★"},
		{"White Star", "☆"},
		{"First Quarter Moon", "☽"},
		{"Last Quarter Moon", "☾"},
		{"Hourglass", "⌛"},
		{"Watch", "⌚"},
		{"House", "⌂"},
		{"Electrical Intersection", "⏧"},
		{"Telephone Location Sign", "✆"},
		{"Black Telephone", "☎"},
		{"White Telephone", "☏"},
		{"Envelope", "✉"},
		{"Ballot Box With Check", "☑"},
		{"Check Mark", "✓"},
		{"Heavy Check Mark", "✔"},
		{"Radical Symbol Bottom", "⎷"},
		{"Not Check Mark", "⍻"},
		{"Heavy Multiplication X", "✖"},
		{"Ballot X", "✗"},
		{"Heavy Ballot X", "✘"},
		{"Ballot Box With X", "☒"},
		{"Multiplication X", "✕"},
		{"Saltire", "☓"},
		{"Hot Beverage", "☕"},
		{"Wheelchair Symbol", "♿"},
		{"Victory Hand", "✌"},
		{"Black Left Pointing Index", "☚"},
		{"Black Right Pointing Index", "☛"},
		{"White Left Pointing Index", "☜"},
		{"White Up Pointing Index", "☝"},
		{"White Right Pointing Index", "☞"},
		{"White Down Pointing Index", "☟"},
		{"White Frowning Face", "☹"},
		{"White Smiling Face", "☺"},
		{"Black Smiling Face", "☻"},
		{"Yin Yang", "☯"},
		{"Flower", "⚘"},
		{"Peace Symbol", "☮"},
		{"Coffin", "⚰"},
		{"Funeral Urn", "⚱"},
		{"Warning Sign", "⚠"},
		{"Skull And Crossbones", "☠"},
		{"Radioactive Sign", "☢"},
		{"Crossed Swords", "⚔"},
		{"Anchor", "⚓"},
		{"Helm Symbol", "⎈"},
		{"Hammer And Pick", "⚒"},
		{"Black Flag", "⚑"},
		{"White Flag", "⚐"},
		{"Caution Sign", "☡"},
		{"Circled Open Centre Eight Pointed Star", "❂"},
		{"Staff Of Aesculapius", "⚕"},
		{"Scales", "⚖"},
		{"Alembic", "⚗"},
		{"Tape Drive", "✇"},
		{"Biohazard Sign", "☣"},
		{"Gear", "⚙"},
		{"Caduceus", "☤"},
		{"Staff Of Hermes", "⚚"},
		{"Atom Symbol", "⚛"},
		{"Fleur De Lis", "⚜"},
		{"Ankh", "☥"},
		{"Latin Cross", "✝"},
		{"Orthodox Cross", "☦"},
		{"Chi Rho", "☧"},
		{"Cross Of Lorraine", "☨"},
		{"Cross Of Jerusalem", "☩"},
		{"Dagger", "†"},
		{"Star And Crescent", "☪"},
		{"Farsi Symbol", "☫"},
		{"Adi Shakti", "☬"},
		{"Hammer And Sickle", "☭"},
		{"Upper Blade Scissors", "✁"},
		{"Black Scissors", "✂"},
		{"Lower Blade Scissors", "✃"},
		{"White Scissors", "✄"},
		{"Writing Hand", "✍"},
		{"Lower Right Pencil", "✎"},
		{"Pencil", "✏"},
		{"Upper Right Pencil", "✐"},
		{"", ""},
		{"White Nib", "✑"},
		{"Black Nib", "✒"},
		{"Outlined Greek Cross", "✙"},
		{"Heavy Greek Cross", "✚"},
		{"Heavy Open Centre Cross", "✜"},
		{"Open Centre Cross", "✛"},
		{"West Syriac Cross", "♰"},
		{"East Syriac Cross", "♱"},
		{"Shadowed White Latin Cross", "✞"},
		{"Outlined Latin Cross", "✟"},
		{"Maltese Cross", "✠"},
		{"Star Of David", "✡"},
		{"Wheel Of Dharma", "☸"},
		{"Four Teardrop Spoked Asterisk", "✢"},
		{"Four Balloon Spoked Asterisk", "✣"},
		{"Heavy Four Balloon Spoked Asterisk", "✤"},
		{"Four Club Spoked Asterisk", "✥"},
		{"Black Four Pointed Star", "✦"},
		{"White Four Pointed Star", "✧"},
		{"Stress Outlined White Star", "✩"},
		{"Circled White Star", "✪"},
		{"Open Centre Black Star", "✫"},
		{"Black Centre White Star", "✬"},
		{"Outlined Black Star", "✭"},
		{"Heavy Outlined Black Star", "✮"},
		{"Pinwheel Star", "✯"},
		{"Shadowed White Star", "✰"},
		{"Open Centre Asterisk", "✲"},
		{"Heavy Asterisk", "✱"},
		{"Eight Spoked Asterisk", "✳"},
		{"Eight Pointed Black Star", "✴"},
		{"Eight Pointed Pinwheel Star", "✵"},
		{"Six Pointed Black Star", "✶"},
		{"Eight Pointed Rectilinear Black Star", "✷"},
		{"Heavy Eight Pointed Rectilinear Black Star", "✸"},
		{"Twelve Pointed Black Star", "✹"},
		{"Sixteen Pointed Asterisk", "✺"},
		{"Teardrop Spoked Asterisk", "✻"},
		{"Open Centre Teardrop Spoked Asterisk", "✼"},
		{"Heavy Teardrop Spoked Asterisk", "✽"},
		{"Six Petalled Black And White Florette", "✾"},
		{"White Florette", "❀"},
		{"Black Florette", "✿"},
		{"Eight Petalled Outlined Black Florette", "❁"},
		{"Heavy Teardrop Spoked Pinwheel Asterisk", "❃"},
		{"Sparkle", "❇"},
		{"Heavy Sparkle", "❈"},
		{"Balloon Spoked Asterisk", "❉"},
		{"Eight Teardrop Spoked Propeller Asterisk", "❊"},
		{"Heavy Eight Teardrop Spoked Propeller Asterisk", "❋"},
		{"Flower Punctuation Mark", "⁕"},
		{"Shamrock", "☘"},
		{"Floral Heart", "❦"},
		{"Rotated Floral Heart Bullet", "❧"},
		{"Reversed Rotated Floral Heart Bullet", "☙"},
		{"Heavy Exclamation Mark Ornament", "❢"},
		{"Heavy Heart Exclamation Mark Ornament", "❣"},
		{"Female Sign", "♀"},
		{"Male Sign", "♂"},
		{"Neuter", "⚲"},
		{"Doubled Female Sign", "⚢"},
		{"Doubled Male Sign", "⚣"},
		{"Interlocked Female And Male Sign", "⚤"},
		{"Male And Female Sign", "⚥"},
		{"Male With Stroke Sign", "⚦"},
		{"Male With Stroke And Male And Female Sign", "⚧"},
		{"Vertical Male With Stroke Sign", "⚨"},
		{"Horizontal Male With Stroke Sign", "⚩"},
		{"Mercury", "☿"},
		{"Earth", "♁"},
		{"Unmarried Partnership Symbol", "⚯"},
		{"Black Chess Queen", "♛"},
		{"White Chess Queen", "♕"},
		{"Black Chess King", "♚"},
		{"White Chess King", "♔"},
		{"Black Chess Rook", "♜"},
		{"White Chess Rook", "♖"},
		{"Black Chess Bishop", "♝"},
		{"White Chess Bishop", "♗"},
		{"Black Chess Knight", "♞"},
		{"White Chess Knight", "♘"},
		{"Black Chess Pawn", "♟"},
		{"White Chess Pawn", "♙"},
		{"Black Shogi Piece", "☗"},
		{"White Shogi Piece", "☖"},
		{"Black Spade Suit", "♠"},
		{"Black Club Suit", "♣"},
		{"Black Diamond Suit", "♦"},
		{"Black Heart Suit", "♥"},
		{"Heavy Black Heart", "❤"},
		{"Rotated Heavy Black Heart Bullet", "❥"},
		{"White Heart Suit", "♡"},
		{"White Diamond Suit", "♢"},
		{"White Spade Suit", "♤"},
		{"White Club Suit", "♧"},
		{"Die Face 1", "⚀"},
		{"Die Face 2", "⚁"},
		{"Die Face 3", "⚂"},
		{"Die Face 4", "⚃"},
		{"Die Face 5", "⚄"},
		{"Die Face 6", "⚅"},
		{"White Circle With Two Dots", "⚇"},
		{"White Circle With Dot Right", "⚆"},
		{"Black Circle With White Dot Right", "⚈"},
		{"Black Circle With Two White Dots", "⚉"},
		{"Hot Springs", "♨"},
		{"Quarter Note", "♩"},
		{"Eighth Note", "♪"},
		{"Beamed Eighth Notes", "♫"},
		{"Beamed Sixteenth Notes", "♬"},
		{"Music Flat Sign", "♭"},
		{"Music Natural Sign", "♮"},
		{"Music Sharp Sign", "♯"},
		{"", ""},
		{"Keyboard", "⌨"},
		{"Eject Symbol", "⏏"},
		{"Previous Page", "⎗"},
		{"Next Page", "⎘"},
		{"Print Screen Symbol", "⎙"},
		{"Clear Screen Symbol", "⎚"},
		{"Option Key", "⌥"},
		{"Alternative Key Symbol", "⎇"},
		{"Place Of Interest Sign", "⌘"},
		{"Erase To The Right", "⌦"},
		{"Erase To The Left", "⌫"},
		{"X In A Rectangle Box", "⌧"},
		{"Universal Recycling Symbol", "♲"},
		{"Recycling Symbol For Type 1 Plastics", "♳"},
		{"Recycling Symbol For Type 2 Plastics", "♴"},
		{"Recycling Symbol For Type 3 Plastics", "♵"},
		{"Recycling Symbol For Type 4 Plastics", "♶"},
		{"Recycling Symbol For Type 5 Plastics", "♷"},
		{"Recycling Symbol For Type 6 Plastics", "♸"},
		{"Recycling Symbol For Type 7 Plastics", "♹"},
		{"Recycling Symbol For Generic Materials", "♺"},
		{"Black Universal Recycling Symbol", "♻"},
		{"Recycled Paper Symbol", "♼"},
		{"Partially Recycled Paper Symbol", "♽"},
		{"Black Leftwards Bullet", "⁌"},
		{"Black Rightwards Bullet", "⁍"},
		{"Undo Symbol", "⎌"},
		{"Wavy Line", "⌇"},
		{"Conical Taper", "⌲"},
		{"Apl Functional Symbol Up Shoe Jot", "⍝"},
		{"Apl Functional Symbol Circle Star", "⍟"},
		{"Apl Functional Symbol Star Diaeresis", "⍣"},
		{"Apl Functional Symbol Jot Diaeresis", "⍤"},
		{"Apl Functional Symbol Circle Diaeresis", "⍥"},
		{"Apl Functional Symbol Tilde Diaeresis", "⍨"},
		{"Apl Functional Symbol Greater Than Diaeresis", "⍩"},
		{"Broken Circle With Northwest Arrow", "⎋"},
		{"Jupiter", "♃"},
		{"Saturn", "♄"},
		{"Uranus", "♅"},
		{"Neptune", "♆"},
		{"Pluto", "♇"},
		{"Aries", "♈"},
		{"Taurus", "♉"},
		{"Gemini", "♊"},
		{"Cancer", "♋"},
		{"Leo", "♌"},
		{"Virgo", "♍"},
		{"Libra", "♎"},
		{"Scorpius", "♏"},
		{"Sagittarius", "♐"},
		{"Capricorn", "♑"},
		{"Aquarius", "♒"},
		{"Pisces", "♓"},
		{"Earth Ground", "⏚"},
		{"Fuse", "⏛"},
	}
	i = v.NewItem("Symbols")
	i.SetIcon(iconPath("symbols", "Yin Yang"))
	i.SetActionRunsInBackground(false)
	i.SetActionReturnsItems(true)
	i.SetRun(func(c *Context) Items {
		items := NewItems()
		for _, row := range chars {
			i := NewItem(row[0])
			i.SetIcon(iconPath("symbols", row[0]))
			i.Run("showChar", row[0], row[1])
			i.SetAction("unisym")
			i.SetActionReturnsItems(true)
			i.SetActionRunsInBackground(false)
			items.Add(i)
		}
		return *items
	})
}
