package odoo

// AccountDeferredReportHandler represents account.deferred.report.handler model.
type AccountDeferredReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountDeferredReportHandlers represents array of account.deferred.report.handler model.
type AccountDeferredReportHandlers []AccountDeferredReportHandler

// AccountDeferredReportHandlerModel is the odoo model name.
const AccountDeferredReportHandlerModel = "account.deferred.report.handler"

// Many2One convert AccountDeferredReportHandler to *Many2One.
func (adrh *AccountDeferredReportHandler) Many2One() *Many2One {
	return NewMany2One(adrh.Id.Get(), "")
}

// CreateAccountDeferredReportHandler creates a new account.deferred.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredReportHandler(adrh *AccountDeferredReportHandler) (int64, error) {
	ids, err := c.CreateAccountDeferredReportHandlers([]*AccountDeferredReportHandler{adrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDeferredReportHandler creates a new account.deferred.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredReportHandlers(adrhs []*AccountDeferredReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range adrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountDeferredReportHandlerModel, vv, nil)
}

// UpdateAccountDeferredReportHandler updates an existing account.deferred.report.handler record.
func (c *Client) UpdateAccountDeferredReportHandler(adrh *AccountDeferredReportHandler) error {
	return c.UpdateAccountDeferredReportHandlers([]int64{adrh.Id.Get()}, adrh)
}

// UpdateAccountDeferredReportHandlers updates existing account.deferred.report.handler records.
// All records (represented by ids) will be updated by adrh values.
func (c *Client) UpdateAccountDeferredReportHandlers(ids []int64, adrh *AccountDeferredReportHandler) error {
	return c.Update(AccountDeferredReportHandlerModel, ids, adrh, nil)
}

// DeleteAccountDeferredReportHandler deletes an existing account.deferred.report.handler record.
func (c *Client) DeleteAccountDeferredReportHandler(id int64) error {
	return c.DeleteAccountDeferredReportHandlers([]int64{id})
}

// DeleteAccountDeferredReportHandlers deletes existing account.deferred.report.handler records.
func (c *Client) DeleteAccountDeferredReportHandlers(ids []int64) error {
	return c.Delete(AccountDeferredReportHandlerModel, ids)
}

// GetAccountDeferredReportHandler gets account.deferred.report.handler existing record.
func (c *Client) GetAccountDeferredReportHandler(id int64) (*AccountDeferredReportHandler, error) {
	adrhs, err := c.GetAccountDeferredReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*adrhs)[0]), nil
}

// GetAccountDeferredReportHandlers gets account.deferred.report.handler existing records.
func (c *Client) GetAccountDeferredReportHandlers(ids []int64) (*AccountDeferredReportHandlers, error) {
	adrhs := &AccountDeferredReportHandlers{}
	if err := c.Read(AccountDeferredReportHandlerModel, ids, nil, adrhs); err != nil {
		return nil, err
	}
	return adrhs, nil
}

// FindAccountDeferredReportHandler finds account.deferred.report.handler record by querying it with criteria.
func (c *Client) FindAccountDeferredReportHandler(criteria *Criteria) (*AccountDeferredReportHandler, error) {
	adrhs := &AccountDeferredReportHandlers{}
	if err := c.SearchRead(AccountDeferredReportHandlerModel, criteria, NewOptions().Limit(1), adrhs); err != nil {
		return nil, err
	}
	return &((*adrhs)[0]), nil
}

// FindAccountDeferredReportHandlers finds account.deferred.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredReportHandlers(criteria *Criteria, options *Options) (*AccountDeferredReportHandlers, error) {
	adrhs := &AccountDeferredReportHandlers{}
	if err := c.SearchRead(AccountDeferredReportHandlerModel, criteria, options, adrhs); err != nil {
		return nil, err
	}
	return adrhs, nil
}

// FindAccountDeferredReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDeferredReportHandlerModel, criteria, options)
}

// FindAccountDeferredReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountDeferredReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDeferredReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
