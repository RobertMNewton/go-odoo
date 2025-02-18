package odoo

// AccountBalanceSheetReportHandler represents account.balance.sheet.report.handler model.
type AccountBalanceSheetReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountBalanceSheetReportHandlers represents array of account.balance.sheet.report.handler model.
type AccountBalanceSheetReportHandlers []AccountBalanceSheetReportHandler

// AccountBalanceSheetReportHandlerModel is the odoo model name.
const AccountBalanceSheetReportHandlerModel = "account.balance.sheet.report.handler"

// Many2One convert AccountBalanceSheetReportHandler to *Many2One.
func (absrh *AccountBalanceSheetReportHandler) Many2One() *Many2One {
	return NewMany2One(absrh.Id.Get(), "")
}

// CreateAccountBalanceSheetReportHandler creates a new account.balance.sheet.report.handler model and returns its id.
func (c *Client) CreateAccountBalanceSheetReportHandler(absrh *AccountBalanceSheetReportHandler) (int64, error) {
	ids, err := c.CreateAccountBalanceSheetReportHandlers([]*AccountBalanceSheetReportHandler{absrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBalanceSheetReportHandler creates a new account.balance.sheet.report.handler model and returns its id.
func (c *Client) CreateAccountBalanceSheetReportHandlers(absrhs []*AccountBalanceSheetReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range absrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountBalanceSheetReportHandlerModel, vv, nil)
}

// UpdateAccountBalanceSheetReportHandler updates an existing account.balance.sheet.report.handler record.
func (c *Client) UpdateAccountBalanceSheetReportHandler(absrh *AccountBalanceSheetReportHandler) error {
	return c.UpdateAccountBalanceSheetReportHandlers([]int64{absrh.Id.Get()}, absrh)
}

// UpdateAccountBalanceSheetReportHandlers updates existing account.balance.sheet.report.handler records.
// All records (represented by ids) will be updated by absrh values.
func (c *Client) UpdateAccountBalanceSheetReportHandlers(ids []int64, absrh *AccountBalanceSheetReportHandler) error {
	return c.Update(AccountBalanceSheetReportHandlerModel, ids, absrh, nil)
}

// DeleteAccountBalanceSheetReportHandler deletes an existing account.balance.sheet.report.handler record.
func (c *Client) DeleteAccountBalanceSheetReportHandler(id int64) error {
	return c.DeleteAccountBalanceSheetReportHandlers([]int64{id})
}

// DeleteAccountBalanceSheetReportHandlers deletes existing account.balance.sheet.report.handler records.
func (c *Client) DeleteAccountBalanceSheetReportHandlers(ids []int64) error {
	return c.Delete(AccountBalanceSheetReportHandlerModel, ids)
}

// GetAccountBalanceSheetReportHandler gets account.balance.sheet.report.handler existing record.
func (c *Client) GetAccountBalanceSheetReportHandler(id int64) (*AccountBalanceSheetReportHandler, error) {
	absrhs, err := c.GetAccountBalanceSheetReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*absrhs)[0]), nil
}

// GetAccountBalanceSheetReportHandlers gets account.balance.sheet.report.handler existing records.
func (c *Client) GetAccountBalanceSheetReportHandlers(ids []int64) (*AccountBalanceSheetReportHandlers, error) {
	absrhs := &AccountBalanceSheetReportHandlers{}
	if err := c.Read(AccountBalanceSheetReportHandlerModel, ids, nil, absrhs); err != nil {
		return nil, err
	}
	return absrhs, nil
}

// FindAccountBalanceSheetReportHandler finds account.balance.sheet.report.handler record by querying it with criteria.
func (c *Client) FindAccountBalanceSheetReportHandler(criteria *Criteria) (*AccountBalanceSheetReportHandler, error) {
	absrhs := &AccountBalanceSheetReportHandlers{}
	if err := c.SearchRead(AccountBalanceSheetReportHandlerModel, criteria, NewOptions().Limit(1), absrhs); err != nil {
		return nil, err
	}
	return &((*absrhs)[0]), nil
}

// FindAccountBalanceSheetReportHandlers finds account.balance.sheet.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBalanceSheetReportHandlers(criteria *Criteria, options *Options) (*AccountBalanceSheetReportHandlers, error) {
	absrhs := &AccountBalanceSheetReportHandlers{}
	if err := c.SearchRead(AccountBalanceSheetReportHandlerModel, criteria, options, absrhs); err != nil {
		return nil, err
	}
	return absrhs, nil
}

// FindAccountBalanceSheetReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBalanceSheetReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBalanceSheetReportHandlerModel, criteria, options)
}

// FindAccountBalanceSheetReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountBalanceSheetReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBalanceSheetReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
