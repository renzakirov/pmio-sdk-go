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

// JSON API response with a collection of the tokens in an array
type TokenCollection struct {

	Data []Token `json:"data,omitempty"`

	Meta Meta `json:"meta,omitempty"`
}
