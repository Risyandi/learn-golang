package old-main
import (
	iris "github.com/kataras/iris/v12"
)
func main() {
	app := iris.New()

	// grouping routes version 1
	v1 := app.Party("/v1")
	{
		v1.Use(iris.Compression)

		// GET: http://localhost:8080/v1/books
		// POST: http://localhost:8080/v1/books

		v1.Get("/books", list)
		v1.Post("/books", create)
	}

	// grouping routes version 2
	v2 := app.Party("/v2")
	{
		v2.Use(iris.Compression)

		// GET: http://localhost:8080/v2/books
		// POST: http://localhost:8080/v2/books

		v2.Get("/books", list)
		v2.Post("/books", create)
	}

	app.Listen(":8080")
}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)

	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
