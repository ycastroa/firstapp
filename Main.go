package main

import (
	"fmt"
	"strconv"
)

const (
	_  = iota // i dont care what value is discard
	KB = 1 << (10 * iota)
	MG
	GB
	TB
	PB
	EB
	ZB
	YT
	y       = iota // counter 0
	x int16 = 27   //counter 1
	z       = iota //counter 2
)

func main() {
	var (
		actorName string = "one artist"
		companion string = "second actor"
	)
	fmt.Println(actorName, companion)
	var i int = 42
	fmt.Printf("%v, %T,\n", i, i)

	var o string
	o = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", o, o)

	// var n bool =true
	// fmt.Printf("%v,%T\n",n,n)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v,%T\n", n, n)
	fmt.Printf("%v,%T\n", m, m)

	var c uint16 = 42
	fmt.Printf("%v,%T\n", c, c)

	a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  //1001  if they have one bit set
	fmt.Println(a &^ b) //0100 if they have

	d := 8
	fmt.Println(d << 3) // 2^4 * 2^3 = 2^6
	fmt.Println(d >> 3) //2^3 / 2^3 = 2^0

	var e float64 = 3.14
	e = 13.7e72
	fmt.Println(e)

	var f complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", real(f), real(f))
	fmt.Printf("%v, %T\n", imag(f), imag(f))

	var g complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", real(g), real(g))
	fmt.Printf("%v, %T\n", imag(g), imag(g))

	h := "This is a string"
	//h[2] ="u" you cant
	fmt.Printf("%v,%T\n", h[2], h[2])

	j := "this is a string"
	j2 := "this is also a string"
	fmt.Printf("%v,%T\n", j+j2, j+j2)

	k := "this is string"
	l := []byte(k)
	fmt.Printf("%v,%T\n", l, l)
	fmt.Println(len(k), " arr: ", len(l))

	r := 'a'
	fmt.Printf("%v,%T\n", r, r)
	var r1 rune = 'a'
	fmt.Printf("%v,%T\n", r1, r1)

	const myConst int = 42
	fmt.Printf("%v,%T\n", myConst, myConst)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n,%.1f\n", fileSize/GB, 2.)

	a1 := [] int{}
	a1 = append(a1, 1)
	fmt.Printf("Length: %v\n", len(a1))

	number := 50
	guess := 50
	if guess < number { fmt.Println("too low")}
	if guess == number {fmt.Println("you got it!")}
	if guess > number { fmt.Println("too high")}
	returnTrue()
	i2 :=10
	switch {
	case i2 <=10:
		fmt.Println("one")
	case i2 <=20:   // you can also use case 1 , 3 , 4 but they all have to be integers
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
	var i3 interface{} =1
	switch i3.(type) {
	case int:
		fmt.Println("i is an int")
		break
		//fmt.Println("this will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("is is string")
	default:
		fmt.Println("default is another type")
	}

}
func returnTrue() bool {
	fmt.Println("returning true")
	return true;
}