package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Item struct {
	InventoryId int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomId int    `json:"room_id"`
	Name   string `json:"name"`
}

// Items in the Meeting Room
func itemsInMeetingRoom(items []*Item) {
	var result []*Item
	for _, v := range items {
		if strings.ToLower(v.Placement.Name) == "meeting room" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Items in the Meeting Room")
	fmt.Println(string(data))
}

// All electronic devices
func allElectronicDevices(items []*Item) {
	var result []*Item
	for _, v := range items {
		if strings.ToLower(v.Type) == "electronic" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All electronic devices")
	fmt.Println(string(data))
}

// All furniture
func allFurnitureType(items []*Item) {
	var result []*Item
	for _, v := range items {
		if strings.ToLower(v.Type) == "furniture" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All furniture type")
	fmt.Println(string(data))
}

// All items with brown color
func allBrownColor(items []*Item) {
	var result []*Item
	for _, v := range items {
		for _, tag := range v.Tags {
			if strings.ToLower(tag) == "brown" {
				result = append(result, v)
				break
			}
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All items with brown color")
	fmt.Println(string(data))
}

func main() {
	jsonFile, err := os.Open("inventory_list.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success")

	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)

	var items []*Item

	err = json.Unmarshal(file, &items)
	if err != nil {
		panic(err)
	}

	itemsInMeetingRoom(items)
	allElectronicDevices(items)
	allFurnitureType(items)
	allBrownColor(items)
}
