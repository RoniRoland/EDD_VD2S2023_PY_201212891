package Lists

import (
	"EDD_VD2S2023_PY_201212891/estructuras/GenerarArchivos"
	"fmt"
	"strconv"
)

type CircularList struct {
	Start  *CircularNode
	Lenght int
}

func (l *CircularList) Add(studentID int, name string, class string, score int) {
	newTutor := &Tutor{StudentID: studentID, Name: name, Class: class, Score: score}
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

func (l *CircularList) Show() {
	aux := l.Start
	cont := 1
	for cont <= l.Lenght {
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Println(cont, ".- ", " ", aux.Tutor.Class, " -> ", aux.Tutor.Name)
		fmt.Println("═══════════════════════════════════════════════════════════════")
		aux = aux.Next
		cont++
	}
}

func (l *CircularList) Find(class string) bool {
	if l.Lenght == 0 {
		return false
	} else {
		aux := l.Start
		cont := 1
		for l.Lenght >= cont {
			if aux.Tutor.Class == class {
				return true
			}
			aux = aux.Next
			cont++
		}
	}
	return false
}

func (l *CircularList) FindTutor(class string) *CircularNode {
	aux := l.Start
	cont := 1
	for l.Lenght >= cont {
		if aux.Tutor.Class == class {
			return aux
		}
		aux = aux.Next
		cont++
	}
	return nil
}

func (l *CircularList) Reportev2() {
	nombreArchivo := "./listadoblecircular.dot"
	nombreImagen := "./listadoblecircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Start
	contador := 0
	for i := 0; i < l.Lenght; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"{" + strconv.Itoa(aux.Tutor.StudentID) + " | " + aux.Tutor.Name + "}\"];\n"
		aux = aux.Next
	}
	for i := 0; i < l.Lenght-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0 \n"
	texto += "nodo0 -> " + "nodo" + strconv.Itoa(contador) + "\n"
	texto += "}"
	GenerarArchivos.CrearArchivo(nombreArchivo)
	GenerarArchivos.EscribirArchivo(texto, nombreArchivo)
	GenerarArchivos.Ejecutar(nombreImagen, nombreArchivo)
}
