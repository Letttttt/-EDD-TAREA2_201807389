package main

import (
	"bufio"
	"fmt"
	"main/simple"
	"os"
	"strings"
)

func main() {
	decision := 0
	reader := bufio.NewReader(os.Stdin)
	lista := simple.ListaSimple{}

	for decision != 3 {
		fmt.Println("Bienvenido")
		fmt.Println("Seleccione la opcion:")
		fmt.Println("1. Crear Usuario")
		fmt.Println("2. Modificar Usuario")
		fmt.Println("3. Salir")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		_, err := fmt.Sscan(text, &decision)
		if err != nil {
			fmt.Println("Lo que ingresó no es un número.")
		}
		switch decision {
		case 1:
			fmt.Println("Ingrese nombre: ")
			textnombre, _ := reader.ReadString('\n')
			textnombre = strings.TrimSpace(textnombre)
			fmt.Println("Ingrese apellido: ")
			textapellido, _ := reader.ReadString('\n')
			textapellido = strings.TrimSpace(textapellido)
			fmt.Println("Ingrese edad: ")
			textedad, _ := reader.ReadString('\n')
			textedad = strings.TrimSpace(textedad)
			fmt.Println("Ingrese dirección: ")
			textdireccion, _ := reader.ReadString('\n')
			textdireccion = strings.TrimSpace(textdireccion)
			fmt.Println("Ingrese carrera: ")
			textcarrera, _ := reader.ReadString('\n')
			textcarrera = strings.TrimSpace(textcarrera)
			fmt.Println("Ingrese curso: ")
			textcurso, _ := reader.ReadString('\n')
			textcurso = strings.TrimSpace(textcurso)
			fmt.Println("Ingrese estado: ")
			textestado, _ := reader.ReadString('\n')
			textestado = strings.TrimSpace(textestado)
			if VerificarTexto(textnombre) && VerificarTexto(textapellido) && VerificarTexto(textedad) && VerificarTexto(textdireccion) && VerificarTexto(textcarrera) && VerificarTexto(textcurso) && VerificarTexto(textestado) {
				if strings.EqualFold(textestado, "activo") || strings.EqualFold(textestado, "inactivo") {
					lista.InsertarUsuario(textnombre, textapellido, textedad, textdireccion, textcarrera, textcurso, textestado)
					fmt.Println("Usuario se ha ingresado correctamente.")
				} else {
					fmt.Println("No existe ese estado.")
				}
			} else {
				fmt.Println("Verifique los datos ingresados.")
			}
		case 2:
			if lista.Vacio() {
				fmt.Println("Esta vacía la lista de usuarios.")
			} else {
				lista.MostrarUsuarios()
				fmt.Println("Elija al número del usuario que desee modificar: ")
				eleccion, _ := reader.ReadString('\n')
				eleccion = strings.TrimSpace(eleccion)
				numero := 0
				fmt.Sscan(eleccion, &numero)
				usuarioAux := lista.ModificarUsuario(numero)
				fmt.Println("Nombre: %s\n", usuarioAux)
				// Por qué imprime todos los valores solo usando la variable usuarioAux,
				// Por ejemplo: fmt.Println(usuarioAux)
				// pero al tratar de acceder a un atributo no lo hace?
				// fmt.Println("Nombre: %s\n", usuarioAux.apellido)
			}

		}
	}
}

func VerificarTexto(texto string) bool {
	return len([]rune(texto)) > 0
}
