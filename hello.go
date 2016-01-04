package main

import "fmt"

//import "net/url"

func test() (string, bool) {
	return "hello", true
}

func main() {
	fmt.Printf("hello, world!\n")
	fmt.Println("how are things there!")

	vcid := "6771549831164675055"
	var mode string
	VoyagerBase := "http://voyager.goibibo.com/api/v1/hotels/get_hotels_data_by_city/"
	s := fmt.Sprintf("{\"city_id\":\"%s\",\"mode\":\"%s\"}", vcid, mode)
	fmt.Println(VoyagerBase + "?params=" + s)

	//Testing if interface{} can hold anything
	var a interface{}
	b := "Hello there!"
	a = b
	fmt.Println("a: ", a)

	var t string
	t = "default"
	var ok bool
	t, ok = test()
	if ok {
		fmt.Println("t: ", t)
	} else {
		fmt.Println("t: ", t)
	}
}
