package models

import (
	"gorm.io/gorm"
	"time"
)

type Example struct {
	ID                 string         `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ParentID           *string        `gorm:"column:parent_id" json:"parent_id"`
	Name               string         `gorm:"column:name" json:"name"`
	WorkflowCategoryID string         `gorm:"column:workflow_category_id" json:"workflow_category_id"`
	WorkflowID         string         `gorm:"column:workflow_id" json:"workflow_id"`
	StatusID           string         `gorm:"column:status_id" json:"status_id"`
	StatusType         string         `gorm:"column:status_type" json:"status_type"`
	Status             string         `gorm:"column:status" json:"status"`
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	OrgID              *string        `gorm:"column:org_id" json:"org_id"`
	StructID           *string        `gorm:"column:struct_id" json:"struct_id"`
	EmpID              *string        `gorm:"column:emp_id" json:"emp_id"`
	UserID             *string        `gorm:"column:user_id" json:"user_id"`
}

func (e *Example) TableName() string {
	return "public.example"
}

type ExampleFullAccessUser struct {
	ID         string         `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID      *string        `gorm:"column:org_id" json:"org_id"`
	StructID   *string        `gorm:"column:struct_id" json:"struct_id"`
	EmpID      *string        `gorm:"column:emp_id" json:"emp_id"`
	UserID     *string        `gorm:"column:user_id" json:"user_id"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ShowAll    int            `gorm:"column:show_all" json:"show_all"`
	ChangeStep int            `gorm:"column:change_step" json:"change_step"`
}

func (e *ExampleFullAccessUser) TableName() string {
	return "public.example_full_access_user"
}

type ExampleChildConfig struct {
	ID       string `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ParentID string `gorm:"column:parent_id" json:"parent_id"`
	ChildID  string `gorm:"column:child_id" json:"child_id"`
	Title    string `gorm:"column:title" json:"title"`
}

func (e *ExampleChildConfig) TableName() string {
	return "public.example_child_config"
}
