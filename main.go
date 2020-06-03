package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jim.zhu/ginessential/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
	fmt.Println("test")

}
