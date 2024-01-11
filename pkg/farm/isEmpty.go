package farm

func (room *Room) Is_Empty() bool {
	return room.is_empty
}

func (room *Room) Set_Empty() {
	room.is_empty = true
}

func (room *Room) Set_Not_Empty() {
	room.is_empty = false
}
