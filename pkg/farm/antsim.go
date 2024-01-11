package farm

import (
	"fmt"
	"sort"
)

var (
	one_ant_moved = false
)

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step(toggle bool) {
	ants_to_work_on := make([]*Ant, farm.number_of_ants)
	copy(ants_to_work_on, farm.ants)

	/* Ants with less tunnels will go first */
	sort.SliceStable(ants_to_work_on, func(i, j int) bool {
		return farm.Number_of_Tunnels(ants_to_work_on[i].room) < farm.Number_of_Tunnels(ants_to_work_on[j].room)
	})

	/* Loop throgh each ant */
	for ant_idx := 0; ant_idx < len(ants_to_work_on); ant_idx++ {
		alt_tun := farm.Find_Min_Path(ants_to_work_on[ant_idx], toggle)

		/* Will cause the ant to hold in one place so that in the next step when the tunnel is empty, it will take it */
		/* Distance of steps is increased because one ant has taken this path, or will take it in the next step */
		if alt_tun != farm.end_room || ants_to_work_on[ant_idx].room.start {
			farm.distances[alt_tun]++
		}
		if check_moving_possiblity(ants_to_work_on[ant_idx], alt_tun) {

			ants_to_work_on[ant_idx].room.locked_tunnels[alt_tun.name] = true               // Lock the tunnel from beign used by other ant until step is finished
			ants_to_work_on[ant_idx].discovered_rooms[ants_to_work_on[ant_idx].room] = true // remember the current room
			ants_to_work_on[ant_idx].room.is_empty = true                                   // flag current room as empty
			ants_to_work_on[ant_idx].room = alt_tun                                         // go to next room
			alt_tun.is_empty = false                                                        // flag next room as not empty
			if ants_to_work_on[ant_idx].room.end {
				ants_to_work_on[ant_idx].moving = false
			}
			ants_to_work_on[ant_idx].moved = true
			one_ant_moved = true
		}

		if ant_idx == len(ants_to_work_on)-1 && one_ant_moved {
			ant_idx = 0
			one_ant_moved = false
		}
	}
	// for debuggin purposes
	//farm.PrintDistances()
	/*if farm.Ants_Overlap() {
		fmt.Println("Ants overlaps!!!")
		os.Exit(0)
	}*/
	farm.Unlock_Locked_Tunnel()
	farm.AntBFS()

}

/* Starts the simulator until all ants are at the end room */
func (farm *Farm) AntSim() {
	toggler_on_steps := 0
	toggler_off_steps := 0
	toggler_on_string := ""
	toggler_off_string := ""

	step := 0
	/*step_string := ""
	step := 0*/

	for !farm.Ants_At_End() {
		toggler_off_steps++
		//fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step(false)
		toggler_off_string += farm.Print_Ants_Locations()
	}

	farm.Re_InitAnts()
	for !farm.Ants_At_End() {
		step++
		toggler_on_steps++
		//fmt.Printf("\nAnts moves step %d:\n", Step)
		// stop the infinite loop
		if toggler_on_steps > toggler_off_steps {
			break
		}
		farm.AntSim_Step(true)
		toggler_on_string += farm.Print_Ants_Locations()
	}

	if toggler_off_steps == toggler_on_steps {
		fmt.Print(toggler_on_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_on_steps)
	} else if toggler_off_steps < toggler_on_steps {
		fmt.Print(toggler_off_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_off_steps)
	} else {
		fmt.Print(toggler_on_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_on_steps)
	}

	/*fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)*/
}

/* Custom iterations for Ant simulator */
/*func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	step_string := ""
	step := 0
	for ant := 0; ant <= farm.number_of_ants; ant++ {
		for !farm.Ants_At_End() && iteration < iter {
			step++
			fmt.Printf("\nAnts moves step %d:\n", step)
			farm.AntSim_Step()
			step_string += farm.Print_Ants_Locations()
			iteration++
		}
	}
	fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)
}*/

/* This function will find the minimum path the ant can take */
func (farm *Farm) Find_Min_Path(ant *Ant, toggle bool) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {

		if toggle {
			if farm.distances[tunnel.room] < min && !ant.discovered_rooms[tunnel.room] {
				temp = tunnel.room
				min = farm.distances[temp]

			}
		} else {
			if farm.distances[tunnel.room] <= min && !ant.discovered_rooms[tunnel.room] {
				temp = tunnel.room
				min = farm.distances[temp]

			}
		}

		tunnel = tunnel.next
	}
	return temp
}

/* Checks the possibility of the ant to take a specified tunnel */
func check_moving_possiblity(ant *Ant, tunnel *Room) bool {
	return (tunnel.is_empty || tunnel.end) && !ant.discovered_rooms[tunnel] && !tunnel.start && !ant.room.locked_tunnels[tunnel.name] && ant.moving && !tunnel.dead_end && !ant.moved
}

/* Counts the tunnels that have same distance */
func (farm *Farm) same_distance_tunnels(ant *Ant) int {
	tunnel := ant.room.tunnels.head
	similar_distance_count := 0

	for tunnel != nil {
		compare_tunnel := tunnel.next
		for compare_tunnel != nil {
			if farm.distances[tunnel.room] == farm.distances[compare_tunnel.room] && (farm.distances[tunnel.room] <= farm.distances[ant.room]) && !ant.discovered_rooms[tunnel.room] {
				similar_distance_count++
			}
			compare_tunnel = compare_tunnel.next
		}
		tunnel = tunnel.next
	}

	return similar_distance_count
}

func (farm *Farm) Ants_Overlap() bool {
	for _, ant := range farm.ants {
		for _, next_ant := range farm.ants {
			if ant.ant_number != next_ant.ant_number {
				if ant.room.name == next_ant.room.name && ant.room != farm.start_room && ant.room != farm.end_room && next_ant.room != farm.start_room && next_ant.room != farm.end_room {
					return true
				}
			}
		}
	}
	return false
}
