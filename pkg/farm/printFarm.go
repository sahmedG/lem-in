package farm

import "fmt"

/* Prints the farm and rooms connections */
func (farm *Farm) PrintFarm() {
	fmt.Println("Rooms connections:")
	for key, room := range farm.rooms {
		fmt.Printf("%s ->", key)

		tunnel := room.tunnels.head
		for tunnel != nil {
			fmt.Printf(" + %s", tunnel.room.name)
			tunnel = tunnel.next
		}
		fmt.Print("\n")
	}
}
