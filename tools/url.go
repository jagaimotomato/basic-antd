package tools

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func IdStrToIdsGroup(key string, c *gin.Context) (Ids []int) {
	ids := strings.Split(c.Param(key), ",")
	for i := 0; i < len(ids); i++ {
		Id, _ := StringToInt(ids[i])
		Ids = append(Ids, Id)
	}
	return
}

func StrToStrGroup(key string, c *gin.Context) (str []string) {
	str = strings.Split(c.Request.FormValue(key), ",")
	return
}
