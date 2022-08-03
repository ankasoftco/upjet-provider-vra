package models 
type CloudAccountVmcRegionEnumerationSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// VMC API access key. Either provide apiKey or provide a cloudAccountId of an existing account.
	APIKey string `json:"apiKey,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or apiKey, sddcId, username, password, hostName, nsxHostName.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// The host name of the CSP service.
	// Example: console-stg.cloud.vmware.com
	CspHostName string `json:"cspHostName,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	DcID string `json:"dcId,omitempty"`

	// Enter the IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.  Either provide hostName or provide a cloudAccountId of an existing account.
	// Example: vc1.vmware.com
	HostName string `json:"hostName,omitempty"`

	// The IP address of the NSX Manager server in the specified SDDC / FQDN.Either provide nsxHostName or provide a cloudAccountId of an existing account.
	// Example: nsxManager.sddc-52-12-8-145.vmwaretest.com
	NsxHostName string `json:"nsxHostName,omitempty"`

	// Password for the user used to authenticate with the cloud Account. Either provide password or provide a cloudAccountId of an existing account.
	// Example: cndhjslacd90ascdbasyoucbdh
	Password string `json:"password,omitempty"`

	// Identifier of the on-premise SDDC to be used by this cloud account. Note that NSX-V SDDCs are not supported. Either provide sddcId or provide a cloudAccountId of an existing account.
	// Example: CMBU-PRD-NSXT-M8GA-090319
	SddcID string `json:"sddcId,omitempty"`

	// vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.Either provide username or provide a cloudAccountId of an existing account.
	// Example: administrator@mycompany.com
	Username string `json:"username,omitempty"`
}

