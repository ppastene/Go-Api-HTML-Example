package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Estadio Nacional", Artist: "Los Prisioneros", Year: 2002},
	{ID: "2", Title: "Californication", Artist: "Red Hot Chili Peppers", Year: 1999},
	{ID: "3", Title: "Moving Pictures", Artist: "Rush", Year: 1981},
	{ID: "4", Title: "Take Off Your Pants and Jacket", Artist: "Blink-182", Year: 2001},
	{ID: "5", Title: "Rush in Rio", Artist: "Rush", Year: 2003},
	{ID: "6", Title: "City of Evil", Artist: "Avenged Sevenfold", Year: 2005},
	{ID: "7", Title: "Ocean Avenue", Artist: "Yellowcard", Year: 1999},
	{ID: "8", Title: "Get Ready", Artist: "Rare Earth", Year: 1969},
	{ID: "9", Title: "Number of The Beast", Artist: "Iron Maiden", Year: 1982},
	{ID: "10", Title: "Outlandos d'Amour", Artist: "The Police", Year: 1978},
	{ID: "11", Title: "Bleed America", Artist: "Jimmy Eat World", Year: 2001},
	{ID: "12", Title: "Audioslave", Artist: "Audioslave", Year: 2002},
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func newAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "My Favourite Albums",
		})
	})
	albumsRoute := router.Group("/albums")
	{
		albumsRoute.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "albums.html", gin.H{
				"title":  "My Favourite Albums",
				"albums": albums,
			})
		})
		albumsRoute.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			for _, album := range albums {
				if album.ID == id {
					fmt.Println(album)
					c.HTML(http.StatusOK, "album.html", gin.H{
						"album": album,
					})
					return
				}
			}
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"title":   "Album not found",
				"message": "Album not found",
			})
		})
	}
	// API
	api := router.Group("/api")
	{
		api.GET("albums/:id", getAlbum)
		api.GET("albums", getAlbums)
		api.POST("albums", newAlbum)
	}

	router.Run(":3000")
	fmt.Println("Server running at port 3000")
}
