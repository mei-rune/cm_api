/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Arguments to run bandwidth diagnostics as part of performance inspector. Requires iperf3 package installed on hosts.
type ApiPerfInspectorBandwidthArgs struct {
	// Optional flag to run bandwidth diagnostics test. Exercise caution, running bandwidth test will have an impact on currently running workloads. If not specified, defaults to false.
	RunBandwidthDiagnostics bool `json:"runBandwidthDiagnostics,omitempty"`
	// Timeout in seconds for the bandwidth request to each target host. If not specified, defaults to 10 seconds.
	BandwidthTimeoutSecs float64 `json:"bandwidthTimeoutSecs,omitempty"`
}
