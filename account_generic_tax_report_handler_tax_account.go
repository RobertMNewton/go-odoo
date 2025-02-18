package odoo

// AccountGenericTaxReportHandlerTaxAccount represents account.generic.tax.report.handler.tax.account model.
type AccountGenericTaxReportHandlerTaxAccount struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountGenericTaxReportHandlerTaxAccounts represents array of account.generic.tax.report.handler.tax.account model.
type AccountGenericTaxReportHandlerTaxAccounts []AccountGenericTaxReportHandlerTaxAccount

// AccountGenericTaxReportHandlerTaxAccountModel is the odoo model name.
const AccountGenericTaxReportHandlerTaxAccountModel = "account.generic.tax.report.handler.tax.account"

// Many2One convert AccountGenericTaxReportHandlerTaxAccount to *Many2One.
func (agtrhta *AccountGenericTaxReportHandlerTaxAccount) Many2One() *Many2One {
	return NewMany2One(agtrhta.Id.Get(), "")
}

// CreateAccountGenericTaxReportHandlerTaxAccount creates a new account.generic.tax.report.handler.tax.account model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandlerTaxAccount(agtrhta *AccountGenericTaxReportHandlerTaxAccount) (int64, error) {
	ids, err := c.CreateAccountGenericTaxReportHandlerTaxAccounts([]*AccountGenericTaxReportHandlerTaxAccount{agtrhta})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountGenericTaxReportHandlerTaxAccount creates a new account.generic.tax.report.handler.tax.account model and returns its id.
func (c *Client) CreateAccountGenericTaxReportHandlerTaxAccounts(agtrhtas []*AccountGenericTaxReportHandlerTaxAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range agtrhtas {
		vv = append(vv, v)
	}
	return c.Create(AccountGenericTaxReportHandlerTaxAccountModel, vv, nil)
}

// UpdateAccountGenericTaxReportHandlerTaxAccount updates an existing account.generic.tax.report.handler.tax.account record.
func (c *Client) UpdateAccountGenericTaxReportHandlerTaxAccount(agtrhta *AccountGenericTaxReportHandlerTaxAccount) error {
	return c.UpdateAccountGenericTaxReportHandlerTaxAccounts([]int64{agtrhta.Id.Get()}, agtrhta)
}

// UpdateAccountGenericTaxReportHandlerTaxAccounts updates existing account.generic.tax.report.handler.tax.account records.
// All records (represented by ids) will be updated by agtrhta values.
func (c *Client) UpdateAccountGenericTaxReportHandlerTaxAccounts(ids []int64, agtrhta *AccountGenericTaxReportHandlerTaxAccount) error {
	return c.Update(AccountGenericTaxReportHandlerTaxAccountModel, ids, agtrhta, nil)
}

// DeleteAccountGenericTaxReportHandlerTaxAccount deletes an existing account.generic.tax.report.handler.tax.account record.
func (c *Client) DeleteAccountGenericTaxReportHandlerTaxAccount(id int64) error {
	return c.DeleteAccountGenericTaxReportHandlerTaxAccounts([]int64{id})
}

// DeleteAccountGenericTaxReportHandlerTaxAccounts deletes existing account.generic.tax.report.handler.tax.account records.
func (c *Client) DeleteAccountGenericTaxReportHandlerTaxAccounts(ids []int64) error {
	return c.Delete(AccountGenericTaxReportHandlerTaxAccountModel, ids)
}

// GetAccountGenericTaxReportHandlerTaxAccount gets account.generic.tax.report.handler.tax.account existing record.
func (c *Client) GetAccountGenericTaxReportHandlerTaxAccount(id int64) (*AccountGenericTaxReportHandlerTaxAccount, error) {
	agtrhtas, err := c.GetAccountGenericTaxReportHandlerTaxAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*agtrhtas)[0]), nil
}

// GetAccountGenericTaxReportHandlerTaxAccounts gets account.generic.tax.report.handler.tax.account existing records.
func (c *Client) GetAccountGenericTaxReportHandlerTaxAccounts(ids []int64) (*AccountGenericTaxReportHandlerTaxAccounts, error) {
	agtrhtas := &AccountGenericTaxReportHandlerTaxAccounts{}
	if err := c.Read(AccountGenericTaxReportHandlerTaxAccountModel, ids, nil, agtrhtas); err != nil {
		return nil, err
	}
	return agtrhtas, nil
}

// FindAccountGenericTaxReportHandlerTaxAccount finds account.generic.tax.report.handler.tax.account record by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandlerTaxAccount(criteria *Criteria) (*AccountGenericTaxReportHandlerTaxAccount, error) {
	agtrhtas := &AccountGenericTaxReportHandlerTaxAccounts{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerTaxAccountModel, criteria, NewOptions().Limit(1), agtrhtas); err != nil {
		return nil, err
	}
	return &((*agtrhtas)[0]), nil
}

// FindAccountGenericTaxReportHandlerTaxAccounts finds account.generic.tax.report.handler.tax.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlerTaxAccounts(criteria *Criteria, options *Options) (*AccountGenericTaxReportHandlerTaxAccounts, error) {
	agtrhtas := &AccountGenericTaxReportHandlerTaxAccounts{}
	if err := c.SearchRead(AccountGenericTaxReportHandlerTaxAccountModel, criteria, options, agtrhtas); err != nil {
		return nil, err
	}
	return agtrhtas, nil
}

// FindAccountGenericTaxReportHandlerTaxAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportHandlerTaxAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountGenericTaxReportHandlerTaxAccountModel, criteria, options)
}

// FindAccountGenericTaxReportHandlerTaxAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportHandlerTaxAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGenericTaxReportHandlerTaxAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
