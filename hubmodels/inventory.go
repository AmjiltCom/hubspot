package hubmodels

import (
	"gorm.io/gorm"
	"time"
)

type HubspotInventories struct {
	ID              string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HubConfigSpotID string    `gorm:"column:hub_config_spot_id" json:"hub_config_spot_id"`
	CompanyID       string    `gorm:"column:company_id" json:"company_id"`
	InventoryID     string    `gorm:"column:inventory_id" json:"inventory_id"`
	HubSpotID       string    `gorm:"column:hub_spot_id" json:"hub_spot_id"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HubspotInventories) TableName() string {
	return "hubspot_inventories"
}

type Inventory struct {
	ID                   string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CompanyID            string         `gorm:"column:company_id" json:"company_id"`
	UserID               string         `gorm:"column:user_id" json:"user_id"`
	Code                 int            `gorm:"column:code" json:"code"`
	InventoryName        string         `gorm:"column:inventory_name" json:"inventory_name"`
	InventoryDescription *string        `gorm:"column:inventory_description" json:"inventory_description"`
	MeasureUnitID        int            `gorm:"column:measure_unit_id" json:"measure_unit_id"`
	CategoryID           string         `gorm:"column:category_id" json:"category_id"`
	CategorySubID        string         `gorm:"column:category_sub_id" json:"category_sub_id"`
	CategorySubSubID     string         `gorm:"column:category_sub_sub_id" json:"category_sub_sub_id"`
	TagID                *string        `gorm:"column:tag_id" json:"tag_id"`
	PhotoFront           *string        `gorm:"column:photo_front" json:"photo_front"`
	PhotoOther           *string        `gorm:"column:photo_other" json:"photo_other"`
	Price                float32        `gorm:"column:price" json:"price"`
	InventoryCost        float32        `gorm:"column:inventory_cost" json:"inventory_cost"`
	WholeSaleQty         int            `gorm:"column:whole_sale_qty" json:"whole_sale_qty"`
	WholeSalePrice       float32        `gorm:"column:whole_sale_price" json:"whole_sale_price"`
	VatType              string         `gorm:"column:vat_type" json:"vat_type"`
	HasCityTax           bool           `gorm:"column:has_city_tax" json:"has_city_tax"`
	Balance              float32        `gorm:"column:balance" json:"balance"`
	AverageCost          float32        `gorm:"column:average_cost" json:"average_cost"`
	TotalCost            float32        `gorm:"column:total_cost" json:"total_cost"`
	CreatedAt            *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	StructID             string         `gorm:"column:struct_id" json:"struct_id"`
}

func (i *Inventory) TableName() string {
	return "inventory"
}

type ViewInventory struct {
	ID                   string         `gorm:"column:id" json:"id"`
	UserID               *string        `gorm:"column:user_id" json:"user_id"`
	Code                 int            `gorm:"column:code" json:"code"`
	InventoryName        string         `gorm:"column:inventory_name" json:"inventory_name"`
	InventoryDescription string         `gorm:"column:inventory_description" json:"inventory_description"`
	MeasureUnitID        *int           `gorm:"column:measure_unit_id" json:"measure_unit_id"`
	CategoryID           *string        `gorm:"column:category_id" json:"category_id"`
	CategorySubID        *string        `gorm:"column:category_sub_id" json:"category_sub_id"`
	CategorySubSubID     *string        `gorm:"column:category_sub_sub_id" json:"category_sub_sub_id"`
	TagID                *string        `gorm:"column:tag_id" json:"tag_id"`
	PhotoFront           string         `gorm:"column:photo_front" json:"photo_front"`
	PhotoOther           *string        `gorm:"column:photo_other" json:"photo_other"`
	CreatedAt            *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Price                float32        `gorm:"column:price" json:"price"`
	InventoryCost        float32        `gorm:"column:inventory_cost" json:"inventory_cost"`
	WholeSaleQty         *int           `gorm:"column:whole_sale_qty" json:"whole_sale_qty"`
	WholeSalePrice       *float32       `gorm:"column:whole_sale_price" json:"whole_sale_price"`
	VatType              *string        `gorm:"column:vat_type" json:"vat_type"`
	HasCityTax           bool           `gorm:"column:has_city_tax" json:"has_city_tax"`
	Balance              *float32       `gorm:"column:balance" json:"balance"`
	TotalCost            *float32       `gorm:"column:total_cost" json:"total_cost"`
	CompanyID            *string        `gorm:"column:company_id" json:"company_id"`
	AverageCost          *float32       `gorm:"column:average_cost" json:"average_cost"`
	StructID             *string        `gorm:"column:struct_id" json:"struct_id"`
	CategoryName         string         `gorm:"column:category_name" json:"category_name"`
	SubCategoryName      string         `gorm:"column:sub_category_name" json:"sub_category_name"`
	SubSubCategoryName   string         `gorm:"column:sub_sub_category_name" json:"sub_sub_category_name"`
	MeasureUnit          *string        `gorm:"column:measure_unit" json:"measure_unit"`
	TagName              *string        `gorm:"column:tag_name" json:"tag_name"`
	VatTypeTitle         *string        `gorm:"column:vat_type_title" json:"vat_type_title"`
	CityTax              *string        `gorm:"column:city_tax" json:"city_tax"`
	TotalAmount          *float32       `gorm:"column:total_amount" json:"total_amount"`
}

func (v *ViewInventory) TableName() string {
	return "view_inventory"
}
