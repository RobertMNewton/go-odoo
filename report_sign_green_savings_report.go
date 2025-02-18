package odoo

// ReportSignGreenSavingsReport represents report.sign.green_savings_report model.
type ReportSignGreenSavingsReport struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportSignGreenSavingsReports represents array of report.sign.green_savings_report model.
type ReportSignGreenSavingsReports []ReportSignGreenSavingsReport

// ReportSignGreenSavingsReportModel is the odoo model name.
const ReportSignGreenSavingsReportModel = "report.sign.green_savings_report"

// Many2One convert ReportSignGreenSavingsReport to *Many2One.
func (rsg *ReportSignGreenSavingsReport) Many2One() *Many2One {
	return NewMany2One(rsg.Id.Get(), "")
}

// CreateReportSignGreenSavingsReport creates a new report.sign.green_savings_report model and returns its id.
func (c *Client) CreateReportSignGreenSavingsReport(rsg *ReportSignGreenSavingsReport) (int64, error) {
	ids, err := c.CreateReportSignGreenSavingsReports([]*ReportSignGreenSavingsReport{rsg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportSignGreenSavingsReport creates a new report.sign.green_savings_report model and returns its id.
func (c *Client) CreateReportSignGreenSavingsReports(rsgs []*ReportSignGreenSavingsReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range rsgs {
		vv = append(vv, v)
	}
	return c.Create(ReportSignGreenSavingsReportModel, vv, nil)
}

// UpdateReportSignGreenSavingsReport updates an existing report.sign.green_savings_report record.
func (c *Client) UpdateReportSignGreenSavingsReport(rsg *ReportSignGreenSavingsReport) error {
	return c.UpdateReportSignGreenSavingsReports([]int64{rsg.Id.Get()}, rsg)
}

// UpdateReportSignGreenSavingsReports updates existing report.sign.green_savings_report records.
// All records (represented by ids) will be updated by rsg values.
func (c *Client) UpdateReportSignGreenSavingsReports(ids []int64, rsg *ReportSignGreenSavingsReport) error {
	return c.Update(ReportSignGreenSavingsReportModel, ids, rsg, nil)
}

// DeleteReportSignGreenSavingsReport deletes an existing report.sign.green_savings_report record.
func (c *Client) DeleteReportSignGreenSavingsReport(id int64) error {
	return c.DeleteReportSignGreenSavingsReports([]int64{id})
}

// DeleteReportSignGreenSavingsReports deletes existing report.sign.green_savings_report records.
func (c *Client) DeleteReportSignGreenSavingsReports(ids []int64) error {
	return c.Delete(ReportSignGreenSavingsReportModel, ids)
}

// GetReportSignGreenSavingsReport gets report.sign.green_savings_report existing record.
func (c *Client) GetReportSignGreenSavingsReport(id int64) (*ReportSignGreenSavingsReport, error) {
	rsgs, err := c.GetReportSignGreenSavingsReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rsgs)[0]), nil
}

// GetReportSignGreenSavingsReports gets report.sign.green_savings_report existing records.
func (c *Client) GetReportSignGreenSavingsReports(ids []int64) (*ReportSignGreenSavingsReports, error) {
	rsgs := &ReportSignGreenSavingsReports{}
	if err := c.Read(ReportSignGreenSavingsReportModel, ids, nil, rsgs); err != nil {
		return nil, err
	}
	return rsgs, nil
}

// FindReportSignGreenSavingsReport finds report.sign.green_savings_report record by querying it with criteria.
func (c *Client) FindReportSignGreenSavingsReport(criteria *Criteria) (*ReportSignGreenSavingsReport, error) {
	rsgs := &ReportSignGreenSavingsReports{}
	if err := c.SearchRead(ReportSignGreenSavingsReportModel, criteria, NewOptions().Limit(1), rsgs); err != nil {
		return nil, err
	}
	return &((*rsgs)[0]), nil
}

// FindReportSignGreenSavingsReports finds report.sign.green_savings_report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportSignGreenSavingsReports(criteria *Criteria, options *Options) (*ReportSignGreenSavingsReports, error) {
	rsgs := &ReportSignGreenSavingsReports{}
	if err := c.SearchRead(ReportSignGreenSavingsReportModel, criteria, options, rsgs); err != nil {
		return nil, err
	}
	return rsgs, nil
}

// FindReportSignGreenSavingsReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportSignGreenSavingsReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportSignGreenSavingsReportModel, criteria, options)
}

// FindReportSignGreenSavingsReportId finds record id by querying it with criteria.
func (c *Client) FindReportSignGreenSavingsReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportSignGreenSavingsReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
