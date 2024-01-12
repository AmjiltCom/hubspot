package hubmodels

import "time"

type HubspotIntegration struct {
	ID           string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CompanyID    string    `gorm:"column:company_id" json:"company_id"`
	AccessToken  string    `gorm:"column:access_token" json:"access_token"`
	SyncPartner  bool      `gorm:"column:sync_partner" json:"sync_partner"`
	SyncProduct  bool      `gorm:"column:sync_product" json:"sync_product"`
	SyncPosOrder bool      `gorm:"column:sync_pos_order" json:"sync_pos_order"`
	SyncSales    bool      `gorm:"column:sync_sales" json:"sync_sales"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HubspotIntegration) TableName() string {
	return "hubspot_integration"
}

type HubspotContacts struct {
	ID              string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HubConfigSpotID string    `gorm:"column:hub_config_spot_id" json:"hub_config_spot_id"`
	CompanyID       string    `gorm:"column:company_id" json:"company_id"`
	PartnerID       string    `gorm:"column:partner_id" json:"partner_id"`
	HubSpotID       string    `gorm:"column:hub_spot_id" json:"hub_spot_id"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HubspotContacts) TableName() string {
	return "hubspot_contacts"
}

type ContactProperties struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	JobTitle  string `json:"jobtitle"`
}

type Contact struct {
	Properties ContactProperties `json:"properties"`
}

type ContactResponse struct {
	ID         string `json:"id"`
	Properties struct {
		Createdate                        time.Time `json:"createdate"`
		Email                             string    `json:"email"`
		Firstname                         string    `json:"firstname"`
		HsAllContactVids                  string    `json:"hs_all_contact_vids"`
		HsEmailDomain                     string    `json:"hs_email_domain"`
		HsIsContact                       string    `json:"hs_is_contact"`
		HsIsUnworked                      string    `json:"hs_is_unworked"`
		HsLifecyclestageLeadDate          time.Time `json:"hs_lifecyclestage_lead_date"`
		HsMarketableStatus                string    `json:"hs_marketable_status"`
		HsMarketableUntilRenewal          string    `json:"hs_marketable_until_renewal"`
		HsObjectID                        string    `json:"hs_object_id"`
		HsObjectSource                    string    `json:"hs_object_source"`
		HsObjectSourceID                  string    `json:"hs_object_source_id"`
		HsObjectSourceLabel               string    `json:"hs_object_source_label"`
		HsPipeline                        string    `json:"hs_pipeline"`
		HsSearchableCalculatedPhoneNumber string    `json:"hs_searchable_calculated_phone_number"`
		Lastmodifieddate                  time.Time `json:"lastmodifieddate"`
		Lastname                          string    `json:"lastname"`
		Lifecyclestage                    string    `json:"lifecyclestage"`
		Phone                             string    `json:"phone"`
	} `json:"properties"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Archived  bool      `json:"archived"`
}

type ProductResponse struct {
	ID         string `json:"id"`
	Properties struct {
		Createdate          time.Time `json:"createdate"`
		Description         string    `json:"description"`
		HsCostOfGoodsSold   string    `json:"hs_cost_of_goods_sold"`
		HsImages            string    `json:"hs_images"`
		HsLastmodifieddate  time.Time `json:"hs_lastmodifieddate"`
		HsObjectID          string    `json:"hs_object_id"`
		HsObjectSource      string    `json:"hs_object_source"`
		HsObjectSourceID    string    `json:"hs_object_source_id"`
		HsObjectSourceLabel string    `json:"hs_object_source_label"`
		HsProductType       string    `json:"hs_product_type"`
		HsSku               string    `json:"hs_sku"`
		Name                string    `json:"name"`
		Price               string    `json:"price"`
		Tax                 string    `json:"tax"`
	} `json:"properties"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Archived  bool      `json:"archived"`
}
type Product struct {
	Properties ProductProperties `json:"properties"`
}

type ProductProperties struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Price             string `json:"price"`
	HsCostOfGoodsSold string `json:"hs_cost_of_goods_sold"`
	HsSku             string `json:"hs_sku"`
	HsProductType     string `json:"hs_product_type"`
	HsImages          string `json:"hs_images"`
}
