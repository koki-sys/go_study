package samplelib

import "fmt"

// ✅ 先頭が大文字 → public（他パッケージからアクセス可能）
type Dish struct {
	Cuisine string // 公開フィールド
	food string // ❌ 非公開フィールド
}

// ✅ 公開メソッド
func (d Dish) Decide() {
	fmt.Println("The chef is deciding which cuisine to cook, such as:", d.Cuisine)
	d.cook() // 同じパッケージ内なのでprivate関数を呼べる
}

// ❌ 非公開メソッド
func (d Dish) cook() {
	ds := Dish{food: "xiaolongbao"}
	fmt.Println("Chef cooking:", ds.food)
}
