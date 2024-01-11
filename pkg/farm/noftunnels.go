package farm

/* Counts number of tunnels a room is connected to. (Counts tunnels that are going to the end room */
func (farm *Farm) Number_of_Tunnels(room *Room) int {

	counter := 0

	tunnel := room.tunnels.head

	for tunnel != nil {
		if farm.distances[tunnel.room] < farm.distances[room] {
			counter++
		}
		tunnel = tunnel.next
	}

	return counter
}
