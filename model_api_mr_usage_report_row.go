/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

type ApiMrUsageReportRow struct {
	// The time period over which this report is generated.
	TimePeriod string `json:"timePeriod,omitempty"`
	// The user being reported.
	User string `json:"user,omitempty"`
	// The group this user belongs to.
	Group string `json:"group,omitempty"`
	// Amount of CPU time (in seconds) taken up this user's MapReduce jobs.
	CpuSec float64 `json:"cpuSec,omitempty"`
	// The sum of physical memory used (collected as a snapshot) by this user's MapReduce jobs.
	MemoryBytes float64 `json:"memoryBytes,omitempty"`
	// Number of jobs.
	JobCount float64 `json:"jobCount,omitempty"`
	// Number of tasks.
	TaskCount float64 `json:"taskCount,omitempty"`
	// Total duration of this user's MapReduce jobs.
	DurationSec float64 `json:"durationSec,omitempty"`
	// Failed maps of this user's MapReduce jobs. Available since v11.
	FailedMaps float64 `json:"failedMaps,omitempty"`
	// Total maps of this user's MapReduce jobs. Available since v11.
	TotalMaps float64 `json:"totalMaps,omitempty"`
	// Failed reduces of this user's MapReduce jobs. Available since v11.
	FailedReduces float64 `json:"failedReduces,omitempty"`
	// Total reduces of this user's MapReduce jobs. Available since v11.
	TotalReduces float64 `json:"totalReduces,omitempty"`
	// Map input bytes of this user's MapReduce jobs. Available since v11.
	MapInputBytes float64 `json:"mapInputBytes,omitempty"`
	// Map output bytes of this user's MapReduce jobs. Available since v11.
	MapOutputBytes float64 `json:"mapOutputBytes,omitempty"`
	// HDFS bytes read of this user's MapReduce jobs. Available since v11.
	HdfsBytesRead float64 `json:"hdfsBytesRead,omitempty"`
	// HDFS bytes written of this user's MapReduce jobs. Available since v11.
	HdfsBytesWritten float64 `json:"hdfsBytesWritten,omitempty"`
	// Local bytes read of this user's MapReduce jobs. Available since v11.
	LocalBytesRead float64 `json:"localBytesRead,omitempty"`
	// Local bytes written of this user's MapReduce jobs. Available since v11.
	LocalBytesWritten float64 `json:"localBytesWritten,omitempty"`
	// Data local maps of this user's MapReduce jobs. Available since v11.
	DataLocalMaps float64 `json:"dataLocalMaps,omitempty"`
	// Rack local maps of this user's MapReduce jobs. Available since v11.
	RackLocalMaps float64 `json:"rackLocalMaps,omitempty"`
}
