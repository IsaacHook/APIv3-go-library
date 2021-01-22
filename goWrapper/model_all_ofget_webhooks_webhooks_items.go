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

type AllOfgetWebhooksWebhooksItems struct {
	// URL of the webhook
	Url string `json:"url"`
	// ID of the webhook
	Id int64 `json:"id"`
	// Description of the webhook
	Description string `json:"description"`
	Events []string `json:"events"`
	// Type of webhook (marketing or transac)
	Type_ string `json:"type"`
	// Creation UTC date-time of the webhook (YYYY-MM-DDTHH:mm:ss.SSSZ)
	CreatedAt string `json:"createdAt"`
	// Last modification UTC date-time of the webhook (YYYY-MM-DDTHH:mm:ss.SSSZ)
	ModifiedAt string `json:"modifiedAt"`
}
