package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	phone string
}

type Student struct {
	Person
	education string
	major_subject string
}

type Teacher struct {
	Person
    qualification string
	department string
}

//A human method to say hi
func (person Person) sayHi() {
	fmt.Printf("Hi I am %s and I am %d years old, you can call me on %s \n", person.name, person.age, person.phone)
}

//A human can sing a song
func (person Person) sing(lyrics string) {
	fmt.Println("La la la la.. \n", lyrics)
}

//Student method overrides Human's one
func (student Student) sayHi() {
	fmt.Printf("Hi I am %s, my education is %s and my major major subject is %s \n", student.name, student.education, student.major_subject)
}

//Teacher method overides Person's one
func (teacher Teacher) sayHi() {
	fmt.Printf("Hi I am %s, my qualification is %s and I am working in  %s department \n", teacher.name, teacher.qualification, teacher.department)
}	

// Interface Men is implemented by Person, Student and Teacher
// because it contains methods implemented by them.
type Men interface {
	sayHi()
	sing(lyrics string)
}

var canStoreAnything interface{}

 func main() {
 	mike := Student{Person{"Mike", 25, "222-222-XXX"}, "B.E", "Computer Science"}
    paul := Student{Person{"Paul", 26, "111-222-XXX"}, "M.Tec", "Information Technology"}
    sam := Teacher{Person{"Sam", 36, "444-222-XXX"}, "PHD", "CSE"}
    Tom := Teacher{Person{"Sam", 36, "444-222-XXX"}, "PHD", "ECE"}

    var i Men

 	i = mike
 	fmt.Println("This is mike, a Student")
 	i.sayHi()
 	i.sing("November rain")

 	i = Tom
 	fmt.Println("This is Tom, I am a teacher")
 	i.sayHi()
 	i.sing("Born to be wild")

 	//a slice of Men
 	fmt.Println("Let's use a slice of Men and see what happens")
 	x := []Men{paul, sam, Tom}

 	for _, value := range x {
 		value.sayHi()
 	}

 	fmt.Println("checking the usage of empty interface")
 	r := 5
 	s := "Hello World!"

 	canStoreAnything = r
 	fmt.Println("canStoreAnything after r assignment ", canStoreAnything)

 	canStoreAnything = s
 	fmt.Println("canStoreAnything after s assignment ", canStoreAnything)
 }