package lead

import (
	"github.com/Sourjaya/go-crm/database"
	l "github.com/Sourjaya/go-crm/logger"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	ctx.JSON(leads)
}
func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	ctx.JSON(lead)
}
func NewLead(ctx *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := ctx.BodyParser(lead); err != nil {
		l.Error("Body could not be parsed")
		ctx.Status(503).Send(err)
		return
	}

	var checkData Validator
	checkData = Email{lead.Email}
	if _, err := checkData.isValid(); err != nil {
		l.Error("Email invalid")
		ctx.Status(503).Send(err)
		return
	}
	checkData = Phone{lead.Phone}
	if _, err := checkData.isValid(); err != nil {
		l.Error("Phone invalid")
		ctx.Status(400).Send(err)
		return
	}
	db.Create(&lead)
	ctx.JSON(lead)
}
func DeleteLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		l.Error("No Lead found")
		ctx.Status(503).Send("No Lead found with ID")
		return
	}
	db.Delete(&lead)
	l.Info("Lead deleted")
	ctx.Send("Lead succesfully deleted")
}
