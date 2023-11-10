resource "newtest_zone" "my_zone" {
  account_id               = 1
  auto_recover_power_state = false
  code                     = "mycloud"
  config                   = {}
  credential = {
    id   = 5
    type = "...my_type..."
  }
  description       = "...my_description..."
  enabled           = true
  group_id          = 3
  linked_account_id = 9
  location          = "US East"
  name              = "My Cloud"
  scale_priority    = 10
  security_mode     = "...my_security_mode..."
  visibility        = "public"
  zone_type = {
    zone_create_1 = {
      id = 12
    }
  }
}