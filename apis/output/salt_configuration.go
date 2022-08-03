package models 
type SaltConfiguration struct {

	// Additional auth params that can be passed in for provisioning the salt minion.
	// Refer: https://docs.saltproject.io/en/master/topics/cloud/profiles.html
	AdditionalAuthParams map[string]string `json:"additionalAuthParams,omitempty"`

	// Additional configuration parameters for the salt minion, to be passed in as dictionary.
	// Refer: https://docs.saltproject.io/en/latest/ref/configuration/minion.html
	AdditionalMinionParams map[string]string `json:"additionalMinionParams,omitempty"`

	// Salt minion installer file name on the master.
	// This property is currently not being used by any SaltStack operation.
	InstallerFileName string `json:"installerFileName,omitempty"`

	// Salt master id to which the Salt minion will be connected to.
	MasterID string `json:"masterId,omitempty"`

	// Salt minion ID to be assigned to the deployed minion.
	MinionID string `json:"minionId,omitempty"`

	// Pillar environment to use when running state files.
	// Refer: https://docs.saltproject.io/en/latest/ref/modules/all/salt.modules.state.html
	PillarEnvironment string `json:"pillarEnvironment,omitempty"`

	// Salt environment to use when running state files.
	SaltEnvironment string `json:"saltEnvironment,omitempty"`

	// List of state files to run on the deployed minion.
	StateFiles []string `json:"stateFiles"`

	// Parameters required by the state file to run on the deployed minion.
	Variables map[string]string `json:"variables,omitempty"`
}

