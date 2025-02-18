package odoo

// ReportMrpAccountEnterpriseMrpCostStructure represents report.mrp_account_enterprise.mrp_cost_structure model.
type ReportMrpAccountEnterpriseMrpCostStructure struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportMrpAccountEnterpriseMrpCostStructures represents array of report.mrp_account_enterprise.mrp_cost_structure model.
type ReportMrpAccountEnterpriseMrpCostStructures []ReportMrpAccountEnterpriseMrpCostStructure

// ReportMrpAccountEnterpriseMrpCostStructureModel is the odoo model name.
const ReportMrpAccountEnterpriseMrpCostStructureModel = "report.mrp_account_enterprise.mrp_cost_structure"

// Many2One convert ReportMrpAccountEnterpriseMrpCostStructure to *Many2One.
func (rmm *ReportMrpAccountEnterpriseMrpCostStructure) Many2One() *Many2One {
	return NewMany2One(rmm.Id.Get(), "")
}

// CreateReportMrpAccountEnterpriseMrpCostStructure creates a new report.mrp_account_enterprise.mrp_cost_structure model and returns its id.
func (c *Client) CreateReportMrpAccountEnterpriseMrpCostStructure(rmm *ReportMrpAccountEnterpriseMrpCostStructure) (int64, error) {
	ids, err := c.CreateReportMrpAccountEnterpriseMrpCostStructures([]*ReportMrpAccountEnterpriseMrpCostStructure{rmm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportMrpAccountEnterpriseMrpCostStructure creates a new report.mrp_account_enterprise.mrp_cost_structure model and returns its id.
func (c *Client) CreateReportMrpAccountEnterpriseMrpCostStructures(rmms []*ReportMrpAccountEnterpriseMrpCostStructure) ([]int64, error) {
	var vv []interface{}
	for _, v := range rmms {
		vv = append(vv, v)
	}
	return c.Create(ReportMrpAccountEnterpriseMrpCostStructureModel, vv, nil)
}

// UpdateReportMrpAccountEnterpriseMrpCostStructure updates an existing report.mrp_account_enterprise.mrp_cost_structure record.
func (c *Client) UpdateReportMrpAccountEnterpriseMrpCostStructure(rmm *ReportMrpAccountEnterpriseMrpCostStructure) error {
	return c.UpdateReportMrpAccountEnterpriseMrpCostStructures([]int64{rmm.Id.Get()}, rmm)
}

// UpdateReportMrpAccountEnterpriseMrpCostStructures updates existing report.mrp_account_enterprise.mrp_cost_structure records.
// All records (represented by ids) will be updated by rmm values.
func (c *Client) UpdateReportMrpAccountEnterpriseMrpCostStructures(ids []int64, rmm *ReportMrpAccountEnterpriseMrpCostStructure) error {
	return c.Update(ReportMrpAccountEnterpriseMrpCostStructureModel, ids, rmm, nil)
}

// DeleteReportMrpAccountEnterpriseMrpCostStructure deletes an existing report.mrp_account_enterprise.mrp_cost_structure record.
func (c *Client) DeleteReportMrpAccountEnterpriseMrpCostStructure(id int64) error {
	return c.DeleteReportMrpAccountEnterpriseMrpCostStructures([]int64{id})
}

// DeleteReportMrpAccountEnterpriseMrpCostStructures deletes existing report.mrp_account_enterprise.mrp_cost_structure records.
func (c *Client) DeleteReportMrpAccountEnterpriseMrpCostStructures(ids []int64) error {
	return c.Delete(ReportMrpAccountEnterpriseMrpCostStructureModel, ids)
}

// GetReportMrpAccountEnterpriseMrpCostStructure gets report.mrp_account_enterprise.mrp_cost_structure existing record.
func (c *Client) GetReportMrpAccountEnterpriseMrpCostStructure(id int64) (*ReportMrpAccountEnterpriseMrpCostStructure, error) {
	rmms, err := c.GetReportMrpAccountEnterpriseMrpCostStructures([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rmms)[0]), nil
}

// GetReportMrpAccountEnterpriseMrpCostStructures gets report.mrp_account_enterprise.mrp_cost_structure existing records.
func (c *Client) GetReportMrpAccountEnterpriseMrpCostStructures(ids []int64) (*ReportMrpAccountEnterpriseMrpCostStructures, error) {
	rmms := &ReportMrpAccountEnterpriseMrpCostStructures{}
	if err := c.Read(ReportMrpAccountEnterpriseMrpCostStructureModel, ids, nil, rmms); err != nil {
		return nil, err
	}
	return rmms, nil
}

// FindReportMrpAccountEnterpriseMrpCostStructure finds report.mrp_account_enterprise.mrp_cost_structure record by querying it with criteria.
func (c *Client) FindReportMrpAccountEnterpriseMrpCostStructure(criteria *Criteria) (*ReportMrpAccountEnterpriseMrpCostStructure, error) {
	rmms := &ReportMrpAccountEnterpriseMrpCostStructures{}
	if err := c.SearchRead(ReportMrpAccountEnterpriseMrpCostStructureModel, criteria, NewOptions().Limit(1), rmms); err != nil {
		return nil, err
	}
	return &((*rmms)[0]), nil
}

// FindReportMrpAccountEnterpriseMrpCostStructures finds report.mrp_account_enterprise.mrp_cost_structure records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpAccountEnterpriseMrpCostStructures(criteria *Criteria, options *Options) (*ReportMrpAccountEnterpriseMrpCostStructures, error) {
	rmms := &ReportMrpAccountEnterpriseMrpCostStructures{}
	if err := c.SearchRead(ReportMrpAccountEnterpriseMrpCostStructureModel, criteria, options, rmms); err != nil {
		return nil, err
	}
	return rmms, nil
}

// FindReportMrpAccountEnterpriseMrpCostStructureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpAccountEnterpriseMrpCostStructureIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportMrpAccountEnterpriseMrpCostStructureModel, criteria, options)
}

// FindReportMrpAccountEnterpriseMrpCostStructureId finds record id by querying it with criteria.
func (c *Client) FindReportMrpAccountEnterpriseMrpCostStructureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportMrpAccountEnterpriseMrpCostStructureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
