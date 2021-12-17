package staff

type Base struct {
	Name		string
	ShowName	string
	Peroration	string
	LV			int
}

type Jack struct {
	Base
	IntegerArray []int
}

type Rose struct {
	Base
	StringArray []string
}

type PhoenixLegend struct {
	Base
	Song	string
}

type FallOutBoy struct {
	Base
	Song	string
}