package main

import "fmt"

func main() {
	var nilaiAkhir = 100
	var nilaiAbsen = 90

	var lulusNilaiAkhir bool = nilaiAkhir < 80
	var lulusNilaiAbsen bool = nilaiAbsen < 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsen
	fmt.Println(lulus)
}
