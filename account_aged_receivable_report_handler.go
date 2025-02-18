package odoo

// AccountAgedReceivableReportHandler represents account.aged.receivable.report.handler model.
type AccountAgedReceivableReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAgedReceivableReportHandlers represents array of account.aged.receivable.report.handler model.
type AccountAgedReceivableReportHandlers []AccountAgedReceivableReportHandler

// AccountAgedReceivableReportHandlerModel is the odoo model name.
const AccountAgedReceivableReportHandlerModel = "account.aged.receivable.report.handler"

// Many2One convert AccountAgedReceivableReportHandler to *Many2One.
func (aarrh *AccountAgedReceivableReportHandler) Many2One() *Many2One {
	return NewMany2One(aarrh.Id.Get(), "")
}

// CreateAccountAgedReceivableReportHandler creates a new account.aged.receivable.report.handler model and returns its id.
func (c *Client) CreateAccountAgedReceivableReportHandler(aarrh *AccountAgedReceivableReportHandler) (int64, error) {
	ids, err := c.CreateAccountAgedReceivableReportHandlers([]*AccountAgedReceivableReportHandler{aarrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAgedReceivableReportHandler creates a new account.aged.receivable.report.handler model and returns its id.
func (c *Client) CreateAccountAgedReceivableReportHandlers(aarrhs []*AccountAgedReceivableReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aarrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountAgedReceivableReportHandlerModel, vv, nil)
}

// UpdateAccountAgedReceivableReportHandler updates an existing account.aged.receivable.report.handler record.
func (c *Client) UpdateAccountAgedReceivableReportHandler(aarrh *AccountAgedReceivableReportHandler) error {
	return c.UpdateAccountAgedReceivableReportHandlers([]int64{aarrh.Id.Get()}, aarrh)
}

// UpdateAccountAgedReceivableReportHandlers updates existing account.aged.receivable.report.handler records.
// All records (represented by ids) will be updated by aarrh values.
func (c *Client) UpdateAccountAgedReceivableReportHandlers(ids []int64, aarrh *AccountAgedReceivableReportHandler) error {
	return c.Update(AccountAgedReceivableReportHandlerModel, ids, aarrh, nil)
}

// DeleteAccountAgedReceivableReportHandler deletes an existing account.aged.receivable.report.handler record.
func (c *Client) DeleteAccountAgedReceivableReportHandler(id int64) error {
	return c.DeleteAccountAgedReceivableReportHandlers([]int64{id})
}

// DeleteAccountAgedReceivableReportHandlers deletes existing account.aged.receivable.report.handler records.
func (c *Client) DeleteAccountAgedReceivableReportHandlers(ids []int64) error {
	return c.Delete(AccountAgedReceivableReportHandlerModel, ids)
}

// GetAccountAgedReceivableReportHandler gets account.aged.receivable.report.handler existing record.
func (c *Client) GetAccountAgedReceivableReportHandler(id int64) (*AccountAgedReceivableReportHandler, error) {
	aarrhs, err := c.GetAccountAgedReceivableReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aarrhs)[0]), nil
}

// GetAccountAgedReceivableReportHandlers gets account.aged.receivable.report.handler existing records.
func (c *Client) GetAccountAgedReceivableReportHandlers(ids []int64) (*AccountAgedReceivableReportHandlers, error) {
	aarrhs := &AccountAgedReceivableReportHandlers{}
	if err := c.Read(AccountAgedReceivableReportHandlerModel, ids, nil, aarrhs); err != nil {
		return nil, err
	}
	return aarrhs, nil
}

// FindAccountAgedReceivableReportHandler finds account.aged.receivable.report.handler record by querying it with criteria.
func (c *Client) FindAccountAgedReceivableReportHandler(criteria *Criteria) (*AccountAgedReceivableReportHandler, error) {
	aarrhs := &AccountAgedReceivableReportHandlers{}
	if err := c.SearchRead(AccountAgedReceivableReportHandlerModel, criteria, NewOptions().Limit(1), aarrhs); err != nil {
		return nil, err
	}
	return &((*aarrhs)[0]), nil
}

// FindAccountAgedReceivableReportHandlers finds account.aged.receivable.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedReceivableReportHandlers(criteria *Criteria, options *Options) (*AccountAgedReceivableReportHandlers, error) {
	aarrhs := &AccountAgedReceivableReportHandlers{}
	if err := c.SearchRead(AccountAgedReceivableReportHandlerModel, criteria, options, aarrhs); err != nil {
		return nil, err
	}
	return aarrhs, nil
}

// FindAccountAgedReceivableReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedReceivableReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAgedReceivableReportHandlerModel, criteria, options)
}

// FindAccountAgedReceivableReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedReceivableReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedReceivableReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
