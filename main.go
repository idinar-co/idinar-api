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

        router.GET("/api/v1/dinars", func(c *gin.Context) {
               data := DinarData()
               c.JSON(http.StatusOK, data)
        })
        router.GET("/api/v1/dirhams", func(c *gin.Context) {
               data := DirhamData()
               c.JSON(http.StatusOK, data)
        })
        router.GET("/api/v1/dinars/:currency", func(c *gin.Context) {
		currency := c.Param("currency")
                data := DinarToFiat(currency)
                c.JSON(http.StatusOK, data)
        })
        router.GET("/api/v1/dirhams/:currency", func(c *gin.Context) {
                currency := c.Param("currency")
                data := DirhamToFiat(currency)
                c.JSON(http.StatusOK, data)
        })

	router.Run(":8080")
        // router.RunUnix("/tmp/idinar.gin.sock")
}

func DinarData() []Dinar {
        dinars := []Dinar{}
        db.DBCon.Select("DISTINCT *").Where("currency IN (?)", []string{"MYR", "USD", "EUR"}).Order("id desc").Limit(3).Find(&dinars)
        return dinars
}

func DirhamData() []Dirham {
       dirhams := []Dirham{}
       db.DBCon.Select("DISTINCT *").Where("currency IN  (?)", []string{"MYR", "USD", "EUR"}).Order("id desc").Limit(3).Find(&dirhams)
       return dirhams
}

func DinarToFiat(currency string) Dinar {
      dinar := Dinar{}
      db.DBCon.Where("currency = ?", currency).Last(&dinar)
      return dinar
} 

func DirhamToFiat(currency string) Dirham {
      dirham := Dirham{}
      db.DBCon.Where("currency = ?", currency).Last(&dirham)
      return dirham
}
