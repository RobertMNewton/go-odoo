package odoo

// AccountDisallowedExpensesRate represents account.disallowed.expenses.rate model.
type AccountDisallowedExpensesRate struct {
	CategoryId  *Many2One `xmlrpc:"category_id,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DateFrom    *Time     `xmlrpc:"date_from,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Rate        *Float    `xmlrpc:"rate,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountDisallowedExpensesRates represents array of account.disallowed.expenses.rate model.
type AccountDisallowedExpensesRates []AccountDisallowedExpensesRate

// AccountDisallowedExpensesRateModel is the odoo model name.
const AccountDisallowedExpensesRateModel = "account.disallowed.expenses.rate"

// Many2One convert AccountDisallowedExpensesRate to *Many2One.
func (ader *AccountDisallowedExpensesRate) Many2One() *Many2One {
	return NewMany2One(ader.Id.Get(), "")
}

// CreateAccountDisallowedExpensesRate creates a new account.disallowed.expenses.rate model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesRate(ader *AccountDisallowedExpensesRate) (int64, error) {
	ids, err := c.CreateAccountDisallowedExpensesRates([]*AccountDisallowedExpensesRate{ader})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDisallowedExpensesRate creates a new account.disallowed.expenses.rate model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesRates(aders []*AccountDisallowedExpensesRate) ([]int64, error) {
	var vv []interface{}
	for _, v := range aders {
		vv = append(vv, v)
	}
	return c.Create(AccountDisallowedExpensesRateModel, vv, nil)
}

// UpdateAccountDisallowedExpensesRate updates an existing account.disallowed.expenses.rate record.
func (c *Client) UpdateAccountDisallowedExpensesRate(ader *AccountDisallowedExpensesRate) error {
	return c.UpdateAccountDisallowedExpensesRates([]int64{ader.Id.Get()}, ader)
}

// UpdateAccountDisallowedExpensesRates updates existing account.disallowed.expenses.rate records.
// All records (represented by ids) will be updated by ader values.
func (c *Client) UpdateAccountDisallowedExpensesRates(ids []int64, ader *AccountDisallowedExpensesRate) error {
	return c.Update(AccountDisallowedExpensesRateModel, ids, ader, nil)
}

// DeleteAccountDisallowedExpensesRate deletes an existing account.disallowed.expenses.rate record.
func (c *Client) DeleteAccountDisallowedExpensesRate(id int64) error {
	return c.DeleteAccountDisallowedExpensesRates([]int64{id})
}

// DeleteAccountDisallowedExpensesRates deletes existing account.disallowed.expenses.rate records.
func (c *Client) DeleteAccountDisallowedExpensesRates(ids []int64) error {
	return c.Delete(AccountDisallowedExpensesRateModel, ids)
}

// GetAccountDisallowedExpensesRate gets account.disallowed.expenses.rate existing record.
func (c *Client) GetAccountDisallowedExpensesRate(id int64) (*AccountDisallowedExpensesRate, error) {
	aders, err := c.GetAccountDisallowedExpensesRates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aders)[0]), nil
}

// GetAccountDisallowedExpensesRates gets account.disallowed.expenses.rate existing records.
func (c *Client) GetAccountDisallowedExpensesRates(ids []int64) (*AccountDisallowedExpensesRates, error) {
	aders := &AccountDisallowedExpensesRates{}
	if err := c.Read(AccountDisallowedExpensesRateModel, ids, nil, aders); err != nil {
		return nil, err
	}
	return aders, nil
}

// FindAccountDisallowedExpensesRate finds account.disallowed.expenses.rate record by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesRate(criteria *Criteria) (*AccountDisallowedExpensesRate, error) {
	aders := &AccountDisallowedExpensesRates{}
	if err := c.SearchRead(AccountDisallowedExpensesRateModel, criteria, NewOptions().Limit(1), aders); err != nil {
		return nil, err
	}
	return &((*aders)[0]), nil
}

// FindAccountDisallowedExpensesRates finds account.disallowed.expenses.rate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesRates(criteria *Criteria, options *Options) (*AccountDisallowedExpensesRates, error) {
	aders := &AccountDisallowedExpensesRates{}
	if err := c.SearchRead(AccountDisallowedExpensesRateModel, criteria, options, aders); err != nil {
		return nil, err
	}
	return aders, nil
}

// FindAccountDisallowedExpensesRateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesRateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDisallowedExpensesRateModel, criteria, options)
}

// FindAccountDisallowedExpensesRateId finds record id by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesRateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDisallowedExpensesRateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
