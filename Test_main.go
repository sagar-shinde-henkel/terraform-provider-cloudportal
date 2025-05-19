package main

import (
	"encoding/json"
	"fmt"
)

// Define the struct based on your schema
type TicketInventory struct {
	TicketloudCosts                     float64 `json:"ticketloudCosts"`
	TicketOneTimeCostsUtilized          float64 `json:"ticketOneTimeCostsUtilized"`
	TicketOneTimeCostsUtilizedThisMonth float64 `json:"ticketOneTimeCostsUtilizedThisMonth"`
	TicketOneTimeCost                   float64 `json:"ticketOneTimeCost"`
	TicketClarityCodeOTCCode            string  `json:"ticketClarityCodeOTCCode"`
	TicketClarityCodeCostCenter         string  `json:"ticketClarityCodeCostCenter"`
	TicketClarityCodeCode               string  `json:"ticketClarityCodeCode"`
	TicketClarityCodeOTCCostCenter      string  `json:"ticketClarityCodeOTCCostCenter"`
	TicketCostCenter                    string  `json:"ticketCostCenter"`
	TicketCostCenterOTC                 string  `json:"ticketCostCenterOTC"`
	TicketChangedAt                     string  `json:"ticketChangedAt"`
	TicketChangedBy                     string  `json:"ticketChangedBy"`
	TicketServiceProviderValue          string  `json:"ticketServiceProviderValue"`
	TicketRequester                     string  `json:"ticketRequester"`
	TicketServiceProvider               string  `json:"ticketServiceProvider"`
	TicketCloudPlatform                 string  `json:"ticketCloudPlatform"`
	TicketArchitect                     string  `json:"ticketArchitect"`
	TicketNo                            int     `json:"ticketNo"`
	TicketTitle                         string  `json:"ticketTitle"`
	TicketAppOwner                      string  `json:"ticketAppOwner"`
	TicketAppName                       *string `json:"ticketAppName"`
	TicketAppID                         *string `json:"ticketAppID"`
	TicketApplicationManager            *string `json:"ticketApplicationManager"`
	TicketCyberRiskCategory             *string `json:"ticketCyberRiskCategory"`
	ResourceID                          string  `json:"resourceID"`
	ResourceStatus                      string  `json:"resourceStatus"`
	ResourceUpdated                     string  `json:"resourceUpdated"`
	TicketRelatedUnit                   *string `json:"ticketRelatedUnit"`
	TicketResponsibleGroupEmail         *string `json:"ticketResponsibleGroupEmail"`
	TicketCloudCosts                    float64 `json:"ticketCloudCosts"`
	ResourceAppID                       string  `json:"resourceAppID"`
	ResourceAppName                     string  `json:"resourceAppName"`
	ResourceAppOwner                    string  `json:"resourceAppOwner"`
	ResourceRelatedUnit                 string  `json:"resourceRelatedUnit"`
	ResourceName                        string  `json:"resourceName"`
	ResourceType                        string  `json:"resourceType"`
	ResourceContainer                   string  `json:"resourceContainer"`
	SnowGroup                           string  `json:"snowGroup"`
	CloudTemplateMaster                 string  `json:"cloudTemplateMaster"`
	TicketStatus                        string  `json:"ticketStatus"`
	CataLogresourceID                   string  `json:"cataLogresourceId"`
	ResourceContractName                *string `json:"resourceContractName"`
	CatalogApplicationSource            string  `json:"catalogApplicationSource"`
}

// Wrap in top-level response struct
type InventoryResponse struct {
	InventoryTickets []TicketInventory `json:"inventoryTickets"`
}

func _main() {
	// Sample JSON as a string (replace with your actual data or file read)
	jsonStr := `{"inventoryTickets":[{"ticketloudCosts":0.0,"ticketOneTimeCostsUtilized":0.0,"ticketOneTimeCostsUtilizedThisMonth":0.0,"ticketOneTimeCost":0.0,"ticketClarityCodeOTCCode":"","ticketClarityCodeCostCenter":"","ticketClarityCodeCode":"","ticketClarityCodeOTCCostCenter":"","ticketCostCenter":"","ticketCostCenterOTC":"","ticketChangedAt":"9/28/2022 1:20:47 PM +02:00","ticketChangedBy":"simon.vansti","ticketServiceProviderValue":"HybridCloud","ticketRequester":"alix.vondoernberg@hk.com","ticketServiceProvider":"","ticketCloudPlatform":"GCP","ticketArchitect":"simon.vanstiphout@hk.com","ticketNo":429,"ticketTitle":"Google Maps API","ticketAppOwner":"alix.vondoernberg@hk.com","ticketAppName":null,"ticketAppID":null,"ticketApplicationManager":null,"ticketCyberRiskCategory":null,"resourceID":"Updated by Provider","resourceStatus":"new","resourceUpdated":"","ticketRelatedUnit":null,"ticketResponsibleGroupEmail":null,"ticketCloudCosts":0.0,"resourceAppID":"string","resourceAppName":"string","resourceAppOwner":"string","resourceRelatedUnit":"new","resourceName":"adhesives fr maps api","resourceType":"GCP Project","resourceContainer":"","snowGroup":"","cloudTemplateMaster":"","ticketStatus":"demandAtRequester","cataLogresourceId":"4c7d223","resourceContractName":null,"catalogApplicationSource":"eb"},{"ticketloudCosts":0.0,"ticketOneTimeCostsUtilized":16500.0,"ticketOneTimeCostsUtilizedThisMonth":10278.0,"ticketOneTimeCost":15000.0,"ticketClarityCodeOTCCode":"","ticketClarityCodeCostCenter":"67001","ticketClarityCodeCode":"P081","ticketClarityCodeOTCCostCenter":"","ticketCostCenter":"","ticketCostCenterOTC":"","ticketChangedAt":"7/16/202 10:21:52 PM +02:00","ticketChangedBy":"cloud@hk.com","ticketServiceProviderValue":"HybridCloud","ticketRequester":"jens.brauer@hk.com","ticketServiceProvider":"","ticketCloudPlatform":"GCP","ticketArchitect":"simon.vanstiphout@hk.com","ticketNo":444,"ticketTitle":"InPlant MES Prod environement","ticketAppOwner":"Jens.Brauer@hk.com","ticketAppName":"Inplant MES production system","ticketAppID":"PR683","ticketApplicationManager":null,"ticketCyberRiskCategory":null,"resourceID":"string","resourceStatus":"new","resourceUpdated":"","ticketRelatedUnit":"dxA","ticketResponsibleGroupEmail":"Jens.Brauer@hk.com","ticketCloudCosts":5000.0,"resourceAppID":"string","resourceAppName":"string","resourceAppOwner":"string","resourceRelatedUnit":"new","resourceName":"der-a-de-prj-trt","resourceType":"GCP Project","resourceContainer":"","snowGroup":"","cloudTemplateMaster":"","ticketStatus":"underMaintenance","cataLogresourceId":"3fa8566afa6","resourceContractName":null,"catalogApplicationSource":"xl"}]}`
	var response InventoryResponse

	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		panic(err)
	}

	// Print a summary
	for _, ticket := range response.InventoryTickets {
		fmt.Printf("Ticket #%d: %s (Platform: %s)\n", ticket.TicketNo, ticket.TicketTitle, ticket.TicketCloudPlatform)
	}
}
