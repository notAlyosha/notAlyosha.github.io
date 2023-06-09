package snake

import (
	"image/color"
	"snake/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

type Snake struct {
	el     []types.SnakeElement
	dir    int
	size   int
	length int
}

func (s *Snake) Init(w, h, size int) {
	s.el = make([]types.SnakeElement, 1)
	s.el[0] = types.SnakeElement{X: 3, Y: 2, Color: color.RGBA{R: 40, G: 170, B: 190, A: 255}}
	s.dir = down
	s.size = size
}

func (s *Snake) NewStatement(w, h, size int) {
	s.outOfBounds(w, h, size)
	s.changeDir()
	s.move()
	s.shift()
}
func (s *Snake) Increase() {
	s.el = append(s.el, types.SnakeElement{})
}
func (s *Snake) EatsApple(a types.Position) bool {
	if s.el[0].X == a.X && s.el[0].Y == a.Y {
		return true
	}
	return false
}

func (s *Snake) outOfBounds(w, h, size int) {

	if s.el[0].X > w/size {
		s.el[0].X = 0
		return
	}

	if s.el[0].X < 0 {
		s.el[0].X = w / size
		return
	}

	if s.el[0].Y > h/size {
		s.el[0].Y = 0
		return
	}

	if s.el[0].Y < 0 {
		s.el[0].Y = h / size
		return
	}
}

func (s *Snake) move() {

	if s.dir == up {
		s.el[0].Y -= 1
		return
	}
	if s.dir == down {
		s.el[0].Y += 1
		return
	}
	if s.dir == left {
		s.el[0].X -= 1
		return
	}
	if s.dir == right {
		s.el[0].X += 1
		return
	}
}

func (s *Snake) changeDir() {

	if ebiten.IsKeyPressed(ebiten.KeyW) && s.dir != down {
		s.dir = up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && s.dir != up {
		s.dir = down
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && s.dir != right {
		s.dir = left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && s.dir != left {
		s.dir = right
	}
}

func (s *Snake) EatsItSelf() bool {
	for i := 2; i < len(s.el); i += 1 {
		if s.el[0].X == s.el[i].X && s.el[0].Y == s.el[i].Y {
			return true
		}
	}
	return false
}

func (s *Snake) shift() {
	for i := len(s.el) - 1; i > 0; i -= 1 {
		s.el[i] = s.el[i-1]
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for i := 0; i < len(s.el); i += 1 {
		ebitenutil.DrawRect(screen, float64(s.el[i].X)*float64(s.size), float64(s.el[i].Y)*float64(s.size), float64(s.size), float64(s.size), s.el[i].Color)
	}
}
