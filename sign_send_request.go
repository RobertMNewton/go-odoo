package odoo

// SignSendRequest represents sign.send.request model.
type SignSendRequest struct {
	ActivityId         *Many2One `xmlrpc:"activity_id,omitempty"`
	AttachmentIds      *Relation `xmlrpc:"attachment_ids,omitempty"`
	CcPartnerIds       *Relation `xmlrpc:"cc_partner_ids,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Filename           *String   `xmlrpc:"filename,omitempty"`
	HasDefaultTemplate *Bool     `xmlrpc:"has_default_template,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	IsUserSigner       *Bool     `xmlrpc:"is_user_signer,omitempty"`
	Message            *String   `xmlrpc:"message,omitempty"`
	MessageCc          *String   `xmlrpc:"message_cc,omitempty"`
	ReferenceDoc       *String   `xmlrpc:"reference_doc,omitempty"`
	Reminder           *Int      `xmlrpc:"reminder,omitempty"`
	ReminderEnabled    *Bool     `xmlrpc:"reminder_enabled,omitempty"`
	SetSignOrder       *Bool     `xmlrpc:"set_sign_order,omitempty"`
	SignerId           *Many2One `xmlrpc:"signer_id,omitempty"`
	SignerIds          *Relation `xmlrpc:"signer_ids,omitempty"`
	SignersCount       *Int      `xmlrpc:"signers_count,omitempty"`
	Subject            *String   `xmlrpc:"subject,omitempty"`
	TemplateId         *Many2One `xmlrpc:"template_id,omitempty"`
	Validity           *Time     `xmlrpc:"validity,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignSendRequests represents array of sign.send.request model.
type SignSendRequests []SignSendRequest

// SignSendRequestModel is the odoo model name.
const SignSendRequestModel = "sign.send.request"

// Many2One convert SignSendRequest to *Many2One.
func (ssr *SignSendRequest) Many2One() *Many2One {
	return NewMany2One(ssr.Id.Get(), "")
}

// CreateSignSendRequest creates a new sign.send.request model and returns its id.
func (c *Client) CreateSignSendRequest(ssr *SignSendRequest) (int64, error) {
	ids, err := c.CreateSignSendRequests([]*SignSendRequest{ssr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignSendRequest creates a new sign.send.request model and returns its id.
func (c *Client) CreateSignSendRequests(ssrs []*SignSendRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrs {
		vv = append(vv, v)
	}
	return c.Create(SignSendRequestModel, vv, nil)
}

// UpdateSignSendRequest updates an existing sign.send.request record.
func (c *Client) UpdateSignSendRequest(ssr *SignSendRequest) error {
	return c.UpdateSignSendRequests([]int64{ssr.Id.Get()}, ssr)
}

// UpdateSignSendRequests updates existing sign.send.request records.
// All records (represented by ids) will be updated by ssr values.
func (c *Client) UpdateSignSendRequests(ids []int64, ssr *SignSendRequest) error {
	return c.Update(SignSendRequestModel, ids, ssr, nil)
}

// DeleteSignSendRequest deletes an existing sign.send.request record.
func (c *Client) DeleteSignSendRequest(id int64) error {
	return c.DeleteSignSendRequests([]int64{id})
}

// DeleteSignSendRequests deletes existing sign.send.request records.
func (c *Client) DeleteSignSendRequests(ids []int64) error {
	return c.Delete(SignSendRequestModel, ids)
}

// GetSignSendRequest gets sign.send.request existing record.
func (c *Client) GetSignSendRequest(id int64) (*SignSendRequest, error) {
	ssrs, err := c.GetSignSendRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssrs)[0]), nil
}

// GetSignSendRequests gets sign.send.request existing records.
func (c *Client) GetSignSendRequests(ids []int64) (*SignSendRequests, error) {
	ssrs := &SignSendRequests{}
	if err := c.Read(SignSendRequestModel, ids, nil, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSignSendRequest finds sign.send.request record by querying it with criteria.
func (c *Client) FindSignSendRequest(criteria *Criteria) (*SignSendRequest, error) {
	ssrs := &SignSendRequests{}
	if err := c.SearchRead(SignSendRequestModel, criteria, NewOptions().Limit(1), ssrs); err != nil {
		return nil, err
	}
	return &((*ssrs)[0]), nil
}

// FindSignSendRequests finds sign.send.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignSendRequests(criteria *Criteria, options *Options) (*SignSendRequests, error) {
	ssrs := &SignSendRequests{}
	if err := c.SearchRead(SignSendRequestModel, criteria, options, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSignSendRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignSendRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignSendRequestModel, criteria, options)
}

// FindSignSendRequestId finds record id by querying it with criteria.
func (c *Client) FindSignSendRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignSendRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
