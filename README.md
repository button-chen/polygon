# polygon
判断点是否在一个多边形区域内， 支持凸多边形与凹多边形(算法源于QT的QPolygonF)

## Example
server:

```go
package main

import "polygon"
import "log"

func main() {
	var pg polygon.Polygon
	pg.Append(polygon.Point{-8234.09, 3247.55})
	pg.Append(polygon.Point{-9207.86, -9216.69})
	pg.Append(polygon.Point{179.271, -14085.5})
	pg.Append(polygon.Point{10657, -9995.71})
	pg.Append(polygon.Point{11903.5, 2234.83})
	pg.Append(polygon.Point{2360.52, 7376.33})
	pg.Append(polygon.Point{-8234.09, 3247.55})
	// 此点在多边形内
	p1In := polygon.Point{-7212, -7941}
	// 此点不在多边形内
	p1Out := polygon.Point{-7455, -12956}
	if pg.ContainsPoint(p1In, polygon.OddEvenFill) {
		log.Println("点p1In在多边形区域内")
	} else {
		log.Println("点p1In不在多边形区域内")
	}
	if pg.ContainsPoint(p1Out, polygon.OddEvenFill) {
		log.Println("点p1Out在多边形区域内")
	} else {
		log.Println("点p1Out不在多边形区域内")
	}
}
```
