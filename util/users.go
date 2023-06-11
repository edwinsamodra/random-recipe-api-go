package util

import "webapp1/model"

var Users = map[string]*model.User{
	"aditira": {
		Password: "password1",
		Role:     "admin",
	},
	"dito": {
		Password: "password2",
		Role:     "student",
	},
}
