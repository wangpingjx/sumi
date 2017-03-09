package controllers

import (
    "goblin"
    "log"
    "sumi/models"
 )

 type BooksController struct {
     goblin.Controller
 }

 func (this *BooksController) Index() {
    params := this.Ctx.Params   // 参数
    book   := new(models.Book)  // 实例化
    result := book.List(params["perpage"], params["page"])

    this.RenderJSON(result)
 }

 func (this *BooksController) Show() {
     params := this.Ctx.Params
     book   := new(models.Book)  // 实例化
     result := book.Show(params["id"])

     this.RenderJSON(result)
 }

 func (this *BooksController) Create() {
     log.Println("=> create new book")
     log.Printf("=> Params: %v", this.Ctx.Params)

     this.RenderText("success")
 }
