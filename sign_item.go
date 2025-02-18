package odoo

// SignItem represents sign.item model.
type SignItem struct {
	Alignment     *String   `xmlrpc:"alignment,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Height        *Float    `xmlrpc:"height,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	NumOptions    *Int      `xmlrpc:"num_options,omitempty"`
	OptionIds     *Relation `xmlrpc:"option_ids,omitempty"`
	Page          *Int      `xmlrpc:"page,omitempty"`
	PosX          *Float    `xmlrpc:"posX,omitempty"`
	PosY          *Float    `xmlrpc:"posY,omitempty"`
	RadioSetId    *Many2One `xmlrpc:"radio_set_id,omitempty"`
	Required      *Bool     `xmlrpc:"required,omitempty"`
	ResponsibleId *Many2One `xmlrpc:"responsible_id,omitempty"`
	TemplateId    *Many2One `xmlrpc:"template_id,omitempty"`
	TransactionId *Int      `xmlrpc:"transaction_id,omitempty"`
	TypeId        *Many2One `xmlrpc:"type_id,omitempty"`
	Width         *Float    `xmlrpc:"width,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignItems represents array of sign.item model.
type SignItems []SignItem

// SignItemModel is the odoo model name.
const SignItemModel = "sign.item"

// Many2One convert SignItem to *Many2One.
func (si *SignItem) Many2One() *Many2One {
	return NewMany2One(si.Id.Get(), "")
}

// CreateSignItem creates a new sign.item model and returns its id.
func (c *Client) CreateSignItem(si *SignItem) (int64, error) {
	ids, err := c.CreateSignItems([]*SignItem{si})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignItem creates a new sign.item model and returns its id.
func (c *Client) CreateSignItems(sis []*SignItem) ([]int64, error) {
	var vv []interface{}
	for _, v := range sis {
		vv = append(vv, v)
	}
	return c.Create(SignItemModel, vv, nil)
}

// UpdateSignItem updates an existing sign.item record.
func (c *Client) UpdateSignItem(si *SignItem) error {
	return c.UpdateSignItems([]int64{si.Id.Get()}, si)
}

// UpdateSignItems updates existing sign.item records.
// All records (represented by ids) will be updated by si values.
func (c *Client) UpdateSignItems(ids []int64, si *SignItem) error {
	return c.Update(SignItemModel, ids, si, nil)
}

// DeleteSignItem deletes an existing sign.item record.
func (c *Client) DeleteSignItem(id int64) error {
	return c.DeleteSignItems([]int64{id})
}

// DeleteSignItems deletes existing sign.item records.
func (c *Client) DeleteSignItems(ids []int64) error {
	return c.Delete(SignItemModel, ids)
}

// GetSignItem gets sign.item existing record.
func (c *Client) GetSignItem(id int64) (*SignItem, error) {
	sis, err := c.GetSignItems([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sis)[0]), nil
}

// GetSignItems gets sign.item existing records.
func (c *Client) GetSignItems(ids []int64) (*SignItems, error) {
	sis := &SignItems{}
	if err := c.Read(SignItemModel, ids, nil, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSignItem finds sign.item record by querying it with criteria.
func (c *Client) FindSignItem(criteria *Criteria) (*SignItem, error) {
	sis := &SignItems{}
	if err := c.SearchRead(SignItemModel, criteria, NewOptions().Limit(1), sis); err != nil {
		return nil, err
	}
	return &((*sis)[0]), nil
}

// FindSignItems finds sign.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItems(criteria *Criteria, options *Options) (*SignItems, error) {
	sis := &SignItems{}
	if err := c.SearchRead(SignItemModel, criteria, options, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSignItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignItemModel, criteria, options)
}

// FindSignItemId finds record id by querying it with criteria.
func (c *Client) FindSignItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
