package odoo

// PlanningRecurrency represents planning.recurrency model.
type PlanningRecurrency struct {
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	LastGeneratedEndDatetime *Time      `xmlrpc:"last_generated_end_datetime,omitempty"`
	RepeatInterval           *Int       `xmlrpc:"repeat_interval,omitempty"`
	RepeatNumber             *Int       `xmlrpc:"repeat_number,omitempty"`
	RepeatType               *Selection `xmlrpc:"repeat_type,omitempty"`
	RepeatUnit               *Selection `xmlrpc:"repeat_unit,omitempty"`
	RepeatUntil              *Time      `xmlrpc:"repeat_until,omitempty"`
	SlotIds                  *Relation  `xmlrpc:"slot_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PlanningRecurrencys represents array of planning.recurrency model.
type PlanningRecurrencys []PlanningRecurrency

// PlanningRecurrencyModel is the odoo model name.
const PlanningRecurrencyModel = "planning.recurrency"

// Many2One convert PlanningRecurrency to *Many2One.
func (pr *PlanningRecurrency) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePlanningRecurrency creates a new planning.recurrency model and returns its id.
func (c *Client) CreatePlanningRecurrency(pr *PlanningRecurrency) (int64, error) {
	ids, err := c.CreatePlanningRecurrencys([]*PlanningRecurrency{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningRecurrency creates a new planning.recurrency model and returns its id.
func (c *Client) CreatePlanningRecurrencys(prs []*PlanningRecurrency) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(PlanningRecurrencyModel, vv, nil)
}

// UpdatePlanningRecurrency updates an existing planning.recurrency record.
func (c *Client) UpdatePlanningRecurrency(pr *PlanningRecurrency) error {
	return c.UpdatePlanningRecurrencys([]int64{pr.Id.Get()}, pr)
}

// UpdatePlanningRecurrencys updates existing planning.recurrency records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePlanningRecurrencys(ids []int64, pr *PlanningRecurrency) error {
	return c.Update(PlanningRecurrencyModel, ids, pr, nil)
}

// DeletePlanningRecurrency deletes an existing planning.recurrency record.
func (c *Client) DeletePlanningRecurrency(id int64) error {
	return c.DeletePlanningRecurrencys([]int64{id})
}

// DeletePlanningRecurrencys deletes existing planning.recurrency records.
func (c *Client) DeletePlanningRecurrencys(ids []int64) error {
	return c.Delete(PlanningRecurrencyModel, ids)
}

// GetPlanningRecurrency gets planning.recurrency existing record.
func (c *Client) GetPlanningRecurrency(id int64) (*PlanningRecurrency, error) {
	prs, err := c.GetPlanningRecurrencys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetPlanningRecurrencys gets planning.recurrency existing records.
func (c *Client) GetPlanningRecurrencys(ids []int64) (*PlanningRecurrencys, error) {
	prs := &PlanningRecurrencys{}
	if err := c.Read(PlanningRecurrencyModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPlanningRecurrency finds planning.recurrency record by querying it with criteria.
func (c *Client) FindPlanningRecurrency(criteria *Criteria) (*PlanningRecurrency, error) {
	prs := &PlanningRecurrencys{}
	if err := c.SearchRead(PlanningRecurrencyModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindPlanningRecurrencys finds planning.recurrency records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningRecurrencys(criteria *Criteria, options *Options) (*PlanningRecurrencys, error) {
	prs := &PlanningRecurrencys{}
	if err := c.SearchRead(PlanningRecurrencyModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPlanningRecurrencyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningRecurrencyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningRecurrencyModel, criteria, options)
}

// FindPlanningRecurrencyId finds record id by querying it with criteria.
func (c *Client) FindPlanningRecurrencyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningRecurrencyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
