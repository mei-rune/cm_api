/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Arguments used for Rolling Restart commands.
type ApiImpalaRoleDiagnosticsArgs struct {
	// The support ticket number to attach to this data collection.
	TicketNumber string `json:"ticketNumber,omitempty"`
	// Comments to include with this data collection.
	Comments    string  `json:"comments,omitempty"`
	StacksCount float64 `json:"stacksCount,omitempty"`
	// Interval between stack collections.
	StacksIntervalSeconds    float64 `json:"stacksIntervalSeconds,omitempty"`
	Jmap                     bool    `json:"jmap,omitempty"`
	Gcore                    bool    `json:"gcore,omitempty"`
	MinidumpsCount           float64 `json:"minidumpsCount,omitempty"`
	MinidumpsIntervalSeconds float64 `json:"minidumpsIntervalSeconds,omitempty"`
	PhoneHome                bool    `json:"phoneHome,omitempty"`
}
