package main
import (
        "runtime"

        "gopkg.in/gin-gonic/gin.v1"
        "net/http"
        "./db"
)

func main() {
	router := gin.Default()

        // Do not close DB connection
        defer db.DBCon.Close()

        // Let's turn up the cores, baby!
        runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU Cores

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
