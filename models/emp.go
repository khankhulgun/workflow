package models

type Employee struct {
	ID             string  `gorm:"column:id" json:"id"`
	OrgID          string  `gorm:"column:org_id" json:"org_id"`
	StructID       string  `gorm:"column:struct_id" json:"struct_id"`
	JobID          string  `gorm:"column:job_id" json:"job_id"`
	StatusID       int     `gorm:"column:status_id" json:"status_id"`
	UserID         string  `gorm:"column:user_id" json:"user_id"`
	FirstName      string  `gorm:"column:first_name" json:"first_name"`
	LastName       string  `gorm:"column:last_name" json:"last_name"`
	Email          string  `gorm:"column:email" json:"email"`
	Phone          string  `gorm:"column:phone" json:"phone"`
	ShiftID        string  `gorm:"column:shift_id" json:"shift_id"`
	WebImage       *string `gorm:"column:web_image" json:"web_image"`
	Organization   string  `gorm:"column:organization" json:"organization"`
	Struct         string  `gorm:"column:struct" json:"struct"`
	Job            string  `gorm:"column:job" json:"job"`
	EmployeeStatus string  `gorm:"column:employee_status" json:"employee_status"`
	Login          string  `gorm:"column:login" json:"login"`
	Avatar         *string `gorm:"column:avatar" json:"avatar"`
}

func (v *Employee) TableName() string {
	return "organization.view_emp_for_process"
}
