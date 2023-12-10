/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A role represents a specific entity that participate in a service. Examples are JobTrackers, DataNodes, HBase Masters. Each role is assigned a host where it runs on.
type ApiRole struct {
	// The name of the role. Optional when creating a role since API v6. If not specified, a name will be automatically generated for the role.
	Name string `json:"name,omitempty"`
	// The type of the role, e.g. NAMENODE, DATANODE, TASKTRACKER.
	Type_           string              `json:"type,omitempty"`
	HostRef         *ApiHostRef         `json:"hostRef,omitempty"`
	ServiceRef      *ApiServiceRef      `json:"serviceRef,omitempty"`
	RoleState       *ApiRoleState       `json:"roleState,omitempty"`
	CommissionState *ApiCommissionState `json:"commissionState,omitempty"`
	HealthSummary   *ApiHealthSummary   `json:"healthSummary,omitempty"`
	// Readonly. Expresses whether the role configuration is stale.
	ConfigStale           bool                      `json:"configStale,omitempty"`
	ConfigStalenessStatus *ApiConfigStalenessStatus `json:"configStalenessStatus,omitempty"`
	// Readonly. The list of health checks of this service.
	HealthChecks []ApiHealthCheck `json:"healthChecks,omitempty"`
	HaStatus     *HaStatus        `json:"haStatus,omitempty"`
	// Readonly. Link into the Cloudera Manager web UI for this specific role.
	RoleUrl string `json:"roleUrl,omitempty"`
	// Readonly. Whether the role is in maintenance mode. Available since API v2.
	MaintenanceMode bool `json:"maintenanceMode,omitempty"`
	// Readonly. The list of objects that trigger this role to be in maintenance mode. Available since API v2.
	MaintenanceOwners   []ApiEntityType        `json:"maintenanceOwners,omitempty"`
	Config              *ApiConfigList         `json:"config,omitempty"`
	RoleConfigGroupRef  *ApiRoleConfigGroupRef `json:"roleConfigGroupRef,omitempty"`
	ZooKeeperServerMode *ZooKeeperServerMode   `json:"zooKeeperServerMode,omitempty"`
	EntityStatus        *ApiEntityStatus       `json:"entityStatus,omitempty"`
}