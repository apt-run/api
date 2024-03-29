package routes

import (
	"fmt"
	"main/handlers"
	"net/http"

	"github.com/sudo-adduser-jordan/gcolor"
	"github.com/sudo-adduser-jordan/loghttp"
)

func SetupRoutes() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", loghttp.Logger(handlers.TestRoot))
	mux.HandleFunc("/search", loghttp.Logger(handlers.Search))
	mux.HandleFunc("/package", loghttp.Logger(handlers.Package))
	// mux.HandleFunc("/sponsored", loghttp.Logger(handlers.TestRoot))
	// mux.HandleFunc("/popular", loghttp.Logger(handlers.TestRoot))

	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	fmt.Print("\n=> http server started on ")
	fmt.Println(gcolor.GreenText("http://localhost" + server.Addr))
	fmt.Println()

	server.ListenAndServe()
}
