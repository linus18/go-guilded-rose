package refactored

import "testing"

var (
	normal    = []*Item{&items[0], &items[2]}
	brie      = &items[1]
	sulfuras  = &items[3]
	backstage = &items[4]
	conjured  = &items[5]
)

const totalIterations = 1000

func TestGuildedRoseNormal(t *testing.T) {
	for _, item := range normal {
		t.Log(item)
		sell := item.sellIn
		q := item.quality
		for i := 0; i < totalIterations; i++ {
			GildedRose()
			//update sell-in days
			sell--
			if item.sellIn != sell {
				t.Errorf("sell-in was %d but should be %d", item.sellIn, sell)
			}
			//update expected quality
			if item.sellIn >= 0 {
				q--
			} else {
				q -= 2
			}
			if q < 0 {
				q = 0
			}
			if item.quality != q {
				t.Errorf("quality was %d but should be %d", item.quality, q)
			}
		}
	}
}
func TestGuildedRoseBrie(t *testing.T) {
	sell := brie.sellIn
	q := brie.quality
	for i := 0; i < totalIterations; i++ {
		GildedRose()
		t.Log(brie)
		//update sell-in days
		sell--
		if brie.sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", brie.sellIn, sell)
		}
		//update expected quality
		if brie.sellIn >= 0 {
			q++
		} else {
			q += 2
		}
		if q > 50 {
			q = 50
		}
		if brie.quality != q {
			t.Errorf("quality was %d but should be %d", brie.quality, q)
		}
	}
}
func TestGuildedRoseSulfuras(t *testing.T) {
	sell := sulfuras.sellIn
	q := sulfuras.quality
	for i := 0; i < totalIterations; i++ {
		GildedRose()
		t.Log(sulfuras)
		if sulfuras.sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", brie.sellIn, sell)
		}
		if sulfuras.quality != q {
			t.Errorf("quality was %d but should be %d", sulfuras.quality, q)
		}
	}
}
func TestGuildedRoseBackstage(t *testing.T) {
	sell := backstage.sellIn
	q := backstage.quality
	for i := 0; i < totalIterations; i++ {
		GildedRose()
		t.Log(backstage)
		//update sell-in days
		sell--
		if backstage.sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", backstage.sellIn, sell)
		}
		//update expected quality
		if backstage.sellIn > 9 {
			q++
		} else if backstage.sellIn > 4 {
			q += 2
		} else if backstage.sellIn >= 0 {
			q += 3
		} else {
			q = 0
		}
		if q > 50 {
			q = 50
		}
		if backstage.quality != q {
			t.Errorf("quality was %d but should be %d", backstage.quality, q)
		}
	}
}
func TestGuildedRoseConjured(t *testing.T) {
	sell := conjured.sellIn
	q := conjured.quality
	for i := 0; i < totalIterations; i++ {
		GildedRose()
		//update sell-in days
		sell--
		if conjured.sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", conjured.sellIn, sell)
		}
		//update expected quality
		if conjured.sellIn >= 0 {
			q -= 2
		} else {
			q -= 4
		}
		if q < 0 {
			q = 0
		}
		if conjured.quality != q {
			t.Errorf("quality was %d but should be %d", conjured.quality, q)
		}
	}
}
