// Copyright 2019 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build example
// +build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	emptyImage = ebiten.NewImage(3, 3)

	// emptySubImage is an internal sub image of emptyImage.
	// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func drawEbitenText(screen *ebiten.Image, x, y int, scale float32, line bool) {
	var path vector.Path

	// E
	path.MoveTo(20, 20)
	path.LineTo(20, 70)
	path.LineTo(70, 70)
	path.LineTo(70, 60)
	path.LineTo(30, 60)
	path.LineTo(30, 50)
	path.LineTo(70, 50)
	path.LineTo(70, 40)
	path.LineTo(30, 40)
	path.LineTo(30, 30)
	path.LineTo(70, 30)
	path.LineTo(70, 20)

	// B
	path.MoveTo(80, 20)
	path.LineTo(80, 70)
	path.LineTo(100, 70)
	path.QuadTo(150, 57.5, 100, 45)
	path.QuadTo(150, 32.5, 100, 20)

	// I
	path.MoveTo(140, 20)
	path.LineTo(140, 70)
	path.LineTo(150, 70)
	path.LineTo(150, 20)

	// T
	path.MoveTo(160, 20)
	path.LineTo(160, 30)
	path.LineTo(180, 30)
	path.LineTo(180, 70)
	path.LineTo(190, 70)
	path.LineTo(190, 30)
	path.LineTo(210, 30)
	path.LineTo(210, 20)

	// E
	path.MoveTo(220, 20)
	path.LineTo(220, 70)
	path.LineTo(270, 70)
	path.LineTo(270, 60)
	path.LineTo(230, 60)
	path.LineTo(230, 50)
	path.LineTo(270, 50)
	path.LineTo(270, 40)
	path.LineTo(230, 40)
	path.LineTo(230, 30)
	path.LineTo(270, 30)
	path.LineTo(270, 20)

	// N
	path.MoveTo(280, 20)
	path.LineTo(280, 70)
	path.LineTo(290, 70)
	path.LineTo(290, 35)
	path.LineTo(320, 70)
	path.LineTo(330, 70)
	path.LineTo(330, 20)
	path.LineTo(320, 20)
	path.LineTo(320, 55)
	path.LineTo(290, 20)

	var vs []ebiten.Vertex
	var is []uint16
	if line {
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	} else {
		vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	}

	for i := range vs {
		vs[i].DstX = (vs[i].DstX + float32(x)) * scale
		vs[i].DstY = (vs[i].DstY + float32(y)) * scale
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}

	op := &ebiten.DrawTrianglesOptions{}
	if !line {
		op.FillRule = ebiten.EvenOdd
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

func drawEbitenLogo(screen *ebiten.Image, x, y int, scale float32, line bool) {
	const unit = 16

	var path vector.Path

	// TODO: Add curves
	path.MoveTo(0, 4*unit)
	path.LineTo(0, 6*unit)
	path.LineTo(2*unit, 6*unit)
	path.LineTo(2*unit, 5*unit)
	path.LineTo(3*unit, 5*unit)
	path.LineTo(3*unit, 4*unit)
	path.LineTo(4*unit, 4*unit)
	path.LineTo(4*unit, 2*unit)
	path.LineTo(6*unit, 2*unit)
	path.LineTo(6*unit, 1*unit)
	path.LineTo(5*unit, 1*unit)
	path.LineTo(5*unit, 0)
	path.LineTo(4*unit, 0)
	path.LineTo(4*unit, 2*unit)
	path.LineTo(2*unit, 2*unit)
	path.LineTo(2*unit, 3*unit)
	path.LineTo(unit, 3*unit)
	path.LineTo(unit, 4*unit)

	var vs []ebiten.Vertex
	var is []uint16
	if line {
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	} else {
		vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	}

	for i := range vs {
		vs[i].DstX = (vs[i].DstX + float32(x)) * scale
		vs[i].DstY = (vs[i].DstY + float32(y)) * scale
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}

	op := &ebiten.DrawTrianglesOptions{}
	if !line {
		op.FillRule = ebiten.EvenOdd
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

func drawArc(screen *ebiten.Image, count int, scale float32, line bool) {
	var path vector.Path

	path.MoveTo(350, 100)
	const cx, cy, r = 450, 100, 70
	theta1 := math.Pi * float64(count) / 180
	x := cx + r*math.Cos(theta1)
	y := cy + r*math.Sin(theta1)
	path.ArcTo(450, 100, float32(x), float32(y), 30)

	theta2 := math.Pi * float64(count) / 180 / 3
	path.MoveTo(550, 100)
	path.Arc(550, 100, 50, float32(theta1), float32(theta2), vector.Clockwise)

	var vs []ebiten.Vertex
	var is []uint16
	if line {
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	} else {
		vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	}

	for i := range vs {
		vs[i].DstX *= scale
		vs[i].DstY *= scale
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0x33 / float32(0xff)
		vs[i].ColorG = 0xcc / float32(0xff)
		vs[i].ColorB = 0x66 / float32(0xff)
	}

	op := &ebiten.DrawTrianglesOptions{}
	if !line {
		op.FillRule = ebiten.EvenOdd
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

func maxCounter(index int) int {
	return 128 + (17*index+32)%64
}

func drawWave(screen *ebiten.Image, counter int, scale float32, line bool) {
	var path vector.Path

	const npoints = 8
	indexToPoint := func(i int, counter int) (float32, float32) {
		x, y := float32(i*screenWidth/(npoints-1)), float32(screenHeight/2)
		y += float32(30 * math.Sin(float64(counter)*2*math.Pi/float64(maxCounter(i))))
		return x, y
	}

	for i := 0; i <= npoints; i++ {
		if i == 0 {
			path.MoveTo(indexToPoint(i, counter))
			continue
		}
		cpx0, cpy0 := indexToPoint(i-1, counter)
		x, y := indexToPoint(i, counter)
		cpx1, cpy1 := x, y
		cpx0 += 30
		cpx1 -= 30
		path.CubicTo(cpx0, cpy0, cpx1, cpy1, x, y)
	}
	path.LineTo(screenWidth, screenHeight)
	path.LineTo(0, screenHeight)

	var vs []ebiten.Vertex
	var is []uint16
	if line {
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	} else {
		vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	}

	for i := range vs {
		vs[i].DstX *= scale
		vs[i].DstY *= scale
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0x33 / float32(0xff)
		vs[i].ColorG = 0x66 / float32(0xff)
		vs[i].ColorB = 0xff / float32(0xff)
	}

	op := &ebiten.DrawTrianglesOptions{}
	if !line {
		op.FillRule = ebiten.EvenOdd
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

type Game struct {
	counter int

	aa        bool
	line      bool
	offscreen *ebiten.Image
}

func (g *Game) Update() error {
	g.counter++

	// Switch anti-alias.
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.aa = !g.aa
	}

	// Switch lines.
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		g.line = !g.line
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.offscreen != nil {
		w, h := screen.Size()
		if ow, oh := g.offscreen.Size(); ow != w*2 || oh != h*2 {
			g.offscreen.Dispose()
			g.offscreen = nil
		}
	}
	if g.aa && g.offscreen == nil {
		w, h := screen.Size()
		g.offscreen = ebiten.NewImage(w*2, h*2)
	}

	scale := float32(1)
	dst := screen
	if g.aa {
		scale = 2
		dst = g.offscreen
	}

	dst.Fill(color.RGBA{0xe0, 0xe0, 0xe0, 0xff})
	drawEbitenText(dst, 0, 50, scale, g.line)
	drawEbitenLogo(dst, 20, 150, scale, g.line)
	drawArc(dst, g.counter, scale, g.line)
	drawWave(dst, g.counter, scale, g.line)

	if g.aa {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.5, 0.5)
		op.Filter = ebiten.FilterLinear
		screen.DrawImage(g.offscreen, op)
	}

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
	msg += "\nPress A to switch anti-alias."
	msg += "\nPress L to switch the fill mode and the line mode."
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{counter: 0}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Vector (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
