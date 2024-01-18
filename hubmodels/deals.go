package hubmodels

import "time"

type HubspotDeals struct {
	ID              string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HubConfigSpotID string    `gorm:"column:hub_config_spot_id" json:"hub_config_spot_id"`
	CompanyID       string    `gorm:"column:company_id" json:"company_id"`
	IssueID         string    `gorm:"column:issue_id" json:"issue_id"`
	HubSpotID       string    `gorm:"column:hub_spot_id" json:"hub_spot_id"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HubspotDeals) TableName() string {
	return "hubspot_deals"
}

type HubspotDealsLines struct {
	ID              string     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HubConfigSpotID string     `gorm:"column:hub_config_spot_id" json:"hub_config_spot_id"`
	CompanyID       string     `gorm:"column:company_id" json:"company_id"`
	IssueID         string     `gorm:"column:issue_id" json:"issue_id"`
	HubSpotID       string     `gorm:"column:hub_spot_id" json:"hub_spot_id"`
	HubSlotLineID   string     `gorm:"column:hub_slot_line_id" json:"hub_slot_line_id"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HubspotDealsLines) TableName() string {
	return "hubspot_deals_lines"
}

type Line struct {
	Properties   LineProperties    `json:"properties"`
	Associations []LineAssociation `json:"associations"`
}
type LineAssociation struct {
	To    LineTo      `json:"to"`
	Types []LineTypes `json:"types"`
}
type LineTo struct {
	ID int64 `json:"id"`
}
type LineTypes struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int    `json:"associationTypeId"`
}
type LineProperties struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

type Deal struct {
	Properties   DealProperties `json:"properties"`
	Associations []DealContact  `json:"associations"`
}

type DealContact struct {
	To    DealContactID `json:"to"`
	Types []DealTypes   `json:"types"`
}
type DealContactID struct {
	ID int `json:"id"`
}
type DealTypes struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int    `json:"associationTypeId"`
}

type DealProperties struct {
	Dealname  string    `json:"dealname"`
	Pipeline  string    `json:"pipeline"`
	Dealstage string    `json:"dealstage"`
	Amount    string    `json:"amount"`
	Closedate time.Time `json:"closedate"`
	Dealtype  string    `json:"dealtype"`
}

type DealResponse struct {
	ID         string `json:"id"`
	Properties struct {
		Amount                          string    `json:"amount"`
		AmountInHomeCurrency            string    `json:"amount_in_home_currency"`
		Closedate                       time.Time `json:"closedate"`
		Createdate                      time.Time `json:"createdate"`
		DaysToClose                     string    `json:"days_to_close"`
		Dealname                        string    `json:"dealname"`
		Dealstage                       string    `json:"dealstage"`
		Dealtype                        string    `json:"dealtype"`
		HsClosedAmount                  string    `json:"hs_closed_amount"`
		HsClosedAmountInHomeCurrency    string    `json:"hs_closed_amount_in_home_currency"`
		HsCreatedate                    time.Time `json:"hs_createdate"`
		HsDaysToCloseRaw                string    `json:"hs_days_to_close_raw"`
		HsDealStageProbabilityShadow    string    `json:"hs_deal_stage_probability_shadow"`
		HsForecastAmount                string    `json:"hs_forecast_amount"`
		HsIsActiveSharedDeal            string    `json:"hs_is_active_shared_deal"`
		HsIsClosed                      string    `json:"hs_is_closed"`
		HsIsClosedWon                   string    `json:"hs_is_closed_won"`
		HsIsDealSplit                   string    `json:"hs_is_deal_split"`
		HsIsOpenCount                   string    `json:"hs_is_open_count"`
		HsLastmodifieddate              time.Time `json:"hs_lastmodifieddate"`
		HsObjectID                      string    `json:"hs_object_id"`
		HsObjectSource                  string    `json:"hs_object_source"`
		HsObjectSourceID                string    `json:"hs_object_source_id"`
		HsObjectSourceLabel             string    `json:"hs_object_source_label"`
		HsProjectedAmount               string    `json:"hs_projected_amount"`
		HsProjectedAmountInHomeCurrency string    `json:"hs_projected_amount_in_home_currency"`
		Pipeline                        string    `json:"pipeline"`
	} `json:"properties"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Archived  bool      `json:"archived"`
}
type LineResponse struct {
	ID         string `json:"id"`
	Properties struct {
		Amount                             string    `json:"amount"`
		Createdate                         time.Time `json:"createdate"`
		HsAcv                              string    `json:"hs_acv"`
		HsArr                              string    `json:"hs_arr"`
		HsLastmodifieddate                 time.Time `json:"hs_lastmodifieddate"`
		HsMargin                           string    `json:"hs_margin"`
		HsMarginAcv                        string    `json:"hs_margin_acv"`
		HsMarginArr                        string    `json:"hs_margin_arr"`
		HsMarginMrr                        string    `json:"hs_margin_mrr"`
		HsMarginTcv                        string    `json:"hs_margin_tcv"`
		HsMrr                              string    `json:"hs_mrr"`
		HsObjectID                         string    `json:"hs_object_id"`
		HsObjectSource                     string    `json:"hs_object_source"`
		HsObjectSourceID                   string    `json:"hs_object_source_id"`
		HsObjectSourceLabel                string    `json:"hs_object_source_label"`
		HsPositionOnQuote                  string    `json:"hs_position_on_quote"`
		HsPostTaxAmount                    string    `json:"hs_post_tax_amount"`
		HsPreDiscountAmount                string    `json:"hs_pre_discount_amount"`
		HsRecurringBillingNumberOfPayments string    `json:"hs_recurring_billing_number_of_payments"`
		HsTcv                              string    `json:"hs_tcv"`
		HsTotalDiscount                    string    `json:"hs_total_discount"`
		Name                               string    `json:"name"`
		Price                              string    `json:"price"`
		Quantity                           string    `json:"quantity"`
	} `json:"properties"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Archived  bool      `json:"archived"`
}
