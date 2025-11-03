package main

import (
	"strc_sample/samplelib" // 自作パッケージをimport
)

func main() {
	// ✅ Publicなフィールド・関数にはアクセス可能
	dish := samplelib.Dish{Cuisine: "Chinese"}
	dish.Decide() // OK: 大文字で始まる → 公開されている

	// ❌ Privateな関数・フィールドにはアクセスできない
	// dish := samplelib.Dish{Cuisine: "Chinese", food: "xiaolongbao"}
	// foodはprivate（小文字）なので、mainからはアクセスできない！
	// fmt.Println(dish.food) // コンパイルエラー:上と同じ理由 
}
