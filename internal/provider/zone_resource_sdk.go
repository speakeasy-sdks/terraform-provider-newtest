// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/testing/terraform-provider-newtest/internal/sdk/pkg/models/operations"
	"github.com/testing/terraform-provider-newtest/internal/sdk/pkg/models/shared"
	"time"
)

func (r *ZoneResourceModel) ToSharedZoneCreate() *shared.ZoneCreate {
	accountID := new(int64)
	if !r.AccountID.IsUnknown() && !r.AccountID.IsNull() {
		*accountID = r.AccountID.ValueInt64()
	} else {
		accountID = nil
	}
	autoRecoverPowerState := new(bool)
	if !r.AutoRecoverPowerState.IsUnknown() && !r.AutoRecoverPowerState.IsNull() {
		*autoRecoverPowerState = r.AutoRecoverPowerState.ValueBool()
	} else {
		autoRecoverPowerState = nil
	}
	code := new(string)
	if !r.Code.IsUnknown() && !r.Code.IsNull() {
		*code = r.Code.ValueString()
	} else {
		code = nil
	}
	var config *shared.ZoneCreateConfig
	if r.Config != nil {
		config = &shared.ZoneCreateConfig{}
	}
	var credential *shared.ZoneCreateCredential
	if r.Credential != nil {
		id := new(int64)
		if !r.Credential.ID.IsUnknown() && !r.Credential.ID.IsNull() {
			*id = r.Credential.ID.ValueInt64()
		} else {
			id = nil
		}
		typeVar := new(string)
		if !r.Credential.Type.IsUnknown() && !r.Credential.Type.IsNull() {
			*typeVar = r.Credential.Type.ValueString()
		} else {
			typeVar = nil
		}
		credential = &shared.ZoneCreateCredential{
			ID:   id,
			Type: typeVar,
		}
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	groupID := r.GroupID.ValueInt64()
	linkedAccountID := new(int64)
	if !r.LinkedAccountID.IsUnknown() && !r.LinkedAccountID.IsNull() {
		*linkedAccountID = r.LinkedAccountID.ValueInt64()
	} else {
		linkedAccountID = nil
	}
	location := new(string)
	if !r.Location.IsUnknown() && !r.Location.IsNull() {
		*location = r.Location.ValueString()
	} else {
		location = nil
	}
	name := r.Name.ValueString()
	scalePriority := new(int64)
	if !r.ScalePriority.IsUnknown() && !r.ScalePriority.IsNull() {
		*scalePriority = r.ScalePriority.ValueInt64()
	} else {
		scalePriority = nil
	}
	securityMode := new(string)
	if !r.SecurityMode.IsUnknown() && !r.SecurityMode.IsNull() {
		*securityMode = r.SecurityMode.ValueString()
	} else {
		securityMode = nil
	}
	visibility := new(shared.Visibility)
	if !r.Visibility.IsUnknown() && !r.Visibility.IsNull() {
		*visibility = shared.Visibility(r.Visibility.ValueString())
	} else {
		visibility = nil
	}
	var zoneType shared.ZoneCreateZoneType
	var zoneCreate1 *shared.ZoneCreate1
	if r.ZoneType.One != nil {
		id1 := new(int64)
		if !r.ZoneType.One.ID.IsUnknown() && !r.ZoneType.One.ID.IsNull() {
			*id1 = r.ZoneType.One.ID.ValueInt64()
		} else {
			id1 = nil
		}
		zoneCreate1 = &shared.ZoneCreate1{
			ID: id1,
		}
	}
	if zoneCreate1 != nil {
		zoneType = shared.ZoneCreateZoneType{
			ZoneCreate1: zoneCreate1,
		}
	}
	var zoneCreate2 *shared.ZoneCreate2
	if r.ZoneType.Two != nil {
		code1 := new(string)
		if !r.ZoneType.Two.Code.IsUnknown() && !r.ZoneType.Two.Code.IsNull() {
			*code1 = r.ZoneType.Two.Code.ValueString()
		} else {
			code1 = nil
		}
		zoneCreate2 = &shared.ZoneCreate2{
			Code: code1,
		}
	}
	if zoneCreate2 != nil {
		zoneType = shared.ZoneCreateZoneType{
			ZoneCreate2: zoneCreate2,
		}
	}
	out := shared.ZoneCreate{
		AccountID:             accountID,
		AutoRecoverPowerState: autoRecoverPowerState,
		Code:                  code,
		Config:                config,
		Credential:            credential,
		Description:           description,
		Enabled:               enabled,
		GroupID:               groupID,
		LinkedAccountID:       linkedAccountID,
		Location:              location,
		Name:                  name,
		ScalePriority:         scalePriority,
		SecurityMode:          securityMode,
		Visibility:            visibility,
		ZoneType:              zoneType,
	}
	return &out
}

func (r *ZoneResourceModel) RefreshFromOperationsAddCloudsResponseBody(resp *operations.AddCloudsResponseBody) {
	if resp != nil {
		r.Success = types.BoolPointerValue(resp.Success)
		if resp.Zone == nil {
			r.Zone = nil
		} else {
			r.Zone = &Zone{}
			if resp.Zone.Account == nil {
				r.Zone.Account = nil
			} else {
				r.Zone.Account = &Account{}
				r.Zone.Account.ID = types.Int64PointerValue(resp.Zone.Account.ID)
				r.Zone.Account.Name = types.StringPointerValue(resp.Zone.Account.Name)
			}
			r.Zone.AccountID = types.Int64PointerValue(resp.Zone.AccountID)
			r.Zone.AgentMode = types.StringPointerValue(resp.Zone.AgentMode)
			r.Zone.APIProxy = types.StringPointerValue(resp.Zone.APIProxy)
			r.Zone.AutoRecoverPowerState = types.BoolPointerValue(resp.Zone.AutoRecoverPowerState)
			r.Zone.Code = types.StringPointerValue(resp.Zone.Code)
			if resp.Zone.Config == nil {
				r.Zone.Config = nil
			} else {
				r.Zone.Config = &Config{}
				if resp.Zone.Config.ZoneAwsConfig != nil {
					r.Zone.Config.ZoneAwsConfig = &ZoneAwsConfig{}
					r.Zone.Config.ZoneAwsConfig.UseHostCredentials = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.UseHostCredentials)
					r.Zone.Config.ZoneAwsConfig.AccessKey = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.AccessKey)
					r.Zone.Config.ZoneAwsConfig.ApplianceURL = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.ApplianceURL)
					r.Zone.Config.ZoneAwsConfig.BackupMode = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.BackupMode)
					r.Zone.Config.ZoneAwsConfig.CertificateProvider = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CertificateProvider)
					r.Zone.Config.ZoneAwsConfig.ConfigCmdbDiscovery = types.BoolPointerValue(resp.Zone.Config.ZoneAwsConfig.ConfigCmdbDiscovery)
					r.Zone.Config.ZoneAwsConfig.ConfigManagementID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.ConfigManagementID)
					r.Zone.Config.ZoneAwsConfig.CostingAccessKey = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingAccessKey)
					r.Zone.Config.ZoneAwsConfig.CostingBucket = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingBucket)
					r.Zone.Config.ZoneAwsConfig.CostingBucketName = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingBucketName)
					r.Zone.Config.ZoneAwsConfig.CostingFolder = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingFolder)
					r.Zone.Config.ZoneAwsConfig.CostingRegion = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingRegion)
					r.Zone.Config.ZoneAwsConfig.CostingReport = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingReport)
					r.Zone.Config.ZoneAwsConfig.CostingReportName = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingReportName)
					r.Zone.Config.ZoneAwsConfig.CostingSecretKey = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingSecretKey)
					r.Zone.Config.ZoneAwsConfig.CostingSecretKeyHash = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.CostingSecretKeyHash)
					r.Zone.Config.ZoneAwsConfig.DatacenterName = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.DatacenterName)
					r.Zone.Config.ZoneAwsConfig.DNSIntegrationID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.DNSIntegrationID)
					r.Zone.Config.ZoneAwsConfig.EbsEncryption = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.EbsEncryption)
					r.Zone.Config.ZoneAwsConfig.Endpoint = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.Endpoint)
					r.Zone.Config.ZoneAwsConfig.ImageStoreID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.ImageStoreID)
					r.Zone.Config.ZoneAwsConfig.IsVpc = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.IsVpc)
					if resp.Zone.Config.ZoneAwsConfig.NetworkServer == nil {
						r.Zone.Config.ZoneAwsConfig.NetworkServer = nil
					} else {
						r.Zone.Config.ZoneAwsConfig.NetworkServer = &ZoneAwsConfigNetworkServer{}
						r.Zone.Config.ZoneAwsConfig.NetworkServer.ID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.NetworkServer.ID)
					}
					r.Zone.Config.ZoneAwsConfig.NetworkServerID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.NetworkServerID)
					r.Zone.Config.ZoneAwsConfig.ReplicationMode = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.ReplicationMode)
					r.Zone.Config.ZoneAwsConfig.SecretKey = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.SecretKey)
					r.Zone.Config.ZoneAwsConfig.SecretKeyHash = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.SecretKeyHash)
					r.Zone.Config.ZoneAwsConfig.SecurityServer = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.SecurityServer)
					r.Zone.Config.ZoneAwsConfig.ServiceRegistryID = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.ServiceRegistryID)
					r.Zone.Config.ZoneAwsConfig.StsAssumeRole = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.StsAssumeRole)
					r.Zone.Config.ZoneAwsConfig.Vpc = types.StringPointerValue(resp.Zone.Config.ZoneAwsConfig.Vpc)
				}
				if resp.Zone.Config.ZoneAzureConfig != nil {
					r.Zone.Config.ZoneAzureConfig = &ZoneAzureConfig{}
					r.Zone.Config.ZoneAzureConfig.AccountType = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.AccountType)
					r.Zone.Config.ZoneAzureConfig.ApplianceURL = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ApplianceURL)
					r.Zone.Config.ZoneAzureConfig.AzureCostingMode = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.AzureCostingMode)
					r.Zone.Config.ZoneAzureConfig.BackupMode = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.BackupMode)
					r.Zone.Config.ZoneAzureConfig.CertificateProvider = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CertificateProvider)
					r.Zone.Config.ZoneAzureConfig.ClientID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ClientID)
					r.Zone.Config.ZoneAzureConfig.ClientSecret = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ClientSecret)
					r.Zone.Config.ZoneAzureConfig.ClientSecretHash = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ClientSecretHash)
					r.Zone.Config.ZoneAzureConfig.CloudType = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CloudType)
					r.Zone.Config.ZoneAzureConfig.ConfigCmdbDiscovery = types.BoolPointerValue(resp.Zone.Config.ZoneAzureConfig.ConfigCmdbDiscovery)
					r.Zone.Config.ZoneAzureConfig.ConfigCmdbID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ConfigCmdbID)
					r.Zone.Config.ZoneAzureConfig.ConfigManagementID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ConfigManagementID)
					r.Zone.Config.ZoneAzureConfig.CspClientID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CspClientID)
					r.Zone.Config.ZoneAzureConfig.CspClientSecret = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CspClientSecret)
					r.Zone.Config.ZoneAzureConfig.CspClientSecretHash = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CspClientSecretHash)
					r.Zone.Config.ZoneAzureConfig.CspCustomer = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CspCustomer)
					r.Zone.Config.ZoneAzureConfig.CspTenantID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.CspTenantID)
					r.Zone.Config.ZoneAzureConfig.DatacenterName = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.DatacenterName)
					r.Zone.Config.ZoneAzureConfig.DiskEncryption = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.DiskEncryption)
					r.Zone.Config.ZoneAzureConfig.DNSIntegrationID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.DNSIntegrationID)
					r.Zone.Config.ZoneAzureConfig.EncryptionSet = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.EncryptionSet)
					r.Zone.Config.ZoneAzureConfig.ImportExisting = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ImportExisting)
					r.Zone.Config.ZoneAzureConfig.InventoryLevel = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.InventoryLevel)
					if resp.Zone.Config.ZoneAzureConfig.NetworkServer == nil {
						r.Zone.Config.ZoneAzureConfig.NetworkServer = nil
					} else {
						r.Zone.Config.ZoneAzureConfig.NetworkServer = &ZoneAwsConfigNetworkServer{}
						r.Zone.Config.ZoneAzureConfig.NetworkServer.ID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.NetworkServer.ID)
					}
					r.Zone.Config.ZoneAzureConfig.NetworkServerID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.NetworkServerID)
					r.Zone.Config.ZoneAzureConfig.ReplicationMode = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ReplicationMode)
					r.Zone.Config.ZoneAzureConfig.ResourceGroup = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ResourceGroup)
					r.Zone.Config.ZoneAzureConfig.RPCMode = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.RPCMode)
					r.Zone.Config.ZoneAzureConfig.SecurityMode = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.SecurityMode)
					r.Zone.Config.ZoneAzureConfig.SecurityServer = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.SecurityServer)
					r.Zone.Config.ZoneAzureConfig.ServiceRegistryID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.ServiceRegistryID)
					r.Zone.Config.ZoneAzureConfig.SubscriberID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.SubscriberID)
					r.Zone.Config.ZoneAzureConfig.TenantID = types.StringPointerValue(resp.Zone.Config.ZoneAzureConfig.TenantID)
				}
				if resp.Zone.Config.ZoneGcpConfig != nil {
					r.Zone.Config.ZoneGcpConfig = &ZoneGcpConfig{}
					r.Zone.Config.ZoneGcpConfig.ApplianceURL = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ApplianceURL)
					r.Zone.Config.ZoneGcpConfig.BackupMode = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.BackupMode)
					r.Zone.Config.ZoneGcpConfig.CertificateProvider = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.CertificateProvider)
					r.Zone.Config.ZoneGcpConfig.ClientEmail = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ClientEmail)
					r.Zone.Config.ZoneGcpConfig.ConfigManagementID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ConfigManagementID)
					r.Zone.Config.ZoneGcpConfig.DatacenterName = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.DatacenterName)
					r.Zone.Config.ZoneGcpConfig.DNSIntegrationID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.DNSIntegrationID)
					r.Zone.Config.ZoneGcpConfig.GoogleRegionID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.GoogleRegionID)
					r.Zone.Config.ZoneGcpConfig.ImportExisting = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ImportExisting)
					if resp.Zone.Config.ZoneGcpConfig.NetworkServer == nil {
						r.Zone.Config.ZoneGcpConfig.NetworkServer = nil
					} else {
						r.Zone.Config.ZoneGcpConfig.NetworkServer = &ZoneAwsConfigNetworkServer{}
						r.Zone.Config.ZoneGcpConfig.NetworkServer.ID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.NetworkServer.ID)
					}
					r.Zone.Config.ZoneGcpConfig.NetworkServerID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.NetworkServerID)
					r.Zone.Config.ZoneGcpConfig.PrivateKey = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.PrivateKey)
					r.Zone.Config.ZoneGcpConfig.PrivateKeyHash = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.PrivateKeyHash)
					r.Zone.Config.ZoneGcpConfig.ProjectID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ProjectID)
					r.Zone.Config.ZoneGcpConfig.ReplicationMode = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ReplicationMode)
					r.Zone.Config.ZoneGcpConfig.SecurityServer = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.SecurityServer)
					r.Zone.Config.ZoneGcpConfig.ServiceRegistryID = types.StringPointerValue(resp.Zone.Config.ZoneGcpConfig.ServiceRegistryID)
				}
				if resp.Zone.Config.ZoneVcenterConfig != nil {
					r.Zone.Config.ZoneVcenterConfig = &ZoneVcenterConfig{}
					r.Zone.Config.ZoneVcenterConfig.EnableNetworkTypeSelection = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.EnableNetworkTypeSelection)
					r.Zone.Config.ZoneVcenterConfig.APIURL = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.APIURL)
					r.Zone.Config.ZoneVcenterConfig.APIVersion = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.APIVersion)
					r.Zone.Config.ZoneVcenterConfig.ApplianceURL = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ApplianceURL)
					r.Zone.Config.ZoneVcenterConfig.BackupMode = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.BackupMode)
					r.Zone.Config.ZoneVcenterConfig.CertificateProvider = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.CertificateProvider)
					r.Zone.Config.ZoneVcenterConfig.Cluster = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.Cluster)
					r.Zone.Config.ZoneVcenterConfig.ConfigCmdbDiscovery = types.BoolPointerValue(resp.Zone.Config.ZoneVcenterConfig.ConfigCmdbDiscovery)
					r.Zone.Config.ZoneVcenterConfig.ConfigCmdbID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ConfigCmdbID)
					r.Zone.Config.ZoneVcenterConfig.ConfigCmID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ConfigCmID)
					r.Zone.Config.ZoneVcenterConfig.ConfigManagementID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ConfigManagementID)
					r.Zone.Config.ZoneVcenterConfig.Datacenter = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.Datacenter)
					r.Zone.Config.ZoneVcenterConfig.DatacenterID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.DatacenterID)
					r.Zone.Config.ZoneVcenterConfig.DatacenterName = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.DatacenterName)
					r.Zone.Config.ZoneVcenterConfig.DiskStorageType = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.DiskStorageType)
					r.Zone.Config.ZoneVcenterConfig.DistributedWorkerID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.DistributedWorkerID)
					r.Zone.Config.ZoneVcenterConfig.DNSIntegrationID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.DNSIntegrationID)
					r.Zone.Config.ZoneVcenterConfig.EnableDiskTypeSelection = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.EnableDiskTypeSelection)
					r.Zone.Config.ZoneVcenterConfig.EnableVnc = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.EnableVnc)
					r.Zone.Config.ZoneVcenterConfig.HideHostSelection = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.HideHostSelection)
					r.Zone.Config.ZoneVcenterConfig.ImportExisting = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ImportExisting)
					r.Zone.Config.ZoneVcenterConfig.KubeURL = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.KubeURL)
					if resp.Zone.Config.ZoneVcenterConfig.NetworkServer == nil {
						r.Zone.Config.ZoneVcenterConfig.NetworkServer = nil
					} else {
						r.Zone.Config.ZoneVcenterConfig.NetworkServer = &ZoneAwsConfigNetworkServer{}
						r.Zone.Config.ZoneVcenterConfig.NetworkServer.ID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.NetworkServer.ID)
					}
					r.Zone.Config.ZoneVcenterConfig.NetworkServerID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.NetworkServerID)
					r.Zone.Config.ZoneVcenterConfig.Password = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.Password)
					r.Zone.Config.ZoneVcenterConfig.PasswordHash = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.PasswordHash)
					r.Zone.Config.ZoneVcenterConfig.ReplicationMode = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ReplicationMode)
					r.Zone.Config.ZoneVcenterConfig.ResourcePool = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ResourcePool)
					r.Zone.Config.ZoneVcenterConfig.ResourcePoolID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ResourcePoolID)
					r.Zone.Config.ZoneVcenterConfig.RPCMode = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.RPCMode)
					r.Zone.Config.ZoneVcenterConfig.SecurityMode = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.SecurityMode)
					r.Zone.Config.ZoneVcenterConfig.SecurityServer = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.SecurityServer)
					r.Zone.Config.ZoneVcenterConfig.ServiceRegistryID = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.ServiceRegistryID)
					r.Zone.Config.ZoneVcenterConfig.Username = types.StringPointerValue(resp.Zone.Config.ZoneVcenterConfig.Username)
				}
			}
			r.Zone.ConsoleKeymap = types.StringPointerValue(resp.Zone.ConsoleKeymap)
			r.Zone.ContainerMode = types.StringPointerValue(resp.Zone.ContainerMode)
			r.Zone.CostingMode = types.StringPointerValue(resp.Zone.CostingMode)
			if resp.Zone.CostLastSync != nil {
				r.Zone.CostLastSync = types.StringValue(resp.Zone.CostLastSync.Format(time.RFC3339Nano))
			} else {
				r.Zone.CostLastSync = types.StringNull()
			}
			r.Zone.CostLastSyncDuration = types.Int64PointerValue(resp.Zone.CostLastSyncDuration)
			r.Zone.CostStatus = types.StringPointerValue(resp.Zone.CostStatus)
			if resp.Zone.CostStatusDate != nil {
				r.Zone.CostStatusDate = types.StringValue(resp.Zone.CostStatusDate.Format(time.RFC3339Nano))
			} else {
				r.Zone.CostStatusDate = types.StringNull()
			}
			r.Zone.CostStatusMessage = types.StringPointerValue(resp.Zone.CostStatusMessage)
			if resp.Zone.Credential == nil {
				r.Zone.Credential = nil
			} else {
				r.Zone.Credential = &Credential{}
				if resp.Zone.Credential.One != nil {
					r.Zone.Credential.One = &One{}
					r.Zone.Credential.One.Type = types.StringPointerValue(resp.Zone.Credential.One.Type)
				}
				if resp.Zone.Credential.Two != nil {
					r.Zone.Credential.Two = &Two{}
					r.Zone.Credential.Two.ID = types.Int64PointerValue(resp.Zone.Credential.Two.ID)
					r.Zone.Credential.Two.Name = types.StringPointerValue(resp.Zone.Credential.Two.Name)
					r.Zone.Credential.Two.Type = types.StringPointerValue(resp.Zone.Credential.Two.Type)
				}
			}
			r.Zone.DarkImagePath = types.StringPointerValue(resp.Zone.DarkImagePath)
			if resp.Zone.DateCreated != nil {
				r.Zone.DateCreated = types.StringValue(resp.Zone.DateCreated.Format(time.RFC3339Nano))
			} else {
				r.Zone.DateCreated = types.StringNull()
			}
			r.Zone.DomainName = types.StringPointerValue(resp.Zone.DomainName)
			r.Zone.Enabled = types.BoolPointerValue(resp.Zone.Enabled)
			r.Zone.ExternalID = types.StringPointerValue(resp.Zone.ExternalID)
			if len(r.Zone.Groups) > len(resp.Zone.Groups) {
				r.Zone.Groups = r.Zone.Groups[:len(resp.Zone.Groups)]
			}
			for groupsCount, groupsItem := range resp.Zone.Groups {
				var groups1 Groups
				groups1.AccountID = types.Int64PointerValue(groupsItem.AccountID)
				groups1.ID = types.Int64PointerValue(groupsItem.ID)
				groups1.Name = types.StringPointerValue(groupsItem.Name)
				if groupsCount+1 > len(r.Zone.Groups) {
					r.Zone.Groups = append(r.Zone.Groups, groups1)
				} else {
					r.Zone.Groups[groupsCount].AccountID = groups1.AccountID
					r.Zone.Groups[groupsCount].ID = groups1.ID
					r.Zone.Groups[groupsCount].Name = groups1.Name
				}
			}
			r.Zone.GuidanceMode = types.StringPointerValue(resp.Zone.GuidanceMode)
			r.Zone.ID = types.Int64PointerValue(resp.Zone.ID)
			r.Zone.ImagePath = types.StringPointerValue(resp.Zone.ImagePath)
			r.Zone.InventoryLevel = types.StringPointerValue(resp.Zone.InventoryLevel)
			if resp.Zone.LastSync != nil {
				r.Zone.LastSync = types.StringValue(resp.Zone.LastSync.Format(time.RFC3339Nano))
			} else {
				r.Zone.LastSync = types.StringNull()
			}
			r.Zone.LastSyncDuration = types.Int64PointerValue(resp.Zone.LastSyncDuration)
			if resp.Zone.LastUpdated != nil {
				r.Zone.LastUpdated = types.StringValue(resp.Zone.LastUpdated.Format(time.RFC3339Nano))
			} else {
				r.Zone.LastUpdated = types.StringNull()
			}
			r.Zone.Location = types.StringPointerValue(resp.Zone.Location)
			r.Zone.Name = types.StringPointerValue(resp.Zone.Name)
			if resp.Zone.NetworkDomain == nil {
				r.Zone.NetworkDomain = nil
			} else {
				r.Zone.NetworkDomain = &Account{}
				r.Zone.NetworkDomain.ID = types.Int64PointerValue(resp.Zone.NetworkDomain.ID)
				r.Zone.NetworkDomain.Name = types.StringPointerValue(resp.Zone.NetworkDomain.Name)
			}
			if resp.Zone.NetworkServer == nil {
				r.Zone.NetworkServer = nil
			} else {
				r.Zone.NetworkServer = &Account{}
				r.Zone.NetworkServer.ID = types.Int64PointerValue(resp.Zone.NetworkServer.ID)
				r.Zone.NetworkServer.Name = types.StringPointerValue(resp.Zone.NetworkServer.Name)
			}
			if resp.Zone.NextRunDate != nil {
				r.Zone.NextRunDate = types.StringValue(resp.Zone.NextRunDate.Format(time.RFC3339Nano))
			} else {
				r.Zone.NextRunDate = types.StringNull()
			}
			if resp.Zone.Owner == nil {
				r.Zone.Owner = nil
			} else {
				r.Zone.Owner = &Account{}
				r.Zone.Owner.ID = types.Int64PointerValue(resp.Zone.Owner.ID)
				r.Zone.Owner.Name = types.StringPointerValue(resp.Zone.Owner.Name)
			}
			r.Zone.ProvisioningProxy = types.StringPointerValue(resp.Zone.ProvisioningProxy)
			r.Zone.RegionCode = types.StringPointerValue(resp.Zone.RegionCode)
			r.Zone.ScalePriority = types.Int64PointerValue(resp.Zone.ScalePriority)
			r.Zone.SecurityMode = types.StringPointerValue(resp.Zone.SecurityMode)
			if resp.Zone.SecurityServer == nil {
				r.Zone.SecurityServer = nil
			} else {
				r.Zone.SecurityServer = &Account{}
				r.Zone.SecurityServer.ID = types.Int64PointerValue(resp.Zone.SecurityServer.ID)
				r.Zone.SecurityServer.Name = types.StringPointerValue(resp.Zone.SecurityServer.Name)
			}
			r.Zone.ServerCount = types.Int64PointerValue(resp.Zone.ServerCount)
			r.Zone.ServiceVersion = types.StringPointerValue(resp.Zone.ServiceVersion)
			if resp.Zone.Stats == nil {
				r.Zone.Stats = nil
			} else {
				r.Zone.Stats = &Stats{}
				if resp.Zone.Stats.ServerCounts == nil {
					r.Zone.Stats.ServerCounts = nil
				} else {
					r.Zone.Stats.ServerCounts = &ServerCounts{}
					r.Zone.Stats.ServerCounts.All = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.All)
					r.Zone.Stats.ServerCounts.Baremetal = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.Baremetal)
					r.Zone.Stats.ServerCounts.ContainerHost = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.ContainerHost)
					r.Zone.Stats.ServerCounts.Host = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.Host)
					r.Zone.Stats.ServerCounts.Hypervisor = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.Hypervisor)
					r.Zone.Stats.ServerCounts.Unmanaged = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.Unmanaged)
					r.Zone.Stats.ServerCounts.VM = types.Int64PointerValue(resp.Zone.Stats.ServerCounts.VM)
				}
			}
			r.Zone.Status = types.StringPointerValue(resp.Zone.Status)
			if resp.Zone.StatusDate != nil {
				r.Zone.StatusDate = types.StringValue(resp.Zone.StatusDate.Format(time.RFC3339Nano))
			} else {
				r.Zone.StatusDate = types.StringNull()
			}
			r.Zone.StatusMessage = types.StringPointerValue(resp.Zone.StatusMessage)
			r.Zone.StorageMode = types.StringPointerValue(resp.Zone.StorageMode)
			r.Zone.Timezone = types.StringPointerValue(resp.Zone.Timezone)
			r.Zone.UserDataLinux = types.StringPointerValue(resp.Zone.UserDataLinux)
			r.Zone.UserDataWindows = types.StringPointerValue(resp.Zone.UserDataWindows)
			r.Zone.UUID = types.StringPointerValue(resp.Zone.UUID)
			r.Zone.Visibility = types.StringPointerValue(resp.Zone.Visibility)
			if resp.Zone.ZoneType == nil {
				r.Zone.ZoneType = nil
			} else {
				r.Zone.ZoneType = &ZoneType{}
				r.Zone.ZoneType.Code = types.StringPointerValue(resp.Zone.ZoneType.Code)
				r.Zone.ZoneType.ID = types.Int64PointerValue(resp.Zone.ZoneType.ID)
				r.Zone.ZoneType.Name = types.StringPointerValue(resp.Zone.ZoneType.Name)
			}
			r.Zone.ZoneTypeID = types.Int64PointerValue(resp.Zone.ZoneTypeID)
		}
	}
}
