package farm

import "fmt"

/* Prints the farm and rooms connections */
func (farm *Farm) PrintDistances() {
	fmt.Println("Rooms distances:")
	for key, room := range farm.rooms {
		fmt.Printf("%s -> %d", key, farm.distances[room])
		fmt.Print("\n")
	}
}
