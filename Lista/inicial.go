package inicial

import (
	"fmt"
)

// Declaración de la estructura
type InicialList struct {
	head *Node
}

func (list *InicialList) Insert(value int) {
	newNode := &Node{value: value, next: nil}
	if list.head == nil {
		list.head = newNode
	} else {
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

func (list *InicialList) Print() {
	temp := list.head
	for temp.next != nil {
		fmt.Printf("%d, ", temp.value)
		temp = temp.next
	}
	fmt.Printf("%d\n", temp.value)
}

func (list *InicialList) Find(value int) (result bool, index int) {
	temp := list.head
	index = 0
	result = false
	for temp != nil {
		if temp.value == value {
			result = true
		}
		index++
		temp = temp.next
	}

	if !result {
		index = -1
	}

	return result, index
}

func (list *InicialList) Delete(value int) (result bool) {
	result = false
	// Si es el primero de la lista
	if list.head.value == value {
		result = true
		if list.head.next == nil {
			list.head = nil
		} else {
			list.head = list.head.next
		}
	} else {
		temp := list.head
		for temp != nil {
			if temp.next != nil {
				if temp.next.value == value {
					break
				}
			}
			temp = temp.next
		}
		if temp == nil {
			return false
		} else {
			temp.next = temp.next.next
			result = true
		}
	}
	return result
}
