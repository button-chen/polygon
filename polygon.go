package polygon

import "math"

/*
OddEvenFill:
Specifies that the region is filled using the odd even fill rule.
With this rule, we determine whether a point is inside the shape
by using the following method. Draw a horizontal line from the point
to a location outside the shape, and count the number of intersections.
If the number of intersections is an odd number, the point is inside the shape.
This mode is the default.

WindingFill:
Specifies that the region is filled using the non zero winding rule.
With this rule, we determine whether a point is inside the shape by
using the following method. Draw a horizontal line from the point to
a location outside the shape. Determine whether the direction of the
line at each intersection point is up or down. The winding number is
determined by summing the direction of each intersection. If the number
is non zero, the point is inside the shape. This fill mode can also in most
cases be considered as the intersection of closed shapes.
*/

const (
	OddEvenFill = iota
	WindingFill
)

type Point struct {
	X float32
	Y float32
}

type Polygon struct {
	points []Point
}

func (this *Polygon) Append(p Point) {
	this.points = append(this.points, p)
}

func (this *Polygon) Clear() {
	this.points = this.points[0:0]
}

func (this *Polygon) ContainsPoint(pt Point, fillrule int) bool {
	if len(this.points) == 0 {
		return false
	}
	winding_number := 0
	last_pt := this.points[0]
	last_start := this.points[0]
	for _, e := range this.points {
		this.polygonIsectLine(last_pt, e, pt, &winding_number)
		last_pt = e
	}
	if last_pt != last_start {
		this.polygonIsectLine(last_pt, last_start, pt, &winding_number)
	}
	if fillrule == OddEvenFill {
		return (winding_number != 0)
	}
	return (winding_number%2 != 0)
}

func (this *Polygon) fuzzyCompare(p1, p2 float32) bool {
	return (math.Abs(float64(p1-p2)) <= 0.00001*math.Min(math.Abs(float64(p1)), math.Abs(float64(p2))))
}

func (this *Polygon) polygonIsectLine(p1, p2, pos Point, winding *int) {
	x1, y1, x2, y2, y := p1.X, p1.Y, p2.X, p2.Y, pos.Y
	dir := 1
	if this.fuzzyCompare(y1, y2) {
		return
	} else if y2 < y1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	if y >= y1 && y < y2 {
		x := x1 + ((x2-x1)/(y2-y1))*(y-y1)
		if x <= pos.X {
			(*winding) += dir
		}
	}
}
