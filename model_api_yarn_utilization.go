/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Utilization report information of a Yarn application service.
type ApiYarnUtilization struct {
	// Average number of VCores used by YARN applications during the report window.
	AvgCpuUtilization float64 `json:"avgCpuUtilization,omitempty"`
	// Maximum number of VCores used by YARN applications during the report window.
	MaxCpuUtilization float64 `json:"maxCpuUtilization,omitempty"`
	// Average daily peak VCores used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean.
	AvgCpuDailyPeak float64 `json:"avgCpuDailyPeak,omitempty"`
	// Timestamp corresponds to maximum number of VCores used by YARN applications during the report window.
	MaxCpuUtilizationTimestampMs float64 `json:"maxCpuUtilizationTimestampMs,omitempty"`
	// Average percentage of VCores used by YARN applications during the report window.
	AvgCpuUtilizationPercentage float64 `json:"avgCpuUtilizationPercentage,omitempty"`
	// Maximum percentage of VCores used by YARN applications during the report window.
	MaxCpuUtilizationPercentage float64 `json:"maxCpuUtilizationPercentage,omitempty"`
	// Average daily peak percentage of VCores used by YARN applications during the report window.
	AvgCpuDailyPeakPercentage float64 `json:"avgCpuDailyPeakPercentage,omitempty"`
	// Average memory used by YARN applications during the report window.
	AvgMemoryUtilization float64 `json:"avgMemoryUtilization,omitempty"`
	// Maximum memory used by YARN applications during the report window.
	MaxMemoryUtilization float64 `json:"maxMemoryUtilization,omitempty"`
	// Average daily peak memory used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean.
	AvgMemoryDailyPeak float64 `json:"avgMemoryDailyPeak,omitempty"`
	// Timestamp corresponds to maximum memory used by YARN applications during the report window.
	MaxMemoryUtilizationTimestampMs float64 `json:"maxMemoryUtilizationTimestampMs,omitempty"`
	// Average percentage memory used by YARN applications during the report window.
	AvgMemoryUtilizationPercentage float64 `json:"avgMemoryUtilizationPercentage,omitempty"`
	// Maximum percentage of memory used by YARN applications during the report window.
	MaxMemoryUtilizationPercentage float64 `json:"maxMemoryUtilizationPercentage,omitempty"`
	// Average daily peak percentage of memory used by YARN applications during the report window.
	AvgMemoryDailyPeakPercentage float64                       `json:"avgMemoryDailyPeakPercentage,omitempty"`
	TenantUtilizations           *ApiYarnTenantUtilizationList `json:"tenantUtilizations,omitempty"`
	// error message of utilization report.
	ErrorMessage string `json:"errorMessage,omitempty"`
}
