package main

import (
	"bookservice/server"
	"context"
	"flag"
	f "fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

func main() {

	// получаем аргументы , их должно быть не от 2 до 3
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enought arguments")
	} else if flag.NArg() < 3 {
		log.Fatal("too many arguments")
	}

	//Запись переменных из аргументов
	fnc := flag.Arg(0)
	like, err := strconv.ParseBool(flag.Arg(2))
	name := flag.Arg(1)
	if err != nil {
		log.Fatal(err)
	}

	// подключение к серверу
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := server.NewBaseClient(conn)

	res, err := Answer(c, like, fnc, name)
	if err != nil {
		log.Fatal(err)
	}
	println(res)

}

// Функция запроса серверу
func Answer(c server.BaseClient, like bool, fnc, name string) (string, error) {
	if strings.ToUpper(fnc) == "SEARCHBYAUTHOR" {
		if flag.NArg() == 2 {
			resp, err := c.SearchByAuthor(context.Background(), &server.RequestByAuthor{Like: false, Name: name})
			if err != nil {
				return f.Sprintf("error : %s", err), err
			}
			return f.Sprintf("Книги автора '%s' : %s", name, resp.GetAnswer()), nil
		} else {
			resp, err := c.SearchByAuthor(context.Background(), &server.RequestByAuthor{Like: like, Name: name})
			if err != nil {
				return f.Sprintf("error : %s", err), err
			}
			return f.Sprintf("Книги автора '%s' : %s", name, resp.GetAnswer()), nil
		}
	} else if strings.ToUpper(fnc) == "SEARCHBYBOOK" {
		resp, err := c.SearchByBook(context.Background(), &server.RequestByBook{Like: false, Title: name})
		if err != nil {
			return f.Sprintf("error : %s", err), err
		}
		return f.Sprintf("Автор книги '%s' : %s", name, resp.GetAnswer()), nil
	} else {
		return f.Sprintf("error : %s", "wrong method"), f.Errorf("error : %s", "wrong method")
	}
}
