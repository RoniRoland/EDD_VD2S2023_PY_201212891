package AVLTree

import (
	"EDD_VD2S2023_PY_201212891/estructuras/GenerarArchivos"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

type AVLTree struct {
	Root *TreeNode
}

/***************************************/
type Curso struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type DatosCursos struct {
	Cursos []Curso `json:"Cursos"`
}

/***************************************/

func (a *AVLTree) altura(raiz *TreeNode) int {
	if raiz == nil {
		return 0
	}
	return raiz.Height
}

func (a *AVLTree) equilibrio(raiz *TreeNode) int {
	if raiz == nil {
		return 0
	}
	return (a.altura(raiz.Right) - a.altura(raiz.Left)) // 1 - 0
}

func (a *AVLTree) rotacionI(raiz *TreeNode) *TreeNode { //Root = 10
	raiz_derecho := raiz.Right          // 10.derecho = 15
	hijo_izquierdo := raiz_derecho.Left // 10.derecho.izquierdo = null
	raiz_derecho.Left = raiz            // 15.izquierdo = 10
	raiz.Right = hijo_izquierdo         // 10.derecho = null
	/*Calcular nuevamente alturas de raiz*/
	numeroMax := math.Max(float64(a.altura(raiz.Left)), float64(a.altura(raiz.Right)))
	raiz.Height = 1 + int(numeroMax)
	raiz.Equilibrium_Factor = a.equilibrio(raiz)
	/*Calcular nuevamente alturas de raiz.derecho*/
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Left)), float64(a.altura(raiz_derecho.Right)))
	raiz_derecho.Height = 1 + int(numeroMax)
	raiz_derecho.Equilibrium_Factor = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *AVLTree) rotacionD(raiz *TreeNode) *TreeNode { //Root = 20
	raiz_izquierdo := raiz.Left          // 20.izquierdo = 15
	hijo_derecho := raiz_izquierdo.Right // 20.izquierdo.derecho = null
	raiz_izquierdo.Right = raiz          //15.derecho = 20
	raiz.Left = hijo_derecho             // 20.izquierdo = null
	/*Calcular nuevamente alturas de raiz.derecho*/
	numeroMax := math.Max(float64(a.altura(raiz.Left)), float64(a.altura(raiz.Right)))
	raiz.Height = 1 + int(numeroMax)
	raiz.Equilibrium_Factor = a.equilibrio(raiz)
	numeroMax = math.Max(float64(a.altura(raiz_izquierdo.Left)), float64(a.altura(raiz_izquierdo.Right)))
	raiz_izquierdo.Height = 1 + int(numeroMax)
	raiz_izquierdo.Equilibrium_Factor = a.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (a *AVLTree) insertarNodo(raiz *TreeNode, nuevoNodo *TreeNode) *TreeNode {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Data > nuevoNodo.Data {
			raiz.Left = a.insertarNodo(raiz.Left, nuevoNodo)
		} else {
			raiz.Right = a.insertarNodo(raiz.Right, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(a.altura(raiz.Left)), float64(a.altura(raiz.Right)))
	raiz.Height = 1 + int(numeroMax)
	balanceo := a.equilibrio(raiz)
	raiz.Equilibrium_Factor = balanceo
	if balanceo > 1 && nuevoNodo.Data > raiz.Right.Data {
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Data < raiz.Left.Data {
		return a.rotacionD(raiz)
	} else if balanceo > 1 && nuevoNodo.Data < raiz.Right.Data {
		raiz.Right = a.rotacionD(raiz.Right)
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Data > raiz.Left.Data {
		raiz.Left = a.rotacionI(raiz.Left)
		return a.rotacionD(raiz)
	}
	return raiz
}

func (a *AVLTree) InsertarElemento(valor string) {
	nuevoNodo := &TreeNode{Data: valor}
	a.Root = a.insertarNodo(a.Root, nuevoNodo)
}

func (a *AVLTree) busqueda_arbol(valor string, raiz *TreeNode) *TreeNode {
	var valorEncontro *TreeNode
	if raiz != nil {
		if raiz.Data == valor {
			valorEncontro = raiz
		} else {
			if raiz.Data > valor {
				valorEncontro = a.busqueda_arbol(valor, raiz.Left)
			} else {
				valorEncontro = a.busqueda_arbol(valor, raiz.Right)
			}
		}
	}
	return valorEncontro
}

func (a *AVLTree) Busqueda(valor string) bool {
	buscarElemento := a.busqueda_arbol(valor, a.Root)
	if buscarElemento != nil {
		return true
	}
	return false
}

func (a *AVLTree) LeerJson(ruta string) {
	data, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("\nError! Archivo incorrecto")
	}

	var datos DatosCursos
	err = json.Unmarshal(data, &datos)
	if err != nil {
		fmt.Println("\nError! Lineas del JSON INCORRECTAS")
	}

	for _, curso := range datos.Cursos {
		a.InsertarElemento(curso.Codigo)
	}
	fmt.Println("\n*********Archivo CARGADO EXITOSAMENTE********** ")
}

func (a *AVLTree) Graficar() {
	cadena := ""
	nombre_archivo := "./ArbolAVL.dot"
	nombre_imagen := "ArbolAVL.jpg"
	if a.Root != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Root, 0)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *AVLTree) retornarValoresArbol(raiz *TreeNode, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Data
		cadena += "\" ;"
		if raiz.Left != nil && raiz.Right != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Left, numero)
			cadena += "\""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Right, numero)
			cadena += "{rank=same" + "\"" + (raiz.Left.Data) + "\"" + " -> " + "\"" + (raiz.Right.Data) + "\"" + " [style=invis]}; "
		} else if raiz.Left != nil && raiz.Right == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Left, numero)
			cadena += "\""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Left.Data) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Left == nil && raiz.Right != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Data
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Right, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Right.Data) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}
