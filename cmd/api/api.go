package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yousefzinsazk78/simple_note_api/internal/database"
	"github.com/yousefzinsazk78/simple_note_api/internal/handler"
	"github.com/yousefzinsazk78/simple_note_api/internal/utils"
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
		db           = database.NewDatabase(dbConn)
		noteStore    = database.NewNoteStore(db)
		postStore    = database.NewPostStore(db)
		roleStore    = database.NewRoleStore(db)
		userStore    = database.NewUserStore(db)
		commentStore = database.NewCommentStore(db)
		helper       = handler.NewHandler(noteStore, postStore, roleStore, userStore, commentStore)
		r            = gin.Default()
		authrouter   = r.Group("/auth/user")
		adminrouter  = r.Group("/admin")
		v2           = r.Group("/api/v2")
	)
	//8 mb
	r.MaxMultipartMemory = 8 << 20
	r.Static("/public/images", "./public/images")

	authrouter.POST("/register", helper.Register)
	authrouter.POST("/login", helper.Login)

	v2.Use(utils.JWTAuthAuthor(userStore))
	//post api
	v2.POST("/posts", helper.HandleInsertPost)
	v2.GET("/posts", helper.HandleGetPosts)
	v2.GET("/posts/title/:title", helper.HandleGetPostByTitle)
	v2.DELETE("/posts/delete/:id", helper.HandleDeletePost)
	v2.PUT("/posts/update", helper.HandlePutPost)
	v2.POST("/upload/:id", helper.HandleUpload)
	//comment api
	v2.POST("/comment", helper.HandleInsertComment)
	v2.GET("/comments", helper.HandleGetComments)
	v2.GET("/comments/:id", helper.HandleGetCommentsByID)
	v2.GET("/comments/body/:body", helper.HandleGetCommentsByBody)
	v2.PUT("/comment/:id", helper.HandleUpdateComments)
	v2.DELETE("/comment/delete/:id", helper.HandleDeleteComment)

	adminrouter.Use(utils.JWTAuth(userStore))
	adminrouter.GET("/users", helper.GetUsers)
	adminrouter.GET("/users/:id", helper.GetUser)
	adminrouter.PUT("/user/:id", helper.UpdateUser)

	///http://localhost:8000/api/v2/posts
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
