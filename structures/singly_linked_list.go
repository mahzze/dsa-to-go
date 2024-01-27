package structures

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type SinglyLinkedList struct {
	Head   Node
	Tail   Node
	Length int
}

func (sll SinglyLinkedList) Get(i int) int {
	if i == 0 {
		return sll.Head.Value
	}
	if i == sll.Length+1 {
		return sll.Tail.Value
	}

	cur := &sll.Head
	for k := 0; k < i; k++ {
		cur = cur.next
	}
	return cur.Value
}

func (sll SinglyLinkedList) Append(v int) {
	tmp := Node{Value: v, prev: &sll.Tail}
	sll.Tail.next = &tmp
	sll.Length++
}

func (sll SinglyLinkedList) Prepend(v int) {
	tmp := Node{Value: v, next: &sll.Head}
	sll.Head.next = &tmp
	sll.Length++
}

func (sll SinglyLinkedList) InsertAt(i, v int) {
}

func (sll SinglyLinkedList) Remove() {
}

func (sll SinglyLinkedList) RemoveAt(i, v int) {
}
