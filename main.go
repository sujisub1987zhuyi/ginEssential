package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
	fmt.Println("test")

}






