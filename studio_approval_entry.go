package odoo

// StudioApprovalEntry represents studio.approval.entry model.
type StudioApprovalEntry struct {
	ActionId    *Many2One `xmlrpc:"action_id,omitempty"`
	Approved    *Bool     `xmlrpc:"approved,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Method      *String   `xmlrpc:"method,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Reference   *String   `xmlrpc:"reference,omitempty"`
	ResId       *Many2One `xmlrpc:"res_id,omitempty"`
	RuleId      *Many2One `xmlrpc:"rule_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioApprovalEntrys represents array of studio.approval.entry model.
type StudioApprovalEntrys []StudioApprovalEntry

// StudioApprovalEntryModel is the odoo model name.
const StudioApprovalEntryModel = "studio.approval.entry"

// Many2One convert StudioApprovalEntry to *Many2One.
func (sae *StudioApprovalEntry) Many2One() *Many2One {
	return NewMany2One(sae.Id.Get(), "")
}

// CreateStudioApprovalEntry creates a new studio.approval.entry model and returns its id.
func (c *Client) CreateStudioApprovalEntry(sae *StudioApprovalEntry) (int64, error) {
	ids, err := c.CreateStudioApprovalEntrys([]*StudioApprovalEntry{sae})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalEntry creates a new studio.approval.entry model and returns its id.
func (c *Client) CreateStudioApprovalEntrys(saes []*StudioApprovalEntry) ([]int64, error) {
	var vv []interface{}
	for _, v := range saes {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalEntryModel, vv, nil)
}

// UpdateStudioApprovalEntry updates an existing studio.approval.entry record.
func (c *Client) UpdateStudioApprovalEntry(sae *StudioApprovalEntry) error {
	return c.UpdateStudioApprovalEntrys([]int64{sae.Id.Get()}, sae)
}

// UpdateStudioApprovalEntrys updates existing studio.approval.entry records.
// All records (represented by ids) will be updated by sae values.
func (c *Client) UpdateStudioApprovalEntrys(ids []int64, sae *StudioApprovalEntry) error {
	return c.Update(StudioApprovalEntryModel, ids, sae, nil)
}

// DeleteStudioApprovalEntry deletes an existing studio.approval.entry record.
func (c *Client) DeleteStudioApprovalEntry(id int64) error {
	return c.DeleteStudioApprovalEntrys([]int64{id})
}

// DeleteStudioApprovalEntrys deletes existing studio.approval.entry records.
func (c *Client) DeleteStudioApprovalEntrys(ids []int64) error {
	return c.Delete(StudioApprovalEntryModel, ids)
}

// GetStudioApprovalEntry gets studio.approval.entry existing record.
func (c *Client) GetStudioApprovalEntry(id int64) (*StudioApprovalEntry, error) {
	saes, err := c.GetStudioApprovalEntrys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*saes)[0]), nil
}

// GetStudioApprovalEntrys gets studio.approval.entry existing records.
func (c *Client) GetStudioApprovalEntrys(ids []int64) (*StudioApprovalEntrys, error) {
	saes := &StudioApprovalEntrys{}
	if err := c.Read(StudioApprovalEntryModel, ids, nil, saes); err != nil {
		return nil, err
	}
	return saes, nil
}

// FindStudioApprovalEntry finds studio.approval.entry record by querying it with criteria.
func (c *Client) FindStudioApprovalEntry(criteria *Criteria) (*StudioApprovalEntry, error) {
	saes := &StudioApprovalEntrys{}
	if err := c.SearchRead(StudioApprovalEntryModel, criteria, NewOptions().Limit(1), saes); err != nil {
		return nil, err
	}
	return &((*saes)[0]), nil
}

// FindStudioApprovalEntrys finds studio.approval.entry records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalEntrys(criteria *Criteria, options *Options) (*StudioApprovalEntrys, error) {
	saes := &StudioApprovalEntrys{}
	if err := c.SearchRead(StudioApprovalEntryModel, criteria, options, saes); err != nil {
		return nil, err
	}
	return saes, nil
}

// FindStudioApprovalEntryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalEntryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioApprovalEntryModel, criteria, options)
}

// FindStudioApprovalEntryId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalEntryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalEntryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
