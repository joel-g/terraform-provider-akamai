provider "akamai" {
  edgerc = "../../common/testutils/edgerc"
}

data "akamai_cloudlets_edge_redirector_match_rule" "test" {
  match_rules {
    name                      = "rule"
    start                     = 10
    end                       = 10000
    redirect_url              = "/abc/sss"
    status_code               = 307
    use_incoming_query_string = false
    use_relative_url          = "copy_scheme_hostname"
    matches_always            = true
  }
}