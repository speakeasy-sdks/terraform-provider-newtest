// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/testing/terraform-provider-newtest/internal/provider/types"
	"github.com/testing/terraform-provider-newtest/internal/sdk"
	"github.com/testing/terraform-provider-newtest/internal/sdk/models/operations"
	"github.com/testing/terraform-provider-newtest/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneResource{}
var _ resource.ResourceWithImportState = &ZoneResource{}

func NewZoneResource() resource.Resource {
	return &ZoneResource{}
}

// ZoneResource defines the resource implementation.
type ZoneResource struct {
	client *sdk.Newtest
}

// ZoneResourceModel describes the resource data model.
type ZoneResourceModel struct {
	AccountID             types.Int64                   `tfsdk:"account_id"`
	AutoRecoverPowerState types.Bool                    `tfsdk:"auto_recover_power_state"`
	Code                  types.String                  `tfsdk:"code"`
	Config                *tfTypes.ZoneCreateConfig     `tfsdk:"config"`
	Credential            *tfTypes.ZoneCreateCredential `tfsdk:"credential"`
	Description           types.String                  `tfsdk:"description"`
	Enabled               types.Bool                    `tfsdk:"enabled"`
	GroupID               types.Int64                   `tfsdk:"group_id"`
	LinkedAccountID       types.Int64                   `tfsdk:"linked_account_id"`
	Location              types.String                  `tfsdk:"location"`
	Name                  types.String                  `tfsdk:"name"`
	ScalePriority         types.Int64                   `tfsdk:"scale_priority"`
	SecurityMode          types.String                  `tfsdk:"security_mode"`
	Success               types.Bool                    `tfsdk:"success"`
	Visibility            types.String                  `tfsdk:"visibility"`
	Zone                  *tfTypes.Zone                 `tfsdk:"zone"`
	ZoneType              tfTypes.ZoneCreateZoneType    `tfsdk:"zone_type"`
}

