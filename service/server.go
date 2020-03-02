package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

//NewServer Configures and returns a Server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n

}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/products/{productId}",
		getFullfillmentStatusHandler(formatter)).Methods("GET")
}
