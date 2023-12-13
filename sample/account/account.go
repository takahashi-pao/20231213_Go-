package account

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type UserName interface {
	Get() string
}

type userNameImpl struct {
	UserName
	username string
}

func (i *userNameImpl) Get() string {
	return i.username
}

func validateUsername(s string) error {
	s = strings.Trim(s, " ")
	c := utf8.RuneCountInString(s)
	switch {
	case c == 0:
		return fmt.Errorf("名前を入力してください")
	case c > 10:
		return fmt.Errorf("名前は10文字以内で入力してください")
	}
	return nil
}

func NewUserName(s string) (UserName, error) {
	err := validateUsername(s)
	if err != nil {
		return nil, err
	}
	return &userNameImpl{username: s}, nil
}
