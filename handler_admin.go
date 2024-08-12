package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createAdmin(login string, email string, c context.Context) {
	completed, err := apiCfg.DB.CheckAdmin(c)
	if err != nil {
		log.Fatal("Couldn't check Admin:", err)
		return
	}
	if completed {
		fmt.Println("Admin exist. Creating finished.")
		return
	}
	fmt.Println("Admin doesn't exist. Creating...")

	created_admin, err := apiCfg.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      login,
		Email:     email,
		Usertype:  "admin",
	})
	if err != nil {
		log.Fatal("Couldn't create Admin:", err)
		return
	}

	err = apiCfg.DB.MarkAdminAsCreated(c)
	if err != nil {
		log.Fatal("Couldn't mark Admin as created:", err)
		return
	}

	fmt.Println("Created admin with this ApiKey(collect it):", created_admin.ApiKey)
}
