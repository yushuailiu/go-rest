package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "models"
	"log"
	"fmt"
	"strconv"
	"database/sql"
)

func IndexHandler(context *gin.Context) {
	context.String(http.StatusOK, "首页")
}

func GetPersonHandler(context *gin.Context) {
	cid := context.Param("id")
	id, err := strconv.Atoi(cid)
	if err == strconv.ErrSyntax {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "id must be int",
		})
		return
	}
	person := Person{Id: id}
	err = person.Get()
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    "not found",
			})
			return
		}
		log.Println(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status": 0,
		"person": person,
	})

}

func AddPersonHandler(context *gin.Context) {
	lastName := context.Param("last_name")
	firstName := context.Param("first_name")

	p := Person{LastName: lastName, FirstName: firstName}
	id, err := p.Add()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "add person fail",
		})
		return
	}
	msg := fmt.Sprintf("insert success %d", id)
	context.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    msg,
	})
}

func DeletePersonHandler(context *gin.Context) {
	cid := context.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "id must be int",
		})
		return
	}
	person := Person{Id: id}
	ra, err := person.Delete()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":     "delete fail",
		})
		return
	}
	if ra == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":     "person not exists",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status": 0,
			"id":     id,
		})
	}

}

func UpdatePersonHandler(context *gin.Context) {
	cid := context.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"id":     "id must be int",
		})
		return
	}
	p := Person{Id: id}
	context.Bind(&p)
	ra, err := p.Update()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":     "update fail",
		})
		return
	}
	if ra == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "not found",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status": 0,
			"id":     id,
		})
	}

}

func ListPersonHandler(context *gin.Context) {
	cpage := context.Param("page")
	cnumber := context.Param("number")
	page, err1 := strconv.ParseInt(cpage, 10, 64)
	number, err2 := strconv.ParseInt(cnumber, 10, 64)
	if err1 != nil || err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg": "page and number must be int",
		})
		return
	}
	var p Person
	rs, err := p.List(page, number)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":     "update fail",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  0,
		"persons": rs,
	})
}
