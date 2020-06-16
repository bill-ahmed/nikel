package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getFood(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.FoodData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var food []Food
	json.Unmarshal(res, &food)
	if len(food) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, food)
	}
}
