package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"bug-tracker/internal/handler"
	"bug-tracker/internal/repository"
	"bug-tracker/internal/service"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_go")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	issueRepo := repository.NewIssueRepository(db)
	issueService := service.NewIssueService(issueRepo)
	issueHandler := handler.NewIssueHandler(issueService)

	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			issueHandler.CreateIssue(w, r)
		case http.MethodGet:
			issueHandler.GetAllIssues(w, r)
		default:
			http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/issues/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			issueHandler.UpdateIssue(w, r)
		case http.MethodDelete:
			issueHandler.DeleteIssue(w, r)
		case http.MethodGet:
			issueHandler.GetIssueByID(w, r)
		default:
			http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
