package odoo

// ReportQualityControlQualityWorksheetInternal represents report.quality_control.quality_worksheet_internal model.
type ReportQualityControlQualityWorksheetInternal struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportQualityControlQualityWorksheetInternals represents array of report.quality_control.quality_worksheet_internal model.
type ReportQualityControlQualityWorksheetInternals []ReportQualityControlQualityWorksheetInternal

// ReportQualityControlQualityWorksheetInternalModel is the odoo model name.
const ReportQualityControlQualityWorksheetInternalModel = "report.quality_control.quality_worksheet_internal"

// Many2One convert ReportQualityControlQualityWorksheetInternal to *Many2One.
func (rqq *ReportQualityControlQualityWorksheetInternal) Many2One() *Many2One {
	return NewMany2One(rqq.Id.Get(), "")
}

// CreateReportQualityControlQualityWorksheetInternal creates a new report.quality_control.quality_worksheet_internal model and returns its id.
func (c *Client) CreateReportQualityControlQualityWorksheetInternal(rqq *ReportQualityControlQualityWorksheetInternal) (int64, error) {
	ids, err := c.CreateReportQualityControlQualityWorksheetInternals([]*ReportQualityControlQualityWorksheetInternal{rqq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportQualityControlQualityWorksheetInternal creates a new report.quality_control.quality_worksheet_internal model and returns its id.
func (c *Client) CreateReportQualityControlQualityWorksheetInternals(rqqs []*ReportQualityControlQualityWorksheetInternal) ([]int64, error) {
	var vv []interface{}
	for _, v := range rqqs {
		vv = append(vv, v)
	}
	return c.Create(ReportQualityControlQualityWorksheetInternalModel, vv, nil)
}

// UpdateReportQualityControlQualityWorksheetInternal updates an existing report.quality_control.quality_worksheet_internal record.
func (c *Client) UpdateReportQualityControlQualityWorksheetInternal(rqq *ReportQualityControlQualityWorksheetInternal) error {
	return c.UpdateReportQualityControlQualityWorksheetInternals([]int64{rqq.Id.Get()}, rqq)
}

// UpdateReportQualityControlQualityWorksheetInternals updates existing report.quality_control.quality_worksheet_internal records.
// All records (represented by ids) will be updated by rqq values.
func (c *Client) UpdateReportQualityControlQualityWorksheetInternals(ids []int64, rqq *ReportQualityControlQualityWorksheetInternal) error {
	return c.Update(ReportQualityControlQualityWorksheetInternalModel, ids, rqq, nil)
}

// DeleteReportQualityControlQualityWorksheetInternal deletes an existing report.quality_control.quality_worksheet_internal record.
func (c *Client) DeleteReportQualityControlQualityWorksheetInternal(id int64) error {
	return c.DeleteReportQualityControlQualityWorksheetInternals([]int64{id})
}

// DeleteReportQualityControlQualityWorksheetInternals deletes existing report.quality_control.quality_worksheet_internal records.
func (c *Client) DeleteReportQualityControlQualityWorksheetInternals(ids []int64) error {
	return c.Delete(ReportQualityControlQualityWorksheetInternalModel, ids)
}

// GetReportQualityControlQualityWorksheetInternal gets report.quality_control.quality_worksheet_internal existing record.
func (c *Client) GetReportQualityControlQualityWorksheetInternal(id int64) (*ReportQualityControlQualityWorksheetInternal, error) {
	rqqs, err := c.GetReportQualityControlQualityWorksheetInternals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rqqs)[0]), nil
}

// GetReportQualityControlQualityWorksheetInternals gets report.quality_control.quality_worksheet_internal existing records.
func (c *Client) GetReportQualityControlQualityWorksheetInternals(ids []int64) (*ReportQualityControlQualityWorksheetInternals, error) {
	rqqs := &ReportQualityControlQualityWorksheetInternals{}
	if err := c.Read(ReportQualityControlQualityWorksheetInternalModel, ids, nil, rqqs); err != nil {
		return nil, err
	}
	return rqqs, nil
}

// FindReportQualityControlQualityWorksheetInternal finds report.quality_control.quality_worksheet_internal record by querying it with criteria.
func (c *Client) FindReportQualityControlQualityWorksheetInternal(criteria *Criteria) (*ReportQualityControlQualityWorksheetInternal, error) {
	rqqs := &ReportQualityControlQualityWorksheetInternals{}
	if err := c.SearchRead(ReportQualityControlQualityWorksheetInternalModel, criteria, NewOptions().Limit(1), rqqs); err != nil {
		return nil, err
	}
	return &((*rqqs)[0]), nil
}

// FindReportQualityControlQualityWorksheetInternals finds report.quality_control.quality_worksheet_internal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportQualityControlQualityWorksheetInternals(criteria *Criteria, options *Options) (*ReportQualityControlQualityWorksheetInternals, error) {
	rqqs := &ReportQualityControlQualityWorksheetInternals{}
	if err := c.SearchRead(ReportQualityControlQualityWorksheetInternalModel, criteria, options, rqqs); err != nil {
		return nil, err
	}
	return rqqs, nil
}

// FindReportQualityControlQualityWorksheetInternalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportQualityControlQualityWorksheetInternalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportQualityControlQualityWorksheetInternalModel, criteria, options)
}

// FindReportQualityControlQualityWorksheetInternalId finds record id by querying it with criteria.
func (c *Client) FindReportQualityControlQualityWorksheetInternalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportQualityControlQualityWorksheetInternalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
