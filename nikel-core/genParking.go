package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
)

func getParking(c *gin.Context) {
	offset := parseInt(c.Query("offset"), 0, math.MaxInt64, 0)
	limit := parseInt(c.Query("limit"), 0, TOPLIMIT, DEFAULTLIMIT)

	data := autoQuery(
		database.ParkingData,
		c.Request.URL.Query(),
		limit,
		offset,
	)

	res, _ := json.Marshal(data)
	var parking []Parking
	json.Unmarshal(res, &parking)
	if len(parking) == 0 {
		sendEmptySuccess(c)
	} else {
		sendSuccess(c, parking)
	}
}
