package odoo

// SignRequestItemValue represents sign.request.item.value model.
type SignRequestItemValue struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	FrameHasHash      *Bool     `xmlrpc:"frame_has_hash,omitempty"`
	FrameValue        *String   `xmlrpc:"frame_value,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	SignItemId        *Many2One `xmlrpc:"sign_item_id,omitempty"`
	SignRequestId     *Many2One `xmlrpc:"sign_request_id,omitempty"`
	SignRequestItemId *Many2One `xmlrpc:"sign_request_item_id,omitempty"`
	Value             *String   `xmlrpc:"value,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignRequestItemValues represents array of sign.request.item.value model.
type SignRequestItemValues []SignRequestItemValue

// SignRequestItemValueModel is the odoo model name.
const SignRequestItemValueModel = "sign.request.item.value"

// Many2One convert SignRequestItemValue to *Many2One.
func (sriv *SignRequestItemValue) Many2One() *Many2One {
	return NewMany2One(sriv.Id.Get(), "")
}

// CreateSignRequestItemValue creates a new sign.request.item.value model and returns its id.
func (c *Client) CreateSignRequestItemValue(sriv *SignRequestItemValue) (int64, error) {
	ids, err := c.CreateSignRequestItemValues([]*SignRequestItemValue{sriv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignRequestItemValue creates a new sign.request.item.value model and returns its id.
func (c *Client) CreateSignRequestItemValues(srivs []*SignRequestItemValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range srivs {
		vv = append(vv, v)
	}
	return c.Create(SignRequestItemValueModel, vv, nil)
}

// UpdateSignRequestItemValue updates an existing sign.request.item.value record.
func (c *Client) UpdateSignRequestItemValue(sriv *SignRequestItemValue) error {
	return c.UpdateSignRequestItemValues([]int64{sriv.Id.Get()}, sriv)
}

// UpdateSignRequestItemValues updates existing sign.request.item.value records.
// All records (represented by ids) will be updated by sriv values.
func (c *Client) UpdateSignRequestItemValues(ids []int64, sriv *SignRequestItemValue) error {
	return c.Update(SignRequestItemValueModel, ids, sriv, nil)
}

// DeleteSignRequestItemValue deletes an existing sign.request.item.value record.
func (c *Client) DeleteSignRequestItemValue(id int64) error {
	return c.DeleteSignRequestItemValues([]int64{id})
}

// DeleteSignRequestItemValues deletes existing sign.request.item.value records.
func (c *Client) DeleteSignRequestItemValues(ids []int64) error {
	return c.Delete(SignRequestItemValueModel, ids)
}

// GetSignRequestItemValue gets sign.request.item.value existing record.
func (c *Client) GetSignRequestItemValue(id int64) (*SignRequestItemValue, error) {
	srivs, err := c.GetSignRequestItemValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srivs)[0]), nil
}

// GetSignRequestItemValues gets sign.request.item.value existing records.
func (c *Client) GetSignRequestItemValues(ids []int64) (*SignRequestItemValues, error) {
	srivs := &SignRequestItemValues{}
	if err := c.Read(SignRequestItemValueModel, ids, nil, srivs); err != nil {
		return nil, err
	}
	return srivs, nil
}

// FindSignRequestItemValue finds sign.request.item.value record by querying it with criteria.
func (c *Client) FindSignRequestItemValue(criteria *Criteria) (*SignRequestItemValue, error) {
	srivs := &SignRequestItemValues{}
	if err := c.SearchRead(SignRequestItemValueModel, criteria, NewOptions().Limit(1), srivs); err != nil {
		return nil, err
	}
	return &((*srivs)[0]), nil
}

// FindSignRequestItemValues finds sign.request.item.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequestItemValues(criteria *Criteria, options *Options) (*SignRequestItemValues, error) {
	srivs := &SignRequestItemValues{}
	if err := c.SearchRead(SignRequestItemValueModel, criteria, options, srivs); err != nil {
		return nil, err
	}
	return srivs, nil
}

// FindSignRequestItemValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignRequestItemValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignRequestItemValueModel, criteria, options)
}

// FindSignRequestItemValueId finds record id by querying it with criteria.
func (c *Client) FindSignRequestItemValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignRequestItemValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
