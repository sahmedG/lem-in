package farm

func (farm *Farm) SetAnts(ants []*Ant) {
	farm.ants = ants

	for ant_idx := 0; ant_idx < len(ants); ant_idx++ {
		farm.ants[ant_idx].room = farm.rooms["Room1"]
	}
}
