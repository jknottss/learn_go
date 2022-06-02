package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func flagParse() map[string]*bool {
	flags := make(map[string]*bool, 4)
	flags["f"] = flag.Bool("f", false, "find files")
	flags["sl"] = flag.Bool("sl", false, "find simlinks")
	flags["d"] = flag.Bool("d", false, "find directory")
	flags["ext"] = flag.Bool("ext", false, "print -f extension")
	flag.Parse()
	return flags
}

func myFind(path, ext string) map[string]string {
	result := make(map[string]string, 5)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			dir := info.IsDir()
			if dir {
				result["d"] += path + "\n"
			} else {
				link, err := os.Readlink(path)
				if err == nil {
					result["sl"] += path + " -> " + link + "\n"
					return nil
				} else {
					result["f"] += path + "\n"
					if strings.HasSuffix(path, ext) {
						result["ext"] += path + "\n"
					}
					return nil
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return result
}

func printResult(flags map[string]*bool, results map[string]string) {
	if flags == nil {
		delete(results, "ext")
		for _, val := range results {
			fmt.Print(val)
		}
		return
	}
	if *flags["d"] {
		fmt.Println(results["d"])
	}
	if *flags["sl"] {
		fmt.Println(results["sl"])
	}
	if *flags["ext"] && *flags["f"] {
		fmt.Println(results["ext"])
		return
	} else if *flags["f"] {
		fmt.Println(results["f"])
	}

}
func main() {
	var path, ext string
	flags := flagParse()
	if *flags["ext"] {
		ext = flag.Arg(0)
		path = flag.Arg(1)
	} else if len(os.Args) == 2 {
		path = os.Args[1]
		flags = nil
	} else {
		path = flag.Arg(0)
	}
	result := myFind(path, ext)
	printResult(flags, result)
}
