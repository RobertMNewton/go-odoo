package odoo

// AccountLoanLine represents account.loan.line model.
type AccountLoanLine struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                        *Time      `xmlrpc:"date,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	GeneratedMoveIds            *Relation  `xmlrpc:"generated_move_ids,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	Interest                    *Float     `xmlrpc:"interest,omitempty"`
	IsPaymentMovePosted         *Bool      `xmlrpc:"is_payment_move_posted,omitempty"`
	LoanAssetGroupId            *Many2One  `xmlrpc:"loan_asset_group_id,omitempty"`
	LoanDate                    *Time      `xmlrpc:"loan_date,omitempty"`
	LoanId                      *Many2One  `xmlrpc:"loan_id,omitempty"`
	LoanName                    *String    `xmlrpc:"loan_name,omitempty"`
	LoanState                   *Selection `xmlrpc:"loan_state,omitempty"`
	LongTermTheoreticalBalance  *Float     `xmlrpc:"long_term_theoretical_balance,omitempty"`
	OutstandingBalance          *Float     `xmlrpc:"outstanding_balance,omitempty"`
	Payment                     *Float     `xmlrpc:"payment,omitempty"`
	Principal                   *Float     `xmlrpc:"principal,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	ShortTermTheoreticalBalance *Float     `xmlrpc:"short_term_theoretical_balance,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountLoanLines represents array of account.loan.line model.
type AccountLoanLines []AccountLoanLine

// AccountLoanLineModel is the odoo model name.
const AccountLoanLineModel = "account.loan.line"

// Many2One convert AccountLoanLine to *Many2One.
func (all *AccountLoanLine) Many2One() *Many2One {
	return NewMany2One(all.Id.Get(), "")
}

// CreateAccountLoanLine creates a new account.loan.line model and returns its id.
func (c *Client) CreateAccountLoanLine(all *AccountLoanLine) (int64, error) {
	ids, err := c.CreateAccountLoanLines([]*AccountLoanLine{all})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLoanLine creates a new account.loan.line model and returns its id.
func (c *Client) CreateAccountLoanLines(alls []*AccountLoanLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range alls {
		vv = append(vv, v)
	}
	return c.Create(AccountLoanLineModel, vv, nil)
}

// UpdateAccountLoanLine updates an existing account.loan.line record.
func (c *Client) UpdateAccountLoanLine(all *AccountLoanLine) error {
	return c.UpdateAccountLoanLines([]int64{all.Id.Get()}, all)
}

// UpdateAccountLoanLines updates existing account.loan.line records.
// All records (represented by ids) will be updated by all values.
func (c *Client) UpdateAccountLoanLines(ids []int64, all *AccountLoanLine) error {
	return c.Update(AccountLoanLineModel, ids, all, nil)
}

// DeleteAccountLoanLine deletes an existing account.loan.line record.
func (c *Client) DeleteAccountLoanLine(id int64) error {
	return c.DeleteAccountLoanLines([]int64{id})
}

// DeleteAccountLoanLines deletes existing account.loan.line records.
func (c *Client) DeleteAccountLoanLines(ids []int64) error {
	return c.Delete(AccountLoanLineModel, ids)
}

// GetAccountLoanLine gets account.loan.line existing record.
func (c *Client) GetAccountLoanLine(id int64) (*AccountLoanLine, error) {
	alls, err := c.GetAccountLoanLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*alls)[0]), nil
}

// GetAccountLoanLines gets account.loan.line existing records.
func (c *Client) GetAccountLoanLines(ids []int64) (*AccountLoanLines, error) {
	alls := &AccountLoanLines{}
	if err := c.Read(AccountLoanLineModel, ids, nil, alls); err != nil {
		return nil, err
	}
	return alls, nil
}

// FindAccountLoanLine finds account.loan.line record by querying it with criteria.
func (c *Client) FindAccountLoanLine(criteria *Criteria) (*AccountLoanLine, error) {
	alls := &AccountLoanLines{}
	if err := c.SearchRead(AccountLoanLineModel, criteria, NewOptions().Limit(1), alls); err != nil {
		return nil, err
	}
	return &((*alls)[0]), nil
}

// FindAccountLoanLines finds account.loan.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanLines(criteria *Criteria, options *Options) (*AccountLoanLines, error) {
	alls := &AccountLoanLines{}
	if err := c.SearchRead(AccountLoanLineModel, criteria, options, alls); err != nil {
		return nil, err
	}
	return alls, nil
}

// FindAccountLoanLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountLoanLineModel, criteria, options)
}

// FindAccountLoanLineId finds record id by querying it with criteria.
func (c *Client) FindAccountLoanLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLoanLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
