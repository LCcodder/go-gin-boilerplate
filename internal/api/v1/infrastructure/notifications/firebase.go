package notifications

import (
	"context"
	"log"

	fcm "github.com/appleboy/go-fcm"
)

var FirebaseClient *fcm.Client

func InitClient(ctx context.Context) {
	client, err := fcm.NewClient(
		ctx,
		fcm.WithCredentialsFile("path/to/serviceAccountKey.json"),
		// initial with service account
		// fcm.WithServiceAccount("my-client-id@my-project-id.iam.gserviceaccount.com"),
	)
	if err != nil {
		log.Fatal(err)
	}

	FirebaseClient = client
}
