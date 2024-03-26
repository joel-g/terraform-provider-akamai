provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

data "akamai_edgeworkers_resource_tier" "test" {
  contract_id        = "ctr_1-599K"
  resource_tier_name = "Basic Compute"
}