package polygon

import "testing"

func BenchmarkLoops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pg Polygon
		pg.Append(Point{-8234.09, 3247.55})
		pg.Append(Point{-9207.86, -9216.69})
		pg.Append(Point{179.271, -14085.5})
		pg.Append(Point{10657, -9995.71})
		pg.Append(Point{11903.5, 2234.83})
		pg.Append(Point{2360.52, 7376.33})
		pg.Append(Point{-8234.09, 3247.55})
		p1In := Point{-7212, -7941}
		p1Out := Point{-7455, -12956}

		pg.ContainsPoint(p1In, OddEvenFill)
		pg.ContainsPoint(p1Out, OddEvenFill)
	}
}
