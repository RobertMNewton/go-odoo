package odoo

// SignItemRole represents sign.item.role model.
type SignItemRole struct {
	AuthMethod       *Selection `xmlrpc:"auth_method,omitempty"`
	ChangeAuthorized *Bool      `xmlrpc:"change_authorized,omitempty"`
	Color            *Int       `xmlrpc:"color,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Default          *Bool      `xmlrpc:"default,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	Sequence         *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SignItemRoles represents array of sign.item.role model.
type SignItemRoles []SignItemRole

// SignItemRoleModel is the odoo model name.
const SignItemRoleModel = "sign.item.role"

// Many2One convert SignItemRole to *Many2One.
func (sir *SignItemRole) Many2One() *Many2One {
	return NewMany2One(sir.Id.Get(), "")
}

// CreateSignItemRole creates a new sign.item.role model and returns its id.
func (c *Client) CreateSignItemRole(sir *SignItemRole) (int64, error) {
	ids, err := c.CreateSignItemRoles([]*SignItemRole{sir})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignItemRole creates a new sign.item.role model and returns its id.
func (c *Client) CreateSignItemRoles(sirs []*SignItemRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range sirs {
		vv = append(vv, v)
	}
	return c.Create(SignItemRoleModel, vv, nil)
}

// UpdateSignItemRole updates an existing sign.item.role record.
func (c *Client) UpdateSignItemRole(sir *SignItemRole) error {
	return c.UpdateSignItemRoles([]int64{sir.Id.Get()}, sir)
}

// UpdateSignItemRoles updates existing sign.item.role records.
// All records (represented by ids) will be updated by sir values.
func (c *Client) UpdateSignItemRoles(ids []int64, sir *SignItemRole) error {
	return c.Update(SignItemRoleModel, ids, sir, nil)
}

// DeleteSignItemRole deletes an existing sign.item.role record.
func (c *Client) DeleteSignItemRole(id int64) error {
	return c.DeleteSignItemRoles([]int64{id})
}

// DeleteSignItemRoles deletes existing sign.item.role records.
func (c *Client) DeleteSignItemRoles(ids []int64) error {
	return c.Delete(SignItemRoleModel, ids)
}

// GetSignItemRole gets sign.item.role existing record.
func (c *Client) GetSignItemRole(id int64) (*SignItemRole, error) {
	sirs, err := c.GetSignItemRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sirs)[0]), nil
}

// GetSignItemRoles gets sign.item.role existing records.
func (c *Client) GetSignItemRoles(ids []int64) (*SignItemRoles, error) {
	sirs := &SignItemRoles{}
	if err := c.Read(SignItemRoleModel, ids, nil, sirs); err != nil {
		return nil, err
	}
	return sirs, nil
}

// FindSignItemRole finds sign.item.role record by querying it with criteria.
func (c *Client) FindSignItemRole(criteria *Criteria) (*SignItemRole, error) {
	sirs := &SignItemRoles{}
	if err := c.SearchRead(SignItemRoleModel, criteria, NewOptions().Limit(1), sirs); err != nil {
		return nil, err
	}
	return &((*sirs)[0]), nil
}

// FindSignItemRoles finds sign.item.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemRoles(criteria *Criteria, options *Options) (*SignItemRoles, error) {
	sirs := &SignItemRoles{}
	if err := c.SearchRead(SignItemRoleModel, criteria, options, sirs); err != nil {
		return nil, err
	}
	return sirs, nil
}

// FindSignItemRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignItemRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignItemRoleModel, criteria, options)
}

// FindSignItemRoleId finds record id by querying it with criteria.
func (c *Client) FindSignItemRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignItemRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
