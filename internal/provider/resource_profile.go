package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
)

type profileType struct {
}

func (profileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					staticComputed{},
				},
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"user_license_id": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
					tfsdk.RequiresReplace(),
				},
			},
			// permissions
			"permissions_email_single":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_mass":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_task":                          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_event":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_export_report":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_personal":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_data_export":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_users":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_filters":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_templates":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_all_data":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_cases":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_mass_inline_edit":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_knowledge":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_knowledge":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_solutions":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_customize_application":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_readonly_fields":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_run_reports":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_setup":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_entity":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_new_report_builder":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activate_contract":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activate_order":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_leads":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_leads":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_lead":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_data":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_documents":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_encrypted_data":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_brand_templates":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_html_templates":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_internal_user":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_encryption_keys":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delete_activated_contract":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_invite_external_users":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_sit_requests":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_remote_access":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_use_new_dashboard_builder":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_categories":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_convert_leads":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_password_never_expires":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_team_reassign_wizards":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_activated_orders":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_install_multiforce":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_publish_multiforce":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_own_groups":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_opp_line_item_unit_price":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_multiforce":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bulk_api_hard_delete":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_solution_import":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_call_centers":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_synonyms":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_content":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_email_client_config":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_enable_notifications":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_data_integrations":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_distribute_from_pers_wksp":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_categories":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_data_categories":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_author_apex":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_mobile":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_api_enabled":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_custom_report_types":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_case_comments":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_case":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_administrator":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_workspaces":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_permissions":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_properties":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_types":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_exchange_config":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_analytic_snapshots":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_schedule_reports":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_business_hour_holidays":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_entitlements":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_dynamic_dashboards":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_custom_sidebar_on_all_pages":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_interaction":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_my_teams_dashboards":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_moderate_chatter":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_reset_passwords":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_flow_ufl_required":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_insert_feed_system_fields":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activities_access":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_knowledge_import_export":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_template_management":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_administration":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_chatter_messages":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_email_ic":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_file_link":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_force_two_factor":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_event_log_files":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_networks":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_auth_providers":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_run_flow":                           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_dashboards":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_dashboard_folders":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_public_dashboards":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_dashbds_in_pub_folders":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_reports":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_report_folders":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_public_reports":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_reports_in_pub_folders":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_my_dashboards":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_my_reports":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_users":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_universal_search":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_connect_org_to_environment_hub":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_work_calibration_user":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_filters":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_work_dot_com_user_perm":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_hub_user":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_govern_networks":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sales_console":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_two_factor_api":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delete_topics":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_topics":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_topics":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_assign_topics":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_identity_enabled":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_identity_connect":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_view_knowledge":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_workspaces":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_work_badge_definition":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_search_promotion_rules":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_custom_mobile_apps_access":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_help_link":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_profiles_permissionsets":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_assign_permission_sets":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_roles":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_ip_addresses":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_sharing":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_internal_users":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_password_policies":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_login_access_policies":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_platform_events":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_custom_permissions":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_verify_comment":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_unlisted_groups":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_std_automatic_activity_capture":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_dashboard_editor":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_two_factor":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_user":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_admin":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_elt_editor":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_upload_user":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_create_application":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lightning_experience_user":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_leakage_events":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_config_custom_recs":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_submit_macros_allowed":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bulk_macros_allowed":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_share_internal_articles":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_session_permission_sets":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_templated_app":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_templated_app":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_announcement_emails":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_edit_own_post":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_edit_own_record_post":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sales_analytics_user":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_service_analytics_user":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_wave_tabular_download":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_automatic_activity_capture":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_custom_objects":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delegated_two_factor":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_compose_ui_codesnippet":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_select_files_from_salesforce":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_moderate_network_users":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_merge_topics":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_to_lightning_reports":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_pvt_rpts_and_dashbds":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_lightning_login":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_campaign_influence2":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_assessment":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_remove_direct_message_members":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_approve_feed_post":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_direct_message_members":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_view_edit_converted_leads":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_show_company_name_as_user_badge":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_access_cmc":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_health_check":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_health_check":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_packaging2":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_certificates":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_report_in_lightning":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_prevent_classic_experience":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_hide_read_by_list":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_smart_data_discovery":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_get_smart_data_discovery":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_update_sdd_dataset":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_update_sdd_story":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_smart_data_discovery":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_share_smart_data_discovery_story":   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_smart_data_discovery_model":  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_list_email_send":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_feed_pinning":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_change_dashboard_colors":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_iot_user":                           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_recommendation_strategies":   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_propositions":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_close_conversations":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_report_roles_grps":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_dashboard_roles_grps":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_web_link":                       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_has_unlimited_nba_executions":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_only_embedded_app_user":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_adoption_analytics_user":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_activities":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_report_to_other_users":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lightning_console_allowed_for_user": {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_reports_run_as_user":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_to_lightning_dashboards":  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_dashboard_to_other_users": {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_ltng_temp_in_pub":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transactional_email_send":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_private_static_resources":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_ltng_temp_folder":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_apex_rest_services":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_enable_community_app_launcher":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_give_recognition_badge":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_run_analysis":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_my_search":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ltng_promo_reserved01_user_perm":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_subscriptions":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_wave_manage_private_assets_user":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_edit_data_prep_recipe":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_analytics_remote_connections":   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_surveys":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_assistant_dialog":               {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_query_suggestions":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_record_visibility_api":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_roles":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_smart_data_discovery_for_community": {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_manage_maps":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_story_on_ds_with_predicate":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lm_outbound_messaging_user_perm":    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_data_classification":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_privacy_data_access":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_query_all_files":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_metadata":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_cms":                         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sandbox_testing_in_community_app":   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_edit_prompts":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_user_pii":                      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_hub_connections":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_b_2_b_marketing_analytics_user":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_trace_xds_queries":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_security_command_center":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_security_command_center":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_custom_settings":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_foreign_key_names":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_survey_advanced_features":     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_wave_notification_recipients":   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_headless_cms_access":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lm_end_messaging_session_user_perm": {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_consent_api_update":                 {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_payments_api_user":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_access_content_builder":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_account_switcher_user":              {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_anomaly_events":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_c_360_a_connections":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_release_updates":             {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_profiles":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_skip_identity_confirmation":         {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_learning_manager":                   {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_custom_notifications":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_packaging2_delete":                  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_autonomous_analytics_privacy":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sonic_consumer":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bot_manage_bots":                    {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bot_manage_bots_training_data":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_learning_reporting":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_c_to_c_user":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_access":                     {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_lex":                        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_get_smart_data_discovery_external":  {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_quip_metrics_access":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_quip_user_engagement_metrics":       {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_external_connections":        {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_subscription_emails":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ai_view_insight_objects":            {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ai_create_insight_objects":          {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lifecycle_management_api_user":      {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_native_webview_scrolling":           {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_developer_name":                {Type: BoolMarshalerType{}, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
		},
	}, nil
}

func (p profileType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(p)}
	}
	return &profileResource{
		Resource: Resource{
			Client:              provider.client,
			Data:                &profileResourceData{},
			NeedsGetAfterUpsert: true,
		},
	}, nil
}

type profileResource struct {
	Resource
}

type profileResourceData struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   *string      `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
	Id            types.String `tfsdk:"id" force:"-"`
	// permissions
	PermissionsEmailSingle                    BoolMarshaler `tfsdk:"permissions_email_single" force:",omitempty"`
	PermissionsEmailMass                      BoolMarshaler `tfsdk:"permissions_email_mass" force:",omitempty"`
	PermissionsEditTask                       BoolMarshaler `tfsdk:"permissions_edit_task" force:",omitempty"`
	PermissionsEditEvent                      BoolMarshaler `tfsdk:"permissions_edit_event" force:",omitempty"`
	PermissionsExportReport                   BoolMarshaler `tfsdk:"permissions_export_report" force:",omitempty"`
	PermissionsImportPersonal                 BoolMarshaler `tfsdk:"permissions_import_personal" force:",omitempty"`
	PermissionsDataExport                     BoolMarshaler `tfsdk:"permissions_data_export" force:",omitempty"`
	PermissionsManageUsers                    BoolMarshaler `tfsdk:"permissions_manage_users" force:",omitempty"`
	PermissionsEditPublicFilters              BoolMarshaler `tfsdk:"permissions_edit_public_filters" force:",omitempty"`
	PermissionsEditPublicTemplates            BoolMarshaler `tfsdk:"permissions_edit_public_templates" force:",omitempty"`
	PermissionsModifyAllData                  BoolMarshaler `tfsdk:"permissions_modify_all_data" force:",omitempty"`
	PermissionsManageCases                    BoolMarshaler `tfsdk:"permissions_manage_cases" force:",omitempty"`
	PermissionsMassInlineEdit                 BoolMarshaler `tfsdk:"permissions_mass_inline_edit" force:",omitempty"`
	PermissionsEditKnowledge                  BoolMarshaler `tfsdk:"permissions_edit_knowledge" force:",omitempty"`
	PermissionsManageKnowledge                BoolMarshaler `tfsdk:"permissions_manage_knowledge" force:",omitempty"`
	PermissionsManageSolutions                BoolMarshaler `tfsdk:"permissions_manage_solutions" force:",omitempty"`
	PermissionsCustomizeApplication           BoolMarshaler `tfsdk:"permissions_customize_application" force:",omitempty"`
	PermissionsEditReadonlyFields             BoolMarshaler `tfsdk:"permissions_edit_readonly_fields" force:",omitempty"`
	PermissionsRunReports                     BoolMarshaler `tfsdk:"permissions_run_reports" force:",omitempty"`
	PermissionsViewSetup                      BoolMarshaler `tfsdk:"permissions_view_setup" force:",omitempty"`
	PermissionsTransferAnyEntity              BoolMarshaler `tfsdk:"permissions_transfer_any_entity" force:",omitempty"`
	PermissionsNewReportBuilder               BoolMarshaler `tfsdk:"permissions_new_report_builder" force:",omitempty"`
	PermissionsActivateContract               BoolMarshaler `tfsdk:"permissions_activate_contract" force:",omitempty"`
	PermissionsActivateOrder                  BoolMarshaler `tfsdk:"permissions_activate_order" force:",omitempty"`
	PermissionsImportLeads                    BoolMarshaler `tfsdk:"permissions_import_leads" force:",omitempty"`
	PermissionsManageLeads                    BoolMarshaler `tfsdk:"permissions_manage_leads" force:",omitempty"`
	PermissionsTransferAnyLead                BoolMarshaler `tfsdk:"permissions_transfer_any_lead" force:",omitempty"`
	PermissionsViewAllData                    BoolMarshaler `tfsdk:"permissions_view_all_data" force:",omitempty"`
	PermissionsEditPublicDocuments            BoolMarshaler `tfsdk:"permissions_edit_public_documents" force:",omitempty"`
	PermissionsViewEncryptedData              BoolMarshaler `tfsdk:"permissions_view_encrypted_data" force:",omitempty"`
	PermissionsEditBrandTemplates             BoolMarshaler `tfsdk:"permissions_edit_brand_templates" force:",omitempty"`
	PermissionsEditHtmlTemplates              BoolMarshaler `tfsdk:"permissions_edit_html_templates" force:",omitempty"`
	PermissionsChatterInternalUser            BoolMarshaler `tfsdk:"permissions_chatter_internal_user" force:",omitempty"`
	PermissionsManageEncryptionKeys           BoolMarshaler `tfsdk:"permissions_manage_encryption_keys" force:",omitempty"`
	PermissionsDeleteActivatedContract        BoolMarshaler `tfsdk:"permissions_delete_activated_contract" force:",omitempty"`
	PermissionsChatterInviteExternalUsers     BoolMarshaler `tfsdk:"permissions_chatter_invite_external_users" force:",omitempty"`
	PermissionsSendSitRequests                BoolMarshaler `tfsdk:"permissions_send_sit_requests" force:",omitempty"`
	PermissionsManageRemoteAccess             BoolMarshaler `tfsdk:"permissions_manage_remote_access" force:",omitempty"`
	PermissionsCanUseNewDashboardBuilder      BoolMarshaler `tfsdk:"permissions_can_use_new_dashboard_builder" force:",omitempty"`
	PermissionsManageCategories               BoolMarshaler `tfsdk:"permissions_manage_categories" force:",omitempty"`
	PermissionsConvertLeads                   BoolMarshaler `tfsdk:"permissions_convert_leads" force:",omitempty"`
	PermissionsPasswordNeverExpires           BoolMarshaler `tfsdk:"permissions_password_never_expires" force:",omitempty"`
	PermissionsUseTeamReassignWizards         BoolMarshaler `tfsdk:"permissions_use_team_reassign_wizards" force:",omitempty"`
	PermissionsEditActivatedOrders            BoolMarshaler `tfsdk:"permissions_edit_activated_orders" force:",omitempty"`
	PermissionsInstallMultiforce              BoolMarshaler `tfsdk:"permissions_install_multiforce" force:",omitempty"`
	PermissionsPublishMultiforce              BoolMarshaler `tfsdk:"permissions_publish_multiforce" force:",omitempty"`
	PermissionsChatterOwnGroups               BoolMarshaler `tfsdk:"permissions_chatter_own_groups" force:",omitempty"`
	PermissionsEditOppLineItemUnitPrice       BoolMarshaler `tfsdk:"permissions_edit_opp_line_item_unit_price" force:",omitempty"`
	PermissionsCreateMultiforce               BoolMarshaler `tfsdk:"permissions_create_multiforce" force:",omitempty"`
	PermissionsBulkApiHardDelete              BoolMarshaler `tfsdk:"permissions_bulk_api_hard_delete" force:",omitempty"`
	PermissionsSolutionImport                 BoolMarshaler `tfsdk:"permissions_solution_import" force:",omitempty"`
	PermissionsManageCallCenters              BoolMarshaler `tfsdk:"permissions_manage_call_centers" force:",omitempty"`
	PermissionsManageSynonyms                 BoolMarshaler `tfsdk:"permissions_manage_synonyms" force:",omitempty"`
	PermissionsViewContent                    BoolMarshaler `tfsdk:"permissions_view_content" force:",omitempty"`
	PermissionsManageEmailClientConfig        BoolMarshaler `tfsdk:"permissions_manage_email_client_config" force:",omitempty"`
	PermissionsEnableNotifications            BoolMarshaler `tfsdk:"permissions_enable_notifications" force:",omitempty"`
	PermissionsManageDataIntegrations         BoolMarshaler `tfsdk:"permissions_manage_data_integrations" force:",omitempty"`
	PermissionsDistributeFromPersWksp         BoolMarshaler `tfsdk:"permissions_distribute_from_pers_wksp" force:",omitempty"`
	PermissionsViewDataCategories             BoolMarshaler `tfsdk:"permissions_view_data_categories" force:",omitempty"`
	PermissionsManageDataCategories           BoolMarshaler `tfsdk:"permissions_manage_data_categories" force:",omitempty"`
	PermissionsAuthorApex                     BoolMarshaler `tfsdk:"permissions_author_apex" force:",omitempty"`
	PermissionsManageMobile                   BoolMarshaler `tfsdk:"permissions_manage_mobile" force:",omitempty"`
	PermissionsApiEnabled                     BoolMarshaler `tfsdk:"permissions_api_enabled" force:",omitempty"`
	PermissionsManageCustomReportTypes        BoolMarshaler `tfsdk:"permissions_manage_custom_report_types" force:",omitempty"`
	PermissionsEditCaseComments               BoolMarshaler `tfsdk:"permissions_edit_case_comments" force:",omitempty"`
	PermissionsTransferAnyCase                BoolMarshaler `tfsdk:"permissions_transfer_any_case" force:",omitempty"`
	PermissionsContentAdministrator           BoolMarshaler `tfsdk:"permissions_content_administrator" force:",omitempty"`
	PermissionsCreateWorkspaces               BoolMarshaler `tfsdk:"permissions_create_workspaces" force:",omitempty"`
	PermissionsManageContentPermissions       BoolMarshaler `tfsdk:"permissions_manage_content_permissions" force:",omitempty"`
	PermissionsManageContentProperties        BoolMarshaler `tfsdk:"permissions_manage_content_properties" force:",omitempty"`
	PermissionsManageContentTypes             BoolMarshaler `tfsdk:"permissions_manage_content_types" force:",omitempty"`
	PermissionsManageExchangeConfig           BoolMarshaler `tfsdk:"permissions_manage_exchange_config" force:",omitempty"`
	PermissionsManageAnalyticSnapshots        BoolMarshaler `tfsdk:"permissions_manage_analytic_snapshots" force:",omitempty"`
	PermissionsScheduleReports                BoolMarshaler `tfsdk:"permissions_schedule_reports" force:",omitempty"`
	PermissionsManageBusinessHourHolidays     BoolMarshaler `tfsdk:"permissions_manage_business_hour_holidays" force:",omitempty"`
	PermissionsManageEntitlements             BoolMarshaler `tfsdk:"permissions_manage_entitlements" force:",omitempty"`
	PermissionsManageDynamicDashboards        BoolMarshaler `tfsdk:"permissions_manage_dynamic_dashboards" force:",omitempty"`
	PermissionsCustomSidebarOnAllPages        BoolMarshaler `tfsdk:"permissions_custom_sidebar_on_all_pages" force:",omitempty"`
	PermissionsManageInteraction              BoolMarshaler `tfsdk:"permissions_manage_interaction" force:",omitempty"`
	PermissionsViewMyTeamsDashboards          BoolMarshaler `tfsdk:"permissions_view_my_teams_dashboards" force:",omitempty"`
	PermissionsModerateChatter                BoolMarshaler `tfsdk:"permissions_moderate_chatter" force:",omitempty"`
	PermissionsResetPasswords                 BoolMarshaler `tfsdk:"permissions_reset_passwords" force:",omitempty"`
	PermissionsFlowUFLRequired                BoolMarshaler `tfsdk:"permissions_flow_ufl_required" force:",omitempty"`
	PermissionsCanInsertFeedSystemFields      BoolMarshaler `tfsdk:"permissions_can_insert_feed_system_fields" force:",omitempty"`
	PermissionsActivitiesAccess               BoolMarshaler `tfsdk:"permissions_activities_access" force:",omitempty"`
	PermissionsManageKnowledgeImportExport    BoolMarshaler `tfsdk:"permissions_manage_knowledge_import_export" force:",omitempty"`
	PermissionsEmailTemplateManagement        BoolMarshaler `tfsdk:"permissions_email_template_management" force:",omitempty"`
	PermissionsEmailAdministration            BoolMarshaler `tfsdk:"permissions_email_administration" force:",omitempty"`
	PermissionsManageChatterMessages          BoolMarshaler `tfsdk:"permissions_manage_chatter_messages" force:",omitempty"`
	PermissionsAllowEmailIC                   BoolMarshaler `tfsdk:"permissions_allow_email_ic" force:",omitempty"`
	PermissionsChatterFileLink                BoolMarshaler `tfsdk:"permissions_chatter_file_link" force:",omitempty"`
	PermissionsForceTwoFactor                 BoolMarshaler `tfsdk:"permissions_force_two_factor" force:",omitempty"`
	PermissionsViewEventLogFiles              BoolMarshaler `tfsdk:"permissions_view_event_log_files" force:",omitempty"`
	PermissionsManageNetworks                 BoolMarshaler `tfsdk:"permissions_manage_networks" force:",omitempty"`
	PermissionsManageAuthProviders            BoolMarshaler `tfsdk:"permissions_manage_auth_providers" force:",omitempty"`
	PermissionsRunFlow                        BoolMarshaler `tfsdk:"permissions_run_flow" force:",omitempty"`
	PermissionsCreateCustomizeDashboards      BoolMarshaler `tfsdk:"permissions_create_customize_dashboards" force:",omitempty"`
	PermissionsCreateDashboardFolders         BoolMarshaler `tfsdk:"permissions_create_dashboard_folders" force:",omitempty"`
	PermissionsViewPublicDashboards           BoolMarshaler `tfsdk:"permissions_view_public_dashboards" force:",omitempty"`
	PermissionsManageDashbdsInPubFolders      BoolMarshaler `tfsdk:"permissions_manage_dashbds_in_pub_folders" force:",omitempty"`
	PermissionsCreateCustomizeReports         BoolMarshaler `tfsdk:"permissions_create_customize_reports" force:",omitempty"`
	PermissionsCreateReportFolders            BoolMarshaler `tfsdk:"permissions_create_report_folders" force:",omitempty"`
	PermissionsViewPublicReports              BoolMarshaler `tfsdk:"permissions_view_public_reports" force:",omitempty"`
	PermissionsManageReportsInPubFolders      BoolMarshaler `tfsdk:"permissions_manage_reports_in_pub_folders" force:",omitempty"`
	PermissionsEditMyDashboards               BoolMarshaler `tfsdk:"permissions_edit_my_dashboards" force:",omitempty"`
	PermissionsEditMyReports                  BoolMarshaler `tfsdk:"permissions_edit_my_reports" force:",omitempty"`
	PermissionsViewAllUsers                   BoolMarshaler `tfsdk:"permissions_view_all_users" force:",omitempty"`
	PermissionsAllowUniversalSearch           BoolMarshaler `tfsdk:"permissions_allow_universal_search" force:",omitempty"`
	PermissionsConnectOrgToEnvironmentHub     BoolMarshaler `tfsdk:"permissions_connect_org_to_environment_hub" force:",omitempty"`
	PermissionsWorkCalibrationUser            BoolMarshaler `tfsdk:"permissions_work_calibration_user" force:",omitempty"`
	PermissionsCreateCustomizeFilters         BoolMarshaler `tfsdk:"permissions_create_customize_filters" force:",omitempty"`
	PermissionsWorkDotComUserPerm             BoolMarshaler `tfsdk:"permissions_work_dot_com_user_perm" force:",omitempty"`
	PermissionsContentHubUser                 BoolMarshaler `tfsdk:"permissions_content_hub_user" force:",omitempty"`
	PermissionsGovernNetworks                 BoolMarshaler `tfsdk:"permissions_govern_networks" force:",omitempty"`
	PermissionsSalesConsole                   BoolMarshaler `tfsdk:"permissions_sales_console" force:",omitempty"`
	PermissionsTwoFactorApi                   BoolMarshaler `tfsdk:"permissions_two_factor_api" force:",omitempty"`
	PermissionsDeleteTopics                   BoolMarshaler `tfsdk:"permissions_delete_topics" force:",omitempty"`
	PermissionsEditTopics                     BoolMarshaler `tfsdk:"permissions_edit_topics" force:",omitempty"`
	PermissionsCreateTopics                   BoolMarshaler `tfsdk:"permissions_create_topics" force:",omitempty"`
	PermissionsAssignTopics                   BoolMarshaler `tfsdk:"permissions_assign_topics" force:",omitempty"`
	PermissionsIdentityEnabled                BoolMarshaler `tfsdk:"permissions_identity_enabled" force:",omitempty"`
	PermissionsIdentityConnect                BoolMarshaler `tfsdk:"permissions_identity_connect" force:",omitempty"`
	PermissionsAllowViewKnowledge             BoolMarshaler `tfsdk:"permissions_allow_view_knowledge" force:",omitempty"`
	PermissionsContentWorkspaces              BoolMarshaler `tfsdk:"permissions_content_workspaces" force:",omitempty"`
	PermissionsCreateWorkBadgeDefinition      BoolMarshaler `tfsdk:"permissions_create_work_badge_definition" force:",omitempty"`
	PermissionsManageSearchPromotionRules     BoolMarshaler `tfsdk:"permissions_manage_search_promotion_rules" force:",omitempty"`
	PermissionsCustomMobileAppsAccess         BoolMarshaler `tfsdk:"permissions_custom_mobile_apps_access" force:",omitempty"`
	PermissionsViewHelpLink                   BoolMarshaler `tfsdk:"permissions_view_help_link" force:",omitempty"`
	PermissionsManageProfilesPermissionsets   BoolMarshaler `tfsdk:"permissions_manage_profiles_permissionsets" force:",omitempty"`
	PermissionsAssignPermissionSets           BoolMarshaler `tfsdk:"permissions_assign_permission_sets" force:",omitempty"`
	PermissionsManageRoles                    BoolMarshaler `tfsdk:"permissions_manage_roles" force:",omitempty"`
	PermissionsManageIpAddresses              BoolMarshaler `tfsdk:"permissions_manage_ip_addresses" force:",omitempty"`
	PermissionsManageSharing                  BoolMarshaler `tfsdk:"permissions_manage_sharing" force:",omitempty"`
	PermissionsManageInternalUsers            BoolMarshaler `tfsdk:"permissions_manage_internal_users" force:",omitempty"`
	PermissionsManagePasswordPolicies         BoolMarshaler `tfsdk:"permissions_manage_password_policies" force:",omitempty"`
	PermissionsManageLoginAccessPolicies      BoolMarshaler `tfsdk:"permissions_manage_login_access_policies" force:",omitempty"`
	PermissionsViewPlatformEvents             BoolMarshaler `tfsdk:"permissions_view_platform_events" force:",omitempty"`
	PermissionsManageCustomPermissions        BoolMarshaler `tfsdk:"permissions_manage_custom_permissions" force:",omitempty"`
	PermissionsCanVerifyComment               BoolMarshaler `tfsdk:"permissions_can_verify_comment" force:",omitempty"`
	PermissionsManageUnlistedGroups           BoolMarshaler `tfsdk:"permissions_manage_unlisted_groups" force:",omitempty"`
	PermissionsStdAutomaticActivityCapture    BoolMarshaler `tfsdk:"permissions_std_automatic_activity_capture" force:",omitempty"`
	PermissionsInsightsAppDashboardEditor     BoolMarshaler `tfsdk:"permissions_insights_app_dashboard_editor" force:",omitempty"`
	PermissionsManageTwoFactor                BoolMarshaler `tfsdk:"permissions_manage_two_factor" force:",omitempty"`
	PermissionsInsightsAppUser                BoolMarshaler `tfsdk:"permissions_insights_app_user" force:",omitempty"`
	PermissionsInsightsAppAdmin               BoolMarshaler `tfsdk:"permissions_insights_app_admin" force:",omitempty"`
	PermissionsInsightsAppEltEditor           BoolMarshaler `tfsdk:"permissions_insights_app_elt_editor" force:",omitempty"`
	PermissionsInsightsAppUploadUser          BoolMarshaler `tfsdk:"permissions_insights_app_upload_user" force:",omitempty"`
	PermissionsInsightsCreateApplication      BoolMarshaler `tfsdk:"permissions_insights_create_application" force:",omitempty"`
	PermissionsLightningExperienceUser        BoolMarshaler `tfsdk:"permissions_lightning_experience_user" force:",omitempty"`
	PermissionsViewDataLeakageEvents          BoolMarshaler `tfsdk:"permissions_view_data_leakage_events" force:",omitempty"`
	PermissionsConfigCustomRecs               BoolMarshaler `tfsdk:"permissions_config_custom_recs" force:",omitempty"`
	PermissionsSubmitMacrosAllowed            BoolMarshaler `tfsdk:"permissions_submit_macros_allowed" force:",omitempty"`
	PermissionsBulkMacrosAllowed              BoolMarshaler `tfsdk:"permissions_bulk_macros_allowed" force:",omitempty"`
	PermissionsShareInternalArticles          BoolMarshaler `tfsdk:"permissions_share_internal_articles" force:",omitempty"`
	PermissionsManageSessionPermissionSets    BoolMarshaler `tfsdk:"permissions_manage_session_permission_sets" force:",omitempty"`
	PermissionsManageTemplatedApp             BoolMarshaler `tfsdk:"permissions_manage_templated_app" force:",omitempty"`
	PermissionsUseTemplatedApp                BoolMarshaler `tfsdk:"permissions_use_templated_app" force:",omitempty"`
	PermissionsSendAnnouncementEmails         BoolMarshaler `tfsdk:"permissions_send_announcement_emails" force:",omitempty"`
	PermissionsChatterEditOwnPost             BoolMarshaler `tfsdk:"permissions_chatter_edit_own_post" force:",omitempty"`
	PermissionsChatterEditOwnRecordPost       BoolMarshaler `tfsdk:"permissions_chatter_edit_own_record_post" force:",omitempty"`
	PermissionsSalesAnalyticsUser             BoolMarshaler `tfsdk:"permissions_sales_analytics_user" force:",omitempty"`
	PermissionsServiceAnalyticsUser           BoolMarshaler `tfsdk:"permissions_service_analytics_user" force:",omitempty"`
	PermissionsWaveTabularDownload            BoolMarshaler `tfsdk:"permissions_wave_tabular_download" force:",omitempty"`
	PermissionsAutomaticActivityCapture       BoolMarshaler `tfsdk:"permissions_automatic_activity_capture" force:",omitempty"`
	PermissionsImportCustomObjects            BoolMarshaler `tfsdk:"permissions_import_custom_objects" force:",omitempty"`
	PermissionsDelegatedTwoFactor             BoolMarshaler `tfsdk:"permissions_delegated_two_factor" force:",omitempty"`
	PermissionsChatterComposeUiCodesnippet    BoolMarshaler `tfsdk:"permissions_chatter_compose_ui_codesnippet" force:",omitempty"`
	PermissionsSelectFilesFromSalesforce      BoolMarshaler `tfsdk:"permissions_select_files_from_salesforce" force:",omitempty"`
	PermissionsModerateNetworkUsers           BoolMarshaler `tfsdk:"permissions_moderate_network_users" force:",omitempty"`
	PermissionsMergeTopics                    BoolMarshaler `tfsdk:"permissions_merge_topics" force:",omitempty"`
	PermissionsSubscribeToLightningReports    BoolMarshaler `tfsdk:"permissions_subscribe_to_lightning_reports" force:",omitempty"`
	PermissionsManagePvtRptsAndDashbds        BoolMarshaler `tfsdk:"permissions_manage_pvt_rpts_and_dashbds" force:",omitempty"`
	PermissionsAllowLightningLogin            BoolMarshaler `tfsdk:"permissions_allow_lightning_login" force:",omitempty"`
	PermissionsCampaignInfluence2             BoolMarshaler `tfsdk:"permissions_campaign_influence2" force:",omitempty"`
	PermissionsViewDataAssessment             BoolMarshaler `tfsdk:"permissions_view_data_assessment" force:",omitempty"`
	PermissionsRemoveDirectMessageMembers     BoolMarshaler `tfsdk:"permissions_remove_direct_message_members" force:",omitempty"`
	PermissionsCanApproveFeedPost             BoolMarshaler `tfsdk:"permissions_can_approve_feed_post" force:",omitempty"`
	PermissionsAddDirectMessageMembers        BoolMarshaler `tfsdk:"permissions_add_direct_message_members" force:",omitempty"`
	PermissionsAllowViewEditConvertedLeads    BoolMarshaler `tfsdk:"permissions_allow_view_edit_converted_leads" force:",omitempty"`
	PermissionsShowCompanyNameAsUserBadge     BoolMarshaler `tfsdk:"permissions_show_company_name_as_user_badge" force:",omitempty"`
	PermissionsAccessCMC                      BoolMarshaler `tfsdk:"permissions_access_cmc" force:",omitempty"`
	PermissionsViewHealthCheck                BoolMarshaler `tfsdk:"permissions_view_health_check" force:",omitempty"`
	PermissionsManageHealthCheck              BoolMarshaler `tfsdk:"permissions_manage_health_check" force:",omitempty"`
	PermissionsPackaging2                     BoolMarshaler `tfsdk:"permissions_packaging2" force:",omitempty"`
	PermissionsManageCertificates             BoolMarshaler `tfsdk:"permissions_manage_certificates" force:",omitempty"`
	PermissionsCreateReportInLightning        BoolMarshaler `tfsdk:"permissions_create_report_in_lightning" force:",omitempty"`
	PermissionsPreventClassicExperience       BoolMarshaler `tfsdk:"permissions_prevent_classic_experience" force:",omitempty"`
	PermissionsHideReadByList                 BoolMarshaler `tfsdk:"permissions_hide_read_by_list" force:",omitempty"`
	PermissionsUseSmartDataDiscovery          BoolMarshaler `tfsdk:"permissions_use_smart_data_discovery" force:",omitempty"`
	PermissionsGetSmartDataDiscovery          BoolMarshaler `tfsdk:"permissions_get_smart_data_discovery" force:",omitempty"`
	PermissionsCreateUpdateSDDDataset         BoolMarshaler `tfsdk:"permissions_create_update_sdd_dataset" force:",omitempty"`
	PermissionsCreateUpdateSDDStory           BoolMarshaler `tfsdk:"permissions_create_update_sdd_story" force:",omitempty"`
	PermissionsManageSmartDataDiscovery       BoolMarshaler `tfsdk:"permissions_manage_smart_data_discovery" force:",omitempty"`
	PermissionsShareSmartDataDiscoveryStory   BoolMarshaler `tfsdk:"permissions_share_smart_data_discovery_story" force:",omitempty"`
	PermissionsManageSmartDataDiscoveryModel  BoolMarshaler `tfsdk:"permissions_manage_smart_data_discovery_model" force:",omitempty"`
	PermissionsListEmailSend                  BoolMarshaler `tfsdk:"permissions_list_email_send" force:",omitempty"`
	PermissionsFeedPinning                    BoolMarshaler `tfsdk:"permissions_feed_pinning" force:",omitempty"`
	PermissionsChangeDashboardColors          BoolMarshaler `tfsdk:"permissions_change_dashboard_colors" force:",omitempty"`
	PermissionsIotUser                        BoolMarshaler `tfsdk:"permissions_iot_user" force:",omitempty"`
	PermissionsManageRecommendationStrategies BoolMarshaler `tfsdk:"permissions_manage_recommendation_strategies" force:",omitempty"`
	PermissionsManagePropositions             BoolMarshaler `tfsdk:"permissions_manage_propositions" force:",omitempty"`
	PermissionsCloseConversations             BoolMarshaler `tfsdk:"permissions_close_conversations" force:",omitempty"`
	PermissionsSubscribeReportRolesGrps       BoolMarshaler `tfsdk:"permissions_subscribe_report_roles_grps" force:",omitempty"`
	PermissionsSubscribeDashboardRolesGrps    BoolMarshaler `tfsdk:"permissions_subscribe_dashboard_roles_grps" force:",omitempty"`
	PermissionsUseWebLink                     BoolMarshaler `tfsdk:"permissions_use_web_link" force:",omitempty"`
	PermissionsHasUnlimitedNBAExecutions      BoolMarshaler `tfsdk:"permissions_has_unlimited_nba_executions" force:",omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        BoolMarshaler `tfsdk:"permissions_view_only_embedded_app_user" force:",omitempty"`
	PermissionsAdoptionAnalyticsUser          BoolMarshaler `tfsdk:"permissions_adoption_analytics_user" force:",omitempty"`
	PermissionsViewAllActivities              BoolMarshaler `tfsdk:"permissions_view_all_activities" force:",omitempty"`
	PermissionsSubscribeReportToOtherUsers    BoolMarshaler `tfsdk:"permissions_subscribe_report_to_other_users" force:",omitempty"`
	PermissionsLightningConsoleAllowedForUser BoolMarshaler `tfsdk:"permissions_lightning_console_allowed_for_user" force:",omitempty"`
	PermissionsSubscribeReportsRunAsUser      BoolMarshaler `tfsdk:"permissions_subscribe_reports_run_as_user" force:",omitempty"`
	PermissionsSubscribeToLightningDashboards BoolMarshaler `tfsdk:"permissions_subscribe_to_lightning_dashboards" force:",omitempty"`
	PermissionsSubscribeDashboardToOtherUsers BoolMarshaler `tfsdk:"permissions_subscribe_dashboard_to_other_users" force:",omitempty"`
	PermissionsCreateLtngTempInPub            BoolMarshaler `tfsdk:"permissions_create_ltng_temp_in_pub" force:",omitempty"`
	PermissionsTransactionalEmailSend         BoolMarshaler `tfsdk:"permissions_transactional_email_send" force:",omitempty"`
	PermissionsViewPrivateStaticResources     BoolMarshaler `tfsdk:"permissions_view_private_static_resources" force:",omitempty"`
	PermissionsCreateLtngTempFolder           BoolMarshaler `tfsdk:"permissions_create_ltng_temp_folder" force:",omitempty"`
	PermissionsApexRestServices               BoolMarshaler `tfsdk:"permissions_apex_rest_services" force:",omitempty"`
	PermissionsEnableCommunityAppLauncher     BoolMarshaler `tfsdk:"permissions_enable_community_app_launcher" force:",omitempty"`
	PermissionsGiveRecognitionBadge           BoolMarshaler `tfsdk:"permissions_give_recognition_badge" force:",omitempty"`
	PermissionsCanRunAnalysis                 BoolMarshaler `tfsdk:"permissions_can_run_analysis" force:",omitempty"`
	PermissionsUseMySearch                    BoolMarshaler `tfsdk:"permissions_use_my_search" force:",omitempty"`
	PermissionsLtngPromoReserved01UserPerm    BoolMarshaler `tfsdk:"permissions_ltng_promo_reserved01_user_perm" force:",omitempty"`
	PermissionsManageSubscriptions            BoolMarshaler `tfsdk:"permissions_manage_subscriptions" force:",omitempty"`
	PermissionsWaveManagePrivateAssetsUser    BoolMarshaler `tfsdk:"permissions_wave_manage_private_assets_user" force:",omitempty"`
	PermissionsCanEditDataPrepRecipe          BoolMarshaler `tfsdk:"permissions_can_edit_data_prep_recipe" force:",omitempty"`
	PermissionsAddAnalyticsRemoteConnections  BoolMarshaler `tfsdk:"permissions_add_analytics_remote_connections" force:",omitempty"`
	PermissionsManageSurveys                  BoolMarshaler `tfsdk:"permissions_manage_surveys" force:",omitempty"`
	PermissionsUseAssistantDialog             BoolMarshaler `tfsdk:"permissions_use_assistant_dialog" force:",omitempty"`
	PermissionsUseQuerySuggestions            BoolMarshaler `tfsdk:"permissions_use_query_suggestions" force:",omitempty"`
	PermissionsRecordVisibilityAPI            BoolMarshaler `tfsdk:"permissions_record_visibility_api" force:",omitempty"`
	PermissionsViewRoles                      BoolMarshaler `tfsdk:"permissions_view_roles" force:",omitempty"`
	PermissionsSmartDataDiscoveryForCommunity BoolMarshaler `tfsdk:"permissions_smart_data_discovery_for_community" force:",omitempty"`
	PermissionsCanManageMaps                  BoolMarshaler `tfsdk:"permissions_can_manage_maps" force:",omitempty"`
	PermissionsStoryOnDSWithPredicate         BoolMarshaler `tfsdk:"permissions_story_on_ds_with_predicate" force:",omitempty"`
	PermissionsLMOutboundMessagingUserPerm    BoolMarshaler `tfsdk:"permissions_lm_outbound_messaging_user_perm" force:",omitempty"`
	PermissionsModifyDataClassification       BoolMarshaler `tfsdk:"permissions_modify_data_classification" force:",omitempty"`
	PermissionsPrivacyDataAccess              BoolMarshaler `tfsdk:"permissions_privacy_data_access" force:",omitempty"`
	PermissionsQueryAllFiles                  BoolMarshaler `tfsdk:"permissions_query_all_files" force:",omitempty"`
	PermissionsModifyMetadata                 BoolMarshaler `tfsdk:"permissions_modify_metadata" force:",omitempty"`
	PermissionsManageCMS                      BoolMarshaler `tfsdk:"permissions_manage_cms" force:",omitempty"`
	PermissionsSandboxTestingInCommunityApp   BoolMarshaler `tfsdk:"permissions_sandbox_testing_in_community_app" force:",omitempty"`
	PermissionsCanEditPrompts                 BoolMarshaler `tfsdk:"permissions_can_edit_prompts" force:",omitempty"`
	PermissionsViewUserPII                    BoolMarshaler `tfsdk:"permissions_view_user_pii" force:",omitempty"`
	PermissionsManageHubConnections           BoolMarshaler `tfsdk:"permissions_manage_hub_connections" force:",omitempty"`
	PermissionsB2BMarketingAnalyticsUser      BoolMarshaler `tfsdk:"permissions_b_2_b_marketing_analytics_user" force:",omitempty"`
	PermissionsTraceXdsQueries                BoolMarshaler `tfsdk:"permissions_trace_xds_queries" force:",omitempty"`
	PermissionsViewSecurityCommandCenter      BoolMarshaler `tfsdk:"permissions_view_security_command_center" force:",omitempty"`
	PermissionsManageSecurityCommandCenter    BoolMarshaler `tfsdk:"permissions_manage_security_command_center" force:",omitempty"`
	PermissionsViewAllCustomSettings          BoolMarshaler `tfsdk:"permissions_view_all_custom_settings" force:",omitempty"`
	PermissionsViewAllForeignKeyNames         BoolMarshaler `tfsdk:"permissions_view_all_foreign_key_names" force:",omitempty"`
	PermissionsAllowSurveyAdvancedFeatures    BoolMarshaler `tfsdk:"permissions_allow_survey_advanced_features" force:",omitempty"`
	PermissionsAddWaveNotificationRecipients  BoolMarshaler `tfsdk:"permissions_add_wave_notification_recipients" force:",omitempty"`
	PermissionsHeadlessCMSAccess              BoolMarshaler `tfsdk:"permissions_headless_cms_access" force:",omitempty"`
	PermissionsLMEndMessagingSessionUserPerm  BoolMarshaler `tfsdk:"permissions_lm_end_messaging_session_user_perm" force:",omitempty"`
	PermissionsConsentApiUpdate               BoolMarshaler `tfsdk:"permissions_consent_api_update" force:",omitempty"`
	PermissionsPaymentsAPIUser                BoolMarshaler `tfsdk:"permissions_payments_api_user" force:",omitempty"`
	PermissionsAccessContentBuilder           BoolMarshaler `tfsdk:"permissions_access_content_builder" force:",omitempty"`
	PermissionsAccountSwitcherUser            BoolMarshaler `tfsdk:"permissions_account_switcher_user" force:",omitempty"`
	PermissionsViewAnomalyEvents              BoolMarshaler `tfsdk:"permissions_view_anomaly_events" force:",omitempty"`
	PermissionsManageC360AConnections         BoolMarshaler `tfsdk:"permissions_manage_c_360_a_connections" force:",omitempty"`
	PermissionsManageReleaseUpdates           BoolMarshaler `tfsdk:"permissions_manage_release_updates" force:",omitempty"`
	PermissionsViewAllProfiles                BoolMarshaler `tfsdk:"permissions_view_all_profiles" force:",omitempty"`
	PermissionsSkipIdentityConfirmation       BoolMarshaler `tfsdk:"permissions_skip_identity_confirmation" force:",omitempty"`
	PermissionsLearningManager                BoolMarshaler `tfsdk:"permissions_learning_manager" force:",omitempty"`
	PermissionsSendCustomNotifications        BoolMarshaler `tfsdk:"permissions_send_custom_notifications" force:",omitempty"`
	PermissionsPackaging2Delete               BoolMarshaler `tfsdk:"permissions_packaging2_delete" force:",omitempty"`
	PermissionsAutonomousAnalyticsPrivacy     BoolMarshaler `tfsdk:"permissions_autonomous_analytics_privacy" force:",omitempty"`
	PermissionsSonicConsumer                  BoolMarshaler `tfsdk:"permissions_sonic_consumer" force:",omitempty"`
	PermissionsBotManageBots                  BoolMarshaler `tfsdk:"permissions_bot_manage_bots" force:",omitempty"`
	PermissionsBotManageBotsTrainingData      BoolMarshaler `tfsdk:"permissions_bot_manage_bots_training_data" force:",omitempty"`
	PermissionsManageLearningReporting        BoolMarshaler `tfsdk:"permissions_manage_learning_reporting" force:",omitempty"`
	PermissionsIsotopeCToCUser                BoolMarshaler `tfsdk:"permissions_isotope_c_to_c_user" force:",omitempty"`
	PermissionsIsotopeAccess                  BoolMarshaler `tfsdk:"permissions_isotope_access" force:",omitempty"`
	PermissionsIsotopeLEX                     BoolMarshaler `tfsdk:"permissions_isotope_lex" force:",omitempty"`
	PermissionsGetSmartDataDiscoveryExternal  BoolMarshaler `tfsdk:"permissions_get_smart_data_discovery_external" force:",omitempty"`
	PermissionsQuipMetricsAccess              BoolMarshaler `tfsdk:"permissions_quip_metrics_access" force:",omitempty"`
	PermissionsQuipUserEngagementMetrics      BoolMarshaler `tfsdk:"permissions_quip_user_engagement_metrics" force:",omitempty"`
	PermissionsManageExternalConnections      BoolMarshaler `tfsdk:"permissions_manage_external_connections" force:",omitempty"`
	PermissionsUseSubscriptionEmails          BoolMarshaler `tfsdk:"permissions_use_subscription_emails" force:",omitempty"`
	PermissionsAIViewInsightObjects           BoolMarshaler `tfsdk:"permissions_ai_view_insight_objects" force:",omitempty"`
	PermissionsAICreateInsightObjects         BoolMarshaler `tfsdk:"permissions_ai_create_insight_objects" force:",omitempty"`
	PermissionsLifecycleManagementAPIUser     BoolMarshaler `tfsdk:"permissions_lifecycle_management_api_user" force:",omitempty"`
	PermissionsNativeWebviewScrolling         BoolMarshaler `tfsdk:"permissions_native_webview_scrolling" force:",omitempty"`
	PermissionsViewDeveloperName              BoolMarshaler `tfsdk:"permissions_view_developer_name" force:",omitempty"`
}

// In order to support json omitempty for boolean fields, a new type has to be copied into
type profileResourceDataJSON struct {
	Name          string  `force:",omitempty"`
	Description   *string `force:",omitempty"`
	UserLicenseId string  `force:",omitempty"`
	// permissions
	PermissionsEmailSingle                    *bool `force:",omitempty"`
	PermissionsEmailMass                      *bool `force:",omitempty"`
	PermissionsEditTask                       *bool `force:",omitempty"`
	PermissionsEditEvent                      *bool `force:",omitempty"`
	PermissionsExportReport                   *bool `force:",omitempty"`
	PermissionsImportPersonal                 *bool `force:",omitempty"`
	PermissionsDataExport                     *bool `force:",omitempty"`
	PermissionsManageUsers                    *bool `force:",omitempty"`
	PermissionsEditPublicFilters              *bool `force:",omitempty"`
	PermissionsEditPublicTemplates            *bool `force:",omitempty"`
	PermissionsModifyAllData                  *bool `force:",omitempty"`
	PermissionsManageCases                    *bool `force:",omitempty"`
	PermissionsMassInlineEdit                 *bool `force:",omitempty"`
	PermissionsEditKnowledge                  *bool `force:",omitempty"`
	PermissionsManageKnowledge                *bool `force:",omitempty"`
	PermissionsManageSolutions                *bool `force:",omitempty"`
	PermissionsCustomizeApplication           *bool `force:",omitempty"`
	PermissionsEditReadonlyFields             *bool `force:",omitempty"`
	PermissionsRunReports                     *bool `force:",omitempty"`
	PermissionsViewSetup                      *bool `force:",omitempty"`
	PermissionsTransferAnyEntity              *bool `force:",omitempty"`
	PermissionsNewReportBuilder               *bool `force:",omitempty"`
	PermissionsActivateContract               *bool `force:",omitempty"`
	PermissionsActivateOrder                  *bool `force:",omitempty"`
	PermissionsImportLeads                    *bool `force:",omitempty"`
	PermissionsManageLeads                    *bool `force:",omitempty"`
	PermissionsTransferAnyLead                *bool `force:",omitempty"`
	PermissionsViewAllData                    *bool `force:",omitempty"`
	PermissionsEditPublicDocuments            *bool `force:",omitempty"`
	PermissionsViewEncryptedData              *bool `force:",omitempty"`
	PermissionsEditBrandTemplates             *bool `force:",omitempty"`
	PermissionsEditHtmlTemplates              *bool `force:",omitempty"`
	PermissionsChatterInternalUser            *bool `force:",omitempty"`
	PermissionsManageEncryptionKeys           *bool `force:",omitempty"`
	PermissionsDeleteActivatedContract        *bool `force:",omitempty"`
	PermissionsChatterInviteExternalUsers     *bool `force:",omitempty"`
	PermissionsSendSitRequests                *bool `force:",omitempty"`
	PermissionsManageRemoteAccess             *bool `force:",omitempty"`
	PermissionsCanUseNewDashboardBuilder      *bool `force:",omitempty"`
	PermissionsManageCategories               *bool `force:",omitempty"`
	PermissionsConvertLeads                   *bool `force:",omitempty"`
	PermissionsPasswordNeverExpires           *bool `force:",omitempty"`
	PermissionsUseTeamReassignWizards         *bool `force:",omitempty"`
	PermissionsEditActivatedOrders            *bool `force:",omitempty"`
	PermissionsInstallMultiforce              *bool `force:",omitempty"`
	PermissionsPublishMultiforce              *bool `force:",omitempty"`
	PermissionsChatterOwnGroups               *bool `force:",omitempty"`
	PermissionsEditOppLineItemUnitPrice       *bool `force:",omitempty"`
	PermissionsCreateMultiforce               *bool `force:",omitempty"`
	PermissionsBulkApiHardDelete              *bool `force:",omitempty"`
	PermissionsSolutionImport                 *bool `force:",omitempty"`
	PermissionsManageCallCenters              *bool `force:",omitempty"`
	PermissionsManageSynonyms                 *bool `force:",omitempty"`
	PermissionsViewContent                    *bool `force:",omitempty"`
	PermissionsManageEmailClientConfig        *bool `force:",omitempty"`
	PermissionsEnableNotifications            *bool `force:",omitempty"`
	PermissionsManageDataIntegrations         *bool `force:",omitempty"`
	PermissionsDistributeFromPersWksp         *bool `force:",omitempty"`
	PermissionsViewDataCategories             *bool `force:",omitempty"`
	PermissionsManageDataCategories           *bool `force:",omitempty"`
	PermissionsAuthorApex                     *bool `force:",omitempty"`
	PermissionsManageMobile                   *bool `force:",omitempty"`
	PermissionsApiEnabled                     *bool `force:",omitempty"`
	PermissionsManageCustomReportTypes        *bool `force:",omitempty"`
	PermissionsEditCaseComments               *bool `force:",omitempty"`
	PermissionsTransferAnyCase                *bool `force:",omitempty"`
	PermissionsContentAdministrator           *bool `force:",omitempty"`
	PermissionsCreateWorkspaces               *bool `force:",omitempty"`
	PermissionsManageContentPermissions       *bool `force:",omitempty"`
	PermissionsManageContentProperties        *bool `force:",omitempty"`
	PermissionsManageContentTypes             *bool `force:",omitempty"`
	PermissionsManageExchangeConfig           *bool `force:",omitempty"`
	PermissionsManageAnalyticSnapshots        *bool `force:",omitempty"`
	PermissionsScheduleReports                *bool `force:",omitempty"`
	PermissionsManageBusinessHourHolidays     *bool `force:",omitempty"`
	PermissionsManageEntitlements             *bool `force:",omitempty"`
	PermissionsManageDynamicDashboards        *bool `force:",omitempty"`
	PermissionsCustomSidebarOnAllPages        *bool `force:",omitempty"`
	PermissionsManageInteraction              *bool `force:",omitempty"`
	PermissionsViewMyTeamsDashboards          *bool `force:",omitempty"`
	PermissionsModerateChatter                *bool `force:",omitempty"`
	PermissionsResetPasswords                 *bool `force:",omitempty"`
	PermissionsFlowUFLRequired                *bool `force:",omitempty"`
	PermissionsCanInsertFeedSystemFields      *bool `force:",omitempty"`
	PermissionsActivitiesAccess               *bool `force:",omitempty"`
	PermissionsManageKnowledgeImportExport    *bool `force:",omitempty"`
	PermissionsEmailTemplateManagement        *bool `force:",omitempty"`
	PermissionsEmailAdministration            *bool `force:",omitempty"`
	PermissionsManageChatterMessages          *bool `force:",omitempty"`
	PermissionsAllowEmailIC                   *bool `force:",omitempty"`
	PermissionsChatterFileLink                *bool `force:",omitempty"`
	PermissionsForceTwoFactor                 *bool `force:",omitempty"`
	PermissionsViewEventLogFiles              *bool `force:",omitempty"`
	PermissionsManageNetworks                 *bool `force:",omitempty"`
	PermissionsManageAuthProviders            *bool `force:",omitempty"`
	PermissionsRunFlow                        *bool `force:",omitempty"`
	PermissionsCreateCustomizeDashboards      *bool `force:",omitempty"`
	PermissionsCreateDashboardFolders         *bool `force:",omitempty"`
	PermissionsViewPublicDashboards           *bool `force:",omitempty"`
	PermissionsManageDashbdsInPubFolders      *bool `force:",omitempty"`
	PermissionsCreateCustomizeReports         *bool `force:",omitempty"`
	PermissionsCreateReportFolders            *bool `force:",omitempty"`
	PermissionsViewPublicReports              *bool `force:",omitempty"`
	PermissionsManageReportsInPubFolders      *bool `force:",omitempty"`
	PermissionsEditMyDashboards               *bool `force:",omitempty"`
	PermissionsEditMyReports                  *bool `force:",omitempty"`
	PermissionsViewAllUsers                   *bool `force:",omitempty"`
	PermissionsAllowUniversalSearch           *bool `force:",omitempty"`
	PermissionsConnectOrgToEnvironmentHub     *bool `force:",omitempty"`
	PermissionsWorkCalibrationUser            *bool `force:",omitempty"`
	PermissionsCreateCustomizeFilters         *bool `force:",omitempty"`
	PermissionsWorkDotComUserPerm             *bool `force:",omitempty"`
	PermissionsContentHubUser                 *bool `force:",omitempty"`
	PermissionsGovernNetworks                 *bool `force:",omitempty"`
	PermissionsSalesConsole                   *bool `force:",omitempty"`
	PermissionsTwoFactorApi                   *bool `force:",omitempty"`
	PermissionsDeleteTopics                   *bool `force:",omitempty"`
	PermissionsEditTopics                     *bool `force:",omitempty"`
	PermissionsCreateTopics                   *bool `force:",omitempty"`
	PermissionsAssignTopics                   *bool `force:",omitempty"`
	PermissionsIdentityEnabled                *bool `force:",omitempty"`
	PermissionsIdentityConnect                *bool `force:",omitempty"`
	PermissionsAllowViewKnowledge             *bool `force:",omitempty"`
	PermissionsContentWorkspaces              *bool `force:",omitempty"`
	PermissionsCreateWorkBadgeDefinition      *bool `force:",omitempty"`
	PermissionsManageSearchPromotionRules     *bool `force:",omitempty"`
	PermissionsCustomMobileAppsAccess         *bool `force:",omitempty"`
	PermissionsViewHelpLink                   *bool `force:",omitempty"`
	PermissionsManageProfilesPermissionsets   *bool `force:",omitempty"`
	PermissionsAssignPermissionSets           *bool `force:",omitempty"`
	PermissionsManageRoles                    *bool `force:",omitempty"`
	PermissionsManageIpAddresses              *bool `force:",omitempty"`
	PermissionsManageSharing                  *bool `force:",omitempty"`
	PermissionsManageInternalUsers            *bool `force:",omitempty"`
	PermissionsManagePasswordPolicies         *bool `force:",omitempty"`
	PermissionsManageLoginAccessPolicies      *bool `force:",omitempty"`
	PermissionsViewPlatformEvents             *bool `force:",omitempty"`
	PermissionsManageCustomPermissions        *bool `force:",omitempty"`
	PermissionsCanVerifyComment               *bool `force:",omitempty"`
	PermissionsManageUnlistedGroups           *bool `force:",omitempty"`
	PermissionsStdAutomaticActivityCapture    *bool `force:",omitempty"`
	PermissionsInsightsAppDashboardEditor     *bool `force:",omitempty"`
	PermissionsManageTwoFactor                *bool `force:",omitempty"`
	PermissionsInsightsAppUser                *bool `force:",omitempty"`
	PermissionsInsightsAppAdmin               *bool `force:",omitempty"`
	PermissionsInsightsAppEltEditor           *bool `force:",omitempty"`
	PermissionsInsightsAppUploadUser          *bool `force:",omitempty"`
	PermissionsInsightsCreateApplication      *bool `force:",omitempty"`
	PermissionsLightningExperienceUser        *bool `force:",omitempty"`
	PermissionsViewDataLeakageEvents          *bool `force:",omitempty"`
	PermissionsConfigCustomRecs               *bool `force:",omitempty"`
	PermissionsSubmitMacrosAllowed            *bool `force:",omitempty"`
	PermissionsBulkMacrosAllowed              *bool `force:",omitempty"`
	PermissionsShareInternalArticles          *bool `force:",omitempty"`
	PermissionsManageSessionPermissionSets    *bool `force:",omitempty"`
	PermissionsManageTemplatedApp             *bool `force:",omitempty"`
	PermissionsUseTemplatedApp                *bool `force:",omitempty"`
	PermissionsSendAnnouncementEmails         *bool `force:",omitempty"`
	PermissionsChatterEditOwnPost             *bool `force:",omitempty"`
	PermissionsChatterEditOwnRecordPost       *bool `force:",omitempty"`
	PermissionsSalesAnalyticsUser             *bool `force:",omitempty"`
	PermissionsServiceAnalyticsUser           *bool `force:",omitempty"`
	PermissionsWaveTabularDownload            *bool `force:",omitempty"`
	PermissionsAutomaticActivityCapture       *bool `force:",omitempty"`
	PermissionsImportCustomObjects            *bool `force:",omitempty"`
	PermissionsDelegatedTwoFactor             *bool `force:",omitempty"`
	PermissionsChatterComposeUiCodesnippet    *bool `force:",omitempty"`
	PermissionsSelectFilesFromSalesforce      *bool `force:",omitempty"`
	PermissionsModerateNetworkUsers           *bool `force:",omitempty"`
	PermissionsMergeTopics                    *bool `force:",omitempty"`
	PermissionsSubscribeToLightningReports    *bool `force:",omitempty"`
	PermissionsManagePvtRptsAndDashbds        *bool `force:",omitempty"`
	PermissionsAllowLightningLogin            *bool `force:",omitempty"`
	PermissionsCampaignInfluence2             *bool `force:",omitempty"`
	PermissionsViewDataAssessment             *bool `force:",omitempty"`
	PermissionsRemoveDirectMessageMembers     *bool `force:",omitempty"`
	PermissionsCanApproveFeedPost             *bool `force:",omitempty"`
	PermissionsAddDirectMessageMembers        *bool `force:",omitempty"`
	PermissionsAllowViewEditConvertedLeads    *bool `force:",omitempty"`
	PermissionsShowCompanyNameAsUserBadge     *bool `force:",omitempty"`
	PermissionsAccessCMC                      *bool `force:",omitempty"`
	PermissionsViewHealthCheck                *bool `force:",omitempty"`
	PermissionsManageHealthCheck              *bool `force:",omitempty"`
	PermissionsPackaging2                     *bool `force:",omitempty"`
	PermissionsManageCertificates             *bool `force:",omitempty"`
	PermissionsCreateReportInLightning        *bool `force:",omitempty"`
	PermissionsPreventClassicExperience       *bool `force:",omitempty"`
	PermissionsHideReadByList                 *bool `force:",omitempty"`
	PermissionsUseSmartDataDiscovery          *bool `force:",omitempty"`
	PermissionsGetSmartDataDiscovery          *bool `force:",omitempty"`
	PermissionsCreateUpdateSDDDataset         *bool `force:",omitempty"`
	PermissionsCreateUpdateSDDStory           *bool `force:",omitempty"`
	PermissionsManageSmartDataDiscovery       *bool `force:",omitempty"`
	PermissionsShareSmartDataDiscoveryStory   *bool `force:",omitempty"`
	PermissionsManageSmartDataDiscoveryModel  *bool `force:",omitempty"`
	PermissionsListEmailSend                  *bool `force:",omitempty"`
	PermissionsFeedPinning                    *bool `force:",omitempty"`
	PermissionsChangeDashboardColors          *bool `force:",omitempty"`
	PermissionsIotUser                        *bool `force:",omitempty"`
	PermissionsManageRecommendationStrategies *bool `force:",omitempty"`
	PermissionsManagePropositions             *bool `force:",omitempty"`
	PermissionsCloseConversations             *bool `force:",omitempty"`
	PermissionsSubscribeReportRolesGrps       *bool `force:",omitempty"`
	PermissionsSubscribeDashboardRolesGrps    *bool `force:",omitempty"`
	PermissionsUseWebLink                     *bool `force:",omitempty"`
	PermissionsHasUnlimitedNBAExecutions      *bool `force:",omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        *bool `force:",omitempty"`
	PermissionsAdoptionAnalyticsUser          *bool `force:",omitempty"`
	PermissionsViewAllActivities              *bool `force:",omitempty"`
	PermissionsSubscribeReportToOtherUsers    *bool `force:",omitempty"`
	PermissionsLightningConsoleAllowedForUser *bool `force:",omitempty"`
	PermissionsSubscribeReportsRunAsUser      *bool `force:",omitempty"`
	PermissionsSubscribeToLightningDashboards *bool `force:",omitempty"`
	PermissionsSubscribeDashboardToOtherUsers *bool `force:",omitempty"`
	PermissionsCreateLtngTempInPub            *bool `force:",omitempty"`
	PermissionsTransactionalEmailSend         *bool `force:",omitempty"`
	PermissionsViewPrivateStaticResources     *bool `force:",omitempty"`
	PermissionsCreateLtngTempFolder           *bool `force:",omitempty"`
	PermissionsApexRestServices               *bool `force:",omitempty"`
	PermissionsEnableCommunityAppLauncher     *bool `force:",omitempty"`
	PermissionsGiveRecognitionBadge           *bool `force:",omitempty"`
	PermissionsCanRunAnalysis                 *bool `force:",omitempty"`
	PermissionsUseMySearch                    *bool `force:",omitempty"`
	PermissionsLtngPromoReserved01UserPerm    *bool `force:",omitempty"`
	PermissionsManageSubscriptions            *bool `force:",omitempty"`
	PermissionsWaveManagePrivateAssetsUser    *bool `force:",omitempty"`
	PermissionsCanEditDataPrepRecipe          *bool `force:",omitempty"`
	PermissionsAddAnalyticsRemoteConnections  *bool `force:",omitempty"`
	PermissionsManageSurveys                  *bool `force:",omitempty"`
	PermissionsUseAssistantDialog             *bool `force:",omitempty"`
	PermissionsUseQuerySuggestions            *bool `force:",omitempty"`
	PermissionsRecordVisibilityAPI            *bool `force:",omitempty"`
	PermissionsViewRoles                      *bool `force:",omitempty"`
	PermissionsSmartDataDiscoveryForCommunity *bool `force:",omitempty"`
	PermissionsCanManageMaps                  *bool `force:",omitempty"`
	PermissionsStoryOnDSWithPredicate         *bool `force:",omitempty"`
	PermissionsLMOutboundMessagingUserPerm    *bool `force:",omitempty"`
	PermissionsModifyDataClassification       *bool `force:",omitempty"`
	PermissionsPrivacyDataAccess              *bool `force:",omitempty"`
	PermissionsQueryAllFiles                  *bool `force:",omitempty"`
	PermissionsModifyMetadata                 *bool `force:",omitempty"`
	PermissionsManageCMS                      *bool `force:",omitempty"`
	PermissionsSandboxTestingInCommunityApp   *bool `force:",omitempty"`
	PermissionsCanEditPrompts                 *bool `force:",omitempty"`
	PermissionsViewUserPII                    *bool `force:",omitempty"`
	PermissionsManageHubConnections           *bool `force:",omitempty"`
	PermissionsB2BMarketingAnalyticsUser      *bool `force:",omitempty"`
	PermissionsTraceXdsQueries                *bool `force:",omitempty"`
	PermissionsViewSecurityCommandCenter      *bool `force:",omitempty"`
	PermissionsManageSecurityCommandCenter    *bool `force:",omitempty"`
	PermissionsViewAllCustomSettings          *bool `force:",omitempty"`
	PermissionsViewAllForeignKeyNames         *bool `force:",omitempty"`
	PermissionsAllowSurveyAdvancedFeatures    *bool `force:",omitempty"`
	PermissionsAddWaveNotificationRecipients  *bool `force:",omitempty"`
	PermissionsHeadlessCMSAccess              *bool `force:",omitempty"`
	PermissionsLMEndMessagingSessionUserPerm  *bool `force:",omitempty"`
	PermissionsConsentApiUpdate               *bool `force:",omitempty"`
	PermissionsPaymentsAPIUser                *bool `force:",omitempty"`
	PermissionsAccessContentBuilder           *bool `force:",omitempty"`
	PermissionsAccountSwitcherUser            *bool `force:",omitempty"`
	PermissionsViewAnomalyEvents              *bool `force:",omitempty"`
	PermissionsManageC360AConnections         *bool `force:",omitempty"`
	PermissionsManageReleaseUpdates           *bool `force:",omitempty"`
	PermissionsViewAllProfiles                *bool `force:",omitempty"`
	PermissionsSkipIdentityConfirmation       *bool `force:",omitempty"`
	PermissionsLearningManager                *bool `force:",omitempty"`
	PermissionsSendCustomNotifications        *bool `force:",omitempty"`
	PermissionsPackaging2Delete               *bool `force:",omitempty"`
	PermissionsAutonomousAnalyticsPrivacy     *bool `force:",omitempty"`
	PermissionsSonicConsumer                  *bool `force:",omitempty"`
	PermissionsBotManageBots                  *bool `force:",omitempty"`
	PermissionsBotManageBotsTrainingData      *bool `force:",omitempty"`
	PermissionsManageLearningReporting        *bool `force:",omitempty"`
	PermissionsIsotopeCToCUser                *bool `force:",omitempty"`
	PermissionsIsotopeAccess                  *bool `force:",omitempty"`
	PermissionsIsotopeLEX                     *bool `force:",omitempty"`
	PermissionsGetSmartDataDiscoveryExternal  *bool `force:",omitempty"`
	PermissionsQuipMetricsAccess              *bool `force:",omitempty"`
	PermissionsQuipUserEngagementMetrics      *bool `force:",omitempty"`
	PermissionsManageExternalConnections      *bool `force:",omitempty"`
	PermissionsUseSubscriptionEmails          *bool `force:",omitempty"`
	PermissionsAIViewInsightObjects           *bool `force:",omitempty"`
	PermissionsAICreateInsightObjects         *bool `force:",omitempty"`
	PermissionsLifecycleManagementAPIUser     *bool `force:",omitempty"`
	PermissionsNativeWebviewScrolling         *bool `force:",omitempty"`
	PermissionsViewDeveloperName              *bool `force:",omitempty"`
}

func (profileResourceDataJSON) ApiName() string {
	return "Profile"
}

func (profileResourceDataJSON) ExternalIdApiName() string {
	return ""
}

func (profileResourceData) ApiName() string {
	return "Profile"
}

func (profileResourceData) ExternalIdApiName() string {
	return ""
}

func (p *profileResourceData) Instance() force.SObject {
	return p
}

func (p *profileResourceData) toJSONData() profileResourceDataJSON {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	var data profileResourceDataJSON
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *profileResourceData) Insertable() force.SObject {
	return p.toJSONData()
}

func (p *profileResourceData) Updatable() force.SObject {
	updatable := p.toJSONData()
	updatable.UserLicenseId = ""
	return updatable
}

func (p *profileResourceData) GetId() string {
	return p.Id.Value
}

func (p *profileResourceData) SetId(id string) {
	p.Id = types.String{Value: id}
}
