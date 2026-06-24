package main

import (
	"os"
	"testing"
)

func TestCreateBranch(t *testing.T) {

	os.MkdirAll(".minigit/branches", 0755)

	createBranch("testing")

	if _, err := os.Stat(".minigit/branches/testing.txt"); os.IsNotExist(err) {
		t.Error("La rama no fue creada")
	}
}
