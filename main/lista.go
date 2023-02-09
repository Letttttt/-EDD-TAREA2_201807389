package main

import "fmt"

type ListaSimple struct {
	cabeza *Node
}

func (lista *ListaSimple) InsertarUsuario(nombre string, apellido string, edad int, direccion string, carrera string, curso string, estado string) {
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
	for temp != nil {
		fmt.Println("%d, ", temp.nombre, temp.apellido, temp.edad, temp.direccion, temp.carrera, temp.curso, temp.estado)
		temp = temp.siguiente
	}
	//fmt.Println("%d, " , temp.nombre, temp.apellido, temp.edad, temp.direccion, temp.carrera, temp.curso, temp.estado)

}

func (lista *ListaSimple) ModificarUsuario(opcion int) {
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
