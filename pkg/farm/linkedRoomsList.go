package farm

type RoomNode struct {
	room *Room
	next *RoomNode
}

type LinkedRoomsList struct {
	head *RoomNode
}

func (list *LinkedRoomsList) AddToList(new_room *Room) {
	newNode := &RoomNode{room: new_room, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
