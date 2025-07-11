package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/models"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
)

func History(c *fiber.Ctx) error {

	id := c.Params("id")

	var lastDone models.ProcessStatusHistory
	var statusHistories []models.ProcessStatusHistory
	var actionNumbers []models.ProcessStatusHistoryActionNum
	var votingPeople []models.ProcessVotingPeople

	DB.DB.Where("row_id = ? AND is_done = ? AND status_type = ?", id, 1, "END").Or("row_id = ? AND is_done = ? AND status_type = ?", id, 1, "CANCEL").First(&lastDone)

	if (lastDone.ID != "" && lastDone.StatusType == "END") || (lastDone.ID != "" && lastDone.StatusType == "CANCEL") {

		DB.DB.Where("row_id = ? AND is_done = ?", id, 1).Order("action_num, updated_at ASC").Find(&statusHistories)
		DB.DB.Select("action_num").Where("row_id = ? AND is_done = ?", id, 1).Order("action_num ASC").Group("action_num").Find(&actionNumbers)
	} else {
		DB.DB.Where("row_id = ? AND status_type != ? AND status_type != ?", id, "END", "CANCEL").Order("action_num, updated_at ASC").Find(&statusHistories)
		DB.DB.Select("action_num").Where("row_id = ? AND status_type != ? AND status_type != ?", id, "END", "CANCEL").Order("action_num ASC").Group("action_num").Find(&actionNumbers)
	}

	DB.DB.Where("row_id = ?", id).Order("signature_date ASC").Find(&votingPeople)

	for i, _ := range statusHistories {

		if statusHistories[i].EmpID != nil {
			if *statusHistories[i].EmpID != "" {
				var emp models.Employee
				DB.DB.Where("id = ?", statusHistories[i].EmpID).Find(&emp)

				statusHistories[i].Emp = &emp
			}
		}

	}

	for i, _ := range votingPeople {

		if votingPeople[i].EmpID != nil {
			if *votingPeople[i].EmpID != "" {
				var emp models.Employee
				DB.DB.Where("id = ?", statusHistories[i].EmpID).Find(&emp)

				votingPeople[i].Emp = &emp
			}
		}

	}
	return c.JSON(map[string]interface{}{
		"statusHistories": statusHistories,
		"actionNumbers":   actionNumbers,
		"votingPeople":    votingPeople,
	})

}

func HistoryWithUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var lastDone models.ProcessStatusHistory
	var statusHistories []models.ProcessStatusHistory
	var actionNumbers []models.ProcessStatusHistoryActionNum
	var votingPeople []models.ProcessVotingPeople

	DB.DB.Where("row_id = ? AND is_done = ? AND status_type = ?", id, 1, "END").Or("row_id = ? AND is_done = ? AND status_type = ?", id, 1, "CANCEL").First(&lastDone)

	if (lastDone.ID != "" && lastDone.StatusType == "END") || (lastDone.ID != "" && lastDone.StatusType == "CANCEL") {

		DB.DB.Where("row_id = ? AND is_done = ?", id, 1).Order("action_num, updated_at ASC").Find(&statusHistories)
		DB.DB.Select("action_num").Where("row_id = ? AND is_done = ?", id, 1).Order("action_num ASC").Group("action_num").Find(&actionNumbers)
	} else {
		DB.DB.Where("row_id = ? AND status_type != ? AND status_type != ?", id, "END", "CANCEL").Order("action_num, updated_at ASC").Find(&statusHistories)
		DB.DB.Select("action_num").Where("row_id = ? AND status_type != ? AND status_type != ?", id, "END", "CANCEL").Order("action_num ASC").Group("action_num").Find(&actionNumbers)
	}

	DB.DB.Where("row_id = ?", id).Order("signature_date ASC").Find(&votingPeople)

	for i, _ := range statusHistories {

		if statusHistories[i].UserID != "" {
			var user models.WorkflowUser
			DB.DB.Where("id = ?", statusHistories[i].UserID).Find(&user)

			statusHistories[i].User = &user
		}

	}

	for i, _ := range votingPeople {

		if votingPeople[i].UserID != "" {
			var user models.WorkflowUser
			DB.DB.Where("id = ?", votingPeople[i].UserID).Find(&user)

			votingPeople[i].User = &user
		}

	}
	return c.JSON(map[string]interface{}{
		"statusHistories": statusHistories,
		"actionNumbers":   actionNumbers,
		"votingPeople":    votingPeople,
	})

}

func Steps(c *fiber.Ctx) error {
	// Declare a slice to hold the ProcessSteps
	var ProcessSteps []models.ProcessStep

	// Use GORM to fetch ProcessSteps and preload SubProcessStepPorts
	if err := DB.DB.Preload("SubProcessStepPorts").Find(&ProcessSteps).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch process steps",
			"details": err.Error(),
		})
	}

	// Return the JSON response
	return c.JSON(ProcessSteps)
}

func GetWorkflowsByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")

	userPre, err := agentUtils.AuthUserObject(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}
	var workflows []models.Workflow
	if userPre["org_id"] != nil {
		orgID := userPre["org_id"].(string)
		if orgID != "" {
			if err := DB.DB.Where("category_id = ? AND org_id", categoryID, orgID).Find(&workflows).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Failed to fetch workflows",
					"details": err.Error(),
				})
			}

		}
	}

	return c.JSON(workflows)
}
