// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ZoneGcpConfigNetworkServer struct {
	ID *string `json:"id,omitempty"`
}

type ZoneAzureConfig struct {
	AccountType         *string                       `json:"accountType,omitempty"`
	ApplianceURL        *string                       `json:"applianceUrl,omitempty"`
	AzureCostingMode    *string                       `json:"azureCostingMode,omitempty"`
	BackupMode          *string                       `json:"backupMode,omitempty"`
	CertificateProvider *string                       `json:"certificateProvider,omitempty"`
	ClientID            *string                       `json:"clientId,omitempty"`
	ClientSecret        *string                       `json:"clientSecret,omitempty"`
	ClientSecretHash    *string                       `json:"clientSecretHash,omitempty"`
	CloudType           *string                       `json:"cloudType,omitempty"`
	ConfigCmdbDiscovery *bool                         `json:"configCmdbDiscovery,omitempty"`
	ConfigCmdbID        *string                       `json:"configCmdbId,omitempty"`
	ConfigManagementID  *string                       `json:"configManagementId,omitempty"`
	CspClientID         *string                       `json:"cspClientId,omitempty"`
	CspClientSecret     *string                       `json:"cspClientSecret,omitempty"`
	CspClientSecretHash *string                       `json:"cspClientSecretHash,omitempty"`
	CspCustomer         *string                       `json:"cspCustomer,omitempty"`
	CspTenantID         *string                       `json:"cspTenantId,omitempty"`
	DatacenterName      *string                       `json:"datacenterName,omitempty"`
	DiskEncryption      *string                       `json:"diskEncryption,omitempty"`
	DNSIntegrationID    *string                       `json:"dnsIntegrationId,omitempty"`
	EncryptionSet       *string                       `json:"encryptionSet,omitempty"`
	ImportExisting      *string                       `json:"importExisting,omitempty"`
	InventoryLevel      *string                       `json:"inventoryLevel,omitempty"`
	NetworkServer       *ZoneAzureConfigNetworkServer `json:"networkServer,omitempty"`
	NetworkServerID     *string                       `json:"networkServer.id,omitempty"`
	ReplicationMode     *string                       `json:"replicationMode,omitempty"`
	ResourceGroup       *string                       `json:"resourceGroup,omitempty"`
	RPCMode             *string                       `json:"rpcMode,omitempty"`
	SecurityMode        *string                       `json:"securityMode,omitempty"`
	SecurityServer      *string                       `json:"securityServer,omitempty"`
	ServiceRegistryID   *string                       `json:"serviceRegistryId,omitempty"`
	SubscriberID        *string                       `json:"subscriberId,omitempty"`
	TenantID            *string                       `json:"tenantId,omitempty"`
}
