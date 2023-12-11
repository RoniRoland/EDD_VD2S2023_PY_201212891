package main

import (
	"EDD_VD2S2023_PY_201212891/estructuras/Lists"
	"EDD_VD2S2023_PY_201212891/estructuras/PriorityQueue"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var circularList *Lists.CircularList = &Lists.CircularList{Start: nil, Lenght: 0}
var doubleList *Lists.DoubleList = &Lists.DoubleList{Start: nil, Lenght: 0}
var priorityQueue *PriorityQueue.Queue = &PriorityQueue.Queue{First: nil, Lenght: 0}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("No se puede limpiar la pantalla en este sistema operativo.")
	}
}

func pressEnterToContinue() {
	fmt.Print("\nPresiona Enter para continuar...")
	fmt.Scanln()
}

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
		clearScreen()
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
			clearScreen()
			Login()
		case 2:
			exit = true
		}
	}
}

func Login() {
	usuario := ""
	password := ""
	fmt.Println("=====================================================")
	fmt.Println("||                   LOGIN                         ||")
	fmt.Println("=====================================================")
	fmt.Print("               Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("               Password: ")
	fmt.Scanln(&password)

	if usuario == "admin" && password == "admin" {
		fmt.Println("\n*****Sesion Iniciada Correctamente******")
		pressEnterToContinue()
		MenuAdmin()
	} else if doubleList.Find(usuario, password) {
		fmt.Println("Bienvenido alumno ", usuario)
	} else {
		fmt.Println("\n Contraseña y usuario INCORRECTOS")
		pressEnterToContinue()
		clearScreen()
	}
}

func MenuAdmin() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("═══════════════════════════════════════════════════════════════════")
		fmt.Println("||                   	MENU ADMINISTRADOR                       ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Carga de Estudiantes Tutores                 ||")
		fmt.Println("||              2.- Carga de Estudiantes                         ||")
		fmt.Println("||              3.- Cargar de Cursos                             ||")
		fmt.Println("||              4.- Control de Estudiantes                       ||")
		fmt.Println("||              5.- Reportes                                     ||")
		fmt.Println("||              6.- Salir                                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("===================================================================")
		fmt.Print("               Ingrese opcion: ")
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
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("||              CARGA DE TUTORES                    ||")
	fmt.Println("=====================================================")
	fmt.Print("             Nombre del Archivo CSV: ")
	fmt.Scanln(&ruta)
	priorityQueue.LeerCSV(ruta)
	fmt.Println("\n*********Archivo CARGADO EXITOSAMENTE********** ")
	pressEnterToContinue()
}

func CargaEstudiantes() {
	ruta := ""
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("||              CARGA DE ESTUDIANTES               ||")
	fmt.Println("=====================================================")
	fmt.Print("          Nombre del Archivo CSV: ")
	fmt.Scanln(&ruta)
	doubleList.ReadCSV(ruta)
	fmt.Println("\n*********Archivo CARGADO EXITOSAMENTE********** ")
	pressEnterToContinue()
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
