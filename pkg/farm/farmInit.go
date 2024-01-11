package farm

/* Allocates an initial space for the farm */
func (farm *Farm) InitFarm() {
	farm.rooms = make(map[string]*Room)
	farm.distances = make(map[*Room]int)

}

/* Puts all the ants in the start room */
func (farm *Farm) InitAnts(ants_number int) {
	farm.number_of_ants = ants_number
	farm.ants = make([]*Ant, farm.number_of_ants)

	for i := 0; i < ants_number; i++ {
		farm.ants[i] = new(Ant)
		farm.ants[i].room = farm.start_room
		farm.ants[i].discovered_rooms = make(map[*Room]bool)
		farm.ants[i].discovered_rooms[farm.start_room] = false
		farm.ants[i].moving = true
		farm.ants[i].ant_number = i + 1
	}
}

func (farm *Farm) InitDistances() {
	for room_idx := range farm.rooms {
		farm.distances[farm.rooms[room_idx]] = 99999
	}
}

func (farm *Farm) InitTunnels() {
	for room_idx := range farm.rooms {
		farm.rooms[room_idx].tunnels = &LinkedRoomsList{}
		farm.rooms[room_idx].locked_tunnels = make(map[string]bool)
		farm.rooms[room_idx].dead_end = false
	}
}
