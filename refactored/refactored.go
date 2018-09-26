package refactored

import (
	"strings"
)

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

type strategy func(it Item) Item

func Cheese(i Item) Item {
	i.sellIn--
	if i.sellIn >= 0 {
		i.quality++
	} else {
		i.quality += 2
	}
	if i.quality > 50 {
		i.quality = 50
	}
	return i
}

func Noop(i Item) Item { return i }

func Tics(i Item) Item {
	i.sellIn--
	if i.sellIn > 9 {
		i.quality++
	} else if i.sellIn > 4 {
		i.quality += 2
	} else if i.sellIn >= 0 {
		i.quality += 3
	} else {
		i.quality = 0
	}
	if i.quality > 50 {
		i.quality = 50
	}
	return i
}

func Conjured(i Item) Item {
	i.sellIn--
	if i.sellIn >= 0 {
		i.quality -= 2
	} else {
		i.quality -= 4
	}
	if i.quality < 0 {
		i.quality = 0
	}
	return i
}

var m = map[string]strategy{
	"brie":      Cheese,
	"sulfuras":  Noop,
	"backstage": Tics,
	"conjured":  Conjured,
}

func factory(it Item) strategy {
	for key, value := range m {
		if strings.Contains(strings.ToLower(it.name), key) {
			return value
		}
	}
	return func(i Item) Item {
		i.sellIn--
		if i.sellIn >= 0 {
			i.quality--
		} else {
			i.quality -= 2
		}
		if i.quality < 0 {
			i.quality = 0
		}
		return i
	}
}

func GildedRose(items []Item) []Item {
	//copy
	tmp := make([]Item, len(items))
	for i := 0; i < len(items); i++ {
		s := factory(items[i])
		tmp[i] = s(items[i])
	}
	return tmp
}
