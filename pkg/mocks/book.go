package mocks

import "go-gorm-orm/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
	{
		Id:     2,
		Title:  "CLR via C#",
		Author: "Richter",
		Desc:   "Book for study C#",
	},
	{
		Id:     3,
		Title:  "C# 8.0",
		Author: "Ben Alblahari",
		Desc:   "Bag documentation",
	},
}
