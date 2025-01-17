package migrations

import (
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func MigrateLookupTables() {
	// Create lut_line_style table
	createSubjectTypeTable := `
	CREATE TABLE IF NOT EXISTS "workflow_and_process"."workflow_subject_types" (
		"subject_type" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"subject" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
		"subject_order" int2 NOT NULL DEFAULT 0,
		CONSTRAINT "workflow_subject_types_pkey" PRIMARY KEY ("subject_type")
	);
	`
	if err := DB.DB.Exec(createSubjectTypeTable).Error; err != nil {
		log.Fatalf("Failed to create workflow_subject_types table: %v", err)
	}
}
