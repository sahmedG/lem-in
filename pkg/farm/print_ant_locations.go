package farm

import (
	"fmt"
	"sort"
)

/*
Prints the location of each ant
*/
func (farm *Farm) Print_Ants_Locations() string {
	locations_as_string := ""
	ants_to_print := make([]*Ant, farm.number_of_ants)
	copy(ants_to_print, farm.ants)
	sort.SliceStable(ants_to_print, func(i, j int) bool {
		return ants_to_print[i].ant_number < ants_to_print[j].ant_number
	})

	for _, ant := range ants_to_print {
		if ant.moved {
			//fmt.Printf("L%d-%s ", i+1, ant.room.name)
			locations_as_string += fmt.Sprintf("L%d-%s ", ant.ant_number, ant.room.name)

			ant.moved = false
		}
	}

	return locations_as_string + "\n"
}
