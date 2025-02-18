package odoo

// AccountBankStatementLineTransient represents account.bank.statement.line.transient model.
type AccountBankStatementLineTransient struct {
	AccountNumber               *String    `xmlrpc:"account_number,omitempty"`
	Amount                      *Float     `xmlrpc:"amount,omitempty"`
	AmountCurrency              *Float     `xmlrpc:"amount_currency,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	ForeignCurrencyId           *Many2One  `xmlrpc:"foreign_currency_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omitempty"`
	OnlineAccountId             *Many2One  `xmlrpc:"online_account_id,omitempty"`
	OnlineTransactionIdentifier *String    `xmlrpc:"online_transaction_identifier,omitempty"`
	PartnerName                 *String    `xmlrpc:"partner_name,omitempty"`
	PaymentRef                  *String    `xmlrpc:"payment_ref,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TransactionDetails          *String    `xmlrpc:"transaction_details,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountBankStatementLineTransients represents array of account.bank.statement.line.transient model.
type AccountBankStatementLineTransients []AccountBankStatementLineTransient

// AccountBankStatementLineTransientModel is the odoo model name.
const AccountBankStatementLineTransientModel = "account.bank.statement.line.transient"

// Many2One convert AccountBankStatementLineTransient to *Many2One.
func (abslt *AccountBankStatementLineTransient) Many2One() *Many2One {
	return NewMany2One(abslt.Id.Get(), "")
}

// CreateAccountBankStatementLineTransient creates a new account.bank.statement.line.transient model and returns its id.
func (c *Client) CreateAccountBankStatementLineTransient(abslt *AccountBankStatementLineTransient) (int64, error) {
	ids, err := c.CreateAccountBankStatementLineTransients([]*AccountBankStatementLineTransient{abslt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBankStatementLineTransient creates a new account.bank.statement.line.transient model and returns its id.
func (c *Client) CreateAccountBankStatementLineTransients(abslts []*AccountBankStatementLineTransient) ([]int64, error) {
	var vv []interface{}
	for _, v := range abslts {
		vv = append(vv, v)
	}
	return c.Create(AccountBankStatementLineTransientModel, vv, nil)
}

// UpdateAccountBankStatementLineTransient updates an existing account.bank.statement.line.transient record.
func (c *Client) UpdateAccountBankStatementLineTransient(abslt *AccountBankStatementLineTransient) error {
	return c.UpdateAccountBankStatementLineTransients([]int64{abslt.Id.Get()}, abslt)
}

// UpdateAccountBankStatementLineTransients updates existing account.bank.statement.line.transient records.
// All records (represented by ids) will be updated by abslt values.
func (c *Client) UpdateAccountBankStatementLineTransients(ids []int64, abslt *AccountBankStatementLineTransient) error {
	return c.Update(AccountBankStatementLineTransientModel, ids, abslt, nil)
}

// DeleteAccountBankStatementLineTransient deletes an existing account.bank.statement.line.transient record.
func (c *Client) DeleteAccountBankStatementLineTransient(id int64) error {
	return c.DeleteAccountBankStatementLineTransients([]int64{id})
}

// DeleteAccountBankStatementLineTransients deletes existing account.bank.statement.line.transient records.
func (c *Client) DeleteAccountBankStatementLineTransients(ids []int64) error {
	return c.Delete(AccountBankStatementLineTransientModel, ids)
}

// GetAccountBankStatementLineTransient gets account.bank.statement.line.transient existing record.
func (c *Client) GetAccountBankStatementLineTransient(id int64) (*AccountBankStatementLineTransient, error) {
	abslts, err := c.GetAccountBankStatementLineTransients([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abslts)[0]), nil
}

// GetAccountBankStatementLineTransients gets account.bank.statement.line.transient existing records.
func (c *Client) GetAccountBankStatementLineTransients(ids []int64) (*AccountBankStatementLineTransients, error) {
	abslts := &AccountBankStatementLineTransients{}
	if err := c.Read(AccountBankStatementLineTransientModel, ids, nil, abslts); err != nil {
		return nil, err
	}
	return abslts, nil
}

// FindAccountBankStatementLineTransient finds account.bank.statement.line.transient record by querying it with criteria.
func (c *Client) FindAccountBankStatementLineTransient(criteria *Criteria) (*AccountBankStatementLineTransient, error) {
	abslts := &AccountBankStatementLineTransients{}
	if err := c.SearchRead(AccountBankStatementLineTransientModel, criteria, NewOptions().Limit(1), abslts); err != nil {
		return nil, err
	}
	return &((*abslts)[0]), nil
}

// FindAccountBankStatementLineTransients finds account.bank.statement.line.transient records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineTransients(criteria *Criteria, options *Options) (*AccountBankStatementLineTransients, error) {
	abslts := &AccountBankStatementLineTransients{}
	if err := c.SearchRead(AccountBankStatementLineTransientModel, criteria, options, abslts); err != nil {
		return nil, err
	}
	return abslts, nil
}

// FindAccountBankStatementLineTransientIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineTransientIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBankStatementLineTransientModel, criteria, options)
}

// FindAccountBankStatementLineTransientId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementLineTransientId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementLineTransientModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
