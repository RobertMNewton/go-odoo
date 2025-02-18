package odoo

// SignItemOption represents sign.item.option model.
type SignItemOption struct {
	Available   *Bool     `xmlrpc:"available,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignItemOptions represents array of sign.item.option model.
type SignItemOptions []SignItemOption

// SignItemOptionModel is the odoo model name.
const SignItemOptionModel = "sign.item.option"

// Many2One convert SignItemOption to *Many2One.
func (sio *SignItemOption) Many2One() *Many2One {
	return NewMany2One(sio.Id.Get(), "")
}

// CreateSignItemOption creates a new sign.item.option model and returns its id.
func (c *Client) CreateSignItemOption(sio *SignItemOption) (int64, error) {
	ids, err := c.CreateSignItemOptions([]*SignItemOption{sio})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignItemOption creates a new sign.item.option model and returns its id.
func (c *Client) CreateSignItemOptions(sios []*SignItemOption) ([]int64, error) {
	var vv []interface{}
	for _, v := range sios {
		vv = append(vv, v)
	}
	return c.Create(SignItemOptionModel, vv, nil)
}

// UpdateSignItemOption updates an existing sign.item.option record.
func (c *Client) UpdateSignItemOption(sio *SignItemOption) error {
	return c.UpdateSignItemOptions([]int64{sio.Id.Get()}, sio)
}

// UpdateSignItemOptions updates existing sign.item.option records.
// All records (represented by ids) will be updated by sio values.
func (c *Client) UpdateSignItemOptions(ids []int64, sio *SignItemOption) error {
	return c.Update(SignItemOptionModel, ids, sio, nil)
}

// DeleteSignItemOption deletes an existing sign.item.option record.
func (c *Client) DeleteSignItemOption(id int64) error {
	return c.DeleteSignItemOptions([]int64{id})
}

// DeleteSignItemOptions deletes existing sign.item.option records.
func (c *Client) DeleteSignItemOptions(ids []int64) error {
	return c.Delete(SignItemOptionModel, ids)
}

// GetSignItemOption gets sign.item.option existing record.
func (c *Client) GetSignItemOption(id int64) (*SignItemOption, error) {
	sios, err := c.GetSignItemOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sios)[0]), nil
}

// GetSignItemOptions gets sign.item.option existing records.
func (c *Client) GetSignItemOptions(ids []int64) (*SignItemOptions, error) {
	sios := &SignItemOptions{}
	if err := c.Read(SignItemOptionModel, ids, nil, sios); err != nil {
		return nil, err
	}
	return sios, nil
}

// FindSignItemOption finds sign.item.option record by querying it with criteria.
func (c *Client) FindSignItemOption(criteria *Criteria) (*SignItemOption, error) {
	sios := &SignItemOptions{}
	if err := c.SearchRead(SignItemOptionModel, criteria, NewOptions().Limit(1), sios); err != nil {
		return nil, err
	}
	return &((*sios)[0]), nil
}

// FindSignItemOptions finds sign.item.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemOptions(criteria *Criteria, options *Options) (*SignItemOptions, error) {
	sios := &SignItemOptions{}
	if err := c.SearchRead(SignItemOptionModel, criteria, options, sios); err != nil {
		return nil, err
	}
	return sios, nil
}

// FindSignItemOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignItemOptionModel, criteria, options)
}

// FindSignItemOptionId finds record id by querying it with criteria.
func (c *Client) FindSignItemOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignItemOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
