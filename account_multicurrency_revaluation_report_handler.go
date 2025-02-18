package odoo

// AccountMulticurrencyRevaluationReportHandler represents account.multicurrency.revaluation.report.handler model.
type AccountMulticurrencyRevaluationReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountMulticurrencyRevaluationReportHandlers represents array of account.multicurrency.revaluation.report.handler model.
type AccountMulticurrencyRevaluationReportHandlers []AccountMulticurrencyRevaluationReportHandler

// AccountMulticurrencyRevaluationReportHandlerModel is the odoo model name.
const AccountMulticurrencyRevaluationReportHandlerModel = "account.multicurrency.revaluation.report.handler"

// Many2One convert AccountMulticurrencyRevaluationReportHandler to *Many2One.
func (amrrh *AccountMulticurrencyRevaluationReportHandler) Many2One() *Many2One {
	return NewMany2One(amrrh.Id.Get(), "")
}

// CreateAccountMulticurrencyRevaluationReportHandler creates a new account.multicurrency.revaluation.report.handler model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationReportHandler(amrrh *AccountMulticurrencyRevaluationReportHandler) (int64, error) {
	ids, err := c.CreateAccountMulticurrencyRevaluationReportHandlers([]*AccountMulticurrencyRevaluationReportHandler{amrrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMulticurrencyRevaluationReportHandler creates a new account.multicurrency.revaluation.report.handler model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationReportHandlers(amrrhs []*AccountMulticurrencyRevaluationReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range amrrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountMulticurrencyRevaluationReportHandlerModel, vv, nil)
}

// UpdateAccountMulticurrencyRevaluationReportHandler updates an existing account.multicurrency.revaluation.report.handler record.
func (c *Client) UpdateAccountMulticurrencyRevaluationReportHandler(amrrh *AccountMulticurrencyRevaluationReportHandler) error {
	return c.UpdateAccountMulticurrencyRevaluationReportHandlers([]int64{amrrh.Id.Get()}, amrrh)
}

// UpdateAccountMulticurrencyRevaluationReportHandlers updates existing account.multicurrency.revaluation.report.handler records.
// All records (represented by ids) will be updated by amrrh values.
func (c *Client) UpdateAccountMulticurrencyRevaluationReportHandlers(ids []int64, amrrh *AccountMulticurrencyRevaluationReportHandler) error {
	return c.Update(AccountMulticurrencyRevaluationReportHandlerModel, ids, amrrh, nil)
}

// DeleteAccountMulticurrencyRevaluationReportHandler deletes an existing account.multicurrency.revaluation.report.handler record.
func (c *Client) DeleteAccountMulticurrencyRevaluationReportHandler(id int64) error {
	return c.DeleteAccountMulticurrencyRevaluationReportHandlers([]int64{id})
}

// DeleteAccountMulticurrencyRevaluationReportHandlers deletes existing account.multicurrency.revaluation.report.handler records.
func (c *Client) DeleteAccountMulticurrencyRevaluationReportHandlers(ids []int64) error {
	return c.Delete(AccountMulticurrencyRevaluationReportHandlerModel, ids)
}

// GetAccountMulticurrencyRevaluationReportHandler gets account.multicurrency.revaluation.report.handler existing record.
func (c *Client) GetAccountMulticurrencyRevaluationReportHandler(id int64) (*AccountMulticurrencyRevaluationReportHandler, error) {
	amrrhs, err := c.GetAccountMulticurrencyRevaluationReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amrrhs)[0]), nil
}

// GetAccountMulticurrencyRevaluationReportHandlers gets account.multicurrency.revaluation.report.handler existing records.
func (c *Client) GetAccountMulticurrencyRevaluationReportHandlers(ids []int64) (*AccountMulticurrencyRevaluationReportHandlers, error) {
	amrrhs := &AccountMulticurrencyRevaluationReportHandlers{}
	if err := c.Read(AccountMulticurrencyRevaluationReportHandlerModel, ids, nil, amrrhs); err != nil {
		return nil, err
	}
	return amrrhs, nil
}

// FindAccountMulticurrencyRevaluationReportHandler finds account.multicurrency.revaluation.report.handler record by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationReportHandler(criteria *Criteria) (*AccountMulticurrencyRevaluationReportHandler, error) {
	amrrhs := &AccountMulticurrencyRevaluationReportHandlers{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationReportHandlerModel, criteria, NewOptions().Limit(1), amrrhs); err != nil {
		return nil, err
	}
	return &((*amrrhs)[0]), nil
}

// FindAccountMulticurrencyRevaluationReportHandlers finds account.multicurrency.revaluation.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationReportHandlers(criteria *Criteria, options *Options) (*AccountMulticurrencyRevaluationReportHandlers, error) {
	amrrhs := &AccountMulticurrencyRevaluationReportHandlers{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationReportHandlerModel, criteria, options, amrrhs); err != nil {
		return nil, err
	}
	return amrrhs, nil
}

// FindAccountMulticurrencyRevaluationReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMulticurrencyRevaluationReportHandlerModel, criteria, options)
}

// FindAccountMulticurrencyRevaluationReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMulticurrencyRevaluationReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
