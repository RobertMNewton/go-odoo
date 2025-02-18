package odoo

// AccountMulticurrencyRevaluationWizard represents account.multicurrency.revaluation.wizard model.
type AccountMulticurrencyRevaluationWizard struct {
	CompanyId                 *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	Date                      *Time     `xmlrpc:"date,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	ExpenseProvisionAccountId *Many2One `xmlrpc:"expense_provision_account_id,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	IncomeProvisionAccountId  *Many2One `xmlrpc:"income_provision_account_id,omitempty"`
	JournalId                 *Many2One `xmlrpc:"journal_id,omitempty"`
	PreviewData               *String   `xmlrpc:"preview_data,omitempty"`
	ReversalDate              *Time     `xmlrpc:"reversal_date,omitempty"`
	ShowWarningMoveId         *Many2One `xmlrpc:"show_warning_move_id,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountMulticurrencyRevaluationWizards represents array of account.multicurrency.revaluation.wizard model.
type AccountMulticurrencyRevaluationWizards []AccountMulticurrencyRevaluationWizard

// AccountMulticurrencyRevaluationWizardModel is the odoo model name.
const AccountMulticurrencyRevaluationWizardModel = "account.multicurrency.revaluation.wizard"

// Many2One convert AccountMulticurrencyRevaluationWizard to *Many2One.
func (amrw *AccountMulticurrencyRevaluationWizard) Many2One() *Many2One {
	return NewMany2One(amrw.Id.Get(), "")
}

// CreateAccountMulticurrencyRevaluationWizard creates a new account.multicurrency.revaluation.wizard model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationWizard(amrw *AccountMulticurrencyRevaluationWizard) (int64, error) {
	ids, err := c.CreateAccountMulticurrencyRevaluationWizards([]*AccountMulticurrencyRevaluationWizard{amrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMulticurrencyRevaluationWizard creates a new account.multicurrency.revaluation.wizard model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationWizards(amrws []*AccountMulticurrencyRevaluationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amrws {
		vv = append(vv, v)
	}
	return c.Create(AccountMulticurrencyRevaluationWizardModel, vv, nil)
}

// UpdateAccountMulticurrencyRevaluationWizard updates an existing account.multicurrency.revaluation.wizard record.
func (c *Client) UpdateAccountMulticurrencyRevaluationWizard(amrw *AccountMulticurrencyRevaluationWizard) error {
	return c.UpdateAccountMulticurrencyRevaluationWizards([]int64{amrw.Id.Get()}, amrw)
}

// UpdateAccountMulticurrencyRevaluationWizards updates existing account.multicurrency.revaluation.wizard records.
// All records (represented by ids) will be updated by amrw values.
func (c *Client) UpdateAccountMulticurrencyRevaluationWizards(ids []int64, amrw *AccountMulticurrencyRevaluationWizard) error {
	return c.Update(AccountMulticurrencyRevaluationWizardModel, ids, amrw, nil)
}

// DeleteAccountMulticurrencyRevaluationWizard deletes an existing account.multicurrency.revaluation.wizard record.
func (c *Client) DeleteAccountMulticurrencyRevaluationWizard(id int64) error {
	return c.DeleteAccountMulticurrencyRevaluationWizards([]int64{id})
}

// DeleteAccountMulticurrencyRevaluationWizards deletes existing account.multicurrency.revaluation.wizard records.
func (c *Client) DeleteAccountMulticurrencyRevaluationWizards(ids []int64) error {
	return c.Delete(AccountMulticurrencyRevaluationWizardModel, ids)
}

// GetAccountMulticurrencyRevaluationWizard gets account.multicurrency.revaluation.wizard existing record.
func (c *Client) GetAccountMulticurrencyRevaluationWizard(id int64) (*AccountMulticurrencyRevaluationWizard, error) {
	amrws, err := c.GetAccountMulticurrencyRevaluationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amrws)[0]), nil
}

// GetAccountMulticurrencyRevaluationWizards gets account.multicurrency.revaluation.wizard existing records.
func (c *Client) GetAccountMulticurrencyRevaluationWizards(ids []int64) (*AccountMulticurrencyRevaluationWizards, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.Read(AccountMulticurrencyRevaluationWizardModel, ids, nil, amrws); err != nil {
		return nil, err
	}
	return amrws, nil
}

// FindAccountMulticurrencyRevaluationWizard finds account.multicurrency.revaluation.wizard record by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationWizard(criteria *Criteria) (*AccountMulticurrencyRevaluationWizard, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationWizardModel, criteria, NewOptions().Limit(1), amrws); err != nil {
		return nil, err
	}
	return &((*amrws)[0]), nil
}

// FindAccountMulticurrencyRevaluationWizards finds account.multicurrency.revaluation.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationWizards(criteria *Criteria, options *Options) (*AccountMulticurrencyRevaluationWizards, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationWizardModel, criteria, options, amrws); err != nil {
		return nil, err
	}
	return amrws, nil
}

// FindAccountMulticurrencyRevaluationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMulticurrencyRevaluationWizardModel, criteria, options)
}

// FindAccountMulticurrencyRevaluationWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMulticurrencyRevaluationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
