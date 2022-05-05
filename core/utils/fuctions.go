package utils

import (
	"github.com/aymerick/raymond"
	"github.com/nade-harlow/E-commerce/core/models"
	"io/ioutil"
)

func ParseTemplate(Data models.RecoverPassword) string {
	template, _ := ioutil.ReadFile("core/template/reset_password.html")
	body := raymond.MustRender(string(template), Data)
	return body
}
