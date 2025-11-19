package utils

import (
	"strings"
)

type Canvas struct{ cells [][]rune }

func NewCanvas(w, h int) *Canvas {
	c := &Canvas{cells: make([][]rune, h)}
	for i := range c.cells {
		c.cells[i] = make([]rune, w)
		for j := range c.cells[i] {
			c.cells[i][j] = ' '
		}
	}
	return c
}

func (c *Canvas) Blit(x, y int, lines []string) {
	for dy, line := range lines {
		if y+dy < 0 || y+dy >= len(c.cells) {
			continue
		}
		row := c.cells[y+dy]
		for dx, r := range []rune(line) {
			if x+dx < 0 || x+dx >= len(row) {
				continue
			}
			row[x+dx] = r
		}
	}
}

func (c *Canvas) BlitTransparent(x, y int, lines []string, transparent rune) {
	for dy, line := range lines {
		ry := y + dy
		if ry < 0 || ry >= len(c.cells) {
			continue
		}
		row := c.cells[ry]
		for dx, r := range []rune(line) {
			rx := x + dx
			if rx < 0 || rx >= len(row) {
				continue
			}
			if r == transparent {
				continue
			}
			row[rx] = r
		}
	}
}

func (c *Canvas) String() string {
	out := make([]string, len(c.cells))
	for i, row := range c.cells {
		out[i] = string(row)
	}
	return strings.Join(out, "\n")
}
