## Membuat Web Services Go dan Gin
1. Buat folder baru dengan nama "web-service-gin"  
2. Buka terminal didalam folder yang sebelumnya dibuat.  
 Lalu ketikan perintah pada terminal untuk membuat modul yang akan me manage dependency. Caranya sebagai berikut:   
	- `$ go mod init \`   
Lalu enter dan masukanpath modul kamu   
	- `$ example.com/web-service-gin`  
Lalu enter dan file go berekstensi ***.mod*** akan dibuat, dan tampilan hasil nya akan seperti ini  

![gambar: hasil pembuatan module](https://link)

3. buat file baru dengan nama main.go
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represent data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "john coltrane", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "gerry mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and clifford brown", Artist: "sarah vaughan", Price: 39.99},
}

// getAlbums responds with the list of all album
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
        var newAlbum album

        // Call BindJSON to bind the received JSON to
        // newAlbum.
        if err := c.BindJSON(&newAlbum); err != nil {
                return
        }

        // Add the new album to the slice.
        albums = append(albums, newAlbum)
        c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Loop over the list of albums, looking for
        // an album whose ID value matches the parameter.
        for _, a := range albums {
                if a.ID == id {
                        c.IndentedJSON(http.StatusOK, a)
                        return
                }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

```
4. ketikan perintah `$ go get .` untuk menginstall modul gin yang diperlukan. setelah proses download dan penambahan modul selesai, maka akan muncul tampilan seperti ini:

![gambar: sukses install module](https://link)

5. lalu jalankan perintah berikut ini untuk menjalankan server.  
    - `$ go run .` atau `$ go run main.go`

6. tes server dengan mengetik perintah berikut:  
    - `$ curl -i http://localhost:8080/albums`
    - `$curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: \
    application/json" --request \
    "POST" --data '{"id": \
    "4","title": "The Modern Sound \
    of Betty Carter","artist": \
    "Betty Carter","price": 49.99}'`