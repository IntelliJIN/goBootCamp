package main

import (
	"testing"
)

var data = Json{
	Tasks: []Task{
		{Title: "first", Important: true},
		{Title: "first number two", Important: true},
		{Title: "second"},
		{Title: "third"},
	},
	Auth: struct {
		SessionID string `json:"session_id"`
		CsrfToken string `json:"csrf_token"`
	}{
		SessionID: "423j345k34h5lh425h34k54343454353rfcwedf",
		CsrfToken: "jflkwnlhtl24hrt3kl4t3k4bt34",
	},
	User: User{
		Name: "Vasya",
		Age:  10303,
		Children: []User{
			{Name: "Petya", Age: 1, Children: []User{
				{Name: "Ivan", Age: 7},
				{Name: "Valeriy", Age: 89, Children: []User{
					{Name: "Ivan", Age: 7},
					{Name: "Valeriy", Age: 89, Children: []User{
						{Name: "Petya", Age: 1, Children: []User{
							{Name: "Ivan", Age: 7},
							{Name: "Valeriy", Age: 89, Children: []User{
								{Name: "Ivan", Age: 7},
								{Name: "Valeriy", Age: 89},
							}},
						}},
						{Name: "Ivan", Age: 7},
						{Name: "Valeriy", Age: 89, Children: []User{
							{Name: "Petya", Age: 1, Children: []User{
								{Name: "Ivan", Age: 7},
								{Name: "Valeriy", Age: 89, Children: []User{
									{Name: "Ivan", Age: 7},
									{Name: "Valeriy", Age: 89},
								}},
							}},
							{Name: "Ivan", Age: 7},
							{Name: "Valeriy", Age: 89},
						}},
					}},
				}},
			}},
			{Name: "Ivan", Age: 7},
			{Name: "Valeriy", Age: 89},
		},
	},
}

//func BenchmarkGenerated(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			Generated(data)
//		}
//	})
//}

func BenchmarkReflect(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Reflect(data)
		}
	})
}
