package main

import "fmt"

type (
	Payment struct{}
	repo1   struct{}
	repo2   struct{}

	RepoInterface interface {
		GetPayment(id string) (Payment, error)
		SavePayment(payment Payment) error
	}
)

func (*repo1) GetPayment(id string) (Payment, error) {
	return Payment{}, nil
}

func (*repo1) SavePayment(payment Payment) error {
	return nil
}

func (repo2) SavePayment(payment Payment) error {
	return nil
}

func (repo2) GetPayment(id string) (Payment, error) {
	return Payment{}, nil
}

func New() RepoInterface {
	return &repo1{}
}

func main() {
	var _ RepoInterface = &repo1{}
	var r RepoInterface = repo2{}

	_, _ = New().GetPayment("!23")

	fmt.Println("run")

	switch r.(type) {
	case RepoInterface:
	default:
	}
}
