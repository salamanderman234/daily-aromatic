package domain

import (
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type NotificationService interface {
	GetUserNotifications(user_id uint) ([]entity.Notification, error)
	UpdateReadFlagNotification(id uint) error
	UpdateNotificationMessage(id uint, message string) error
}

type NotificationRepository interface {
	GetNotificationByUserID(user_id uint) ([]entity.Notification, error)
	UpdateNotification(id uint, updateField model.Notification)
}
