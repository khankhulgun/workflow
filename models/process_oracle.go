package models

import (
	"time"

	"github.com/khankhulgun/workflow/resolver"
)

// ProcessStatusHistoryOracle — Oracle EOFFICE.PROCESS_STATUS_HISTORY
type ProcessStatusHistoryOracle struct {
	ID                 string                 `gorm:"column:ID;primaryKey" json:"id"`
	RowID              string                 `gorm:"column:ROW_ID" json:"row_id"`
	WorkflowCategoryID string                 `gorm:"column:WORKFLOW_CATEGORY_ID" json:"workflow_category_id"`
	Status             string                 `gorm:"column:STATUS" json:"status"`
	StatusType         string                 `gorm:"column:STATUS_TYPE" json:"status_type"`
	StatusID           string                 `gorm:"column:STATUS_ID" json:"status_id"`
	SubjectType        string                 `gorm:"column:SUBJECT_TYPE" json:"subject_type"`
	Description        *string                `gorm:"column:DESCRIPTION" json:"description"`
	CreatedAt          time.Time              `gorm:"column:CREATED_AT" json:"created_at"`
	UpdatedAt          time.Time              `gorm:"column:UPDATED_AT" json:"updated_at"`
	IsDone             int                    `gorm:"column:IS_DONE" json:"is_done"`
	Signature          *string                `gorm:"column:SIGNATURE" json:"signature"`
	StructID           *int                   `gorm:"column:STRUCT_ID" json:"struct_id"`
	EmpID              *int                   `gorm:"column:EMP_ID" json:"emp_id"`
	ActionNum          int                    `gorm:"column:ACTION_NUM" json:"action_num"`
	StatusAction       *string                `gorm:"column:STATUS_ACTION" json:"status_action"`
	Emp                *resolver.EmployeeInfo `gorm:"-" json:"emp"`
}

func (p *ProcessStatusHistoryOracle) TableName() string {
	return "PROCESS_STATUS_HISTORY"
}

// ProcessVotingPeopleOracle — Oracle EOFFICE.PROCESS_VOTING_PEOPLE
type ProcessVotingPeopleOracle struct {
	ID                 string                 `gorm:"column:ID;primaryKey" json:"id"`
	RowID              string                 `gorm:"column:ROW_ID" json:"row_id"`
	WorkflowCategoryID string                 `gorm:"column:WORKFLOW_CATEGORY_ID" json:"workflow_category_id"`
	UserID             *int                   `gorm:"column:USER_ID" json:"user_id"`
	Approve            int                    `gorm:"column:APPROVE" json:"approve"`
	SignatureDate      *time.Time             `gorm:"column:SIGNATURE_DATE" json:"signature_date"`
	SignatureImage     *string                `gorm:"column:SIGNATURE_IMAGE" json:"signature_image"`
	Description        *string                `gorm:"column:DESCRIPTION" json:"description"`
	Voted              int                    `gorm:"column:VOTED" json:"voted"`
	Recreate           *int                   `gorm:"column:RECREATE" json:"recreate"`
	StructID           *int                   `gorm:"column:STRUCT_ID" json:"struct_id"`
	EmpID              *int                   `gorm:"column:EMP_ID" json:"emp_id"`
	PreDescription     *string                `gorm:"column:PRE_DESCRIPTION" json:"pre_description"`
	Emp                *resolver.EmployeeInfo `gorm:"-" json:"emp"`
}

func (p *ProcessVotingPeopleOracle) TableName() string {
	return "PROCESS_VOTING_PEOPLE"
}

// ProcessStatusHistoryActionNumOracle — Oracle helper
type ProcessStatusHistoryActionNumOracle struct {
	ActionNum int `gorm:"column:ACTION_NUM" json:"action_num"`
}

func (p *ProcessStatusHistoryActionNumOracle) TableName() string {
	return "PROCESS_STATUS_HISTORY"
}
