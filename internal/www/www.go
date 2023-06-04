package www

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/vianamjr/page/internal/www/static"
)

var version = "no_version"

//go:embed templates
var templateFS embed.FS

var home = template.Must(template.ParseFS(templateFS, "templates/*"))

type Config struct {
	Port int
}

func (c Config) verify() error {
	if c.Port <= 0 {
		return errors.New("invalid port number")
	}

	return nil
}

type Service struct {
	config Config
	server *http.Server
}

func New(c Config) (*Service, error) {
	if err := c.verify(); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(static.FS))
	// TODO: add a custom file handler, only accepts *.css || *.js paths.
	// TODO: do not serve the files (folders view) under /static...
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)

	return &Service{
		config: c,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", c.Port),
			Handler: mux,
		},
	}, nil
}

func (s *Service) ListenAndServe() error {
	log.Printf("INFO Starting server port=%d version=%s", s.config.Port, version)
	return s.server.ListenAndServe()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	home.ExecuteTemplate(w, "index.html", nil)
}
