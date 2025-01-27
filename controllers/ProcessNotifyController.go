package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/models"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/config"
	notifyHandler "github.com/lambda-platform/lambda/notify/handler"
	modelsModels "github.com/lambda-platform/lambda/notify/models"
)

func SendNotification(c *fiber.Ctx) error {

	var step models.CurrentStep
	var notification modelsModels.NotificationData
	var notifyConfig models.NotificationConfig

	var users []models.Users
	var employees []models.Employee
	var owners []models.ViewOrganizationUsers

	var userIDs []string

	if err := json.Unmarshal(c.Body(), &step); err != nil {
		return err
	}

	if step.StatusType != "START" && step.StatusType != "" && step.Notify {

		if err := DB.DB.Where("schema_id = ?", step.SchemaID).Order("id DESC").First(&notifyConfig).Error; err != nil {
			return err
		}

		SendCreatorNotification(step, notifyConfig)

		if step.SubjectType != "ANY_EMP" && step.SubjectType != "CURRENT" && step.SubjectType != "CREATOR" {

			FCMNotification := modelsModels.FCMNotification{
				Title: "Мэдэгдэл",
				Body:  fmt.Sprintf("Танд шийдвэрлэх \"%s\" хүсэлт ирлээ.", notifyConfig.RequestName),
			}
			notification.Notification = FCMNotification

			FCMOptions := modelsModels.FCMOptions{
				Link: config.LambdaConfig.Domain + notifyConfig.AdminGrid,
			}

			FCMData := map[string]interface{}{
				"title":      "Мэдэгдэл",
				"body":       fmt.Sprintf("Танд шийдвэрлэх \"%s\" хүсэлт ирлээ.", notifyConfig.RequestName),
				"first_name": "",
				"sound":      config.LambdaConfig.Domain + config.LambdaConfig.Notify.Sound,
				"icon":       config.LambdaConfig.Favicon,
				"link":       notifyConfig.AdminGrid,
			}

			switch step.SubjectType {

			case "DIRECT":

				notification.UsersUUID = []string{step.UserID}

			case "TO_ROLE":

				if err := DB.DB.Where("role = ?", step.RoleID).Find(&users).Error; err != nil {
					return err
				}

				for _, user := range users {
					userIDs = append(userIDs, user.ID)
				}
				notification.UsersUUID = userIDs

			case "VOTERS":

				notification.Notification.Body = fmt.Sprintf("Танд санал өгөх \"%s\" хүсэлт ирлээ.", notifyConfig.RequestName)
				FCMData["body"] = fmt.Sprintf("Танд санал өгөх \"%s\" хүсэлт ирлээ.", notifyConfig.RequestName)

				for _, votingPerson := range step.VotingPeople {

					switch votingPerson.SubjectType {

					case "", "DIRECT":

						userIDs = append(userIDs, votingPerson.UserID)

					case "TO_ORG_EMPLOYEE":

						if err := DB.DB.Where("id = ?", votingPerson.EmpID).Find(&employees).Error; err != nil {
							return err
						}
						userIDs = append(userIDs, employees[0].UserID)

					case "TO_ORG":

						if err := DB.DB.Where("org_id = ? AND access = 'owner'", votingPerson.OrgID).Find(&owners).Error; err != nil {
							return err
						}

						userIDs = append(userIDs, owners[0].UserID)

					default:

						return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
							"error": "User not found!",
						})

					}
				}
				notification.UsersUUID = userIDs

			default:

				switch {

				case step.EmpID != "":

					if err := DB.DB.Where("id = ?", step.EmpID).Find(&employees).Error; err != nil {
						return err
					}
					notification.UsersUUID = []string{employees[0].UserID}

				case step.JobID != "":

					if err := DB.DB.Where("job_id = ?", step.JobID).Find(&employees).Error; err != nil {
						return err
					}

					for _, employee := range employees {
						userIDs = append(userIDs, employee.UserID)
					}
					notification.UsersUUID = userIDs

				case step.StructID != "":

					if err := DB.DB.Where("struct_id = ?", step.StructID).Find(&employees).Error; err != nil {
						return err
					}

					for _, employee := range employees {
						userIDs = append(userIDs, employee.UserID)
					}
					notification.UsersUUID = userIDs

				case step.OrgID != 0:

					if err := DB.DB.Where("org_id = ? AND access = 'owner'", step.OrgID).Find(&owners).Error; err != nil {
						return err
					}

					notification.UsersUUID = []string{owners[0].UserID}

				case step.OrgRoleID != 0:

					if err := DB.DB.Where("role_id = ? AND access = 'owner'", step.OrgRoleID).Find(&owners).Error; err != nil {
						return err
					}

					for _, owner := range owners {
						userIDs = append(userIDs, owner.UserID)
					}
					notification.UsersUUID = userIDs

				default:

					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "User not found!",
					})

				}
			}
			notifyHandler.CreateNotification(notification, FCMOptions, FCMData)
			return c.JSON(fiber.Map{"message": "Notification sent successfully!"})

		}
		return c.JSON(fiber.Map{"message": "No notifications sent to admin"})

	}
	return c.JSON(fiber.Map{"message": "No notifications sent"})

}

func SendCreatorNotification(step models.CurrentStep, notifyConfig models.NotificationConfig) {

	FCMNotification := modelsModels.FCMNotification{
		Title: "Мэдэгдэл",
		Body:  fmt.Sprintf("Таны илгээсэн \"%s\" хүсэлт \"%s\" төлөвт орлоо.", notifyConfig.RequestName, step.Status),
	}

	FCMOptions := modelsModels.FCMOptions{
		Link: config.LambdaConfig.Domain + notifyConfig.UserGrid,
	}

	FCMData := map[string]interface{}{
		"title":      "Мэдэгдэл",
		"body":       fmt.Sprintf("Таны илгээсэн \"%s\" хүсэлт \"%s\" төлөвт орлоо.", notifyConfig.RequestName, step.Status),
		"first_name": "",
		"sound":      config.LambdaConfig.Domain + config.LambdaConfig.Notify.Sound,
		"icon":       config.LambdaConfig.Favicon,
		"link":       notifyConfig.UserGrid,
	}

	notification := modelsModels.NotificationData{
		Users:        []int{step.CreatorID},
		Notification: FCMNotification,
	}

	notifyHandler.CreateNotification(notification, FCMOptions, FCMData)
}
