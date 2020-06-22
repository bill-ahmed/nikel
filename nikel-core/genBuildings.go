// DO NOT TOUCH. This file is automatically generated.
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/common"
)

// getBuildings queries via the buildings endpoint
func getBuildings(c *gin.Context) {
	// params, err := parsers.Courses.Parse([]byte(`{"filter": {"code": "ACMB01H3"}}`))

	// if err != nil {
	// 	panic(err)
	// }

	var res common.Course

	db.Last(&res)

	bytes, _ := json.Marshal(res)

	json.Unmarshal(bytes, &res)

	sendSuccess(c, res)
}
