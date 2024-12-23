package global

var Lines [][]rune

var Command []rune = []rune{':'}

var ScrollOffSet int = 0

var ScrollOffNumber int = 8

var (
	CurrentColumn, CurrentLine int  = 0, 0
	RunningApp                 bool = true
)
