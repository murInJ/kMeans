package kmeans

func floats2Vector(floats []float64) Vector {
	v := make(Vector, len(floats))
	copy(v, floats)
	return v
}

func vector2Floats(vector Vector) []float64 {
	f := make([]float64, len(vector))
	copy(f, vector)
	return f
}

func floatss2Vector(floats [][]float64) []Vector {
	vectors := make([]Vector, len(floats))
	for i, floats := range floats {
		vectors[i] = floats2Vector(floats)
	}
	return vectors
}

func vectors2floatss(vectors []Vector) [][]float64 {
	floatss := make([][]float64, len(vectors))
	for i, vector := range vectors {
		floatss[i] = vector2Floats(vector)
	}
	return floatss
}
