package odoo

// AccountJournalReportHandler represents account.journal.report.handler model.
type AccountJournalReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountJournalReportHandlers represents array of account.journal.report.handler model.
type AccountJournalReportHandlers []AccountJournalReportHandler

// AccountJournalReportHandlerModel is the odoo model name.
const AccountJournalReportHandlerModel = "account.journal.report.handler"

// Many2One convert AccountJournalReportHandler to *Many2One.
func (ajrh *AccountJournalReportHandler) Many2One() *Many2One {
	return NewMany2One(ajrh.Id.Get(), "")
}

// CreateAccountJournalReportHandler creates a new account.journal.report.handler model and returns its id.
func (c *Client) CreateAccountJournalReportHandler(ajrh *AccountJournalReportHandler) (int64, error) {
	ids, err := c.CreateAccountJournalReportHandlers([]*AccountJournalReportHandler{ajrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountJournalReportHandler creates a new account.journal.report.handler model and returns its id.
func (c *Client) CreateAccountJournalReportHandlers(ajrhs []*AccountJournalReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range ajrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountJournalReportHandlerModel, vv, nil)
}

// UpdateAccountJournalReportHandler updates an existing account.journal.report.handler record.
func (c *Client) UpdateAccountJournalReportHandler(ajrh *AccountJournalReportHandler) error {
	return c.UpdateAccountJournalReportHandlers([]int64{ajrh.Id.Get()}, ajrh)
}

// UpdateAccountJournalReportHandlers updates existing account.journal.report.handler records.
// All records (represented by ids) will be updated by ajrh values.
func (c *Client) UpdateAccountJournalReportHandlers(ids []int64, ajrh *AccountJournalReportHandler) error {
	return c.Update(AccountJournalReportHandlerModel, ids, ajrh, nil)
}

// DeleteAccountJournalReportHandler deletes an existing account.journal.report.handler record.
func (c *Client) DeleteAccountJournalReportHandler(id int64) error {
	return c.DeleteAccountJournalReportHandlers([]int64{id})
}

// DeleteAccountJournalReportHandlers deletes existing account.journal.report.handler records.
func (c *Client) DeleteAccountJournalReportHandlers(ids []int64) error {
	return c.Delete(AccountJournalReportHandlerModel, ids)
}

// GetAccountJournalReportHandler gets account.journal.report.handler existing record.
func (c *Client) GetAccountJournalReportHandler(id int64) (*AccountJournalReportHandler, error) {
	ajrhs, err := c.GetAccountJournalReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ajrhs)[0]), nil
}

// GetAccountJournalReportHandlers gets account.journal.report.handler existing records.
func (c *Client) GetAccountJournalReportHandlers(ids []int64) (*AccountJournalReportHandlers, error) {
	ajrhs := &AccountJournalReportHandlers{}
	if err := c.Read(AccountJournalReportHandlerModel, ids, nil, ajrhs); err != nil {
		return nil, err
	}
	return ajrhs, nil
}

// FindAccountJournalReportHandler finds account.journal.report.handler record by querying it with criteria.
func (c *Client) FindAccountJournalReportHandler(criteria *Criteria) (*AccountJournalReportHandler, error) {
	ajrhs := &AccountJournalReportHandlers{}
	if err := c.SearchRead(AccountJournalReportHandlerModel, criteria, NewOptions().Limit(1), ajrhs); err != nil {
		return nil, err
	}
	return &((*ajrhs)[0]), nil
}

// FindAccountJournalReportHandlers finds account.journal.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalReportHandlers(criteria *Criteria, options *Options) (*AccountJournalReportHandlers, error) {
	ajrhs := &AccountJournalReportHandlers{}
	if err := c.SearchRead(AccountJournalReportHandlerModel, criteria, options, ajrhs); err != nil {
		return nil, err
	}
	return ajrhs, nil
}

// FindAccountJournalReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountJournalReportHandlerModel, criteria, options)
}

// FindAccountJournalReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountJournalReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
