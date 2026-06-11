package models

import (
	"time"

	"gorm.io/gorm"
)

// WorkflowOracle — Oracle EOFFICE.WORKFLOW
type WorkflowOracle struct {
	ID           string                       `gorm:"column:ID;primaryKey" json:"id"`
	SystemTypeID string                       `gorm:"column:SYSTEM_TYPE_ID" json:"system_type_id"`
	CategoryID   string                       `gorm:"column:CATEGORY_ID" json:"category_id"`
	FlowName     string                       `gorm:"column:FLOW_NAME" json:"flow_name"`
	Description  *string                      `gorm:"column:DESCRIPTION" json:"description"`
	FlowData     *string                      `gorm:"column:FLOW_DATA" json:"flow_data"`
	CreatedAt    time.Time                    `gorm:"column:CREATED_AT" json:"created_at"`
	UpdatedAt    time.Time                    `gorm:"column:UPDATED_AT" json:"updated_at"`
	DeletedAt    gorm.DeletedAt               `gorm:"column:DELETED_AT" json:"deleted_at"`
	VotingPeople []WorkflowVotingPeopleOracle `gorm:"-" json:"workflow_voting_people"`
}

func (w *WorkflowOracle) TableName() string {
	return "WORKFLOW"
}

// WorkflowCategoryOracle — Oracle EOFFICE.WORKFLOW_CATEGORY
type WorkflowCategoryOracle struct {
	ID           string  `gorm:"column:ID;primaryKey" json:"id"`
	SystemTypeID string  `gorm:"column:SYSTEM_TYPE_ID" json:"system_type_id"`
	Category     string  `gorm:"column:CATEGORY" json:"category"`
	AdminURL     *string `gorm:"column:ADMIN_URL" json:"admin_url"`
	UserURL      *string `gorm:"column:USER_URL" json:"user_url"`
}

func (w *WorkflowCategoryOracle) TableName() string {
	return "WORKFLOW_CATEGORY"
}

// WorkflowSystemTypeOracle — Oracle EOFFICE.WORKFLOW_SYSTEM_TYPE
type WorkflowSystemTypeOracle struct {
	ID         string `gorm:"column:ID;primaryKey" json:"id"`
	SystemType string `gorm:"column:SYSTEM_TYPE" json:"system_type"`
}

func (w *WorkflowSystemTypeOracle) TableName() string {
	return "WORKFLOW_SYSTEM_TYPE"
}

// WorkflowVotingPeopleOracle — Oracle EOFFICE.WORKFLOW_VOTING_PEOPLE
type WorkflowVotingPeopleOracle struct {
	ID          string `gorm:"column:ID;primaryKey" json:"id"`
	WorkflowID  string `gorm:"column:WORKFLOW_ID" json:"workflow_id"`
	SubjectType string `gorm:"column:SUBJECT_TYPE" json:"subject_type"`
	UserID      *int   `gorm:"column:USER_ID" json:"user_id"`
	StructID    *int   `gorm:"column:STRUCT_ID" json:"struct_id"`
	EmpID       *int   `gorm:"column:EMP_ID" json:"emp_id"`
}

func (w *WorkflowVotingPeopleOracle) TableName() string {
	return "WORKFLOW_VOTING_PEOPLE"
}
