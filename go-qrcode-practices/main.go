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
