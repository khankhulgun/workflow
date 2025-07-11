package models

import (
	"gorm.io/gorm"
	"time"
)

type Workflow struct {
	ID           string                 `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID        string                 `gorm:"column:org_id;type:uuid" json:"org_id"`
	SystemTypeID string                 `gorm:"column:system_type_id;type:uuid" json:"system_type_id"`
	CategoryID   string                 `gorm:"column:category_id;type:uuid" json:"category_id"`
	FlowName     string                 `gorm:"column:flow_name" json:"flow_name"`
	Description  *string                `gorm:"column:description" json:"description"`
	FlowData     *string                `gorm:"column:flow_data" json:"flow_data"`
	CreatedAt    time.Time              `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time              `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt         `gorm:"column:deleted_at" json:"deleted_at"`
	VotingPeople []WorkflowVotingPeople `gorm:"-" json:"workflow_voting_people"`
}

func (w *Workflow) TableName() string {
	return "workflow_and_process.workflow"
}

type WorkflowCategory struct {
	ID           string `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SystemTypeID string `gorm:"column:system_type_id;type:uuid" json:"system_type_id"`
	Category     string `gorm:"column:category" json:"category"`
}

func (w *WorkflowCategory) TableName() string {
	return "workflow_and_process.workflow_category"
}

type WorkflowSystemType struct {
	ID         string `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SystemType string `gorm:"column:system_type" json:"system_type"`
}

func (w *WorkflowSystemType) TableName() string {
	return "workflow_and_process.workflow_system_type"
}

type WorkflowVotingPeople struct {
	ID          string  `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	WorkflowID  string  `gorm:"column:workflow_id;type:uuid" json:"workflow_id"`
	SubjectType string  `gorm:"column:subject_type" json:"subject_type"`
	UserID      *string `gorm:"column:user_id;type:uuid" json:"user_id"`
	RoleID      *int    `gorm:"column:role_id" json:"role_id"`
	OrgRoleID   *int    `gorm:"column:org_role_id" json:"org_role_id"`
	OrgID       *string `gorm:"column:org_id;type:uuid" json:"org_id"`
	StructID    *string `gorm:"column:struct_id;type:uuid" json:"struct_id"`
	JobID       *string `gorm:"column:job_id;type:uuid" json:"job_id"`
	EmpID       *string `gorm:"column:emp_id;type:uuid" json:"emp_id"`
	OrgRole     *int    `gorm:"column:org_role" json:"org_role"`
	Org         *string `gorm:"column:org;type:uuid" json:"org"`
	Struct      *string `gorm:"column:struct;type:uuid" json:"struct"`
	Job         *string `gorm:"column:job;type:uuid" json:"job"`
	Emp         *string `gorm:"column:emp;type:uuid" json:"emp"`
}

func (w *WorkflowVotingPeople) TableName() string {
	return "workflow_and_process.workflow_voting_people"
}
