package Lists

import (
	"EDD_VD2S2023_PY_201212891/estructuras/GenerarArchivos"
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
		fmt.Println("\nError! Archivo incorrecto")
		return
	}
	defer file.Close()

	readCSV := csv.NewReader(file)
	readCSV.Comma = ','
	header := true
	for {
		linea, err := readCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("\nError! Lineas del csv INCORRECTAS")
			continue
		}
		if header {
			header = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		l.Add(valor, linea[1])
	}
	fmt.Println("\n*********Archivo CARGADO EXITOSAMENTE********** ")
}

func (l *DoubleList) Reporte() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Start
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Lenght; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"{" + strconv.Itoa(aux.Alumn.StudentID) + " | " + aux.Alumn.Name + "}\"];\n"
		aux = aux.Next
	}
	for i := 0; i < l.Lenght-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	GenerarArchivos.CrearArchivo(nombreArchivo)
	GenerarArchivos.EscribirArchivo(texto, nombreArchivo)
	GenerarArchivos.Ejecutar(nombreImagen, nombreArchivo)
}
