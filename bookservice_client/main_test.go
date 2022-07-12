package main

import (
	"bookservice/server"
	"log"
	"testing"

	"google.golang.org/grpc"
)

// Тестирование метода получения ответа от сервера
func TestAnswer(t *testing.T) {

	// тестовые кейсы
	tests := []struct {
		like bool
		fnc  string
		name string
		want string
	}{
		{
			like: false,
			fnc:  "searchBYbook",
			name: "1984",
			want: "Автор книги '1984' : Джордж Оруэлл ,",
		},
		{
			like: false,
			fnc:  "searchBYbook",
			name: "На Западном фронте без перемен",
			want: "Автор книги 'На Западном фронте без перемен' : Эрих Мария Ремарк ,",
		},
		{
			like: true,
			fnc:  "searchbyauthor",
			name: "Оскар",
			want: "Книги автора 'Оскар' : Портрет Дориана Грея ,Сказки Оскара Уайльда ,Рыбак и его душа ,",
		},
		{
			like: false,
			fnc:  "SEARCHBYauthor",
			name: "Фёдор Достоевский",
			want: "Книги автора 'Фёдор Достоевский' : Преступление и наказание ,",
		},
	}

	for _, tt := range tests {
		// подключение к серверу
		conn, err := grpc.Dial(":8085", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		// создание клиента
		c := server.NewBaseClient(conn)
		resp, err := Answer(c, tt.like, tt.fnc, tt.name)
		if err != nil {
			t.Errorf("answer(*) got unexpected error : %v", err)
		}
		if resp != tt.want {
			t.Errorf("answer(%v,%v,%v)=%v, wanted %v", tt.fnc, tt.name, tt.like, resp, tt.want)
		}
	}

}
