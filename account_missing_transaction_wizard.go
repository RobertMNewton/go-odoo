package odoo

// AccountMissingTransactionWizard represents account.missing.transaction.wizard model.
type AccountMissingTransactionWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountMissingTransactionWizards represents array of account.missing.transaction.wizard model.
type AccountMissingTransactionWizards []AccountMissingTransactionWizard

// AccountMissingTransactionWizardModel is the odoo model name.
const AccountMissingTransactionWizardModel = "account.missing.transaction.wizard"

// Many2One convert AccountMissingTransactionWizard to *Many2One.
func (amtw *AccountMissingTransactionWizard) Many2One() *Many2One {
	return NewMany2One(amtw.Id.Get(), "")
}

// CreateAccountMissingTransactionWizard creates a new account.missing.transaction.wizard model and returns its id.
func (c *Client) CreateAccountMissingTransactionWizard(amtw *AccountMissingTransactionWizard) (int64, error) {
	ids, err := c.CreateAccountMissingTransactionWizards([]*AccountMissingTransactionWizard{amtw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMissingTransactionWizard creates a new account.missing.transaction.wizard model and returns its id.
func (c *Client) CreateAccountMissingTransactionWizards(amtws []*AccountMissingTransactionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amtws {
		vv = append(vv, v)
	}
	return c.Create(AccountMissingTransactionWizardModel, vv, nil)
}

// UpdateAccountMissingTransactionWizard updates an existing account.missing.transaction.wizard record.
func (c *Client) UpdateAccountMissingTransactionWizard(amtw *AccountMissingTransactionWizard) error {
	return c.UpdateAccountMissingTransactionWizards([]int64{amtw.Id.Get()}, amtw)
}

// UpdateAccountMissingTransactionWizards updates existing account.missing.transaction.wizard records.
// All records (represented by ids) will be updated by amtw values.
func (c *Client) UpdateAccountMissingTransactionWizards(ids []int64, amtw *AccountMissingTransactionWizard) error {
	return c.Update(AccountMissingTransactionWizardModel, ids, amtw, nil)
}

// DeleteAccountMissingTransactionWizard deletes an existing account.missing.transaction.wizard record.
func (c *Client) DeleteAccountMissingTransactionWizard(id int64) error {
	return c.DeleteAccountMissingTransactionWizards([]int64{id})
}

// DeleteAccountMissingTransactionWizards deletes existing account.missing.transaction.wizard records.
func (c *Client) DeleteAccountMissingTransactionWizards(ids []int64) error {
	return c.Delete(AccountMissingTransactionWizardModel, ids)
}

// GetAccountMissingTransactionWizard gets account.missing.transaction.wizard existing record.
func (c *Client) GetAccountMissingTransactionWizard(id int64) (*AccountMissingTransactionWizard, error) {
	amtws, err := c.GetAccountMissingTransactionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amtws)[0]), nil
}

// GetAccountMissingTransactionWizards gets account.missing.transaction.wizard existing records.
func (c *Client) GetAccountMissingTransactionWizards(ids []int64) (*AccountMissingTransactionWizards, error) {
	amtws := &AccountMissingTransactionWizards{}
	if err := c.Read(AccountMissingTransactionWizardModel, ids, nil, amtws); err != nil {
		return nil, err
	}
	return amtws, nil
}

// FindAccountMissingTransactionWizard finds account.missing.transaction.wizard record by querying it with criteria.
func (c *Client) FindAccountMissingTransactionWizard(criteria *Criteria) (*AccountMissingTransactionWizard, error) {
	amtws := &AccountMissingTransactionWizards{}
	if err := c.SearchRead(AccountMissingTransactionWizardModel, criteria, NewOptions().Limit(1), amtws); err != nil {
		return nil, err
	}
	return &((*amtws)[0]), nil
}

// FindAccountMissingTransactionWizards finds account.missing.transaction.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMissingTransactionWizards(criteria *Criteria, options *Options) (*AccountMissingTransactionWizards, error) {
	amtws := &AccountMissingTransactionWizards{}
	if err := c.SearchRead(AccountMissingTransactionWizardModel, criteria, options, amtws); err != nil {
		return nil, err
	}
	return amtws, nil
}

// FindAccountMissingTransactionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMissingTransactionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMissingTransactionWizardModel, criteria, options)
}

// FindAccountMissingTransactionWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMissingTransactionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMissingTransactionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
