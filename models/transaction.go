package models

type Transaction struct {
	Data []byte
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{Data: data}
}

func (t *Transaction) toStr() string {
	return "Data: " + string(t.Data)
}
