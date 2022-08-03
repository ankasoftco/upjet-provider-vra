package models 
type CloudAccountGcpRegionEnumerationSpecification struct {

	// GCP Client email. Either provide clientEmail or provide a cloudAccountId of an existing account.
	// Example: 321743978432-compute@developer.gserviceaccount.com
	ClientEmail string `json:"clientEmail,omitempty"`

	// Existing cloud account id. Either provide id of existing account, or cloud account credentials: projectId, privateKeyId, privateKey and clientEmail.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// GCP Private key. Either provide privateKey or provide a cloudAccountId of an existing account.
	// Example: -----BEGIN PRIVATE KEY-----\nMIICXgIHAASBgSDHikastc8+I81zCg/qWW8dMr8mqvXQ3qbPAmu0RjxoZVI47tvs\nkYlFAXOf0sPrhO2nUuooJngnHV0639iTTEYG1vckNaW2R6U5QTdQ5Rq5u+uV3pMk\n7w7Vs4n3urQ4jnqt7rTXbC1DNa/PFeAZatbf7ffBBy0IGO0zc128IshYcwIDAQAB\nAoGBALTNl2JxTvq4SDW/3VH0fZkQXWH1MM10oeMbB2qO5beWb11FGaOO77nGKfWc\nbYgfp5Ogrql2yhBvLAXnxH8bcqqwORtFhlyV68U1y4R+8WxDNh0aevxH8hRS/1X5\n963DJm1JlU0E+vStiktN0tC3ebH5hE+1OxbIHSZ+WOWLYX7JAkEA5uigRgKp8ScG\nauUijvdOLZIhHWq9y5Wz+nOHUuDw8P7wOTKU34QJAoWEe771p9Pf/GTA/kr0BQnP\nQvWUDxGzJwJBAN05C6krwPeryFKrKtjOGJIbiIoY72wRnoNcdEEs3HDRhf48YWFo\nriRbZylzzzNFy/gmzT6XJQTfktGqq+FZD9UCQGIJaGrxHJgfmpDuAhMzGsUsYtTr\niRox0D1Iqa7dhE693t5aBG010OF6MLqdZA1CXrn5SRtuVVaCSLZEL/2J5UcCQQDA\nd3MXucNnN4NPuS/L9HMYJWD7lPoosaORcgyK77bSSNgk+u9WSjbH1uYIAIPSffUZ\nbti+jc2dUg5wb+aeZlgJAkEAurrpmpqj5vg087ZngKfFGR5rozDiTsK5DceTV97K\na1Y+Nzl+XWTxDBWk4YPh2ZlKv402hZEfWBYxUDn5ZkH/bw==\n-----END PRIVATE KEY-----\n
	PrivateKey string `json:"privateKey,omitempty"`

	// GCP Private key ID. Either provide privateKeyId or provide a cloudAccountId of an existing account.
	// Example: 027f73d50a19452eedf5775a9b42c5083678abdf
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// GCP Project ID. Either provide projectId or provide a cloudAccountId of an existing account.
	// Example: example-gcp-project
	ProjectID string `json:"projectId,omitempty"`
}

