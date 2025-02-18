package odoo

// StudioApprovalRuleDelegate represents studio.approval.rule.delegate model.
type StudioApprovalRuleDelegate struct {
	ApprovalRuleId *Many2One `xmlrpc:"approval_rule_id,omitempty"`
	ApproverIds    *Relation `xmlrpc:"approver_ids,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DateTo         *Time     `xmlrpc:"date_to,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	UsersToNotify  *Relation `xmlrpc:"users_to_notify,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioApprovalRuleDelegates represents array of studio.approval.rule.delegate model.
type StudioApprovalRuleDelegates []StudioApprovalRuleDelegate

// StudioApprovalRuleDelegateModel is the odoo model name.
const StudioApprovalRuleDelegateModel = "studio.approval.rule.delegate"

// Many2One convert StudioApprovalRuleDelegate to *Many2One.
func (sard *StudioApprovalRuleDelegate) Many2One() *Many2One {
	return NewMany2One(sard.Id.Get(), "")
}

// CreateStudioApprovalRuleDelegate creates a new studio.approval.rule.delegate model and returns its id.
func (c *Client) CreateStudioApprovalRuleDelegate(sard *StudioApprovalRuleDelegate) (int64, error) {
	ids, err := c.CreateStudioApprovalRuleDelegates([]*StudioApprovalRuleDelegate{sard})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRuleDelegate creates a new studio.approval.rule.delegate model and returns its id.
func (c *Client) CreateStudioApprovalRuleDelegates(sards []*StudioApprovalRuleDelegate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sards {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRuleDelegateModel, vv, nil)
}

// UpdateStudioApprovalRuleDelegate updates an existing studio.approval.rule.delegate record.
func (c *Client) UpdateStudioApprovalRuleDelegate(sard *StudioApprovalRuleDelegate) error {
	return c.UpdateStudioApprovalRuleDelegates([]int64{sard.Id.Get()}, sard)
}

// UpdateStudioApprovalRuleDelegates updates existing studio.approval.rule.delegate records.
// All records (represented by ids) will be updated by sard values.
func (c *Client) UpdateStudioApprovalRuleDelegates(ids []int64, sard *StudioApprovalRuleDelegate) error {
	return c.Update(StudioApprovalRuleDelegateModel, ids, sard, nil)
}

// DeleteStudioApprovalRuleDelegate deletes an existing studio.approval.rule.delegate record.
func (c *Client) DeleteStudioApprovalRuleDelegate(id int64) error {
	return c.DeleteStudioApprovalRuleDelegates([]int64{id})
}

// DeleteStudioApprovalRuleDelegates deletes existing studio.approval.rule.delegate records.
func (c *Client) DeleteStudioApprovalRuleDelegates(ids []int64) error {
	return c.Delete(StudioApprovalRuleDelegateModel, ids)
}

// GetStudioApprovalRuleDelegate gets studio.approval.rule.delegate existing record.
func (c *Client) GetStudioApprovalRuleDelegate(id int64) (*StudioApprovalRuleDelegate, error) {
	sards, err := c.GetStudioApprovalRuleDelegates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sards)[0]), nil
}

// GetStudioApprovalRuleDelegates gets studio.approval.rule.delegate existing records.
func (c *Client) GetStudioApprovalRuleDelegates(ids []int64) (*StudioApprovalRuleDelegates, error) {
	sards := &StudioApprovalRuleDelegates{}
	if err := c.Read(StudioApprovalRuleDelegateModel, ids, nil, sards); err != nil {
		return nil, err
	}
	return sards, nil
}

// FindStudioApprovalRuleDelegate finds studio.approval.rule.delegate record by querying it with criteria.
func (c *Client) FindStudioApprovalRuleDelegate(criteria *Criteria) (*StudioApprovalRuleDelegate, error) {
	sards := &StudioApprovalRuleDelegates{}
	if err := c.SearchRead(StudioApprovalRuleDelegateModel, criteria, NewOptions().Limit(1), sards); err != nil {
		return nil, err
	}
	return &((*sards)[0]), nil
}

// FindStudioApprovalRuleDelegates finds studio.approval.rule.delegate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleDelegates(criteria *Criteria, options *Options) (*StudioApprovalRuleDelegates, error) {
	sards := &StudioApprovalRuleDelegates{}
	if err := c.SearchRead(StudioApprovalRuleDelegateModel, criteria, options, sards); err != nil {
		return nil, err
	}
	return sards, nil
}

// FindStudioApprovalRuleDelegateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleDelegateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioApprovalRuleDelegateModel, criteria, options)
}

// FindStudioApprovalRuleDelegateId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRuleDelegateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRuleDelegateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
