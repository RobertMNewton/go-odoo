package odoo

// AccountCustomerStatementReportHandler represents account.customer.statement.report.handler model.
type AccountCustomerStatementReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountCustomerStatementReportHandlers represents array of account.customer.statement.report.handler model.
type AccountCustomerStatementReportHandlers []AccountCustomerStatementReportHandler

// AccountCustomerStatementReportHandlerModel is the odoo model name.
const AccountCustomerStatementReportHandlerModel = "account.customer.statement.report.handler"

// Many2One convert AccountCustomerStatementReportHandler to *Many2One.
func (acsrh *AccountCustomerStatementReportHandler) Many2One() *Many2One {
	return NewMany2One(acsrh.Id.Get(), "")
}

// CreateAccountCustomerStatementReportHandler creates a new account.customer.statement.report.handler model and returns its id.
func (c *Client) CreateAccountCustomerStatementReportHandler(acsrh *AccountCustomerStatementReportHandler) (int64, error) {
	ids, err := c.CreateAccountCustomerStatementReportHandlers([]*AccountCustomerStatementReportHandler{acsrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountCustomerStatementReportHandler creates a new account.customer.statement.report.handler model and returns its id.
func (c *Client) CreateAccountCustomerStatementReportHandlers(acsrhs []*AccountCustomerStatementReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range acsrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountCustomerStatementReportHandlerModel, vv, nil)
}

// UpdateAccountCustomerStatementReportHandler updates an existing account.customer.statement.report.handler record.
func (c *Client) UpdateAccountCustomerStatementReportHandler(acsrh *AccountCustomerStatementReportHandler) error {
	return c.UpdateAccountCustomerStatementReportHandlers([]int64{acsrh.Id.Get()}, acsrh)
}

// UpdateAccountCustomerStatementReportHandlers updates existing account.customer.statement.report.handler records.
// All records (represented by ids) will be updated by acsrh values.
func (c *Client) UpdateAccountCustomerStatementReportHandlers(ids []int64, acsrh *AccountCustomerStatementReportHandler) error {
	return c.Update(AccountCustomerStatementReportHandlerModel, ids, acsrh, nil)
}

// DeleteAccountCustomerStatementReportHandler deletes an existing account.customer.statement.report.handler record.
func (c *Client) DeleteAccountCustomerStatementReportHandler(id int64) error {
	return c.DeleteAccountCustomerStatementReportHandlers([]int64{id})
}

// DeleteAccountCustomerStatementReportHandlers deletes existing account.customer.statement.report.handler records.
func (c *Client) DeleteAccountCustomerStatementReportHandlers(ids []int64) error {
	return c.Delete(AccountCustomerStatementReportHandlerModel, ids)
}

// GetAccountCustomerStatementReportHandler gets account.customer.statement.report.handler existing record.
func (c *Client) GetAccountCustomerStatementReportHandler(id int64) (*AccountCustomerStatementReportHandler, error) {
	acsrhs, err := c.GetAccountCustomerStatementReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acsrhs)[0]), nil
}

// GetAccountCustomerStatementReportHandlers gets account.customer.statement.report.handler existing records.
func (c *Client) GetAccountCustomerStatementReportHandlers(ids []int64) (*AccountCustomerStatementReportHandlers, error) {
	acsrhs := &AccountCustomerStatementReportHandlers{}
	if err := c.Read(AccountCustomerStatementReportHandlerModel, ids, nil, acsrhs); err != nil {
		return nil, err
	}
	return acsrhs, nil
}

// FindAccountCustomerStatementReportHandler finds account.customer.statement.report.handler record by querying it with criteria.
func (c *Client) FindAccountCustomerStatementReportHandler(criteria *Criteria) (*AccountCustomerStatementReportHandler, error) {
	acsrhs := &AccountCustomerStatementReportHandlers{}
	if err := c.SearchRead(AccountCustomerStatementReportHandlerModel, criteria, NewOptions().Limit(1), acsrhs); err != nil {
		return nil, err
	}
	return &((*acsrhs)[0]), nil
}

// FindAccountCustomerStatementReportHandlers finds account.customer.statement.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCustomerStatementReportHandlers(criteria *Criteria, options *Options) (*AccountCustomerStatementReportHandlers, error) {
	acsrhs := &AccountCustomerStatementReportHandlers{}
	if err := c.SearchRead(AccountCustomerStatementReportHandlerModel, criteria, options, acsrhs); err != nil {
		return nil, err
	}
	return acsrhs, nil
}

// FindAccountCustomerStatementReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCustomerStatementReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountCustomerStatementReportHandlerModel, criteria, options)
}

// FindAccountCustomerStatementReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountCustomerStatementReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCustomerStatementReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
