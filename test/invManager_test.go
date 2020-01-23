package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	tInv = Inventory{
		Inven: map[string][]Item{
			"meats":  []Item{{"chicken", true, "p"}, {"beef", false, "p"}},
			"fruits": []Item{{"apple", true, "p"}, {"orange", false, "p"}},
		},
	}
)

func TestCheck(t *testing.T) {
	b := tInv.check("chicken")
	if b != true {
		t.Fatalf("TestCheck failed, gave %s, got %t, expected %s", "chicken", b, "true")
	}
	b = tInv.check("liver")
	if b != false {
		t.Fatalf("TestCheck failed, gave %s, got %t, expected %s", "live", b, "false")
	}
}

func TestFind(t *testing.T) {
	str, _ := tInv.find("chicken")
	if str == "" {
		t.Fatalf("TestFind failed")
	}
	str, _ = tInv.find("liver")
	if str != "" {
		t.Fatalf("TestFind2 failed")
	}
}

func TestAdd(t *testing.T) {
	arry := []string{"meats", "chicken"}
	tInv.Add(arry)
	if tInv.Inven["meats"][2].Name != "chicken" {
		t.Fatalf("TestAdd failed")
	}
	tInv.Add(arry[1:])
	_, b := tInv.Inven[arry[1]]
	if !b {
		t.Fatalf("TestAdd2 failed")
	}
}

func TestRemove(t *testing.T) {
	arry := []string{"apple", "orange"}
	tInv.Remove(arry)
	if len(tInv.Inven["fruits"]) > 0 {
		fmt.Print(tInv.Inven["fruits"])
		t.Fatalf("TestRemove failed")
	}
}

func TestChangeKey(t *testing.T) {
	arry := []string{"fruits", "chicken"}

	tInv.ChangeKey(arry)
	if len(tInv.Inven["fruits"]) > 2 {
		t.Fatalf("TestChangeKey failed")
	}
	if len(tInv.Inven["meats"]) < 2 {
		t.Fatalf("TestChangeKey2 failed")
	}
}
func TestRemoveKey(t *testing.T) {
	tInv.RemoveKey("chicken")
	_, b := tInv.Inven["chicken"]
	if b {
		t.Fatalf("TestRemoveKey failed")
	}

}

func TestUpdateDate(t *testing.T) {
	date := time.Now().Format("Mon Jan 2")
	tInv.updateDate("chicken")
	if date != tInv.Inven["meats"][0].DateEntered {
		t.Fatalf("TestUpdateDate failed")
	}
}

func TestAddGrocery(t *testing.T) {
	arry := []string{"beef", "chicken"}
	tInv.AddGrocery(arry)
	if !tInv.Inven["meats"][0].ForceList {
		t.Fatalf("TestAddGrocery failed")
	}
	if !tInv.Inven["meats"][1].ForceList {
		t.Fatalf("TestAddGrocery2 failed")
	}

}

func TestRemoveGrocery(t *testing.T) {
	arry := []string{"beef", "chicken"}
	tInv.RemoveGrocery(arry)
	if tInv.Inven["meats"][0].ForceList {
		t.Fatalf("TestRemoveGrocery failed")
	}
	if tInv.Inven["meats"][1].ForceList {
		t.Fatalf("TestRemoveGrocery2 failed")
	}
}
