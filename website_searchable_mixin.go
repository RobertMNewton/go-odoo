package odoo

// WebsiteSearchableMixin represents website.searchable.mixin model.
type WebsiteSearchableMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// WebsiteSearchableMixins represents array of website.searchable.mixin model.
type WebsiteSearchableMixins []WebsiteSearchableMixin

// WebsiteSearchableMixinModel is the odoo model name.
const WebsiteSearchableMixinModel = "website.searchable.mixin"

// Many2One convert WebsiteSearchableMixin to *Many2One.
func (wsm *WebsiteSearchableMixin) Many2One() *Many2One {
	return NewMany2One(wsm.Id.Get(), "")
}

// CreateWebsiteSearchableMixin creates a new website.searchable.mixin model and returns its id.
func (c *Client) CreateWebsiteSearchableMixin(wsm *WebsiteSearchableMixin) (int64, error) {
	ids, err := c.CreateWebsiteSearchableMixins([]*WebsiteSearchableMixin{wsm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteSearchableMixin creates a new website.searchable.mixin model and returns its id.
func (c *Client) CreateWebsiteSearchableMixins(wsms []*WebsiteSearchableMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range wsms {
		vv = append(vv, v)
	}
	return c.Create(WebsiteSearchableMixinModel, vv, nil)
}

// UpdateWebsiteSearchableMixin updates an existing website.searchable.mixin record.
func (c *Client) UpdateWebsiteSearchableMixin(wsm *WebsiteSearchableMixin) error {
	return c.UpdateWebsiteSearchableMixins([]int64{wsm.Id.Get()}, wsm)
}

// UpdateWebsiteSearchableMixins updates existing website.searchable.mixin records.
// All records (represented by ids) will be updated by wsm values.
func (c *Client) UpdateWebsiteSearchableMixins(ids []int64, wsm *WebsiteSearchableMixin) error {
	return c.Update(WebsiteSearchableMixinModel, ids, wsm, nil)
}

// DeleteWebsiteSearchableMixin deletes an existing website.searchable.mixin record.
func (c *Client) DeleteWebsiteSearchableMixin(id int64) error {
	return c.DeleteWebsiteSearchableMixins([]int64{id})
}

// DeleteWebsiteSearchableMixins deletes existing website.searchable.mixin records.
func (c *Client) DeleteWebsiteSearchableMixins(ids []int64) error {
	return c.Delete(WebsiteSearchableMixinModel, ids)
}

// GetWebsiteSearchableMixin gets website.searchable.mixin existing record.
func (c *Client) GetWebsiteSearchableMixin(id int64) (*WebsiteSearchableMixin, error) {
	wsms, err := c.GetWebsiteSearchableMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wsms)[0]), nil
}

// GetWebsiteSearchableMixins gets website.searchable.mixin existing records.
func (c *Client) GetWebsiteSearchableMixins(ids []int64) (*WebsiteSearchableMixins, error) {
	wsms := &WebsiteSearchableMixins{}
	if err := c.Read(WebsiteSearchableMixinModel, ids, nil, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSearchableMixin finds website.searchable.mixin record by querying it with criteria.
func (c *Client) FindWebsiteSearchableMixin(criteria *Criteria) (*WebsiteSearchableMixin, error) {
	wsms := &WebsiteSearchableMixins{}
	if err := c.SearchRead(WebsiteSearchableMixinModel, criteria, NewOptions().Limit(1), wsms); err != nil {
		return nil, err
	}
	return &((*wsms)[0]), nil
}

// FindWebsiteSearchableMixins finds website.searchable.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSearchableMixins(criteria *Criteria, options *Options) (*WebsiteSearchableMixins, error) {
	wsms := &WebsiteSearchableMixins{}
	if err := c.SearchRead(WebsiteSearchableMixinModel, criteria, options, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSearchableMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSearchableMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteSearchableMixinModel, criteria, options)
}

// FindWebsiteSearchableMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSearchableMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSearchableMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
