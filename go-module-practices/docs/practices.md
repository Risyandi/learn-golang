## Membuat Module Dan Mengakses Module Lain
1. Pertama buat terlebih dahulu folder untuk module greetings dengan nama "`greetings`"
2. Buka terminal di folder greetings
3. Ketik perintah `$ go mod init example.com/greetings`, untuk menginisiasi dependency pada module greetings.
4. Kemudian buat file dengan nama `greetings.go` dan ketikan kode berikut : 
    ```go
    package greetings

    import "fmt"

    // Hello returns a greeting for the named person.
    func Hello(name string) string {
        // Return a greeting that embeds the name in a message.
        message := fmt.Sprintf("Hi, %v. Welcome!", name)
        return message
    }
    ```
5. setelah membuat module `greetings` selanjutnya membuat folder baru untuk module `hello`.
6. Buka terminal di folder `hello`
7. Ketik perintah `go mod init example.com/hello`, untuk menginisiasi dependency pada module hello.
8. Kemudian buat file dengan nama `hello.go` dan ketikan kode berikut :
    ```go
    package main

    import (
        "fmt"

        "example.com/greetings"
    )

    func main() {
        // Get a greeting message and print it.
        message := greetings.Hello("Gladys")
        fmt.Println(message)
    }
    ```
9. setelah membuat 2 module, "`greetings`" dan "`hello`". selanjutnya kita edit module hello agar bisa mengakses atau mengimport module greetings.  
caranya adalah mengetik perintah  
`$ go mod edit -replace example.com/greetings=../greetings`  
dan jika sukses akan muncul pada `go.mod` hello seperti berikut ini :  
    ```go
        module example.com/hello

        go 1.17

        // akan muncul ini
        replace example.com/greetings => ../greetings
    ```
10. lalu selanjutnya tahap akhir mengsinkronkan dependency yang sebelumnya diubah selanjutnya ketikan perintah berikut  
`$ go mod tidy`
11. untuk menjalankan nya, buka terminal di folder `hello` ketikan perintah  
`$ go run .`