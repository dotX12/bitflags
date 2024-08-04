# bitflags

`bitflags` is a Go library for managing sets of flags.

## Installation

```bash
go get github.com/dotX12/bitflags
```

## Example slice

```go
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
```

## Example map

```go
package main

import (
	"fmt"
	"github.com/dotX12/bitflags"
)

func main() {
	flags := map[string]uint{
		"flag1": 1 << 0, // 1
		"flag2": 1 << 1, // 2
		"flag3": 1 << 2, // 4
		"flag4": 1 << 3, // 8
		"flag5": 1 << 4, // 16
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
}
```