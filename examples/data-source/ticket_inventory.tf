data "cloudportal_datasource_ticket_inventory" "inventory" {
  
}

# Output the ticket details
output "ticket_inventory" {
  value = data.cloudportal_datasource_ticket_inventory.exampletwo
}