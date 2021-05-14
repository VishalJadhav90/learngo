package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true
	var games []string
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)
	s.Show("another games", []string{})

	games = []string{"pacman", "mario", "terminal", "doom", "tetris"}
	s.Show("games", games)

	part := games[0:]
	s.Show("part", part)
	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	games = games[len(games):]
	s.Show("games", games)
}
