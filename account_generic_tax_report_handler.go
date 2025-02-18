package odoo

// AccountGenericTaxReportHandler represents account.generic.tax.report.handler model.
type AccountGenericTaxReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountGenericTaxReportHandlers represents array of account.generic.tax.report.handler model.
type AccountGenericTaxReportHandlers []AccountGenericTaxReportHandler

// AccountGenericTaxReportHandlerModel is the odoo model name.
const AccountGenericTaxReportHandlerModel = "account.generic.tax.report.handler"

// Many2One convert AccountGenericTaxReportHandler to *Many2One.
func (agtrh *AccountGenericTaxReportHandler) Many2One() *Many2One {
	return NewMany2One(agtrh.Id.Get(), "")
}

// CreateAccountGenericTaxReportHandler creates a new account.generic.tax.report.handler model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandler(agtrh *AccountGenericTaxReportHandler) (int64, error) {
	ids, err := c.CreateAccountGenericTaxReportHandlers([]*AccountGenericTaxReportHandler{agtrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGenericTaxReportHandler creates a new account.generic.tax.report.handler model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandlers(agtrhs []*AccountGenericTaxReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range agtrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountGenericTaxReportHandlerModel, vv, nil)
}

// UpdateAccountGenericTaxReportHandler updates an existing account.generic.tax.report.handler record.
func (c *Client) UpdateAccountGenericTaxReportHandler(agtrh *AccountGenericTaxReportHandler) error {
	return c.UpdateAccountGenericTaxReportHandlers([]int64{agtrh.Id.Get()}, agtrh)
}

// UpdateAccountGenericTaxReportHandlers updates existing account.generic.tax.report.handler records.
// All records (represented by ids) will be updated by agtrh values.
func (c *Client) UpdateAccountGenericTaxReportHandlers(ids []int64, agtrh *AccountGenericTaxReportHandler) error {
	return c.Update(AccountGenericTaxReportHandlerModel, ids, agtrh, nil)
}

// DeleteAccountGenericTaxReportHandler deletes an existing account.generic.tax.report.handler record.
func (c *Client) DeleteAccountGenericTaxReportHandler(id int64) error {
	return c.DeleteAccountGenericTaxReportHandlers([]int64{id})
}

// DeleteAccountGenericTaxReportHandlers deletes existing account.generic.tax.report.handler records.
func (c *Client) DeleteAccountGenericTaxReportHandlers(ids []int64) error {
	return c.Delete(AccountGenericTaxReportHandlerModel, ids)
}

// GetAccountGenericTaxReportHandler gets account.generic.tax.report.handler existing record.
func (c *Client) GetAccountGenericTaxReportHandler(id int64) (*AccountGenericTaxReportHandler, error) {
	agtrhs, err := c.GetAccountGenericTaxReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*agtrhs)[0]), nil
}

// GetAccountGenericTaxReportHandlers gets account.generic.tax.report.handler existing records.
func (c *Client) GetAccountGenericTaxReportHandlers(ids []int64) (*AccountGenericTaxReportHandlers, error) {
	agtrhs := &AccountGenericTaxReportHandlers{}
	if err := c.Read(AccountGenericTaxReportHandlerModel, ids, nil, agtrhs); err != nil {
		return nil, err
	}
	return agtrhs, nil
}

// FindAccountGenericTaxReportHandler finds account.generic.tax.report.handler record by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandler(criteria *Criteria) (*AccountGenericTaxReportHandler, error) {
	agtrhs := &AccountGenericTaxReportHandlers{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerModel, criteria, NewOptions().Limit(1), agtrhs); err != nil {
		return nil, err
	}
	return &((*agtrhs)[0]), nil
}

// FindAccountGenericTaxReportHandlers finds account.generic.tax.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlers(criteria *Criteria, options *Options) (*AccountGenericTaxReportHandlers, error) {
	agtrhs := &AccountGenericTaxReportHandlers{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerModel, criteria, options, agtrhs); err != nil {
		return nil, err
	}
	return agtrhs, nil
}

// FindAccountGenericTaxReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGenericTaxReportHandlerModel, criteria, options)
}

// FindAccountGenericTaxReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGenericTaxReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
