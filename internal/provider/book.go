package provider

import (
	"app/pkg/book"

	"github.com/google/wire"
)

type BookSet struct {
	BookRepository *book.BookRepository
}

var ProvideBook = wire.NewSet(
	book.NewBookRepository,
	wire.Struct(new(BookSet), "*"),
)
