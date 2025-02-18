package odoo

// SignItemType represents sign.item.type model.
type SignItemType struct {
	AutoField     *String    `xmlrpc:"auto_field,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultHeight *Float     `xmlrpc:"default_height,omitempty"`
	DefaultWidth  *Float     `xmlrpc:"default_width,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Icon          *String    `xmlrpc:"icon,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	ItemType      *Selection `xmlrpc:"item_type,omitempty"`
	ModelId       *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelName     *String    `xmlrpc:"model_name,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	Placeholder   *String    `xmlrpc:"placeholder,omitempty"`
	Tip           *String    `xmlrpc:"tip,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SignItemTypes represents array of sign.item.type model.
type SignItemTypes []SignItemType

// SignItemTypeModel is the odoo model name.
const SignItemTypeModel = "sign.item.type"

// Many2One convert SignItemType to *Many2One.
func (sit *SignItemType) Many2One() *Many2One {
	return NewMany2One(sit.Id.Get(), "")
}

// CreateSignItemType creates a new sign.item.type model and returns its id.
func (c *Client) CreateSignItemType(sit *SignItemType) (int64, error) {
	ids, err := c.CreateSignItemTypes([]*SignItemType{sit})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignItemType creates a new sign.item.type model and returns its id.
func (c *Client) CreateSignItemTypes(sits []*SignItemType) ([]int64, error) {
	var vv []interface{}
	for _, v := range sits {
		vv = append(vv, v)
	}
	return c.Create(SignItemTypeModel, vv, nil)
}

// UpdateSignItemType updates an existing sign.item.type record.
func (c *Client) UpdateSignItemType(sit *SignItemType) error {
	return c.UpdateSignItemTypes([]int64{sit.Id.Get()}, sit)
}

// UpdateSignItemTypes updates existing sign.item.type records.
// All records (represented by ids) will be updated by sit values.
func (c *Client) UpdateSignItemTypes(ids []int64, sit *SignItemType) error {
	return c.Update(SignItemTypeModel, ids, sit, nil)
}

// DeleteSignItemType deletes an existing sign.item.type record.
func (c *Client) DeleteSignItemType(id int64) error {
	return c.DeleteSignItemTypes([]int64{id})
}

// DeleteSignItemTypes deletes existing sign.item.type records.
func (c *Client) DeleteSignItemTypes(ids []int64) error {
	return c.Delete(SignItemTypeModel, ids)
}

// GetSignItemType gets sign.item.type existing record.
func (c *Client) GetSignItemType(id int64) (*SignItemType, error) {
	sits, err := c.GetSignItemTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sits)[0]), nil
}

// GetSignItemTypes gets sign.item.type existing records.
func (c *Client) GetSignItemTypes(ids []int64) (*SignItemTypes, error) {
	sits := &SignItemTypes{}
	if err := c.Read(SignItemTypeModel, ids, nil, sits); err != nil {
		return nil, err
	}
	return sits, nil
}

// FindSignItemType finds sign.item.type record by querying it with criteria.
func (c *Client) FindSignItemType(criteria *Criteria) (*SignItemType, error) {
	sits := &SignItemTypes{}
	if err := c.SearchRead(SignItemTypeModel, criteria, NewOptions().Limit(1), sits); err != nil {
		return nil, err
	}
	return &((*sits)[0]), nil
}

// FindSignItemTypes finds sign.item.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemTypes(criteria *Criteria, options *Options) (*SignItemTypes, error) {
	sits := &SignItemTypes{}
	if err := c.SearchRead(SignItemTypeModel, criteria, options, sits); err != nil {
		return nil, err
	}
	return sits, nil
}

// FindSignItemTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignItemTypeModel, criteria, options)
}

// FindSignItemTypeId finds record id by querying it with criteria.
func (c *Client) FindSignItemTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignItemTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
