package global

var Lines [][]rune

var Command []rune = []rune{':'}

var (
	CurrentFileName            string = ""
	CurrentColumn, CurrentLine int    = 0, 0
	RunningApp                 bool   = true
)
