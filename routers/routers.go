package routers

import (
    "goblin"
    "sumi/controllers"
)

func init() {
    // session
    goblin.Get("/signup",                    &controllers.SessionsController{}, "Signup")
    goblin.Post("/signup",                   &controllers.SessionsController{}, "Signup")
    goblin.Get("/login",                     &controllers.SessionsController{}, "Login")
    goblin.Post("/login",                    &controllers.SessionsController{}, "Login")
    goblin.Get("/welcome",                   &controllers.SessionsController{}, "Welcome")

    // book
    goblin.Get("/books",                     &controllers.BooksController{}, "Index")
    goblin.Get("/users/:user_id/books/:id",  &controllers.BooksController{}, "Index")
    goblin.Get("/books/:id",                 &controllers.BooksController{}, "Show")
    goblin.Post("/books",                    &controllers.BooksController{}, "Create")
}
