package odoo

// IrMinCronMixin represents ir.min.cron.mixin model.
type IrMinCronMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrMinCronMixins represents array of ir.min.cron.mixin model.
type IrMinCronMixins []IrMinCronMixin

// IrMinCronMixinModel is the odoo model name.
const IrMinCronMixinModel = "ir.min.cron.mixin"

// Many2One convert IrMinCronMixin to *Many2One.
func (imcm *IrMinCronMixin) Many2One() *Many2One {
	return NewMany2One(imcm.Id.Get(), "")
}

// CreateIrMinCronMixin creates a new ir.min.cron.mixin model and returns its id.
func (c *Client) CreateIrMinCronMixin(imcm *IrMinCronMixin) (int64, error) {
	ids, err := c.CreateIrMinCronMixins([]*IrMinCronMixin{imcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrMinCronMixin creates a new ir.min.cron.mixin model and returns its id.
func (c *Client) CreateIrMinCronMixins(imcms []*IrMinCronMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range imcms {
		vv = append(vv, v)
	}
	return c.Create(IrMinCronMixinModel, vv, nil)
}

// UpdateIrMinCronMixin updates an existing ir.min.cron.mixin record.
func (c *Client) UpdateIrMinCronMixin(imcm *IrMinCronMixin) error {
	return c.UpdateIrMinCronMixins([]int64{imcm.Id.Get()}, imcm)
}

// UpdateIrMinCronMixins updates existing ir.min.cron.mixin records.
// All records (represented by ids) will be updated by imcm values.
func (c *Client) UpdateIrMinCronMixins(ids []int64, imcm *IrMinCronMixin) error {
	return c.Update(IrMinCronMixinModel, ids, imcm, nil)
}

// DeleteIrMinCronMixin deletes an existing ir.min.cron.mixin record.
func (c *Client) DeleteIrMinCronMixin(id int64) error {
	return c.DeleteIrMinCronMixins([]int64{id})
}

// DeleteIrMinCronMixins deletes existing ir.min.cron.mixin records.
func (c *Client) DeleteIrMinCronMixins(ids []int64) error {
	return c.Delete(IrMinCronMixinModel, ids)
}

// GetIrMinCronMixin gets ir.min.cron.mixin existing record.
func (c *Client) GetIrMinCronMixin(id int64) (*IrMinCronMixin, error) {
	imcms, err := c.GetIrMinCronMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imcms)[0]), nil
}

// GetIrMinCronMixins gets ir.min.cron.mixin existing records.
func (c *Client) GetIrMinCronMixins(ids []int64) (*IrMinCronMixins, error) {
	imcms := &IrMinCronMixins{}
	if err := c.Read(IrMinCronMixinModel, ids, nil, imcms); err != nil {
		return nil, err
	}
	return imcms, nil
}

// FindIrMinCronMixin finds ir.min.cron.mixin record by querying it with criteria.
func (c *Client) FindIrMinCronMixin(criteria *Criteria) (*IrMinCronMixin, error) {
	imcms := &IrMinCronMixins{}
	if err := c.SearchRead(IrMinCronMixinModel, criteria, NewOptions().Limit(1), imcms); err != nil {
		return nil, err
	}
	return &((*imcms)[0]), nil
}

// FindIrMinCronMixins finds ir.min.cron.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMinCronMixins(criteria *Criteria, options *Options) (*IrMinCronMixins, error) {
	imcms := &IrMinCronMixins{}
	if err := c.SearchRead(IrMinCronMixinModel, criteria, options, imcms); err != nil {
		return nil, err
	}
	return imcms, nil
}

// FindIrMinCronMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMinCronMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrMinCronMixinModel, criteria, options)
}

// FindIrMinCronMixinId finds record id by querying it with criteria.
func (c *Client) FindIrMinCronMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrMinCronMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
