package main

import "fmt"

// 構造体定義
type Player struct {
	Name  string
	Score int
}

// 値を受け取る関数（コピーが渡される）
func addScoreByValue(p Player) {
	p.Score += 10
	fmt.Println("関数内（値渡し）:", p.Score)
}

// ポインタを受け取る関数（参照が渡される）
func addScoreByPointer(p *Player) {
	p.Score += 10
	fmt.Println("関数内（ポインタ渡し）:", p.Score)
}

func main() {
	playerA := Player{Name: "Hassy", Score: 50}
	playerB := Player{Name: "Hassy", Score: 50}

	fmt.Println("=== 値渡し ===")
	addScoreByValue(playerA)
	fmt.Println("関数呼び出し後:", playerA.Score) // ← 元の値は変わらない

	fmt.Println("\n=== ポインタ渡し ===")
	addScoreByPointer(&playerB)
	fmt.Println("関数呼び出し後:", playerB.Score) // ← 元の値が変わる
}
