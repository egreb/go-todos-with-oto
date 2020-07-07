package definitions

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type TodosService interface {
	Create(CreateRequest) CreateResponse
	All(AllRequest) AllResponse
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
