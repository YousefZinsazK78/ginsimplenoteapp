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
		postStore = database.NewPostStore(db)
		roleStore = database.NewRoleStore(db)
		userStore = database.NewUserStore(db)
		helper    = handler.NewHandler(noteStore, postStore, roleStore, userStore)
		r         = gin.Default()
		v2        = r.Group("/api/v2")
	)

	/// users : admin(manage user, manage role) ,
	///  author(manage post, manage comments) , anonymous-unauthorized user

	///post api v2
	v2.POST("/posts", helper.HandleInsertPost)
	v2.GET("/posts", helper.HandleGetPosts)
	v2.GET("/posts/title/:title", helper.HandleGetPostByTitle)
	v2.DELETE("/posts/delete/:id", helper.HandleDeletePost)
	v2.PUT("/posts/update", helper.HandlePutPost)

	///http://localhost:8000/api/v2/posts
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
