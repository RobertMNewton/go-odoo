package odoo

// AccountBatchErrorWizardLine represents account.batch.error.wizard.line model.
type AccountBatchErrorWizardLine struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Description      *String   `xmlrpc:"description,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	ErrorWizardId    *Many2One `xmlrpc:"error_wizard_id,omitempty"`
	HelpMessage      *String   `xmlrpc:"help_message,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	PaymentIds       *Relation `xmlrpc:"payment_ids,omitempty"`
	ShowRemoveButton *Bool     `xmlrpc:"show_remove_button,omitempty"`
	WarningWizardId  *Many2One `xmlrpc:"warning_wizard_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountBatchErrorWizardLines represents array of account.batch.error.wizard.line model.
type AccountBatchErrorWizardLines []AccountBatchErrorWizardLine

// AccountBatchErrorWizardLineModel is the odoo model name.
const AccountBatchErrorWizardLineModel = "account.batch.error.wizard.line"

// Many2One convert AccountBatchErrorWizardLine to *Many2One.
func (abewl *AccountBatchErrorWizardLine) Many2One() *Many2One {
	return NewMany2One(abewl.Id.Get(), "")
}

// CreateAccountBatchErrorWizardLine creates a new account.batch.error.wizard.line model and returns its id.
func (c *Client) CreateAccountBatchErrorWizardLine(abewl *AccountBatchErrorWizardLine) (int64, error) {
	ids, err := c.CreateAccountBatchErrorWizardLines([]*AccountBatchErrorWizardLine{abewl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchErrorWizardLine creates a new account.batch.error.wizard.line model and returns its id.
func (c *Client) CreateAccountBatchErrorWizardLines(abewls []*AccountBatchErrorWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range abewls {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchErrorWizardLineModel, vv, nil)
}

// UpdateAccountBatchErrorWizardLine updates an existing account.batch.error.wizard.line record.
func (c *Client) UpdateAccountBatchErrorWizardLine(abewl *AccountBatchErrorWizardLine) error {
	return c.UpdateAccountBatchErrorWizardLines([]int64{abewl.Id.Get()}, abewl)
}

// UpdateAccountBatchErrorWizardLines updates existing account.batch.error.wizard.line records.
// All records (represented by ids) will be updated by abewl values.
func (c *Client) UpdateAccountBatchErrorWizardLines(ids []int64, abewl *AccountBatchErrorWizardLine) error {
	return c.Update(AccountBatchErrorWizardLineModel, ids, abewl, nil)
}

// DeleteAccountBatchErrorWizardLine deletes an existing account.batch.error.wizard.line record.
func (c *Client) DeleteAccountBatchErrorWizardLine(id int64) error {
	return c.DeleteAccountBatchErrorWizardLines([]int64{id})
}

// DeleteAccountBatchErrorWizardLines deletes existing account.batch.error.wizard.line records.
func (c *Client) DeleteAccountBatchErrorWizardLines(ids []int64) error {
	return c.Delete(AccountBatchErrorWizardLineModel, ids)
}

// GetAccountBatchErrorWizardLine gets account.batch.error.wizard.line existing record.
func (c *Client) GetAccountBatchErrorWizardLine(id int64) (*AccountBatchErrorWizardLine, error) {
	abewls, err := c.GetAccountBatchErrorWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abewls)[0]), nil
}

// GetAccountBatchErrorWizardLines gets account.batch.error.wizard.line existing records.
func (c *Client) GetAccountBatchErrorWizardLines(ids []int64) (*AccountBatchErrorWizardLines, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.Read(AccountBatchErrorWizardLineModel, ids, nil, abewls); err != nil {
		return nil, err
	}
	return abewls, nil
}

// FindAccountBatchErrorWizardLine finds account.batch.error.wizard.line record by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizardLine(criteria *Criteria) (*AccountBatchErrorWizardLine, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.SearchRead(AccountBatchErrorWizardLineModel, criteria, NewOptions().Limit(1), abewls); err != nil {
		return nil, err
	}
	return &((*abewls)[0]), nil
}

// FindAccountBatchErrorWizardLines finds account.batch.error.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizardLines(criteria *Criteria, options *Options) (*AccountBatchErrorWizardLines, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.SearchRead(AccountBatchErrorWizardLineModel, criteria, options, abewls); err != nil {
		return nil, err
	}
	return abewls, nil
}

// FindAccountBatchErrorWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBatchErrorWizardLineModel, criteria, options)
}

// FindAccountBatchErrorWizardLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchErrorWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
