/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  | 
 *
 * API version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package newtest

type GetExtendedContactDetails struct {
	// Email address of the contact for which you requested the details
	Email string `json:"email"`
	// ID of the contact for which you requested the details
	Id int64 `json:"id"`
	// Blacklist status for email campaigns (true=blacklisted, false=not blacklisted)
	EmailBlacklisted bool `json:"emailBlacklisted"`
	// Blacklist status for SMS campaigns (true=blacklisted, false=not blacklisted)
	SmsBlacklisted bool `json:"smsBlacklisted"`
	// Creation UTC date-time of the contact (YYYY-MM-DDTHH:mm:ss.SSSZ)
	CreatedAt string `json:"createdAt"`
	// Last modification UTC date-time of the contact (YYYY-MM-DDTHH:mm:ss.SSSZ)
	ModifiedAt string `json:"modifiedAt"`
	ListIds []int64 `json:"listIds"`
	ListUnsubscribed []int64 `json:"listUnsubscribed,omitempty"`
	// Set of attributes of the contact
	Attributes *interface{} `json:"attributes"`
	Statistics *GetExtendedContactDetailsStatistics `json:"statistics"`
}
