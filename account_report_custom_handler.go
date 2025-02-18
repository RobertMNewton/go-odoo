package odoo

// AccountReportCustomHandler represents account.report.custom.handler model.
type AccountReportCustomHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountReportCustomHandlers represents array of account.report.custom.handler model.
type AccountReportCustomHandlers []AccountReportCustomHandler

// AccountReportCustomHandlerModel is the odoo model name.
const AccountReportCustomHandlerModel = "account.report.custom.handler"

// Many2One convert AccountReportCustomHandler to *Many2One.
func (arch *AccountReportCustomHandler) Many2One() *Many2One {
	return NewMany2One(arch.Id.Get(), "")
}

// CreateAccountReportCustomHandler creates a new account.report.custom.handler model and returns its id.
func (c *Client) CreateAccountReportCustomHandler(arch *AccountReportCustomHandler) (int64, error) {
	ids, err := c.CreateAccountReportCustomHandlers([]*AccountReportCustomHandler{arch})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportCustomHandler creates a new account.report.custom.handler model and returns its id.
func (c *Client) CreateAccountReportCustomHandlers(archs []*AccountReportCustomHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range archs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportCustomHandlerModel, vv, nil)
}

// UpdateAccountReportCustomHandler updates an existing account.report.custom.handler record.
func (c *Client) UpdateAccountReportCustomHandler(arch *AccountReportCustomHandler) error {
	return c.UpdateAccountReportCustomHandlers([]int64{arch.Id.Get()}, arch)
}

// UpdateAccountReportCustomHandlers updates existing account.report.custom.handler records.
// All records (represented by ids) will be updated by arch values.
func (c *Client) UpdateAccountReportCustomHandlers(ids []int64, arch *AccountReportCustomHandler) error {
	return c.Update(AccountReportCustomHandlerModel, ids, arch, nil)
}

// DeleteAccountReportCustomHandler deletes an existing account.report.custom.handler record.
func (c *Client) DeleteAccountReportCustomHandler(id int64) error {
	return c.DeleteAccountReportCustomHandlers([]int64{id})
}

// DeleteAccountReportCustomHandlers deletes existing account.report.custom.handler records.
func (c *Client) DeleteAccountReportCustomHandlers(ids []int64) error {
	return c.Delete(AccountReportCustomHandlerModel, ids)
}

// GetAccountReportCustomHandler gets account.report.custom.handler existing record.
func (c *Client) GetAccountReportCustomHandler(id int64) (*AccountReportCustomHandler, error) {
	archs, err := c.GetAccountReportCustomHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*archs)[0]), nil
}

// GetAccountReportCustomHandlers gets account.report.custom.handler existing records.
func (c *Client) GetAccountReportCustomHandlers(ids []int64) (*AccountReportCustomHandlers, error) {
	archs := &AccountReportCustomHandlers{}
	if err := c.Read(AccountReportCustomHandlerModel, ids, nil, archs); err != nil {
		return nil, err
	}
	return archs, nil
}

// FindAccountReportCustomHandler finds account.report.custom.handler record by querying it with criteria.
func (c *Client) FindAccountReportCustomHandler(criteria *Criteria) (*AccountReportCustomHandler, error) {
	archs := &AccountReportCustomHandlers{}
	if err := c.SearchRead(AccountReportCustomHandlerModel, criteria, NewOptions().Limit(1), archs); err != nil {
		return nil, err
	}
	return &((*archs)[0]), nil
}

// FindAccountReportCustomHandlers finds account.report.custom.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportCustomHandlers(criteria *Criteria, options *Options) (*AccountReportCustomHandlers, error) {
	archs := &AccountReportCustomHandlers{}
	if err := c.SearchRead(AccountReportCustomHandlerModel, criteria, options, archs); err != nil {
		return nil, err
	}
	return archs, nil
}

// FindAccountReportCustomHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportCustomHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportCustomHandlerModel, criteria, options)
}

// FindAccountReportCustomHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountReportCustomHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportCustomHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
