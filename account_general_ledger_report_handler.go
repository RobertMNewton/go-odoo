package odoo

// AccountGeneralLedgerReportHandler represents account.general.ledger.report.handler model.
type AccountGeneralLedgerReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountGeneralLedgerReportHandlers represents array of account.general.ledger.report.handler model.
type AccountGeneralLedgerReportHandlers []AccountGeneralLedgerReportHandler

// AccountGeneralLedgerReportHandlerModel is the odoo model name.
const AccountGeneralLedgerReportHandlerModel = "account.general.ledger.report.handler"

// Many2One convert AccountGeneralLedgerReportHandler to *Many2One.
func (aglrh *AccountGeneralLedgerReportHandler) Many2One() *Many2One {
	return NewMany2One(aglrh.Id.Get(), "")
}

// CreateAccountGeneralLedgerReportHandler creates a new account.general.ledger.report.handler model and returns its id.
func (c *Client) CreateAccountGeneralLedgerReportHandler(aglrh *AccountGeneralLedgerReportHandler) (int64, error) {
	ids, err := c.CreateAccountGeneralLedgerReportHandlers([]*AccountGeneralLedgerReportHandler{aglrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGeneralLedgerReportHandler creates a new account.general.ledger.report.handler model and returns its id.
func (c *Client) CreateAccountGeneralLedgerReportHandlers(aglrhs []*AccountGeneralLedgerReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aglrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountGeneralLedgerReportHandlerModel, vv, nil)
}

// UpdateAccountGeneralLedgerReportHandler updates an existing account.general.ledger.report.handler record.
func (c *Client) UpdateAccountGeneralLedgerReportHandler(aglrh *AccountGeneralLedgerReportHandler) error {
	return c.UpdateAccountGeneralLedgerReportHandlers([]int64{aglrh.Id.Get()}, aglrh)
}

// UpdateAccountGeneralLedgerReportHandlers updates existing account.general.ledger.report.handler records.
// All records (represented by ids) will be updated by aglrh values.
func (c *Client) UpdateAccountGeneralLedgerReportHandlers(ids []int64, aglrh *AccountGeneralLedgerReportHandler) error {
	return c.Update(AccountGeneralLedgerReportHandlerModel, ids, aglrh, nil)
}

// DeleteAccountGeneralLedgerReportHandler deletes an existing account.general.ledger.report.handler record.
func (c *Client) DeleteAccountGeneralLedgerReportHandler(id int64) error {
	return c.DeleteAccountGeneralLedgerReportHandlers([]int64{id})
}

// DeleteAccountGeneralLedgerReportHandlers deletes existing account.general.ledger.report.handler records.
func (c *Client) DeleteAccountGeneralLedgerReportHandlers(ids []int64) error {
	return c.Delete(AccountGeneralLedgerReportHandlerModel, ids)
}

// GetAccountGeneralLedgerReportHandler gets account.general.ledger.report.handler existing record.
func (c *Client) GetAccountGeneralLedgerReportHandler(id int64) (*AccountGeneralLedgerReportHandler, error) {
	aglrhs, err := c.GetAccountGeneralLedgerReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aglrhs)[0]), nil
}

// GetAccountGeneralLedgerReportHandlers gets account.general.ledger.report.handler existing records.
func (c *Client) GetAccountGeneralLedgerReportHandlers(ids []int64) (*AccountGeneralLedgerReportHandlers, error) {
	aglrhs := &AccountGeneralLedgerReportHandlers{}
	if err := c.Read(AccountGeneralLedgerReportHandlerModel, ids, nil, aglrhs); err != nil {
		return nil, err
	}
	return aglrhs, nil
}

// FindAccountGeneralLedgerReportHandler finds account.general.ledger.report.handler record by querying it with criteria.
func (c *Client) FindAccountGeneralLedgerReportHandler(criteria *Criteria) (*AccountGeneralLedgerReportHandler, error) {
	aglrhs := &AccountGeneralLedgerReportHandlers{}
	if err := c.SearchRead(AccountGeneralLedgerReportHandlerModel, criteria, NewOptions().Limit(1), aglrhs); err != nil {
		return nil, err
	}
	return &((*aglrhs)[0]), nil
}

// FindAccountGeneralLedgerReportHandlers finds account.general.ledger.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgerReportHandlers(criteria *Criteria, options *Options) (*AccountGeneralLedgerReportHandlers, error) {
	aglrhs := &AccountGeneralLedgerReportHandlers{}
	if err := c.SearchRead(AccountGeneralLedgerReportHandlerModel, criteria, options, aglrhs); err != nil {
		return nil, err
	}
	return aglrhs, nil
}

// FindAccountGeneralLedgerReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgerReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGeneralLedgerReportHandlerModel, criteria, options)
}

// FindAccountGeneralLedgerReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountGeneralLedgerReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGeneralLedgerReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
