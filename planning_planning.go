package odoo

// PlanningPlanning represents planning.planning model.
type PlanningPlanning struct {
	AccessToken            *String   `xmlrpc:"access_token,omitempty"`
	AllowSelfUnassign      *Bool     `xmlrpc:"allow_self_unassign,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DateEnd                *Time     `xmlrpc:"date_end,omitempty"`
	DateStart              *Time     `xmlrpc:"date_start,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	EndDatetime            *Time     `xmlrpc:"end_datetime,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	IncludeUnassigned      *Bool     `xmlrpc:"include_unassigned,omitempty"`
	SelfUnassignDaysBefore *Int      `xmlrpc:"self_unassign_days_before,omitempty"`
	StartDatetime          *Time     `xmlrpc:"start_datetime,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PlanningPlannings represents array of planning.planning model.
type PlanningPlannings []PlanningPlanning

// PlanningPlanningModel is the odoo model name.
const PlanningPlanningModel = "planning.planning"

// Many2One convert PlanningPlanning to *Many2One.
func (pp *PlanningPlanning) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePlanningPlanning creates a new planning.planning model and returns its id.
func (c *Client) CreatePlanningPlanning(pp *PlanningPlanning) (int64, error) {
	ids, err := c.CreatePlanningPlannings([]*PlanningPlanning{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningPlanning creates a new planning.planning model and returns its id.
func (c *Client) CreatePlanningPlannings(pps []*PlanningPlanning) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PlanningPlanningModel, vv, nil)
}

// UpdatePlanningPlanning updates an existing planning.planning record.
func (c *Client) UpdatePlanningPlanning(pp *PlanningPlanning) error {
	return c.UpdatePlanningPlannings([]int64{pp.Id.Get()}, pp)
}

// UpdatePlanningPlannings updates existing planning.planning records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePlanningPlannings(ids []int64, pp *PlanningPlanning) error {
	return c.Update(PlanningPlanningModel, ids, pp, nil)
}

// DeletePlanningPlanning deletes an existing planning.planning record.
func (c *Client) DeletePlanningPlanning(id int64) error {
	return c.DeletePlanningPlannings([]int64{id})
}

// DeletePlanningPlannings deletes existing planning.planning records.
func (c *Client) DeletePlanningPlannings(ids []int64) error {
	return c.Delete(PlanningPlanningModel, ids)
}

// GetPlanningPlanning gets planning.planning existing record.
func (c *Client) GetPlanningPlanning(id int64) (*PlanningPlanning, error) {
	pps, err := c.GetPlanningPlannings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetPlanningPlannings gets planning.planning existing records.
func (c *Client) GetPlanningPlannings(ids []int64) (*PlanningPlannings, error) {
	pps := &PlanningPlannings{}
	if err := c.Read(PlanningPlanningModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPlanningPlanning finds planning.planning record by querying it with criteria.
func (c *Client) FindPlanningPlanning(criteria *Criteria) (*PlanningPlanning, error) {
	pps := &PlanningPlannings{}
	if err := c.SearchRead(PlanningPlanningModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindPlanningPlannings finds planning.planning records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningPlannings(criteria *Criteria, options *Options) (*PlanningPlannings, error) {
	pps := &PlanningPlannings{}
	if err := c.SearchRead(PlanningPlanningModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPlanningPlanningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningPlanningIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningPlanningModel, criteria, options)
}

// FindPlanningPlanningId finds record id by querying it with criteria.
func (c *Client) FindPlanningPlanningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningPlanningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
