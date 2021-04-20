package common

import "github.com/volatiletech/null"

type Product struct {
	Id   int
	Name string
}

type Fat struct {
	Id   int
	Type string
}

type Info struct {
	Id        int
	ProductId int
	FatId     int
	Price     float64

	Product Product `json:"-"`
	Fat     Fat `json:"-"`
}

type Sale struct {
	InfoId    int
	Amount    int
	Sold      null.Time
	Purchased null.Time

	Info Info `json:"-"`
}
