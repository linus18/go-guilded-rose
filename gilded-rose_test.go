package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ATest_GildedRose(t *testing.T) {
	newItem := Item{"Our new item", 0, 5}
	//printSlice(items)
	items = append(items, newItem)
	//printSlice(items2)

	main()
	if items[len(items)-1].quality != 3 {
		t.Errorf("quality should've decresed twice as fast: %d", items[len(items)-1].quality)
	}
	main()
	main()
	printSlice(items)
	for i := 0; i < 100; i++ {
		fmt.Printf("loop %d...", i)
		main()
		printSlice(items)
	}
}

func Test_Brie(t *testing.T) {
	var brie Item = getItemFromSlice("Aged Brie")

	fmt.Println(brie)
	assert.Equal(t, 0, brie.quality, "should be two")

	main()

	assert.Equal(t, 1, getItemFromSlice("Aged Brie").quality, "should be one")

	main()

	assert.Equal(t, 2, getItemFromSlice("Aged Brie").quality, "should be one")

	main()

	assert.Equal(t, 4, getItemFromSlice("Aged Brie").quality, "should be one")
}

func getItemFromSlice(itemName string) Item {
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i].name)
		if items[i].name == itemName {
			return items[i]
		}

	}
	return Item{"", 0, 0}
}

func printSlice(s []Item) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
