// Проблема исходного кода была в том, что, взяв слайс от
// большой строки, переменной justString присваивался слайс меньшей длины,
// но массив, который лежит в основе этого слайса оставался таким же большим,
// каким и был изначально. Это приводит к утечке памяти.
package main

var justString string

func getSubstring(s string, start uint, end uint) string {
	s_slice := []rune(s)
	substring := make([]rune, 0)

	for i := start; i < end; i++ {
		substring = append(substring, s_slice[i])
	}

	return string(substring)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = getSubstring(v, 0, 100)
}

func createHugeString(n uint64) string {
	return string(make([]rune, n))
}

func main() {
	someFunc()
}
