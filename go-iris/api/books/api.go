package books

import (
	"go-iris/user"

	iris "github.com/kataras/iris/v12"
)

type API struct {
	// exported field so api/router.go#api.RegisterDependency can bind it.
	Users user.Repository
}

type Book struct {
	Title string `json:"title"`
}

func (api *API) Configure(r iris.Party) {
	r.Post("/create", api.create)
	r.Get("/", api.list)
}

func (api *API) list(ctx iris.Context) {
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

func (api *API) create(ctx iris.Context) {
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
	ctx.JSON(b.Title)
	ctx.StatusCode(iris.StatusCreated)
}
