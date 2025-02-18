package odoo

// AccountPartnerLedgerReportHandler represents account.partner.ledger.report.handler model.
type AccountPartnerLedgerReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountPartnerLedgerReportHandlers represents array of account.partner.ledger.report.handler model.
type AccountPartnerLedgerReportHandlers []AccountPartnerLedgerReportHandler

// AccountPartnerLedgerReportHandlerModel is the odoo model name.
const AccountPartnerLedgerReportHandlerModel = "account.partner.ledger.report.handler"

// Many2One convert AccountPartnerLedgerReportHandler to *Many2One.
func (aplrh *AccountPartnerLedgerReportHandler) Many2One() *Many2One {
	return NewMany2One(aplrh.Id.Get(), "")
}

// CreateAccountPartnerLedgerReportHandler creates a new account.partner.ledger.report.handler model and returns its id.
func (c *Client) CreateAccountPartnerLedgerReportHandler(aplrh *AccountPartnerLedgerReportHandler) (int64, error) {
	ids, err := c.CreateAccountPartnerLedgerReportHandlers([]*AccountPartnerLedgerReportHandler{aplrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPartnerLedgerReportHandler creates a new account.partner.ledger.report.handler model and returns its id.
func (c *Client) CreateAccountPartnerLedgerReportHandlers(aplrhs []*AccountPartnerLedgerReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aplrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountPartnerLedgerReportHandlerModel, vv, nil)
}

// UpdateAccountPartnerLedgerReportHandler updates an existing account.partner.ledger.report.handler record.
func (c *Client) UpdateAccountPartnerLedgerReportHandler(aplrh *AccountPartnerLedgerReportHandler) error {
	return c.UpdateAccountPartnerLedgerReportHandlers([]int64{aplrh.Id.Get()}, aplrh)
}

// UpdateAccountPartnerLedgerReportHandlers updates existing account.partner.ledger.report.handler records.
// All records (represented by ids) will be updated by aplrh values.
func (c *Client) UpdateAccountPartnerLedgerReportHandlers(ids []int64, aplrh *AccountPartnerLedgerReportHandler) error {
	return c.Update(AccountPartnerLedgerReportHandlerModel, ids, aplrh, nil)
}

// DeleteAccountPartnerLedgerReportHandler deletes an existing account.partner.ledger.report.handler record.
func (c *Client) DeleteAccountPartnerLedgerReportHandler(id int64) error {
	return c.DeleteAccountPartnerLedgerReportHandlers([]int64{id})
}

// DeleteAccountPartnerLedgerReportHandlers deletes existing account.partner.ledger.report.handler records.
func (c *Client) DeleteAccountPartnerLedgerReportHandlers(ids []int64) error {
	return c.Delete(AccountPartnerLedgerReportHandlerModel, ids)
}

// GetAccountPartnerLedgerReportHandler gets account.partner.ledger.report.handler existing record.
func (c *Client) GetAccountPartnerLedgerReportHandler(id int64) (*AccountPartnerLedgerReportHandler, error) {
	aplrhs, err := c.GetAccountPartnerLedgerReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aplrhs)[0]), nil
}

// GetAccountPartnerLedgerReportHandlers gets account.partner.ledger.report.handler existing records.
func (c *Client) GetAccountPartnerLedgerReportHandlers(ids []int64) (*AccountPartnerLedgerReportHandlers, error) {
	aplrhs := &AccountPartnerLedgerReportHandlers{}
	if err := c.Read(AccountPartnerLedgerReportHandlerModel, ids, nil, aplrhs); err != nil {
		return nil, err
	}
	return aplrhs, nil
}

// FindAccountPartnerLedgerReportHandler finds account.partner.ledger.report.handler record by querying it with criteria.
func (c *Client) FindAccountPartnerLedgerReportHandler(criteria *Criteria) (*AccountPartnerLedgerReportHandler, error) {
	aplrhs := &AccountPartnerLedgerReportHandlers{}
	if err := c.SearchRead(AccountPartnerLedgerReportHandlerModel, criteria, NewOptions().Limit(1), aplrhs); err != nil {
		return nil, err
	}
	return &((*aplrhs)[0]), nil
}

// FindAccountPartnerLedgerReportHandlers finds account.partner.ledger.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartnerLedgerReportHandlers(criteria *Criteria, options *Options) (*AccountPartnerLedgerReportHandlers, error) {
	aplrhs := &AccountPartnerLedgerReportHandlers{}
	if err := c.SearchRead(AccountPartnerLedgerReportHandlerModel, criteria, options, aplrhs); err != nil {
		return nil, err
	}
	return aplrhs, nil
}

// FindAccountPartnerLedgerReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartnerLedgerReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPartnerLedgerReportHandlerModel, criteria, options)
}

// FindAccountPartnerLedgerReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountPartnerLedgerReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPartnerLedgerReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
