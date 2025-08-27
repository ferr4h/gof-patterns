package structural

import (
	"fmt"
)

type Renderer interface {
	RenderCircle(x, y, radius float64) string
	RenderRectangle(x, y, width, height float64) string
}

type ASCIIRenderer struct {
}

func (r *ASCIIRenderer) RenderCircle(x, y, radius float64) string {
	return fmt.Sprintf("ASCII: circle at (%.2f, %.2f), r=%.2f", x, y, radius)
}

func (r *ASCIIRenderer) RenderRectangle(x, y, width, height float64) string {
	return fmt.Sprintf("ASCII: rect at (%.2f, %.2f), w=%.2f, h=%.2f", x, y, width, height)
}

type DebugRenderer struct {
}

func (r *DebugRenderer) RenderCircle(x, y, radius float64) string {
	return fmt.Sprintf("DEBUG: CIRCLE{x:%.2f,y:%.2f,r:%.2f}", x, y, radius)
}

func (r *DebugRenderer) RenderRectangle(x, y, width, height float64) string {
	return fmt.Sprintf("DEBUG: RECT{x:%.2f,y:%.2f,w:%.2f,h:%.2f}", x, y, width, height)
}

type Shape interface {
	Draw() string
	SetRenderer(Renderer)
}

type Circle struct {
	renderer Renderer
	x, y     float64
	radius   float64
}

func NewCircle(renderer Renderer, x, y, radius float64) *Circle {
	return &Circle{renderer, x, y, radius}
}

func (c *Circle) SetRenderer(renderer Renderer) {
	c.renderer = renderer
}

func (c *Circle) Draw() string {
	if c.renderer == nil {
		return "no renderer"
	}
	return c.renderer.RenderCircle(c.x, c.y, c.radius)
}

type Rectangle struct {
	renderer      Renderer
	x, y          float64
	width, height float64
}

func NewRectangle(renderer Renderer, x, y, width, height float64) *Rectangle {
	return &Rectangle{renderer, x, y, width, height}
}

func (r *Rectangle) SetRenderer(renderer Renderer) {
	r.renderer = renderer
}

func (r *Rectangle) Draw() string {
	if r.renderer == nil {
		return "no renderer"
	}
	return r.renderer.RenderRectangle(r.x, r.y, r.width, r.height)
}
