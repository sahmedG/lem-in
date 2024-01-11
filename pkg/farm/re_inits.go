package farm

/* Reset ants memory and puts them all in the starting room */
func (farm *Farm) Re_InitAnts() {

	for i := 0; i < farm.number_of_ants; i++ {
		farm.ants[i].room = farm.start_room

		for room_idx := range farm.ants[i].discovered_rooms {
			farm.ants[i].discovered_rooms[room_idx] = false
		}
		farm.ants[i].discovered_rooms[farm.start_room] = false
		farm.ants[i].moving = true
	}

	farm.Unlock_Locked_Tunnel()
}

/* Resets locked tunnels rooms */
func (farm *Farm) Unlock_Locked_Tunnel() {
	rooms := farm.rooms

	for room_idx := range rooms {
		room := rooms[room_idx]

		for locked_tunnel_idx := range room.locked_tunnels {
			room.locked_tunnels[locked_tunnel_idx] = false
		}
	}
}
