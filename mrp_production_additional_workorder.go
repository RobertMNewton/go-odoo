package odoo

// MrpProductionAdditionalWorkorder represents mrp_production.additional.workorder model.
type MrpProductionAdditionalWorkorder struct {
	BlockedByWorkorderId *Many2One `xmlrpc:"blocked_by_workorder_id,omitempty"`
	CompanyId            *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DateStart            *Time     `xmlrpc:"date_start,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	DurationExpected     *Float    `xmlrpc:"duration_expected,omitempty"`
	EmployeeAssignedIds  *Relation `xmlrpc:"employee_assigned_ids,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	ProductionId         *Many2One `xmlrpc:"production_id,omitempty"`
	WorkcenterId         *Many2One `xmlrpc:"workcenter_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionAdditionalWorkorders represents array of mrp_production.additional.workorder model.
type MrpProductionAdditionalWorkorders []MrpProductionAdditionalWorkorder

// MrpProductionAdditionalWorkorderModel is the odoo model name.
const MrpProductionAdditionalWorkorderModel = "mrp_production.additional.workorder"

// Many2One convert MrpProductionAdditionalWorkorder to *Many2One.
func (maw *MrpProductionAdditionalWorkorder) Many2One() *Many2One {
	return NewMany2One(maw.Id.Get(), "")
}

// CreateMrpProductionAdditionalWorkorder creates a new mrp_production.additional.workorder model and returns its id.
func (c *Client) CreateMrpProductionAdditionalWorkorder(maw *MrpProductionAdditionalWorkorder) (int64, error) {
	ids, err := c.CreateMrpProductionAdditionalWorkorders([]*MrpProductionAdditionalWorkorder{maw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionAdditionalWorkorder creates a new mrp_production.additional.workorder model and returns its id.
func (c *Client) CreateMrpProductionAdditionalWorkorders(maws []*MrpProductionAdditionalWorkorder) ([]int64, error) {
	var vv []interface{}
	for _, v := range maws {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionAdditionalWorkorderModel, vv, nil)
}

// UpdateMrpProductionAdditionalWorkorder updates an existing mrp_production.additional.workorder record.
func (c *Client) UpdateMrpProductionAdditionalWorkorder(maw *MrpProductionAdditionalWorkorder) error {
	return c.UpdateMrpProductionAdditionalWorkorders([]int64{maw.Id.Get()}, maw)
}

// UpdateMrpProductionAdditionalWorkorders updates existing mrp_production.additional.workorder records.
// All records (represented by ids) will be updated by maw values.
func (c *Client) UpdateMrpProductionAdditionalWorkorders(ids []int64, maw *MrpProductionAdditionalWorkorder) error {
	return c.Update(MrpProductionAdditionalWorkorderModel, ids, maw, nil)
}

// DeleteMrpProductionAdditionalWorkorder deletes an existing mrp_production.additional.workorder record.
func (c *Client) DeleteMrpProductionAdditionalWorkorder(id int64) error {
	return c.DeleteMrpProductionAdditionalWorkorders([]int64{id})
}

// DeleteMrpProductionAdditionalWorkorders deletes existing mrp_production.additional.workorder records.
func (c *Client) DeleteMrpProductionAdditionalWorkorders(ids []int64) error {
	return c.Delete(MrpProductionAdditionalWorkorderModel, ids)
}

// GetMrpProductionAdditionalWorkorder gets mrp_production.additional.workorder existing record.
func (c *Client) GetMrpProductionAdditionalWorkorder(id int64) (*MrpProductionAdditionalWorkorder, error) {
	maws, err := c.GetMrpProductionAdditionalWorkorders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*maws)[0]), nil
}

// GetMrpProductionAdditionalWorkorders gets mrp_production.additional.workorder existing records.
func (c *Client) GetMrpProductionAdditionalWorkorders(ids []int64) (*MrpProductionAdditionalWorkorders, error) {
	maws := &MrpProductionAdditionalWorkorders{}
	if err := c.Read(MrpProductionAdditionalWorkorderModel, ids, nil, maws); err != nil {
		return nil, err
	}
	return maws, nil
}

// FindMrpProductionAdditionalWorkorder finds mrp_production.additional.workorder record by querying it with criteria.
func (c *Client) FindMrpProductionAdditionalWorkorder(criteria *Criteria) (*MrpProductionAdditionalWorkorder, error) {
	maws := &MrpProductionAdditionalWorkorders{}
	if err := c.SearchRead(MrpProductionAdditionalWorkorderModel, criteria, NewOptions().Limit(1), maws); err != nil {
		return nil, err
	}
	return &((*maws)[0]), nil
}

// FindMrpProductionAdditionalWorkorders finds mrp_production.additional.workorder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionAdditionalWorkorders(criteria *Criteria, options *Options) (*MrpProductionAdditionalWorkorders, error) {
	maws := &MrpProductionAdditionalWorkorders{}
	if err := c.SearchRead(MrpProductionAdditionalWorkorderModel, criteria, options, maws); err != nil {
		return nil, err
	}
	return maws, nil
}

// FindMrpProductionAdditionalWorkorderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionAdditionalWorkorderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionAdditionalWorkorderModel, criteria, options)
}

// FindMrpProductionAdditionalWorkorderId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionAdditionalWorkorderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionAdditionalWorkorderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
