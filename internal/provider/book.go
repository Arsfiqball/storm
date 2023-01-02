package provider

import (
	"app/pkg/book"

	"github.com/google/wire"
)

type BookSet struct {
	BookRepository *book.BookRepository
	BookService    *book.BookService
	BookHandler    *book.BookHandler
}

var ProvideBook = wire.NewSet(
	book.NewBookRepository,
	book.NewBookService,
	book.NewBookHandler,
	wire.Struct(new(BookSet), "*"),
)
