package seeds

import (
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func SeedLookupTables() {
	SubjectTypes := []struct {
		SubjectType  string
		Subject      string
		SubjectOrder int
	}{
		{"TO_ROLE", "Системийн Хандах эрхээр", 1},
		{"DIRECT", "Системийн хэрэглэгч", 2},
		{"ANY_EMP", "Бүртгэлтэй хэрэглэгч", 3},
		{"CREATOR", "Боловсруулсан хэрэглэгч", 4},
	}

	for _, subjectType := range SubjectTypes {
		query := `
		INSERT INTO "workflow_and_process"."workflow_subject_types" ("subject_type", "subject", "subject_order")
		VALUES (?, ?, ?)
		ON CONFLICT ("subject_type") DO NOTHING;
		`
		if err := DB.DB.Exec(query, subjectType.SubjectType, subjectType.Subject, subjectType.SubjectOrder).Error; err != nil {
			log.Printf("Failed to seed workflow_subject_types: %v", err)
		}
	}
}
