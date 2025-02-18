package odoo

// AccountAgedPayableReportHandler represents account.aged.payable.report.handler model.
type AccountAgedPayableReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAgedPayableReportHandlers represents array of account.aged.payable.report.handler model.
type AccountAgedPayableReportHandlers []AccountAgedPayableReportHandler

// AccountAgedPayableReportHandlerModel is the odoo model name.
const AccountAgedPayableReportHandlerModel = "account.aged.payable.report.handler"

// Many2One convert AccountAgedPayableReportHandler to *Many2One.
func (aaprh *AccountAgedPayableReportHandler) Many2One() *Many2One {
	return NewMany2One(aaprh.Id.Get(), "")
}

// CreateAccountAgedPayableReportHandler creates a new account.aged.payable.report.handler model and returns its id.
func (c *Client) CreateAccountAgedPayableReportHandler(aaprh *AccountAgedPayableReportHandler) (int64, error) {
	ids, err := c.CreateAccountAgedPayableReportHandlers([]*AccountAgedPayableReportHandler{aaprh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAgedPayableReportHandler creates a new account.aged.payable.report.handler model and returns its id.
func (c *Client) CreateAccountAgedPayableReportHandlers(aaprhs []*AccountAgedPayableReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaprhs {
		vv = append(vv, v)
	}
	return c.Create(AccountAgedPayableReportHandlerModel, vv, nil)
}

// UpdateAccountAgedPayableReportHandler updates an existing account.aged.payable.report.handler record.
func (c *Client) UpdateAccountAgedPayableReportHandler(aaprh *AccountAgedPayableReportHandler) error {
	return c.UpdateAccountAgedPayableReportHandlers([]int64{aaprh.Id.Get()}, aaprh)
}

// UpdateAccountAgedPayableReportHandlers updates existing account.aged.payable.report.handler records.
// All records (represented by ids) will be updated by aaprh values.
func (c *Client) UpdateAccountAgedPayableReportHandlers(ids []int64, aaprh *AccountAgedPayableReportHandler) error {
	return c.Update(AccountAgedPayableReportHandlerModel, ids, aaprh, nil)
}

// DeleteAccountAgedPayableReportHandler deletes an existing account.aged.payable.report.handler record.
func (c *Client) DeleteAccountAgedPayableReportHandler(id int64) error {
	return c.DeleteAccountAgedPayableReportHandlers([]int64{id})
}

// DeleteAccountAgedPayableReportHandlers deletes existing account.aged.payable.report.handler records.
func (c *Client) DeleteAccountAgedPayableReportHandlers(ids []int64) error {
	return c.Delete(AccountAgedPayableReportHandlerModel, ids)
}

// GetAccountAgedPayableReportHandler gets account.aged.payable.report.handler existing record.
func (c *Client) GetAccountAgedPayableReportHandler(id int64) (*AccountAgedPayableReportHandler, error) {
	aaprhs, err := c.GetAccountAgedPayableReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaprhs)[0]), nil
}

// GetAccountAgedPayableReportHandlers gets account.aged.payable.report.handler existing records.
func (c *Client) GetAccountAgedPayableReportHandlers(ids []int64) (*AccountAgedPayableReportHandlers, error) {
	aaprhs := &AccountAgedPayableReportHandlers{}
	if err := c.Read(AccountAgedPayableReportHandlerModel, ids, nil, aaprhs); err != nil {
		return nil, err
	}
	return aaprhs, nil
}

// FindAccountAgedPayableReportHandler finds account.aged.payable.report.handler record by querying it with criteria.
func (c *Client) FindAccountAgedPayableReportHandler(criteria *Criteria) (*AccountAgedPayableReportHandler, error) {
	aaprhs := &AccountAgedPayableReportHandlers{}
	if err := c.SearchRead(AccountAgedPayableReportHandlerModel, criteria, NewOptions().Limit(1), aaprhs); err != nil {
		return nil, err
	}
	return &((*aaprhs)[0]), nil
}

// FindAccountAgedPayableReportHandlers finds account.aged.payable.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPayableReportHandlers(criteria *Criteria, options *Options) (*AccountAgedPayableReportHandlers, error) {
	aaprhs := &AccountAgedPayableReportHandlers{}
	if err := c.SearchRead(AccountAgedPayableReportHandlerModel, criteria, options, aaprhs); err != nil {
		return nil, err
	}
	return aaprhs, nil
}

// FindAccountAgedPayableReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPayableReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAgedPayableReportHandlerModel, criteria, options)
}

// FindAccountAgedPayableReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedPayableReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedPayableReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
