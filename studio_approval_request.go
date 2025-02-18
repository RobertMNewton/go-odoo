package odoo

// StudioApprovalRequest represents studio.approval.request model.
type StudioApprovalRequest struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	MailActivityId *Many2One `xmlrpc:"mail_activity_id,omitempty"`
	ResId          *Many2One `xmlrpc:"res_id,omitempty"`
	RuleId         *Many2One `xmlrpc:"rule_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioApprovalRequests represents array of studio.approval.request model.
type StudioApprovalRequests []StudioApprovalRequest

// StudioApprovalRequestModel is the odoo model name.
const StudioApprovalRequestModel = "studio.approval.request"

// Many2One convert StudioApprovalRequest to *Many2One.
func (sar *StudioApprovalRequest) Many2One() *Many2One {
	return NewMany2One(sar.Id.Get(), "")
}

// CreateStudioApprovalRequest creates a new studio.approval.request model and returns its id.
func (c *Client) CreateStudioApprovalRequest(sar *StudioApprovalRequest) (int64, error) {
	ids, err := c.CreateStudioApprovalRequests([]*StudioApprovalRequest{sar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRequest creates a new studio.approval.request model and returns its id.
func (c *Client) CreateStudioApprovalRequests(sars []*StudioApprovalRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range sars {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRequestModel, vv, nil)
}

// UpdateStudioApprovalRequest updates an existing studio.approval.request record.
func (c *Client) UpdateStudioApprovalRequest(sar *StudioApprovalRequest) error {
	return c.UpdateStudioApprovalRequests([]int64{sar.Id.Get()}, sar)
}

// UpdateStudioApprovalRequests updates existing studio.approval.request records.
// All records (represented by ids) will be updated by sar values.
func (c *Client) UpdateStudioApprovalRequests(ids []int64, sar *StudioApprovalRequest) error {
	return c.Update(StudioApprovalRequestModel, ids, sar, nil)
}

// DeleteStudioApprovalRequest deletes an existing studio.approval.request record.
func (c *Client) DeleteStudioApprovalRequest(id int64) error {
	return c.DeleteStudioApprovalRequests([]int64{id})
}

// DeleteStudioApprovalRequests deletes existing studio.approval.request records.
func (c *Client) DeleteStudioApprovalRequests(ids []int64) error {
	return c.Delete(StudioApprovalRequestModel, ids)
}

// GetStudioApprovalRequest gets studio.approval.request existing record.
func (c *Client) GetStudioApprovalRequest(id int64) (*StudioApprovalRequest, error) {
	sars, err := c.GetStudioApprovalRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// GetStudioApprovalRequests gets studio.approval.request existing records.
func (c *Client) GetStudioApprovalRequests(ids []int64) (*StudioApprovalRequests, error) {
	sars := &StudioApprovalRequests{}
	if err := c.Read(StudioApprovalRequestModel, ids, nil, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRequest finds studio.approval.request record by querying it with criteria.
func (c *Client) FindStudioApprovalRequest(criteria *Criteria) (*StudioApprovalRequest, error) {
	sars := &StudioApprovalRequests{}
	if err := c.SearchRead(StudioApprovalRequestModel, criteria, NewOptions().Limit(1), sars); err != nil {
		return nil, err
	}
	return &((*sars)[0]), nil
}

// FindStudioApprovalRequests finds studio.approval.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRequests(criteria *Criteria, options *Options) (*StudioApprovalRequests, error) {
	sars := &StudioApprovalRequests{}
	if err := c.SearchRead(StudioApprovalRequestModel, criteria, options, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioApprovalRequestModel, criteria, options)
}

// FindStudioApprovalRequestId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
