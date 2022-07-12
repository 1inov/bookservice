package grpc_server

import (
	"bookservice/server"
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestSearchByAuthor(t *testing.T) {
	// запуск сервера
	s := grpc.NewServer()
	srv := &GRPCServer{}
	// регистрация grpc сервера
	server.RegisterBaseServer(s, srv)

	tests := []struct {
		like bool
		name string
		want string
	}{
		{
			like: false,
			name: "Джордж Оруэлл",
			want: "1984 ,Глотнуть воздуха ,Скотный Двор ,",
		},
		{
			like: false,
			name: "Эрих Мария Ремарк",
			want: "Три товарища ,На Западном фронте без перемен ,Возвращение ,",
		},
		{
			like: true,
			name: "Достоевский",
			want: "Преступление и наказание ,",
		},
		{
			like: true,
			name: "Оскар",
			want: "Портрет Дориана Грея ,Сказки Оскара Уайльда ,Рыбак и его душа ,",
		},
	}

	for _, tt := range tests {
		req := &server.RequestByAuthor{Like: tt.like, Name: tt.name}
		resp, err := srv.SearchByAuthor(context.Background(), req)
		if err != nil {
			t.Errorf("SearchByAuthor(%v) got unexpected error", err)
		}
		if resp.Answer != tt.want {
			t.Errorf("SearchByAuthor(%v,%v)=%v, wanted %v", tt.like, tt.name, resp.Answer, tt.want)
		}
	}
}

func TestSearchByBook(t *testing.T) {
	// запуск сервера
	s := grpc.NewServer()
	srv := &GRPCServer{}
	// регистрация grpc сервера
	server.RegisterBaseServer(s, srv)

	tests := []struct {
		like  bool
		title string
		want  string
	}{
		{
			like:  false,
			title: "1984",
			want:  "Джордж Оруэлл ,",
		},
		{
			like:  false,
			title: "На Западном фронте без перемен",
			want:  "Эрих Мария Ремарк ,",
		},
		{
			like:  true,
			title: "Рыбак и его душа",
			want:  "Оскар Уайльд ,",
		},
		{
			like:  true,
			title: "Преступление и наказание",
			want:  "Фёдор Достоевский ,",
		},
	}

	for _, tt := range tests {
		req := &server.RequestByBook{Like: tt.like, Title: tt.title}
		resp, err := srv.SearchByBook(context.Background(), req)
		if err != nil {
			t.Errorf("SearchByAuthor(%v) got unexpected error", err)
		}
		if resp.Answer != tt.want {
			t.Errorf("SearchByAuthor(%v,%v)=%v, wanted %v", tt.like, tt.title, resp.Answer, tt.want)
		}
	}
}
