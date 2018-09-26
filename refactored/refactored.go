package refactored

import (
	"errors"
	"strings"

	. "github.com/yanatan16/itertools"
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

func ConvertToGeneric(items []Item) []interface{} {
	tmp := make([]interface{}, len(items))
	for i := range items {
		tmp[i] = items[i]
	}
	return tmp
}

func ConvertToItem(generic []interface{}) ([]Item, error) {
	tmp := make([]Item, len(generic))
	for i := 0; i < len(generic); i++ {
		f, ok := (generic[i].(Item))
		if ok {
			tmp[i] = f
		} else {
			return nil, errors.New("not item")
		}
	}
	return tmp, nil
}

func GildedRose(items []Item) []Item {

	mapper := func(i interface{}) interface{} {
		item, ok := (i.(Item))
		if ok {
			return factory(item)(item)
		}
		return errors.New("unknown type")
	}

	it := Map(mapper, New(ConvertToGeneric(items)...))

	newItems, _ := ConvertToItem(List(it))
	return newItems
}
