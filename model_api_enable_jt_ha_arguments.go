/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Arguments used for enable JT HA command.
type ApiEnableJtHaArguments struct {
	// Id of host on which second JobTracker role will be added.
	NewJtHostId string `json:"newJtHostId,omitempty"`
	// Initialize the ZNode even if it already exists. This can happen if JobTracker HA was enabled before and then disabled. Disable operation doesn't delete this ZNode. Defaults to true.
	ForceInitZNode bool `json:"forceInitZNode,omitempty"`
	// Name of the ZooKeeper service that will be used for auto-failover. This is an optional parameter if the MapReduce to ZooKeeper dependency is already set in CM.
	ZkServiceName string `json:"zkServiceName,omitempty"`
	// Name of the second JobTracker role to be created (Optional)
	NewJtRoleName string `json:"newJtRoleName,omitempty"`
	// Name of first Failover Controller role to be created. This is the Failover Controller co-located with the current JobTracker (Optional)
	Fc1RoleName string `json:"fc1RoleName,omitempty"`
	// Name of second Failover Controller role to be created. This is the Failover Controller co-located with the new JobTracker (Optional)
	Fc2RoleName string `json:"fc2RoleName,omitempty"`
	// Logical name of the JobTracker pair. If value is not provided, \"logicaljt\" is used as the default. The name can contain only alphanumeric characters and \"-\". <p> Available since API v8.
	LogicalName string `json:"logicalName,omitempty"`
}
