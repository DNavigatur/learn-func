package main

/*

func main() {
	//this is to run the foo func the "()" would have contained an argument if we had decleared a type while making the func foo. like "func foo(s string){...}"
	foo()

	//1 param, no returns
	bar("Henry")

	// 1 param, 1 return
	s := pepsi("Pepsi Whiskey")
	fmt.Println(s)

	//2 params, 2 return
	n, a := DodYears("Gabriel", 32)
	fmt.Println(n, a)

	x := gift(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)
}

// func (r reciver) identifier(p parameter) (return) {...}
func foo() {
	fmt.Println("You are a bad boy")
}

// 1 param, no returns
func bar(s string) {
	fmt.Println("I am", s)
}

// 1 param, 1 return
func pepsi(s string) string {
	return fmt.Sprint("My best drink is ", s)
}

// 2 params, 2 returns
func DodYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " in dog years, is"), age
}

// variadic parameter
func gift(vari ...int) int {
	fmt.Println(vari)
	fmt.Printf("%T\n", vari)

	// //this is unfurling
	//  xi := []int {1,2,3,4,5,6,7,8,9}
	//  //to unfurl it we use the "..." as a surfix of the slice, when declearing a variable
	//  x := gift(xi...)

	n := 0
	for _, v := range vari {
		n += v
	}
	return n
}

type person struct {
	first string
}

// what makes it a method is tha t we decleared a receiver
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Kenny",
	}

	p1.speak()
	p2.speak()
}

//Hands on exercise 1 ==========================

// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results


// call both funcs
// print out their results
func main() {
	foo(45)
	bar(27, "Travis Scott")

}

//  create a func with the identifier foo that returns an int
func foo(n int) {
	fmt.Printf("foo gives %v", n)
}

// create a func with the identifier bar that returns an int and a string

func bar(m int, s string) {
	fmt.Printf("The int is %v, the string is %v", m, s)
}


//Hands on exercise 2 ==========================

//1  create a func with the identifier foo that
// 	 takes in a variadic parameter of type int
// 	 pass in a value of type []int into your func (unfurl the []int)
//2 	 returns the sum of all values of type int passed in
//  create a func with the identifier bar that
// 	 takes in a parameter of type []int
// 	 returns the sum of all values of type int passed in

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(num...) //unfurling the int slice
	fmt.Println("sum of foo is ", n)

	m := bar([]int{1, 2, 3, 4, 5})
	fmt.Println("sum of bar is ", m)
}

func foo(num ...int) int {

	n := 0
	for _, v := range num {
		n += v
	}
	return n
}

func bar(n []int) int {

	m := 0
	for _, v := range n {
		m += v
	}
	return m
}


//Hands on exercise 3 ==========================
//1 “defer” multiple functions in main
// 	 show that a deferred func runs after the func containing it exits.
// 	 determine the order in which the multiple defer funcs run

func main() {
	defer eze("Henry")         //4th
	defer tagarean("Dynereas") //3rd
	stark("Brandon")           //1st
	defer lanister("Imp")      //2nd
}

func stark(a string) {
	fmt.Printf("First of his name %v\n", a)
}
func lanister(b string) {
	fmt.Printf("Second of his name %v\n", b)
}
func tagarean(c string) {
	fmt.Printf("Third of his name %v\n", c)
}
func eze(d string) {
	fmt.Printf("Fourth of his name %v\n", d)
}

//Hands on exercise 4 ==========================

//1 Create a user defined struct with
// 	 the identifier “person”
// 	 the fields:
// 		 first
// 		 age
//2  attach a method to type person with
// 	 the identifier “speak”
// 	 the method should have the person say their name and age
//3  create a value of type person
//4  call the method from the value of type person

type person struct {
	first string
	age   int
}

func main() {

	a := person{
		first: "Dope-Cesear",
		age:   26,
	}

	a.speak()
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, " and my age is ", p.age)

}

// Hands on exercise 5 ==========================

type square struct {
	length float64
	width  float64
}

func (s square) Area() float64 {
	//L * W
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	//PI*r^2
	return math.Pi * math.Pow(c.radius, 2)

}

type shape interface {
	Area() float64
}

func info(s shape) float64 {
	return s.Area()
}

func main() {
	a := square{
		length: 20.5,
		width:  5.12,
	}
	a.Area()

	b := circle{
		radius: 4,
	}
	b.Area()

	fmt.Println(info(a))
	fmt.Println(info(b))
}

// anonymous func hands  on
func main() {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		}()

		//x is assigned a func
		x := func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
		}

		x()
	}

	//hands on exercise
	func main() {
		//y is asigned a func
		y := outer()
		fmt.Println(y())

	}

	func outer() func() int {
		return func() int {
			return 42
		}
	}

*/

//hands-on callback exercise
