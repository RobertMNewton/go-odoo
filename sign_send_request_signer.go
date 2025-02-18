package odoo

// SignSendRequestSigner represents sign.send.request.signer model.
type SignSendRequestSigner struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	MailSentOrder     *Int      `xmlrpc:"mail_sent_order,omitempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omitempty"`
	RoleId            *Many2One `xmlrpc:"role_id,omitempty"`
	SignSendRequestId *Many2One `xmlrpc:"sign_send_request_id,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignSendRequestSigners represents array of sign.send.request.signer model.
type SignSendRequestSigners []SignSendRequestSigner

// SignSendRequestSignerModel is the odoo model name.
const SignSendRequestSignerModel = "sign.send.request.signer"

// Many2One convert SignSendRequestSigner to *Many2One.
func (ssrs *SignSendRequestSigner) Many2One() *Many2One {
	return NewMany2One(ssrs.Id.Get(), "")
}

// CreateSignSendRequestSigner creates a new sign.send.request.signer model and returns its id.
func (c *Client) CreateSignSendRequestSigner(ssrs *SignSendRequestSigner) (int64, error) {
	ids, err := c.CreateSignSendRequestSigners([]*SignSendRequestSigner{ssrs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignSendRequestSigner creates a new sign.send.request.signer model and returns its id.
func (c *Client) CreateSignSendRequestSigners(ssrss []*SignSendRequestSigner) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrss {
		vv = append(vv, v)
	}
	return c.Create(SignSendRequestSignerModel, vv, nil)
}

// UpdateSignSendRequestSigner updates an existing sign.send.request.signer record.
func (c *Client) UpdateSignSendRequestSigner(ssrs *SignSendRequestSigner) error {
	return c.UpdateSignSendRequestSigners([]int64{ssrs.Id.Get()}, ssrs)
}

// UpdateSignSendRequestSigners updates existing sign.send.request.signer records.
// All records (represented by ids) will be updated by ssrs values.
func (c *Client) UpdateSignSendRequestSigners(ids []int64, ssrs *SignSendRequestSigner) error {
	return c.Update(SignSendRequestSignerModel, ids, ssrs, nil)
}

// DeleteSignSendRequestSigner deletes an existing sign.send.request.signer record.
func (c *Client) DeleteSignSendRequestSigner(id int64) error {
	return c.DeleteSignSendRequestSigners([]int64{id})
}

// DeleteSignSendRequestSigners deletes existing sign.send.request.signer records.
func (c *Client) DeleteSignSendRequestSigners(ids []int64) error {
	return c.Delete(SignSendRequestSignerModel, ids)
}

// GetSignSendRequestSigner gets sign.send.request.signer existing record.
func (c *Client) GetSignSendRequestSigner(id int64) (*SignSendRequestSigner, error) {
	ssrss, err := c.GetSignSendRequestSigners([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssrss)[0]), nil
}

// GetSignSendRequestSigners gets sign.send.request.signer existing records.
func (c *Client) GetSignSendRequestSigners(ids []int64) (*SignSendRequestSigners, error) {
	ssrss := &SignSendRequestSigners{}
	if err := c.Read(SignSendRequestSignerModel, ids, nil, ssrss); err != nil {
		return nil, err
	}
	return ssrss, nil
}

// FindSignSendRequestSigner finds sign.send.request.signer record by querying it with criteria.
func (c *Client) FindSignSendRequestSigner(criteria *Criteria) (*SignSendRequestSigner, error) {
	ssrss := &SignSendRequestSigners{}
	if err := c.SearchRead(SignSendRequestSignerModel, criteria, NewOptions().Limit(1), ssrss); err != nil {
		return nil, err
	}
	return &((*ssrss)[0]), nil
}

// FindSignSendRequestSigners finds sign.send.request.signer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignSendRequestSigners(criteria *Criteria, options *Options) (*SignSendRequestSigners, error) {
	ssrss := &SignSendRequestSigners{}
	if err := c.SearchRead(SignSendRequestSignerModel, criteria, options, ssrss); err != nil {
		return nil, err
	}
	return ssrss, nil
}

// FindSignSendRequestSignerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignSendRequestSignerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignSendRequestSignerModel, criteria, options)
}

// FindSignSendRequestSignerId finds record id by querying it with criteria.
func (c *Client) FindSignSendRequestSignerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignSendRequestSignerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
