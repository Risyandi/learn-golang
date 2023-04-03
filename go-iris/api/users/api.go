package users

import (
	"go-iris/user"

	iris "github.com/kataras/iris/v12"
)

type API struct {
	// exported field so api/router.go#api.RegisterDependency can bind it.
	Users user.Repository
}

func (api *API) Configure(r iris.Party) {
	// Add middlewares such as user verification by bearer token here.
	// Authenticated routes...
	r.Post("/signup", api.signUp)
	r.Post("/signin", api.signIn)
	r.Get("/", api.getInfo)
}

func (api *API) getInfo(ctx iris.Context) {
	ctx.WriteString("getInfoFunction")
}

func (api *API) signUp(ctx iris.Context) {
	ctx.WriteString("signupFunction")
}

func (api *API) signIn(ctx iris.Context) {
	ctx.WriteString("signinFunction")
}