func (r *ZoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *ZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Zone Resource",

		Attributes: map[string]schema.Attribute{
			"account_id": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Specifies which Tenant this cloud should be assigned to. Requires replacement if changed. `,
			},
			"auto_recover_power_state": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `Automatically Power on VMs. Requires replacement if changed. ; Default: false`,
			},
			"code": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional code for use with policies. Requires replacement if changed. `,
			},
			"config": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `Map containing zone configuration settings. See the section on specific zone types for details. Requires replacement if changed. `,
			},
			"credential": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
					},
					"type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Default:     stringdefault.StaticString("local"),
						Description: `Requires replacement if changed. ; Default: "local"`,
					},
				},
				Description: `Map containing Credential ID. Setting ` + "`" + `type` + "`" + ` to ` + "`" + `local` + "`" + ` means use the values set in the local cloud config instead of associating a credential. Requires replacement if changed. `,
			},
			"description": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional description field if you want to put more info there. Requires replacement if changed. `,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Can be used to disable the cloud. Requires replacement if changed. ; Default: true`,
			},
			"group_id": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Specifies which Server group this cloud should be assigned to. Requires replacement if changed. `,
			},
			"linked_account_id": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Linked Account ID (enter commercial ID to get costing for AWS Govcloud). Requires replacement if changed. `,
			},
			"location": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional location for your cloud. Requires replacement if changed. `,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `A unique name scoped to your account for the cloud. Requires replacement if changed. `,
			},
			"scale_priority": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: `Scale Priority. Requires replacement if changed. ; Default: 1`,
			},
			"security_mode": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("off"),
				Description: `host firewall. ` + "`" + `off` + "`" + ` or ` + "`" + `internal` + "`" + `. a.k.a. "local firewall". Requires replacement if changed. ; Default: "off"`,
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
			"visibility": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("private"),
				Description: `private or public. Requires replacement if changed. ; must be one of ["private", "public"]; Default: "private"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"private",
						"public",
					),
				},
			},
			"zone": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"account": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"account_id": schema.Int64Attribute{
						Computed: true,
					},
					"agent_mode": schema.StringAttribute{
						Computed: true,
					},
					"api_proxy": schema.StringAttribute{
						Computed: true,
					},
					"auto_recover_power_state": schema.BoolAttribute{
						Computed: true,
					},
					"code": schema.StringAttribute{
						Computed: true,
					},
					"config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"zone_aws_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"use_host_credentials": schema.StringAttribute{
										Computed: true,
									},
									"access_key": schema.StringAttribute{
										Computed: true,
									},
									"appliance_url": schema.StringAttribute{
										Computed: true,
									},
									"backup_mode": schema.StringAttribute{
										Computed: true,
									},
									"certificate_provider": schema.StringAttribute{
										Computed: true,
									},
									"config_cmdb_discovery": schema.BoolAttribute{
										Computed: true,
									},
									"config_management_id": schema.StringAttribute{
										Computed: true,
									},
									"costing_access_key": schema.StringAttribute{
										Computed: true,
									},
									"costing_bucket": schema.StringAttribute{
										Computed: true,
									},
									"costing_bucket_name": schema.StringAttribute{
										Computed: true,
									},
									"costing_folder": schema.StringAttribute{
										Computed: true,
									},
									"costing_region": schema.StringAttribute{
										Computed: true,
									},
									"costing_report": schema.StringAttribute{
										Computed: true,
									},
									"costing_report_name": schema.StringAttribute{
										Computed: true,
									},
									"costing_secret_key": schema.StringAttribute{
										Computed: true,
									},
									"costing_secret_key_hash": schema.StringAttribute{
										Computed: true,
									},
									"datacenter_name": schema.StringAttribute{
										Computed: true,
									},
									"dns_integration_id": schema.StringAttribute{
										Computed: true,
									},
									"ebs_encryption": schema.StringAttribute{
										Computed: true,
									},
									"endpoint": schema.StringAttribute{
										Computed: true,
									},
									"image_store_id": schema.StringAttribute{
										Computed: true,
									},
									"is_vpc": schema.StringAttribute{
										Computed: true,
									},
									"network_server": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"network_server_id": schema.StringAttribute{
										Computed: true,
									},
									"replication_mode": schema.StringAttribute{
										Computed: true,
									},
									"secret_key": schema.StringAttribute{
										Computed: true,
									},
									"secret_key_hash": schema.StringAttribute{
										Computed: true,
									},
									"security_server": schema.StringAttribute{
										Computed: true,
									},
									"service_registry_id": schema.StringAttribute{
										Computed: true,
									},
									"sts_assume_role": schema.StringAttribute{
										Computed: true,
									},
									"vpc": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"zone_azure_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"account_type": schema.StringAttribute{
										Computed: true,
									},
									"appliance_url": schema.StringAttribute{
										Computed: true,
									},
									"azure_costing_mode": schema.StringAttribute{
										Computed: true,
									},
									"backup_mode": schema.StringAttribute{
										Computed: true,
									},
									"certificate_provider": schema.StringAttribute{
										Computed: true,
									},
									"client_id": schema.StringAttribute{
										Computed: true,
									},
									"client_secret": schema.StringAttribute{
										Computed: true,
									},
									"client_secret_hash": schema.StringAttribute{
										Computed: true,
									},
									"cloud_type": schema.StringAttribute{
										Computed: true,
									},
									"config_cmdb_discovery": schema.BoolAttribute{
										Computed: true,
									},
									"config_cmdb_id": schema.StringAttribute{
										Computed: true,
									},
									"config_management_id": schema.StringAttribute{
										Computed: true,
									},
									"csp_client_id": schema.StringAttribute{
										Computed: true,
									},
									"csp_client_secret": schema.StringAttribute{
										Computed: true,
									},
									"csp_client_secret_hash": schema.StringAttribute{
										Computed: true,
									},
									"csp_customer": schema.StringAttribute{
										Computed: true,
									},
									"csp_tenant_id": schema.StringAttribute{
										Computed: true,
									},
									"datacenter_name": schema.StringAttribute{
										Computed: true,
									},
									"disk_encryption": schema.StringAttribute{
										Computed: true,
									},
									"dns_integration_id": schema.StringAttribute{
										Computed: true,
									},
									"encryption_set": schema.StringAttribute{
										Computed: true,
									},
									"import_existing": schema.StringAttribute{
										Computed: true,
									},
									"inventory_level": schema.StringAttribute{
										Computed: true,
									},
									"network_server": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"network_server_id": schema.StringAttribute{
										Computed: true,
									},
									"replication_mode": schema.StringAttribute{
										Computed: true,
									},
									"resource_group": schema.StringAttribute{
										Computed: true,
									},
									"rpc_mode": schema.StringAttribute{
										Computed: true,
									},
									"security_mode": schema.StringAttribute{
										Computed: true,
									},
									"security_server": schema.StringAttribute{
										Computed: true,
									},
									"service_registry_id": schema.StringAttribute{
										Computed: true,
									},
									"subscriber_id": schema.StringAttribute{
										Computed: true,
									},
									"tenant_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"zone_gcp_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"appliance_url": schema.StringAttribute{
										Computed: true,
									},
									"backup_mode": schema.StringAttribute{
										Computed: true,
									},
									"certificate_provider": schema.StringAttribute{
										Computed: true,
									},
									"client_email": schema.StringAttribute{
										Computed: true,
									},
									"config_management_id": schema.StringAttribute{
										Computed: true,
									},
									"datacenter_name": schema.StringAttribute{
										Computed: true,
									},
									"dns_integration_id": schema.StringAttribute{
										Computed: true,
									},
									"google_region_id": schema.StringAttribute{
										Computed: true,
									},
									"import_existing": schema.StringAttribute{
										Computed: true,
									},
									"network_server": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"network_server_id": schema.StringAttribute{
										Computed: true,
									},
									"private_key": schema.StringAttribute{
										Computed: true,
									},
									"private_key_hash": schema.StringAttribute{
										Computed: true,
									},
									"project_id": schema.StringAttribute{
										Computed: true,
									},
									"replication_mode": schema.StringAttribute{
										Computed: true,
									},
									"security_server": schema.StringAttribute{
										Computed: true,
									},
									"service_registry_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"zone_vcenter_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"enable_network_type_selection": schema.StringAttribute{
										Computed: true,
									},
									"api_url": schema.StringAttribute{
										Computed: true,
									},
									"api_version": schema.StringAttribute{
										Computed: true,
									},
									"appliance_url": schema.StringAttribute{
										Computed: true,
									},
									"backup_mode": schema.StringAttribute{
										Computed: true,
									},
									"certificate_provider": schema.StringAttribute{
										Computed: true,
									},
									"cluster": schema.StringAttribute{
										Computed: true,
									},
									"config_cmdb_discovery": schema.BoolAttribute{
										Computed: true,
									},
									"config_cmdb_id": schema.StringAttribute{
										Computed: true,
									},
									"config_cm_id": schema.StringAttribute{
										Computed: true,
									},
									"config_management_id": schema.StringAttribute{
										Computed: true,
									},
									"datacenter": schema.StringAttribute{
										Computed: true,
									},
									"datacenter_id": schema.StringAttribute{
										Computed: true,
									},
									"datacenter_name": schema.StringAttribute{
										Computed: true,
									},
									"disk_storage_type": schema.StringAttribute{
										Computed: true,
									},
									"distributed_worker_id": schema.StringAttribute{
										Computed: true,
									},
									"dns_integration_id": schema.StringAttribute{
										Computed: true,
									},
									"enable_disk_type_selection": schema.StringAttribute{
										Computed: true,
									},
									"enable_vnc": schema.StringAttribute{
										Computed: true,
									},
									"hide_host_selection": schema.StringAttribute{
										Computed: true,
									},
									"import_existing": schema.StringAttribute{
										Computed: true,
									},
									"kube_url": schema.StringAttribute{
										Computed: true,
									},
									"network_server": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"network_server_id": schema.StringAttribute{
										Computed: true,
									},
									"password": schema.StringAttribute{
										Computed: true,
									},
									"password_hash": schema.StringAttribute{
										Computed: true,
									},
									"replication_mode": schema.StringAttribute{
										Computed: true,
									},
									"resource_pool": schema.StringAttribute{
										Computed: true,
									},
									"resource_pool_id": schema.StringAttribute{
										Computed: true,
									},
									"rpc_mode": schema.StringAttribute{
										Computed: true,
									},
									"security_mode": schema.StringAttribute{
										Computed: true,
									},
									"security_server": schema.StringAttribute{
										Computed: true,
									},
									"service_registry_id": schema.StringAttribute{
										Computed: true,
									},
									"username": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"console_keymap": schema.StringAttribute{
						Computed: true,
					},
					"container_mode": schema.StringAttribute{
						Computed: true,
					},
					"costing_mode": schema.StringAttribute{
						Computed: true,
					},
					"cost_last_sync": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"cost_last_sync_duration": schema.Int64Attribute{
						Computed: true,
					},
					"cost_status": schema.StringAttribute{
						Computed: true,
					},
					"cost_status_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"cost_status_message": schema.StringAttribute{
						Computed: true,
					},
					"credential": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"one": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"two": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"type": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"dark_image_path": schema.StringAttribute{
						Computed:    true,
						Description: `Dark logo image URL`,
					},
					"date_created": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"domain_name": schema.StringAttribute{
						Computed: true,
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"external_id": schema.StringAttribute{
						Computed: true,
					},
					"groups": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"account_id": schema.Int64Attribute{
									Computed: true,
								},
								"id": schema.Int64Attribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"guidance_mode": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.Int64Attribute{
						Computed: true,
					},
					"image_path": schema.StringAttribute{
						Computed:    true,
						Description: `Logo image URL`,
					},
					"inventory_level": schema.StringAttribute{
						Computed: true,
					},
					"last_sync": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"last_sync_duration": schema.Int64Attribute{
						Computed: true,
					},
					"last_updated": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"location": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_domain": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"network_server": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"next_run_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"owner": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"provisioning_proxy": schema.StringAttribute{
						Computed: true,
					},
					"region_code": schema.StringAttribute{
						Computed: true,
					},
					"scale_priority": schema.Int64Attribute{
						Computed: true,
					},
					"security_mode": schema.StringAttribute{
						Computed: true,
					},
					"security_server": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"server_count": schema.Int64Attribute{
						Computed: true,
					},
					"service_version": schema.StringAttribute{
						Computed: true,
					},
					"stats": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"server_counts": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"all": schema.Int64Attribute{
										Computed: true,
									},
									"baremetal": schema.Int64Attribute{
										Computed: true,
									},
									"container_host": schema.Int64Attribute{
										Computed: true,
									},
									"host": schema.Int64Attribute{
										Computed: true,
									},
									"hypervisor": schema.Int64Attribute{
										Computed: true,
									},
									"unmanaged": schema.Int64Attribute{
										Computed: true,
									},
									"vm": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
					"status_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"status_message": schema.StringAttribute{
						Computed: true,
					},
					"storage_mode": schema.StringAttribute{
						Computed: true,
					},
					"timezone": schema.StringAttribute{
						Computed: true,
					},
					"user_data_linux": schema.StringAttribute{
						Computed: true,
					},
					"user_data_windows": schema.StringAttribute{
						Computed: true,
					},
					"uuid": schema.StringAttribute{
						Computed: true,
					},
					"visibility": schema.StringAttribute{
						Computed: true,
					},
					"zone_type": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"code": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"zone_type_id": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"zone_type": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"one": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.RequiresReplaceIfConfigured(),
								},
								Optional:    true,
								Description: `Requires replacement if changed. `,
							},
						},
						Description: `Map containing the Cloud (zone) type ID. See the zone-types API to fetch a list of all available Cloud (zone) types and their IDs. Requires replacement if changed. `,
					},
					"two": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"code": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Optional:    true,
								Description: `Requires replacement if changed. `,
							},
						},
						Description: `Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes. Requires replacement if changed. `,
					},
				},
				Description: `Requires replacement if changed. `,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
		},
	}
}

func (r *ZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Newtest)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Newtest, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZoneResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var request *operations.AddCloudsRequestBody
	zone := *data.ToSharedZoneCreate()
	request = &operations.AddCloudsRequestBody{
		Zone: zone,
	}
	res, err := r.client.Clouds.AddClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Object == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsAddCloudsResponseBody(res.Object)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZoneResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *ZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource zone.")
}
