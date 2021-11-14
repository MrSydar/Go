package shamir

import (
	"errors"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type Part struct {
	val       int
	realIndex int
}

func GetPartsRand(s int, n int, t int, p int) ([]Part, error) {
	if t < 1 {
		return nil, errors.New("t can't be less than 1")
	}

	a := make([]int, t)

	a[0] = s
	for i := 1; i < t; i++ {
		a[i] = rand.Int()
	}

	return GetParts(a, n, p)
}

func GetParts(a []int, n int, p int) ([]Part, error) {
	if n < 3 {
		return nil, errors.New("n can't be less than 3")
	}

	parts := make([]Part, n)

	for x := 1; x <= n; x++ {
		var sum int
		for i, si := range a {
			sum += si * int(math.Pow(float64(x), float64(i)))
		}
		parts[x-1] = Part{sum % p, x}
	}

	return parts, nil
}

func flatten(f [][]float64) (r, c int, d []float64) {
	r = len(f)
	if r == 0 {
		panic("bad test: no row")
	}
	c = len(f[0])
	d = make([]float64, 0, r*c)
	for _, row := range f {
		if len(row) != c {
			panic("bad test: ragged input")
		}
		d = append(d, row...)
	}
	return r, c, d
}

func GetSecret(parts []Part, t int, p int) (int, error) {
	partsNum := len(parts)

	if partsNum < 3 {
		return -1, errors.New("minimum three parts must be selected")
	}

	a := make([][]float64, partsNum)
	b := make([][]float64, partsNum)
	for i, p := range parts {
		a[i] = make([]float64, t)
		for j := 0; j < t; j++ {
			a[i][j] = math.Pow(float64(p.realIndex), float64(j))
		}
		b[i] = make([]float64, 1)
		b[i][0] = float64(p.val)
	}

	ma := mat.NewDense(flatten(a))
	mb := mat.NewDense(flatten(b))
	var x mat.Dense
	x.Solve(ma, mb)

	return int(x.At(0, 0)) % p, nil
}
