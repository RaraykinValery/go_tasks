package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ans_map := make(map[int][]float64)

	init_arr := []float64{-25.4, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sort.Float64s(init_arr)

	cur_ten := 0

	for _, v := range init_arr {
		v_ten := int(v - math.Mod(v, 10.0))

		if cur_ten != v_ten {
			cur_ten = v_ten
		}

		if _, ok := ans_map[cur_ten]; !ok {
			ans_map[cur_ten] = make([]float64, 0)
		}

		ans_map[cur_ten] = append(ans_map[cur_ten], v)
	}

	fmt.Println(ans_map)
}
