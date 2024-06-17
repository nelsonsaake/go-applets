package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func newClient() (*messaging.Client, error) {

	ctx := context.Background()

	// Initialize Firebase Admin SDK with your service account key JSON file.
	opt := option.WithCredentialsFile(".cred/.serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
		return nil, err
	}

	// Access the FCM client from the Firebase app.
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error getting FCM client: %v", err)
		return nil, err
	}

	return client, err
}
