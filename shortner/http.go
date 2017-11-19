package shortner

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

type HttpServer struct {
	App Container
}

func (s HttpServer) Listen(port string) {
	http.HandleFunc("/", s.FrontController)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Error returned from HTTP server: %s", err.Error())
	}
}

func (s HttpServer) FrontController(w http.ResponseWriter, r *http.Request) {

	log.Printf(fmt.Sprintf("Handling request to: %s", r.URL.String()))

	switch r.URL.Path {
	case "/shorten":
		serveShorten(s.App.repository, r, w)
		break

	default:
		serveRedirect(s.App.repository, r, w)
		break
	}
}

func serveShorten(repo Repository, r *http.Request, w http.ResponseWriter) {
	long := r.URL.Query().Get("long")

	if long == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "You must provide a `long` query string parameter")
		return
	}

	alias := strconv.Itoa(rand.Int())

	repo.addShortUrl(alias, long)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Your short URL is <a href='/?alias=%s'>/?alias=%s</a>", alias, alias)
}

func serveRedirect(repo Repository, r *http.Request, w http.ResponseWriter) {

	alias := r.URL.Query().Get("alias")

	location, err := repo.getLocation(alias)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Redirect alias not found")
		return
	}

	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusTemporaryRedirect)

	fmt.Fprintf(w, "Redirecting you to URL is <a href='%'>%s</a>\n", location, location)
}
