package odoo

// PlanningSlotTemplate represents planning.slot.template model.
type PlanningSlotTemplate struct {
	Active       *Bool     `xmlrpc:"active,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	DurationDays *Int      `xmlrpc:"duration_days,omitempty"`
	EndTime      *Float    `xmlrpc:"end_time,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	RoleId       *Many2One `xmlrpc:"role_id,omitempty"`
	Sequence     *Int      `xmlrpc:"sequence,omitempty"`
	StartTime    *Float    `xmlrpc:"start_time,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PlanningSlotTemplates represents array of planning.slot.template model.
type PlanningSlotTemplates []PlanningSlotTemplate

// PlanningSlotTemplateModel is the odoo model name.
const PlanningSlotTemplateModel = "planning.slot.template"

// Many2One convert PlanningSlotTemplate to *Many2One.
func (pst *PlanningSlotTemplate) Many2One() *Many2One {
	return NewMany2One(pst.Id.Get(), "")
}

// CreatePlanningSlotTemplate creates a new planning.slot.template model and returns its id.
func (c *Client) CreatePlanningSlotTemplate(pst *PlanningSlotTemplate) (int64, error) {
	ids, err := c.CreatePlanningSlotTemplates([]*PlanningSlotTemplate{pst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSlotTemplate creates a new planning.slot.template model and returns its id.
func (c *Client) CreatePlanningSlotTemplates(psts []*PlanningSlotTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range psts {
		vv = append(vv, v)
	}
	return c.Create(PlanningSlotTemplateModel, vv, nil)
}

// UpdatePlanningSlotTemplate updates an existing planning.slot.template record.
func (c *Client) UpdatePlanningSlotTemplate(pst *PlanningSlotTemplate) error {
	return c.UpdatePlanningSlotTemplates([]int64{pst.Id.Get()}, pst)
}

// UpdatePlanningSlotTemplates updates existing planning.slot.template records.
// All records (represented by ids) will be updated by pst values.
func (c *Client) UpdatePlanningSlotTemplates(ids []int64, pst *PlanningSlotTemplate) error {
	return c.Update(PlanningSlotTemplateModel, ids, pst, nil)
}

// DeletePlanningSlotTemplate deletes an existing planning.slot.template record.
func (c *Client) DeletePlanningSlotTemplate(id int64) error {
	return c.DeletePlanningSlotTemplates([]int64{id})
}

// DeletePlanningSlotTemplates deletes existing planning.slot.template records.
func (c *Client) DeletePlanningSlotTemplates(ids []int64) error {
	return c.Delete(PlanningSlotTemplateModel, ids)
}

// GetPlanningSlotTemplate gets planning.slot.template existing record.
func (c *Client) GetPlanningSlotTemplate(id int64) (*PlanningSlotTemplate, error) {
	psts, err := c.GetPlanningSlotTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*psts)[0]), nil
}

// GetPlanningSlotTemplates gets planning.slot.template existing records.
func (c *Client) GetPlanningSlotTemplates(ids []int64) (*PlanningSlotTemplates, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.Read(PlanningSlotTemplateModel, ids, nil, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindPlanningSlotTemplate finds planning.slot.template record by querying it with criteria.
func (c *Client) FindPlanningSlotTemplate(criteria *Criteria) (*PlanningSlotTemplate, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.SearchRead(PlanningSlotTemplateModel, criteria, NewOptions().Limit(1), psts); err != nil {
		return nil, err
	}
	return &((*psts)[0]), nil
}

// FindPlanningSlotTemplates finds planning.slot.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotTemplates(criteria *Criteria, options *Options) (*PlanningSlotTemplates, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.SearchRead(PlanningSlotTemplateModel, criteria, options, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindPlanningSlotTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningSlotTemplateModel, criteria, options)
}

// FindPlanningSlotTemplateId finds record id by querying it with criteria.
func (c *Client) FindPlanningSlotTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSlotTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
