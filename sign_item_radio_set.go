package odoo

// SignItemRadioSet represents sign.item.radio.set model.
type SignItemRadioSet struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	NumOptions  *Int      `xmlrpc:"num_options,omitempty"`
	RadioItems  *Relation `xmlrpc:"radio_items,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignItemRadioSets represents array of sign.item.radio.set model.
type SignItemRadioSets []SignItemRadioSet

// SignItemRadioSetModel is the odoo model name.
const SignItemRadioSetModel = "sign.item.radio.set"

// Many2One convert SignItemRadioSet to *Many2One.
func (sirs *SignItemRadioSet) Many2One() *Many2One {
	return NewMany2One(sirs.Id.Get(), "")
}

// CreateSignItemRadioSet creates a new sign.item.radio.set model and returns its id.
func (c *Client) CreateSignItemRadioSet(sirs *SignItemRadioSet) (int64, error) {
	ids, err := c.CreateSignItemRadioSets([]*SignItemRadioSet{sirs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignItemRadioSet creates a new sign.item.radio.set model and returns its id.
func (c *Client) CreateSignItemRadioSets(sirss []*SignItemRadioSet) ([]int64, error) {
	var vv []interface{}
	for _, v := range sirss {
		vv = append(vv, v)
	}
	return c.Create(SignItemRadioSetModel, vv, nil)
}

// UpdateSignItemRadioSet updates an existing sign.item.radio.set record.
func (c *Client) UpdateSignItemRadioSet(sirs *SignItemRadioSet) error {
	return c.UpdateSignItemRadioSets([]int64{sirs.Id.Get()}, sirs)
}

// UpdateSignItemRadioSets updates existing sign.item.radio.set records.
// All records (represented by ids) will be updated by sirs values.
func (c *Client) UpdateSignItemRadioSets(ids []int64, sirs *SignItemRadioSet) error {
	return c.Update(SignItemRadioSetModel, ids, sirs, nil)
}

// DeleteSignItemRadioSet deletes an existing sign.item.radio.set record.
func (c *Client) DeleteSignItemRadioSet(id int64) error {
	return c.DeleteSignItemRadioSets([]int64{id})
}

// DeleteSignItemRadioSets deletes existing sign.item.radio.set records.
func (c *Client) DeleteSignItemRadioSets(ids []int64) error {
	return c.Delete(SignItemRadioSetModel, ids)
}

// GetSignItemRadioSet gets sign.item.radio.set existing record.
func (c *Client) GetSignItemRadioSet(id int64) (*SignItemRadioSet, error) {
	sirss, err := c.GetSignItemRadioSets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sirss)[0]), nil
}

// GetSignItemRadioSets gets sign.item.radio.set existing records.
func (c *Client) GetSignItemRadioSets(ids []int64) (*SignItemRadioSets, error) {
	sirss := &SignItemRadioSets{}
	if err := c.Read(SignItemRadioSetModel, ids, nil, sirss); err != nil {
		return nil, err
	}
	return sirss, nil
}

// FindSignItemRadioSet finds sign.item.radio.set record by querying it with criteria.
func (c *Client) FindSignItemRadioSet(criteria *Criteria) (*SignItemRadioSet, error) {
	sirss := &SignItemRadioSets{}
	if err := c.SearchRead(SignItemRadioSetModel, criteria, NewOptions().Limit(1), sirss); err != nil {
		return nil, err
	}
	return &((*sirss)[0]), nil
}

// FindSignItemRadioSets finds sign.item.radio.set records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemRadioSets(criteria *Criteria, options *Options) (*SignItemRadioSets, error) {
	sirss := &SignItemRadioSets{}
	if err := c.SearchRead(SignItemRadioSetModel, criteria, options, sirss); err != nil {
		return nil, err
	}
	return sirss, nil
}

// FindSignItemRadioSetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemRadioSetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignItemRadioSetModel, criteria, options)
}

// FindSignItemRadioSetId finds record id by querying it with criteria.
func (c *Client) FindSignItemRadioSetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignItemRadioSetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
