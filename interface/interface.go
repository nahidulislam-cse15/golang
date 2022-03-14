package main
import (
	"fmt"
)
//keyword-reserved word (type) identifier->variable... name(person) type-(struct)
type person struct {
	name string
	age int

}
type secretAgent struct{
	person
	id int
}
// func speak(p person){
// 	fmt.Println("i am ",p)

// }
func (s secretAgent) speak() {
	fmt.Println("i am ",s.person.name," the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("i am ",p.name," the person speak")
}
type human interface{
	speak()
}
func test(h human){
//swith on type
	switch h.(type){
	case person:
		fmt.Println("i am from",h.(person).name)
	case secretAgent:
		fmt.Println("i am from",h.(secretAgent).name)
	}

	//fmt.Println("test",h)
}
func main() {
	p1:=person{"nahid",24}
	fmt.Println(p1)
	s1:=secretAgent{
		person: person{
			name: "Jakib",
			age: 24,
		},
		id:01}
	fmt.Println(s1)
//	speak(p1)
	//speak(s1) will get error
//	speak(s1)
	//speak() get parametre type person now we want to pass secretAgent type also
	//so we need interface ->alowed polymorphism,to specify defined behavior of object 
		s1.speak()
		p1.speak()
		fmt.Printf("%T\n",p1)
		fmt.Printf("%T\n",s1)
	//	fmt.Printf("%T",p1.speak())
	test(p1)
	test(s1)


}