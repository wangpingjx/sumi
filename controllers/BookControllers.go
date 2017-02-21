package controllers

import (
    "goblin"
    "log"
 )

 type BooksController struct {
     goblin.Controller
 }

 func (this *BooksController) Index() {
     log.Println("=> in book list")
     log.Printf("=> Params: %v", this.Ctx.Params)

     this.Data["id"]   = 1
     this.Data["name"] = "Anne"
 }

 func (this *BooksController) Show() {
     log.Println("=> in book show")
     log.Printf("=> Params: %v", this.Ctx.Params)
 }

 func (this *BooksController) Create() {
     log.Println("=> create new book")
     log.Printf("=> Params: %v", this.Ctx.Params)
 }
