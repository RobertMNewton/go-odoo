package odoo

// StudioApprovalRuleApprover represents studio.approval.rule.approver model.
type StudioApprovalRuleApprover struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DateTo       *Time     `xmlrpc:"date_to,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	IsDelegation *Bool     `xmlrpc:"is_delegation,omitempty"`
	RuleId       *Many2One `xmlrpc:"rule_id,omitempty"`
	UserId       *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioApprovalRuleApprovers represents array of studio.approval.rule.approver model.
type StudioApprovalRuleApprovers []StudioApprovalRuleApprover

// StudioApprovalRuleApproverModel is the odoo model name.
const StudioApprovalRuleApproverModel = "studio.approval.rule.approver"

// Many2One convert StudioApprovalRuleApprover to *Many2One.
func (sara *StudioApprovalRuleApprover) Many2One() *Many2One {
	return NewMany2One(sara.Id.Get(), "")
}

// CreateStudioApprovalRuleApprover creates a new studio.approval.rule.approver model and returns its id.
func (c *Client) CreateStudioApprovalRuleApprover(sara *StudioApprovalRuleApprover) (int64, error) {
	ids, err := c.CreateStudioApprovalRuleApprovers([]*StudioApprovalRuleApprover{sara})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRuleApprover creates a new studio.approval.rule.approver model and returns its id.
func (c *Client) CreateStudioApprovalRuleApprovers(saras []*StudioApprovalRuleApprover) ([]int64, error) {
	var vv []interface{}
	for _, v := range saras {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRuleApproverModel, vv, nil)
}

// UpdateStudioApprovalRuleApprover updates an existing studio.approval.rule.approver record.
func (c *Client) UpdateStudioApprovalRuleApprover(sara *StudioApprovalRuleApprover) error {
	return c.UpdateStudioApprovalRuleApprovers([]int64{sara.Id.Get()}, sara)
}

// UpdateStudioApprovalRuleApprovers updates existing studio.approval.rule.approver records.
// All records (represented by ids) will be updated by sara values.
func (c *Client) UpdateStudioApprovalRuleApprovers(ids []int64, sara *StudioApprovalRuleApprover) error {
	return c.Update(StudioApprovalRuleApproverModel, ids, sara, nil)
}

// DeleteStudioApprovalRuleApprover deletes an existing studio.approval.rule.approver record.
func (c *Client) DeleteStudioApprovalRuleApprover(id int64) error {
	return c.DeleteStudioApprovalRuleApprovers([]int64{id})
}

// DeleteStudioApprovalRuleApprovers deletes existing studio.approval.rule.approver records.
func (c *Client) DeleteStudioApprovalRuleApprovers(ids []int64) error {
	return c.Delete(StudioApprovalRuleApproverModel, ids)
}

// GetStudioApprovalRuleApprover gets studio.approval.rule.approver existing record.
func (c *Client) GetStudioApprovalRuleApprover(id int64) (*StudioApprovalRuleApprover, error) {
	saras, err := c.GetStudioApprovalRuleApprovers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*saras)[0]), nil
}

// GetStudioApprovalRuleApprovers gets studio.approval.rule.approver existing records.
func (c *Client) GetStudioApprovalRuleApprovers(ids []int64) (*StudioApprovalRuleApprovers, error) {
	saras := &StudioApprovalRuleApprovers{}
	if err := c.Read(StudioApprovalRuleApproverModel, ids, nil, saras); err != nil {
		return nil, err
	}
	return saras, nil
}

// FindStudioApprovalRuleApprover finds studio.approval.rule.approver record by querying it with criteria.
func (c *Client) FindStudioApprovalRuleApprover(criteria *Criteria) (*StudioApprovalRuleApprover, error) {
	saras := &StudioApprovalRuleApprovers{}
	if err := c.SearchRead(StudioApprovalRuleApproverModel, criteria, NewOptions().Limit(1), saras); err != nil {
		return nil, err
	}
	return &((*saras)[0]), nil
}

// FindStudioApprovalRuleApprovers finds studio.approval.rule.approver records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleApprovers(criteria *Criteria, options *Options) (*StudioApprovalRuleApprovers, error) {
	saras := &StudioApprovalRuleApprovers{}
	if err := c.SearchRead(StudioApprovalRuleApproverModel, criteria, options, saras); err != nil {
		return nil, err
	}
	return saras, nil
}

// FindStudioApprovalRuleApproverIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleApproverIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioApprovalRuleApproverModel, criteria, options)
}

// FindStudioApprovalRuleApproverId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRuleApproverId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRuleApproverModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
