package Lists

type CircularList struct {
	Start  *CircularNode
	Lenght int
}

func (l *CircularList) Add(studentID int, name string, class string, score int) {
	newTutor := &Tutor{StudentID: studentID, Name: name}
	newNode := &CircularNode{Tutor: newTutor, Next: nil, Previous: nil}

	if l.Lenght == 0 {
		l.Start = newNode
		l.Start.Previous = newNode
		l.Start.Next = newNode
		l.Lenght++
	} else {
		aux := l.Start
		cont := 1
		for cont < l.Lenght {
			if l.Start.Tutor.StudentID > studentID {
				newNode.Next = l.Start
				newNode.Previous = l.Start.Previous
				l.Start.Previous = newNode
				l.Start = newNode
				l.Lenght++
				return
			}
			if aux.Tutor.StudentID < studentID {
				aux = aux.Next
			} else {
				newNode.Previous = aux.Previous
				aux.Previous.Next = newNode
				newNode.Next = aux
				aux.Previous = newNode
				l.Lenght++
				return
			}
			cont++
		}
		if aux.Tutor.StudentID > studentID {
			newNode.Next = aux
			newNode.Previous = aux.Previous
			aux.Previous.Next = newNode
			aux.Previous = newNode
			l.Lenght++
			return
		}
		newNode.Previous = aux
		newNode.Next = l.Start
		aux.Next = newNode
		l.Start.Previous = newNode
		l.Lenght++
	}
}
