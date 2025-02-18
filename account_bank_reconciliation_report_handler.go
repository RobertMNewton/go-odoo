package odoo

// AccountBankReconciliationReportHandler represents account.bank.reconciliation.report.handler model.
type AccountBankReconciliationReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountBankReconciliationReportHandlers represents array of account.bank.reconciliation.report.handler model.
type AccountBankReconciliationReportHandlers []AccountBankReconciliationReportHandler

// AccountBankReconciliationReportHandlerModel is the odoo model name.
const AccountBankReconciliationReportHandlerModel = "account.bank.reconciliation.report.handler"

// Many2One convert AccountBankReconciliationReportHandler to *Many2One.
func (abrrh *AccountBankReconciliationReportHandler) Many2One() *Many2One {
	return NewMany2One(abrrh.Id.Get(), "")
}

// CreateAccountBankReconciliationReportHandler creates a new account.bank.reconciliation.report.handler model and returns its id.
func (c *Client) CreateAccountBankReconciliationReportHandler(abrrh *AccountBankReconciliationReportHandler) (int64, error) {
	ids, err := c.CreateAccountBankReconciliationReportHandlers([]*AccountBankReconciliationReportHandler{abrrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankReconciliationReportHandler creates a new account.bank.reconciliation.report.handler model and returns its id.
func (c *Client) CreateAccountBankReconciliationReportHandlers(abrrhs []*AccountBankReconciliationReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range abrrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountBankReconciliationReportHandlerModel, vv, nil)
}

// UpdateAccountBankReconciliationReportHandler updates an existing account.bank.reconciliation.report.handler record.
func (c *Client) UpdateAccountBankReconciliationReportHandler(abrrh *AccountBankReconciliationReportHandler) error {
	return c.UpdateAccountBankReconciliationReportHandlers([]int64{abrrh.Id.Get()}, abrrh)
}

// UpdateAccountBankReconciliationReportHandlers updates existing account.bank.reconciliation.report.handler records.
// All records (represented by ids) will be updated by abrrh values.
func (c *Client) UpdateAccountBankReconciliationReportHandlers(ids []int64, abrrh *AccountBankReconciliationReportHandler) error {
	return c.Update(AccountBankReconciliationReportHandlerModel, ids, abrrh, nil)
}

// DeleteAccountBankReconciliationReportHandler deletes an existing account.bank.reconciliation.report.handler record.
func (c *Client) DeleteAccountBankReconciliationReportHandler(id int64) error {
	return c.DeleteAccountBankReconciliationReportHandlers([]int64{id})
}

// DeleteAccountBankReconciliationReportHandlers deletes existing account.bank.reconciliation.report.handler records.
func (c *Client) DeleteAccountBankReconciliationReportHandlers(ids []int64) error {
	return c.Delete(AccountBankReconciliationReportHandlerModel, ids)
}

// GetAccountBankReconciliationReportHandler gets account.bank.reconciliation.report.handler existing record.
func (c *Client) GetAccountBankReconciliationReportHandler(id int64) (*AccountBankReconciliationReportHandler, error) {
	abrrhs, err := c.GetAccountBankReconciliationReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abrrhs)[0]), nil
}

// GetAccountBankReconciliationReportHandlers gets account.bank.reconciliation.report.handler existing records.
func (c *Client) GetAccountBankReconciliationReportHandlers(ids []int64) (*AccountBankReconciliationReportHandlers, error) {
	abrrhs := &AccountBankReconciliationReportHandlers{}
	if err := c.Read(AccountBankReconciliationReportHandlerModel, ids, nil, abrrhs); err != nil {
		return nil, err
	}
	return abrrhs, nil
}

// FindAccountBankReconciliationReportHandler finds account.bank.reconciliation.report.handler record by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReportHandler(criteria *Criteria) (*AccountBankReconciliationReportHandler, error) {
	abrrhs := &AccountBankReconciliationReportHandlers{}
	if err := c.SearchRead(AccountBankReconciliationReportHandlerModel, criteria, NewOptions().Limit(1), abrrhs); err != nil {
		return nil, err
	}
	return &((*abrrhs)[0]), nil
}

// FindAccountBankReconciliationReportHandlers finds account.bank.reconciliation.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReportHandlers(criteria *Criteria, options *Options) (*AccountBankReconciliationReportHandlers, error) {
	abrrhs := &AccountBankReconciliationReportHandlers{}
	if err := c.SearchRead(AccountBankReconciliationReportHandlerModel, criteria, options, abrrhs); err != nil {
		return nil, err
	}
	return abrrhs, nil
}

// FindAccountBankReconciliationReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankReconciliationReportHandlerModel, criteria, options)
}

// FindAccountBankReconciliationReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankReconciliationReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
