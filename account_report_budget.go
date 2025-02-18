package odoo

// AccountReportBudget represents account.report.budget model.
type AccountReportBudget struct {
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ItemIds     *Relation `xmlrpc:"item_ids,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReportBudgets represents array of account.report.budget model.
type AccountReportBudgets []AccountReportBudget

// AccountReportBudgetModel is the odoo model name.
const AccountReportBudgetModel = "account.report.budget"

// Many2One convert AccountReportBudget to *Many2One.
func (arb *AccountReportBudget) Many2One() *Many2One {
	return NewMany2One(arb.Id.Get(), "")
}

// CreateAccountReportBudget creates a new account.report.budget model and returns its id.
func (c *Client) CreateAccountReportBudget(arb *AccountReportBudget) (int64, error) {
	ids, err := c.CreateAccountReportBudgets([]*AccountReportBudget{arb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportBudget creates a new account.report.budget model and returns its id.
func (c *Client) CreateAccountReportBudgets(arbs []*AccountReportBudget) ([]int64, error) {
	var vv []interface{}
	for _, v := range arbs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportBudgetModel, vv, nil)
}

// UpdateAccountReportBudget updates an existing account.report.budget record.
func (c *Client) UpdateAccountReportBudget(arb *AccountReportBudget) error {
	return c.UpdateAccountReportBudgets([]int64{arb.Id.Get()}, arb)
}

// UpdateAccountReportBudgets updates existing account.report.budget records.
// All records (represented by ids) will be updated by arb values.
func (c *Client) UpdateAccountReportBudgets(ids []int64, arb *AccountReportBudget) error {
	return c.Update(AccountReportBudgetModel, ids, arb, nil)
}

// DeleteAccountReportBudget deletes an existing account.report.budget record.
func (c *Client) DeleteAccountReportBudget(id int64) error {
	return c.DeleteAccountReportBudgets([]int64{id})
}

// DeleteAccountReportBudgets deletes existing account.report.budget records.
func (c *Client) DeleteAccountReportBudgets(ids []int64) error {
	return c.Delete(AccountReportBudgetModel, ids)
}

// GetAccountReportBudget gets account.report.budget existing record.
func (c *Client) GetAccountReportBudget(id int64) (*AccountReportBudget, error) {
	arbs, err := c.GetAccountReportBudgets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arbs)[0]), nil
}

// GetAccountReportBudgets gets account.report.budget existing records.
func (c *Client) GetAccountReportBudgets(ids []int64) (*AccountReportBudgets, error) {
	arbs := &AccountReportBudgets{}
	if err := c.Read(AccountReportBudgetModel, ids, nil, arbs); err != nil {
		return nil, err
	}
	return arbs, nil
}

// FindAccountReportBudget finds account.report.budget record by querying it with criteria.
func (c *Client) FindAccountReportBudget(criteria *Criteria) (*AccountReportBudget, error) {
	arbs := &AccountReportBudgets{}
	if err := c.SearchRead(AccountReportBudgetModel, criteria, NewOptions().Limit(1), arbs); err != nil {
		return nil, err
	}
	return &((*arbs)[0]), nil
}

// FindAccountReportBudgets finds account.report.budget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportBudgets(criteria *Criteria, options *Options) (*AccountReportBudgets, error) {
	arbs := &AccountReportBudgets{}
	if err := c.SearchRead(AccountReportBudgetModel, criteria, options, arbs); err != nil {
		return nil, err
	}
	return arbs, nil
}

// FindAccountReportBudgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportBudgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportBudgetModel, criteria, options)
}

// FindAccountReportBudgetId finds record id by querying it with criteria.
func (c *Client) FindAccountReportBudgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportBudgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
