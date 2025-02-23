package odoo

// ReportAccountReportInvoice represents report.account.report_invoice model.
type ReportAccountReportInvoice struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountReportInvoices represents array of report.account.report_invoice model.
type ReportAccountReportInvoices []ReportAccountReportInvoice

// ReportAccountReportInvoiceModel is the odoo model name.
const ReportAccountReportInvoiceModel = "report.account.report_invoice"

// Many2One convert ReportAccountReportInvoice to *Many2One.
func (rar *ReportAccountReportInvoice) Many2One() *Many2One {
	return NewMany2One(rar.Id.Get(), "")
}

// CreateReportAccountReportInvoice creates a new report.account.report_invoice model and returns its id.
func (c *Client) CreateReportAccountReportInvoice(rar *ReportAccountReportInvoice) (int64, error) {
	ids, err := c.CreateReportAccountReportInvoices([]*ReportAccountReportInvoice{rar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountReportInvoice creates a new report.account.report_invoice model and returns its id.
func (c *Client) CreateReportAccountReportInvoices(rars []*ReportAccountReportInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range rars {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountReportInvoiceModel, vv, nil)
}

// UpdateReportAccountReportInvoice updates an existing report.account.report_invoice record.
func (c *Client) UpdateReportAccountReportInvoice(rar *ReportAccountReportInvoice) error {
	return c.UpdateReportAccountReportInvoices([]int64{rar.Id.Get()}, rar)
}

// UpdateReportAccountReportInvoices updates existing report.account.report_invoice records.
// All records (represented by ids) will be updated by rar values.
func (c *Client) UpdateReportAccountReportInvoices(ids []int64, rar *ReportAccountReportInvoice) error {
	return c.Update(ReportAccountReportInvoiceModel, ids, rar, nil)
}

// DeleteReportAccountReportInvoice deletes an existing report.account.report_invoice record.
func (c *Client) DeleteReportAccountReportInvoice(id int64) error {
	return c.DeleteReportAccountReportInvoices([]int64{id})
}

// DeleteReportAccountReportInvoices deletes existing report.account.report_invoice records.
func (c *Client) DeleteReportAccountReportInvoices(ids []int64) error {
	return c.Delete(ReportAccountReportInvoiceModel, ids)
}

// GetReportAccountReportInvoice gets report.account.report_invoice existing record.
func (c *Client) GetReportAccountReportInvoice(id int64) (*ReportAccountReportInvoice, error) {
	rars, err := c.GetReportAccountReportInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// GetReportAccountReportInvoices gets report.account.report_invoice existing records.
func (c *Client) GetReportAccountReportInvoices(ids []int64) (*ReportAccountReportInvoices, error) {
	rars := &ReportAccountReportInvoices{}
	if err := c.Read(ReportAccountReportInvoiceModel, ids, nil, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportInvoice finds report.account.report_invoice record by querying it with criteria.
func (c *Client) FindReportAccountReportInvoice(criteria *Criteria) (*ReportAccountReportInvoice, error) {
	rars := &ReportAccountReportInvoices{}
	if err := c.SearchRead(ReportAccountReportInvoiceModel, criteria, NewOptions().Limit(1), rars); err != nil {
		return nil, err
	}
	return &((*rars)[0]), nil
}

// FindReportAccountReportInvoices finds report.account.report_invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportInvoices(criteria *Criteria, options *Options) (*ReportAccountReportInvoices, error) {
	rars := &ReportAccountReportInvoices{}
	if err := c.SearchRead(ReportAccountReportInvoiceModel, criteria, options, rars); err != nil {
		return nil, err
	}
	return rars, nil
}

// FindReportAccountReportInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountReportInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAccountReportInvoiceModel, criteria, options)
}

// FindReportAccountReportInvoiceId finds record id by querying it with criteria.
func (c *Client) FindReportAccountReportInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountReportInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
