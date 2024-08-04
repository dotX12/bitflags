package bitflags

import (
	"testing"
)

func TestNewFlagSetFromMap(t *testing.T) {
	flagMap := map[string]uint64{"flag1": 1, "flag2": 2, "flag3": 4}
	fs := NewFlagSetFromMap(flagMap)

	if len(fs.flagMap) != 3 {
		t.Fatalf("expected 3 flags, got %d", len(fs.flagMap))
	}
}

func TestNewFlagSetFromSlice(t *testing.T) {
	flags := []string{"flag1", "flag2", "flag3"}
	fs := NewFlagSetFromSlice(flags)

	if len(fs.flagMap) != 3 {
		t.Fatalf("expected 3 flags, got %d", len(fs.flagMap))
	}

	if fs.flagMap["flag1"] != 1 || fs.flagMap["flag2"] != 2 || fs.flagMap["flag3"] != 4 {
		t.Fatalf("flag map values are incorrect")
	}
}

func TestFlagSet_GetValue(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})
	if fs.GetValue() != 0 {
		t.Fatalf("expected initial value to be 0, got %d", fs.GetValue())
	}
}

func TestFlagSet_SetByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})

	if err := fs.SetByName("flag1"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	if fs.flags != 1 {
		t.Fatalf("expected flags to be 1, got %d", fs.flags)
	}

	if err := fs.SetByName("unknown"); err == nil {
		t.Fatalf("expected error when setting unknown flag")
	}
}

func TestFlagSet_ClearByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})
	err := fs.SetByName("flag1")

	if err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	if err = fs.ClearByName("flag1"); err != nil {
		t.Fatalf("failed to clear flag: %v", err)
	}

	if fs.flags != 0 {
		t.Fatalf("expected flags to be 0, got %d", fs.flags)
	}

	if err := fs.ClearByName("unknown"); err == nil {
		t.Fatalf("expected error when clearing unknown flag")
	}
}

func TestFlagSet_HasByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})
	err := fs.SetByName("flag1")

	if err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	has, err := fs.HasByName("flag1")
	if err != nil || !has {
		t.Fatalf("expected flag1 to be set")
	}

	has, err = fs.HasByName("flag2")
	if err != nil || has {
		t.Fatalf("expected flag2 to be not set")
	}

	if _, err := fs.HasByName("unknown"); err == nil {
		t.Fatalf("expected error when checking unknown flag")
	}
}

func TestFlagSet_ToggleByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})

	if err := fs.ToggleByName("flag1"); err != nil {
		t.Fatalf("failed to toggle flag: %v", err)
	}

	if fs.flags != 1 {
		t.Fatalf("expected flags to be 1, got %d", fs.flags)
	}

	if err := fs.ToggleByName("flag1"); err != nil {
		t.Fatalf("failed to toggle flag: %v", err)
	}

	if fs.flags != 0 {
		t.Fatalf("expected flags to be 0, got %d", fs.flags)
	}

	if err := fs.ToggleByName("unknown"); err == nil {
		t.Fatalf("expected error when toggling unknown flag")
	}
}

func TestFlagSet_HasAnyByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2", "flag3"})
	err := fs.SetByName("flag1")
	if err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}
	hasAny, err := fs.HasAnyByName("flag1", "flag2")
	if err != nil || !hasAny {
		t.Fatalf("expected flag1 or flag2 to be set")
	}

	hasAny, err = fs.HasAnyByName("flag2", "flag3")
	if err != nil || hasAny {
		t.Fatalf("expected neither flag2 nor flag3 to be set")
	}

	if _, err := fs.HasAnyByName("unknown"); err == nil {
		t.Fatalf("expected error when checking unknown flag")
	}
}

func TestFlagSet_HasAllByName(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})

	if err := fs.SetByName("flag1"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}
	if err := fs.SetByName("flag2"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	hasAll, err := fs.HasAllByName("flag1", "flag2")
	if err != nil || !hasAll {
		t.Fatalf("expected flag1 and flag2 to be set")
	}

	_ = fs.ClearByName("flag2")
	hasAll, err = fs.HasAllByName("flag1", "flag2")
	if err != nil || hasAll {
		t.Fatalf("expected flag1 or flag2 to be not set")
	}

	if _, err := fs.HasAllByName("unknown"); err == nil {
		t.Fatalf("expected error when checking unknown flag")
	}
}

func TestFlagSet_SetByValue(t *testing.T) {
	flagMap := map[string]uint64{"flag1": 1, "flag2": 2, "flag3": 4}
	fs := NewFlagSetFromMap(flagMap)

	if err := fs.SetByValue(3); err != nil {
		t.Fatalf("failed to set flags by value: %v", err)
	}

	if fs.flags != 3 {
		t.Fatalf("expected flags to be 3, got %d", fs.flags)
	}

	if err := fs.SetByValue(8); err == nil {
		t.Fatalf("expected error when setting invalid flag value")
	}
}

func TestFlagSet_String(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})
	if err := fs.SetByName("flag1"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	expected := "00000001"
	if fs.String() != expected {
		t.Fatalf("expected string representation to be %s, got %s", expected, fs.String())
	}
}

func TestFlagSet_GetActiveFlags(t *testing.T) {
	fs := NewFlagSetFromSlice([]string{"flag1", "flag2"})
	if err := fs.SetByName("flag1"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	activeFlags := fs.GetActiveFlags()
	if len(activeFlags) != 1 || activeFlags["flag1"] != 1 {
		t.Fatalf("expected active flags to contain flag1")
	}
}
