package main

import (
	"context"
	"database/sql"
	"fmt"
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

func (ts TodoService) Delete(ctx context.Context, req generated.DeleteRequest) (*generated.DeleteResponse, error) {
	if err := deleteTodo(ts.DB, req.Todo); err != nil {
		return &generated.DeleteResponse{OK: false}, err
	}

	return &generated.DeleteResponse{OK: true}, nil
}

func (ts TodoService) Update(ctx context.Context, req generated.UpdateRequest) (*generated.UpdateResponse, error) {
	if err := updateTodo(ts.DB, req.Todo); err != nil {
		return &generated.UpdateResponse{OK: false}, err
	}

	return &generated.UpdateResponse{OK: true, Todo: req.Todo}, nil
}

func (ts TodoService) All(ctx context.Context, _ generated.AllRequest) (*generated.AllResponse, error) {
	todos, err := getTodos(ts.DB)
	if err != nil {
		return &generated.AllResponse{Error: err.Error()}, nil
	}

	return &generated.AllResponse{Todos: todos}, nil
}

func (ts TodoService) Create(ctx context.Context, req generated.CreateRequest) (*generated.CreateResponse, error) {
	todo := req.Todo
	if todo.Task == "" {
		return &generated.CreateResponse{Error: "No task provided"}, nil
	}
	todo, err := insertTodo(ts.DB, req.Todo)
	if err != nil {
		return &generated.CreateResponse{Error: err.Error()}, nil
	}
	return &generated.CreateResponse{Todo: todo}, nil
}

// SQL functions
func insertTodo(db *sql.DB, todo generated.Todo) (generated.Todo, error) {
	log.Println("Inserting todo record ...")

	query := `INSERT INTO todos(task, completed) VALUES (?, ?)`
	statement, err := db.Prepare(query) // Prepare statement.
	if err != nil {
		return todo, err
	}
	defer statement.Close()

	res, err := statement.Exec(todo.Task, todo.Completed)
	if err != nil {
		return todo, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return todo, err
	}
	todo.ID = int(id)
	return todo, nil
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

func updateTodo(db *sql.DB, todo generated.Todo) error {
	log.Println("Update todo record ...")

	query := fmt.Sprintf(`UPDATE todos SET task = '%s', completed = %v WHERE id=%d`, todo.Task, todo.Completed, todo.ID)
	statement, err := db.Prepare(query) // Prepare statement.
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func deleteTodo(db *sql.DB, todo generated.Todo) error {
	log.Println("Deleting todo record..")

	query := fmt.Sprintf("DELETE FROM todos WHERE id=%d", todo.ID)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
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
