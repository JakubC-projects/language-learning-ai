package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/samber/lo"
)

var Client *firestore.Client

func init() {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: config.C.ProjectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	Client = client
}

func GetChats(ctx context.Context, userId string) ([]models.Chat, error) {
	data, err := Client.Collection("chat").Where("UserId", "==", userId).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("cannot fetch chats: %w", err)
	}
	return lo.Map(data, func(d *firestore.DocumentSnapshot, _ int) models.Chat {
		var chat models.Chat
		d.DataTo(&chat)
		return chat
	}), nil
}

func GetMessages(ctx context.Context, chatId string) ([]models.Message, error) {
	data, err := Client.Collection("chat").Where("ChatId", "==", chatId).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("cannot fetch chats: %w", err)
	}
	return lo.Map(data, func(d *firestore.DocumentSnapshot, _ int) models.Message {
		var chat models.Message
		d.DataTo(&chat)
		return chat
	}), nil
}
