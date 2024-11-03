package controller

import (
	"team-builder/models"
)

func CreateVisitor(v models.Visitor) {
	models.Db.Create(&v)
}

func UpsertVisitor(v models.Visitor) {
	var visitor models.Visitor
	res := models.Db.Where("ip = ?", v.IP).First(&visitor)
	if res.RowsAffected > 0 {
		visitor.Count++
		models.Db.Save(&visitor)
	} else {
		models.Db.Create(&v)
	}
}

func GetVisitors() []models.Visitor {
	var visitors []models.Visitor
	models.Db.Find(&visitors)
	return visitors
}

func DeleteVisitorByIP(ip string) {
	var visitor models.Visitor
	res := models.Db.Where("ip = ?", ip).First(&visitor)
	if res.RowsAffected > 0 {
		models.Db.Delete(&visitor)
	}
}

func DeleteVisitorByID(id uint) {
	var visitor models.Visitor
	res := models.Db.Where("id = ?", id).First(&visitor)
	if res.RowsAffected > 0 {
		models.Db.Delete(&visitor)
	}
}
