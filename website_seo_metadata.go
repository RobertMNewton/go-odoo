package odoo

// WebsiteSeoMetadata represents website.seo.metadata model.
type WebsiteSeoMetadata struct {
	DisplayName            *String `xmlrpc:"display_name,omitempty"`
	Id                     *Int    `xmlrpc:"id,omitempty"`
	IsSeoOptimized         *Bool   `xmlrpc:"is_seo_optimized,omitempty"`
	SeoName                *String `xmlrpc:"seo_name,omitempty"`
	WebsiteMetaDescription *String `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords    *String `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg       *String `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle       *String `xmlrpc:"website_meta_title,omitempty"`
}

// WebsiteSeoMetadatas represents array of website.seo.metadata model.
type WebsiteSeoMetadatas []WebsiteSeoMetadata

// WebsiteSeoMetadataModel is the odoo model name.
const WebsiteSeoMetadataModel = "website.seo.metadata"

// Many2One convert WebsiteSeoMetadata to *Many2One.
func (wsm *WebsiteSeoMetadata) Many2One() *Many2One {
	return NewMany2One(wsm.Id.Get(), "")
}

// CreateWebsiteSeoMetadata creates a new website.seo.metadata model and returns its id.
func (c *Client) CreateWebsiteSeoMetadata(wsm *WebsiteSeoMetadata) (int64, error) {
	ids, err := c.CreateWebsiteSeoMetadatas([]*WebsiteSeoMetadata{wsm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteSeoMetadata creates a new website.seo.metadata model and returns its id.
func (c *Client) CreateWebsiteSeoMetadatas(wsms []*WebsiteSeoMetadata) ([]int64, error) {
	var vv []interface{}
	for _, v := range wsms {
		vv = append(vv, v)
	}
	return c.Create(WebsiteSeoMetadataModel, vv, nil)
}

// UpdateWebsiteSeoMetadata updates an existing website.seo.metadata record.
func (c *Client) UpdateWebsiteSeoMetadata(wsm *WebsiteSeoMetadata) error {
	return c.UpdateWebsiteSeoMetadatas([]int64{wsm.Id.Get()}, wsm)
}

// UpdateWebsiteSeoMetadatas updates existing website.seo.metadata records.
// All records (represented by ids) will be updated by wsm values.
func (c *Client) UpdateWebsiteSeoMetadatas(ids []int64, wsm *WebsiteSeoMetadata) error {
	return c.Update(WebsiteSeoMetadataModel, ids, wsm, nil)
}

// DeleteWebsiteSeoMetadata deletes an existing website.seo.metadata record.
func (c *Client) DeleteWebsiteSeoMetadata(id int64) error {
	return c.DeleteWebsiteSeoMetadatas([]int64{id})
}

// DeleteWebsiteSeoMetadatas deletes existing website.seo.metadata records.
func (c *Client) DeleteWebsiteSeoMetadatas(ids []int64) error {
	return c.Delete(WebsiteSeoMetadataModel, ids)
}

// GetWebsiteSeoMetadata gets website.seo.metadata existing record.
func (c *Client) GetWebsiteSeoMetadata(id int64) (*WebsiteSeoMetadata, error) {
	wsms, err := c.GetWebsiteSeoMetadatas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wsms)[0]), nil
}

// GetWebsiteSeoMetadatas gets website.seo.metadata existing records.
func (c *Client) GetWebsiteSeoMetadatas(ids []int64) (*WebsiteSeoMetadatas, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.Read(WebsiteSeoMetadataModel, ids, nil, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSeoMetadata finds website.seo.metadata record by querying it with criteria.
func (c *Client) FindWebsiteSeoMetadata(criteria *Criteria) (*WebsiteSeoMetadata, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.SearchRead(WebsiteSeoMetadataModel, criteria, NewOptions().Limit(1), wsms); err != nil {
		return nil, err
	}
	return &((*wsms)[0]), nil
}

// FindWebsiteSeoMetadatas finds website.seo.metadata records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSeoMetadatas(criteria *Criteria, options *Options) (*WebsiteSeoMetadatas, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.SearchRead(WebsiteSeoMetadataModel, criteria, options, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSeoMetadataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSeoMetadataIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteSeoMetadataModel, criteria, options)
}

// FindWebsiteSeoMetadataId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSeoMetadataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSeoMetadataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
