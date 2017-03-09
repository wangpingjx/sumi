package routers

import (
    "goblin"
    "sumi/controllers"
)

func init() {
    goblin.Get("/books",                     &controllers.BooksController{}, "Index")
    goblin.Get("/users/:user_id/books/:id",  &controllers.BooksController{}, "Index")
    goblin.Get("/books/:id",                 &controllers.BooksController{}, "Show")
    goblin.Post("/books",                    &controllers.BooksController{}, "Create")
}
