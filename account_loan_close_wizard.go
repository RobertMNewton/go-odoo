package odoo

// AccountLoanCloseWizard represents account.loan.close.wizard model.
type AccountLoanCloseWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LoanId      *Many2One `xmlrpc:"loan_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountLoanCloseWizards represents array of account.loan.close.wizard model.
type AccountLoanCloseWizards []AccountLoanCloseWizard

// AccountLoanCloseWizardModel is the odoo model name.
const AccountLoanCloseWizardModel = "account.loan.close.wizard"

// Many2One convert AccountLoanCloseWizard to *Many2One.
func (alcw *AccountLoanCloseWizard) Many2One() *Many2One {
	return NewMany2One(alcw.Id.Get(), "")
}

// CreateAccountLoanCloseWizard creates a new account.loan.close.wizard model and returns its id.
func (c *Client) CreateAccountLoanCloseWizard(alcw *AccountLoanCloseWizard) (int64, error) {
	ids, err := c.CreateAccountLoanCloseWizards([]*AccountLoanCloseWizard{alcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLoanCloseWizard creates a new account.loan.close.wizard model and returns its id.
func (c *Client) CreateAccountLoanCloseWizards(alcws []*AccountLoanCloseWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range alcws {
		vv = append(vv, v)
	}
	return c.Create(AccountLoanCloseWizardModel, vv, nil)
}

// UpdateAccountLoanCloseWizard updates an existing account.loan.close.wizard record.
func (c *Client) UpdateAccountLoanCloseWizard(alcw *AccountLoanCloseWizard) error {
	return c.UpdateAccountLoanCloseWizards([]int64{alcw.Id.Get()}, alcw)
}

// UpdateAccountLoanCloseWizards updates existing account.loan.close.wizard records.
// All records (represented by ids) will be updated by alcw values.
func (c *Client) UpdateAccountLoanCloseWizards(ids []int64, alcw *AccountLoanCloseWizard) error {
	return c.Update(AccountLoanCloseWizardModel, ids, alcw, nil)
}

// DeleteAccountLoanCloseWizard deletes an existing account.loan.close.wizard record.
func (c *Client) DeleteAccountLoanCloseWizard(id int64) error {
	return c.DeleteAccountLoanCloseWizards([]int64{id})
}

// DeleteAccountLoanCloseWizards deletes existing account.loan.close.wizard records.
func (c *Client) DeleteAccountLoanCloseWizards(ids []int64) error {
	return c.Delete(AccountLoanCloseWizardModel, ids)
}

// GetAccountLoanCloseWizard gets account.loan.close.wizard existing record.
func (c *Client) GetAccountLoanCloseWizard(id int64) (*AccountLoanCloseWizard, error) {
	alcws, err := c.GetAccountLoanCloseWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*alcws)[0]), nil
}

// GetAccountLoanCloseWizards gets account.loan.close.wizard existing records.
func (c *Client) GetAccountLoanCloseWizards(ids []int64) (*AccountLoanCloseWizards, error) {
	alcws := &AccountLoanCloseWizards{}
	if err := c.Read(AccountLoanCloseWizardModel, ids, nil, alcws); err != nil {
		return nil, err
	}
	return alcws, nil
}

// FindAccountLoanCloseWizard finds account.loan.close.wizard record by querying it with criteria.
func (c *Client) FindAccountLoanCloseWizard(criteria *Criteria) (*AccountLoanCloseWizard, error) {
	alcws := &AccountLoanCloseWizards{}
	if err := c.SearchRead(AccountLoanCloseWizardModel, criteria, NewOptions().Limit(1), alcws); err != nil {
		return nil, err
	}
	return &((*alcws)[0]), nil
}

// FindAccountLoanCloseWizards finds account.loan.close.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanCloseWizards(criteria *Criteria, options *Options) (*AccountLoanCloseWizards, error) {
	alcws := &AccountLoanCloseWizards{}
	if err := c.SearchRead(AccountLoanCloseWizardModel, criteria, options, alcws); err != nil {
		return nil, err
	}
	return alcws, nil
}

// FindAccountLoanCloseWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanCloseWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountLoanCloseWizardModel, criteria, options)
}

// FindAccountLoanCloseWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountLoanCloseWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLoanCloseWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
