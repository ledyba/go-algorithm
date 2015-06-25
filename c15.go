package algo

import (
	"log"
	"math"
)

type Point struct {
	X int
	Y int
}

func distance(p1, p2 Point) float32 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func reverse(z []int) []int {
	s := make([]int, len(z))
	copy(s, z)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func BitonicTour(pts []Point) ([]int, float32) {
	var bestR [][][]int
	var dist [][]float32
	for i := 0; i < len(pts); i++ {
		dist = append(dist, make([]float32, i))
		bestR = append(bestR, make([][]int, i))
	}
	dist[1][0] = distance(pts[0], pts[1])
	bestR[1][0] = []int{0, 1}
	for j := 1; j < len(pts); j++ {
		log.Printf("Distance: [%v]", j)
		for i := 0; i < j; i++ {
			if i < j-1 {
				d := dist[j-1][i] + distance(pts[j-1], pts[j])
				dist[j][i] = d
				bestR[j][i] = append(bestR[j-1][i], j)
				log.Printf("Distance: [%v,%v]%v", j, i, dist[j][i])
			} else if j > 1 {
				dist[j][i] = 1000
				log.Printf("%v", dist)
				for k := 0; k < j-1; k++ {
					d := dist[i][k] + distance(pts[k], pts[j])
					if d < dist[j][i] {
						dist[j][i] = d
						bestR[j][i] = append(reverse(bestR[i][k]), j)
					}
					log.Printf("Distance: [%v,%v]%v", j, i, dist[j][i])
				}
			}
		}
	}
	mJ, mI := len(pts)-1, len(pts)-2
	minDist := dist[mJ][mI] + distance(pts[mI], pts[mJ])

	return bestR[len(pts)-1][len(pts)-2], minDist
}
