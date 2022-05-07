package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("testdata")
	defer file.Close()
    rd := bufio.NewReader(file)
    pos := 0
    for {
        str, _ := rd.ReadString(';')
        pos++;
        if str == "0;" {
            fmt.Println(pos)
            break
        }
    }

}

