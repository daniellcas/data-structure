package main

import (
	"fmt"

	simple_linked_list "github.com/daniellcas/data-structure/simple-linked-list/list"
	"github.com/daniellcas/data-structure/simple-linked-list/node"
)

func main() {
	// Abrindo a lista vazia
	list := simple_linked_list.NewSimpleLinkedList[string](nil)
	mostrarConsole(list)

	//Adicionando node na lista
	node1 := node.NewNode[string]("Daniel", nil)
	list.Add(node1)
	mostrarConsole(list)

	//Adicionando mais um node na lista
	node2 := node.NewNode[string]("Ana", nil)
	list.Add(node2)
	mostrarConsole(list)

	//Removendo ultimo item da lista
	list.Pop()
	mostrarConsole(list)
}

func mostrarConsole(v *simple_linked_list.SimpleLinkedList[string]) {
	fmt.Println("----")
	fmt.Println(v.GetHead())
	fmt.Println(v.GetTail())
	fmt.Println(v.GetSize())
}
