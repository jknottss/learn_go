package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var path string

	fileFlag := flag.Bool("f", false, "find files")
	simFlag := flag.Bool("sl", false, "find simlinks")
	flag.Parse()
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		path2 := flag.Args()
		path = path2[0]
	}
	fmt.Println(path)
	if *fileFlag {
		fmt.Println("file flag - true")
	}
	if *simFlag {
		fmt.Println("simlink flag - true")
	}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			link, err := os.Readlink(path)
			if err != nil {
				fmt.Println(path)
				return nil
			}
			fmt.Println(path, "->", link)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
