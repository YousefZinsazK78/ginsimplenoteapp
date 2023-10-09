// package main

// import (
// 	"database/sql"
// 	"os"

// 	_ "github.com/lib/pq"

// 	"github.com/joho/godotenv"
// 	"github.com/yousefzinsazk78/simple_note_api/internal/database"
// 	"github.com/yousefzinsazk78/simple_note_api/internal/models"
// )

// func seedRoleTbl(roleStore database.RoleStorer) error {
// 	/// admin , author , unauthorizeduser

// 	var (
// 		admin = models.Role{
// 			Name:        "admin",
// 			Description: "admin description",
// 		}
// 		author = models.Role{
// 			Name:        "author",
// 			Description: "author description",
// 		}
// 		user = models.Role{
// 			Name:        "user",
// 			Description: "user description",
// 		}
// 	)

// 	if err := roleStore.InsertRole(admin); err != nil {
// 		return err
// 	}
// 	if err := roleStore.InsertRole(author); err != nil {
// 		return err
// 	}
// 	if err := roleStore.InsertRole(user); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func seedUserTbl(userStore database.UserStorer) error {
// 	var (
// 		user1 = models.User{
// 			RoleID:   1,
// 			Username: "yousef",
// 			Password: "123password",
// 			Email:    "yz.1378@gmail.com",
// 		}
// 		user2 = models.User{
// 			RoleID:   2,
// 			Username: "sina",
// 			Password: "password123",
// 			Email:    "sina.1378@gmail.com",
// 		}
// 		user3 = models.User{
// 			RoleID:   3,
// 			Username: "atta",
// 			Password: "passwordatta",
// 			Email:    "atta.1378@gmail.com",
// 		}
// 	)

// 	if err := userStore.InsertUser(user1); err != nil {
// 		return err
// 	}

// 	if err := userStore.InsertUser(user2); err != nil {
// 		return err
// 	}

// 	if err := userStore.InsertUser(user3); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	if err := godotenv.Load("./internal/config/.env"); err != nil {
// 		panic(err)
// 	}

// 	dbConn, err := sql.Open("postgres", os.Getenv("ConnectionStr"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dbConn.Close()

// 	if err := dbConn.Ping(); err != nil {
// 		panic(err)
// 	}

// 	var (
// 		db = database.NewDatabase(dbConn)
// 		// roleStore = database.NewRoleStore(db)

// 		// userstore = database.NewUserStore(db)
// 	)

// 	// seedRoleTbl(roleStore)
// 	// seedUserTbl(userstore)
// }
