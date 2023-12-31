/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Histogram bin of Impala utilization.
type ApiImpalaUtilizationHistogramBin struct {
	// start point (inclusive) of the histogram bin.
	StartPointInclusive float64 `json:"startPointInclusive,omitempty"`
	// end point (exclusive) of the histogram bin.
	EndPointExclusive float64 `json:"endPointExclusive,omitempty"`
	// Number of Impala daemons.
	NumberOfImpalaDaemons float64 `json:"numberOfImpalaDaemons,omitempty"`
}
