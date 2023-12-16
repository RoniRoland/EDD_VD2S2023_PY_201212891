package PriorityQueue

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Queue struct {
	First  *QueueNode
	Lenght int
}

func (c *Queue) AddQue(studentID int, name string, class string, score int) {
	newTutor := &Tutor{StudentID: studentID, Name: name, Class: class, Score: score}
	newNode := &QueueNode{Tutor: newTutor, Next: nil, Priority: 0}

	if score >= 90 && score <= 100 {
		newNode.Priority = 1
	} else if score >= 75 && score <= 89 {
		newNode.Priority = 2
	} else if score >= 65 && score <= 74 {
		newNode.Priority = 3
	} else if score >= 61 && score <= 64 {
		newNode.Priority = 4
	} else {
		return
	}

	if c.Lenght == 0 {
		c.First = newNode
		c.Lenght++
	} else {
		aux := c.First
		for aux.Next != nil {
			if aux.Next.Priority > newNode.Priority && aux.Priority == newNode.Priority {
				newNode.Next = aux.Next
				aux.Next = newNode
				c.Lenght++
				return
			} else if aux.Next.Priority > newNode.Priority && aux.Priority < newNode.Priority {
				newNode.Next = aux.Next
				aux.Next = newNode
				c.Lenght++
				return
			} else {
				aux = aux.Next
			}
		}
		aux.Next = newNode
		c.Lenght++
	}
}

func (c *Queue) Dequeue() {
	if c.Lenght == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.First = c.First.Next
		c.Lenght--
	}
}

func (c *Queue) LeerCSV(ruta string) {
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
		studentID, _ := strconv.Atoi(linea[0])
		score, _ := strconv.Atoi(linea[3])
		c.AddQue(studentID, linea[1], "0"+linea[2], score)
	}
}

func (c *Queue) First_Queue() {
	if c.Lenght == 0 {
		fmt.Println("No hay mas Tutores")
	} else {
		fmt.Println("===============================================================")
		fmt.Println("||              CONTROL DE ESTUDIANTES TUTORES               ||")
		fmt.Println("||                                                           ||")
		fmt.Printf("||     Actual:    %-43d||\n", c.First.Tutor.StudentID)
		fmt.Printf("||     Nombre:    %-43s||\n", c.First.Tutor.Name)
		fmt.Printf("||     Curso:     %-43s||\n", c.First.Tutor.Class)
		fmt.Printf("||     Nota:      %-43d||\n", c.First.Tutor.Score)
		fmt.Printf("||     Prioridad: %-43d||\n", c.First.Priority)
		if c.First.Next != nil {
			fmt.Printf("||     Siguiente: %-43d||\n", c.First.Next.Tutor.StudentID)
			fmt.Println("||                                                           ||")
		} else {
			fmt.Printf("||     Siguiente: No hay mas tutores por evaluar             ||\n")
			fmt.Println("||                                                           ||")
		}
	}
}
