package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var app *firebase.App
var validate *validator.Validate

func init() {
	opt := option.WithCredentialsFile("secrets/albums-api-golang-firebase-adminsdk-9eoxs-fca0ed4986.json")
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	validate = validator.New()
}

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PATCH("/albums/:id", patchAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	client, err := app.Firestore(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	iter := client.Collection("albums").Documents(context.Background())
	var albums []album
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var a album
		doc.DataTo(&a)
		albums = append(albums, a)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	client, err := app.Firestore(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	doc, err := client.Collection("albums").Doc(id).Get(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	var a album
	doc.DataTo(&a)
	c.IndentedJSON(http.StatusOK, a)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	_, err = client.Collection("albums").Doc(newAlbum.ID).Set(context.Background(), newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	_, err = client.Collection("albums").Doc(id).Set(context.Background(), newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newAlbum)
}
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	client, err := app.Firestore(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	_, err = client.Collection("albums").Doc(id).Delete(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
