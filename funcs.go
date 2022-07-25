package main

func badDivision(a int) (x, y int) {
	if a%2 == 0 {
		x = a / 2
		y = x
	} else {
		x = a / 2
		y = a - y
	}
	return
}
