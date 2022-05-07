package main

import (
        "os"
        "io/ioutil"
        "encoding/json"
)

func main() {

    type myStruct struct {
        Students []struct {
            Rating []float64 `json:"Rating"`
        }   `json:"Students"`
    }

    type avg struct {
        Average float64
    }
    var out avg
    var unmsh myStruct
    file, _ := os.Open("test.json")
    data, _ := ioutil.ReadAll(file)
    _ = json.Unmarshal(data, &unmsh)
    students := len(unmsh.Students)
    res := 0.0
    for i := 0; i < students; i++ {
        j := len(unmsh.Students[i].Rating)
        res = res + float64(j)
    }
    out.Avarage = float64(res/float64(students))
    fin, _ := json.MarshalIndent(out, "", "    ")
    _, _ = os.Stdout.Write(fin)
}
