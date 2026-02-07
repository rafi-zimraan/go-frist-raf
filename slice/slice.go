package main

import "fmt"

func main() {

	name := [...]string{"Rafi", "Zimraan", "Arjuna", "Wijaya"}

	name1 := name[2:4]
	fmt.Println(name1)

	name2 := name[0:3]
	fmt.Println(name2)

	name3 := name[:]
	fmt.Println(name3)

	var name4 []string = name[2:4] // pembuatan manual
	fmt.Println(name4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[0:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Senin Malam"
	daySlice1[1] = "Selasa Malam"
	fmt.Println(daySlice1)
}
