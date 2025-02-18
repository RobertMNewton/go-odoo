package odoo

// AccountReportAnnotation represents account.report.annotation model.
type AccountReportAnnotation struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Date             *Time     `xmlrpc:"date,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	FiscalPositionId *Many2One `xmlrpc:"fiscal_position_id,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	LineId           *String   `xmlrpc:"line_id,omitempty"`
	ReportId         *Many2One `xmlrpc:"report_id,omitempty"`
	Text             *String   `xmlrpc:"text,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountReportAnnotations represents array of account.report.annotation model.
type AccountReportAnnotations []AccountReportAnnotation

// AccountReportAnnotationModel is the odoo model name.
const AccountReportAnnotationModel = "account.report.annotation"

// Many2One convert AccountReportAnnotation to *Many2One.
func (ara *AccountReportAnnotation) Many2One() *Many2One {
	return NewMany2One(ara.Id.Get(), "")
}

// CreateAccountReportAnnotation creates a new account.report.annotation model and returns its id.
func (c *Client) CreateAccountReportAnnotation(ara *AccountReportAnnotation) (int64, error) {
	ids, err := c.CreateAccountReportAnnotations([]*AccountReportAnnotation{ara})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportAnnotation creates a new account.report.annotation model and returns its id.
func (c *Client) CreateAccountReportAnnotations(aras []*AccountReportAnnotation) ([]int64, error) {
	var vv []interface{}
	for _, v := range aras {
		vv = append(vv, v)
	}
	return c.Create(AccountReportAnnotationModel, vv, nil)
}

// UpdateAccountReportAnnotation updates an existing account.report.annotation record.
func (c *Client) UpdateAccountReportAnnotation(ara *AccountReportAnnotation) error {
	return c.UpdateAccountReportAnnotations([]int64{ara.Id.Get()}, ara)
}

// UpdateAccountReportAnnotations updates existing account.report.annotation records.
// All records (represented by ids) will be updated by ara values.
func (c *Client) UpdateAccountReportAnnotations(ids []int64, ara *AccountReportAnnotation) error {
	return c.Update(AccountReportAnnotationModel, ids, ara, nil)
}

// DeleteAccountReportAnnotation deletes an existing account.report.annotation record.
func (c *Client) DeleteAccountReportAnnotation(id int64) error {
	return c.DeleteAccountReportAnnotations([]int64{id})
}

// DeleteAccountReportAnnotations deletes existing account.report.annotation records.
func (c *Client) DeleteAccountReportAnnotations(ids []int64) error {
	return c.Delete(AccountReportAnnotationModel, ids)
}

// GetAccountReportAnnotation gets account.report.annotation existing record.
func (c *Client) GetAccountReportAnnotation(id int64) (*AccountReportAnnotation, error) {
	aras, err := c.GetAccountReportAnnotations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aras)[0]), nil
}

// GetAccountReportAnnotations gets account.report.annotation existing records.
func (c *Client) GetAccountReportAnnotations(ids []int64) (*AccountReportAnnotations, error) {
	aras := &AccountReportAnnotations{}
	if err := c.Read(AccountReportAnnotationModel, ids, nil, aras); err != nil {
		return nil, err
	}
	return aras, nil
}

// FindAccountReportAnnotation finds account.report.annotation record by querying it with criteria.
func (c *Client) FindAccountReportAnnotation(criteria *Criteria) (*AccountReportAnnotation, error) {
	aras := &AccountReportAnnotations{}
	if err := c.SearchRead(AccountReportAnnotationModel, criteria, NewOptions().Limit(1), aras); err != nil {
		return nil, err
	}
	return &((*aras)[0]), nil
}

// FindAccountReportAnnotations finds account.report.annotation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportAnnotations(criteria *Criteria, options *Options) (*AccountReportAnnotations, error) {
	aras := &AccountReportAnnotations{}
	if err := c.SearchRead(AccountReportAnnotationModel, criteria, options, aras); err != nil {
		return nil, err
	}
	return aras, nil
}

// FindAccountReportAnnotationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportAnnotationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportAnnotationModel, criteria, options)
}

// FindAccountReportAnnotationId finds record id by querying it with criteria.
func (c *Client) FindAccountReportAnnotationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportAnnotationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
