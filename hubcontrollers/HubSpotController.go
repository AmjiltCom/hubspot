package hubcontrollers

import (
	"bytes"
	"encoding/json"
	"fmt"
	hubSpotProto "github.com/AmjiltCom/hubspot/grpc/hubSpotProto"
	"github.com/AmjiltCom/hubspot/grpc/hubclient"
	"github.com/AmjiltCom/hubspot/hubmodels"
	"github.com/lambda-platform/lambda/DB"
	"io/ioutil"

	"net/http"
	"strconv"
)

func SyncPartners(companyID, partnerID string) {

	config := GetHubSpotConfig(companyID)

	if config != nil {

		if config.SyncPartner {

			partner := hubmodels.Partner{}
			var partnerContacts []hubmodels.SubPartnerPosition

			DB.DB.Where("id = ?", partnerID).Find(&partner)

			if partner.ID != "" {
				DB.DB.Where("partner_id = ?", partnerID).Find(&partnerContacts)

				contactProperties := hubmodels.ContactProperties{}

				partnerContactName := ""
				partnerContactPhone := ""
				partnerContactPosition := ""

				if len(partnerContacts) >= 1 {
					if partnerContacts[0].Name != nil {
						partnerContactName = *partnerContacts[0].Name
					}
					if partnerContacts[0].Position != nil {
						partnerContactPosition = *partnerContacts[0].Position
					}
					if partnerContacts[0].Phone != nil {
						prePhone := strconv.Itoa(*partnerContacts[0].Phone)

						if len(prePhone) <= 8 {
							prePhone = "+976" + prePhone
						}
						partnerContactPhone = prePhone
					}

				}

				if partner.PartnerCategory == "civil" {
					contactProperties.FirstName = partner.PartnerCompanyName
					if partner.PartnerCivilSurname != nil {
						contactProperties.LastName = *partner.PartnerCivilSurname
					}
					contactProperties.JobTitle = partnerContactPosition

				} else if partner.PartnerCategory == "company" {
					contactProperties.FirstName = partnerContactName
					contactProperties.LastName = partner.PartnerCompanyName
					contactProperties.JobTitle = partnerContactPosition
				}
				contactProperties.Phone = partnerContactPhone

				contact := hubmodels.Contact{
					Properties: contactProperties,
				}

				syncContactToHubSpot(partnerID, contact, *config)

			}
		}
	}

}

func GetHubSpotConfig(companyID string) *hubmodels.HubspotIntegration {
	var hubspotIntegration hubmodels.HubspotIntegration

	DB.DB.Where("company_id = ?", companyID).Find(&hubspotIntegration)

	if hubspotIntegration.ID != "" {
		return &hubspotIntegration
	} else {
		return nil
	}
}

func syncContactToHubSpot(partnerID string, contact hubmodels.Contact, config hubmodels.HubspotIntegration) {
	url := "https://api.hubspot.com/crm/v3/objects/contacts"
	method := "POST"

	savedContact := hubmodels.HubspotContacts{}
	DB.DB.Where("partner_id = ?", partnerID).Find(&savedContact)

	if savedContact.ID != "" {
		url = url + "/" + savedContact.HubSpotID
		method = "PATCH"
	}

	requestBody, _ := json.Marshal(contact)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.AccessToken)

		client := &http.Client{}
		resp, err2 := client.Do(req)
		if err2 == nil {
			defer resp.Body.Close()

			body, err3 := ioutil.ReadAll(resp.Body)
			if err3 == nil {

				contactResponse := hubmodels.ContactResponse{}

				err4 := json.Unmarshal(body, &contactResponse)

				if err4 == nil {
					if contactResponse.ID != "" {

						if savedContact.ID == "" {
							newContact := hubmodels.HubspotContacts{
								CompanyID:       config.CompanyID,
								PartnerID:       partnerID,
								HubConfigSpotID: config.ID,
								HubSpotID:       contactResponse.ID,
							}

							DB.DB.Create(&newContact)

						}

					}
				}

			}

		}

	}

}

func SyncInventory(companyID, inventoryID, GRPCURL string) {
	request := hubSpotProto.ConfigRequest{
		CompanyID: companyID,
	}
	config, err := hubclient.GegHubConfig(&request, GRPCURL)

	if err == nil && config.Config.AccessToken != "" && config.Config.SyncPartner {

		inventory := hubmodels.ViewInventory{}
		DB.DB.Where("id = ?", inventoryID).Find(&inventory)

		if inventory.ID != "" {

			productProperties := hubmodels.ProductProperties{}

			productProperties.Name = inventory.InventoryName
			productProperties.Description = fmt.Sprintf("%s > %s > %s. %s", inventory.CategoryName, inventory.SubCategoryName, inventory.SubSubCategoryName, inventory.InventoryDescription)
			productProperties.Price = fmt.Sprintf("%f", inventory.Price)
			productProperties.HsCostOfGoodsSold = fmt.Sprintf("%f", inventory.InventoryCost)
			productProperties.HsSku = fmt.Sprintf("%d", inventory.Code)
			productProperties.HsProductType = "inventory"
			if inventory.PhotoFront != "" {
				productProperties.HsImages = fmt.Sprintf("https://api.tatatonga.com/main%s", inventory.PhotoFront)
			}

			product := hubmodels.Product{
				Properties: productProperties,
			}

			syncProductToHubSpot(inventoryID, product, *config)

		}
	}

}
func syncProductToHubSpot(inventoryID string, product hubmodels.Product, config hubSpotProto.ConfigResponse) {
	url := "https://api.hubspot.com/crm/v3/objects/products"
	method := "POST"

	savedInventory := hubmodels.HubspotInventories{}
	DB.DB.Where("inventory_id = ?", inventoryID).Find(&savedInventory)

	if savedInventory.ID != "" {
		url = url + "/" + savedInventory.HubSpotID
		method = "PATCH"
	}

	requestBody, _ := json.Marshal(product)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.Config.AccessToken)

		client := &http.Client{}
		resp, err2 := client.Do(req)
		if err2 == nil {
			defer resp.Body.Close()

			body, err3 := ioutil.ReadAll(resp.Body)
			if err3 == nil {

				productResponse := hubmodels.ProductResponse{}

				err4 := json.Unmarshal(body, &productResponse)

				if err4 == nil {
					if productResponse.ID != "" {

						if savedInventory.ID == "" {
							newContact := hubmodels.HubspotInventories{
								CompanyID:       config.Config.CompanyID,
								InventoryID:     inventoryID,
								HubConfigSpotID: config.Config.ID,
								HubSpotID:       productResponse.ID,
							}

							DB.DB.Create(&newContact)

						}

					}
				}

			}

		}

	}

}
