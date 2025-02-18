package odoo

// ReportMrpReportMoOverview represents report.mrp.report_mo_overview model.
type ReportMrpReportMoOverview struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportMrpReportMoOverviews represents array of report.mrp.report_mo_overview model.
type ReportMrpReportMoOverviews []ReportMrpReportMoOverview

// ReportMrpReportMoOverviewModel is the odoo model name.
const ReportMrpReportMoOverviewModel = "report.mrp.report_mo_overview"

// Many2One convert ReportMrpReportMoOverview to *Many2One.
func (rmr *ReportMrpReportMoOverview) Many2One() *Many2One {
	return NewMany2One(rmr.Id.Get(), "")
}

// CreateReportMrpReportMoOverview creates a new report.mrp.report_mo_overview model and returns its id.
func (c *Client) CreateReportMrpReportMoOverview(rmr *ReportMrpReportMoOverview) (int64, error) {
	ids, err := c.CreateReportMrpReportMoOverviews([]*ReportMrpReportMoOverview{rmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportMrpReportMoOverview creates a new report.mrp.report_mo_overview model and returns its id.
func (c *Client) CreateReportMrpReportMoOverviews(rmrs []*ReportMrpReportMoOverview) ([]int64, error) {
	var vv []interface{}
	for _, v := range rmrs {
		vv = append(vv, v)
	}
	return c.Create(ReportMrpReportMoOverviewModel, vv, nil)
}

// UpdateReportMrpReportMoOverview updates an existing report.mrp.report_mo_overview record.
func (c *Client) UpdateReportMrpReportMoOverview(rmr *ReportMrpReportMoOverview) error {
	return c.UpdateReportMrpReportMoOverviews([]int64{rmr.Id.Get()}, rmr)
}

// UpdateReportMrpReportMoOverviews updates existing report.mrp.report_mo_overview records.
// All records (represented by ids) will be updated by rmr values.
func (c *Client) UpdateReportMrpReportMoOverviews(ids []int64, rmr *ReportMrpReportMoOverview) error {
	return c.Update(ReportMrpReportMoOverviewModel, ids, rmr, nil)
}

// DeleteReportMrpReportMoOverview deletes an existing report.mrp.report_mo_overview record.
func (c *Client) DeleteReportMrpReportMoOverview(id int64) error {
	return c.DeleteReportMrpReportMoOverviews([]int64{id})
}

// DeleteReportMrpReportMoOverviews deletes existing report.mrp.report_mo_overview records.
func (c *Client) DeleteReportMrpReportMoOverviews(ids []int64) error {
	return c.Delete(ReportMrpReportMoOverviewModel, ids)
}

// GetReportMrpReportMoOverview gets report.mrp.report_mo_overview existing record.
func (c *Client) GetReportMrpReportMoOverview(id int64) (*ReportMrpReportMoOverview, error) {
	rmrs, err := c.GetReportMrpReportMoOverviews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rmrs)[0]), nil
}

// GetReportMrpReportMoOverviews gets report.mrp.report_mo_overview existing records.
func (c *Client) GetReportMrpReportMoOverviews(ids []int64) (*ReportMrpReportMoOverviews, error) {
	rmrs := &ReportMrpReportMoOverviews{}
	if err := c.Read(ReportMrpReportMoOverviewModel, ids, nil, rmrs); err != nil {
		return nil, err
	}
	return rmrs, nil
}

// FindReportMrpReportMoOverview finds report.mrp.report_mo_overview record by querying it with criteria.
func (c *Client) FindReportMrpReportMoOverview(criteria *Criteria) (*ReportMrpReportMoOverview, error) {
	rmrs := &ReportMrpReportMoOverviews{}
	if err := c.SearchRead(ReportMrpReportMoOverviewModel, criteria, NewOptions().Limit(1), rmrs); err != nil {
		return nil, err
	}
	return &((*rmrs)[0]), nil
}

// FindReportMrpReportMoOverviews finds report.mrp.report_mo_overview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpReportMoOverviews(criteria *Criteria, options *Options) (*ReportMrpReportMoOverviews, error) {
	rmrs := &ReportMrpReportMoOverviews{}
	if err := c.SearchRead(ReportMrpReportMoOverviewModel, criteria, options, rmrs); err != nil {
		return nil, err
	}
	return rmrs, nil
}

// FindReportMrpReportMoOverviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpReportMoOverviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportMrpReportMoOverviewModel, criteria, options)
}

// FindReportMrpReportMoOverviewId finds record id by querying it with criteria.
func (c *Client) FindReportMrpReportMoOverviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportMrpReportMoOverviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
