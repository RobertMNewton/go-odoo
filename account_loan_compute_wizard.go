package odoo

// AccountLoanComputeWizard represents account.loan.compute.wizard model.
type AccountLoanComputeWizard struct {
	CompoundingMethod *Selection `xmlrpc:"compounding_method,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	FirstPaymentDate  *Time      `xmlrpc:"first_payment_date,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	InterestRate      *Float     `xmlrpc:"interest_rate,omitempty"`
	LoanAmount        *Float     `xmlrpc:"loan_amount,omitempty"`
	LoanId            *Many2One  `xmlrpc:"loan_id,omitempty"`
	LoanTerm          *Int       `xmlrpc:"loan_term,omitempty"`
	PaymentEndOfMonth *Selection `xmlrpc:"payment_end_of_month,omitempty"`
	Preview           *String    `xmlrpc:"preview,omitempty"`
	StartDate         *Time      `xmlrpc:"start_date,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountLoanComputeWizards represents array of account.loan.compute.wizard model.
type AccountLoanComputeWizards []AccountLoanComputeWizard

// AccountLoanComputeWizardModel is the odoo model name.
const AccountLoanComputeWizardModel = "account.loan.compute.wizard"

// Many2One convert AccountLoanComputeWizard to *Many2One.
func (alcw *AccountLoanComputeWizard) Many2One() *Many2One {
	return NewMany2One(alcw.Id.Get(), "")
}

// CreateAccountLoanComputeWizard creates a new account.loan.compute.wizard model and returns its id.
func (c *Client) CreateAccountLoanComputeWizard(alcw *AccountLoanComputeWizard) (int64, error) {
	ids, err := c.CreateAccountLoanComputeWizards([]*AccountLoanComputeWizard{alcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLoanComputeWizard creates a new account.loan.compute.wizard model and returns its id.
func (c *Client) CreateAccountLoanComputeWizards(alcws []*AccountLoanComputeWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range alcws {
		vv = append(vv, v)
	}
	return c.Create(AccountLoanComputeWizardModel, vv, nil)
}

// UpdateAccountLoanComputeWizard updates an existing account.loan.compute.wizard record.
func (c *Client) UpdateAccountLoanComputeWizard(alcw *AccountLoanComputeWizard) error {
	return c.UpdateAccountLoanComputeWizards([]int64{alcw.Id.Get()}, alcw)
}

// UpdateAccountLoanComputeWizards updates existing account.loan.compute.wizard records.
// All records (represented by ids) will be updated by alcw values.
func (c *Client) UpdateAccountLoanComputeWizards(ids []int64, alcw *AccountLoanComputeWizard) error {
	return c.Update(AccountLoanComputeWizardModel, ids, alcw, nil)
}

// DeleteAccountLoanComputeWizard deletes an existing account.loan.compute.wizard record.
func (c *Client) DeleteAccountLoanComputeWizard(id int64) error {
	return c.DeleteAccountLoanComputeWizards([]int64{id})
}

// DeleteAccountLoanComputeWizards deletes existing account.loan.compute.wizard records.
func (c *Client) DeleteAccountLoanComputeWizards(ids []int64) error {
	return c.Delete(AccountLoanComputeWizardModel, ids)
}

// GetAccountLoanComputeWizard gets account.loan.compute.wizard existing record.
func (c *Client) GetAccountLoanComputeWizard(id int64) (*AccountLoanComputeWizard, error) {
	alcws, err := c.GetAccountLoanComputeWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*alcws)[0]), nil
}

// GetAccountLoanComputeWizards gets account.loan.compute.wizard existing records.
func (c *Client) GetAccountLoanComputeWizards(ids []int64) (*AccountLoanComputeWizards, error) {
	alcws := &AccountLoanComputeWizards{}
	if err := c.Read(AccountLoanComputeWizardModel, ids, nil, alcws); err != nil {
		return nil, err
	}
	return alcws, nil
}

// FindAccountLoanComputeWizard finds account.loan.compute.wizard record by querying it with criteria.
func (c *Client) FindAccountLoanComputeWizard(criteria *Criteria) (*AccountLoanComputeWizard, error) {
	alcws := &AccountLoanComputeWizards{}
	if err := c.SearchRead(AccountLoanComputeWizardModel, criteria, NewOptions().Limit(1), alcws); err != nil {
		return nil, err
	}
	return &((*alcws)[0]), nil
}

// FindAccountLoanComputeWizards finds account.loan.compute.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanComputeWizards(criteria *Criteria, options *Options) (*AccountLoanComputeWizards, error) {
	alcws := &AccountLoanComputeWizards{}
	if err := c.SearchRead(AccountLoanComputeWizardModel, criteria, options, alcws); err != nil {
		return nil, err
	}
	return alcws, nil
}

// FindAccountLoanComputeWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLoanComputeWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountLoanComputeWizardModel, criteria, options)
}

// FindAccountLoanComputeWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountLoanComputeWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLoanComputeWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
