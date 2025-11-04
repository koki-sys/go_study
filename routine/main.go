package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        fmt.Println(s, ":", i)
        time.Sleep(300 * time.Millisecond)
    }
}

func main() {
    // goroutineを2つ起動
    go say("hello")
    go say("world")

    // main関数がすぐ終わらないように待つ
    time.Sleep(2 * time.Second)
    fmt.Println("main done")
}
