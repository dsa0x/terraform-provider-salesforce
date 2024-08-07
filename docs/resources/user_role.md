---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "salesforce_user_role Resource - terraform-provider-salesforce"
subcategory: ""
description: |-
  User Role Resource for the Salesforce Provider
---

# salesforce_user_role (Resource)

User Role Resource for the Salesforce Provider

## Example Usage

```terraform
resource "salesforce_user_role" "parent" {
  name           = "parent"
  developer_name = "parent"
}

resource "salesforce_user_role" "child" {
  name           = "child"
  developer_name = "child"
  parent_role_id = resource.salesforce_user_role.parent.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **developer_name** (String) The unique name of the object in the API. This name can contain only underscores and alphanumeric characters, and must be unique in your org. It must begin with a letter, not include spaces, not end with an underscore, and not contain two consecutive underscores. In managed packages, this field prevents naming conflicts on package installations. With this field, a developer can change the object’s name in a managed package and the changes are reflected in a subscriber’s organization. Corresponds to Role Name in the user interface.
- **name** (String) Name of the role. Corresponds to Label on the user interface.

### Optional

- **parent_role_id** (String) The ID of the parent role.

### Read-Only

- **id** (String) ID of the resource.

## Import

Import is supported using the following syntax:

```shell
terraform import salesforce_user_role.example 00AB0000000abc1AAA
```
