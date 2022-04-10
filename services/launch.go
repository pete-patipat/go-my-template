package services

import (
	"strings"

	"git.wndv.co/sharp/app/models"
)

func FormatEmail(e models.Email, c models.Campaign) string {
	r := strings.NewReplacer("{{firstname}}", e.Firstname, "{{lastname}}", e.Lastname)
	return r.Replace(c.BodyTemplate)
}
