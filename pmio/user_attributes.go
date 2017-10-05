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

type UserAttributes struct {

	Username string `json:"username"`

	Password string `json:"password"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	Email string `json:"email"`

	ExpiresAt string `json:"expires_at,omitempty"`

	Status string `json:"status,omitempty"`

	Country string `json:"country,omitempty"`

	City string `json:"city,omitempty"`

	Location string `json:"location,omitempty"`

	Address string `json:"address,omitempty"`

	Phone string `json:"phone,omitempty"`

	Fax string `json:"fax,omitempty"`

	Cellular string `json:"cellular,omitempty"`

	ZipCode string `json:"zip_code,omitempty"`

	Position string `json:"position,omitempty"`

	Resume string `json:"resume,omitempty"`

	BirthdayAt string `json:"birthday_at,omitempty"`

	BookmarkStartCases string `json:"bookmark_start_cases,omitempty"`

	TimeZone string `json:"time_zone,omitempty"`

	DefaultLang string `json:"default_lang,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	Clients []int32 `json:"clients,omitempty"`
}
