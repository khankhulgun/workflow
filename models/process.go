package models

import "time"

type ProcessStatusHistory struct {
	ID                 string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RowID              string    `gorm:"column:row_id;type:uuid" json:"row_id"`
	WorkflowCategoryID string    `gorm:"column:workflow_category_id;type:uuid" json:"workflow_category_id"`
	Status             string    `gorm:"column:status" json:"status"`
	StatusType         string    `gorm:"column:status_type" json:"status_type"`
	StatusID           string    `gorm:"column:status_id;type:uuid" json:"status_id"`
	SubjectType        string    `gorm:"column:subject_type" json:"subject_type"`
	Description        *string   `gorm:"column:description" json:"description"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDone             int       `gorm:"column:is_done" json:"is_done"`
	Signature          *string   `gorm:"column:signature" json:"signature"`
	RoleID             *int      `gorm:"column:role_id" json:"role_id"`
	UserID             *string   `gorm:"column:user_id;type:uuid" json:"user_id"`
	OrgID              *string   `gorm:"column:org_id;type:uuid" json:"org_id"`
	StructID           *string   `gorm:"column:struct_id;type:uuid" json:"struct_id"`
	JobID              *string   `gorm:"column:job_id;type:uuid" json:"job_id"`
	EmpID              *string   `gorm:"column:emp_id;type:uuid" json:"emp_id"`
	Org                *string   `gorm:"column:org;type:uuid" json:"org"`
	Struct             *string   `gorm:"column:struct;type:uuid" json:"struct"`
	Job                *string   `gorm:"column:job;type:uuid" json:"job"`
	Emp                *string   `gorm:"column:emp;type:uuid" json:"emp"`
}

func (p *ProcessStatusHistory) TableName() string {
	return "workflow_and_process.process_status_history"
}

type ProcessVotingPeople struct {
	ID                 string     `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RowID              string     `gorm:"column:row_id;type:uuid" json:"row_id"`
	WorkflowCategoryID string     `gorm:"column:workflow_category_id;type:uuid" json:"workflow_category_id"`
	Approve            int        `gorm:"column:approve" json:"approve"`
	SignatureDate      *time.Time `gorm:"column:signature_date" json:"signature_date"`
	SignatureImage     *string    `gorm:"column:signature_image" json:"signature_image"`
	Description        *string    `gorm:"column:description" json:"description"`
	Voted              int        `gorm:"column:voted" json:"voted"`
	Recreate           *int       `gorm:"column:recreate" json:"recreate"`
	SubjectType        string     `gorm:"column:subject_type" json:"subject_type"`
	RoleID             *int       `gorm:"column:role_id" json:"role_id"`
	UserID             *string    `gorm:"column:user_id;type:uuid" json:"user_id"`
	OrgID              *string    `gorm:"column:org_id;type:uuid" json:"org_id"`
	StructID           *string    `gorm:"column:struct_id;type:uuid" json:"struct_id"`
	JobID              *string    `gorm:"column:job_id;type:uuid" json:"job_id"`
	EmpID              *string    `gorm:"column:emp_id;type:uuid" json:"emp_id"`
	Org                *string    `gorm:"column:org;type:uuid" json:"org"`
	Struct             *string    `gorm:"column:struct;type:uuid" json:"struct"`
	Job                *string    `gorm:"column:job;type:uuid" json:"job"`
	Emp                *string    `gorm:"column:emp;type:uuid" json:"emp"`
}

func (p *ProcessVotingPeople) TableName() string {
	return "workflow_and_process.process_voting_people"
}
