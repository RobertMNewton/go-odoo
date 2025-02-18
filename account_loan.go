package odoo

// AccountLoan represents account.loan model.
type AccountLoan struct {
	Active                   *Bool       `xmlrpc:"active,omitempty"`
	AmountBorrowed           *Float      `xmlrpc:"amount_borrowed,omitempty"`
	AmountBorrowedDifference *Float      `xmlrpc:"amount_borrowed_difference,omitempty"`
	AssetGroupId             *Many2One   `xmlrpc:"asset_group_id,omitempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omitempty"`
	CountLinkedAssets        *Int        `xmlrpc:"count_linked_assets,omitempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                     *Time       `xmlrpc:"date,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	Duration                 *Int        `xmlrpc:"duration,omitempty"`
	DurationDifference       *Int        `xmlrpc:"duration_difference,omitempty"`
	EndDate                  *Time       `xmlrpc:"end_date,omitempty"`
	ExpenseAccountId         *Many2One   `xmlrpc:"expense_account_id,omitempty"`
	HasMessage               *Bool       `xmlrpc:"has_message,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	Interest                 *Float      `xmlrpc:"interest,omitempty"`
	InterestDifference       *Float      `xmlrpc:"interest_difference,omitempty"`
	IsWrongDate              *Bool       `xmlrpc:"is_wrong_date,omitempty"`
	JournalId                *Many2One   `xmlrpc:"journal_id,omitempty"`
	LineIds                  *Relation   `xmlrpc:"line_ids,omitempty"`
	LinkedAssetsIds          *Relation   `xmlrpc:"linked_assets_ids,omitempty"`
	LoanProperties           interface{} `xmlrpc:"loan_properties,omitempty"`
	LongTermAccountId        *Many2One   `xmlrpc:"long_term_account_id,omitempty"`
	MessageAttachmentCount   *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String     `xmlrpc:"name,omitempty"`
	NbPostedEntries          *Int        `xmlrpc:"nb_posted_entries,omitempty"`
	OutstandingBalance       *Float      `xmlrpc:"outstanding_balance,omitempty"`
	RatingIds                *Relation   `xmlrpc:"rating_ids,omitempty"`
	ShortTermAccountId       *Many2One   `xmlrpc:"short_term_account_id,omitempty"`
	SkipUntilDate            *Time       `xmlrpc:"skip_until_date,omitempty"`
	StartDate                *Time       `xmlrpc:"start_date,omitempty"`
	State                    *Selection  `xmlrpc:"state,omitempty"`
	WebsiteMessageIds        *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountLoans represents array of account.loan model.
type AccountLoans []AccountLoan

// AccountLoanModel is the odoo model name.
const AccountLoanModel = "account.loan"

// Many2One convert AccountLoan to *Many2One.
func (al *AccountLoan) Many2One() *Many2One {
	return NewMany2One(al.Id.Get(), "")
}

// CreateAccountLoan creates a new account.loan model and returns its id.
func (c *Client) CreateAccountLoan(al *AccountLoan) (int64, error) {
	ids, err := c.CreateAccountLoans([]*AccountLoan{al})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLoan creates a new account.loan model and returns its id.
func (c *Client) CreateAccountLoans(als []*AccountLoan) ([]int64, error) {
	var vv []interface{}
	for _, v := range als {
		vv = append(vv, v)
	}
	return c.Create(AccountLoanModel, vv, nil)
}

// UpdateAccountLoan updates an existing account.loan record.
func (c *Client) UpdateAccountLoan(al *AccountLoan) error {
	return c.UpdateAccountLoans([]int64{al.Id.Get()}, al)
}

// UpdateAccountLoans updates existing account.loan records.
// All records (represented by ids) will be updated by al values.
func (c *Client) UpdateAccountLoans(ids []int64, al *AccountLoan) error {
	return c.Update(AccountLoanModel, ids, al, nil)
}

// DeleteAccountLoan deletes an existing account.loan record.
func (c *Client) DeleteAccountLoan(id int64) error {
	return c.DeleteAccountLoans([]int64{id})
}

// DeleteAccountLoans deletes existing account.loan records.
func (c *Client) DeleteAccountLoans(ids []int64) error {
	return c.Delete(AccountLoanModel, ids)
}

// GetAccountLoan gets account.loan existing record.
func (c *Client) GetAccountLoan(id int64) (*AccountLoan, error) {
	als, err := c.GetAccountLoans([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*als)[0]), nil
}

// GetAccountLoans gets account.loan existing records.
func (c *Client) GetAccountLoans(ids []int64) (*AccountLoans, error) {
	als := &AccountLoans{}
	if err := c.Read(AccountLoanModel, ids, nil, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAccountLoan finds account.loan record by querying it with criteria.
func (c *Client) FindAccountLoan(criteria *Criteria) (*AccountLoan, error) {
	als := &AccountLoans{}
	if err := c.SearchRead(AccountLoanModel, criteria, NewOptions().Limit(1), als); err != nil {
		return nil, err
	}
	return &((*als)[0]), nil
}

// FindAccountLoans finds account.loan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoans(criteria *Criteria, options *Options) (*AccountLoans, error) {
	als := &AccountLoans{}
	if err := c.SearchRead(AccountLoanModel, criteria, options, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAccountLoanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountLoanModel, criteria, options)
}

// FindAccountLoanId finds record id by querying it with criteria.
func (c *Client) FindAccountLoanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLoanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
