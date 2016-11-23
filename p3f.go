package main

import "fmt"
import "time"

func primes(n int) []int {
	p := make([]int, 0, 1000000)
	p = append(p, 2, 3, 5)
	m := n/6 + 1
	for i := 7; i < m; i = i + 2 {
		lenp := len(p)
		for j := 0; j < lenp; j++ {
			if i%p[j] == 0 {
				break
			} else if int64(p[j])*int64(p[j]) > int64(i) {
				p = append(p, i)
				break
			}
		}
	}
	fmt.Println("nprimes = ", len(p))
	return p
}

func prime3factor(num int, p []int) (int64, int64) {
	sum := int64(0)
	counter := sum
	pn := len(p)
	for ai := 0; ai < pn-2; ai++ {
		a := int64(p[ai])
		ax := a
		for { // powers of a
			for bi := ai + 1; bi < pn-1; bi++ {
				b := int64(p[bi])
				axby := ax * b
				if axby >= int64(num) {
					break
				}
				for { // powers of b
					for ci := bi + 1; ci < pn; ci++ {
						c := int64(p[ci])
						nn := axby * c
						if nn >= int64(num) {
							break
						}
						for { // powers of c
							counter++
							sum += nn
							nn = nn * c
							if nn >= int64(num) {
								break
							}
						}
					}
					axby = axby * b
					if axby >= int64(num) || axby*int64(p[bi+1]) >= int64(num) {
						break
					}
				}
			}
			ax = ax * a
			if ax >= int64(num) ||
				ax*int64(p[ai+1]) >= int64(num) ||
				ax*int64(p[ai+1])*int64(p[ai+2]) >= int64(num) {
				break
			}
		}
	}
	return sum, counter
}

func main() {
	n := 1000000000
	st := time.Now()
	p := primes(n)
	pt := time.Now()
	fmt.Println("setup: ", pt.Sub(st))
	s, c := prime3factor(n, p)
	fmt.Println("calc / counter / sum : ", time.Now().Sub(pt), " / ", c, " / ", s)
}
