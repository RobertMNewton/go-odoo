package odoo

// ReportAccountBatchPaymentPrintBatchPayment represents report.account_batch_payment.print_batch_payment model.
type ReportAccountBatchPaymentPrintBatchPayment struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportAccountBatchPaymentPrintBatchPayments represents array of report.account_batch_payment.print_batch_payment model.
type ReportAccountBatchPaymentPrintBatchPayments []ReportAccountBatchPaymentPrintBatchPayment

// ReportAccountBatchPaymentPrintBatchPaymentModel is the odoo model name.
const ReportAccountBatchPaymentPrintBatchPaymentModel = "report.account_batch_payment.print_batch_payment"

// Many2One convert ReportAccountBatchPaymentPrintBatchPayment to *Many2One.
func (rap *ReportAccountBatchPaymentPrintBatchPayment) Many2One() *Many2One {
	return NewMany2One(rap.Id.Get(), "")
}

// CreateReportAccountBatchPaymentPrintBatchPayment creates a new report.account_batch_payment.print_batch_payment model and returns its id.
func (c *Client) CreateReportAccountBatchPaymentPrintBatchPayment(rap *ReportAccountBatchPaymentPrintBatchPayment) (int64, error) {
	ids, err := c.CreateReportAccountBatchPaymentPrintBatchPayments([]*ReportAccountBatchPaymentPrintBatchPayment{rap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAccountBatchPaymentPrintBatchPayment creates a new report.account_batch_payment.print_batch_payment model and returns its id.
func (c *Client) CreateReportAccountBatchPaymentPrintBatchPayments(raps []*ReportAccountBatchPaymentPrintBatchPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range raps {
		vv = append(vv, v)
	}
	return c.Create(ReportAccountBatchPaymentPrintBatchPaymentModel, vv, nil)
}

// UpdateReportAccountBatchPaymentPrintBatchPayment updates an existing report.account_batch_payment.print_batch_payment record.
func (c *Client) UpdateReportAccountBatchPaymentPrintBatchPayment(rap *ReportAccountBatchPaymentPrintBatchPayment) error {
	return c.UpdateReportAccountBatchPaymentPrintBatchPayments([]int64{rap.Id.Get()}, rap)
}

// UpdateReportAccountBatchPaymentPrintBatchPayments updates existing report.account_batch_payment.print_batch_payment records.
// All records (represented by ids) will be updated by rap values.
func (c *Client) UpdateReportAccountBatchPaymentPrintBatchPayments(ids []int64, rap *ReportAccountBatchPaymentPrintBatchPayment) error {
	return c.Update(ReportAccountBatchPaymentPrintBatchPaymentModel, ids, rap, nil)
}

// DeleteReportAccountBatchPaymentPrintBatchPayment deletes an existing report.account_batch_payment.print_batch_payment record.
func (c *Client) DeleteReportAccountBatchPaymentPrintBatchPayment(id int64) error {
	return c.DeleteReportAccountBatchPaymentPrintBatchPayments([]int64{id})
}

// DeleteReportAccountBatchPaymentPrintBatchPayments deletes existing report.account_batch_payment.print_batch_payment records.
func (c *Client) DeleteReportAccountBatchPaymentPrintBatchPayments(ids []int64) error {
	return c.Delete(ReportAccountBatchPaymentPrintBatchPaymentModel, ids)
}

// GetReportAccountBatchPaymentPrintBatchPayment gets report.account_batch_payment.print_batch_payment existing record.
func (c *Client) GetReportAccountBatchPaymentPrintBatchPayment(id int64) (*ReportAccountBatchPaymentPrintBatchPayment, error) {
	raps, err := c.GetReportAccountBatchPaymentPrintBatchPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*raps)[0]), nil
}

// GetReportAccountBatchPaymentPrintBatchPayments gets report.account_batch_payment.print_batch_payment existing records.
func (c *Client) GetReportAccountBatchPaymentPrintBatchPayments(ids []int64) (*ReportAccountBatchPaymentPrintBatchPayments, error) {
	raps := &ReportAccountBatchPaymentPrintBatchPayments{}
	if err := c.Read(ReportAccountBatchPaymentPrintBatchPaymentModel, ids, nil, raps); err != nil {
		return nil, err
	}
	return raps, nil
}

// FindReportAccountBatchPaymentPrintBatchPayment finds report.account_batch_payment.print_batch_payment record by querying it with criteria.
func (c *Client) FindReportAccountBatchPaymentPrintBatchPayment(criteria *Criteria) (*ReportAccountBatchPaymentPrintBatchPayment, error) {
	raps := &ReportAccountBatchPaymentPrintBatchPayments{}
	if err := c.SearchRead(ReportAccountBatchPaymentPrintBatchPaymentModel, criteria, NewOptions().Limit(1), raps); err != nil {
		return nil, err
	}
	return &((*raps)[0]), nil
}

// FindReportAccountBatchPaymentPrintBatchPayments finds report.account_batch_payment.print_batch_payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountBatchPaymentPrintBatchPayments(criteria *Criteria, options *Options) (*ReportAccountBatchPaymentPrintBatchPayments, error) {
	raps := &ReportAccountBatchPaymentPrintBatchPayments{}
	if err := c.SearchRead(ReportAccountBatchPaymentPrintBatchPaymentModel, criteria, options, raps); err != nil {
		return nil, err
	}
	return raps, nil
}

// FindReportAccountBatchPaymentPrintBatchPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAccountBatchPaymentPrintBatchPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAccountBatchPaymentPrintBatchPaymentModel, criteria, options)
}

// FindReportAccountBatchPaymentPrintBatchPaymentId finds record id by querying it with criteria.
func (c *Client) FindReportAccountBatchPaymentPrintBatchPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAccountBatchPaymentPrintBatchPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
