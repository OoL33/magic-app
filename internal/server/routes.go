package server

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"time"

	"magic-app/cmd/web"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.IndexPageHandler)

	r.HandleFunc("/health", s.healthHandler)

	r.HandleFunc("/websocket", s.websocketHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	r.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(web.Index()).ServeHTTP(w, r)
	})

	r.HandleFunc("/index", web.IndexWebHandler)

	r.HandleFunc("/login", s.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", s.LoginPostHandler).Methods("POST")

	r.HandleFunc("/signup", s.SignupPageHandler).Methods("GET")
	r.HandleFunc("/signup", s.SignupPostHandler).Methods("POST")

	return r
}

func (s *Server) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	err := web.Index().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	err := web.LoginForm().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	err := web.LoginPost(username).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	err := web.SignupForm().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	err := web.SignupPost(username).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
