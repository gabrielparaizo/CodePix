package model

import (
	uuid "github.com/satori/uuid.go"
	"time"
)

type Bank struct {
	Base
	Code string `json:"code"`
	Name string `json:"name"`
}

func (b *Bank) isValid() error {
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
