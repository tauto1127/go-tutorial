package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	// jsonではフィールド名が小文字なので，指定する
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
}

func main() {
	router := gin.Default()
	// ルーティング
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// gin.Contextはリクエストとレスポンスの情報を持つ構造体
func getAlbums(c *gin.Context) {
	// structをjsonに変換 こっちはデバックしやすい
	//c.IndentedJSON(200, albums)
	c.JSON(200, albums) //は通常のjson
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(200, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// jsonをstructに変換
	// ;を使って一行にまとめることができる
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
