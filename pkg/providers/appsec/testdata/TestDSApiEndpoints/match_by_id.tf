provider "akamai" {
  edgerc = "~/.edgerc"
}




data "akamai_appsec_api_endpoints" "test" {
  config_id = 43253
    version = 7

 // name = var.api_endpoint_name
}