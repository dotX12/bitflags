package main

import (
	"fmt"
	"github.com/dotX12/bitflags"
)

func main() {
	flags := []string{"flag1", "flag2", "flag3"}
	fs := bitflags.NewFlagSetFromSlice(flags)

	err := fs.SetByName("flag1")
	if err != nil {
		fmt.Println("failed to set flag:", err)
	}

	has, err := fs.HasByName("flag1")
	if err != nil {
		fmt.Println("failed to check flag:", err)
	}
	fmt.Println("flag1 set:", has)

	err = fs.SetByName("flag1")
	if err != nil {
		fmt.Println("failed to clear flag:", err)
	}

	has, err = fs.HasByName("flag1")
	if err != nil {
		fmt.Println("failed to check flag:", err)
	}
	fmt.Println("flag1 set:", has) // Output: flag1 set: false

	err = fs.ToggleByName("flag2")
	if err != nil {
		fmt.Println("failed to toggle flag:", err)
	}

	fmt.Println("current flags:", fs.GetActiveFlags())

	err = fs.ToggleByName("flag2")
	if err != nil {
		fmt.Println("failed to toggle flag:", err)
	}

	fmt.Println("current flags:", fs.GetActiveFlags())

}
