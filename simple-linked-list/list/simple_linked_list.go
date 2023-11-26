package simple_linked_list

import (
	"github.com/daniellcas/data-structure/simple-linked-list/node"
)

type simpleLinkedListInputTypes interface {
	string | int | float64 | bool
}

type SimpleLinkedList[v simpleLinkedListInputTypes] struct {
	head *node.Node[v]
	tail *node.Node[v]
	size int
}

func NewSimpleLinkedList[v simpleLinkedListInputTypes](head *node.Node[v]) *SimpleLinkedList[v] {
	return &SimpleLinkedList[v]{
		head: head,
		tail: nil,
		size: 0,
	}
}

// Adicionar node na ultima posicao da lista
func (s *SimpleLinkedList[v]) Add(nodeInput *node.Node[v]) {
	// Validando caso ainda nao exista a primeira posicao
	if s.head == nil {
		s.head = nodeInput
		return
	}

	// Pegando o ultimo da lista atual e apontando para o novo ultimo
	ultimoNode := s.head
	for i := 0; i < s.size; i++ {
		ultimoNode = ultimoNode.GetNext()
	}
	ultimoNode.SetNext(nodeInput)

	// Definindo o node de entrada como o ultimo
	s.tail = nodeInput

	//Aumentando tamanho da lista
	s.size += 1
}

// Remove node da ultima posicao da lista
func (s *SimpleLinkedList[v]) Pop() {
	//Pegar o penultimo elemento
	penultimoNode := s.head
	for i := 0; i < s.size-1; i++ {
		penultimoNode = penultimoNode.GetNext()
	}
	penultimoNode.SetNext(nil)

	//Definindo o penultimo node como ultimo
	s.tail = penultimoNode

	//Diminuindo tamanho da lista
	s.size -= 1
}

func (s *SimpleLinkedList[v]) GetHead() *node.Node[v] {
	return s.head
}
func (s *SimpleLinkedList[v]) GetTail() *node.Node[v] {
	return s.tail
}
func (s *SimpleLinkedList[v]) GetSize() int {
	return s.size
}
