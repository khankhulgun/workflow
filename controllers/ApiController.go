package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/dbutil"
	"github.com/khankhulgun/workflow/models"
	"github.com/khankhulgun/workflow/resolver"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
)

var empResolver resolver.EmployeeResolver

func SetResolver(r resolver.EmployeeResolver) {
	empResolver = r
}

func History(c *fiber.Ctx) error {
	id := c.Params("id")

	if dbutil.IsOracle() {
		return historyOracle(c, id)
	}
	return historyPostgres(c, id)
}

func HistoryWithUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if dbutil.IsOracle() {
		// Oracle: HistoryWithUser uses same emp-based logic as History
		return historyOracle(c, id)
	}
	return historyWithUserPostgres(c, id)
}

// ===== PostgreSQL implementations =====

func historyPostgres(c *fiber.Ctx, id string) error {
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

	for i := range statusHistories {
		if statusHistories[i].EmpID != "" {
			var emp models.Employee
			DB.DB.Where("id = ?", statusHistories[i].EmpID).Find(&emp)
			statusHistories[i].Emp = &emp
		}
	}

	for i := range votingPeople {
		if votingPeople[i].EmpID != "" {
			var emp models.Employee
			DB.DB.Where("id = ?", votingPeople[i].EmpID).Find(&emp)
			votingPeople[i].Emp = &emp
		}
	}

	return c.JSON(map[string]interface{}{
		"statusHistories": statusHistories,
		"actionNumbers":   actionNumbers,
		"votingPeople":    votingPeople,
	})
}

func historyWithUserPostgres(c *fiber.Ctx, id string) error {
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

	for i := range statusHistories {
		if statusHistories[i].UserID != "" {
			var user models.WorkflowUser
			DB.DB.Where("id = ?", statusHistories[i].UserID).Find(&user)
			statusHistories[i].User = &user
		}
	}

	for i := range votingPeople {
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

// ===== Oracle implementations =====

func historyOracle(c *fiber.Ctx, id string) error {
	var lastDone models.ProcessStatusHistoryOracle
	var statusHistories []models.ProcessStatusHistoryOracle
	var actionNumbers []models.ProcessStatusHistoryActionNumOracle
	var votingPeople []models.ProcessVotingPeopleOracle

	DB.DB.Where("ROW_ID = ? AND IS_DONE = ? AND STATUS_TYPE = ?", id, 1, "END").Or("ROW_ID = ? AND IS_DONE = ? AND STATUS_TYPE = ?", id, 1, "CANCEL").First(&lastDone)

	if (lastDone.ID != "" && lastDone.StatusType == "END") || (lastDone.ID != "" && lastDone.StatusType == "CANCEL") {
		DB.DB.Where("ROW_ID = ? AND IS_DONE = ?", id, 1).Order("ACTION_NUM, UPDATED_AT ASC").Find(&statusHistories)
		DB.DB.Select("ACTION_NUM").Where("ROW_ID = ? AND IS_DONE = ?", id, 1).Order("ACTION_NUM ASC").Group("ACTION_NUM").Find(&actionNumbers)
	} else {
		DB.DB.Where("ROW_ID = ? AND STATUS_TYPE != ? AND STATUS_TYPE != ?", id, "END", "CANCEL").Order("ACTION_NUM, UPDATED_AT ASC").Find(&statusHistories)
		DB.DB.Select("ACTION_NUM").Where("ROW_ID = ? AND STATUS_TYPE != ? AND STATUS_TYPE != ?", id, "END", "CANCEL").Order("ACTION_NUM ASC").Group("ACTION_NUM").Find(&actionNumbers)
	}

	DB.DB.Where("ROW_ID = ?", id).Order("SIGNATURE_DATE ASC").Find(&votingPeople)

	// Use EmployeeResolver to enrich employee data
	if empResolver != nil {
		for i := range statusHistories {
			if statusHistories[i].EmpID != nil && *statusHistories[i].EmpID > 0 {
				emp, err := empResolver.GetByEmpID(*statusHistories[i].EmpID)
				if err == nil {
					statusHistories[i].Emp = emp
				}
			}
		}

		for i := range votingPeople {
			if votingPeople[i].EmpID != nil && *votingPeople[i].EmpID > 0 {
				emp, err := empResolver.GetByEmpID(*votingPeople[i].EmpID)
				if err == nil {
					votingPeople[i].Emp = emp
				}
			}
		}
	}

	return c.JSON(map[string]interface{}{
		"statusHistories": statusHistories,
		"actionNumbers":   actionNumbers,
		"votingPeople":    votingPeople,
	})
}

// ===== Shared handlers =====

func Steps(c *fiber.Ctx) error {
	var ProcessSteps []models.ProcessStep

	if err := DB.DB.Preload("SubProcessStepPorts").Find(&ProcessSteps).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch process steps",
			"details": err.Error(),
		})
	}

	return c.JSON(ProcessSteps)
}

func GetWorkflowsByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")

	userPre, err := agentUtils.AuthUserObject(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}

	if dbutil.IsOracle() {
		return getWorkflowsByCategoryOracle(c, categoryID, userPre)
	}
	return getWorkflowsByCategoryPostgres(c, categoryID, userPre)
}

func getWorkflowsByCategoryPostgres(c *fiber.Ctx, categoryID string, userPre map[string]interface{}) error {
	var workflows []models.Workflow
	if userPre["org_id"] != nil {
		orgID := userPre["org_id"].(string)
		if orgID != "" {
			if err := DB.DB.Where("category_id = ? AND org_id = ?", categoryID, orgID).Order("created_at DESC").Find(&workflows).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Failed to fetch workflows",
					"details": err.Error(),
				})
			}

			for i := range workflows {
				var votingPeople []models.WorkflowVotingPeople
				if err := DB.DB.Where("workflow_id = ?", workflows[i].ID).Find(&votingPeople).Error; err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to fetch voting people",
						"details": err.Error(),
					})
				}
				workflows[i].VotingPeople = votingPeople
			}
		}
	}

	return c.JSON(workflows)
}

func getWorkflowsByCategoryOracle(c *fiber.Ctx, categoryID string, userPre map[string]interface{}) error {
	var workflows []models.WorkflowOracle

	// Oracle doesn't have org_id in WORKFLOW table, so fetch by category only
	if err := DB.DB.Where("CATEGORY_ID = ?", categoryID).Order("CREATED_AT DESC").Find(&workflows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch workflows",
			"details": err.Error(),
		})
	}

	for i := range workflows {
		var votingPeople []models.WorkflowVotingPeopleOracle
		if err := DB.DB.Where("WORKFLOW_ID = ?", workflows[i].ID).Find(&votingPeople).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to fetch voting people",
				"details": err.Error(),
			})
		}
		workflows[i].VotingPeople = votingPeople
	}

	return c.JSON(workflows)
}
