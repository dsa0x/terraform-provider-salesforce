---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "salesforce_user Resource - terraform-provider-salesforce"
subcategory: ""
description: |-
  
---

# salesforce_user (Resource)



## Example Usage

```terraform
data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

resource "salesforce_user" "example" {
  alias               = "example"
  email               = "user@example.com"
  last_name           = "example"
  username            = "user@example.com"
  profile_id          = data.salesforce_profile.chatter_free.id
  email_encoding_key  = "UTF-8"
  language_locale_key = "en_US"
  time_zone_sid_key   = "America/Chicago"
  locale_sid_key      = "en_US"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **alias** (String)
- **email** (String)
- **last_name** (String)
- **profile_id** (String)
- **username** (String)

### Optional

- **email_encoding_key** (String)
- **language_locale_key** (String)
- **locale_sid_key** (String)
- **time_zone_sid_key** (String)

### Read-Only

- **id** (String) The ID of this resource.

