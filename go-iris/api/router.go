package api

import (
	"time"

	"go-iris/api/books"
	"go-iris/api/users"
	"go-iris/pkg/database"
	"go-iris/user"

	"github.com/kataras/iris/v12/middleware/modrevision"
)

// buildRouter is the most important part of your server.
// All root endpoints are registered here.
func (srv *Server) buildRouter() {

	// Add a simple health route.
	secondsEastOfUTC := int((7 * time.Hour).Seconds())
	srv.Any("/health", modrevision.New(modrevision.Options{
		ServerName:   srv.config.ServerName,
		Env:          srv.config.Env,
		Developer:    "Risyandi",
		TimeLocation: time.FixedZone("indonesian time", secondsEastOfUTC),
	}))

	// grouping api version
	api := srv.Party("/api/v1")
	api.RegisterDependency(
		database.Open(srv.config.ConnString),
		user.NewRepository,
	)

	// parent endpoint api
	api.PartyConfigure("/user", new(users.API))
	api.PartyConfigure("/books", new(books.API))
}
