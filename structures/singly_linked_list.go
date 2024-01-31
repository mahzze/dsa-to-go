package structures

import "fmt"

type Node struct {
	Value int
	next  *Node
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func newSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func (sll SinglyLinkedList) Get(i int) (int, error) {
	if i > sll.Length {
		err := fmt.Errorf("Index (%v) is bigger than the list length (%v)!", i, sll.Length)
		return 0, err
	}
	if i == 0 {
		return sll.Head.Value, nil
	}
	if i == sll.Length+1 {
		return sll.Tail.Value, nil
	}

	cur := sll.Head
	for cur.next != nil {
		cur = cur.next
	}
	return cur.Value, nil
}

func (sll *SinglyLinkedList) Append(v int) {
	tmp := &Node{Value: v, next: nil}
	if sll.Length == 0 {
		sll.Head = tmp
		sll.Tail = tmp
	} else {
		sll.Tail.next = tmp
		sll.Tail = sll.Tail.next
	}

	sll.Length++
}

func (sll *SinglyLinkedList) Prepend(v int) {
	tmp := sll.Head
	sll.Head = &Node{Value: v, next: tmp}
	if sll.Length == 0 {
		sll.Tail = sll.Head
	}
	sll.Length++
}

func (sll SinglyLinkedList) InsertAt(i, v int) {
}

func (sll SinglyLinkedList) Remove() {
}

func (sll SinglyLinkedList) RemoveAt(i, v int) {
}
