package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/egreb/go-oto-todos/generated"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	sqliteDatabase, _ := sql.Open("sqlite3", "./todos.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                           // Defer Closing the database
	createTable(sqliteDatabase)

	todoService := TodoService{sqliteDatabase}
	server := otohttp.NewServer()
	generated.RegisterTodosService(server, todoService)

	http.Handle("/oto/", server)
	http.Handle("/", http.FileServer(http.Dir("frontend/public")))
	log.Println("Listening at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create the TodoService interfaces defined in the definitions package.
type TodoService struct {
	*sql.DB
}

func (ts TodoService) All(ctx context.Context, _ generated.AllRequest) (*generated.AllResponse, error) {
	todos, err := getTodos(ts.DB)
	if err != nil {
		return &generated.AllResponse{Error: err.Error()}, nil
	}

	return &generated.AllResponse{Todos: todos}, nil
}

func (ts TodoService) Create(ctx context.Context, req generated.CreateRequest) (*generated.CreateResponse, error) {
	if req.Todo.Task == "" {
		return &generated.CreateResponse{Error: "No task provided"}, nil
	}
	if err := insertTodo(ts.DB, req.Todo); err != nil {
		return &generated.CreateResponse{Error: err.Error()}, nil
	}
	return &generated.CreateResponse{Todo: req.Todo}, nil
}

func insertTodo(db *sql.DB, todo generated.Todo) error {
	log.Println("Inserting todo record ...")
	query := `INSERT INTO todos(task, completed) VALUES (?, ?)`
	statement, err := db.Prepare(query) // Prepare statement.
	if err != nil {
		return err
	}

	_, err = statement.Exec(todo.Task, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func getTodos(db *sql.DB) ([]generated.Todo, error) {
	row, err := db.Query("SELECT * FROM todos ORDER BY completed")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	todos := []generated.Todo{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var task string
		var completed bool

		row.Scan(&id, &task, &completed)
		t := generated.Todo{id, task, completed}
		todos = append(todos, t)
	}
	return todos, nil
}

// create the todo table
func createTable(db *sql.DB) {
	createTodosTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT,
		"completed" BOOL
	  );` // SQL Statement for Create Table

	log.Println("Create todo table...")
	statement, err := db.Prepare(createTodosTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("todo table created")
}
