package odoo

// AccountEcSalesReportHandler represents account.ec.sales.report.handler model.
type AccountEcSalesReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountEcSalesReportHandlers represents array of account.ec.sales.report.handler model.
type AccountEcSalesReportHandlers []AccountEcSalesReportHandler

// AccountEcSalesReportHandlerModel is the odoo model name.
const AccountEcSalesReportHandlerModel = "account.ec.sales.report.handler"

// Many2One convert AccountEcSalesReportHandler to *Many2One.
func (aesrh *AccountEcSalesReportHandler) Many2One() *Many2One {
	return NewMany2One(aesrh.Id.Get(), "")
}

// CreateAccountEcSalesReportHandler creates a new account.ec.sales.report.handler model and returns its id.
func (c *Client) CreateAccountEcSalesReportHandler(aesrh *AccountEcSalesReportHandler) (int64, error) {
	ids, err := c.CreateAccountEcSalesReportHandlers([]*AccountEcSalesReportHandler{aesrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountEcSalesReportHandler creates a new account.ec.sales.report.handler model and returns its id.
func (c *Client) CreateAccountEcSalesReportHandlers(aesrhs []*AccountEcSalesReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aesrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountEcSalesReportHandlerModel, vv, nil)
}

// UpdateAccountEcSalesReportHandler updates an existing account.ec.sales.report.handler record.
func (c *Client) UpdateAccountEcSalesReportHandler(aesrh *AccountEcSalesReportHandler) error {
	return c.UpdateAccountEcSalesReportHandlers([]int64{aesrh.Id.Get()}, aesrh)
}

// UpdateAccountEcSalesReportHandlers updates existing account.ec.sales.report.handler records.
// All records (represented by ids) will be updated by aesrh values.
func (c *Client) UpdateAccountEcSalesReportHandlers(ids []int64, aesrh *AccountEcSalesReportHandler) error {
	return c.Update(AccountEcSalesReportHandlerModel, ids, aesrh, nil)
}

// DeleteAccountEcSalesReportHandler deletes an existing account.ec.sales.report.handler record.
func (c *Client) DeleteAccountEcSalesReportHandler(id int64) error {
	return c.DeleteAccountEcSalesReportHandlers([]int64{id})
}

// DeleteAccountEcSalesReportHandlers deletes existing account.ec.sales.report.handler records.
func (c *Client) DeleteAccountEcSalesReportHandlers(ids []int64) error {
	return c.Delete(AccountEcSalesReportHandlerModel, ids)
}

// GetAccountEcSalesReportHandler gets account.ec.sales.report.handler existing record.
func (c *Client) GetAccountEcSalesReportHandler(id int64) (*AccountEcSalesReportHandler, error) {
	aesrhs, err := c.GetAccountEcSalesReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aesrhs)[0]), nil
}

// GetAccountEcSalesReportHandlers gets account.ec.sales.report.handler existing records.
func (c *Client) GetAccountEcSalesReportHandlers(ids []int64) (*AccountEcSalesReportHandlers, error) {
	aesrhs := &AccountEcSalesReportHandlers{}
	if err := c.Read(AccountEcSalesReportHandlerModel, ids, nil, aesrhs); err != nil {
		return nil, err
	}
	return aesrhs, nil
}

// FindAccountEcSalesReportHandler finds account.ec.sales.report.handler record by querying it with criteria.
func (c *Client) FindAccountEcSalesReportHandler(criteria *Criteria) (*AccountEcSalesReportHandler, error) {
	aesrhs := &AccountEcSalesReportHandlers{}
	if err := c.SearchRead(AccountEcSalesReportHandlerModel, criteria, NewOptions().Limit(1), aesrhs); err != nil {
		return nil, err
	}
	return &((*aesrhs)[0]), nil
}

// FindAccountEcSalesReportHandlers finds account.ec.sales.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEcSalesReportHandlers(criteria *Criteria, options *Options) (*AccountEcSalesReportHandlers, error) {
	aesrhs := &AccountEcSalesReportHandlers{}
	if err := c.SearchRead(AccountEcSalesReportHandlerModel, criteria, options, aesrhs); err != nil {
		return nil, err
	}
	return aesrhs, nil
}

// FindAccountEcSalesReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEcSalesReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountEcSalesReportHandlerModel, criteria, options)
}

// FindAccountEcSalesReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountEcSalesReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEcSalesReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
