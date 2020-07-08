package definitions

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type TodosService interface {
	Update(UpdateRequest) UpdateResponse
	Create(CreateRequest) CreateResponse
	All(AllRequest) AllResponse
	Delete(DeleteRequest) DeleteResponse
}

type DeleteRequest struct {
	Todo Todo
}

type DeleteResponse struct {
	OK bool
}

type UpdateRequest struct {
	Todo Todo
}

type UpdateResponse struct {
	OK   bool
	Todo Todo
}

type CreateRequest struct {
	Todo Todo
}

type CreateResponse struct {
	Todo Todo
}

type AllRequest struct{}

type AllResponse struct {
	Todos []Todo
}
