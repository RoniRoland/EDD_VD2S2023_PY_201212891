package main

import (
	"Fase2/estructuras/ArbolB"
	"Fase2/estructuras/tablaHash"
)

var tablaAlumnos *tablaHash.TablaHash

func main() {
	tablaAlumnos = &tablaHash.TablaHash{Tabla: make(map[int])}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", validar)
	app.listen(":4000")

}

func Validar(c *fiber.Ctx) error {
	var usuario Peticiones.PeticionLogin
	c.BodyParser(&usuario)
	if usuario.UserName == "2012" {
		if usuario.Password == "admin"{
			return c.JSON(&fiber.Map{
				"status": 200,
				"message": "login correcto"
				"rol": 1
			})
		}
	} else {
		if usuario.Tutor {
			// buscar arbol b
		} else {
			// buscar en tabla hash
			if tablaAlumnos.Buscar(usuario.UserName, usuario.Password) {
				return c.JSON(&fiber.Map{
					"status": 200,
					"message": "login correcto"
					"rol": 3,

				})
			}
		}
	}
	return c.JSON(&fiber.Map{
		"status": 400,
		"message": "login incorrecto"
	})
}