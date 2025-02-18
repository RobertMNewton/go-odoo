package odoo

// AccountBankSelection represents account.bank.selection model.
type AccountBankSelection struct {
	AccountOnlineAccountIds *Relation `xmlrpc:"account_online_account_ids,omitempty"`
	AccountOnlineLinkId     *Many2One `xmlrpc:"account_online_link_id,omitempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	InstitutionName         *String   `xmlrpc:"institution_name,omitempty"`
	SelectedAccount         *Many2One `xmlrpc:"selected_account,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountBankSelections represents array of account.bank.selection model.
type AccountBankSelections []AccountBankSelection

// AccountBankSelectionModel is the odoo model name.
const AccountBankSelectionModel = "account.bank.selection"

// Many2One convert AccountBankSelection to *Many2One.
func (abs *AccountBankSelection) Many2One() *Many2One {
	return NewMany2One(abs.Id.Get(), "")
}

// CreateAccountBankSelection creates a new account.bank.selection model and returns its id.
func (c *Client) CreateAccountBankSelection(abs *AccountBankSelection) (int64, error) {
	ids, err := c.CreateAccountBankSelections([]*AccountBankSelection{abs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankSelection creates a new account.bank.selection model and returns its id.
func (c *Client) CreateAccountBankSelections(abss []*AccountBankSelection) ([]int64, error) {
	var vv []interface{}
	for _, v := range abss {
		vv = append(vv, v)
	}
	return c.Create(AccountBankSelectionModel, vv, nil)
}

// UpdateAccountBankSelection updates an existing account.bank.selection record.
func (c *Client) UpdateAccountBankSelection(abs *AccountBankSelection) error {
	return c.UpdateAccountBankSelections([]int64{abs.Id.Get()}, abs)
}

// UpdateAccountBankSelections updates existing account.bank.selection records.
// All records (represented by ids) will be updated by abs values.
func (c *Client) UpdateAccountBankSelections(ids []int64, abs *AccountBankSelection) error {
	return c.Update(AccountBankSelectionModel, ids, abs, nil)
}

// DeleteAccountBankSelection deletes an existing account.bank.selection record.
func (c *Client) DeleteAccountBankSelection(id int64) error {
	return c.DeleteAccountBankSelections([]int64{id})
}

// DeleteAccountBankSelections deletes existing account.bank.selection records.
func (c *Client) DeleteAccountBankSelections(ids []int64) error {
	return c.Delete(AccountBankSelectionModel, ids)
}

// GetAccountBankSelection gets account.bank.selection existing record.
func (c *Client) GetAccountBankSelection(id int64) (*AccountBankSelection, error) {
	abss, err := c.GetAccountBankSelections([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abss)[0]), nil
}

// GetAccountBankSelections gets account.bank.selection existing records.
func (c *Client) GetAccountBankSelections(ids []int64) (*AccountBankSelections, error) {
	abss := &AccountBankSelections{}
	if err := c.Read(AccountBankSelectionModel, ids, nil, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankSelection finds account.bank.selection record by querying it with criteria.
func (c *Client) FindAccountBankSelection(criteria *Criteria) (*AccountBankSelection, error) {
	abss := &AccountBankSelections{}
	if err := c.SearchRead(AccountBankSelectionModel, criteria, NewOptions().Limit(1), abss); err != nil {
		return nil, err
	}
	return &((*abss)[0]), nil
}

// FindAccountBankSelections finds account.bank.selection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankSelections(criteria *Criteria, options *Options) (*AccountBankSelections, error) {
	abss := &AccountBankSelections{}
	if err := c.SearchRead(AccountBankSelectionModel, criteria, options, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankSelectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankSelectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankSelectionModel, criteria, options)
}

// FindAccountBankSelectionId finds record id by querying it with criteria.
func (c *Client) FindAccountBankSelectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankSelectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
