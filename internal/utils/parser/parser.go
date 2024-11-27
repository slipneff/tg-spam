package parser

import (
	"bufio"
	"os"
	"strings"
)

type Proxy struct {
	Address  string `json:"address`
	Login    string `json:"login`
	Password string `json:"password`
}

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Recovery string `json:"recovery"`
	Proxy Proxy
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

func (a *Accounts) Add(account Account) {
	a.Accounts = append(a.Accounts, account)
}

func (a *Accounts) Parse(data string) Account {
	f := strings.Split(data, "|")
	d := strings.Split(f[0], ":")
	p := strings.Split(f[1], "@")

	return Account{Email: d[0], Password: d[1], Recovery: d[2], Proxy: Proxy{Address: p[0], Login: p[1], Password: p[2]}}
}

func ReadAccountsFile(file string) (*Accounts, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	accounts := &Accounts{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		accounts.Add(accounts.Parse(scanner.Text()))
	}
	return accounts, nil
}
