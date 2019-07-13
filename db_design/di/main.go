package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Khigashiguchi/go-snippets/db_design/di/models"
)

// Env implements handlers.
type Env struct {
	DB *sql.DB
}

func main() {
	db, err := models.NewDB("sqlite3", "./simple.sqlite3")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{DB: db}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	bks, err := models.GetAllBooks(env.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %d円\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
