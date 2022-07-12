package grpc_server

import (
	"bookservice/databases"
	srv "bookservice/server"
	"context"
	f "fmt"
)

type GRPCServer struct{}

// метод обработки grpc метода поиска по книге
func (s *GRPCServer) SearchByAuthor(ctx context.Context, rq *srv.RequestByAuthor) (*srv.Response, error) {
	db := databases.NewConfig("mysql", "root", "myPassWord", "127.0.0.1", "3307")
	con, err := databases.Open("books", db)
	defer databases.Close(con)
	databases.Takeerr(err, "search by author")
	sel, err := databases.DBSearchByAuthor(con, rq.Name, rq.Like)
	databases.Takeerr(err, "search by author")
	b := ""
	for _, str := range sel {
		b += f.Sprintf("%s ,", str)
	}
	a := &srv.Response{Answer: b}
	return a, nil
}

// метод обработки grpc метода поиска по книге
func (s *GRPCServer) SearchByBook(ctx context.Context, rq *srv.RequestByBook) (*srv.Response, error) {
	db := databases.NewConfig("mysql", "root", "myPassWord", "127.0.0.1", "3307")
	con, err := databases.Open("books", db)
	defer databases.Close(con)
	databases.Takeerr(err, "search by book")
	sel, err := databases.DBSearchByBook(con, rq.Title, false)
	databases.Takeerr(err, "search by book")
	b := ""
	for _, str := range sel {
		b += f.Sprintf("%s ,", str)
	}
	a := &srv.Response{Answer: b}
	return a, nil
}
