package polygon

import "testing"

func TestContainsPoint(t *testing.T) {
	var pg Polygon
	// 凸多边形
	pg.Append(Point{-8234.09, 3247.55})
	pg.Append(Point{-9207.86, -9216.69})
	pg.Append(Point{179.271, -14085.5})
	pg.Append(Point{10657, -9995.71})
	pg.Append(Point{11903.5, 2234.83})
	pg.Append(Point{2360.52, 7376.33})
	pg.Append(Point{-8234.09, 3247.55})
	p1In := Point{-7212, -7941}
	p1Out := Point{-7455, -12956}
	if pg.ContainsPoint(p1In, OddEvenFill) != true {
		t.Errorf("p1In OddEvenFill error")
	}
	if pg.ContainsPoint(p1Out, OddEvenFill) != false {
		t.Errorf("p1Out OddEvenFill error")
	}
	if pg.ContainsPoint(p1In, WindingFill) != true {
		t.Errorf("p1In WindingFill error")
	}
	if pg.ContainsPoint(p1Out, WindingFill) != false {
		t.Errorf("p1Out WindingFill error")
	}

	// 凹多边形
	pg.Clear()
	pg.Append(Point{-10137.8, 3846.43})
	pg.Append(Point{-10683.1, -11422.3})
	pg.Append(Point{11986.2, -10175.8})
	pg.Append(Point{11674.6, 4976})
	pg.Append(Point{223.118, -3710.02})
	pg.Append(Point{-10137.8, 3846.43})
	p2In := Point{-7212, -7941}
	p2Out := Point{-7455, -12956}
	if pg.ContainsPoint(p2In, OddEvenFill) != true {
		t.Errorf("p2In OddEvenFill error")
	}
	if pg.ContainsPoint(p2Out, OddEvenFill) != false {
		t.Errorf("p2Out OddEvenFill error")
	}
	if pg.ContainsPoint(p2In, WindingFill) != true {
		t.Errorf("p2In WindingFill error")
	}
	if pg.ContainsPoint(p2Out, WindingFill) != false {
		t.Errorf("p2Out WindingFill error")
	}
}
