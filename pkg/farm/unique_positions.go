package farm

func (farm Farm) Unique_Positions() bool {
	for _, room := range farm.rooms {
		for _, next_room := range farm.rooms {
			if room.name != next_room.name {
				if room.pos_x == next_room.pos_x && room.pos_y == next_room.pos_y {
					return false
				}
			}
		}
	}
	return true
}
