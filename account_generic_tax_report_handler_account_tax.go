package odoo

// AccountGenericTaxReportHandlerAccountTax represents account.generic.tax.report.handler.account.tax model.
type AccountGenericTaxReportHandlerAccountTax struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountGenericTaxReportHandlerAccountTaxs represents array of account.generic.tax.report.handler.account.tax model.
type AccountGenericTaxReportHandlerAccountTaxs []AccountGenericTaxReportHandlerAccountTax

// AccountGenericTaxReportHandlerAccountTaxModel is the odoo model name.
const AccountGenericTaxReportHandlerAccountTaxModel = "account.generic.tax.report.handler.account.tax"

// Many2One convert AccountGenericTaxReportHandlerAccountTax to *Many2One.
func (agtrhat *AccountGenericTaxReportHandlerAccountTax) Many2One() *Many2One {
	return NewMany2One(agtrhat.Id.Get(), "")
}

// CreateAccountGenericTaxReportHandlerAccountTax creates a new account.generic.tax.report.handler.account.tax model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandlerAccountTax(agtrhat *AccountGenericTaxReportHandlerAccountTax) (int64, error) {
	ids, err := c.CreateAccountGenericTaxReportHandlerAccountTaxs([]*AccountGenericTaxReportHandlerAccountTax{agtrhat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGenericTaxReportHandlerAccountTax creates a new account.generic.tax.report.handler.account.tax model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandlerAccountTaxs(agtrhats []*AccountGenericTaxReportHandlerAccountTax) ([]int64, error) {
	var vv []interface{}
	for _, v := range agtrhats {
		vv = append(vv, v)
	}
	return c.Create(AccountGenericTaxReportHandlerAccountTaxModel, vv, nil)
}

// UpdateAccountGenericTaxReportHandlerAccountTax updates an existing account.generic.tax.report.handler.account.tax record.
func (c *Client) UpdateAccountGenericTaxReportHandlerAccountTax(agtrhat *AccountGenericTaxReportHandlerAccountTax) error {
	return c.UpdateAccountGenericTaxReportHandlerAccountTaxs([]int64{agtrhat.Id.Get()}, agtrhat)
}

// UpdateAccountGenericTaxReportHandlerAccountTaxs updates existing account.generic.tax.report.handler.account.tax records.
// All records (represented by ids) will be updated by agtrhat values.
func (c *Client) UpdateAccountGenericTaxReportHandlerAccountTaxs(ids []int64, agtrhat *AccountGenericTaxReportHandlerAccountTax) error {
	return c.Update(AccountGenericTaxReportHandlerAccountTaxModel, ids, agtrhat, nil)
}

// DeleteAccountGenericTaxReportHandlerAccountTax deletes an existing account.generic.tax.report.handler.account.tax record.
func (c *Client) DeleteAccountGenericTaxReportHandlerAccountTax(id int64) error {
	return c.DeleteAccountGenericTaxReportHandlerAccountTaxs([]int64{id})
}

// DeleteAccountGenericTaxReportHandlerAccountTaxs deletes existing account.generic.tax.report.handler.account.tax records.
func (c *Client) DeleteAccountGenericTaxReportHandlerAccountTaxs(ids []int64) error {
	return c.Delete(AccountGenericTaxReportHandlerAccountTaxModel, ids)
}

// GetAccountGenericTaxReportHandlerAccountTax gets account.generic.tax.report.handler.account.tax existing record.
func (c *Client) GetAccountGenericTaxReportHandlerAccountTax(id int64) (*AccountGenericTaxReportHandlerAccountTax, error) {
	agtrhats, err := c.GetAccountGenericTaxReportHandlerAccountTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*agtrhats)[0]), nil
}

// GetAccountGenericTaxReportHandlerAccountTaxs gets account.generic.tax.report.handler.account.tax existing records.
func (c *Client) GetAccountGenericTaxReportHandlerAccountTaxs(ids []int64) (*AccountGenericTaxReportHandlerAccountTaxs, error) {
	agtrhats := &AccountGenericTaxReportHandlerAccountTaxs{}
	if err := c.Read(AccountGenericTaxReportHandlerAccountTaxModel, ids, nil, agtrhats); err != nil {
		return nil, err
	}
	return agtrhats, nil
}

// FindAccountGenericTaxReportHandlerAccountTax finds account.generic.tax.report.handler.account.tax record by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandlerAccountTax(criteria *Criteria) (*AccountGenericTaxReportHandlerAccountTax, error) {
	agtrhats := &AccountGenericTaxReportHandlerAccountTaxs{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerAccountTaxModel, criteria, NewOptions().Limit(1), agtrhats); err != nil {
		return nil, err
	}
	return &((*agtrhats)[0]), nil
}

// FindAccountGenericTaxReportHandlerAccountTaxs finds account.generic.tax.report.handler.account.tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlerAccountTaxs(criteria *Criteria, options *Options) (*AccountGenericTaxReportHandlerAccountTaxs, error) {
	agtrhats := &AccountGenericTaxReportHandlerAccountTaxs{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerAccountTaxModel, criteria, options, agtrhats); err != nil {
		return nil, err
	}
	return agtrhats, nil
}

// FindAccountGenericTaxReportHandlerAccountTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlerAccountTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGenericTaxReportHandlerAccountTaxModel, criteria, options)
}

// FindAccountGenericTaxReportHandlerAccountTaxId finds record id by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandlerAccountTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGenericTaxReportHandlerAccountTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
