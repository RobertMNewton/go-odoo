package odoo

// ProposeChange represents propose.change model.
type ProposeChange struct {
	ChangeType  *Selection `xmlrpc:"change_type,omitempty"`
	Comment     *String    `xmlrpc:"comment,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Note        *String    `xmlrpc:"note,omitempty"`
	Picture     *String    `xmlrpc:"picture,omitempty"`
	StepId      *Many2One  `xmlrpc:"step_id,omitempty"`
	Title       *String    `xmlrpc:"title,omitempty"`
	WorkorderId *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProposeChanges represents array of propose.change model.
type ProposeChanges []ProposeChange

// ProposeChangeModel is the odoo model name.
const ProposeChangeModel = "propose.change"

// Many2One convert ProposeChange to *Many2One.
func (pc *ProposeChange) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreateProposeChange creates a new propose.change model and returns its id.
func (c *Client) CreateProposeChange(pc *ProposeChange) (int64, error) {
	ids, err := c.CreateProposeChanges([]*ProposeChange{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProposeChange creates a new propose.change model and returns its id.
func (c *Client) CreateProposeChanges(pcs []*ProposeChange) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(ProposeChangeModel, vv, nil)
}

// UpdateProposeChange updates an existing propose.change record.
func (c *Client) UpdateProposeChange(pc *ProposeChange) error {
	return c.UpdateProposeChanges([]int64{pc.Id.Get()}, pc)
}

// UpdateProposeChanges updates existing propose.change records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdateProposeChanges(ids []int64, pc *ProposeChange) error {
	return c.Update(ProposeChangeModel, ids, pc, nil)
}

// DeleteProposeChange deletes an existing propose.change record.
func (c *Client) DeleteProposeChange(id int64) error {
	return c.DeleteProposeChanges([]int64{id})
}

// DeleteProposeChanges deletes existing propose.change records.
func (c *Client) DeleteProposeChanges(ids []int64) error {
	return c.Delete(ProposeChangeModel, ids)
}

// GetProposeChange gets propose.change existing record.
func (c *Client) GetProposeChange(id int64) (*ProposeChange, error) {
	pcs, err := c.GetProposeChanges([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetProposeChanges gets propose.change existing records.
func (c *Client) GetProposeChanges(ids []int64) (*ProposeChanges, error) {
	pcs := &ProposeChanges{}
	if err := c.Read(ProposeChangeModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProposeChange finds propose.change record by querying it with criteria.
func (c *Client) FindProposeChange(criteria *Criteria) (*ProposeChange, error) {
	pcs := &ProposeChanges{}
	if err := c.SearchRead(ProposeChangeModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindProposeChanges finds propose.change records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProposeChanges(criteria *Criteria, options *Options) (*ProposeChanges, error) {
	pcs := &ProposeChanges{}
	if err := c.SearchRead(ProposeChangeModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProposeChangeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProposeChangeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProposeChangeModel, criteria, options)
}

// FindProposeChangeId finds record id by querying it with criteria.
func (c *Client) FindProposeChangeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProposeChangeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
