package odoo

// StudioApprovalRule represents studio.approval.rule model.
type StudioApprovalRule struct {
	ActionId                 *Many2One  `xmlrpc:"action_id,omitempty"`
	ActionXmlid              *String    `xmlrpc:"action_xmlid,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	ApprovalGroupId          *Many2One  `xmlrpc:"approval_group_id,omitempty"`
	ApproverIds              *Relation  `xmlrpc:"approver_ids,omitempty"`
	ApproverLogIds           *Relation  `xmlrpc:"approver_log_ids,omitempty"`
	CanValidate              *Bool      `xmlrpc:"can_validate,omitempty"`
	Conditional              *Bool      `xmlrpc:"conditional,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Domain                   *String    `xmlrpc:"domain,omitempty"`
	EntriesCount             *Int       `xmlrpc:"entries_count,omitempty"`
	EntryIds                 *Relation  `xmlrpc:"entry_ids,omitempty"`
	ExclusiveUser            *Bool      `xmlrpc:"exclusive_user,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	KanbanColor              *Int       `xmlrpc:"kanban_color,omitempty"`
	Message                  *String    `xmlrpc:"message,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Method                   *String    `xmlrpc:"method,omitempty"`
	ModelId                  *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelName                *String    `xmlrpc:"model_name,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	NotificationOrder        *Selection `xmlrpc:"notification_order,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	UsersToNotify            *Relation  `xmlrpc:"users_to_notify,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StudioApprovalRules represents array of studio.approval.rule model.
type StudioApprovalRules []StudioApprovalRule

// StudioApprovalRuleModel is the odoo model name.
const StudioApprovalRuleModel = "studio.approval.rule"

// Many2One convert StudioApprovalRule to *Many2One.
func (sar *StudioApprovalRule) Many2One() *Many2One {
	return NewMany2One(sar.Id.Get(), "")
}

// CreateStudioApprovalRule creates a new studio.approval.rule model and returns its id.
func (c *Client) CreateStudioApprovalRule(sar *StudioApprovalRule) (int64, error) {
	ids, err := c.CreateStudioApprovalRules([]*StudioApprovalRule{sar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRule creates a new studio.approval.rule model and returns its id.
func (c *Client) CreateStudioApprovalRules(sars []*StudioApprovalRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range sars {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRuleModel, vv, nil)
}

// UpdateStudioApprovalRule updates an existing studio.approval.rule record.
func (c *Client) UpdateStudioApprovalRule(sar *StudioApprovalRule) error {
	return c.UpdateStudioApprovalRules([]int64{sar.Id.Get()}, sar)
}

// UpdateStudioApprovalRules updates existing studio.approval.rule records.
// All records (represented by ids) will be updated by sar values.
func (c *Client) UpdateStudioApprovalRules(ids []int64, sar *StudioApprovalRule) error {
	return c.Update(StudioApprovalRuleModel, ids, sar, nil)
}

// DeleteStudioApprovalRule deletes an existing studio.approval.rule record.
func (c *Client) DeleteStudioApprovalRule(id int64) error {
	return c.DeleteStudioApprovalRules([]int64{id})
}

// DeleteStudioApprovalRules deletes existing studio.approval.rule records.
func (c *Client) DeleteStudioApprovalRules(ids []int64) error {
	return c.Delete(StudioApprovalRuleModel, ids)
}

// GetStudioApprovalRule gets studio.approval.rule existing record.
func (c *Client) GetStudioApprovalRule(id int64) (*StudioApprovalRule, error) {
	sars, err := c.GetStudioApprovalRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// GetStudioApprovalRules gets studio.approval.rule existing records.
func (c *Client) GetStudioApprovalRules(ids []int64) (*StudioApprovalRules, error) {
	sars := &StudioApprovalRules{}
	if err := c.Read(StudioApprovalRuleModel, ids, nil, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRule finds studio.approval.rule record by querying it with criteria.
func (c *Client) FindStudioApprovalRule(criteria *Criteria) (*StudioApprovalRule, error) {
	sars := &StudioApprovalRules{}
	if err := c.SearchRead(StudioApprovalRuleModel, criteria, NewOptions().Limit(1), sars); err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// FindStudioApprovalRules finds studio.approval.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRules(criteria *Criteria, options *Options) (*StudioApprovalRules, error) {
	sars := &StudioApprovalRules{}
	if err := c.SearchRead(StudioApprovalRuleModel, criteria, options, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioApprovalRuleModel, criteria, options)
}

// FindStudioApprovalRuleId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
