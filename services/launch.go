package services

import (
	"strings"

	"git.wndv.co/sharp/app/models"
)

func FormatEmail(email models.Email, campaign models.Campaign) string {
	replacer := strings.NewReplacer("{{firstname}}", email.Firstname, "{{lastname}}", email.Lastname)
	return replacer.Replace(campaign.BodyTemplate)
}
