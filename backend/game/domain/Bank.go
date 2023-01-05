package domain

type Bank struct {
	id      string
	account *Account
}

func (b Bank) Account() *Account {
	return b.account
}

func newBank() *Bank {
	return &Bank{id: "Bank", account: createUnlimitedAccount("Bank")}
}
