provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

resource "akamai_property" "test" {
  name        = "test_property"
  contract_id = "ctr_123"
  group_id    = "grp_123"
  product_id  = "prd_123"

  rule_format = "v2023-01-05"
  rules       = file("testdata/TestResProperty/Lifecycle/versionNotes/04_05_rules.json")
}