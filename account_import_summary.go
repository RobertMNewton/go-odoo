package odoo

// AccountImportSummary represents account.import.summary model.
type AccountImportSummary struct {
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	ImportSummaryAccountIds *Relation `xmlrpc:"import_summary_account_ids,omitempty"`
	ImportSummaryHaveData   *Bool     `xmlrpc:"import_summary_have_data,omitempty"`
	ImportSummaryJournalIds *Relation `xmlrpc:"import_summary_journal_ids,omitempty"`
	ImportSummaryLenAccount *Int      `xmlrpc:"import_summary_len_account,omitempty"`
	ImportSummaryLenJournal *Int      `xmlrpc:"import_summary_len_journal,omitempty"`
	ImportSummaryLenMove    *Int      `xmlrpc:"import_summary_len_move,omitempty"`
	ImportSummaryLenPartner *Int      `xmlrpc:"import_summary_len_partner,omitempty"`
	ImportSummaryLenTax     *Int      `xmlrpc:"import_summary_len_tax,omitempty"`
	ImportSummaryMoveIds    *Relation `xmlrpc:"import_summary_move_ids,omitempty"`
	ImportSummaryName       *String   `xmlrpc:"import_summary_name,omitempty"`
	ImportSummaryPartnerIds *Relation `xmlrpc:"import_summary_partner_ids,omitempty"`
	ImportSummaryTaxIds     *Relation `xmlrpc:"import_summary_tax_ids,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountImportSummarys represents array of account.import.summary model.
type AccountImportSummarys []AccountImportSummary

// AccountImportSummaryModel is the odoo model name.
const AccountImportSummaryModel = "account.import.summary"

// Many2One convert AccountImportSummary to *Many2One.
func (ais *AccountImportSummary) Many2One() *Many2One {
	return NewMany2One(ais.Id.Get(), "")
}

// CreateAccountImportSummary creates a new account.import.summary model and returns its id.
func (c *Client) CreateAccountImportSummary(ais *AccountImportSummary) (int64, error) {
	ids, err := c.CreateAccountImportSummarys([]*AccountImportSummary{ais})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountImportSummary creates a new account.import.summary model and returns its id.
func (c *Client) CreateAccountImportSummarys(aiss []*AccountImportSummary) ([]int64, error) {
	var vv []interface{}
	for _, v := range aiss {
		vv = append(vv, v)
	}
	return c.Create(AccountImportSummaryModel, vv, nil)
}

// UpdateAccountImportSummary updates an existing account.import.summary record.
func (c *Client) UpdateAccountImportSummary(ais *AccountImportSummary) error {
	return c.UpdateAccountImportSummarys([]int64{ais.Id.Get()}, ais)
}

// UpdateAccountImportSummarys updates existing account.import.summary records.
// All records (represented by ids) will be updated by ais values.
func (c *Client) UpdateAccountImportSummarys(ids []int64, ais *AccountImportSummary) error {
	return c.Update(AccountImportSummaryModel, ids, ais, nil)
}

// DeleteAccountImportSummary deletes an existing account.import.summary record.
func (c *Client) DeleteAccountImportSummary(id int64) error {
	return c.DeleteAccountImportSummarys([]int64{id})
}

// DeleteAccountImportSummarys deletes existing account.import.summary records.
func (c *Client) DeleteAccountImportSummarys(ids []int64) error {
	return c.Delete(AccountImportSummaryModel, ids)
}

// GetAccountImportSummary gets account.import.summary existing record.
func (c *Client) GetAccountImportSummary(id int64) (*AccountImportSummary, error) {
	aiss, err := c.GetAccountImportSummarys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aiss)[0]), nil
}

// GetAccountImportSummarys gets account.import.summary existing records.
func (c *Client) GetAccountImportSummarys(ids []int64) (*AccountImportSummarys, error) {
	aiss := &AccountImportSummarys{}
	if err := c.Read(AccountImportSummaryModel, ids, nil, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountImportSummary finds account.import.summary record by querying it with criteria.
func (c *Client) FindAccountImportSummary(criteria *Criteria) (*AccountImportSummary, error) {
	aiss := &AccountImportSummarys{}
	if err := c.SearchRead(AccountImportSummaryModel, criteria, NewOptions().Limit(1), aiss); err != nil {
		return nil, err
	}
	return &((*aiss)[0]), nil
}

// FindAccountImportSummarys finds account.import.summary records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountImportSummarys(criteria *Criteria, options *Options) (*AccountImportSummarys, error) {
	aiss := &AccountImportSummarys{}
	if err := c.SearchRead(AccountImportSummaryModel, criteria, options, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountImportSummaryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountImportSummaryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountImportSummaryModel, criteria, options)
}

// FindAccountImportSummaryId finds record id by querying it with criteria.
func (c *Client) FindAccountImportSummaryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountImportSummaryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
