package main

import (
	"math"
	"math/rand"
)

type Vector []float64

// 计算两个向量之间的欧几里得距离
func (b Vector) euclideanDistance(a Vector) float64 {
	if len(a) != len(b) {
		panic("两个向量的长度不一致")
	}

	var sum float64
	for i := 0; i < len(a); i++ {
		sum += (a[i] - b[i]) * (a[i] - b[i])
	}
	return math.Sqrt(sum)
}

type cluster struct {
	Center  Vector
	Vectors []Vector
}

func (c *cluster) updateCenter() {
	center := make([]float64, len(c.Vectors[0]))
	n := float64(len(c.Vectors))

	for _, p := range c.Vectors {
		for i := range center {
			center[i] += p[i]
		}
	}

	for i := range center {
		center[i] /= n
	}

	c.Center = center
}

func KMeansPP(points []Vector, k int, iter int) []cluster {
	// Step 1: Choose the first center uniformly at random
	centers := make([]Vector, k)
	centers[0] = points[rand.Intn(len(points))]
	dists := make([]float64, len(points))

	// Step 2: Choose the remaining centers via weighted sampling
	for i := 1; i < k; i++ {
		sumDist := 0.0
		for j, p := range points {
			minDist := math.MaxFloat64
			for _, c := range centers[:i] {
				dist := p.euclideanDistance(c)
				if dist < minDist {
					minDist = dist
				}
			}
			dists[j] = minDist * minDist
			sumDist += dists[j]
		}
		randVal := rand.Float64() * sumDist
		index := 0
		for ; index < len(points)-1 && randVal > 0; index++ {
			randVal -= dists[index]
		}
		centers[i] = points[index]
	}

	// Step 3: Initialize the clusters
	clusters := make([]cluster, k)
	for i := range clusters {
		clusters[i].Center = centers[i]
	}

	// Step 4: Run the iterative algorithm
	cnt := 0
	for {
		if cnt > iter {
			break
		}
		cnt++

		// Update the cluster assignments
		for i := range clusters {
			clusters[i].Vectors = nil
		}
		for _, p := range points {
			minDist := math.MaxFloat64
			index := 0
			for j, c := range clusters {
				dist := p.euclideanDistance(c.Center)
				if dist < minDist {
					minDist = dist
					index = j
				}
			}
			clusters[index].Vectors = append(clusters[index].Vectors, p)
		}

		// Exit condition: no point changes clusters
		changed := false
		for _, c := range clusters {
			oldCenter := c.Center
			c.updateCenter()
			if c.Center.euclideanDistance(oldCenter) > 1e-6 {
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	return clusters
}
