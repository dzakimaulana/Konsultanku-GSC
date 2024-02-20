package main

import (
	"konsultanku-v2/internal/comments"
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/internal/firebase/storage"
	"konsultanku-v2/internal/msme"
	"konsultanku-v2/internal/problems"
	"konsultanku-v2/internal/student"
	"konsultanku-v2/pkg/databases"
	"konsultanku-v2/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	ah, se, err := databases.FirebaseConn()
	if err != nil {
		panic(err.Error())
	}

	db, err := databases.DatabaseConn()
	if err != nil {
		panic(err.Error())
	}
	app := fiber.New()
	app.Use(logger.New())

	// Initialize session middleware
	store := session.New(session.Config{
		Expiration:        3600,
		KeyLookup:         "cookie:my-session",
		CookieSecure:      true,
		CookieSessionOnly: true,
	})

	storageRepo := storage.NewRepo(se.GetStorage())

	authRepo := auth.NewRepo(ah.GetAuth())
	authSvc := auth.NewSvc(authRepo, storageRepo)
	authHandler := auth.NewHandler(authSvc, store)
	routes.AuthRoute(authHandler, app)

	problemRepo := problems.NewRepo(db.GetDB())
	problemSvc := problems.NewSvc(problemRepo)

	commentRepo := comments.NewRepo(db.GetDB())
	commentSvc := comments.NewSvc(commentRepo)

	msmeRepo := msme.NewRepo(db.GetDB())
	msmeSvc := msme.NewSvc(msmeRepo, authRepo)
	msmeHandler := msme.NewHandler(msmeSvc, problemSvc, commentSvc, store)
	routes.MsmeRoute(msmeHandler, app)

	studentRepo := student.NewRepo(db.GetDB())
	studentSvc := student.NewSvc(studentRepo, authRepo)
	studentHandler := student.NewHandler(studentSvc, problemSvc, commentSvc, store)
	routes.StudentRoute(studentHandler, app)

	err = app.Listen(":8080")
	if err != nil {
		panic("Error starting server")
	}

}
