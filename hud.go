package main

import (
	"fmt"
	"math"
	"time"
	"unicode/utf8"

	"github.com/seredot/kepler-22t/color"
	"gonum.org/v1/gonum/interp"
)

func (g *Game) drawHud() {
	g.ResetStyle()

	w := float64(g.width - 1)
	rs := []float64{102, 223, 162, 75, 55}
	gs := []float64{44, 100, 86, 43, 30}
	bs := []float64{55, 100, 206, 137, 97}
	xs := []float64{0, w * 0.2, w * 0.5, w * 0.8, w}
	pr := interp.ClampedCubic{}
	pg := interp.ClampedCubic{}
	pb := interp.ClampedCubic{}
	pr.Fit(xs, rs)
	pg.Fit(xs, gs)
	pb.Fit(xs, bs)

	for x := 0; x < g.width; x++ {
		g.Background(color.NewColorIntRGB(
			uint64(pr.Predict(float64(x))),
			uint64(pg.Predict(float64(x))),
			uint64(pb.Predict(float64(x)))))
		g.PutChar(x, 0, ' ')
	}

	g.ResetStyle()

	var textX, textY int

	// Title
	g.DrawText(1, 0, "Kepler 22t")

	// Score
	score := fmt.Sprintf(" %d", g.score)
	textX = g.width - 1 - utf8.RuneCountInString(score)
	g.Foreground(color.ColorWhite)
	g.DrawText(textX, 0, score)
	textX--
	g.Foreground(color.ColorAlien)
	g.DrawText(textX, 0, "☠")
	g.ResetStyle()

	// Health
	health := fmt.Sprintf(" %d ", int(math.Round(g.health)))
	textX -= 1 + utf8.RuneCountInString(health)
	g.DrawText(textX, 0, health)
	textX--
	g.Foreground(color.ColorCrossRed)
	g.DrawText(textX, 0, "✚")
	g.ResetStyle()

	// Ammo
	ammo := fmt.Sprintf(" %d ", g.ammo)
	textX -= 1 + utf8.RuneCountInString(ammo)
	g.DrawText(textX, 0, ammo)
	textX--
	g.Foreground(color.ColorAmber)
	g.DrawText(textX, 0, "⁍")
	g.ResetStyle()

	// Game over
	if g.state == GameOver {
		var m string
		g.Background(color.ColorRedSpill)
		m = "      GAME OVER       "
		textX = (g.width - utf8.RuneCountInString(m)) / 2
		textY = (g.height / 2)
		g.DrawText(textX, textY, m)
		m = " press enter to start "
		textX = (g.width - utf8.RuneCountInString(m)) / 2
		textY++
		g.DrawText(textX, textY, m)
		g.ResetStyle()
	}
	// Stats
	g.DrawText(2, g.height-1, fmt.Sprintf("FPS %0.2f", float64(time.Second/g.deltaT)))
	// Debug log
	//g.drawText(22, g.height-1, fmt.Sprintf(" Color %d %x ", g.screen.Colors(), style.Hsl2Rgb(242, 26, 43)))
}
