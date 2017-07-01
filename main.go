package main
import (
        "runtime"

        "github.com/jinzhu/gorm"
        "gopkg.in/gin-gonic/gin.v1"
        "net/http"
        "./db"
)

type Dinar struct {
    gorm.Model
    Amount float64 `json:"amount"`
    Buy float64 `json:"buy"`
    Sell float64 `json:"sell"`
    Currency string `json:"currency"`
}

type Dirham struct {
    gorm.Model
    Amount float64 `json:"amount"`
    Buy float64 `json:"buy"`
    Sell float64 `json:"sell"`
    Currency string `json:"currency"`
}

func main() {
	router := gin.Default()

        // Do not close DB connection
        defer db.DBCon.Close()

        db.DBCon.AutoMigrate(&Dinar{}, &Dirham{})

        // Let's turn up the cores, baby!
        runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU Cores

        router.GET("/api/v1/dinars", DinarData)
        router.GET("/api/v1/dirhams", DirhamData)

	router.Run(":8080")
        // router.RunUnix("/tmp/gin.idinar.sock")
}

func DinarData(c *gin.Context) {
        dinars := []Dinar{}
        // db.DBCon.Find(&dinars)
        db.DBCon.Where("currency IN (?)", []string{"MYR", "USD", "EUR"}).Order("id desc").Limit(3).Find(&dinars)

        c.JSON(http.StatusOK, dinars)
}

func DirhamData(c *gin.Context) {
       dirhams := []Dirham{}
       // db.DBCon.Find(&dirhams)
       db.DBCon.Where("currency IN  (?)", []string{"MYR", "USD", "EUR"}).Order("id desc").Limit(3).Find(&dirhams)

       c.JSON(http.StatusOK, dirhams)
}
