package structural

import "fmt"

type HomeTheater struct {
	screen    *Screen
	projector *Projector
	lights    *Lights
}

func NewHomeTheater() *HomeTheater {
	return &HomeTheater{
		&Screen{},
		&Projector{},
		&Lights{},
	}
}

func (h *HomeTheater) WatchMovie(name string) {
	h.screen.Down()
	h.projector.insertDisk(name)
	h.lights.Off()
	h.projector.Play()
}

type Screen struct {
}

func (*Screen) Down() {
	fmt.Println("Screen is down")
}

type Projector struct {
	currentDisk string
}

func (p *Projector) insertDisk(name string) {
	p.currentDisk = name
	fmt.Println("Disk has been inserted")
}

func (p *Projector) Play() {
	fmt.Printf("Movie %s is playing\n", p.currentDisk)
}

type Lights struct {
}

func (l *Lights) Off() {
	fmt.Println("Lights are off")
}
