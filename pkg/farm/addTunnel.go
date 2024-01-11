package farm

import "fmt"

/* Joins two rooms together */
func (farm *Farm) AddTunnel(from_room string, to_room string, bi_direction bool) bool {
	if farm.rooms[from_room] == nil || farm.rooms[to_room] == nil {
		//fmt.Println("One room or two wasn't found!")
		fmt.Println("ERROR: invalid data format")
		return false
	}
	if farm.rooms[from_room].Tunnel_Exists(farm.rooms[to_room]) {
		return false
	}
	if farm.rooms[to_room].Tunnel_Exists(farm.rooms[from_room]) {
		return false
	}
	if bi_direction {
		farm.rooms[from_room].tunnels.AddToList(farm.rooms[to_room])
		farm.rooms[to_room].tunnels.AddToList(farm.rooms[from_room])
	} else {
		farm.rooms[from_room].tunnels.AddToList(farm.rooms[to_room])
	}
	return true
}
