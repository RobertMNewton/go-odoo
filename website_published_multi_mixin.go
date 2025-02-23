package odoo

// WebsitePublishedMultiMixin represents website.published.multi.mixin model.
type WebsitePublishedMultiMixin struct {
	CanPublish       *Bool     `xmlrpc:"can_publish,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	IsPublished      *Bool     `xmlrpc:"is_published,omitempty"`
	WebsiteId        *Many2One `xmlrpc:"website_id,omitempty"`
	WebsitePublished *Bool     `xmlrpc:"website_published,omitempty"`
	WebsiteUrl       *String   `xmlrpc:"website_url,omitempty"`
}

// WebsitePublishedMultiMixins represents array of website.published.multi.mixin model.
type WebsitePublishedMultiMixins []WebsitePublishedMultiMixin

// WebsitePublishedMultiMixinModel is the odoo model name.
const WebsitePublishedMultiMixinModel = "website.published.multi.mixin"

// Many2One convert WebsitePublishedMultiMixin to *Many2One.
func (wpmm *WebsitePublishedMultiMixin) Many2One() *Many2One {
	return NewMany2One(wpmm.Id.Get(), "")
}

// CreateWebsitePublishedMultiMixin creates a new website.published.multi.mixin model and returns its id.
func (c *Client) CreateWebsitePublishedMultiMixin(wpmm *WebsitePublishedMultiMixin) (int64, error) {
	ids, err := c.CreateWebsitePublishedMultiMixins([]*WebsitePublishedMultiMixin{wpmm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsitePublishedMultiMixin creates a new website.published.multi.mixin model and returns its id.
func (c *Client) CreateWebsitePublishedMultiMixins(wpmms []*WebsitePublishedMultiMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range wpmms {
		vv = append(vv, v)
	}
	return c.Create(WebsitePublishedMultiMixinModel, vv, nil)
}

// UpdateWebsitePublishedMultiMixin updates an existing website.published.multi.mixin record.
func (c *Client) UpdateWebsitePublishedMultiMixin(wpmm *WebsitePublishedMultiMixin) error {
	return c.UpdateWebsitePublishedMultiMixins([]int64{wpmm.Id.Get()}, wpmm)
}

// UpdateWebsitePublishedMultiMixins updates existing website.published.multi.mixin records.
// All records (represented by ids) will be updated by wpmm values.
func (c *Client) UpdateWebsitePublishedMultiMixins(ids []int64, wpmm *WebsitePublishedMultiMixin) error {
	return c.Update(WebsitePublishedMultiMixinModel, ids, wpmm, nil)
}

// DeleteWebsitePublishedMultiMixin deletes an existing website.published.multi.mixin record.
func (c *Client) DeleteWebsitePublishedMultiMixin(id int64) error {
	return c.DeleteWebsitePublishedMultiMixins([]int64{id})
}

// DeleteWebsitePublishedMultiMixins deletes existing website.published.multi.mixin records.
func (c *Client) DeleteWebsitePublishedMultiMixins(ids []int64) error {
	return c.Delete(WebsitePublishedMultiMixinModel, ids)
}

// GetWebsitePublishedMultiMixin gets website.published.multi.mixin existing record.
func (c *Client) GetWebsitePublishedMultiMixin(id int64) (*WebsitePublishedMultiMixin, error) {
	wpmms, err := c.GetWebsitePublishedMultiMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wpmms)[0]), nil
}

// GetWebsitePublishedMultiMixins gets website.published.multi.mixin existing records.
func (c *Client) GetWebsitePublishedMultiMixins(ids []int64) (*WebsitePublishedMultiMixins, error) {
	wpmms := &WebsitePublishedMultiMixins{}
	if err := c.Read(WebsitePublishedMultiMixinModel, ids, nil, wpmms); err != nil {
		return nil, err
	}
	return wpmms, nil
}

// FindWebsitePublishedMultiMixin finds website.published.multi.mixin record by querying it with criteria.
func (c *Client) FindWebsitePublishedMultiMixin(criteria *Criteria) (*WebsitePublishedMultiMixin, error) {
	wpmms := &WebsitePublishedMultiMixins{}
	if err := c.SearchRead(WebsitePublishedMultiMixinModel, criteria, NewOptions().Limit(1), wpmms); err != nil {
		return nil, err
	}
	return &((*wpmms)[0]), nil
}

// FindWebsitePublishedMultiMixins finds website.published.multi.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMultiMixins(criteria *Criteria, options *Options) (*WebsitePublishedMultiMixins, error) {
	wpmms := &WebsitePublishedMultiMixins{}
	if err := c.SearchRead(WebsitePublishedMultiMixinModel, criteria, options, wpmms); err != nil {
		return nil, err
	}
	return wpmms, nil
}

// FindWebsitePublishedMultiMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMultiMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsitePublishedMultiMixinModel, criteria, options)
}

// FindWebsitePublishedMultiMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsitePublishedMultiMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePublishedMultiMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
