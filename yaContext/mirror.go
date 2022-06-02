package main 

import (
        "fmt"
        "strconv"
        "strings"
)

func main() {
var h, w, tmp, res string
fmt.Scan(&h, &w)
heigth, _ := strconv.Atoi(h)
//width, _ := strconv.Atoi(w)
for i:=0; i < heigth; i++ {
        fmt.Scan(&tmp)
        res += fmt.Sprintf("%s\n", tmp)
    }
    res = VertMirror(res)
    res = HorMirror(res)
    fmt.Println(res)
}


func myRev(s string) (res string) {
  for _, ch := range s {
    res = string(ch) + res
  }
  return
}

func VertMirror(s string) (res string) {
  str := strings.Split(s, "\n")
  for i, ch := range str {
    str[i] = myRev(ch)
    res = res + "\n" + str[i]
  }
  res = res [1:]
  return
}
func HorMirror(s string) (res string) {
  str := strings.Split(s, "\n")
  for _, ch := range str {
    res = "\n" + string(ch) + res
  }
  res = res[1:]
  return
}
