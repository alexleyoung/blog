package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	"google.golang.org/api/option"
)

var AuthClient *auth.Client
var FirestoreClient *firestore.Client

func InitFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("secrets/service_account_key.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth: %v", err)
	}

	fsClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore: %v", err)
	}

	AuthClient = authClient
	FirestoreClient = fsClient
}
