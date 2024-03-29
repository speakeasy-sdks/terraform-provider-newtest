// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZoneVcenterConfig struct {
	EnableNetworkTypeSelection types.String                `tfsdk:"enable_network_type_selection"`
	APIURL                     types.String                `tfsdk:"api_url"`
	APIVersion                 types.String                `tfsdk:"api_version"`
	ApplianceURL               types.String                `tfsdk:"appliance_url"`
	BackupMode                 types.String                `tfsdk:"backup_mode"`
	CertificateProvider        types.String                `tfsdk:"certificate_provider"`
	Cluster                    types.String                `tfsdk:"cluster"`
	ConfigCmdbDiscovery        types.Bool                  `tfsdk:"config_cmdb_discovery"`
	ConfigCmdbID               types.String                `tfsdk:"config_cmdb_id"`
	ConfigCmID                 types.String                `tfsdk:"config_cm_id"`
	ConfigManagementID         types.String                `tfsdk:"config_management_id"`
	Datacenter                 types.String                `tfsdk:"datacenter"`
	DatacenterID               types.String                `tfsdk:"datacenter_id"`
	DatacenterName             types.String                `tfsdk:"datacenter_name"`
	DiskStorageType            types.String                `tfsdk:"disk_storage_type"`
	DistributedWorkerID        types.String                `tfsdk:"distributed_worker_id"`
	DNSIntegrationID           types.String                `tfsdk:"dns_integration_id"`
	EnableDiskTypeSelection    types.String                `tfsdk:"enable_disk_type_selection"`
	EnableVnc                  types.String                `tfsdk:"enable_vnc"`
	HideHostSelection          types.String                `tfsdk:"hide_host_selection"`
	ImportExisting             types.String                `tfsdk:"import_existing"`
	KubeURL                    types.String                `tfsdk:"kube_url"`
	NetworkServer              *ZoneAwsConfigNetworkServer `tfsdk:"network_server"`
	NetworkServerID            types.String                `tfsdk:"network_server_id"`
	Password                   types.String                `tfsdk:"password"`
	PasswordHash               types.String                `tfsdk:"password_hash"`
	ReplicationMode            types.String                `tfsdk:"replication_mode"`
	ResourcePool               types.String                `tfsdk:"resource_pool"`
	ResourcePoolID             types.String                `tfsdk:"resource_pool_id"`
	RPCMode                    types.String                `tfsdk:"rpc_mode"`
	SecurityMode               types.String                `tfsdk:"security_mode"`
	SecurityServer             types.String                `tfsdk:"security_server"`
	ServiceRegistryID          types.String                `tfsdk:"service_registry_id"`
	Username                   types.String                `tfsdk:"username"`
}
