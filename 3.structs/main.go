// [1,3,5,2,6,3,7,8,2,3,41]

package main

import (
	"fmt"
)

type contactInfo struct {
	mobile int
	email  string
}

type student struct {
	rollNo  int
	name    string
	contact contactInfo
}

func (s student) display() {
	fmt.Printf("%+v \n", s)
}

func (s *student) updateName(name string) {
	(*s).name = name
}

func main() {

	s1 := student{
		rollNo: 14,
		name:   "Yeswanth",
		contact: contactInfo{mobile: 908045678,
			email: "yas@gmail.com",
		},
	}
	s1.display()
	s1.updateName("Yas")
	s1.display()

}
