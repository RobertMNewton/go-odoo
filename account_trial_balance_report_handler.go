package odoo

// AccountTrialBalanceReportHandler represents account.trial.balance.report.handler model.
type AccountTrialBalanceReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountTrialBalanceReportHandlers represents array of account.trial.balance.report.handler model.
type AccountTrialBalanceReportHandlers []AccountTrialBalanceReportHandler

// AccountTrialBalanceReportHandlerModel is the odoo model name.
const AccountTrialBalanceReportHandlerModel = "account.trial.balance.report.handler"

// Many2One convert AccountTrialBalanceReportHandler to *Many2One.
func (atbrh *AccountTrialBalanceReportHandler) Many2One() *Many2One {
	return NewMany2One(atbrh.Id.Get(), "")
}

// CreateAccountTrialBalanceReportHandler creates a new account.trial.balance.report.handler model and returns its id.
func (c *Client) CreateAccountTrialBalanceReportHandler(atbrh *AccountTrialBalanceReportHandler) (int64, error) {
	ids, err := c.CreateAccountTrialBalanceReportHandlers([]*AccountTrialBalanceReportHandler{atbrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTrialBalanceReportHandler creates a new account.trial.balance.report.handler model and returns its id.
func (c *Client) CreateAccountTrialBalanceReportHandlers(atbrhs []*AccountTrialBalanceReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range atbrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountTrialBalanceReportHandlerModel, vv, nil)
}

// UpdateAccountTrialBalanceReportHandler updates an existing account.trial.balance.report.handler record.
func (c *Client) UpdateAccountTrialBalanceReportHandler(atbrh *AccountTrialBalanceReportHandler) error {
	return c.UpdateAccountTrialBalanceReportHandlers([]int64{atbrh.Id.Get()}, atbrh)
}

// UpdateAccountTrialBalanceReportHandlers updates existing account.trial.balance.report.handler records.
// All records (represented by ids) will be updated by atbrh values.
func (c *Client) UpdateAccountTrialBalanceReportHandlers(ids []int64, atbrh *AccountTrialBalanceReportHandler) error {
	return c.Update(AccountTrialBalanceReportHandlerModel, ids, atbrh, nil)
}

// DeleteAccountTrialBalanceReportHandler deletes an existing account.trial.balance.report.handler record.
func (c *Client) DeleteAccountTrialBalanceReportHandler(id int64) error {
	return c.DeleteAccountTrialBalanceReportHandlers([]int64{id})
}

// DeleteAccountTrialBalanceReportHandlers deletes existing account.trial.balance.report.handler records.
func (c *Client) DeleteAccountTrialBalanceReportHandlers(ids []int64) error {
	return c.Delete(AccountTrialBalanceReportHandlerModel, ids)
}

// GetAccountTrialBalanceReportHandler gets account.trial.balance.report.handler existing record.
func (c *Client) GetAccountTrialBalanceReportHandler(id int64) (*AccountTrialBalanceReportHandler, error) {
	atbrhs, err := c.GetAccountTrialBalanceReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atbrhs)[0]), nil
}

// GetAccountTrialBalanceReportHandlers gets account.trial.balance.report.handler existing records.
func (c *Client) GetAccountTrialBalanceReportHandlers(ids []int64) (*AccountTrialBalanceReportHandlers, error) {
	atbrhs := &AccountTrialBalanceReportHandlers{}
	if err := c.Read(AccountTrialBalanceReportHandlerModel, ids, nil, atbrhs); err != nil {
		return nil, err
	}
	return atbrhs, nil
}

// FindAccountTrialBalanceReportHandler finds account.trial.balance.report.handler record by querying it with criteria.
func (c *Client) FindAccountTrialBalanceReportHandler(criteria *Criteria) (*AccountTrialBalanceReportHandler, error) {
	atbrhs := &AccountTrialBalanceReportHandlers{}
	if err := c.SearchRead(AccountTrialBalanceReportHandlerModel, criteria, NewOptions().Limit(1), atbrhs); err != nil {
		return nil, err
	}
	return &((*atbrhs)[0]), nil
}

// FindAccountTrialBalanceReportHandlers finds account.trial.balance.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTrialBalanceReportHandlers(criteria *Criteria, options *Options) (*AccountTrialBalanceReportHandlers, error) {
	atbrhs := &AccountTrialBalanceReportHandlers{}
	if err := c.SearchRead(AccountTrialBalanceReportHandlerModel, criteria, options, atbrhs); err != nil {
		return nil, err
	}
	return atbrhs, nil
}

// FindAccountTrialBalanceReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTrialBalanceReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTrialBalanceReportHandlerModel, criteria, options)
}

// FindAccountTrialBalanceReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountTrialBalanceReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTrialBalanceReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
