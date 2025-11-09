package main

import "fmt"

// Interface を宣言
type Accessor interface {
	GetName() string
	SetName(string)
}

// Accessor を満たす実装
// 明示的な宣言は必要なく、実装と完全に分離している
type User struct {
	name string
}

func (u *User) GetName() string {
	return u.name
}

// 基本的には、ポインタレシーバで実装する。
// func (u User) GetName() string {
// 	return u.name
// }

func (u *User) SetName(name string) {
	u.name = name
}

func main() {
	// Userのインスタンスを直接変更しても値渡しになってしまうのでポインタを使用
	var user *User = &User{}
	user.SetName("user")
	fmt.Println(user.GetName())
	
	//Accessor Interface を実装しているのでAccessor 型に代入可能
	var acsr Accessor = &User{}
	acsr.SetName("accessor")
	fmt.Println(acsr.GetName())
}

