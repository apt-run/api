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
	mux.HandleFunc("/sponsored", loghttp.Logger(handlers.TestRoot))
	mux.HandleFunc("/popular", loghttp.Logger(handlers.TestRoot))

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	fmt.Print("\n=> http server started on ")
	fmt.Println(gcolor.GreenText("localhost" + server.Addr))
	fmt.Println()

	server.ListenAndServe()
}
