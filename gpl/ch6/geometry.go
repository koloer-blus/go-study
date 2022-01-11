package main

import "math"

type Point struct{ Y, X float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//这种p.Distance的表达式叫做选择器，因为他会选择合适的对应p这个对象的Distance方法来执行
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// (*Point).ScaleBy
func (p *Point) ScaleBy(factor float64) {
	p.Y *= factor
	p.X *= factor
}
