package controllers

import (
    "goblin"
    "log"
    "html/template"
    "sumi/models"
    "time"
)

type SessionsController struct {
    goblin.Controller
}

func (this *SessionsController) Signup() {
    log.Printf("=> in SessionsController#Signup")
    if this.Ctx.Request.Method == "GET" {
        log.Printf("=> in GET")
        t, err := template.ParseFiles("/Users/vikki/iwork/Mine/Go/src/sumi/views/signup.gtpl")
        if err != nil {
            log.Printf("error: %v", err)
        } else {
            t.Execute(this.Ctx.ResponseWriter, nil)
        }
    } else if this.Ctx.Request.Method == "POST" {
        log.Printf("=> in POST")
        params := this.Ctx.Params   // 参数
        log.Printf("params is: %v", params)
        user := models.User{ Name: params["username"], Password: params["password"], Created_at: time.Now(), Updated_at: time.Now()}
        models.DB.Create(&user)
        log.Printf("user is: %v", user)
        t, err := template.ParseFiles("/Users/vikki/iwork/Mine/Go/src/sumi/views/login.gtpl")
        if err != nil {
            log.Printf("error: %v", err)
        } else {
            t.Execute(this.Ctx.ResponseWriter, nil)
        }
    }
}

func (this *SessionsController) Login() {
    log.Printf("=> in SessionsController#Login")
    if this.Ctx.Request.Method == "GET" {
        t, err := template.ParseFiles("/Users/vikki/iwork/Mine/Go/src/sumi/views/login.gtpl")
        if err != nil {
            log.Printf("error: %v", err)
        } else {
            t.Execute(this.Ctx.ResponseWriter, nil)
        }
    } else if this.Ctx.Request.Method == "POST" {
        params := this.Ctx.Params   // 参数
        log.Printf("params is: %v", params)
        var user models.User
        models.DB.Where("name = ?", params["username"]).Where("password = ?", params["password"]).Limit(1).Find(&user)
        log.Printf("user is: %v", user)
        if user.Id > 0 {
            t, err := template.ParseFiles("/Users/vikki/iwork/Mine/Go/src/sumi/views/welcome.gtpl")
            if err != nil {
                log.Printf("error: %v", err)
            } else {
                t.Execute(this.Ctx.ResponseWriter, nil)
            }
        }
    }
}

// 登录后才可以访问
func (this *SessionsController) Welcome() {
    log.Printf("=> in SessionsController#Welcome")
    t, err := template.ParseFiles("/Users/vikki/iwork/Mine/Go/src/sumi/views/welcome.gtpl")
    if err != nil {
        log.Printf("error: %v", err)
    } else {
        t.Execute(this.Ctx.ResponseWriter, nil)
    }
}
