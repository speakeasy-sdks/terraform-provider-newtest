// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZoneAzureConfig struct {
	AccountType         types.String                `tfsdk:"account_type"`
	ApplianceURL        types.String                `tfsdk:"appliance_url"`
	AzureCostingMode    types.String                `tfsdk:"azure_costing_mode"`
	BackupMode          types.String                `tfsdk:"backup_mode"`
	CertificateProvider types.String                `tfsdk:"certificate_provider"`
	ClientID            types.String                `tfsdk:"client_id"`
	ClientSecret        types.String                `tfsdk:"client_secret"`
	ClientSecretHash    types.String                `tfsdk:"client_secret_hash"`
	CloudType           types.String                `tfsdk:"cloud_type"`
	ConfigCmdbDiscovery types.Bool                  `tfsdk:"config_cmdb_discovery"`
	ConfigCmdbID        types.String                `tfsdk:"config_cmdb_id"`
	ConfigManagementID  types.String                `tfsdk:"config_management_id"`
	CspClientID         types.String                `tfsdk:"csp_client_id"`
	CspClientSecret     types.String                `tfsdk:"csp_client_secret"`
	CspClientSecretHash types.String                `tfsdk:"csp_client_secret_hash"`
	CspCustomer         types.String                `tfsdk:"csp_customer"`
	CspTenantID         types.String                `tfsdk:"csp_tenant_id"`
	DatacenterName      types.String                `tfsdk:"datacenter_name"`
	DiskEncryption      types.String                `tfsdk:"disk_encryption"`
	DNSIntegrationID    types.String                `tfsdk:"dns_integration_id"`
	EncryptionSet       types.String                `tfsdk:"encryption_set"`
	ImportExisting      types.String                `tfsdk:"import_existing"`
	InventoryLevel      types.String                `tfsdk:"inventory_level"`
	NetworkServer       *ZoneAwsConfigNetworkServer `tfsdk:"network_server"`
	NetworkServerID     types.String                `tfsdk:"network_server_id"`
	ReplicationMode     types.String                `tfsdk:"replication_mode"`
	ResourceGroup       types.String                `tfsdk:"resource_group"`
	RPCMode             types.String                `tfsdk:"rpc_mode"`
	SecurityMode        types.String                `tfsdk:"security_mode"`
	SecurityServer      types.String                `tfsdk:"security_server"`
	ServiceRegistryID   types.String                `tfsdk:"service_registry_id"`
	SubscriberID        types.String                `tfsdk:"subscriber_id"`
	TenantID            types.String                `tfsdk:"tenant_id"`
}
