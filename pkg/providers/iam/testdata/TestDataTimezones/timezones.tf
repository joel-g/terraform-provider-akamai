provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

data "akamai_iam_timezones" "test" {}