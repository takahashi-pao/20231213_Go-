package main

import (
	"fmt"
	"sample/account"
)

func main() {
	// u := account.UserName{} // シンタックスエラー
	// u := account.userNameImpl{username: "hoge"} // シンタックスエラー
	u1, err := account.NewUserName("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u1.Get())
	}
	u2, err := account.NewUserName("01234567890")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u2.Get())
	}
	u3, err := account.NewUserName("hoge")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u3.Get())
	}
}
