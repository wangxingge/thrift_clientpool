package serviceImplement

import (
	"errors"
	"fmt"
	"github.com/wangxingge/thrift_clientpool/examples/entity"
)

var (
	bookStore = make(map[string]entity.Book)
)

type BookServiceImpl struct {
}

func (srv *BookServiceImpl) GetBookById(bookId string) (r *entity.Book, err error) {

	if book, found := bookStore[bookId]; !found {
		return nil, errors.New(fmt.Sprintf("Book %v not found.", bookId))
	} else {
		return &entity.Book{
			BookId:   book.BookId,
			BookName: book.BookName,
			Author:   book.Author,
			Price:    book.Price,
			Date:     book.Price,
			Cover:    book.Cover}, nil
	}
}

func (srv *BookServiceImpl) GetBookByName(bookName string) (r *entity.Book, err error) {

	for _, book := range bookStore {
		if book.BookName == bookName {
			return &entity.Book{
				BookId:   book.BookId,
				BookName: book.BookName,
				Author:   book.Author,
				Price:    book.Price,
				Date:     book.Price,
				Cover:    book.Cover}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Book %v not found.", bookName))
}

func (srv *BookServiceImpl) GetAllBooks() (r []*entity.Book, err error) {
	return
}

func (srv *BookServiceImpl) AddBook(bookInfo *entity.Book) (r bool, err error) {
	return
}

func (srv *BookServiceImpl) RemoveBook(bookId string) (r bool, err error) {
	return
}
