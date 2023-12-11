package main

import "fmt"

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
		fmt.Scan(&opt)

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
	} else if listaDoble.Buscar(usuario, password) {
		fmt.Println("Bienvenido alumno ", usuario)
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar de Cursos")
		fmt.Println("4. Control de Estudiantes")
		fmt.Println("5. Reportes")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
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
			salir = true
		}

	}
}
