package main

import (
	"fmt"
	"sort"
)

/*
median of values:
-sort values
-if odd number of values - return middle
-if even number of values - return average of middle
*/

func main() {
	nums := []int{10, 20, 30}
	fmt.Println("len:", len(nums))
	nums[1] = 200
	fmt.Println("1:", nums[1])

	for i := range nums { //by index
		fmt.Println(i)
	}

	for i, v := range nums { //by index and value
		fmt.Println(i, v)
	}

	for _, v := range nums { //by value
		fmt.Println(v)
	}
	nums = append(nums, 40)
	fmt.Println(nums)

	s := nums[1:3] //slicing
	fmt.Println(s)

	s[0] = 999
	fmt.Println(s)
	fmt.Println(nums)

	fmt.Printf("s: len=%d, cap=%d\n", len(s), cap(s))

	values := []float64{2, 1, 3}
	fmt.Println(median(values)) // 2
	values = append(values, 4)
	fmt.Println(median(values)) // 2.5
	fmt.Println(values)
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	vs := make([]float64, len(values))
	copy(vs, values)

	sort.Float64s(vs) //if we sorted values, it would be reflected outside of function, so we make a copy above
	i := len(vs) / 2
	if len(vs)%2 == 1 {
		return vs[i], nil
	}

	v := (vs[i-1] + vs[i]) / 2
	return v, nil
}
