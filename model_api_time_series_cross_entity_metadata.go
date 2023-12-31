/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A class holding additional metadata to the ApiTimeSeriesAggregateStatistics class that applies specifically to cross-entity aggregate metrics.
type ApiTimeSeriesCrossEntityMetadata struct {
	// The display name of the entity that had the maximum value for the cross-entity aggregate metric.
	MaxEntityDisplayName string `json:"maxEntityDisplayName,omitempty"`
	// The name of the entity that had the maximum value for the cross-entity aggregate metric. <p> Available since API v11.
	MaxEntityName string `json:"maxEntityName,omitempty"`
	// The display name of the entity that had the minimum value for the cross-entity aggregate metric.
	MinEntityDisplayName string `json:"minEntityDisplayName,omitempty"`
	// The name of the entity that had the minimum value for the cross-entity aggregate metric. <p> Available since API v11.
	MinEntityName string `json:"minEntityName,omitempty"`
	// The number of entities covered by this point. For a raw cross-entity point this number is exact. For a rollup point this number is an average, since the number of entities being aggregated can change over the aggregation period.
	NumEntities float64 `json:"numEntities,omitempty"`
}
