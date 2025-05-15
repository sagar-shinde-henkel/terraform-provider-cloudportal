data "cloudportal_datasource_ticket_search" "search" {
  keyword= "gcp"
}

output "search-result"{
  value=data.cloudportal_datasource_ticket_search.search
}