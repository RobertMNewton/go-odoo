package odoo

// AccountReportBudgetItem represents account.report.budget.item model.
type AccountReportBudgetItem struct {
	AccountId   *Many2One `xmlrpc:"account_id,omitempty"`
	Amount      *Float    `xmlrpc:"amount,omitempty"`
	BudgetId    *Many2One `xmlrpc:"budget_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReportBudgetItems represents array of account.report.budget.item model.
type AccountReportBudgetItems []AccountReportBudgetItem

// AccountReportBudgetItemModel is the odoo model name.
const AccountReportBudgetItemModel = "account.report.budget.item"

// Many2One convert AccountReportBudgetItem to *Many2One.
func (arbi *AccountReportBudgetItem) Many2One() *Many2One {
	return NewMany2One(arbi.Id.Get(), "")
}

// CreateAccountReportBudgetItem creates a new account.report.budget.item model and returns its id.
func (c *Client) CreateAccountReportBudgetItem(arbi *AccountReportBudgetItem) (int64, error) {
	ids, err := c.CreateAccountReportBudgetItems([]*AccountReportBudgetItem{arbi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportBudgetItem creates a new account.report.budget.item model and returns its id.
func (c *Client) CreateAccountReportBudgetItems(arbis []*AccountReportBudgetItem) ([]int64, error) {
	var vv []interface{}
	for _, v := range arbis {
		vv = append(vv, v)
	}
	return c.Create(AccountReportBudgetItemModel, vv, nil)
}

// UpdateAccountReportBudgetItem updates an existing account.report.budget.item record.
func (c *Client) UpdateAccountReportBudgetItem(arbi *AccountReportBudgetItem) error {
	return c.UpdateAccountReportBudgetItems([]int64{arbi.Id.Get()}, arbi)
}

// UpdateAccountReportBudgetItems updates existing account.report.budget.item records.
// All records (represented by ids) will be updated by arbi values.
func (c *Client) UpdateAccountReportBudgetItems(ids []int64, arbi *AccountReportBudgetItem) error {
	return c.Update(AccountReportBudgetItemModel, ids, arbi, nil)
}

// DeleteAccountReportBudgetItem deletes an existing account.report.budget.item record.
func (c *Client) DeleteAccountReportBudgetItem(id int64) error {
	return c.DeleteAccountReportBudgetItems([]int64{id})
}

// DeleteAccountReportBudgetItems deletes existing account.report.budget.item records.
func (c *Client) DeleteAccountReportBudgetItems(ids []int64) error {
	return c.Delete(AccountReportBudgetItemModel, ids)
}

// GetAccountReportBudgetItem gets account.report.budget.item existing record.
func (c *Client) GetAccountReportBudgetItem(id int64) (*AccountReportBudgetItem, error) {
	arbis, err := c.GetAccountReportBudgetItems([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arbis)[0]), nil
}

// GetAccountReportBudgetItems gets account.report.budget.item existing records.
func (c *Client) GetAccountReportBudgetItems(ids []int64) (*AccountReportBudgetItems, error) {
	arbis := &AccountReportBudgetItems{}
	if err := c.Read(AccountReportBudgetItemModel, ids, nil, arbis); err != nil {
		return nil, err
	}
	return arbis, nil
}

// FindAccountReportBudgetItem finds account.report.budget.item record by querying it with criteria.
func (c *Client) FindAccountReportBudgetItem(criteria *Criteria) (*AccountReportBudgetItem, error) {
	arbis := &AccountReportBudgetItems{}
	if err := c.SearchRead(AccountReportBudgetItemModel, criteria, NewOptions().Limit(1), arbis); err != nil {
		return nil, err
	}
	return &((*arbis)[0]), nil
}

// FindAccountReportBudgetItems finds account.report.budget.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportBudgetItems(criteria *Criteria, options *Options) (*AccountReportBudgetItems, error) {
	arbis := &AccountReportBudgetItems{}
	if err := c.SearchRead(AccountReportBudgetItemModel, criteria, options, arbis); err != nil {
		return nil, err
	}
	return arbis, nil
}

// FindAccountReportBudgetItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportBudgetItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportBudgetItemModel, criteria, options)
}

// FindAccountReportBudgetItemId finds record id by querying it with criteria.
func (c *Client) FindAccountReportBudgetItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportBudgetItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
