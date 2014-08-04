package main

func init() {
	chars["02200"] = [][2]string{
		{"For All", "\U00002200"},
		{"Complement", "\U00002201"},
		{"Partial Differential", "\U00002202"},
		{"There Exists", "\U00002203"},
		{"There Does Not Exist", "\U00002204"},
		{"Empty Set", "\U00002205"},
		{"Increment", "\U00002206"},
		{"Nabla", "\U00002207"},
		{"Element Of", "\U00002208"},
		{"Not An Element Of", "\U00002209"},
		{"Small Element Of", "\U0000220a"},
		{"Contains As Member", "\U0000220b"},
		{"Does Not Contain As Member", "\U0000220c"},
		{"Small Contains As Member", "\U0000220d"},
		{"End Of Proof", "\U0000220e"},
		{"N-Ary Product", "\U0000220f"},
		{"N-Ary Coproduct", "\U00002210"},
		{"N-Ary Summation", "\U00002211"},
		{"Minus Sign", "\U00002212"},
		{"Minus-Or-Plus Sign", "\U00002213"},
		{"Dot Plus", "\U00002214"},
		{"Division Slash", "\U00002215"},
		{"Set Minus", "\U00002216"},
		{"Asterisk Operator", "\U00002217"},
		{"Ring Operator", "\U00002218"},
		{"Bullet Operator", "\U00002219"},
		{"Square Root", "\U0000221a"},
		{"Cube Root", "\U0000221b"},
		{"Fourth Root", "\U0000221c"},
		{"Proportional To", "\U0000221d"},
		{"Infinity", "\U0000221e"},
		{"Right Angle", "\U0000221f"},
		{"Angle", "\U00002220"},
		{"Measured Angle", "\U00002221"},
		{"Spherical Angle", "\U00002222"},
		{"Divides", "\U00002223"},
		{"Does Not Divide", "\U00002224"},
		{"Parallel To", "\U00002225"},
		{"Not Parallel To", "\U00002226"},
		{"Logical And", "\U00002227"},
		{"Logical Or", "\U00002228"},
		{"Intersection", "\U00002229"},
		{"Union", "\U0000222a"},
		{"Integral", "\U0000222b"},
		{"Double Integral", "\U0000222c"},
		{"Triple Integral", "\U0000222d"},
		{"Contour Integral", "\U0000222e"},
		{"Surface Integral", "\U0000222f"},
		{"Volume Integral", "\U00002230"},
		{"Clockwise Integral", "\U00002231"},
		{"Clockwise Contour Integral", "\U00002232"},
		{"Anticlockwise Contour Integral", "\U00002233"},
		{"Therefore", "\U00002234"},
		{"Because", "\U00002235"},
		{"Ratio", "\U00002236"},
		{"Proportion", "\U00002237"},
		{"Dot Minus", "\U00002238"},
		{"Excess", "\U00002239"},
		{"Geometric Proportion", "\U0000223a"},
		{"Homothetic", "\U0000223b"},
		{"Tilde Operator", "\U0000223c"},
		{"Reversed Tilde", "\U0000223d"},
		{"Inverted Lazy S", "\U0000223e"},
		{"Sine Wave", "\U0000223f"},
		{"Wreath Product", "\U00002240"},
		{"Not Tilde", "\U00002241"},
		{"Minus Tilde", "\U00002242"},
		{"Asymptotically Equal To", "\U00002243"},
		{"Not Asymptotically Equal To", "\U00002244"},
		{"Approximately Equal To", "\U00002245"},
		{"Approximately But Not Actually Equal To", "\U00002246"},
		{"Neither Approximately Nor Actually Equal To", "\U00002247"},
		{"Almost Equal To", "\U00002248"},
		{"Not Almost Equal To", "\U00002249"},
		{"Almost Equal Or Equal To", "\U0000224a"},
		{"Triple Tilde", "\U0000224b"},
		{"All Equal To", "\U0000224c"},
		{"Equivalent To", "\U0000224d"},
		{"Geometrically Equivalent To", "\U0000224e"},
		{"Difference Between", "\U0000224f"},
		{"Approaches The Limit", "\U00002250"},
		{"Geometrically Equal To", "\U00002251"},
		{"Approximately Equal To Or The Image Of", "\U00002252"},
		{"Image Of Or Approximately Equal To", "\U00002253"},
		{"Colon Equals", "\U00002254"},
		{"Equals Colon", "\U00002255"},
		{"Ring In Equal To", "\U00002256"},
		{"Ring Equal To", "\U00002257"},
		{"Corresponds To", "\U00002258"},
		{"Estimates", "\U00002259"},
		{"Equiangular To", "\U0000225a"},
		{"Star Equals", "\U0000225b"},
		{"Delta Equal To", "\U0000225c"},
		{"Equal To By Definition", "\U0000225d"},
		{"Measured By", "\U0000225e"},
		{"Questioned Equal To", "\U0000225f"},
		{"Not Equal To", "\U00002260"},
		{"Identical To", "\U00002261"},
		{"Not Identical To", "\U00002262"},
		{"Strictly Equivalent To", "\U00002263"},
		{"Less-Than Or Equal To", "\U00002264"},
		{"Greater-Than Or Equal To", "\U00002265"},
		{"Less-Than Over Equal To", "\U00002266"},
		{"Greater-Than Over Equal To", "\U00002267"},
		{"Less-Than But Not Equal To", "\U00002268"},
		{"Greater-Than But Not Equal To", "\U00002269"},
		{"Much Less-Than", "\U0000226a"},
		{"Much Greater-Than", "\U0000226b"},
		{"Between", "\U0000226c"},
		{"Not Equivalent To", "\U0000226d"},
		{"Not Less-Than", "\U0000226e"},
		{"Not Greater-Than", "\U0000226f"},
		{"Neither Less-Than Nor Equal To", "\U00002270"},
		{"Neither Greater-Than Nor Equal To", "\U00002271"},
		{"Less-Than Or Equivalent To", "\U00002272"},
		{"Greater-Than Or Equivalent To", "\U00002273"},
		{"Neither Less-Than Nor Equivalent To", "\U00002274"},
		{"Neither Greater-Than Nor Equivalent To", "\U00002275"},
		{"Less-Than Or Greater-Than", "\U00002276"},
		{"Greater-Than Or Less-Than", "\U00002277"},
		{"Neither Less-Than Nor Greater-Than", "\U00002278"},
		{"Neither Greater-Than Nor Less-Than", "\U00002279"},
		{"Precedes", "\U0000227a"},
		{"Succeeds", "\U0000227b"},
		{"Precedes Or Equal To", "\U0000227c"},
		{"Succeeds Or Equal To", "\U0000227d"},
		{"Precedes Or Equivalent To", "\U0000227e"},
		{"Succeeds Or Equivalent To", "\U0000227f"},
		{"Does Not Precede", "\U00002280"},
		{"Does Not Succeed", "\U00002281"},
		{"Subset Of", "\U00002282"},
		{"Superset Of", "\U00002283"},
		{"Not A Subset Of", "\U00002284"},
		{"Not A Superset Of", "\U00002285"},
		{"Subset Of Or Equal To", "\U00002286"},
		{"Superset Of Or Equal To", "\U00002287"},
		{"Neither A Subset Of Nor Equal To", "\U00002288"},
		{"Neither A Superset Of Nor Equal To", "\U00002289"},
		{"Subset Of With Not Equal To", "\U0000228a"},
		{"Superset Of With Not Equal To", "\U0000228b"},
		{"Multiset", "\U0000228c"},
		{"Multiset Multiplication", "\U0000228d"},
		{"Multiset Union", "\U0000228e"},
		{"Square Image Of", "\U0000228f"},
		{"Square Original Of", "\U00002290"},
		{"Square Image Of Or Equal To", "\U00002291"},
		{"Square Original Of Or Equal To", "\U00002292"},
		{"Square Cap", "\U00002293"},
		{"Square Cup", "\U00002294"},
		{"Circled Plus", "\U00002295"},
		{"Circled Minus", "\U00002296"},
		{"Circled Times", "\U00002297"},
		{"Circled Division Slash", "\U00002298"},
		{"Circled Dot Operator", "\U00002299"},
		{"Circled Ring Operator", "\U0000229a"},
		{"Circled Asterisk Operator", "\U0000229b"},
		{"Circled Equals", "\U0000229c"},
		{"Circled Dash", "\U0000229d"},
		{"Squared Plus", "\U0000229e"},
		{"Squared Minus", "\U0000229f"},
		{"Squared Times", "\U000022a0"},
		{"Squared Dot Operator", "\U000022a1"},
		{"Right Tack", "\U000022a2"},
		{"Left Tack", "\U000022a3"},
		{"Down Tack", "\U000022a4"},
		{"Up Tack", "\U000022a5"},
		{"Assertion", "\U000022a6"},
		{"Models", "\U000022a7"},
		{"True", "\U000022a8"},
		{"Forces", "\U000022a9"},
		{"Triple Vertical Bar Right Turnstile", "\U000022aa"},
		{"Double Vertical Bar Double Right Turnstile", "\U000022ab"},
		{"Does Not Prove", "\U000022ac"},
		{"Not True", "\U000022ad"},
		{"Does Not Force", "\U000022ae"},
		{"Negated Double Vertical Bar Double Right Turnstile", "\U000022af"},
		{"Precedes Under Relation", "\U000022b0"},
		{"Succeeds Under Relation", "\U000022b1"},
		{"Normal Subgroup Of", "\U000022b2"},
		{"Contains As Normal Subgroup", "\U000022b3"},
		{"Normal Subgroup Of Or Equal To", "\U000022b4"},
		{"Contains As Normal Subgroup Or Equal To", "\U000022b5"},
		{"Original Of", "\U000022b6"},
		{"Image Of", "\U000022b7"},
		{"Multimap", "\U000022b8"},
		{"Hermitian Conjugate Matrix", "\U000022b9"},
		{"Intercalate", "\U000022ba"},
		{"Xor", "\U000022bb"},
		{"Nand", "\U000022bc"},
		{"Nor", "\U000022bd"},
		{"Right Angle With Arc", "\U000022be"},
		{"Right Triangle", "\U000022bf"},
		{"N-Ary Logical And", "\U000022c0"},
		{"N-Ary Logical Or", "\U000022c1"},
		{"N-Ary Intersection", "\U000022c2"},
		{"N-Ary Union", "\U000022c3"},
		{"Diamond Operator", "\U000022c4"},
		{"Dot Operator", "\U000022c5"},
		{"Star Operator", "\U000022c6"},
		{"Division Times", "\U000022c7"},
		{"Bowtie", "\U000022c8"},
		{"Left Normal Factor Semidirect Product", "\U000022c9"},
		{"Right Normal Factor Semidirect Product", "\U000022ca"},
		{"Left Semidirect Product", "\U000022cb"},
		{"Right Semidirect Product", "\U000022cc"},
		{"Reversed Tilde Equals", "\U000022cd"},
		{"Curly Logical Or", "\U000022ce"},
		{"Curly Logical And", "\U000022cf"},
		{"Double Subset", "\U000022d0"},
		{"Double Superset", "\U000022d1"},
		{"Double Intersection", "\U000022d2"},
		{"Double Union", "\U000022d3"},
		{"Pitchfork", "\U000022d4"},
		{"Equal And Parallel To", "\U000022d5"},
		{"Less-Than With Dot", "\U000022d6"},
		{"Greater-Than With Dot", "\U000022d7"},
		{"Very Much Less-Than", "\U000022d8"},
		{"Very Much Greater-Than", "\U000022d9"},
		{"Less-Than Equal To Or Greater-Than", "\U000022da"},
		{"Greater-Than Equal To Or Less-Than", "\U000022db"},
		{"Equal To Or Less-Than", "\U000022dc"},
		{"Equal To Or Greater-Than", "\U000022dd"},
		{"Equal To Or Precedes", "\U000022de"},
		{"Equal To Or Succeeds", "\U000022df"},
		{"Does Not Precede Or Equal", "\U000022e0"},
		{"Does Not Succeed Or Equal", "\U000022e1"},
		{"Not Square Image Of Or Equal To", "\U000022e2"},
		{"Not Square Original Of Or Equal To", "\U000022e3"},
		{"Square Image Of Or Not Equal To", "\U000022e4"},
		{"Square Original Of Or Not Equal To", "\U000022e5"},
		{"Less-Than But Not Equivalent To", "\U000022e6"},
		{"Greater-Than But Not Equivalent To", "\U000022e7"},
		{"Precedes But Not Equivalent To", "\U000022e8"},
		{"Succeeds But Not Equivalent To", "\U000022e9"},
		{"Not Normal Subgroup Of", "\U000022ea"},
		{"Does Not Contain As Normal Subgroup", "\U000022eb"},
		{"Not Normal Subgroup Of Or Equal To", "\U000022ec"},
		{"Does Not Contain As Normal Subgroup Or Equal", "\U000022ed"},
		{"Vertical Ellipsis", "\U000022ee"},
		{"Midline Horizontal Ellipsis", "\U000022ef"},
		{"Up Right Diagonal Ellipsis", "\U000022f0"},
		{"Down Right Diagonal Ellipsis", "\U000022f1"},
		{"Element Of With Long Horizontal Stroke", "\U000022f2"},
		{"Element Of With Vertical Bar At End Of Horizontal Stroke", "\U000022f3"},
		{"Small Element Of With Vertical Bar At End Of Horizontal Stroke", "\U000022f4"},
		{"Element Of With Dot Above", "\U000022f5"},
		{"Element Of With Overbar", "\U000022f6"},
		{"Small Element Of With Overbar", "\U000022f7"},
		{"Element Of With Underbar", "\U000022f8"},
		{"Element Of With Two Horizontal Strokes", "\U000022f9"},
		{"Contains With Long Horizontal Stroke", "\U000022fa"},
		{"Contains With Vertical Bar At End Of Horizontal Stroke", "\U000022fb"},
		{"Small Contains With Vertical Bar At End Of Horizontal Stroke", "\U000022fc"},
		{"Contains With Overbar", "\U000022fd"},
		{"Small Contains With Overbar", "\U000022fe"},
	}
}
