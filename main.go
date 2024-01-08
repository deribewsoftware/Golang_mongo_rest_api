package main

import (
	"net/http"

	"github.com/deribewsoftware/mongo_golang_rest/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)
func main() {
	r:=httprouter.New()
	uc :=controllers.NewUserController(getsession())
	r.GET("",)
	r.POST("/",)
	r.DELETE("/")

}

func getSession()*mgo.Session{

}