package main

import (
	"billing-engine/internal/config"
	"billing-engine/internal/db"
	dbmodels "billing-engine/internal/models/db"
	"billing-engine/internal/repositories"
	"billing-engine/internal/schedulers"
	"billing-engine/internal/services"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	c, err := config.LoadConfig("./config", "dev.json")
	if err != nil {
		log.Fatalln("Error loading config", err)
	}

	// Initialize db connection
	db, err := db.ConnectDatabase(c.Postgresql)
	if err != nil {
		panic("Failed to connect to database")
	}

	var users []dbmodels.User
	err = db.Find(&users).Error
	if err == nil {
		if len(users) == 0 {
			err = db.Create(&dbmodels.User{
				ID:        1,
				Username:  "admin",
				CreatedAt: time.Now(),
			}).Error
			if err == nil {
				fmt.Println("User created successfully")
			}
		}
	}

	loanRepo := repositories.NewLoanRepository(db)
	billingRepo := repositories.NewBillingRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Initialize the billing service
	// billingService := services.NewBillingService(loanRepo, billingRepo)
	loanService := services.NewLoanService(loanRepo, userRepo)
	billingService := services.NewBillingService(loanRepo, billingRepo)

	// Initialize the router
	r := gin.Default()

	// Create a group
	router := r.Group(c.Server.PrefixPath)

	// Define routes
	router.POST("/loans", loanService.CreateLoan)
	// router.POST("/loans/:id/payments", loanService.MakePayment)
	router.GET("/loans/:id/outstanding", loanService.GetOutstanding)
	router.GET("/loans/:id/delinquent", loanService.IsDelinquent)

	router.POST("/billings", billingService.CreateBilling)

	schedulerService := schedulers.NewScheduler(loanRepo, billingRepo)

	// Start the scheduler
	schedulerService.StartScheduler()

	// Run the server
	port := fmt.Sprintf(":%d", c.Server.Port)
	r.Run(port)
}
