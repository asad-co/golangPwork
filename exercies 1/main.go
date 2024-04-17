package main

import "fmt"

func main() {
	var obj Company = Detail{
		Cname: "Asad-co",
		Tnum:  3,
		Names: []string{"Asad", "Zeeshi", "Wajid"},
	}
	obj.ShowCompanyDetail()
	fmt.Println("Total Employee: ", obj.ShowTotalEmployeeNumber())
	fmt.Println("Employee Names: ", obj.ShowAllEmployeeNames())

	fmt.Printf("%T", obj)
}

type Company interface {
	ShowCompanyDetail()
	ShowTotalEmployeeNumber() int
	ShowAllEmployeeNames() []string
}

type Detail struct {
	Cname string
	Tnum  int
	Names []string
}

func (d Detail) ShowCompanyDetail() {
	fmt.Println("Company Name: ", d.Cname)

}
func (d Detail) ShowTotalEmployeeNumber() int {
	return d.Tnum
}
func (d Detail) ShowAllEmployeeNames() []string {
	return d.Names
}
