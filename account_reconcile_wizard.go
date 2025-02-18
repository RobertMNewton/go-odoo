package odoo

// AccountReconcileWizard represents account.reconcile.wizard model.
type AccountReconcileWizard struct {
	AccountId                      *Many2One `xmlrpc:"account_id,omitempty"`
	AllowPartials                  *Bool     `xmlrpc:"allow_partials,omitempty"`
	Amount                         *Float    `xmlrpc:"amount,omitempty"`
	AmountCurrency                 *Float    `xmlrpc:"amount_currency,omitempty"`
	CompanyCurrencyId              *Many2One `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One `xmlrpc:"create_uid,omitempty"`
	Date                           *Time     `xmlrpc:"date,omitempty"`
	DisplayAllowPartials           *Bool     `xmlrpc:"display_allow_partials,omitempty"`
	DisplayName                    *String   `xmlrpc:"display_name,omitempty"`
	EditMode                       *Bool     `xmlrpc:"edit_mode,omitempty"`
	EditModeAmount                 *Float    `xmlrpc:"edit_mode_amount,omitempty"`
	EditModeAmountCurrency         *Float    `xmlrpc:"edit_mode_amount_currency,omitempty"`
	EditModeRecoCurrencyId         *Many2One `xmlrpc:"edit_mode_reco_currency_id,omitempty"`
	ForcePartials                  *Bool     `xmlrpc:"force_partials,omitempty"`
	Id                             *Int      `xmlrpc:"id,omitempty"`
	IsRecPayAccount                *Bool     `xmlrpc:"is_rec_pay_account,omitempty"`
	IsTransferRequired             *Bool     `xmlrpc:"is_transfer_required,omitempty"`
	IsWriteOffRequired             *Bool     `xmlrpc:"is_write_off_required,omitempty"`
	JournalId                      *Many2One `xmlrpc:"journal_id,omitempty"`
	Label                          *String   `xmlrpc:"label,omitempty"`
	LockDateViolatedWarningMessage *String   `xmlrpc:"lock_date_violated_warning_message,omitempty"`
	MoveLineIds                    *Relation `xmlrpc:"move_line_ids,omitempty"`
	RecoAccountId                  *Many2One `xmlrpc:"reco_account_id,omitempty"`
	RecoCurrencyId                 *Many2One `xmlrpc:"reco_currency_id,omitempty"`
	RecoModelAutocompleteIds       *Relation `xmlrpc:"reco_model_autocomplete_ids,omitempty"`
	RecoModelId                    *Many2One `xmlrpc:"reco_model_id,omitempty"`
	SingleCurrencyMode             *Bool     `xmlrpc:"single_currency_mode,omitempty"`
	TaxId                          *Many2One `xmlrpc:"tax_id,omitempty"`
	ToCheck                        *Bool     `xmlrpc:"to_check,omitempty"`
	ToPartnerId                    *Many2One `xmlrpc:"to_partner_id,omitempty"`
	TransferFromAccountId          *Many2One `xmlrpc:"transfer_from_account_id,omitempty"`
	TransferWarningMessage         *String   `xmlrpc:"transfer_warning_message,omitempty"`
	WriteDate                      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileWizards represents array of account.reconcile.wizard model.
type AccountReconcileWizards []AccountReconcileWizard

// AccountReconcileWizardModel is the odoo model name.
const AccountReconcileWizardModel = "account.reconcile.wizard"

// Many2One convert AccountReconcileWizard to *Many2One.
func (arw *AccountReconcileWizard) Many2One() *Many2One {
	return NewMany2One(arw.Id.Get(), "")
}

// CreateAccountReconcileWizard creates a new account.reconcile.wizard model and returns its id.
func (c *Client) CreateAccountReconcileWizard(arw *AccountReconcileWizard) (int64, error) {
	ids, err := c.CreateAccountReconcileWizards([]*AccountReconcileWizard{arw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileWizard creates a new account.reconcile.wizard model and returns its id.
func (c *Client) CreateAccountReconcileWizards(arws []*AccountReconcileWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range arws {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileWizardModel, vv, nil)
}

// UpdateAccountReconcileWizard updates an existing account.reconcile.wizard record.
func (c *Client) UpdateAccountReconcileWizard(arw *AccountReconcileWizard) error {
	return c.UpdateAccountReconcileWizards([]int64{arw.Id.Get()}, arw)
}

// UpdateAccountReconcileWizards updates existing account.reconcile.wizard records.
// All records (represented by ids) will be updated by arw values.
func (c *Client) UpdateAccountReconcileWizards(ids []int64, arw *AccountReconcileWizard) error {
	return c.Update(AccountReconcileWizardModel, ids, arw, nil)
}

// DeleteAccountReconcileWizard deletes an existing account.reconcile.wizard record.
func (c *Client) DeleteAccountReconcileWizard(id int64) error {
	return c.DeleteAccountReconcileWizards([]int64{id})
}

// DeleteAccountReconcileWizards deletes existing account.reconcile.wizard records.
func (c *Client) DeleteAccountReconcileWizards(ids []int64) error {
	return c.Delete(AccountReconcileWizardModel, ids)
}

// GetAccountReconcileWizard gets account.reconcile.wizard existing record.
func (c *Client) GetAccountReconcileWizard(id int64) (*AccountReconcileWizard, error) {
	arws, err := c.GetAccountReconcileWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arws)[0]), nil
}

// GetAccountReconcileWizards gets account.reconcile.wizard existing records.
func (c *Client) GetAccountReconcileWizards(ids []int64) (*AccountReconcileWizards, error) {
	arws := &AccountReconcileWizards{}
	if err := c.Read(AccountReconcileWizardModel, ids, nil, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountReconcileWizard finds account.reconcile.wizard record by querying it with criteria.
func (c *Client) FindAccountReconcileWizard(criteria *Criteria) (*AccountReconcileWizard, error) {
	arws := &AccountReconcileWizards{}
	if err := c.SearchRead(AccountReconcileWizardModel, criteria, NewOptions().Limit(1), arws); err != nil {
		return nil, err
	}
	return &((*arws)[0]), nil
}

// FindAccountReconcileWizards finds account.reconcile.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileWizards(criteria *Criteria, options *Options) (*AccountReconcileWizards, error) {
	arws := &AccountReconcileWizards{}
	if err := c.SearchRead(AccountReconcileWizardModel, criteria, options, arws); err != nil {
		return nil, err
	}
	return arws, nil
}

// FindAccountReconcileWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileWizardModel, criteria, options)
}

// FindAccountReconcileWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
