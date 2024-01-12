package hubmodels

import (
	"gorm.io/gorm"
	"time"
)

type Partner struct {
	ID                       string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CompanyID                string         `gorm:"column:company_id" json:"company_id"`
	PartnerTypeID            *int           `gorm:"column:partner_type_id" json:"partner_type_id"`
	PartnerCategory          string         `gorm:"column:partner_category" json:"partner_category"`
	PartnerCompanyName       string         `gorm:"column:partner_company_name" json:"partner_company_name"`
	PartnerCompanyRegister   *int           `gorm:"column:partner_company_register" json:"partner_company_register"`
	PartnerCivilSurname      *string        `gorm:"column:partner_civil_surname" json:"partner_civil_surname"`
	PartnerCivilName         *string        `gorm:"column:partner_civil_name" json:"partner_civil_name"`
	PartnerCivilRegister     *string        `gorm:"column:partner_civil_register" json:"partner_civil_register"`
	PartnerVatCheck          *int           `gorm:"column:partner_vat_check" json:"partner_vat_check"`
	PartnerProvince          *int           `gorm:"column:partner_province" json:"partner_province"`
	PartnerSum               *int           `gorm:"column:partner_sum" json:"partner_sum"`
	PartnerRegion            *int           `gorm:"column:partner_region" json:"partner_region"`
	PartnerAddress           *string        `gorm:"column:partner_address" json:"partner_address"`
	PartnerContractCheck     *int           `gorm:"column:partner_contract_check" json:"partner_contract_check"`
	PartnerContractStartDate *time.Time     `gorm:"column:partner_contract_start_date" json:"partner_contract_start_date"`
	PartnerContractEndDate   *time.Time     `gorm:"column:partner_contract_end_date" json:"partner_contract_end_date"`
	CreatedAt                *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	PartnerContractTypeID    *string        `gorm:"column:partner_contract_type_id" json:"partner_contract_type_id"`
	PartnerCompanyID         *string        `gorm:"column:partner_company_id" json:"partner_company_id"`
}

func (p *Partner) TableName() string {
	return "partner"
}

type SubPartnerPosition struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PartnerID *string `gorm:"column:partner_id" json:"partner_id"`
	Position  *string `gorm:"column:position" json:"position"`
	Name      *string `gorm:"column:name" json:"name"`
	Phone     *int    `gorm:"column:phone" json:"phone"`
}

func (s *SubPartnerPosition) TableName() string {
	return "sub_partner_position"
}
