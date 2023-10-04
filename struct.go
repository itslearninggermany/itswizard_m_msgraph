package itswizard_m_msgraph

import "github.com/jinzhu/gorm"

/*
Here are the licence in Office 365
*/
type DbMsGraphSKUReference struct {
	gorm.Model
	SkuID string `gorm:"unique"`
	Name  string
}

type DbAzureSetup struct {
	gorm.Model
	OrganisationID uint
	InstitutionID  uint
	Profile        string
	SkuID          string
	Domain         string
	AuthGroupID    string
}
