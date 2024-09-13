package main

import (
	"encoding/json" // used to encode/decode data between JSON structure and Go structs
	"log"           // used to write output like errors and debug messages to the console
	"net/http"      // provides HTTP client and server functionality
	"os"            // allows interaction with the operating system, such as reading/writing from files

	"github.com/gin-gonic/gin" //  the web framework we’ll be using to build our API
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums []album

func main() {
	// load albums from JSON
	err := loadAlbumsFromJSON("data.json")
	if err != nil {
		log.Fatalf("Could not load albums: %v", err)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)         // /albums 		- lists all albums
	router.GET("/albums/:id", getAlbumsByID) // /albums/{id} - lists album(s) with a certain ID
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func loadAlbumsFromJSON(filename string) error {
	file, err := os.ReadFile(filename) // read file into memory
	if err != nil {
		return err // if file cannot be read, return error
	}

	err = json.Unmarshal(file, &albums) // unmarshal data from file into albums slice
	if err != nil {
		return err
	}
	return nil // no errors indicates success
}

func saveAlbumsToFile(filename string) error {
	file, err := json.MarshalIndent(albums, "", "  ") // marshal into pretty, readable JSON format
	if err != nil {
		return err // if slice cannot be marshalled, return error
	}

	err = os.WriteFile(filename, file, 0644) // write to JSON file with r/w permissions
	if err != nil {
		return err // if file cannot be written, return error
	}
	return nil // save operation successful
}

// getAlbumns - responds with a JSON list of all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // serializes the album struct into JSON
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through albums, looking for one that matches ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind received JSON to newAlbum
	err := c.BindJSON(&newAlbum)
	if err != nil {
		// if error, return a bad request response with an error message
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	albums = append(albums, newAlbum)

	// save to JSON
	err = saveAlbumsToFile("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
