package main

type Node struct {
	nombre    string
	apellido  string
	edad      int
	direccion string
	carrera   string
	curso     string
	estado    string
	siguiente *Node
}
