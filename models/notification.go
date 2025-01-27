package models

import (
	"gorm.io/gorm"
	"time"
)

type CurrentStep struct {
	Notify       bool                  `json:"notify"`
	CreatorID    int                   `json:"creator_id"`
	FormID       string                `json:"form_id"`
	SchemaID     int                   `json:"schema_id"`
	ID           string                `json:"id"`
	RowID        string                `json:"row_id"`
	RoleID       int                   `json:"role_id"`
	UserID       string                `json:"user_id"`
	OrgRoleID    int                   `json:"org_role_id"`
	OrgID        int                   `json:"org_id"`
	StructID     string                `json:"struct_id"`
	JobID        string                `json:"job_id"`
	EmpID        string                `json:"emp_id"`
	Status       string                `json:"status"`
	StatusType   string                `json:"status_type"`
	StatusID     string                `json:"status_id"`
	SubjectType  string                `json:"subject_type"`
	VotingPeople []ProcessVotingPeople `json:"votingPeople"`
}

type NotificationConfig struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SchemaID    int    `gorm:"column:schema_id" json:"schema_id"`
	RequestName string `gorm:"column:request_name" json:"request_name"`
	AdminGrid   string `gorm:"column:admin_grid" json:"admin_grid"`
	UserGrid    string `gorm:"column:user_grid" json:"user_grid"`
}

func (n *NotificationConfig) TableName() string {
	return "public.notification_config"
}

type Users struct {
	ID             string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt      *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Status         *string        `gorm:"column:status" json:"status"`
	Role           int            `gorm:"column:role" json:"role"`
	Login          string         `gorm:"column:login" json:"login"`
	Email          string         `gorm:"column:email" json:"email"`
	RegisterNumber string         `gorm:"column:register_number" json:"register_number"`
	Avatar         *string        `gorm:"column:avatar" json:"avatar"`
	FirstName      *string        `gorm:"column:first_name" json:"first_name"`
	LastName       *string        `gorm:"column:last_name" json:"last_name"`
	Birthday       *time.Time     `gorm:"column:birthday" json:"birthday"`
	Phone          *string        `gorm:"column:phone" json:"phone"`
	Gender         *string        `gorm:"column:gender" json:"gender"`
}

func (u *Users) TableName() string {
	return "public.users"
}

