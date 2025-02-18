package odoo

// AccountCashFlowReportHandler represents account.cash.flow.report.handler model.
type AccountCashFlowReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountCashFlowReportHandlers represents array of account.cash.flow.report.handler model.
type AccountCashFlowReportHandlers []AccountCashFlowReportHandler

// AccountCashFlowReportHandlerModel is the odoo model name.
const AccountCashFlowReportHandlerModel = "account.cash.flow.report.handler"

// Many2One convert AccountCashFlowReportHandler to *Many2One.
func (acfrh *AccountCashFlowReportHandler) Many2One() *Many2One {
	return NewMany2One(acfrh.Id.Get(), "")
}

// CreateAccountCashFlowReportHandler creates a new account.cash.flow.report.handler model and returns its id.
func (c *Client) CreateAccountCashFlowReportHandler(acfrh *AccountCashFlowReportHandler) (int64, error) {
	ids, err := c.CreateAccountCashFlowReportHandlers([]*AccountCashFlowReportHandler{acfrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCashFlowReportHandler creates a new account.cash.flow.report.handler model and returns its id.
func (c *Client) CreateAccountCashFlowReportHandlers(acfrhs []*AccountCashFlowReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range acfrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountCashFlowReportHandlerModel, vv, nil)
}

// UpdateAccountCashFlowReportHandler updates an existing account.cash.flow.report.handler record.
func (c *Client) UpdateAccountCashFlowReportHandler(acfrh *AccountCashFlowReportHandler) error {
	return c.UpdateAccountCashFlowReportHandlers([]int64{acfrh.Id.Get()}, acfrh)
}

// UpdateAccountCashFlowReportHandlers updates existing account.cash.flow.report.handler records.
// All records (represented by ids) will be updated by acfrh values.
func (c *Client) UpdateAccountCashFlowReportHandlers(ids []int64, acfrh *AccountCashFlowReportHandler) error {
	return c.Update(AccountCashFlowReportHandlerModel, ids, acfrh, nil)
}

// DeleteAccountCashFlowReportHandler deletes an existing account.cash.flow.report.handler record.
func (c *Client) DeleteAccountCashFlowReportHandler(id int64) error {
	return c.DeleteAccountCashFlowReportHandlers([]int64{id})
}

// DeleteAccountCashFlowReportHandlers deletes existing account.cash.flow.report.handler records.
func (c *Client) DeleteAccountCashFlowReportHandlers(ids []int64) error {
	return c.Delete(AccountCashFlowReportHandlerModel, ids)
}

// GetAccountCashFlowReportHandler gets account.cash.flow.report.handler existing record.
func (c *Client) GetAccountCashFlowReportHandler(id int64) (*AccountCashFlowReportHandler, error) {
	acfrhs, err := c.GetAccountCashFlowReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acfrhs)[0]), nil
}

// GetAccountCashFlowReportHandlers gets account.cash.flow.report.handler existing records.
func (c *Client) GetAccountCashFlowReportHandlers(ids []int64) (*AccountCashFlowReportHandlers, error) {
	acfrhs := &AccountCashFlowReportHandlers{}
	if err := c.Read(AccountCashFlowReportHandlerModel, ids, nil, acfrhs); err != nil {
		return nil, err
	}
	return acfrhs, nil
}

// FindAccountCashFlowReportHandler finds account.cash.flow.report.handler record by querying it with criteria.
func (c *Client) FindAccountCashFlowReportHandler(criteria *Criteria) (*AccountCashFlowReportHandler, error) {
	acfrhs := &AccountCashFlowReportHandlers{}
	if err := c.SearchRead(AccountCashFlowReportHandlerModel, criteria, NewOptions().Limit(1), acfrhs); err != nil {
		return nil, err
	}
	return &((*acfrhs)[0]), nil
}

// FindAccountCashFlowReportHandlers finds account.cash.flow.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashFlowReportHandlers(criteria *Criteria, options *Options) (*AccountCashFlowReportHandlers, error) {
	acfrhs := &AccountCashFlowReportHandlers{}
	if err := c.SearchRead(AccountCashFlowReportHandlerModel, criteria, options, acfrhs); err != nil {
		return nil, err
	}
	return acfrhs, nil
}

// FindAccountCashFlowReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCashFlowReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCashFlowReportHandlerModel, criteria, options)
}

// FindAccountCashFlowReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountCashFlowReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCashFlowReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
