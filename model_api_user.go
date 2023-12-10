/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// This is the model for user information in the API prior to v18. Post v18, please refer to ApiUser2.java. <p> Note that any method that returns user information will not contain any password information. The password property is only used when creating or updating users.
type ApiUser struct {
	// The username, which is unique within a Cloudera Manager installation.
	Name string `json:"name,omitempty"`
	// Returns the user password. <p> Passwords are not returned when querying user information, so this property will always be empty when reading information from a server.
	Password string `json:"password,omitempty"`
	// A list of roles this user belongs to. <p> In Cloudera Express, possible values are: <ul> <li><b>ROLE_ADMIN</b></li> <li><b>ROLE_USER</b></li> </ul> In Cloudera Enterprise Datahub Edition, additional possible values are: <ul> <li><b>ROLE_LIMITED</b>: Added in Cloudera Manager 5.0</li> <li><b>ROLE_OPERATOR</b>: Added in Cloudera Manager 5.1</li> <li><b>ROLE_CONFIGURATOR</b>: Added in Cloudera Manager 5.1</li> <li><b>ROLE_CLUSTER_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_BDR_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_NAVIGATOR_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_USER_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_KEY_ADMIN</b>: Added in Cloudera Manager 5.5</li> </ul> An empty list implies ROLE_USER. <p> Note that although this interface provides a list of roles, a user should only be assigned a single role at a time.
	Roles []string `json:"roles,omitempty"`
	// NOTE: Only available in the \"export\" view
	PwHash string `json:"pwHash,omitempty"`
	// NOTE: Only available in the \"export\" view
	PwSalt float64 `json:"pwSalt,omitempty"`
	// NOTE: Only available in the \"export\" view
	PwLogin bool `json:"pwLogin,omitempty"`
}
