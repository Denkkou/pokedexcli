package main

import (
	"fmt"
)

func commandExplore(config *config, areaName string) error {
	fmt.Printf("Exploring: %s\n", areaName)
	return nil
}