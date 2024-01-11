package farm

import "fmt"

/* Add a room to the farm */
func (farm *Farm) AddRoom(name string, start_end_normal string, pos_x int, pos_y int) {

	switch start_end_normal {
	case "start":
		farm.rooms[name] = &Room{
			name:           name,
			start:          true,
			end:            false,
			is_empty:       true,
			pos_x:          pos_x,
			pos_y:          pos_y,
			tunnels:        &LinkedRoomsList{},
			locked_tunnels: make(map[string]bool),
		}
		farm.start_room = farm.rooms[name]

	case "end":
		farm.rooms[name] = &Room{
			name:           name,
			start:          false,
			end:            true,
			is_empty:       true,
			pos_x:          pos_x,
			pos_y:          pos_y,
			tunnels:        &LinkedRoomsList{},
			locked_tunnels: make(map[string]bool),
		}
		farm.end_room = farm.rooms[name]

	case "normal":
		farm.rooms[name] = &Room{
			name:           name,
			start:          false,
			end:            false,
			is_empty:       true,
			pos_x:          pos_x,
			pos_y:          pos_y,
			tunnels:        &LinkedRoomsList{},
			locked_tunnels: make(map[string]bool),
		}

	default:
		fmt.Println("Enter room type")
	}
}
