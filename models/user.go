package models

import (
	"gorm.io/gorm"
	"time"
)

type WorkflowUser struct {
	ID          string         `gorm:"column:id" json:"id"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Role        int            `gorm:"column:role" json:"role"`
	Login       string         `gorm:"column:login" json:"login"`
	Email       string         `gorm:"column:email" json:"email"`
	Avatar      *string        `gorm:"column:avatar" json:"avatar"`
	FirstName   *string        `gorm:"column:first_name" json:"first_name"`
	LastName    *string        `gorm:"column:last_name" json:"last_name"`
	Phone       *string        `gorm:"column:phone" json:"phone"`
	Gender      *string        `gorm:"column:gender" json:"gender"`
	DisplayName *string        `gorm:"column:display_name" json:"display_name"`
}

func (v *WorkflowUser) TableName() string {
	return "public.view_users"
}