type Employee struct {
	ID         string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgID      int            `gorm:"column:org_id" json:"org_id"`
	StructID   string         `gorm:"column:struct_id" json:"struct_id"`
	JobID      string         `gorm:"column:job_id" json:"job_id"`
	StatusID   int            `gorm:"column:status_id" json:"status_id"`
	UserID     string         `gorm:"column:user_id" json:"user_id"`
	FirstName  string         `gorm:"column:first_name" json:"first_name"`
	LastName   string         `gorm:"column:last_name" json:"last_name"`
	Register   string         `gorm:"column:register" json:"register"`
	Gender     *string        `gorm:"column:gender" json:"gender"`
	Birthday   *time.Time     `gorm:"column:birthday" json:"birthday"`
	ProvinceID int            `gorm:"column:province_id" json:"province_id"`
	SoumID     int            `gorm:"column:soum_id" json:"soum_id"`
	BaghID     int            `gorm:"column:bagh_id" json:"bagh_id"`
	Address    *string        `gorm:"column:address" json:"address"`
	Email      *string        `gorm:"column:email" json:"email"`
	Phone      string         `gorm:"column:phone" json:"phone"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	EmpID      *string        `gorm:"column:emp_id" json:"emp_id"`
}

func (e *Employee) TableName() string {
	return "organization.employee"
}

type Job struct {
	ID             string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgID          int            `gorm:"column:org_id" json:"org_id"`
	StructID       string         `gorm:"column:struct_id" json:"struct_id"`
	Job            string         `gorm:"column:job" json:"job"`
	JobDescription *string        `gorm:"column:job_description" json:"job_description"`
	Permissions    *string        `gorm:"column:permissions" json:"permissions"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (j *Job) TableName() string {
	return "organization.job"
}

type Struct struct {
	ID                string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgID             int            `gorm:"column:org_id" json:"org_id"`
	Struct            string         `gorm:"column:struct" json:"struct"`
	ParentStructID    *string        `gorm:"column:parent_struct_id" json:"parent_struct_id"`
	StructDescription *string        `gorm:"column:struct_description" json:"struct_description"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	StructType        string         `gorm:"column:struct_type" json:"struct_type"`
	ChildOrgID        *int           `gorm:"column:child_org_id" json:"child_org_id"`
}

func (s *Struct) TableName() string {
	return "organization.struct"
}

type BaiguullagaBurtgel struct {
	ID                         int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	BaiguullagaNer             string         `gorm:"column:baiguullaga_ner" json:"baiguullaga_ner"`
	BaiguullagaRegister        int            `gorm:"column:baiguullaga_register" json:"baiguullaga_register"`
	UlsBurtgelGerchilgeeDugaar *string        `gorm:"column:uls_burtgel_gerchilgee_dugaar" json:"uls_burtgel_gerchilgee_dugaar"`
	UlsBurtgelGerchilgee       *string        `gorm:"column:uls_burtgel_gerchilgee" json:"uls_burtgel_gerchilgee"`
	EOgnoo                     *time.Time     `gorm:"column:e_ognoo" json:"e_ognoo"`
	TzGerchilgeeDugaar         *string        `gorm:"column:tz_gerchilgee_dugaar" json:"tz_gerchilgee_dugaar"`
	TzEOgnoo                   *time.Time     `gorm:"column:tz_e_ognoo" json:"tz_e_ognoo"`
	TzDOgnoo                   *time.Time     `gorm:"column:tz_d_ognoo" json:"tz_d_ognoo"`
	Aimagid                    *int           `gorm:"column:aimagid" json:"aimagid"`
	Sumid                      *int           `gorm:"column:sumid" json:"sumid"`
	Bagid                      *int           `gorm:"column:bagid" json:"bagid"`
	Gudamj                     *string        `gorm:"column:gudamj" json:"gudamj"`
	Bair                       *string        `gorm:"column:bair" json:"bair"`
	ZahiralNer                 *string        `gorm:"column:zahiral_ner" json:"zahiral_ner"`
	Mail                       *string        `gorm:"column:mail" json:"mail"`
	Web                        *string        `gorm:"column:web" json:"web"`
	UtasOne                    *int           `gorm:"column:utas_one" json:"utas_one"`
	UtasTwo                    *int           `gorm:"column:utas_two" json:"utas_two"`
	CreatedAt                  *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt                  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt                  *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID                     *int           `gorm:"column:user_id" json:"user_id"`
	TzEseh                     *int           `gorm:"column:tz_eseh" json:"tz_eseh"`
	RoleID                     *int           `gorm:"column:role_id" json:"role_id"`
	CitizenRoleID              *int           `gorm:"column:citizen_role_id" json:"citizen_role_id"`
	OrgTypeID                  *int           `gorm:"column:org_type_id" json:"org_type_id"`
	MainRoleID                 *int           `gorm:"column:main_role_id" json:"main_role_id"`
	OrgAvartar                 *string        `gorm:"column:org_avartar" json:"org_avartar"`
	OrgCover                   *string        `gorm:"column:org_cover" json:"org_cover"`
}

func (b *BaiguullagaBurtgel) TableName() string {
	return "organization.baiguullaga_burtgel"
}

type ViewOrganizationUsers struct {
	ID                         *string `gorm:"column:id" json:"id"`
	OrgID                      *int    `gorm:"column:org_id" json:"org_id"`
	UserID                     string  `gorm:"column:user_id" json:"user_id"`
	Access                     *string `gorm:"column:access" json:"access"`
	BaiguullagaNer             *string `gorm:"column:baiguullaga_ner" json:"baiguullaga_ner"`
	BaiguullagaRegister        *int    `gorm:"column:baiguullaga_register" json:"baiguullaga_register"`
	UlsBurtgelGerchilgeeDugaar *string `gorm:"column:uls_burtgel_gerchilgee_dugaar" json:"uls_burtgel_gerchilgee_dugaar"`
	UlsBurtgelGerchilgee       *string `gorm:"column:uls_burtgel_gerchilgee" json:"uls_burtgel_gerchilgee"`
	Aimagid                    *int    `gorm:"column:aimagid" json:"aimagid"`
	Sumid                      *int    `gorm:"column:sumid" json:"sumid"`
	Bagid                      *int    `gorm:"column:bagid" json:"bagid"`
	UtasOne                    *int    `gorm:"column:utas_one" json:"utas_one"`
	RoleID                     *int    `gorm:"column:role_id" json:"role_id"`
	OrgAvartar                 *string `gorm:"column:org_avartar" json:"org_avartar"`
	OrgCover                   *string `gorm:"column:org_cover" json:"org_cover"`
}

func (v *ViewOrganizationUsers) TableName() string {
	return "organization.view_organization_users"
}
