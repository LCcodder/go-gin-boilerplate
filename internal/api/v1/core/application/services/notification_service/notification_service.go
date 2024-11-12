package notification_service

import (
	"context"

	"example.com/m/internal/api/v1/core/application/exceptions"
	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/go-fcm"
)

type NotificationService struct {
	fb *fcm.Client
}

func NewNotificationService(fb *fcm.Client) *NotificationService {
	return &NotificationService{
		fb: fb,
	}
}

func (ns *NotificationService) SendNotification(ctx context.Context, token string, data map[string]string) *exceptions.Error_ {
	_, err := ns.fb.Send(
		ctx,
		&messaging.Message{
			Token: token,
			Data:  data,
		},
	)
	if err != nil {
		return &exceptions.ErrServiceUnavailable
	}
	return nil
}
