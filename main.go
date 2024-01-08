package main

import (
	"net/http"
	"github.com/deribewsoftware/mongo_golang_rest/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	
)
func main() {
	r:=httprouter.New()
	uc :=controllers.NewUserController(getSession())
	r.GET("user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
	r.DELETE("/user:id",uc.DeleteUser)
	http.ListenAndServe("localhost:9000",r)

}

func getSession() *mgo.Session{
	s,err:=mgo.Dial("mongodb+srv://deribewtech:golang_rest@cluster0.pbdh8ir.mongodb.net")
	if err != nil{
		panic(err)
	}
	return s

}