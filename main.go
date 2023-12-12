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
		AdminMenu()
	} else if doubleList.Find(usuario, password) {
		fmt.Println("Bienvenido alumno ", usuario)
		pressEnterToContinue()
		StudentMenu()
	} else {
		fmt.Println("\n Contraseña y usuario INCORRECTOS")
		pressEnterToContinue()
		clearScreen()
	}
}

func AdminMenu() {
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
			TutorLoadCSV()
		case 2:
			StudentsLoadCSV()
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

func StudentMenu() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("═══════════════════════════════════════════════════════════════════")
		fmt.Println("||                   	MENU ESTUDIANTE                           ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Ver Tutores disponibles                      ||")
		fmt.Println("||              2.- Asignarse a tutores                          ||")
		fmt.Println("||              3.- Salir                                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("===================================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("tutos disponibles")
		case 2:
			fmt.Println("asignacion")
		case 3:
			exit = true
		}

	}
}

func TutorLoadCSV() {
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

func StudentsLoadCSV() {
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
		clearScreen()
		priorityQueue.First_Queue()
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Println("||                  1.- Aceptar                              ||")
		fmt.Println("||                  2.- Rechazar                             ||")
		fmt.Println("||                  3.- Salir                                ||")
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			circularList.Add(priorityQueue.First.Tutor.StudentID, priorityQueue.First.Tutor.Name, priorityQueue.First.Tutor.Class, priorityQueue.First.Tutor.Score)
			priorityQueue.Dequeue()
		} else if opcion == 2 {
			priorityQueue.Dequeue()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
