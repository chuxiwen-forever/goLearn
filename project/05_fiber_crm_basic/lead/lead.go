package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string
}

func NewLead(*fiber.Ctx) {

}

func DeleteLead(*fiber.Ctx) {

}

func GetLeads(*fiber.Ctx) {

}

func GetLead(*fiber.Ctx) {

}
