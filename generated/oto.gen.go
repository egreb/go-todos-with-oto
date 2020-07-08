// Code generated by oto; DO NOT EDIT.

package generated

import (
	"context"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
)

type TodosService interface {
	All(context.Context, AllRequest) (*AllResponse, error)
	Create(context.Context, CreateRequest) (*CreateResponse, error)
	Delete(context.Context, DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, UpdateRequest) (*UpdateResponse, error)
}

type todosServiceServer struct {
	server       *otohttp.Server
	todosService TodosService
}

func RegisterTodosService(server *otohttp.Server, todosService TodosService) {
	handler := &todosServiceServer{
		server:       server,
		todosService: todosService,
	}
	server.Register("TodosService", "All", handler.handleAll)
	server.Register("TodosService", "Create", handler.handleCreate)
	server.Register("TodosService", "Delete", handler.handleDelete)
	server.Register("TodosService", "Update", handler.handleUpdate)
}

func (s *todosServiceServer) handleAll(w http.ResponseWriter, r *http.Request) {
	var request AllRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.todosService.All(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *todosServiceServer) handleCreate(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.todosService.Create(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *todosServiceServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	var request DeleteRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.todosService.Delete(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *todosServiceServer) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var request UpdateRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.todosService.Update(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type AllRequest struct {
}

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type AllResponse struct {
	Todos []Todo `json:"todos"`
	Error string `json:"error,omitempty"`
}

type CreateRequest struct {
	Todo Todo `json:"todo"`
}

type CreateResponse struct {
	Todo  Todo   `json:"todo"`
	Error string `json:"error,omitempty"`
}

type DeleteRequest struct {
	Todo Todo `json:"todo"`
}

type DeleteResponse struct {
	OK    bool   `json:"oK"`
	Error string `json:"error,omitempty"`
}

type UpdateRequest struct {
	Todo Todo `json:"todo"`
}

type UpdateResponse struct {
	OK    bool   `json:"oK"`
	Todo  Todo   `json:"todo"`
	Error string `json:"error,omitempty"`
}
