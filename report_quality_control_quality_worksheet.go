package odoo

// ReportQualityControlQualityWorksheet represents report.quality_control.quality_worksheet model.
type ReportQualityControlQualityWorksheet struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportQualityControlQualityWorksheets represents array of report.quality_control.quality_worksheet model.
type ReportQualityControlQualityWorksheets []ReportQualityControlQualityWorksheet

// ReportQualityControlQualityWorksheetModel is the odoo model name.
const ReportQualityControlQualityWorksheetModel = "report.quality_control.quality_worksheet"

// Many2One convert ReportQualityControlQualityWorksheet to *Many2One.
func (rqq *ReportQualityControlQualityWorksheet) Many2One() *Many2One {
	return NewMany2One(rqq.Id.Get(), "")
}

// CreateReportQualityControlQualityWorksheet creates a new report.quality_control.quality_worksheet model and returns its id.
func (c *Client) CreateReportQualityControlQualityWorksheet(rqq *ReportQualityControlQualityWorksheet) (int64, error) {
	ids, err := c.CreateReportQualityControlQualityWorksheets([]*ReportQualityControlQualityWorksheet{rqq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportQualityControlQualityWorksheet creates a new report.quality_control.quality_worksheet model and returns its id.
func (c *Client) CreateReportQualityControlQualityWorksheets(rqqs []*ReportQualityControlQualityWorksheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range rqqs {
		vv = append(vv, v)
	}
	return c.Create(ReportQualityControlQualityWorksheetModel, vv, nil)
}

// UpdateReportQualityControlQualityWorksheet updates an existing report.quality_control.quality_worksheet record.
func (c *Client) UpdateReportQualityControlQualityWorksheet(rqq *ReportQualityControlQualityWorksheet) error {
	return c.UpdateReportQualityControlQualityWorksheets([]int64{rqq.Id.Get()}, rqq)
}

// UpdateReportQualityControlQualityWorksheets updates existing report.quality_control.quality_worksheet records.
// All records (represented by ids) will be updated by rqq values.
func (c *Client) UpdateReportQualityControlQualityWorksheets(ids []int64, rqq *ReportQualityControlQualityWorksheet) error {
	return c.Update(ReportQualityControlQualityWorksheetModel, ids, rqq, nil)
}

// DeleteReportQualityControlQualityWorksheet deletes an existing report.quality_control.quality_worksheet record.
func (c *Client) DeleteReportQualityControlQualityWorksheet(id int64) error {
	return c.DeleteReportQualityControlQualityWorksheets([]int64{id})
}

// DeleteReportQualityControlQualityWorksheets deletes existing report.quality_control.quality_worksheet records.
func (c *Client) DeleteReportQualityControlQualityWorksheets(ids []int64) error {
	return c.Delete(ReportQualityControlQualityWorksheetModel, ids)
}

// GetReportQualityControlQualityWorksheet gets report.quality_control.quality_worksheet existing record.
func (c *Client) GetReportQualityControlQualityWorksheet(id int64) (*ReportQualityControlQualityWorksheet, error) {
	rqqs, err := c.GetReportQualityControlQualityWorksheets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rqqs)[0]), nil
}

// GetReportQualityControlQualityWorksheets gets report.quality_control.quality_worksheet existing records.
func (c *Client) GetReportQualityControlQualityWorksheets(ids []int64) (*ReportQualityControlQualityWorksheets, error) {
	rqqs := &ReportQualityControlQualityWorksheets{}
	if err := c.Read(ReportQualityControlQualityWorksheetModel, ids, nil, rqqs); err != nil {
		return nil, err
	}
	return rqqs, nil
}

// FindReportQualityControlQualityWorksheet finds report.quality_control.quality_worksheet record by querying it with criteria.
func (c *Client) FindReportQualityControlQualityWorksheet(criteria *Criteria) (*ReportQualityControlQualityWorksheet, error) {
	rqqs := &ReportQualityControlQualityWorksheets{}
	if err := c.SearchRead(ReportQualityControlQualityWorksheetModel, criteria, NewOptions().Limit(1), rqqs); err != nil {
		return nil, err
	}
	return &((*rqqs)[0]), nil
}

// FindReportQualityControlQualityWorksheets finds report.quality_control.quality_worksheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportQualityControlQualityWorksheets(criteria *Criteria, options *Options) (*ReportQualityControlQualityWorksheets, error) {
	rqqs := &ReportQualityControlQualityWorksheets{}
	if err := c.SearchRead(ReportQualityControlQualityWorksheetModel, criteria, options, rqqs); err != nil {
		return nil, err
	}
	return rqqs, nil
}

// FindReportQualityControlQualityWorksheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportQualityControlQualityWorksheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportQualityControlQualityWorksheetModel, criteria, options)
}

// FindReportQualityControlQualityWorksheetId finds record id by querying it with criteria.
func (c *Client) FindReportQualityControlQualityWorksheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportQualityControlQualityWorksheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
