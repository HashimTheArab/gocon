package example

import "github.com/hashimthearab/gocon"

// Token config is used for saving a refresh token on a users device.
type Token struct {
	Name  string
	Token string
}

func (Token) Path() string {
	return "token.json"
}

func SaveToken(name string, token string) {
	_ = gocon.Save(Token{Name: name, Token: token})
}

func LoadToken() Token {
	token, err := gocon.Load[Token]()
	if err != nil {
		// You can return some default values if nothing is available.
		return Token{}
	}
	return token
}
