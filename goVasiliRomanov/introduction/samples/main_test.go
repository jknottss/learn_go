package main 

import (
         "bufio"
         "strings"
         "bytes"
         "testing"
)


//используем "`" чтобы сохранить перевод строки
var testOk =  `1
2
3
4
5`
//в функции мы добавляем перевод строки к резульату, поэтому к проверочным данным тоже нужно
var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
    in := bufio.NewReader(strings.NewReader(testOk))
    out := new(bytes.Buffer)
    err := uniq(in, out)
    if err != nil {
        t.Errorf("test for OK failed")
    }
    result := out.String()
    if result != testOkResult {
        t.Errorf("test for OK failed - res not match\n [%v] - expected\n  %v", testOkResult, result)
    }
}
//тест на ошибку при подаче неотсортированного массива в функцию uniq
var testFail = `1
2
1`

func TestForError(t *testing.T) {
     in := bufio.NewReader(strings.NewReader(testFail))
    out := new(bytes.Buffer)
    err := uniq(in, out)
    if err == nil {
        t.Errorf("test for OK failed - %v\n", err)
    }
}
