package farm

/*
Checks if ants are stuck and can't move
It will check even the ant that previosly checked and it replied with no tunnels,
pherhaps change the structure to queue or stacks? linked lists?
*/
func (farm *Farm) Ants_At_End() bool {

	for _, ant := range farm.ants {
		if ant.room != farm.end_room {
			return false
		}
	}
	return true
}
