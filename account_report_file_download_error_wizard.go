package odoo

// AccountReportFileDownloadErrorWizard represents account.report.file.download.error.wizard model.
type AccountReportFileDownloadErrorWizard struct {
	ActionableErrors interface{} `xmlrpc:"actionable_errors,omitempty"`
	CreateDate       *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String     `xmlrpc:"display_name,omitempty"`
	FileContent      *String     `xmlrpc:"file_content,omitempty"`
	FileName         *String     `xmlrpc:"file_name,omitempty"`
	Id               *Int        `xmlrpc:"id,omitempty"`
	WriteDate        *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountReportFileDownloadErrorWizards represents array of account.report.file.download.error.wizard model.
type AccountReportFileDownloadErrorWizards []AccountReportFileDownloadErrorWizard

// AccountReportFileDownloadErrorWizardModel is the odoo model name.
const AccountReportFileDownloadErrorWizardModel = "account.report.file.download.error.wizard"

// Many2One convert AccountReportFileDownloadErrorWizard to *Many2One.
func (arfdew *AccountReportFileDownloadErrorWizard) Many2One() *Many2One {
	return NewMany2One(arfdew.Id.Get(), "")
}

// CreateAccountReportFileDownloadErrorWizard creates a new account.report.file.download.error.wizard model and returns its id.
func (c *Client) CreateAccountReportFileDownloadErrorWizard(arfdew *AccountReportFileDownloadErrorWizard) (int64, error) {
	ids, err := c.CreateAccountReportFileDownloadErrorWizards([]*AccountReportFileDownloadErrorWizard{arfdew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportFileDownloadErrorWizard creates a new account.report.file.download.error.wizard model and returns its id.
func (c *Client) CreateAccountReportFileDownloadErrorWizards(arfdews []*AccountReportFileDownloadErrorWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range arfdews {
		vv = append(vv, v)
	}
	return c.Create(AccountReportFileDownloadErrorWizardModel, vv, nil)
}

// UpdateAccountReportFileDownloadErrorWizard updates an existing account.report.file.download.error.wizard record.
func (c *Client) UpdateAccountReportFileDownloadErrorWizard(arfdew *AccountReportFileDownloadErrorWizard) error {
	return c.UpdateAccountReportFileDownloadErrorWizards([]int64{arfdew.Id.Get()}, arfdew)
}

// UpdateAccountReportFileDownloadErrorWizards updates existing account.report.file.download.error.wizard records.
// All records (represented by ids) will be updated by arfdew values.
func (c *Client) UpdateAccountReportFileDownloadErrorWizards(ids []int64, arfdew *AccountReportFileDownloadErrorWizard) error {
	return c.Update(AccountReportFileDownloadErrorWizardModel, ids, arfdew, nil)
}

// DeleteAccountReportFileDownloadErrorWizard deletes an existing account.report.file.download.error.wizard record.
func (c *Client) DeleteAccountReportFileDownloadErrorWizard(id int64) error {
	return c.DeleteAccountReportFileDownloadErrorWizards([]int64{id})
}

// DeleteAccountReportFileDownloadErrorWizards deletes existing account.report.file.download.error.wizard records.
func (c *Client) DeleteAccountReportFileDownloadErrorWizards(ids []int64) error {
	return c.Delete(AccountReportFileDownloadErrorWizardModel, ids)
}

// GetAccountReportFileDownloadErrorWizard gets account.report.file.download.error.wizard existing record.
func (c *Client) GetAccountReportFileDownloadErrorWizard(id int64) (*AccountReportFileDownloadErrorWizard, error) {
	arfdews, err := c.GetAccountReportFileDownloadErrorWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arfdews)[0]), nil
}

// GetAccountReportFileDownloadErrorWizards gets account.report.file.download.error.wizard existing records.
func (c *Client) GetAccountReportFileDownloadErrorWizards(ids []int64) (*AccountReportFileDownloadErrorWizards, error) {
	arfdews := &AccountReportFileDownloadErrorWizards{}
	if err := c.Read(AccountReportFileDownloadErrorWizardModel, ids, nil, arfdews); err != nil {
		return nil, err
	}
	return arfdews, nil
}

// FindAccountReportFileDownloadErrorWizard finds account.report.file.download.error.wizard record by querying it with criteria.
func (c *Client) FindAccountReportFileDownloadErrorWizard(criteria *Criteria) (*AccountReportFileDownloadErrorWizard, error) {
	arfdews := &AccountReportFileDownloadErrorWizards{}
	if err := c.SearchRead(AccountReportFileDownloadErrorWizardModel, criteria, NewOptions().Limit(1), arfdews); err != nil {
		return nil, err
	}
	return &((*arfdews)[0]), nil
}

// FindAccountReportFileDownloadErrorWizards finds account.report.file.download.error.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFileDownloadErrorWizards(criteria *Criteria, options *Options) (*AccountReportFileDownloadErrorWizards, error) {
	arfdews := &AccountReportFileDownloadErrorWizards{}
	if err := c.SearchRead(AccountReportFileDownloadErrorWizardModel, criteria, options, arfdews); err != nil {
		return nil, err
	}
	return arfdews, nil
}

// FindAccountReportFileDownloadErrorWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFileDownloadErrorWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportFileDownloadErrorWizardModel, criteria, options)
}

// FindAccountReportFileDownloadErrorWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountReportFileDownloadErrorWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportFileDownloadErrorWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
