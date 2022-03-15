package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//float64s
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted ->sorted Nan
	sort.Float64s(s)
	fmt.Println(s)
	//output: [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]

	//float64sAreSorted
	s = []float64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	fmt.Println(sort.Float64sAreSorted(s)) //true

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	fmt.Println(sort.Float64sAreSorted(s)) //false

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	fmt.Println(sort.Float64sAreSorted(s)) //false
	//Ints
	s1 := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s1)
	fmt.Println(s1)

	//IntsAreSorted
	s1 = []int{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sort.IntsAreSorted(s1))

	s1 = []int{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sort.IntsAreSorted(s1))

	s1 = []int{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sort.IntsAreSorted(s1))
	//search Ascending order
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	//search descending order
	a = []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x = 6

	i = sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	//SearchFloat64s
	a1 := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}
	x1 := 2.0
	i1 := sort.SearchFloat64s(a1, x1)
	fmt.Printf("found %g at index %d in %v\n", x1, i1, a1)
	x1 = 0.5
	i1 = sort.SearchFloat64s(a1, x1)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", x1, i1, a1) //SearchInts
	a2 := []int{1, 2, 3, 4, 6, 7, 8}

	x2 := 2
	i2 := sort.SearchInts(a2, x2)
	fmt.Printf("found %d at index %d in %v\n", x2, i2, a2)

	x2 = 5
	i2 = sort.SearchInts(a, x)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x2, i2, a2)
	//slice
	a3 := []int{1, 2, 3, 4, 6, 7, 8}

	x3 := 2
	i3 := sort.SearchInts(a3, x3)
	fmt.Printf("found %d at index %d in %v\n", x3, i3, a3)

	x3 = 5
	i3 = sort.SearchInts(a3, x3)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x3, i3, a3)

}
