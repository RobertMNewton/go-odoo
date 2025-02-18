package odoo

// SignRequest represents sign.request model.
type SignRequest struct {
	AccessToken                    *String    `xmlrpc:"access_token,omitempty"`
	Active                         *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId        *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration    *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon          *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon               *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttachmentIds                  *Relation  `xmlrpc:"attachment_ids,omitempty"`
	CcPartnerIds                   *Relation  `xmlrpc:"cc_partner_ids,omitempty"`
	Color                          *Int       `xmlrpc:"color,omitempty"`
	CommunicationCompanyId         *Many2One  `xmlrpc:"communication_company_id,omitempty"`
	CompletedDocument              *String    `xmlrpc:"completed_document,omitempty"`
	CompletedDocumentAttachmentIds *Relation  `xmlrpc:"completed_document_attachment_ids,omitempty"`
	CompletionDate                 *Time      `xmlrpc:"completion_date,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	FavoritedIds                   *Relation  `xmlrpc:"favorited_ids,omitempty"`
	HasMessage                     *Bool      `xmlrpc:"has_message,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	Integrity                      *Bool      `xmlrpc:"integrity,omitempty"`
	LastActionDate                 *Time      `xmlrpc:"last_action_date,omitempty"`
	LastReminder                   *Time      `xmlrpc:"last_reminder,omitempty"`
	Message                        *String    `xmlrpc:"message,omitempty"`
	MessageAttachmentCount         *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageCc                      *String    `xmlrpc:"message_cc,omitempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter         *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError             *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline         *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	NbClosed                       *Int       `xmlrpc:"nb_closed,omitempty"`
	NbTotal                        *Int       `xmlrpc:"nb_total,omitempty"`
	NbWait                         *Int       `xmlrpc:"nb_wait,omitempty"`
	NeedMySignature                *Bool      `xmlrpc:"need_my_signature,omitempty"`
	Progress                       *String    `xmlrpc:"progress,omitempty"`
	RatingIds                      *Relation  `xmlrpc:"rating_ids,omitempty"`
	Reference                      *String    `xmlrpc:"reference,omitempty"`
	ReferenceDoc                   *String    `xmlrpc:"reference_doc,omitempty"`
	Reminder                       *Int       `xmlrpc:"reminder,omitempty"`
	ReminderEnabled                *Bool      `xmlrpc:"reminder_enabled,omitempty"`
	RequestItemIds                 *Relation  `xmlrpc:"request_item_ids,omitempty"`
	RequestItemInfos               *String    `xmlrpc:"request_item_infos,omitempty"`
	ShareLink                      *String    `xmlrpc:"share_link,omitempty"`
	SignLogIds                     *Relation  `xmlrpc:"sign_log_ids,omitempty"`
	StartSign                      *Bool      `xmlrpc:"start_sign,omitempty"`
	State                          *Selection `xmlrpc:"state,omitempty"`
	Subject                        *String    `xmlrpc:"subject,omitempty"`
	TemplateId                     *Many2One  `xmlrpc:"template_id,omitempty"`
	TemplateTags                   *Relation  `xmlrpc:"template_tags,omitempty"`
	Validity                       *Time      `xmlrpc:"validity,omitempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SignRequests represents array of sign.request model.
type SignRequests []SignRequest

// SignRequestModel is the odoo model name.
const SignRequestModel = "sign.request"

// Many2One convert SignRequest to *Many2One.
func (sr *SignRequest) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSignRequest creates a new sign.request model and returns its id.
func (c *Client) CreateSignRequest(sr *SignRequest) (int64, error) {
	ids, err := c.CreateSignRequests([]*SignRequest{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignRequest creates a new sign.request model and returns its id.
func (c *Client) CreateSignRequests(srs []*SignRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(SignRequestModel, vv, nil)
}

// UpdateSignRequest updates an existing sign.request record.
func (c *Client) UpdateSignRequest(sr *SignRequest) error {
	return c.UpdateSignRequests([]int64{sr.Id.Get()}, sr)
}

// UpdateSignRequests updates existing sign.request records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSignRequests(ids []int64, sr *SignRequest) error {
	return c.Update(SignRequestModel, ids, sr, nil)
}

// DeleteSignRequest deletes an existing sign.request record.
func (c *Client) DeleteSignRequest(id int64) error {
	return c.DeleteSignRequests([]int64{id})
}

// DeleteSignRequests deletes existing sign.request records.
func (c *Client) DeleteSignRequests(ids []int64) error {
	return c.Delete(SignRequestModel, ids)
}

// GetSignRequest gets sign.request existing record.
func (c *Client) GetSignRequest(id int64) (*SignRequest, error) {
	srs, err := c.GetSignRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetSignRequests gets sign.request existing records.
func (c *Client) GetSignRequests(ids []int64) (*SignRequests, error) {
	srs := &SignRequests{}
	if err := c.Read(SignRequestModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSignRequest finds sign.request record by querying it with criteria.
func (c *Client) FindSignRequest(criteria *Criteria) (*SignRequest, error) {
	srs := &SignRequests{}
	if err := c.SearchRead(SignRequestModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindSignRequests finds sign.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequests(criteria *Criteria, options *Options) (*SignRequests, error) {
	srs := &SignRequests{}
	if err := c.SearchRead(SignRequestModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSignRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignRequestModel, criteria, options)
}

// FindSignRequestId finds record id by querying it with criteria.
func (c *Client) FindSignRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
