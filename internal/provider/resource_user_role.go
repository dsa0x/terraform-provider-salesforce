package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
)

type userRoleType struct {
}

func (userRoleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"developer_name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				}, // TODO full validation, see requirements in https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_role.htm
			},
			"parent_role_id": {
				Type:     types.StringType,
				Optional: true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
				},
			},
		},
	}, nil
}

func (u userRoleType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(u)}
	}
	r := &userRoleResource{}
	r.Client = provider.client
	r.SObject = r
	return r, nil
}

type userRoleResource struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	DeveloperName string       `tfsdk:"developer_name" force:",omitempty"`
	ParentRoleId  *string      `tfsdk:"parent_role_id" force:",omitempty"`
	Id            types.String `tfsdk:"id" force:"-"`
	Resource      `tfsdk:"-"`
}

func (userRoleResource) ApiName() string {
	return "UserRole"
}

func (userRoleResource) ExternalIdApiName() string {
	return ""
}

func (u *userRoleResource) Instance() force.SObject {
	return u
}

func (u *userRoleResource) Updatable() force.SObject {
	return *u
}

func (u *userRoleResource) GetId() string {
	return u.Id.Value
}

func (u *userRoleResource) SetId(id string) {
	u.Id = types.String{Value: id}
}
