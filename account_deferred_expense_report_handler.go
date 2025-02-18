package odoo

// AccountDeferredExpenseReportHandler represents account.deferred.expense.report.handler model.
type AccountDeferredExpenseReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountDeferredExpenseReportHandlers represents array of account.deferred.expense.report.handler model.
type AccountDeferredExpenseReportHandlers []AccountDeferredExpenseReportHandler

// AccountDeferredExpenseReportHandlerModel is the odoo model name.
const AccountDeferredExpenseReportHandlerModel = "account.deferred.expense.report.handler"

// Many2One convert AccountDeferredExpenseReportHandler to *Many2One.
func (aderh *AccountDeferredExpenseReportHandler) Many2One() *Many2One {
	return NewMany2One(aderh.Id.Get(), "")
}

// CreateAccountDeferredExpenseReportHandler creates a new account.deferred.expense.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredExpenseReportHandler(aderh *AccountDeferredExpenseReportHandler) (int64, error) {
	ids, err := c.CreateAccountDeferredExpenseReportHandlers([]*AccountDeferredExpenseReportHandler{aderh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDeferredExpenseReportHandler creates a new account.deferred.expense.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredExpenseReportHandlers(aderhs []*AccountDeferredExpenseReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aderhs {
		vv = append(vv, v)
	}
	return c.Create(AccountDeferredExpenseReportHandlerModel, vv, nil)
}

// UpdateAccountDeferredExpenseReportHandler updates an existing account.deferred.expense.report.handler record.
func (c *Client) UpdateAccountDeferredExpenseReportHandler(aderh *AccountDeferredExpenseReportHandler) error {
	return c.UpdateAccountDeferredExpenseReportHandlers([]int64{aderh.Id.Get()}, aderh)
}

// UpdateAccountDeferredExpenseReportHandlers updates existing account.deferred.expense.report.handler records.
// All records (represented by ids) will be updated by aderh values.
func (c *Client) UpdateAccountDeferredExpenseReportHandlers(ids []int64, aderh *AccountDeferredExpenseReportHandler) error {
	return c.Update(AccountDeferredExpenseReportHandlerModel, ids, aderh, nil)
}

// DeleteAccountDeferredExpenseReportHandler deletes an existing account.deferred.expense.report.handler record.
func (c *Client) DeleteAccountDeferredExpenseReportHandler(id int64) error {
	return c.DeleteAccountDeferredExpenseReportHandlers([]int64{id})
}

// DeleteAccountDeferredExpenseReportHandlers deletes existing account.deferred.expense.report.handler records.
func (c *Client) DeleteAccountDeferredExpenseReportHandlers(ids []int64) error {
	return c.Delete(AccountDeferredExpenseReportHandlerModel, ids)
}

// GetAccountDeferredExpenseReportHandler gets account.deferred.expense.report.handler existing record.
func (c *Client) GetAccountDeferredExpenseReportHandler(id int64) (*AccountDeferredExpenseReportHandler, error) {
	aderhs, err := c.GetAccountDeferredExpenseReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aderhs)[0]), nil
}

// GetAccountDeferredExpenseReportHandlers gets account.deferred.expense.report.handler existing records.
func (c *Client) GetAccountDeferredExpenseReportHandlers(ids []int64) (*AccountDeferredExpenseReportHandlers, error) {
	aderhs := &AccountDeferredExpenseReportHandlers{}
	if err := c.Read(AccountDeferredExpenseReportHandlerModel, ids, nil, aderhs); err != nil {
		return nil, err
	}
	return aderhs, nil
}

// FindAccountDeferredExpenseReportHandler finds account.deferred.expense.report.handler record by querying it with criteria.
func (c *Client) FindAccountDeferredExpenseReportHandler(criteria *Criteria) (*AccountDeferredExpenseReportHandler, error) {
	aderhs := &AccountDeferredExpenseReportHandlers{}
	if err := c.SearchRead(AccountDeferredExpenseReportHandlerModel, criteria, NewOptions().Limit(1), aderhs); err != nil {
		return nil, err
	}
	return &((*aderhs)[0]), nil
}

// FindAccountDeferredExpenseReportHandlers finds account.deferred.expense.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredExpenseReportHandlers(criteria *Criteria, options *Options) (*AccountDeferredExpenseReportHandlers, error) {
	aderhs := &AccountDeferredExpenseReportHandlers{}
	if err := c.SearchRead(AccountDeferredExpenseReportHandlerModel, criteria, options, aderhs); err != nil {
		return nil, err
	}
	return aderhs, nil
}

// FindAccountDeferredExpenseReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredExpenseReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDeferredExpenseReportHandlerModel, criteria, options)
}

// FindAccountDeferredExpenseReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountDeferredExpenseReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDeferredExpenseReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
