package main

import (
	ants "ants/pkg/farm"
	load_data "ants/pkg/farm_loading"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("USAGE: go run ./cmd <path_to_farm_file>")
		return
	}
	var farm ants.Farm
	farm.InitFarm()
	read_err, data := load_data.Read_Farm_File(args[0], &farm)
	if !read_err {
		return
	}
	if !farm.Unique_Positions() {
		fmt.Println("ERROR: invalid data format")
		return
	}
	farm.InitDistances()
	farm.AntBFS()
	fmt.Println(data)
	//for debugging
	//farm.PrintFarm()
	//farm.PrintDistances()
	farm.AntSim()
}
