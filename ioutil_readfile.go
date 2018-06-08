package main
import (
  "fmt"
  "io/ioutil"
)

 
// 函数原型 func ReadFile(filename string) ([]byte, error)

func main() {

    data, err := ioutil.ReadFile("./a.txt")

    if err != nil {
        fmt.Println("read dir error")
        return
    }

    fmt.Println(string(data))
}

