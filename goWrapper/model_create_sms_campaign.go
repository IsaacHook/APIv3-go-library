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

type CreateSmsCampaign struct {
	// Name of the campaign
	Name string `json:"name"`
	// Name of the sender. **The number of characters is limited to 11 for alphanumeric characters and 15 for numeric characters**
	Sender string `json:"sender"`
	// Content of the message. The maximum characters used per SMS is 160, if used more than that, it will be counted as more than one SMS
	Content string `json:"content"`
	Recipients *CreateSmsCampaignRecipients `json:"recipients,omitempty"`
	// UTC date-time on which the campaign has to run (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result.
	ScheduledAt string `json:"scheduledAt,omitempty"`
}
