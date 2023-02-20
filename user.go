package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"time"
)

type (
	Agent struct {
		Roles []Role
	}
	Role struct {
		ID          int64 `firestore:"id"`
		Level       int   `firestore:"level"`
		GrantedTime int64 `firestore:"granted_time"`
	}

	User struct {
		// ignore id from firestore
		Status int `firestore:"status"`
		Agent
	}
)

func testCase01(client *firestore.Client) {
	_, err := client.Collection("users").Doc("user1").Set(context.Background(),
		User{
			Status: 1,
			Agent: Agent{
				Roles: []Role{
					{
						ID:          123124214,
						Level:       1,
						GrantedTime: time.Now().UnixMilli(),
					},
				},
			},
		})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}

func testCase02(client *firestore.Client) {
	snap, err := client.Collection("users").Doc("user1").Get(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	var user User
	err = snap.DataTo(&user)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	log.Printf("User: %v", user)
}
