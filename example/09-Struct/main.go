package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    type Person struct {
        Name string
        Age int
    }

    Zhangsan := Person{"Zhangsan", 18}
    bytes, err := json.Marshal(Zhangsan)
    if err != nil {
        fmt.Println("转换为 json 时，发生错误", err)
    } else {
        fmt.Println(string(bytes))
    }
}