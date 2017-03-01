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
     log.Println("=> in book list")
     log.Printf("=> Params: %v", this.Ctx.Params)

     // render json
     data := make(map[string]interface{})
     data["id"] = 1
     data["name"] = "Anne"
     this.RenderJSON(data)
 }

 func (this *BooksController) Show() {
     log.Println("=> in book show")
     log.Printf("=> Params: %v", this.Ctx.Params)

     // 调用 model 方法
     book := new(models.Book)  // 实例化
     sql  := book.Sample()

     data := make(map[string]interface{})
     data["params"] = this.Ctx.Params
     data["output"] = sql

     this.RenderJSON(data)
 }

 func (this *BooksController) Create() {
     log.Println("=> create new book")
     log.Printf("=> Params: %v", this.Ctx.Params)

     this.RenderText("success")
 }
