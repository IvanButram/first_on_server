package Http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	Handlers *HTTPHandlers
}

func NewServer(HttpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		Handlers: HttpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.Handlers.CreateHandler)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.Handlers.ReadHandler)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.Handlers.UpdateHandler)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.Handlers.DeleteHandler)

	router.Path("/health").Methods("GET").HandlerFunc(s.Handlers.Health)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	}).Methods("GET")

	fileServer := http.FileServer(http.Dir("./frontend"))
	router.PathPrefix("/").Handler(fileServer)

	return http.ListenAndServe(":9091", router)
}
