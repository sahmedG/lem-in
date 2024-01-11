package farm

import "fmt"

func (room *Room) Tunnel_Exists(to_room *Room) bool {

	tunnel := room.tunnels.head
	for tunnel != nil {
		if tunnel.room.name == to_room.name {
			//fmt.Printf("Tunnel %s -> %s already exist\n", room.name, to_room.name)
			fmt.Println("ERROR: invalid data format")
			return true
		}
		tunnel = tunnel.next
	}
	return false
}
