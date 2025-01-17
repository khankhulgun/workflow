package seeds

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	krudModels "github.com/lambda-platform/lambda/krud/models"
	puzzleModels "github.com/lambda-platform/lambda/models"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Seed() {
	absolutePath := AbsolutePath()

	fileName := "vb_schemas.json"
	filePath := filepath.Join(absolutePath, fileName)

	vbSchemas, err := LoadVBSchemas(filePath)
	if err != nil {
		log.Fatalf("Failed to load seed data: %v", err)
	}

	err = SeedVBSchemas(vbSchemas)
	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	SeedLookupTables()
	fmt.Println("Seed data successfully loaded and updated into the database.")
}

func LoadVBSchemas(filePath string) ([]puzzleModels.VBSchema, error) {
	var vbSchemas []puzzleModels.VBSchema

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&vbSchemas)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return vbSchemas, nil
}

func SeedVBSchemas(vbSchemas []puzzleModels.VBSchema) error {
	for _, vb := range vbSchemas {
		var existingVB puzzleModels.VBSchema
		err := DB.DB.Where("name = ? AND type = ?", vb.Name, vb.Type).First(&existingVB).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := DB.DB.Create(&vb).Error; err != nil {
				return fmt.Errorf("error creating vb_schema with name %s: %w", vb.Name, err)
			}
			existingVB = vb
		} else if err != nil {
			return fmt.Errorf("error querying vb_schema with name %s: %w", vb.Name, err)
		}

		if vb.Type == "graphql" {
			continue
		}

		var krud krudModels.Krud
		err = DB.DB.Where("title = ?", vb.Name).First(&krud).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			krud = krudModels.Krud{
				Title:    vb.Name,
				Template: "canvas",
				Grid:     0,
				Form:     0,
			}
		} else if err != nil {
			return fmt.Errorf("error querying Krud with title %s: %w", vb.Name, err)
		}

		if vb.Type == "grid" {
			krud.Grid = int(existingVB.ID)
		} else if vb.Type == "form" {
			krud.Form = int(existingVB.ID)
		}

		if vb.Name == "Ажлын урсгал" {
			krud.Template = "window"
		}

		if err := DB.DB.Save(&krud).Error; err != nil {
			return fmt.Errorf("error saving Krud for vb_schema with name %s: %w", vb.Name, err)
		}
	}
	return nil
}

func AbsolutePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return filepath.Dir(fileName) + "/"
}
