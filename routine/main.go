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

func talk(word string) {
	fmt.Println(word, "と発言しました。")
	time.Sleep(200 * time.Millisecond)
}

func main() {
    // goroutineを2つ起動
    go say("hello")
	go talk("何て？")
    go say("world")
	go talk("こんにちは")

    // main関数がすぐ終わらないように待つ
    time.Sleep(2 * time.Second)
    fmt.Println("main done")
}
