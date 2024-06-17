package main

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

var token = "cfD8zC3zEvLpKoSymPj2S7:APA91bGn2gtFkF-HVJg3qgQu7jdxXbTn__j8tkYK3GrUonQjKCYV2hIxpgjN9dObMqZ-N9yBBdNJSrnkQuLtM5EBdxejhxTl5rHQLWwKuc7H5qVp4LTCc-8XiSNdQwsXjCafz04vQMnN"

func sendMessage(client *messaging.Client) {
	// Context
	ctx := context.Background()

	// Create an FCM message.
	message := &messaging.Message{
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Notification: &messaging.Notification{
			Title: "Title",
			Body:  "Body",
		},
		Token: token,
	}

	// Send the message.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalf("Error sending FCM message: %v", err)
	}

	log.Printf("Successfully sent FCM message: %v", response)

}
