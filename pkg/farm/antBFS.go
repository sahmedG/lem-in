package farm

import (
	"fmt"
	"os"
)

func (farm *Farm) AntBFS() {
	queue := &Queue{}
	discovered := make(map[*Room]bool)
	visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()
	farm.distances[room] = 1
	visited[room] = true // added to visited list

	for room != nil {
		tunnel := room.tunnels.head
		for tunnel != nil {
			if !visited[tunnel.room] {

				if !tunnel.room.start {
					queue.Enqueue(tunnel.room)
				}
				visited[tunnel.room] = true
				if farm.distances[room]+1 < farm.distances[tunnel.room] {
					farm.distances[tunnel.room] = farm.distances[room] + 1
				}

			}
			if tunnel.room == farm.start_room {
				discovered[tunnel.room] = true
			}
			tunnel = tunnel.next
		}
		//fmt.Println(room.name)
		discovered[room] = true
		room = queue.Dequeue()

	}
	if !discovered[farm.start_room] {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
}
