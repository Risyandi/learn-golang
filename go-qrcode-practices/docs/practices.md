# Membuat Simple QR Code Generator

1. Pertama yang harus dipersiapkan adalah membuat file  *go.mod* sebagai file yang memanajemen package yang digunakan. cara nya adalah dengan mengetikan perintah berikut:  
`$ go mod go-qrcode-practices`  
Lalu file akan secara otomatis ter generate dengan sendirinya dengan file baru dengan nama *go.mod* dan *go.sum*  
2. Buat file baru dengan nama `main.go` dan ketikan kode berikut:  

    ```go
        package main

        // import library qrcode
        import (
            "image/png"
            "net/http"
            "text/template"

            "github.com/boombuler/barcode"
            "github.com/boombuler/barcode/qr"
        )

        // initialize struct Page
        type Page struct {
            Title string
        }

        // main function
        // Call function homeHandler and viewCodeHandler and run server on port 8080
        func main() {
            http.HandleFunc("/", homeHandler)
            http.HandleFunc("/generator/", viewCodeHandler)
            http.ListenAndServe(":8080", nil)
        }

        func homeHandler(w http.ResponseWriter, r *http.Request) {
            p := Page{Title: "QR Code Generator"}
            t, _ := template.ParseFiles("generator.html")
            t.Execute(w, p)
        }

        func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
            dataString := r.FormValue("data")
            qrCode, _ := qr.Encode(r.FormValue(dataString), qr.L, qr.Auto)
            qrCode, _ = barcode.Scale(qrCode, 300, 300)
            png.Encode(w, qrCode)
        }
    ```

3. Lalu install package library "github.com/boombuler/barcode" dengan mengetikan perintah berikut:
`$ go get github.com/boombuler/barcode`
4. Setelah itu buat file html dengan memberikana nama `generator.html` lalu ketikan kode berikut:  

    ```html
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>QR Code Generator</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        <div>Please enter the string you want to generate as Qrcode</div>
        <form action="generator/" method="post">
            <input type="text" name="dataString">
            <input type="submit" value="Submit">
        </form>s
    </body>
    </html>
    ```

5. Setelah selesai, aplikasi QR generator ini bisa dijalankan dengan mengetikan perintah berikut:  
`$ go run main.go`
