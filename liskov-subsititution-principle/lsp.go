package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func (r *Square) GetWidth() int {
	return r.width
}

func (r *Square) SetWidth(width int) {
	r.width = width
	r.height = width
}

func (r *Square) GetHeight() int {
	return r.height
}

func (r *Square) SetHeight(height int) {
	r.height = height
	r.width = height
}

type Square2 struct {
	size int
}

func (sq *Square2) Rectangle() *Rectangle {
	return &Rectangle{width: sq.size, height: sq.size}
}

func UseIt(size Sized) {
	expectedArea := 10 * size.GetWidth()
	size.SetHeight(10)
	actualArea := size.GetHeight() * size.GetWidth()
	fmt.Printf("Expected areas %d and actual area is %d \n", expectedArea, actualArea)
}

func main() {
	fmt.Println("lsp")
	rect := &Rectangle{width: 20, height: 30}
	UseIt(rect)

	// this void lsp
	square := &Square{}
	square.SetHeight(20)
	square.SetWidth(30)
	UseIt(square)

	// via this can handle the same withou breaking lsp
	square2 := Square2{size: 20}
	UseIt(square2.Rectangle())
}
