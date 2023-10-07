package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load("./internal/config/.env"); err != nil {
		panic(err)
	}

	dbConn, err := sql.Open("postgres", os.Getenv("ConnectionStr"))
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	var (
		// db        = database.NewDatabase(dbConn)
		// noteStore = database.NewNoteStore(db)
		// helper    = handler.NewHandler(noteStore)
		r = gin.Default()
		// v1 = r.Group("/api/v1")
	)

	///note api v1
	// v1.POST("/notes", helper.HandleInsertNote)
	// v1.GET("/notes", helper.HandleGetNotes)
	// v1.GET("/notes/title/:title", helper.HandleGetNoteByTitle)
	// v1.DELETE("/notes/delete/:id", helper.HandleDeleteNote)
	// v1.PUT("/notes/update", helper.HandlePutNote)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
