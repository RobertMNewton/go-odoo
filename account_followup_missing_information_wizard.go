package odoo

// AccountFollowupMissingInformationWizard represents account_followup.missing.information.wizard model.
type AccountFollowupMissingInformationWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFollowupMissingInformationWizards represents array of account_followup.missing.information.wizard model.
type AccountFollowupMissingInformationWizards []AccountFollowupMissingInformationWizard

// AccountFollowupMissingInformationWizardModel is the odoo model name.
const AccountFollowupMissingInformationWizardModel = "account_followup.missing.information.wizard"

// Many2One convert AccountFollowupMissingInformationWizard to *Many2One.
func (amiw *AccountFollowupMissingInformationWizard) Many2One() *Many2One {
	return NewMany2One(amiw.Id.Get(), "")
}

// CreateAccountFollowupMissingInformationWizard creates a new account_followup.missing.information.wizard model and returns its id.
func (c *Client) CreateAccountFollowupMissingInformationWizard(amiw *AccountFollowupMissingInformationWizard) (int64, error) {
	ids, err := c.CreateAccountFollowupMissingInformationWizards([]*AccountFollowupMissingInformationWizard{amiw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupMissingInformationWizard creates a new account_followup.missing.information.wizard model and returns its id.
func (c *Client) CreateAccountFollowupMissingInformationWizards(amiws []*AccountFollowupMissingInformationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amiws {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupMissingInformationWizardModel, vv, nil)
}

// UpdateAccountFollowupMissingInformationWizard updates an existing account_followup.missing.information.wizard record.
func (c *Client) UpdateAccountFollowupMissingInformationWizard(amiw *AccountFollowupMissingInformationWizard) error {
	return c.UpdateAccountFollowupMissingInformationWizards([]int64{amiw.Id.Get()}, amiw)
}

// UpdateAccountFollowupMissingInformationWizards updates existing account_followup.missing.information.wizard records.
// All records (represented by ids) will be updated by amiw values.
func (c *Client) UpdateAccountFollowupMissingInformationWizards(ids []int64, amiw *AccountFollowupMissingInformationWizard) error {
	return c.Update(AccountFollowupMissingInformationWizardModel, ids, amiw, nil)
}

// DeleteAccountFollowupMissingInformationWizard deletes an existing account_followup.missing.information.wizard record.
func (c *Client) DeleteAccountFollowupMissingInformationWizard(id int64) error {
	return c.DeleteAccountFollowupMissingInformationWizards([]int64{id})
}

// DeleteAccountFollowupMissingInformationWizards deletes existing account_followup.missing.information.wizard records.
func (c *Client) DeleteAccountFollowupMissingInformationWizards(ids []int64) error {
	return c.Delete(AccountFollowupMissingInformationWizardModel, ids)
}

// GetAccountFollowupMissingInformationWizard gets account_followup.missing.information.wizard existing record.
func (c *Client) GetAccountFollowupMissingInformationWizard(id int64) (*AccountFollowupMissingInformationWizard, error) {
	amiws, err := c.GetAccountFollowupMissingInformationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amiws)[0]), nil
}

// GetAccountFollowupMissingInformationWizards gets account_followup.missing.information.wizard existing records.
func (c *Client) GetAccountFollowupMissingInformationWizards(ids []int64) (*AccountFollowupMissingInformationWizards, error) {
	amiws := &AccountFollowupMissingInformationWizards{}
	if err := c.Read(AccountFollowupMissingInformationWizardModel, ids, nil, amiws); err != nil {
		return nil, err
	}
	return amiws, nil
}

// FindAccountFollowupMissingInformationWizard finds account_followup.missing.information.wizard record by querying it with criteria.
func (c *Client) FindAccountFollowupMissingInformationWizard(criteria *Criteria) (*AccountFollowupMissingInformationWizard, error) {
	amiws := &AccountFollowupMissingInformationWizards{}
	if err := c.SearchRead(AccountFollowupMissingInformationWizardModel, criteria, NewOptions().Limit(1), amiws); err != nil {
		return nil, err
	}
	return &((*amiws)[0]), nil
}

// FindAccountFollowupMissingInformationWizards finds account_followup.missing.information.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupMissingInformationWizards(criteria *Criteria, options *Options) (*AccountFollowupMissingInformationWizards, error) {
	amiws := &AccountFollowupMissingInformationWizards{}
	if err := c.SearchRead(AccountFollowupMissingInformationWizardModel, criteria, options, amiws); err != nil {
		return nil, err
	}
	return amiws, nil
}

// FindAccountFollowupMissingInformationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupMissingInformationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFollowupMissingInformationWizardModel, criteria, options)
}

// FindAccountFollowupMissingInformationWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupMissingInformationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupMissingInformationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
