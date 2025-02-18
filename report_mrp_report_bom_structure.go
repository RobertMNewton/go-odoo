package odoo

// ReportMrpReportBomStructure represents report.mrp.report_bom_structure model.
type ReportMrpReportBomStructure struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportMrpReportBomStructures represents array of report.mrp.report_bom_structure model.
type ReportMrpReportBomStructures []ReportMrpReportBomStructure

// ReportMrpReportBomStructureModel is the odoo model name.
const ReportMrpReportBomStructureModel = "report.mrp.report_bom_structure"

// Many2One convert ReportMrpReportBomStructure to *Many2One.
func (rmr *ReportMrpReportBomStructure) Many2One() *Many2One {
	return NewMany2One(rmr.Id.Get(), "")
}

// CreateReportMrpReportBomStructure creates a new report.mrp.report_bom_structure model and returns its id.
func (c *Client) CreateReportMrpReportBomStructure(rmr *ReportMrpReportBomStructure) (int64, error) {
	ids, err := c.CreateReportMrpReportBomStructures([]*ReportMrpReportBomStructure{rmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportMrpReportBomStructure creates a new report.mrp.report_bom_structure model and returns its id.
func (c *Client) CreateReportMrpReportBomStructures(rmrs []*ReportMrpReportBomStructure) ([]int64, error) {
	var vv []interface{}
	for _, v := range rmrs {
		vv = append(vv, v)
	}
	return c.Create(ReportMrpReportBomStructureModel, vv, nil)
}

// UpdateReportMrpReportBomStructure updates an existing report.mrp.report_bom_structure record.
func (c *Client) UpdateReportMrpReportBomStructure(rmr *ReportMrpReportBomStructure) error {
	return c.UpdateReportMrpReportBomStructures([]int64{rmr.Id.Get()}, rmr)
}

// UpdateReportMrpReportBomStructures updates existing report.mrp.report_bom_structure records.
// All records (represented by ids) will be updated by rmr values.
func (c *Client) UpdateReportMrpReportBomStructures(ids []int64, rmr *ReportMrpReportBomStructure) error {
	return c.Update(ReportMrpReportBomStructureModel, ids, rmr, nil)
}

// DeleteReportMrpReportBomStructure deletes an existing report.mrp.report_bom_structure record.
func (c *Client) DeleteReportMrpReportBomStructure(id int64) error {
	return c.DeleteReportMrpReportBomStructures([]int64{id})
}

// DeleteReportMrpReportBomStructures deletes existing report.mrp.report_bom_structure records.
func (c *Client) DeleteReportMrpReportBomStructures(ids []int64) error {
	return c.Delete(ReportMrpReportBomStructureModel, ids)
}

// GetReportMrpReportBomStructure gets report.mrp.report_bom_structure existing record.
func (c *Client) GetReportMrpReportBomStructure(id int64) (*ReportMrpReportBomStructure, error) {
	rmrs, err := c.GetReportMrpReportBomStructures([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rmrs)[0]), nil
}

// GetReportMrpReportBomStructures gets report.mrp.report_bom_structure existing records.
func (c *Client) GetReportMrpReportBomStructures(ids []int64) (*ReportMrpReportBomStructures, error) {
	rmrs := &ReportMrpReportBomStructures{}
	if err := c.Read(ReportMrpReportBomStructureModel, ids, nil, rmrs); err != nil {
		return nil, err
	}
	return rmrs, nil
}

// FindReportMrpReportBomStructure finds report.mrp.report_bom_structure record by querying it with criteria.
func (c *Client) FindReportMrpReportBomStructure(criteria *Criteria) (*ReportMrpReportBomStructure, error) {
	rmrs := &ReportMrpReportBomStructures{}
	if err := c.SearchRead(ReportMrpReportBomStructureModel, criteria, NewOptions().Limit(1), rmrs); err != nil {
		return nil, err
	}
	return &((*rmrs)[0]), nil
}

// FindReportMrpReportBomStructures finds report.mrp.report_bom_structure records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpReportBomStructures(criteria *Criteria, options *Options) (*ReportMrpReportBomStructures, error) {
	rmrs := &ReportMrpReportBomStructures{}
	if err := c.SearchRead(ReportMrpReportBomStructureModel, criteria, options, rmrs); err != nil {
		return nil, err
	}
	return rmrs, nil
}

// FindReportMrpReportBomStructureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpReportBomStructureIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportMrpReportBomStructureModel, criteria, options)
}

// FindReportMrpReportBomStructureId finds record id by querying it with criteria.
func (c *Client) FindReportMrpReportBomStructureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportMrpReportBomStructureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
