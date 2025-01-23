package migrations

import (
	"github.com/khankhulgun/workflow/models"
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func Migrate() {
	createSchema := `
	CREATE SCHEMA IF NOT EXISTS workflow_and_process;
	`

	err := DB.DB.Exec(createSchema).Error
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	err = DB.DB.AutoMigrate(
		&models.Workflow{},
		&models.WorkflowCategory{},
		&models.WorkflowSystemType{},
		&models.WorkflowVotingPeople{},
		&models.ProcessStatusHistory{},
		&models.ProcessVotingPeople{},
		&models.ProcessStep{},
		&models.SubProcessStepPort{},
		&models.Example{},
		&models.ExampleFullAccessUser{},
		&models.ExampleChildConfig{},
	)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createView := `
	CREATE OR REPLACE VIEW "workflow_and_process"."view_workflow" AS
		SELECT workflow.id,
		workflow.system_type_id,
		workflow.category_id,
		workflow.flow_name,
		workflow.description,
		workflow.flow_data,
		workflow.created_at,
		workflow.updated_at,
		workflow.deleted_at,
		workflow_category.category,
		workflow_system_type.system_type
		FROM workflow_and_process.workflow
	LEFT JOIN workflow_and_process.workflow_category ON workflow.category_id = workflow_category.id
	LEFT JOIN workflow_and_process.workflow_system_type ON workflow.system_type_id = workflow_system_type.id;

	CREATE OR REPLACE VIEW "workflow_and_process"."view_workflow_category" AS
		SELECT workflow_category.id,
		workflow_category.system_type_id,
		workflow_category.category,
		workflow_system_type.system_type
		FROM workflow_and_process.workflow_category
	LEFT JOIN workflow_and_process.workflow_system_type ON workflow_category.system_type_id = workflow_system_type.id;

	CREATE OR REPLACE VIEW "public"."view_users" AS
		SELECT users.id,
		users.created_at,
		users.updated_at,
		users.deleted_at,
		users.status,
		users.role,
		users.login,
		users.email,
		users.register_number,
		users.avatar,
		users.bio,
		users.first_name,
		users.last_name,
		users.birthday,
		users.phone,
		users.gender,
		roles.display_name
	FROM users
	LEFT JOIN roles ON users.role = roles.id;
	`

	err = DB.DB.Exec(createView).Error
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	}

	MigrateLookupTables()
}
