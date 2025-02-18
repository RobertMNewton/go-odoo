package odoo

// XMaterials represents x_materials model.
type XMaterials struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// XMaterialss represents array of x_materials model.
type XMaterialss []XMaterials

// XMaterialsModel is the odoo model name.
const XMaterialsModel = "x_materials"

// Many2One convert XMaterials to *Many2One.
func (x *XMaterials) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXMaterials creates a new x_materials model and returns its id.
func (c *Client) CreateXMaterials(x *XMaterials) (int64, error) {
	ids, err := c.CreateXMaterialss([]*XMaterials{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXMaterials creates a new x_materials model and returns its id.
func (c *Client) CreateXMaterialss(xs []*XMaterials) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XMaterialsModel, vv, nil)
}

// UpdateXMaterials updates an existing x_materials record.
func (c *Client) UpdateXMaterials(x *XMaterials) error {
	return c.UpdateXMaterialss([]int64{x.Id.Get()}, x)
}

// UpdateXMaterialss updates existing x_materials records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXMaterialss(ids []int64, x *XMaterials) error {
	return c.Update(XMaterialsModel, ids, x, nil)
}

// DeleteXMaterials deletes an existing x_materials record.
func (c *Client) DeleteXMaterials(id int64) error {
	return c.DeleteXMaterialss([]int64{id})
}

// DeleteXMaterialss deletes existing x_materials records.
func (c *Client) DeleteXMaterialss(ids []int64) error {
	return c.Delete(XMaterialsModel, ids)
}

// GetXMaterials gets x_materials existing record.
func (c *Client) GetXMaterials(id int64) (*XMaterials, error) {
	xs, err := c.GetXMaterialss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXMaterialss gets x_materials existing records.
func (c *Client) GetXMaterialss(ids []int64) (*XMaterialss, error) {
	xs := &XMaterialss{}
	if err := c.Read(XMaterialsModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMaterials finds x_materials record by querying it with criteria.
func (c *Client) FindXMaterials(criteria *Criteria) (*XMaterials, error) {
	xs := &XMaterialss{}
	if err := c.SearchRead(XMaterialsModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXMaterialss finds x_materials records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMaterialss(criteria *Criteria, options *Options) (*XMaterialss, error) {
	xs := &XMaterialss{}
	if err := c.SearchRead(XMaterialsModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMaterialsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMaterialsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XMaterialsModel, criteria, options)
}

// FindXMaterialsId finds record id by querying it with criteria.
func (c *Client) FindXMaterialsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XMaterialsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
