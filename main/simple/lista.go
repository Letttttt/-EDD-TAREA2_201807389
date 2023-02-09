package simple

import "fmt"

type ListaSimple struct {
	cabeza *Node
}

func (lista *ListaSimple) InsertarUsuario(nombre string, apellido string, edad string, direccion string, carrera string, curso string, estado string) {
	nuevo := &Node{nombre: nombre, apellido: apellido, edad: edad, direccion: direccion, carrera: carrera, curso: curso, estado: estado, siguiente: nil}
	if lista.cabeza == nil {
		lista.cabeza = nuevo
	} else {
		temp := lista.cabeza
		for temp.siguiente != nil {
			temp = temp.siguiente
		}
		temp.siguiente = nuevo
	}
}

func (lista *ListaSimple) MostrarUsuarios() {
	temp := lista.cabeza
	tempOpcion := 1
	fmt.Println("Lista de usuarios:")
	for temp != nil {
		fmt.Println(tempOpcion, temp.nombre, temp.apellido, temp.edad, temp.direccion, temp.carrera, temp.curso, temp.estado)
		temp = temp.siguiente
		tempOpcion++
	}
	//fmt.Println("%d, " , temp.nombre, temp.apellido, temp.edad, temp.direccion, temp.carrera, temp.curso, temp.estado)

}

func (lista *ListaSimple) ModificarUsuario(opcion int) *Node {
	temp := lista.cabeza
	tempOpcion := 1

	for temp != nil {
		if tempOpcion == opcion {
			return temp
		}
		tempOpcion++
		temp = temp.siguiente
	}
	return nil
}

func (lista *ListaSimple) Vacio() bool {
	return lista.cabeza == nil

}
