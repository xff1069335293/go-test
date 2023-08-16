package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"httproxy/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetInfo() {

	req := models.Student{}

	body := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Printf("error: %+v\n", err)
	}

	models.Test()
	resp := "111"
	c.Ctx.WriteString(resp)
	return
}
