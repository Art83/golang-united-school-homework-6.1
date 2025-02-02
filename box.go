package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func NewBox(shapesCapacity int) *box {
	return &box{shapesCapacity: shapesCapacity,
		shapes: make([]Shape, 0)}
}

func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("out of range")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("no such index")
	} else {
		return b.shapes[i], nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("no such Index")
	} else {
		elem := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return elem, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("no such index")
	} else {
		elem := b.shapes[i]
		b.shapes[i] = shape
		return elem, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sm float64
	for _, v := range b.shapes {
		sm += v.CalcPerimeter()
	}
	return sm
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sm float64
	for _, v := range b.shapes {
		sm += v.CalcArea()
	}
	return sm
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	cnt := 0
	cp := b.shapes
	b.shapes = make([]Shape, 0)
	for _, val := range cp {
		_, err := val.(*Circle)
		if err == false {
			b.shapes = append(b.shapes, val)
		} else {
			cnt++
		}
	}
	if cnt == 0 {
		return errors.New("no circles")
	}
	return nil
}
