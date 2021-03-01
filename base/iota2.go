package main

func main() {
	// 每当定义一个常量， iota就会自动累加1
	const (
		a = iota  // 0
		b = iota  // 1
		c = "hhh" // 2
		d         // hhh, iota = 3
		e = 100
		f
		g = iota
	)
	const (
		h = iota
		i
	)
}
