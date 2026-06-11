package resolver

// EmployeeInfo represents a unified employee structure across different database systems.
type EmployeeInfo struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	StructName string  `json:"struct_name"`
	JobName    string  `json:"job_name"`
	Avatar     *string `json:"avatar"`
}

// EmployeeResolver is the interface that each project must implement
// to provide employee lookup functionality for the workflow system.
// This allows the workflow package to work with different employee
// database structures (e.g., MIAT Oracle EO_EMPLOYEES vs khan_codes PostgreSQL view_emp_for_process).
type EmployeeResolver interface {
	// GetByEmpID returns employee info by employee ID.
	GetByEmpID(empID interface{}) (*EmployeeInfo, error)

	// GetByUserID returns employee info by user ID.
	GetByUserID(userID interface{}) (*EmployeeInfo, error)

	// GetUserIDByEmpID returns the user ID associated with an employee ID.
	GetUserIDByEmpID(empID interface{}) (interface{}, error)
}
