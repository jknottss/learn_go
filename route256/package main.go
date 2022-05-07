package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nbr int
	for scanner.Scan() {
		nbr, _ = strconv.Atoi(scanner.Text())
	}
	maskPrice, packPrice, boxPrice := 0.55, 11, 160
	pack := 24
	box := pack * 16
	boxes := nbr / box
	nbr -= box * boxes
	packs := nbr / pack
	nbr -= packs * pack
	if float64(nbr)*maskPrice > float64(packPrice) {
		packs += 1
		nbr = 0
	}
	if packs*packPrice > boxPrice {
		boxes += 1
		packs = 0
	}
	fmt.Println(nbr, packs, boxes)

}
