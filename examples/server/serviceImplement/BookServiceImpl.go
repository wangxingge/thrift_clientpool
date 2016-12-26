package serviceImplement

import (
	"errors"
	"fmt"
	"github.com/wangxingge/thrift_clientpool/examples/entity"
	"log"
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
	for _, book := range bookStore {
		t := book
		r = append(r, &t)
	}
	return
}

func (srv *BookServiceImpl) AddBook(bookInfo *entity.Book) (r bool, err error) {

	if _, ok := bookStore[bookInfo.BookId]; !ok {
		bookStore[bookInfo.BookId] = *bookInfo
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("Dulplicate book %v", bookInfo.BookId))
}

func (srv *BookServiceImpl) RemoveBook(bookId string) (r bool, err error) {

	if _, ok := bookStore[bookId]; ok {
		delete(bookStore, bookId)
	}

	return true, nil
}

func (srv *BookServiceImpl) DefaultKeepAlive(clientId string) (r bool, err error) {

	log.Printf("Client: %v keep alived.", clientId)
	r = true
	return
}
