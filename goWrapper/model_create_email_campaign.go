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

type CreateEmailCampaign struct {
	// Tag of the campaign
	Tag string `json:"tag,omitempty"`
	Sender *CreateEmailCampaignSender `json:"sender"`
	// Name of the campaign
	Name string `json:"name"`
	// Mandatory if htmlUrl and templateId are empty. Body of the message (HTML)
	HtmlContent string `json:"htmlContent,omitempty"`
	// Mandatory if htmlContent and templateId are empty. Url to the message (HTML)
	HtmlUrl string `json:"htmlUrl,omitempty"`
	// Mandatory if htmlContent and htmlUrl are empty. Id of the transactional email template with status 'active'. Used to copy only its content fetched from htmlContent/htmlUrl to an email campaign for RSS feature.
	TemplateId int64 `json:"templateId,omitempty"`
	// Sending UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result. If sendAtBestTime is set to true, your campaign will be sent according to the date passed (ignoring the time part).
	ScheduledAt string `json:"scheduledAt,omitempty"`
	// Subject of the campaign. Mandatory if abTesting is false. Ignored if abTesting is true.
	Subject string `json:"subject,omitempty"`
	// Email on which the campaign recipients will be able to reply to
	ReplyTo string `json:"replyTo,omitempty"`
	// To personalize the «To» Field. If you want to include the first name and last name of your recipient, add {FNAME} {LNAME}. These contact attributes must already exist in your SendinBlue account. If input parameter 'params' used please use {{contact.FNAME}} {{contact.LNAME}} for personalization
	ToField string `json:"toField,omitempty"`
	Recipients *CreateEmailCampaignRecipients `json:"recipients,omitempty"`
	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentUrl string `json:"attachmentUrl,omitempty"`
	// Use true to embedded the images in your email. Final size of the email should be less than 4MB. Campaigns with embedded images can not be sent to more than 5000 contacts
	InlineImageActivation bool `json:"inlineImageActivation,omitempty"`
	// Use true to enable the mirror link
	MirrorActive bool `json:"mirrorActive,omitempty"`
	// Footer of the email campaign
	Footer string `json:"footer,omitempty"`
	// Header of the email campaign
	Header string `json:"header,omitempty"`
	// Customize the utm_campaign value. If this field is empty, the campaign name will be used. Only alphanumeric characters and spaces are allowed
	UtmCampaign string `json:"utmCampaign,omitempty"`
	// Pass the set of attributes to customize the type classic campaign. For example, {\"FNAME\":\"Joe\", \"LNAME\":\"Doe\"}. Only available if 'type' is 'classic'. It's considered only if campaign is in New Template Language format. The New Template Language is dependent on the values of 'subject', 'htmlContent/htmlUrl', 'sender.name' & 'toField'
	Params *interface{} `json:"params,omitempty"`
	// Set this to true if you want to send your campaign at best time.
	SendAtBestTime bool `json:"sendAtBestTime,omitempty"`
	// Status of A/B Test. abTesting = false means it is disabled, & abTesting = true means it is enabled. 'subjectA', 'subjectB', 'splitRule', 'winnerCriteria' & 'winnerDelay' will be considered when abTesting is set to true. 'subjectA' & 'subjectB' are mandatory together & 'subject' if passed is ignored. Can be set to true only if 'sendAtBestTime' is 'false'. You will be able to set up two subject lines for your campaign and send them to a random sample of your total recipients. Half of the test group will receive version A, and the other half will receive version B
	AbTesting bool `json:"abTesting,omitempty"`
	// Subject A of the campaign. Mandatory if abTesting = true. subjectA & subjectB should have unique value
	SubjectA string `json:"subjectA,omitempty"`
	// Subject B of the campaign. Mandatory if abTesting = true. subjectA & subjectB should have unique value
	SubjectB string `json:"subjectB,omitempty"`
	// Add the size of your test groups. Mandatory if abTesting = true & 'recipients' is passed. We'll send version A and B to a random sample of recipients, and then the winning version to everyone else
	SplitRule int64 `json:"splitRule,omitempty"`
	// Choose the metrics that will determinate the winning version. Mandatory if 'splitRule' >= 1 and < 50. If splitRule = 50, 'winnerCriteria' is ignored if passed
	WinnerCriteria string `json:"winnerCriteria,omitempty"`
	// Choose the duration of the test in hours. Maximum is 7 days, pass 24*7 = 168 hours. The winning version will be sent at the end of the test. Mandatory if 'splitRule' >= 1 and < 50. If splitRule = 50, 'winnerDelay' is ignored if passed
	WinnerDelay int64 `json:"winnerDelay,omitempty"`
	// Available for dedicated ip clients. Set this to true if you wish to warm up your ip.
	IpWarmupEnable bool `json:"ipWarmupEnable,omitempty"`
	// Mandatory if ipWarmupEnable is set to true. Set an initial quota greater than 1 for warming up your ip. We recommend you set a value of 3000.
	InitialQuota int64 `json:"initialQuota,omitempty"`
	// Mandatory if ipWarmupEnable is set to true. Set a percentage increase rate for warming up your ip. We recommend you set the increase rate to 30% per day. If you want to send the same number of emails every day, set the daily increase value to 0%.
	IncreaseRate int64 `json:"increaseRate,omitempty"`
}
