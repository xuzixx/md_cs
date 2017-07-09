package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego/logs"

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
// @Title Book all
// @Description get all books
// @Success 200 {object} models.BooksDTO
// @router / [get]
func (c *BookController) All() {
	var books []*models.Book
	// TODO pagination
	count, err := models.GetAll(&models.Book{}, &books, 10, 0)
	if err != nil {
		c.ErrorJSON(err)
		return
	}

	logs.Info("-----------", count)

	c.SuccessJSON(books)
	//c.Ctx.WriteString("all")
}

// Create ...
// @router / [post]
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
// @Title Book one
// @Description get book by id
// @Param id path int true "the bookid you want to get"
// @Success 200 {object} models.BookDTO
// @router /:id [get]
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
