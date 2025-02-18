package odoo

// AccountDisallowedExpensesReportHandler represents account.disallowed.expenses.report.handler model.
type AccountDisallowedExpensesReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountDisallowedExpensesReportHandlers represents array of account.disallowed.expenses.report.handler model.
type AccountDisallowedExpensesReportHandlers []AccountDisallowedExpensesReportHandler

// AccountDisallowedExpensesReportHandlerModel is the odoo model name.
const AccountDisallowedExpensesReportHandlerModel = "account.disallowed.expenses.report.handler"

// Many2One convert AccountDisallowedExpensesReportHandler to *Many2One.
func (aderh *AccountDisallowedExpensesReportHandler) Many2One() *Many2One {
	return NewMany2One(aderh.Id.Get(), "")
}

// CreateAccountDisallowedExpensesReportHandler creates a new account.disallowed.expenses.report.handler model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesReportHandler(aderh *AccountDisallowedExpensesReportHandler) (int64, error) {
	ids, err := c.CreateAccountDisallowedExpensesReportHandlers([]*AccountDisallowedExpensesReportHandler{aderh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDisallowedExpensesReportHandler creates a new account.disallowed.expenses.report.handler model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesReportHandlers(aderhs []*AccountDisallowedExpensesReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aderhs {
		vv = append(vv, v)
	}
	return c.Create(AccountDisallowedExpensesReportHandlerModel, vv, nil)
}

// UpdateAccountDisallowedExpensesReportHandler updates an existing account.disallowed.expenses.report.handler record.
func (c *Client) UpdateAccountDisallowedExpensesReportHandler(aderh *AccountDisallowedExpensesReportHandler) error {
	return c.UpdateAccountDisallowedExpensesReportHandlers([]int64{aderh.Id.Get()}, aderh)
}

// UpdateAccountDisallowedExpensesReportHandlers updates existing account.disallowed.expenses.report.handler records.
// All records (represented by ids) will be updated by aderh values.
func (c *Client) UpdateAccountDisallowedExpensesReportHandlers(ids []int64, aderh *AccountDisallowedExpensesReportHandler) error {
	return c.Update(AccountDisallowedExpensesReportHandlerModel, ids, aderh, nil)
}

// DeleteAccountDisallowedExpensesReportHandler deletes an existing account.disallowed.expenses.report.handler record.
func (c *Client) DeleteAccountDisallowedExpensesReportHandler(id int64) error {
	return c.DeleteAccountDisallowedExpensesReportHandlers([]int64{id})
}

// DeleteAccountDisallowedExpensesReportHandlers deletes existing account.disallowed.expenses.report.handler records.
func (c *Client) DeleteAccountDisallowedExpensesReportHandlers(ids []int64) error {
	return c.Delete(AccountDisallowedExpensesReportHandlerModel, ids)
}

// GetAccountDisallowedExpensesReportHandler gets account.disallowed.expenses.report.handler existing record.
func (c *Client) GetAccountDisallowedExpensesReportHandler(id int64) (*AccountDisallowedExpensesReportHandler, error) {
	aderhs, err := c.GetAccountDisallowedExpensesReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aderhs)[0]), nil
}

// GetAccountDisallowedExpensesReportHandlers gets account.disallowed.expenses.report.handler existing records.
func (c *Client) GetAccountDisallowedExpensesReportHandlers(ids []int64) (*AccountDisallowedExpensesReportHandlers, error) {
	aderhs := &AccountDisallowedExpensesReportHandlers{}
	if err := c.Read(AccountDisallowedExpensesReportHandlerModel, ids, nil, aderhs); err != nil {
		return nil, err
	}
	return aderhs, nil
}

// FindAccountDisallowedExpensesReportHandler finds account.disallowed.expenses.report.handler record by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesReportHandler(criteria *Criteria) (*AccountDisallowedExpensesReportHandler, error) {
	aderhs := &AccountDisallowedExpensesReportHandlers{}
	if err := c.SearchRead(AccountDisallowedExpensesReportHandlerModel, criteria, NewOptions().Limit(1), aderhs); err != nil {
		return nil, err
	}
	return &((*aderhs)[0]), nil
}

// FindAccountDisallowedExpensesReportHandlers finds account.disallowed.expenses.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesReportHandlers(criteria *Criteria, options *Options) (*AccountDisallowedExpensesReportHandlers, error) {
	aderhs := &AccountDisallowedExpensesReportHandlers{}
	if err := c.SearchRead(AccountDisallowedExpensesReportHandlerModel, criteria, options, aderhs); err != nil {
		return nil, err
	}
	return aderhs, nil
}

// FindAccountDisallowedExpensesReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDisallowedExpensesReportHandlerModel, criteria, options)
}

// FindAccountDisallowedExpensesReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDisallowedExpensesReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
