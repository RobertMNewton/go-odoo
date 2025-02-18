package odoo

// PlanningSend represents planning.send model.
type PlanningSend struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	EmployeeIds       *Relation `xmlrpc:"employee_ids,omitempty"`
	EmployeesNoEmail  *Relation `xmlrpc:"employees_no_email,omitempty"`
	EndDatetime       *Time     `xmlrpc:"end_datetime,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	IncludeUnassigned *Bool     `xmlrpc:"include_unassigned,omitempty"`
	Note              *String   `xmlrpc:"note,omitempty"`
	SlotIds           *Relation `xmlrpc:"slot_ids,omitempty"`
	StartDatetime     *Time     `xmlrpc:"start_datetime,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PlanningSends represents array of planning.send model.
type PlanningSends []PlanningSend

// PlanningSendModel is the odoo model name.
const PlanningSendModel = "planning.send"

// Many2One convert PlanningSend to *Many2One.
func (ps *PlanningSend) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePlanningSend creates a new planning.send model and returns its id.
func (c *Client) CreatePlanningSend(ps *PlanningSend) (int64, error) {
	ids, err := c.CreatePlanningSends([]*PlanningSend{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSend creates a new planning.send model and returns its id.
func (c *Client) CreatePlanningSends(pss []*PlanningSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PlanningSendModel, vv, nil)
}

// UpdatePlanningSend updates an existing planning.send record.
func (c *Client) UpdatePlanningSend(ps *PlanningSend) error {
	return c.UpdatePlanningSends([]int64{ps.Id.Get()}, ps)
}

// UpdatePlanningSends updates existing planning.send records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePlanningSends(ids []int64, ps *PlanningSend) error {
	return c.Update(PlanningSendModel, ids, ps, nil)
}

// DeletePlanningSend deletes an existing planning.send record.
func (c *Client) DeletePlanningSend(id int64) error {
	return c.DeletePlanningSends([]int64{id})
}

// DeletePlanningSends deletes existing planning.send records.
func (c *Client) DeletePlanningSends(ids []int64) error {
	return c.Delete(PlanningSendModel, ids)
}

// GetPlanningSend gets planning.send existing record.
func (c *Client) GetPlanningSend(id int64) (*PlanningSend, error) {
	pss, err := c.GetPlanningSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// GetPlanningSends gets planning.send existing records.
func (c *Client) GetPlanningSends(ids []int64) (*PlanningSends, error) {
	pss := &PlanningSends{}
	if err := c.Read(PlanningSendModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSend finds planning.send record by querying it with criteria.
func (c *Client) FindPlanningSend(criteria *Criteria) (*PlanningSend, error) {
	pss := &PlanningSends{}
	if err := c.SearchRead(PlanningSendModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// FindPlanningSends finds planning.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSends(criteria *Criteria, options *Options) (*PlanningSends, error) {
	pss := &PlanningSends{}
	if err := c.SearchRead(PlanningSendModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningSendModel, criteria, options)
}

// FindPlanningSendId finds record id by querying it with criteria.
func (c *Client) FindPlanningSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
