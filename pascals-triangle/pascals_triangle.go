package pascal

const testVersion = 1

func Triangle(i int) [][]int {
	var xs = [][]int{[]int{1}}
	return triangle(i, xs)
}

func triangle(pos int, last [][]int) [][]int {
	if pos == 1 {
		return last
	}
	return triangle(pos-1, append(last, sig(last[len(last)-1])))
}

func sig(xs []int) []int {
	var ws = make([]int, len(xs)+1)
	for i := 0; i < len(xs); i++ {
		ws[i] += xs[i]
		ws[i+1] += xs[i]
	}
	return ws
}
