package refactored

import "testing"

const (
	totalIterations = 1000
	cheese          = 1
	sulfuras        = 3
	backstage       = 4
	conjured        = 5
)

func TestGuildedRoseNormal(t *testing.T) {
	normal := []int{0, 2}
	for _, idx := range normal {
		it := items
		sell := items[idx].sellIn
		q := items[idx].quality
		for i := 0; i < totalIterations; i++ {
			it = GildedRose(it)
			//update sell-in days
			sell--
			if it[idx].sellIn != sell {
				t.Errorf("sell-in was %d but should be %d", it[idx].sellIn, sell)
			}
			//update expected quality
			if it[idx].sellIn >= 0 {
				q--
			} else {
				q -= 2
			}
			if q < 0 {
				q = 0
			}
			if it[idx].quality != q {
				t.Errorf("quality was %d but should be %d", it[idx].quality, q)
			}
		}
	}
}
func TestGuildedRoseBrie(t *testing.T) {
	it := items
	sell := items[cheese].sellIn
	q := items[cheese].quality
	for i := 0; i < totalIterations; i++ {
		it = GildedRose(it)
		t.Log(it[cheese])
		//update sell-in days
		sell--
		if it[cheese].sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", it[cheese].sellIn, sell)
		}
		//update expected quality
		if it[cheese].sellIn >= 0 {
			q++
		} else {
			q += 2
		}
		if q > 50 {
			q = 50
		}
		if it[cheese].quality != q {
			t.Errorf("quality was %d but should be %d", it[cheese].quality, q)
		}
	}
}
func TestGuildedRoseSulfuras(t *testing.T) {
	it := items
	sell := items[sulfuras].sellIn
	q := items[sulfuras].quality
	for i := 0; i < totalIterations; i++ {
		it = GildedRose(it)
		t.Log(it[sulfuras])
		if it[sulfuras].sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", it[sulfuras].sellIn, sell)
		}
		if it[sulfuras].quality != q {
			t.Errorf("quality was %d but should be %d", it[sulfuras].quality, q)
		}
	}
}
func TestGuildedRoseBackstage(t *testing.T) {
	it := items
	sell := items[backstage].sellIn
	q := items[backstage].quality
	for i := 0; i < totalIterations; i++ {
		it = GildedRose(it)
		t.Log(it[backstage])
		//update sell-in days
		sell--
		if it[backstage].sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", it[backstage].sellIn, sell)
		}
		//update expected quality
		if it[backstage].sellIn > 9 {
			q++
		} else if it[backstage].sellIn > 4 {
			q += 2
		} else if it[backstage].sellIn >= 0 {
			q += 3
		} else {
			q = 0
		}
		if q > 50 {
			q = 50
		}
		if it[backstage].quality != q {
			t.Errorf("quality was %d but should be %d", it[backstage].quality, q)
		}
	}
}
func TestGuildedRoseConjured(t *testing.T) {
	it := items
	sell := items[conjured].sellIn
	q := items[conjured].quality
	for i := 0; i < totalIterations; i++ {
		it = GildedRose(it)
		//update sell-in days
		sell--
		if it[conjured].sellIn != sell {
			t.Errorf("sell-in was %d but should be %d", it[conjured].sellIn, sell)
		}
		//update expected quality
		if it[conjured].sellIn >= 0 {
			q -= 2
		} else {
			q -= 4
		}
		if q < 0 {
			q = 0
		}
		if it[conjured].quality != q {
			t.Errorf("quality was %d but should be %d", it[conjured].quality, q)
		}
	}
}
