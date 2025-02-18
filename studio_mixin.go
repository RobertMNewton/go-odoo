package odoo

// StudioMixin represents studio.mixin model.
type StudioMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// StudioMixins represents array of studio.mixin model.
type StudioMixins []StudioMixin

// StudioMixinModel is the odoo model name.
const StudioMixinModel = "studio.mixin"

// Many2One convert StudioMixin to *Many2One.
func (sm *StudioMixin) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateStudioMixin creates a new studio.mixin model and returns its id.
func (c *Client) CreateStudioMixin(sm *StudioMixin) (int64, error) {
	ids, err := c.CreateStudioMixins([]*StudioMixin{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioMixin creates a new studio.mixin model and returns its id.
func (c *Client) CreateStudioMixins(sms []*StudioMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(StudioMixinModel, vv, nil)
}

// UpdateStudioMixin updates an existing studio.mixin record.
func (c *Client) UpdateStudioMixin(sm *StudioMixin) error {
	return c.UpdateStudioMixins([]int64{sm.Id.Get()}, sm)
}

// UpdateStudioMixins updates existing studio.mixin records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateStudioMixins(ids []int64, sm *StudioMixin) error {
	return c.Update(StudioMixinModel, ids, sm, nil)
}

// DeleteStudioMixin deletes an existing studio.mixin record.
func (c *Client) DeleteStudioMixin(id int64) error {
	return c.DeleteStudioMixins([]int64{id})
}

// DeleteStudioMixins deletes existing studio.mixin records.
func (c *Client) DeleteStudioMixins(ids []int64) error {
	return c.Delete(StudioMixinModel, ids)
}

// GetStudioMixin gets studio.mixin existing record.
func (c *Client) GetStudioMixin(id int64) (*StudioMixin, error) {
	sms, err := c.GetStudioMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// GetStudioMixins gets studio.mixin existing records.
func (c *Client) GetStudioMixins(ids []int64) (*StudioMixins, error) {
	sms := &StudioMixins{}
	if err := c.Read(StudioMixinModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStudioMixin finds studio.mixin record by querying it with criteria.
func (c *Client) FindStudioMixin(criteria *Criteria) (*StudioMixin, error) {
	sms := &StudioMixins{}
	if err := c.SearchRead(StudioMixinModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// FindStudioMixins finds studio.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioMixins(criteria *Criteria, options *Options) (*StudioMixins, error) {
	sms := &StudioMixins{}
	if err := c.SearchRead(StudioMixinModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStudioMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioMixinModel, criteria, options)
}

// FindStudioMixinId finds record id by querying it with criteria.
func (c *Client) FindStudioMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
