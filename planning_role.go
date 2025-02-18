package odoo

// PlanningRole represents planning.role model.
type PlanningRole struct {
	Active                   *Bool       `xmlrpc:"active,omitempty"`
	Color                    *Int        `xmlrpc:"color,omitempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	Name                     *String     `xmlrpc:"name,omitempty"`
	ProductIds               *Relation   `xmlrpc:"product_ids,omitempty"`
	ResourceIds              *Relation   `xmlrpc:"resource_ids,omitempty"`
	Sequence                 *Int        `xmlrpc:"sequence,omitempty"`
	SlotPropertiesDefinition interface{} `xmlrpc:"slot_properties_definition,omitempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PlanningRoles represents array of planning.role model.
type PlanningRoles []PlanningRole

// PlanningRoleModel is the odoo model name.
const PlanningRoleModel = "planning.role"

// Many2One convert PlanningRole to *Many2One.
func (pr *PlanningRole) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePlanningRole creates a new planning.role model and returns its id.
func (c *Client) CreatePlanningRole(pr *PlanningRole) (int64, error) {
	ids, err := c.CreatePlanningRoles([]*PlanningRole{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningRole creates a new planning.role model and returns its id.
func (c *Client) CreatePlanningRoles(prs []*PlanningRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(PlanningRoleModel, vv, nil)
}

// UpdatePlanningRole updates an existing planning.role record.
func (c *Client) UpdatePlanningRole(pr *PlanningRole) error {
	return c.UpdatePlanningRoles([]int64{pr.Id.Get()}, pr)
}

// UpdatePlanningRoles updates existing planning.role records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePlanningRoles(ids []int64, pr *PlanningRole) error {
	return c.Update(PlanningRoleModel, ids, pr, nil)
}

// DeletePlanningRole deletes an existing planning.role record.
func (c *Client) DeletePlanningRole(id int64) error {
	return c.DeletePlanningRoles([]int64{id})
}

// DeletePlanningRoles deletes existing planning.role records.
func (c *Client) DeletePlanningRoles(ids []int64) error {
	return c.Delete(PlanningRoleModel, ids)
}

// GetPlanningRole gets planning.role existing record.
func (c *Client) GetPlanningRole(id int64) (*PlanningRole, error) {
	prs, err := c.GetPlanningRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetPlanningRoles gets planning.role existing records.
func (c *Client) GetPlanningRoles(ids []int64) (*PlanningRoles, error) {
	prs := &PlanningRoles{}
	if err := c.Read(PlanningRoleModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPlanningRole finds planning.role record by querying it with criteria.
func (c *Client) FindPlanningRole(criteria *Criteria) (*PlanningRole, error) {
	prs := &PlanningRoles{}
	if err := c.SearchRead(PlanningRoleModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindPlanningRoles finds planning.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningRoles(criteria *Criteria, options *Options) (*PlanningRoles, error) {
	prs := &PlanningRoles{}
	if err := c.SearchRead(PlanningRoleModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPlanningRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningRoleModel, criteria, options)
}

// FindPlanningRoleId finds record id by querying it with criteria.
func (c *Client) FindPlanningRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
