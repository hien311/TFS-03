package fi

func Fi(n int, c chan int64) {

	x := int64(1)
	y := int64(1)
	for i := 1; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
}
