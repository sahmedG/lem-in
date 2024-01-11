package farm

type Room struct {
	tunnels        *LinkedRoomsList
	start          bool
	end            bool
	is_empty       bool
	pos_x          int
	pos_y          int
	name           string
	locked_tunnels map[string]bool
	dead_end       bool
}

type Ant struct {
	ant_number       int
	room             *Room
	discovered_rooms map[*Room]bool
	moving           bool
	moved            bool
	check_again      bool
}

type Farm struct {
	rooms          map[string]*Room
	number_of_ants int
	distances      map[*Room]int
	ants           []*Ant
	start_room     *Room
	end_room       *Room
}

func (farm *Farm) Get_start_room_tunnels() *RoomNode {
	return farm.start_room.tunnels.head
}
