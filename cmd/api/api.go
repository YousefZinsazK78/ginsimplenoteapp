package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yousefzinsazk78/simple_note_api/internal/database"
	"github.com/yousefzinsazk78/simple_note_api/internal/handler"
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
		db        = database.NewDatabase(dbConn)
		noteStore = database.NewNoteStore(db)
		helper    = handler.NewHandler(noteStore)
		r         = gin.Default()
		v2        = r.Group("/api/v2")
	)

	///post api v1
	v2.POST("/posts", helper.HandleInsertNote)
	v2.GET("/posts", helper.HandleGetNotes)
	v2.GET("/posts/title/:title", helper.HandleGetNoteByTitle)
	v2.DELETE("/posts/delete/:id", helper.HandleDeleteNote)
	v2.PUT("/posts/update", helper.HandlePutNote)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
