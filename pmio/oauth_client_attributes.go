/* 
 * example
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package pmio

type OauthClientAttributes struct {

	Secret string `json:"secret,omitempty"`

	Name string `json:"name"`

	PersonalAccessClient bool `json:"personal_access_client,omitempty"`

	PasswordClient bool `json:"password_client,omitempty"`

	Revoked bool `json:"revoked,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`
}