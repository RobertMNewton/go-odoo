package odoo

// AccountTaxReportHandler represents account.tax.report.handler model.
type AccountTaxReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountTaxReportHandlers represents array of account.tax.report.handler model.
type AccountTaxReportHandlers []AccountTaxReportHandler

// AccountTaxReportHandlerModel is the odoo model name.
const AccountTaxReportHandlerModel = "account.tax.report.handler"

// Many2One convert AccountTaxReportHandler to *Many2One.
func (atrh *AccountTaxReportHandler) Many2One() *Many2One {
	return NewMany2One(atrh.Id.Get(), "")
}

// CreateAccountTaxReportHandler creates a new account.tax.report.handler model and returns its id.
func (c *Client) CreateAccountTaxReportHandler(atrh *AccountTaxReportHandler) (int64, error) {
	ids, err := c.CreateAccountTaxReportHandlers([]*AccountTaxReportHandler{atrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTaxReportHandler creates a new account.tax.report.handler model and returns its id.
func (c *Client) CreateAccountTaxReportHandlers(atrhs []*AccountTaxReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range atrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxReportHandlerModel, vv, nil)
}

// UpdateAccountTaxReportHandler updates an existing account.tax.report.handler record.
func (c *Client) UpdateAccountTaxReportHandler(atrh *AccountTaxReportHandler) error {
	return c.UpdateAccountTaxReportHandlers([]int64{atrh.Id.Get()}, atrh)
}

// UpdateAccountTaxReportHandlers updates existing account.tax.report.handler records.
// All records (represented by ids) will be updated by atrh values.
func (c *Client) UpdateAccountTaxReportHandlers(ids []int64, atrh *AccountTaxReportHandler) error {
	return c.Update(AccountTaxReportHandlerModel, ids, atrh, nil)
}

// DeleteAccountTaxReportHandler deletes an existing account.tax.report.handler record.
func (c *Client) DeleteAccountTaxReportHandler(id int64) error {
	return c.DeleteAccountTaxReportHandlers([]int64{id})
}

// DeleteAccountTaxReportHandlers deletes existing account.tax.report.handler records.
func (c *Client) DeleteAccountTaxReportHandlers(ids []int64) error {
	return c.Delete(AccountTaxReportHandlerModel, ids)
}

// GetAccountTaxReportHandler gets account.tax.report.handler existing record.
func (c *Client) GetAccountTaxReportHandler(id int64) (*AccountTaxReportHandler, error) {
	atrhs, err := c.GetAccountTaxReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atrhs)[0]), nil
}

// GetAccountTaxReportHandlers gets account.tax.report.handler existing records.
func (c *Client) GetAccountTaxReportHandlers(ids []int64) (*AccountTaxReportHandlers, error) {
	atrhs := &AccountTaxReportHandlers{}
	if err := c.Read(AccountTaxReportHandlerModel, ids, nil, atrhs); err != nil {
		return nil, err
	}
	return atrhs, nil
}

// FindAccountTaxReportHandler finds account.tax.report.handler record by querying it with criteria.
func (c *Client) FindAccountTaxReportHandler(criteria *Criteria) (*AccountTaxReportHandler, error) {
	atrhs := &AccountTaxReportHandlers{}
	if err := c.SearchRead(AccountTaxReportHandlerModel, criteria, NewOptions().Limit(1), atrhs); err != nil {
		return nil, err
	}
	return &((*atrhs)[0]), nil
}

// FindAccountTaxReportHandlers finds account.tax.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReportHandlers(criteria *Criteria, options *Options) (*AccountTaxReportHandlers, error) {
	atrhs := &AccountTaxReportHandlers{}
	if err := c.SearchRead(AccountTaxReportHandlerModel, criteria, options, atrhs); err != nil {
		return nil, err
	}
	return atrhs, nil
}

// FindAccountTaxReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTaxReportHandlerModel, criteria, options)
}

// FindAccountTaxReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
