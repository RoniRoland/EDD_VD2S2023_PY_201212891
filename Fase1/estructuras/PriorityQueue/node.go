package PriorityQueue

type Tutor struct {
	StudentID int
	Name      string
	Class     string
	Score     int
}

type QueueNode struct {
	Tutor    *Tutor
	Next     *QueueNode
	Priority int
}
