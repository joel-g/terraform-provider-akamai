provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

data "akamai_cps_csr" "test" {
  enrollment_id = 1
}