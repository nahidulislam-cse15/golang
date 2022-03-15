package main

import (
	"fmt"
	"sort"
)
type people []string

func (p people) Len() int           { return len(p) }//o(n)
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i]}//o(nlog(n))
func (p people) Less(i, j int) bool { return p[i] <p[j] } //o(nlog(n))
type Planet struct {
	name string
	mass float64
	diastance float64
}
// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Planet) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}
func main() {
	numbers := []int{1, 8, 3, 7,2, 5,4, 6}
//	names := people{"Nahid","jakib","sabuj","shaun","Roman","Mahram"}
	//fmt.Println(numbers)
	//sort.Ints(numbers)
	//sort.Sort(names)
	sort.Sort(people(numbers))
	fmt.Println(numbers)

	//fmt.Println("")
	//fmt.Println(names)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
	var planets = []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}
	// Closures that order the Planet structure.
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return distance(p2, p1)
	}

	// Sort the planets by the various criteria.
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)

}
	

}