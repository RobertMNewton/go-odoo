package odoo

// AccountDeferredRevenueReportHandler represents account.deferred.revenue.report.handler model.
type AccountDeferredRevenueReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountDeferredRevenueReportHandlers represents array of account.deferred.revenue.report.handler model.
type AccountDeferredRevenueReportHandlers []AccountDeferredRevenueReportHandler

// AccountDeferredRevenueReportHandlerModel is the odoo model name.
const AccountDeferredRevenueReportHandlerModel = "account.deferred.revenue.report.handler"

// Many2One convert AccountDeferredRevenueReportHandler to *Many2One.
func (adrrh *AccountDeferredRevenueReportHandler) Many2One() *Many2One {
	return NewMany2One(adrrh.Id.Get(), "")
}

// CreateAccountDeferredRevenueReportHandler creates a new account.deferred.revenue.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredRevenueReportHandler(adrrh *AccountDeferredRevenueReportHandler) (int64, error) {
	ids, err := c.CreateAccountDeferredRevenueReportHandlers([]*AccountDeferredRevenueReportHandler{adrrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDeferredRevenueReportHandler creates a new account.deferred.revenue.report.handler model and returns its id.
func (c *Client) CreateAccountDeferredRevenueReportHandlers(adrrhs []*AccountDeferredRevenueReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range adrrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountDeferredRevenueReportHandlerModel, vv, nil)
}

// UpdateAccountDeferredRevenueReportHandler updates an existing account.deferred.revenue.report.handler record.
func (c *Client) UpdateAccountDeferredRevenueReportHandler(adrrh *AccountDeferredRevenueReportHandler) error {
	return c.UpdateAccountDeferredRevenueReportHandlers([]int64{adrrh.Id.Get()}, adrrh)
}

// UpdateAccountDeferredRevenueReportHandlers updates existing account.deferred.revenue.report.handler records.
// All records (represented by ids) will be updated by adrrh values.
func (c *Client) UpdateAccountDeferredRevenueReportHandlers(ids []int64, adrrh *AccountDeferredRevenueReportHandler) error {
	return c.Update(AccountDeferredRevenueReportHandlerModel, ids, adrrh, nil)
}

// DeleteAccountDeferredRevenueReportHandler deletes an existing account.deferred.revenue.report.handler record.
func (c *Client) DeleteAccountDeferredRevenueReportHandler(id int64) error {
	return c.DeleteAccountDeferredRevenueReportHandlers([]int64{id})
}

// DeleteAccountDeferredRevenueReportHandlers deletes existing account.deferred.revenue.report.handler records.
func (c *Client) DeleteAccountDeferredRevenueReportHandlers(ids []int64) error {
	return c.Delete(AccountDeferredRevenueReportHandlerModel, ids)
}

// GetAccountDeferredRevenueReportHandler gets account.deferred.revenue.report.handler existing record.
func (c *Client) GetAccountDeferredRevenueReportHandler(id int64) (*AccountDeferredRevenueReportHandler, error) {
	adrrhs, err := c.GetAccountDeferredRevenueReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*adrrhs)[0]), nil
}

// GetAccountDeferredRevenueReportHandlers gets account.deferred.revenue.report.handler existing records.
func (c *Client) GetAccountDeferredRevenueReportHandlers(ids []int64) (*AccountDeferredRevenueReportHandlers, error) {
	adrrhs := &AccountDeferredRevenueReportHandlers{}
	if err := c.Read(AccountDeferredRevenueReportHandlerModel, ids, nil, adrrhs); err != nil {
		return nil, err
	}
	return adrrhs, nil
}

// FindAccountDeferredRevenueReportHandler finds account.deferred.revenue.report.handler record by querying it with criteria.
func (c *Client) FindAccountDeferredRevenueReportHandler(criteria *Criteria) (*AccountDeferredRevenueReportHandler, error) {
	adrrhs := &AccountDeferredRevenueReportHandlers{}
	if err := c.SearchRead(AccountDeferredRevenueReportHandlerModel, criteria, NewOptions().Limit(1), adrrhs); err != nil {
		return nil, err
	}
	return &((*adrrhs)[0]), nil
}

// FindAccountDeferredRevenueReportHandlers finds account.deferred.revenue.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredRevenueReportHandlers(criteria *Criteria, options *Options) (*AccountDeferredRevenueReportHandlers, error) {
	adrrhs := &AccountDeferredRevenueReportHandlers{}
	if err := c.SearchRead(AccountDeferredRevenueReportHandlerModel, criteria, options, adrrhs); err != nil {
		return nil, err
	}
	return adrrhs, nil
}

// FindAccountDeferredRevenueReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDeferredRevenueReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDeferredRevenueReportHandlerModel, criteria, options)
}

// FindAccountDeferredRevenueReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountDeferredRevenueReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDeferredRevenueReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
