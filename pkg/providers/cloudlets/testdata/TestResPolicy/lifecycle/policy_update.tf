provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

resource "akamai_cloudlets_policy" "policy" {
  name          = "test_policy_updated"
  cloudlet_code = "ER"
  description   = "test policy description"
  group_id      = "grp_123"
  match_rules   = <<-EOF
  [
  {
    "name": "r1",
    "type": "erMatchRule",
    "useRelativeUrl": "copy_scheme_hostname",
    "statusCode": 301,
    "redirectURL": "/ddd",
    "matchURL": "abc.com",
    "useIncomingQueryString": false,
    "useIncomingSchemeAndHost": true
  }
]
EOF
}