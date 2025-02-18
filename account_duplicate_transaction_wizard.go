package odoo

// AccountDuplicateTransactionWizard represents account.duplicate.transaction.wizard model.
type AccountDuplicateTransactionWizard struct {
	CreateDate      *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One   `xmlrpc:"create_uid,omitempty"`
	Date            *Time       `xmlrpc:"date,omitempty"`
	DisplayName     *String     `xmlrpc:"display_name,omitempty"`
	FirstIdsInGroup interface{} `xmlrpc:"first_ids_in_group,omitempty"`
	Id              *Int        `xmlrpc:"id,omitempty"`
	JournalId       *Many2One   `xmlrpc:"journal_id,omitempty"`
	TransactionIds  *Relation   `xmlrpc:"transaction_ids,omitempty"`
	WriteDate       *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountDuplicateTransactionWizards represents array of account.duplicate.transaction.wizard model.
type AccountDuplicateTransactionWizards []AccountDuplicateTransactionWizard

// AccountDuplicateTransactionWizardModel is the odoo model name.
const AccountDuplicateTransactionWizardModel = "account.duplicate.transaction.wizard"

// Many2One convert AccountDuplicateTransactionWizard to *Many2One.
func (adtw *AccountDuplicateTransactionWizard) Many2One() *Many2One {
	return NewMany2One(adtw.Id.Get(), "")
}

// CreateAccountDuplicateTransactionWizard creates a new account.duplicate.transaction.wizard model and returns its id.
func (c *Client) CreateAccountDuplicateTransactionWizard(adtw *AccountDuplicateTransactionWizard) (int64, error) {
	ids, err := c.CreateAccountDuplicateTransactionWizards([]*AccountDuplicateTransactionWizard{adtw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDuplicateTransactionWizard creates a new account.duplicate.transaction.wizard model and returns its id.
func (c *Client) CreateAccountDuplicateTransactionWizards(adtws []*AccountDuplicateTransactionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range adtws {
		vv = append(vv, v)
	}
	return c.Create(AccountDuplicateTransactionWizardModel, vv, nil)
}

// UpdateAccountDuplicateTransactionWizard updates an existing account.duplicate.transaction.wizard record.
func (c *Client) UpdateAccountDuplicateTransactionWizard(adtw *AccountDuplicateTransactionWizard) error {
	return c.UpdateAccountDuplicateTransactionWizards([]int64{adtw.Id.Get()}, adtw)
}

// UpdateAccountDuplicateTransactionWizards updates existing account.duplicate.transaction.wizard records.
// All records (represented by ids) will be updated by adtw values.
func (c *Client) UpdateAccountDuplicateTransactionWizards(ids []int64, adtw *AccountDuplicateTransactionWizard) error {
	return c.Update(AccountDuplicateTransactionWizardModel, ids, adtw, nil)
}

// DeleteAccountDuplicateTransactionWizard deletes an existing account.duplicate.transaction.wizard record.
func (c *Client) DeleteAccountDuplicateTransactionWizard(id int64) error {
	return c.DeleteAccountDuplicateTransactionWizards([]int64{id})
}

// DeleteAccountDuplicateTransactionWizards deletes existing account.duplicate.transaction.wizard records.
func (c *Client) DeleteAccountDuplicateTransactionWizards(ids []int64) error {
	return c.Delete(AccountDuplicateTransactionWizardModel, ids)
}

// GetAccountDuplicateTransactionWizard gets account.duplicate.transaction.wizard existing record.
func (c *Client) GetAccountDuplicateTransactionWizard(id int64) (*AccountDuplicateTransactionWizard, error) {
	adtws, err := c.GetAccountDuplicateTransactionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*adtws)[0]), nil
}

// GetAccountDuplicateTransactionWizards gets account.duplicate.transaction.wizard existing records.
func (c *Client) GetAccountDuplicateTransactionWizards(ids []int64) (*AccountDuplicateTransactionWizards, error) {
	adtws := &AccountDuplicateTransactionWizards{}
	if err := c.Read(AccountDuplicateTransactionWizardModel, ids, nil, adtws); err != nil {
		return nil, err
	}
	return adtws, nil
}

// FindAccountDuplicateTransactionWizard finds account.duplicate.transaction.wizard record by querying it with criteria.
func (c *Client) FindAccountDuplicateTransactionWizard(criteria *Criteria) (*AccountDuplicateTransactionWizard, error) {
	adtws := &AccountDuplicateTransactionWizards{}
	if err := c.SearchRead(AccountDuplicateTransactionWizardModel, criteria, NewOptions().Limit(1), adtws); err != nil {
		return nil, err
	}
	return &((*adtws)[0]), nil
}

// FindAccountDuplicateTransactionWizards finds account.duplicate.transaction.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDuplicateTransactionWizards(criteria *Criteria, options *Options) (*AccountDuplicateTransactionWizards, error) {
	adtws := &AccountDuplicateTransactionWizards{}
	if err := c.SearchRead(AccountDuplicateTransactionWizardModel, criteria, options, adtws); err != nil {
		return nil, err
	}
	return adtws, nil
}

// FindAccountDuplicateTransactionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDuplicateTransactionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDuplicateTransactionWizardModel, criteria, options)
}

// FindAccountDuplicateTransactionWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountDuplicateTransactionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDuplicateTransactionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
