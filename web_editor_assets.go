package odoo

// WebEditorAssets represents web_editor.assets model.
type WebEditorAssets struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// WebEditorAssetss represents array of web_editor.assets model.
type WebEditorAssetss []WebEditorAssets

// WebEditorAssetsModel is the odoo model name.
const WebEditorAssetsModel = "web_editor.assets"

// Many2One convert WebEditorAssets to *Many2One.
func (wa *WebEditorAssets) Many2One() *Many2One {
	return NewMany2One(wa.Id.Get(), "")
}

// CreateWebEditorAssets creates a new web_editor.assets model and returns its id.
func (c *Client) CreateWebEditorAssets(wa *WebEditorAssets) (int64, error) {
	ids, err := c.CreateWebEditorAssetss([]*WebEditorAssets{wa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebEditorAssets creates a new web_editor.assets model and returns its id.
func (c *Client) CreateWebEditorAssetss(was []*WebEditorAssets) ([]int64, error) {
	var vv []interface{}
	for _, v := range was {
		vv = append(vv, v)
	}
	return c.Create(WebEditorAssetsModel, vv, nil)
}

// UpdateWebEditorAssets updates an existing web_editor.assets record.
func (c *Client) UpdateWebEditorAssets(wa *WebEditorAssets) error {
	return c.UpdateWebEditorAssetss([]int64{wa.Id.Get()}, wa)
}

// UpdateWebEditorAssetss updates existing web_editor.assets records.
// All records (represented by ids) will be updated by wa values.
func (c *Client) UpdateWebEditorAssetss(ids []int64, wa *WebEditorAssets) error {
	return c.Update(WebEditorAssetsModel, ids, wa, nil)
}

// DeleteWebEditorAssets deletes an existing web_editor.assets record.
func (c *Client) DeleteWebEditorAssets(id int64) error {
	return c.DeleteWebEditorAssetss([]int64{id})
}

// DeleteWebEditorAssetss deletes existing web_editor.assets records.
func (c *Client) DeleteWebEditorAssetss(ids []int64) error {
	return c.Delete(WebEditorAssetsModel, ids)
}

// GetWebEditorAssets gets web_editor.assets existing record.
func (c *Client) GetWebEditorAssets(id int64) (*WebEditorAssets, error) {
	was, err := c.GetWebEditorAssetss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*was)[0]), nil
}

// GetWebEditorAssetss gets web_editor.assets existing records.
func (c *Client) GetWebEditorAssetss(ids []int64) (*WebEditorAssetss, error) {
	was := &WebEditorAssetss{}
	if err := c.Read(WebEditorAssetsModel, ids, nil, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWebEditorAssets finds web_editor.assets record by querying it with criteria.
func (c *Client) FindWebEditorAssets(criteria *Criteria) (*WebEditorAssets, error) {
	was := &WebEditorAssetss{}
	if err := c.SearchRead(WebEditorAssetsModel, criteria, NewOptions().Limit(1), was); err != nil {
		return nil, err
	}
	return &((*was)[0]), nil
}

// FindWebEditorAssetss finds web_editor.assets records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorAssetss(criteria *Criteria, options *Options) (*WebEditorAssetss, error) {
	was := &WebEditorAssetss{}
	if err := c.SearchRead(WebEditorAssetsModel, criteria, options, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWebEditorAssetsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorAssetsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebEditorAssetsModel, criteria, options)
}

// FindWebEditorAssetsId finds record id by querying it with criteria.
func (c *Client) FindWebEditorAssetsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebEditorAssetsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
