package main // определение пакета для текущего файла

import "fmt"

/*
создание функции main
*/
func main() {
	var n, m int
	cnt := 1
	fmt.Scan(&n)
	for fmt.Scan(&m); m != 0; fmt.Scan(&m) {
		if m > n {
			cnt = 1
			n = m
		} else if m == n {
			cnt++
		}
	}
	fmt.Print(cnt)
}
