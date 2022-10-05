package main

type Coords interface {
	Width() int
	Height() int
	Left() int
	Right() int
	Top() int
	Bottom() int
	MouseX() int
	MouseY() int
}

func (g *Game) Width() int {
	return g.width
}

func (g *Game) Height() int {
	return g.height
}

func (g *Game) Left() int {
	return g.left
}

func (g *Game) Right() int {
	return g.right
}

func (g *Game) Top() int {
	return g.top
}

func (g *Game) Bottom() int {
	return g.bottom
}

func (g *Game) MouseX() int {
	return g.mouseX
}

func (g *Game) MouseY() int {
	return g.mouseY
}
