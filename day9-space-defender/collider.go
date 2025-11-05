package main

type Collider interface {
	GetX() float32
	GetY() float32
	GetWidth() float32
	GetHeight() float32
}
