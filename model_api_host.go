/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// This is the model for a host in the system.
type ApiHost struct {
	// A unique host identifier. This is not the same as the hostname (FQDN). It is a distinct value that remains the same even if the hostname changes.
	HostId string `json:"hostId,omitempty"`
	// The host IP address. This field is not mutable after the initial creation.
	IpAddress string `json:"ipAddress,omitempty"`
	// The hostname. This field is not mutable after the initial creation.
	Hostname string `json:"hostname,omitempty"`
	// The rack ID for this host.
	RackId string `json:"rackId,omitempty"`
	// Readonly. Requires \"full\" view. When the host agent sent the last heartbeat.
	LastHeartbeat string `json:"lastHeartbeat,omitempty"`
	// Readonly. Requires \"full\" view. The list of roles assigned to this host.
	RoleRefs      []ApiRoleRef      `json:"roleRefs,omitempty"`
	HealthSummary *ApiHealthSummary `json:"healthSummary,omitempty"`
	// Readonly. Requires \"full\" view. The list of health checks performed on the host, with their results.
	HealthChecks []ApiHealthCheck `json:"healthChecks,omitempty"`
	// Readonly. A URL into the Cloudera Manager web UI for this specific host.
	HostUrl string `json:"hostUrl,omitempty"`
	// Readonly. Whether the host is in maintenance mode. Available since API v2.
	MaintenanceMode bool                `json:"maintenanceMode,omitempty"`
	CommissionState *ApiCommissionState `json:"commissionState,omitempty"`
	// Readonly. The list of objects that trigger this host to be in maintenance mode. Available since API v2.
	MaintenanceOwners []ApiEntityType `json:"maintenanceOwners,omitempty"`
	Config            *ApiConfigList  `json:"config,omitempty"`
	// Readonly. The number of logical CPU cores on this host. Only populated after the host has heartbeated to the server. Available since API v4.
	NumCores float64 `json:"numCores,omitempty"`
	// Readonly. The number of physical CPU cores on this host. Only populated after the host has heartbeated to the server. Available since API v9.
	NumPhysicalCores float64 `json:"numPhysicalCores,omitempty"`
	// Readonly. The amount of physical RAM on this host, in bytes. Only populated after the host has heartbeated to the server. Available since API v4.
	TotalPhysMemBytes float64          `json:"totalPhysMemBytes,omitempty"`
	EntityStatus      *ApiEntityStatus `json:"entityStatus,omitempty"`
	ClusterRef        *ApiClusterRef   `json:"clusterRef,omitempty"`
}
