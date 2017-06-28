package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/xuzixx/md_cs/models"
)

// BookController ...
type BookController struct {
	baseController
}

// URLMapping 性能提升
func (c *BookController) URLMapping() {
	c.Mapping("/v1/books", c.All)
}

// All ...
// @router /v1/books [get]
func (c *BookController) All() {
	c.Ctx.WriteString("all")
}

// Create ...
// @router /v1/books [post]
func (c *BookController) Create() {
	b := &models.Book{}

	log.Printf("%s", c.Ctx.Input.RequestBody)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, b)
	if err != nil {
		c.Data["json"] = fmt.Sprintf(`{"code":1,"msg":%s}`, err.Error())
		c.ServeJSON()
		return
	}

	id, err := models.AddOne(b)
	if err != nil {
		c.Data["json"] = fmt.Sprintf(`{"code":1,"msg":%s}`, err.Error())
	} else {
		c.Data["json"] = fmt.Sprintf(`{"code":0,"data":{"id":%d}}`, id)
	}
	c.ServeJSON()
}

// Get ...
// @router /v1/books/:id [get]
func (c *BookController) Get() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.ErrorJSON(err)
		return
	}

	book := &models.Book{ID: id}
	err = models.GetOne(book)
	if err != nil {
		c.ErrorJSON(err)
		return
	}

	c.SuccessJSON(book)
}

// TODO Update
//c.Ctx.WriteString(id)
