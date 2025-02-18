package odoo

// AccountAutoReconcileWizard represents account.auto.reconcile.wizard model.
type AccountAutoReconcileWizard struct {
	AccountIds  *Relation  `xmlrpc:"account_ids,omitempty"`
	CompanyId   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	FromDate    *Time      `xmlrpc:"from_date,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LineIds     *Relation  `xmlrpc:"line_ids,omitempty"`
	PartnerIds  *Relation  `xmlrpc:"partner_ids,omitempty"`
	SearchMode  *Selection `xmlrpc:"search_mode,omitempty"`
	ToDate      *Time      `xmlrpc:"to_date,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAutoReconcileWizards represents array of account.auto.reconcile.wizard model.
type AccountAutoReconcileWizards []AccountAutoReconcileWizard

// AccountAutoReconcileWizardModel is the odoo model name.
const AccountAutoReconcileWizardModel = "account.auto.reconcile.wizard"

// Many2One convert AccountAutoReconcileWizard to *Many2One.
func (aarw *AccountAutoReconcileWizard) Many2One() *Many2One {
	return NewMany2One(aarw.Id.Get(), "")
}

// CreateAccountAutoReconcileWizard creates a new account.auto.reconcile.wizard model and returns its id.
func (c *Client) CreateAccountAutoReconcileWizard(aarw *AccountAutoReconcileWizard) (int64, error) {
	ids, err := c.CreateAccountAutoReconcileWizards([]*AccountAutoReconcileWizard{aarw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAutoReconcileWizard creates a new account.auto.reconcile.wizard model and returns its id.
func (c *Client) CreateAccountAutoReconcileWizards(aarws []*AccountAutoReconcileWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aarws {
		vv = append(vv, v)
	}
	return c.Create(AccountAutoReconcileWizardModel, vv, nil)
}

// UpdateAccountAutoReconcileWizard updates an existing account.auto.reconcile.wizard record.
func (c *Client) UpdateAccountAutoReconcileWizard(aarw *AccountAutoReconcileWizard) error {
	return c.UpdateAccountAutoReconcileWizards([]int64{aarw.Id.Get()}, aarw)
}

// UpdateAccountAutoReconcileWizards updates existing account.auto.reconcile.wizard records.
// All records (represented by ids) will be updated by aarw values.
func (c *Client) UpdateAccountAutoReconcileWizards(ids []int64, aarw *AccountAutoReconcileWizard) error {
	return c.Update(AccountAutoReconcileWizardModel, ids, aarw, nil)
}

// DeleteAccountAutoReconcileWizard deletes an existing account.auto.reconcile.wizard record.
func (c *Client) DeleteAccountAutoReconcileWizard(id int64) error {
	return c.DeleteAccountAutoReconcileWizards([]int64{id})
}

// DeleteAccountAutoReconcileWizards deletes existing account.auto.reconcile.wizard records.
func (c *Client) DeleteAccountAutoReconcileWizards(ids []int64) error {
	return c.Delete(AccountAutoReconcileWizardModel, ids)
}

// GetAccountAutoReconcileWizard gets account.auto.reconcile.wizard existing record.
func (c *Client) GetAccountAutoReconcileWizard(id int64) (*AccountAutoReconcileWizard, error) {
	aarws, err := c.GetAccountAutoReconcileWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aarws)[0]), nil
}

// GetAccountAutoReconcileWizards gets account.auto.reconcile.wizard existing records.
func (c *Client) GetAccountAutoReconcileWizards(ids []int64) (*AccountAutoReconcileWizards, error) {
	aarws := &AccountAutoReconcileWizards{}
	if err := c.Read(AccountAutoReconcileWizardModel, ids, nil, aarws); err != nil {
		return nil, err
	}
	return aarws, nil
}

// FindAccountAutoReconcileWizard finds account.auto.reconcile.wizard record by querying it with criteria.
func (c *Client) FindAccountAutoReconcileWizard(criteria *Criteria) (*AccountAutoReconcileWizard, error) {
	aarws := &AccountAutoReconcileWizards{}
	if err := c.SearchRead(AccountAutoReconcileWizardModel, criteria, NewOptions().Limit(1), aarws); err != nil {
		return nil, err
	}
	return &((*aarws)[0]), nil
}

// FindAccountAutoReconcileWizards finds account.auto.reconcile.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutoReconcileWizards(criteria *Criteria, options *Options) (*AccountAutoReconcileWizards, error) {
	aarws := &AccountAutoReconcileWizards{}
	if err := c.SearchRead(AccountAutoReconcileWizardModel, criteria, options, aarws); err != nil {
		return nil, err
	}
	return aarws, nil
}

// FindAccountAutoReconcileWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutoReconcileWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAutoReconcileWizardModel, criteria, options)
}

// FindAccountAutoReconcileWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAutoReconcileWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAutoReconcileWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
