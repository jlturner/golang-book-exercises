package main
import "fmt"

func main () {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	
	fmt.Println(elements["Li"], elements)
	name, ok := elements["Un"]
	if ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("oh no")
	}

	more_elements := map[string]string{
		"Rh": "Rhodium",
		"Os": "Osmium",
	}
	fmt.Println(more_elements)
}
