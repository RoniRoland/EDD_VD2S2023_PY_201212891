package Lists

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type DoubleList struct {
	Start  *DoubleListNode
	Lenght int
}

func (l *DoubleList) Add(studentID int, name string) {
	newAlumn := &Alumn{StudentID: studentID, Name: name}
	newNode := &DoubleListNode{Alumn: newAlumn, Next: nil, Previous: nil}

	if l.Lenght == 0 {
		l.Start = newNode
		l.Lenght++
	} else {
		aux := l.Start
		for aux.Next != nil {
			aux = aux.Next
		}
		newNode.Previous = aux
		aux.Next = newNode
		l.Lenght++
	}
}

func (l *DoubleList) Find(studentID string, password string) bool {
	if l.Lenght == 0 {
		return false
	} else {
		aux := l.Start
		for aux != nil {
			if strconv.Itoa(aux.Alumn.StudentID) == studentID && studentID == password {
				return true
			}
			aux = aux.Next
		}
	}

	return false
}

func (l *DoubleList) ReadCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		l.Add(valor, linea[1])
	}
}
