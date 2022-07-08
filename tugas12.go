package main

import (
	"fmt"
	"regexp"
)

func main() {
	var hewan = "kucing kelinci singa"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}
	var tampil = regex.FindAllString(hewan, 3)
	fmt.Println(tampil)

	var tampil2 = regex.FindString(hewan)
	fmt.Println(tampil2)
}
