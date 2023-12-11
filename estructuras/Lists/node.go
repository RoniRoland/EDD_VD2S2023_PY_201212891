package Lists

type Alumn struct {
	StudentID int
	Name      string
}

type Tutor struct {
	StudentID int
	Name      string
	Class     string
	Score     int
}

type DoubleListNode struct {
	Alumn    *Alumn
	Next     *DoubleListNode
	Previous *DoubleListNode
}

type CircularNode struct {
	Tutor    *Tutor
	Next     *CircularNode
	Previous *CircularNode
}
