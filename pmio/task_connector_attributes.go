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

type TaskConnectorAttributes struct {

	TaskId int32 `json:"task_id,omitempty"`

	ConnectorClass string `json:"connector_class,omitempty"`

	InputParameters string `json:"input_parameters,omitempty"`

	OutputParameters string `json:"output_parameters,omitempty"`

	AsyncBefore bool `json:"async_before,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`
}
