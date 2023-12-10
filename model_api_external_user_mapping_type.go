/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// ApiExternalUserMappingType : Enum for external user mapping types
type ApiExternalUserMappingType string

// List of ApiExternalUserMappingType
const (
	LDAP_ApiExternalUserMappingType             ApiExternalUserMappingType = "LDAP"
	SAML_SCRIPT_ApiExternalUserMappingType      ApiExternalUserMappingType = "SAML_SCRIPT"
	SAML_ATTRIBUTE_ApiExternalUserMappingType   ApiExternalUserMappingType = "SAML_ATTRIBUTE"
	EXTERNAL_PROGRAM_ApiExternalUserMappingType ApiExternalUserMappingType = "EXTERNAL_PROGRAM"
)
