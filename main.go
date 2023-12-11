package main

import (
	"EDD_VD2S2023_PY_201212891/estructuras/Lists"
	"EDD_VD2S2023_PY_201212891/estructuras/PriorityQueue"
	"fmt"
)

var circularList *Lists.CircularList = &Lists.CircularList{Start: nil, Lenght: 0}
var doubleList *Lists.DoubleList = &Lists.DoubleList{Start: nil, Lenght: 0}
var priorityQueue *PriorityQueue.Queue = &PriorityQueue.Queue{First: nil, Lenght: 0}

func rowMajor(dimension int, coordenadas []int, tamanio []int) int {
	var index int

	if dimension == 1 {
		index = coordenadas[0]
	} else {
		index = coordenadas[0]
		for i := 1; i < dimension; i++ {
			index = index*tamanio[i] + coordenadas[i]
		}
	}
	return index
}

func operar() {
	dimension := 2
	coordenadas := []int{0, 2}
	tamanio := []int{3, 3}
	resultado := rowMajor(dimension, coordenadas, tamanio)
	fmt.Println(resultado)
}

func main() {
	opt := 0
	exit := false
	for !exit {
		fmt.Println("=====================================================")
		fmt.Println("||                   LOGIN                         ||")
		fmt.Println("||                                                 ||")
		fmt.Println("||              1.- Inicio de Sesion               ||")
		fmt.Println("||              2.- Salir                          ||")
		fmt.Println("||                                                 ||")
		fmt.Println("=====================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			Login()
		case 2:
			exit = true
		}
	}
}

func Login() {
	usuario := ""
	password := ""
	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201212891" && password == "admin" {
		fmt.Println("Administrador Inicio Sesion")
		MenuAdmin()
	} else if doubleList.Find(usuario, password) {
		fmt.Println("Bienvenido alumno ", usuario)
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opt := 0
	exit := false
	for !exit {
		fmt.Println("===================================================================")
		fmt.Println("||                   	MENU ADMINISTRADOR                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Carga de Estudiantes Tutores                 ||")
		fmt.Println("||              2.- Carga de Estudiantes                         ||")
		fmt.Println("||              3.- Cargar de Cursos")
		fmt.Println("||              4.- Control de Estudiantes")
		fmt.Println("||              5.- Reportes")
		fmt.Println("6. Salir")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			fmt.Println("Se cargo los cursos")
		case 4:
			ControlEstudiantes()
		case 5:
			fmt.Println("Mis Reportes")
		case 6:
			exit = true
		}

	}
}

func CargaTutores() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	priorityQueue.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")
}

func CargaEstudiantes() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	doubleList.ReadCSV(ruta)
	fmt.Println("Se cargo los estudiantes")
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		priorityQueue.First_Queue()
		fmt.Println("════════════════════")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			circularList.Add(priorityQueue.First.Tutor.StudentID, priorityQueue.First.Tutor.Name, priorityQueue.First.Tutor.Class, priorityQueue.First.Tutor.Score)
			priorityQueue.Descolar()
		} else if opcion == 2 {
			priorityQueue.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
