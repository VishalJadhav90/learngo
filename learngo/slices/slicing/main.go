package main

import s "github.com/inancgumus/prettyslice"

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("items", items)

	top3 := items[0:3]
	s.Show("top3", top3)
	top3[2] = "tootoo"
	s.Show("items", items)

	last4 := items[9:]
	s.Show("last4", last4)

	top1 := top3[:1]
	s.Show("top1", top1)

	size := 4
	for beg := 0; beg < len(items); beg = beg + size {
		to := beg + size
		if to > len(items) {
			to = len(items)
		}
		slc := items[beg:to]
		s.Show("", slc)
	}
}
