package main
import ( 
	"net/http"
	"github.com/gin-gonic/gin"
)


type metalBar struct {
	ID	  string `json: "id"`
	Title string `json: "title"`
	Price float64 `json: "price"`
	Size float64 `json: "size"`
}
var metalBars = []metalBar{
	{ID: "1", Title: "steel", Price: 80.00, Size: 1},
	{ID: "2", Title: "aluminium", Price: 25.00, Size: 3},
	{ID: "3", Title: "iron", Price: 50.00, Size: 2},
}

func getMetalBars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, metalBars)
}
func getMetalBarsByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range metalBars{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "metalBar not found"})
}
func postMetalBars(c *gin.Context) {
	var newMetalBar metalBar
	if err := c.BindJSON(&newMetalBar); err != nil {
		return
	}
	metalBars = append(metalBars, newMetalBar)
	c.IndentedJSON(http.StatusCreated, newMetalBar)
}
// func deleteMetalBars(c *gin.Context){
// 	id := c.Param("id")
// 	var newMetalBar metalBar
// 	for _, a := range metalBars{
// 		if a.ID == id{
// 			if err := c.BindJSON(&newMetalBar); err != nil{
// 				return
// 			}
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "metalBar not found"})
// }
func main() {
	router := gin.Default()
	router.GET("/metalBars",getMetalBars)
	router.GET("/metalBars/:id",getMetalBarsByID)
	router.POST("/metalBars", postMetalBars)
	//router.DELETE("/metalBars/:id", deleteMetalBars)

	router.Run("localhost:8080")
}


