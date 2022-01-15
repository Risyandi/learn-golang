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

import "net/http"

// album represent data about a record album
type album struct {
	ID 		string `json:"id"`
	Title 	string `json:"title"`
	Artist 	string `json:"artist"`
	Price 	float64 `json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "john coltrane", Price : 56.99},
	{ID: "2", Title: "jeru", Artist: "gerry mulligan", Price : 17.99},
	{ID: "3", Title: "Sarah Vaughan and clifford brown", Artist: "sarah vaughan", Price : 39.99},
}

// getAlbums responds with the list of all album
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
```