package main

import (
	"github.com/gin-gonic/gin"
	"go-tp-docker/db"
	"go-tp-docker/entity"
	"log"
	"net/http"
	"strconv"
)

func handleGetPeople(c *gin.Context) {
	people, err := entity.Get()

	if nil != err {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"people": people})
}

func handleCreatePerson(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := entity.Create(&person)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleGetVisit(c *gin.Context) {
	rdb, ctx := db.GetRedisConnection()
	defer rdb.Close()

	strCount, err := rdb.Get(ctx, "tp-docker:view-count").Result()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	count, _ := strconv.Atoi(strCount)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func handleNewVisit(c *gin.Context) {
	rdb, ctx := db.GetRedisConnection()
	defer rdb.Close()

	strCount, err := rdb.Get(ctx, "tp-docker:view-count").Result()
	if err != nil {
		strCount = "0"
	}

	count, _ := strconv.Atoi(strCount)
	count++

	err = rdb.Set(ctx, "tp-docker:view-count", count, 0).Err()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
