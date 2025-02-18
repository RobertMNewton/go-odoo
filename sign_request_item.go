package odoo

// SignRequestItem represents sign.request.item model.
type SignRequestItem struct {
	AccessToken            *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl              *String    `xmlrpc:"access_url,omitempty"`
	AccessViaLink          *Bool      `xmlrpc:"access_via_link,omitempty"`
	AccessWarning          *String    `xmlrpc:"access_warning,omitempty"`
	ChangeAuthorized       *Bool      `xmlrpc:"change_authorized,omitempty"`
	Color                  *Int       `xmlrpc:"color,omitempty"`
	CommunicationCompanyId *Many2One  `xmlrpc:"communication_company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	FrameHash              *String    `xmlrpc:"frame_hash,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	IsMailSent             *Bool      `xmlrpc:"is_mail_sent,omitempty"`
	Latitude               *Float     `xmlrpc:"latitude,omitempty"`
	Longitude              *Float     `xmlrpc:"longitude,omitempty"`
	MailSentOrder          *Int       `xmlrpc:"mail_sent_order,omitempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omitempty"`
	Reference              *String    `xmlrpc:"reference,omitempty"`
	RoleId                 *Many2One  `xmlrpc:"role_id,omitempty"`
	SignItemValueIds       *Relation  `xmlrpc:"sign_item_value_ids,omitempty"`
	SignRequestId          *Many2One  `xmlrpc:"sign_request_id,omitempty"`
	Signature              *String    `xmlrpc:"signature,omitempty"`
	SignedWithoutExtraAuth *Bool      `xmlrpc:"signed_without_extra_auth,omitempty"`
	SignerEmail            *String    `xmlrpc:"signer_email,omitempty"`
	SigningDate            *Time      `xmlrpc:"signing_date,omitempty"`
	SmsNumber              *String    `xmlrpc:"sms_number,omitempty"`
	SmsToken               *String    `xmlrpc:"sms_token,omitempty"`
	State                  *Selection `xmlrpc:"state,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SignRequestItems represents array of sign.request.item model.
type SignRequestItems []SignRequestItem

// SignRequestItemModel is the odoo model name.
const SignRequestItemModel = "sign.request.item"

// Many2One convert SignRequestItem to *Many2One.
func (sri *SignRequestItem) Many2One() *Many2One {
	return NewMany2One(sri.Id.Get(), "")
}

// CreateSignRequestItem creates a new sign.request.item model and returns its id.
func (c *Client) CreateSignRequestItem(sri *SignRequestItem) (int64, error) {
	ids, err := c.CreateSignRequestItems([]*SignRequestItem{sri})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignRequestItem creates a new sign.request.item model and returns its id.
func (c *Client) CreateSignRequestItems(sris []*SignRequestItem) ([]int64, error) {
	var vv []interface{}
	for _, v := range sris {
		vv = append(vv, v)
	}
	return c.Create(SignRequestItemModel, vv, nil)
}

// UpdateSignRequestItem updates an existing sign.request.item record.
func (c *Client) UpdateSignRequestItem(sri *SignRequestItem) error {
	return c.UpdateSignRequestItems([]int64{sri.Id.Get()}, sri)
}

// UpdateSignRequestItems updates existing sign.request.item records.
// All records (represented by ids) will be updated by sri values.
func (c *Client) UpdateSignRequestItems(ids []int64, sri *SignRequestItem) error {
	return c.Update(SignRequestItemModel, ids, sri, nil)
}

// DeleteSignRequestItem deletes an existing sign.request.item record.
func (c *Client) DeleteSignRequestItem(id int64) error {
	return c.DeleteSignRequestItems([]int64{id})
}

// DeleteSignRequestItems deletes existing sign.request.item records.
func (c *Client) DeleteSignRequestItems(ids []int64) error {
	return c.Delete(SignRequestItemModel, ids)
}

// GetSignRequestItem gets sign.request.item existing record.
func (c *Client) GetSignRequestItem(id int64) (*SignRequestItem, error) {
	sris, err := c.GetSignRequestItems([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sris)[0]), nil
}

// GetSignRequestItems gets sign.request.item existing records.
func (c *Client) GetSignRequestItems(ids []int64) (*SignRequestItems, error) {
	sris := &SignRequestItems{}
	if err := c.Read(SignRequestItemModel, ids, nil, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindSignRequestItem finds sign.request.item record by querying it with criteria.
func (c *Client) FindSignRequestItem(criteria *Criteria) (*SignRequestItem, error) {
	sris := &SignRequestItems{}
	if err := c.SearchRead(SignRequestItemModel, criteria, NewOptions().Limit(1), sris); err != nil {
		return nil, err
	}
	return &((*sris)[0]), nil
}

// FindSignRequestItems finds sign.request.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequestItems(criteria *Criteria, options *Options) (*SignRequestItems, error) {
	sris := &SignRequestItems{}
	if err := c.SearchRead(SignRequestItemModel, criteria, options, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindSignRequestItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequestItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignRequestItemModel, criteria, options)
}

// FindSignRequestItemId finds record id by querying it with criteria.
func (c *Client) FindSignRequestItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignRequestItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
