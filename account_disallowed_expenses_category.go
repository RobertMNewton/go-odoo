package odoo

// AccountDisallowedExpensesCategory represents account.disallowed.expenses.category model.
type AccountDisallowedExpensesCategory struct {
	AccountIds  *Relation `xmlrpc:"account_ids,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentRate *String   `xmlrpc:"current_rate,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	RateIds     *Relation `xmlrpc:"rate_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountDisallowedExpensesCategorys represents array of account.disallowed.expenses.category model.
type AccountDisallowedExpensesCategorys []AccountDisallowedExpensesCategory

// AccountDisallowedExpensesCategoryModel is the odoo model name.
const AccountDisallowedExpensesCategoryModel = "account.disallowed.expenses.category"

// Many2One convert AccountDisallowedExpensesCategory to *Many2One.
func (adec *AccountDisallowedExpensesCategory) Many2One() *Many2One {
	return NewMany2One(adec.Id.Get(), "")
}

// CreateAccountDisallowedExpensesCategory creates a new account.disallowed.expenses.category model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesCategory(adec *AccountDisallowedExpensesCategory) (int64, error) {
	ids, err := c.CreateAccountDisallowedExpensesCategorys([]*AccountDisallowedExpensesCategory{adec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDisallowedExpensesCategory creates a new account.disallowed.expenses.category model and returns its id.
func (c *Client) CreateAccountDisallowedExpensesCategorys(adecs []*AccountDisallowedExpensesCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range adecs {
		vv = append(vv, v)
	}
	return c.Create(AccountDisallowedExpensesCategoryModel, vv, nil)
}

// UpdateAccountDisallowedExpensesCategory updates an existing account.disallowed.expenses.category record.
func (c *Client) UpdateAccountDisallowedExpensesCategory(adec *AccountDisallowedExpensesCategory) error {
	return c.UpdateAccountDisallowedExpensesCategorys([]int64{adec.Id.Get()}, adec)
}

// UpdateAccountDisallowedExpensesCategorys updates existing account.disallowed.expenses.category records.
// All records (represented by ids) will be updated by adec values.
func (c *Client) UpdateAccountDisallowedExpensesCategorys(ids []int64, adec *AccountDisallowedExpensesCategory) error {
	return c.Update(AccountDisallowedExpensesCategoryModel, ids, adec, nil)
}

// DeleteAccountDisallowedExpensesCategory deletes an existing account.disallowed.expenses.category record.
func (c *Client) DeleteAccountDisallowedExpensesCategory(id int64) error {
	return c.DeleteAccountDisallowedExpensesCategorys([]int64{id})
}

// DeleteAccountDisallowedExpensesCategorys deletes existing account.disallowed.expenses.category records.
func (c *Client) DeleteAccountDisallowedExpensesCategorys(ids []int64) error {
	return c.Delete(AccountDisallowedExpensesCategoryModel, ids)
}

// GetAccountDisallowedExpensesCategory gets account.disallowed.expenses.category existing record.
func (c *Client) GetAccountDisallowedExpensesCategory(id int64) (*AccountDisallowedExpensesCategory, error) {
	adecs, err := c.GetAccountDisallowedExpensesCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*adecs)[0]), nil
}

// GetAccountDisallowedExpensesCategorys gets account.disallowed.expenses.category existing records.
func (c *Client) GetAccountDisallowedExpensesCategorys(ids []int64) (*AccountDisallowedExpensesCategorys, error) {
	adecs := &AccountDisallowedExpensesCategorys{}
	if err := c.Read(AccountDisallowedExpensesCategoryModel, ids, nil, adecs); err != nil {
		return nil, err
	}
	return adecs, nil
}

// FindAccountDisallowedExpensesCategory finds account.disallowed.expenses.category record by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesCategory(criteria *Criteria) (*AccountDisallowedExpensesCategory, error) {
	adecs := &AccountDisallowedExpensesCategorys{}
	if err := c.SearchRead(AccountDisallowedExpensesCategoryModel, criteria, NewOptions().Limit(1), adecs); err != nil {
		return nil, err
	}
	return &((*adecs)[0]), nil
}

// FindAccountDisallowedExpensesCategorys finds account.disallowed.expenses.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesCategorys(criteria *Criteria, options *Options) (*AccountDisallowedExpensesCategorys, error) {
	adecs := &AccountDisallowedExpensesCategorys{}
	if err := c.SearchRead(AccountDisallowedExpensesCategoryModel, criteria, options, adecs); err != nil {
		return nil, err
	}
	return adecs, nil
}

// FindAccountDisallowedExpensesCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDisallowedExpensesCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDisallowedExpensesCategoryModel, criteria, options)
}

// FindAccountDisallowedExpensesCategoryId finds record id by querying it with criteria.
func (c *Client) FindAccountDisallowedExpensesCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDisallowedExpensesCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
