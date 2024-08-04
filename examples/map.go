package main

import (
	"fmt"
	"github.com/dotX12/bitflags"
)

func main() {

	flags := map[string]uint{
		"flag1":  1 << 0,  // 1
		"flag2":  1 << 1,  // 2
		"flag3":  1 << 2,  // 4
		"flag4":  1 << 3,  // 8
		"flag5":  1 << 4,  // 16
		"flag6":  1 << 5,  // 32
		"flag7":  1 << 6,  // 64
		"flag8":  1 << 7,  // 128
		"flag9":  1 << 8,  // 256
		"flag10": 1 << 9,  // 512
		"flag11": 1 << 10, // 1024
		"flag12": 1 << 11, // 2048
		"flag13": 1 << 12, // 4096
		"flag14": 1 << 13, // 8192
		"flag15": 1 << 14, // 16384
		"flag16": 1 << 15, // 32768
		"flag17": 1 << 16, // 65536
		"flag18": 1 << 17, // 131072
		"flag19": 1 << 18, // 262144
		"flag20": 1 << 19, // 524288
		"flag21": 1 << 20, // 1048576
		"flag22": 1 << 21, // 2097152
		"flag23": 1 << 22, // 4194304

	}
	fs := bitflags.NewFlagSetFromMap(flags)

	if err := fs.SetByValue(6); err != nil {
		fmt.Println("failed to set flags:", err)
	}
	currentFlags := fs.GetActiveFlags()
	fmt.Println("current flags:", currentFlags)

	hasFlag2, _ := fs.HasAnyByName("flag2")
	fmt.Println("has flag2:", hasFlag2)
	hasFlag3, _ := fs.HasAllByName("flag3")
	fmt.Println("has flag3:", hasFlag3)

	fmt.Println("current flag value:", fs.String())

	if err := fs.ToggleByName("flag2"); err != nil {
		fmt.Println("failed to toggle flag:", err)
	}
	currentFlags = fs.GetActiveFlags()
	fmt.Println("current flags:", currentFlags)
	fmt.Println("current flag value:", fs.String())

	if err := fs.ToggleByName("flag23"); err != nil {
		fmt.Println("failed to toggle flag:", err)
	}

	currentFlags = fs.GetActiveFlags()
	fmt.Println("current flags:", currentFlags)
	fmt.Println("current flag value:", fs.String())
	fmt.Println("current flag value int:", fs.GetValue())

	fs2 := bitflags.NewFlagSetFromMap(flags)
	if err := fs2.SetByValue(4194308); err != nil {
		fmt.Println("failed to set flags:", err)
	}

	currentFlags = fs2.GetActiveFlags()
	fmt.Println("current flags:", currentFlags)
	fmt.Println("current flag value:", fs2.String())
	fmt.Println("current flag value int:", fs2.GetValue())

}
