/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Utilization report information of a tenant of Yarn application.
type ApiYarnTenantUtilization struct {
	// Name of the tenant.
	TenantName string `json:"tenantName,omitempty"`
	// Average number of VCores allocated to YARN applications of the tenant.
	AvgYarnCpuAllocation float64 `json:"avgYarnCpuAllocation,omitempty"`
	// Average number of VCores used by YARN applications of the tenant.
	AvgYarnCpuUtilization float64 `json:"avgYarnCpuUtilization,omitempty"`
	// Average unused VCores of the tenant.
	AvgYarnCpuUnusedCapacity float64 `json:"avgYarnCpuUnusedCapacity,omitempty"`
	// Average steady fair share VCores.
	AvgYarnCpuSteadyFairShare float64 `json:"avgYarnCpuSteadyFairShare,omitempty"`
	// Average allocated Vcores with pending containers.
	AvgYarnPoolAllocatedCpuDuringContention float64 `json:"avgYarnPoolAllocatedCpuDuringContention,omitempty"`
	// Average fair share VCores with pending containers.
	AvgYarnPoolFairShareCpuDuringContention float64 `json:"avgYarnPoolFairShareCpuDuringContention,omitempty"`
	// Average steady fair share VCores with pending containers.
	AvgYarnPoolSteadyFairShareCpuDuringContention float64 `json:"avgYarnPoolSteadyFairShareCpuDuringContention,omitempty"`
	// Average percentage of pending containers for the pool during periods of contention.
	AvgYarnContainerWaitRatio float64 `json:"avgYarnContainerWaitRatio,omitempty"`
	// Average memory allocated to YARN applications of the tenant.
	AvgYarnMemoryAllocation float64 `json:"avgYarnMemoryAllocation,omitempty"`
	// Average memory used by YARN applications of the tenant.
	AvgYarnMemoryUtilization float64 `json:"avgYarnMemoryUtilization,omitempty"`
	// Average unused memory of the tenant.
	AvgYarnMemoryUnusedCapacity float64 `json:"avgYarnMemoryUnusedCapacity,omitempty"`
	// Average steady fair share memory.
	AvgYarnMemorySteadyFairShare float64 `json:"avgYarnMemorySteadyFairShare,omitempty"`
	// Average allocated memory with pending containers.
	AvgYarnPoolAllocatedMemoryDuringContention float64 `json:"avgYarnPoolAllocatedMemoryDuringContention,omitempty"`
	// Average fair share memory with pending containers.
	AvgYarnPoolFairShareMemoryDuringContention float64 `json:"avgYarnPoolFairShareMemoryDuringContention,omitempty"`
	// Average steady fair share memory with pending containers.
	AvgYarnPoolSteadyFairShareMemoryDuringContention float64 `json:"avgYarnPoolSteadyFairShareMemoryDuringContention,omitempty"`
}
