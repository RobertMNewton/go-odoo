package odoo

// AccountFollowupReport represents account.followup.report model.
type AccountFollowupReport struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountFollowupReports represents array of account.followup.report model.
type AccountFollowupReports []AccountFollowupReport

// AccountFollowupReportModel is the odoo model name.
const AccountFollowupReportModel = "account.followup.report"

// Many2One convert AccountFollowupReport to *Many2One.
func (afr *AccountFollowupReport) Many2One() *Many2One {
	return NewMany2One(afr.Id.Get(), "")
}

// CreateAccountFollowupReport creates a new account.followup.report model and returns its id.
func (c *Client) CreateAccountFollowupReport(afr *AccountFollowupReport) (int64, error) {
	ids, err := c.CreateAccountFollowupReports([]*AccountFollowupReport{afr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupReport creates a new account.followup.report model and returns its id.
func (c *Client) CreateAccountFollowupReports(afrs []*AccountFollowupReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range afrs {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupReportModel, vv, nil)
}

// UpdateAccountFollowupReport updates an existing account.followup.report record.
func (c *Client) UpdateAccountFollowupReport(afr *AccountFollowupReport) error {
	return c.UpdateAccountFollowupReports([]int64{afr.Id.Get()}, afr)
}

// UpdateAccountFollowupReports updates existing account.followup.report records.
// All records (represented by ids) will be updated by afr values.
func (c *Client) UpdateAccountFollowupReports(ids []int64, afr *AccountFollowupReport) error {
	return c.Update(AccountFollowupReportModel, ids, afr, nil)
}

// DeleteAccountFollowupReport deletes an existing account.followup.report record.
func (c *Client) DeleteAccountFollowupReport(id int64) error {
	return c.DeleteAccountFollowupReports([]int64{id})
}

// DeleteAccountFollowupReports deletes existing account.followup.report records.
func (c *Client) DeleteAccountFollowupReports(ids []int64) error {
	return c.Delete(AccountFollowupReportModel, ids)
}

// GetAccountFollowupReport gets account.followup.report existing record.
func (c *Client) GetAccountFollowupReport(id int64) (*AccountFollowupReport, error) {
	afrs, err := c.GetAccountFollowupReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afrs)[0]), nil
}

// GetAccountFollowupReports gets account.followup.report existing records.
func (c *Client) GetAccountFollowupReports(ids []int64) (*AccountFollowupReports, error) {
	afrs := &AccountFollowupReports{}
	if err := c.Read(AccountFollowupReportModel, ids, nil, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFollowupReport finds account.followup.report record by querying it with criteria.
func (c *Client) FindAccountFollowupReport(criteria *Criteria) (*AccountFollowupReport, error) {
	afrs := &AccountFollowupReports{}
	if err := c.SearchRead(AccountFollowupReportModel, criteria, NewOptions().Limit(1), afrs); err != nil {
		return nil, err
	}
	return &((*afrs)[0]), nil
}

// FindAccountFollowupReports finds account.followup.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReports(criteria *Criteria, options *Options) (*AccountFollowupReports, error) {
	afrs := &AccountFollowupReports{}
	if err := c.SearchRead(AccountFollowupReportModel, criteria, options, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFollowupReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFollowupReportModel, criteria, options)
}

// FindAccountFollowupReportId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
