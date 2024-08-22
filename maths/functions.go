package maths

func Linearregression(xaxis, yaxis []float64) (float64, float64) {
	n := float64(len(xaxis))
	sumx := Sumxax(xaxis)
	sumy := Sumyax(yaxis)
	sumxy := Xymultsum(xaxis, yaxis)
	sumxsq := Sumxsqrd(xaxis)

	b := (n*sumxy - sumx*sumy) / (n*sumxsq - sumx*sumx)
	a := (sumy - b*sumx) / n
	return b, a
}

func Sumyax(yaxis []float64) float64 {
	temp := yaxis[0]

	for i := 1; i < len(yaxis); i++ {
		temp += yaxis[i]
	}
	return temp
}

func Sumxax(xaxis []float64) float64 {
	temp := xaxis[0]

	for i := 1; i < len(xaxis); i++ {
		temp += xaxis[i]
	}
	return temp
}

func Sumxsqrd(xaxis []float64) float64 {
	temp := 0.0

	for i := 0; i < len(xaxis); i++ {
		temp += xaxis[i] * xaxis[i]
	}
	return temp
}

func Xymultsum(xaxis, yaxis []float64) float64 {
	temp := 0.0

	for i := 0; i < len(xaxis); i++ {
		temp += xaxis[i] * yaxis[i]
	}
	return temp
}
