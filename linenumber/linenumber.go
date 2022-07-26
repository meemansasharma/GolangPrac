/*
We are working on a security system for a badged-access room in our company's building.

Given an ordered list of employees who used their badge to enter or exit the room, write a function that returns two collections:

1. All employees who didn't use their badge while exiting the room - they recorded an enter without a matching exit. (All employees are required to leave the room before the log ends.)

2. All employees who didn't use their badge while entering the room - they recorded an exit without a matching enter. (The room is empty when the log begins.)

Each collection should contain no duplicates, regardless of how many times a given employee matches the criteria for belonging to it.

records1 = [
  ["Martha",   "exit"],
  ["Paul",     "enter"],
  ["Martha",   "enter"],
  ["Steve",    "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "enter"],
  ["Paul",     "enter"],
  ["Curtis",   "exit"],
  ["Curtis",   "enter"],
  ["Paul",     "exit"],
  ["Martha",   "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "exit"],
  ["Paul",     "enter"],
  ["Paul",     "enter"],
  ["Martha",   "exit"],
  ["Paul",     "enter"],
  ["Paul",     "enter"],
  ["Paul",     "exit"],
  ["Paul",     "exit"]
]

Expected output: ["Paul", "Curtis", "Steve"], ["Martha", "Curtis", "Paul"]

Other test cases:

records2 = [
  ["Paul", "enter"],
  ["Paul", "exit"],
]

Expected output: [], []

records3 = [
  ["Paul", "enter"],
  ["Paul", "enter"],
  ["Paul", "exit"],
  ["Paul", "exit"],
]

Expected output: ["Paul"], ["Paul"]

records4 = [
  ["Paul", "enter"],
  ["Paul", "exit"],
  ["Paul", "exit"],
  ["Paul", "enter"],
]

Expected output: ["Paul"], ["Paul"]

All Test Cases:
mismatches(records1) => ["Paul", "Curtis", "Steve"], ["Martha", "Curtis", "Paul"]
mismatches(records2) => [], []
mismatches(records3) => ["Paul"], ["Paul"]
mismatches(records4) => ["Paul"], ["Paul"]

n: length of the badge records array
*/

package main

import "fmt"

func main() {
	records1 := [][]string{
		[]string{"Martha", "exit"},
		[]string{"Paul", "enter"},
		[]string{"Martha", "enter"},
		[]string{"Steve", "enter"},
		[]string{"Martha", "exit"},
		[]string{"Jennifer", "enter"},
		[]string{"Paul", "enter"},
		[]string{"Curtis", "exit"},
		[]string{"Curtis", "enter"},
		[]string{"Paul", "exit"},
		[]string{"Martha", "enter"},
		[]string{"Martha", "exit"},
		[]string{"Jennifer", "exit"},
		[]string{"Paul", "enter"},
		[]string{"Paul", "enter"},
		[]string{"Martha", "exit"},
		[]string{"Paul", "enter"},
		[]string{"Paul", "enter"},
		[]string{"Paul", "exit"},
		[]string{"Paul", "exit"},
	}

	records2 := [][]string{
		[]string{"Paul", "enter"},
		[]string{"Paul", "exit"},
	}

	records3 := [][]string{
		[]string{"Paul", "enter"},
		[]string{"Paul", "enter"},
		[]string{"Paul", "exit"},
		[]string{"Paul", "exit"},
	}

	records4 := [][]string{
		[]string{"Paul", "enter"},
		[]string{"Paul", "exit"},
		[]string{"Paul", "exit"},
		[]string{"Paul", "enter"},
	}
	fmt.Println(counterFunc(records1))
	fmt.Println(counterFunc(records2))
	fmt.Println(counterFunc(records3))
	fmt.Println(counterFunc(records4))
}

func counterFunc(inp [][]string) ([]string, []string) {

	NoExitList := make([]string, 1)
	NoEntryList := make([]string, 1)

	counterMap := make(map[string]int)

	for _, v := range inp {
		person, ok := counterMap[v[0]]
		if ok {
			if v[1] == "entry" && person > 0 {
				counterMap[v[0]] = 1
				NoExitList = append(NoExitList, v[0])
			} else if v[1] == "exit" && person < 0 {
				counterMap[v[0]] = -1
				NoEntryList = append(NoEntryList, v[0])
			} else if v[1] == "entry" && person < 0 {
				counterMap[v[0]] = 1
			} else {
				counterMap[v[0]] = -1
			}
		} else {
			if v[1] == "entry" {
				counterMap[v[0]] = 1

			} else {
				counterMap[v[0]] = -1
				NoEntryList = append(NoEntryList, v[0])
			}
		}
		for key, val := range counterMap {
			if val > 0 {
				NoExitList = append(NoExitList, key)
			}
		}
	}
	return NoEntryList, NoExitList
}
