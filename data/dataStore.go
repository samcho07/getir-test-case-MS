package data

import (
	"github.com/samcho07/getir-test-case-MS/data/store"
)

type My_store struct {
	store.My_store
}

func New(holder store.My_store) *My_store {
	return &My_store{
		holder,
	}
}
