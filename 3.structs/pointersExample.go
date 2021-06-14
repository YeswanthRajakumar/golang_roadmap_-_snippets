package main

import "fmt"

type person struct {
	name string
}
type personList []person

func (p *personList) updatePerson(name string,nameToUpdate string)  {
	fmt.Println(*p)
	pIndex := -1
	for i,person := range *p {
		if person.name == name{
			pIndex = i
		}
	}
	if pIndex != -1{
		(*p)[pIndex].name = nameToUpdate
	}
}


func main() {

	people := personList{{name: "Ravi"}, {name: "Shyam"}, {name: "Kamal"},{name: "Charlie"}}
	people.updatePerson("Kamal","Kamal Raj")
	fmt.Print(people)
}
