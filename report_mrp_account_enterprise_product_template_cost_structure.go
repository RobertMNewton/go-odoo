package odoo

// ReportMrpAccountEnterpriseProductTemplateCostStructure represents report.mrp_account_enterprise.product_template_cost_structure model.
type ReportMrpAccountEnterpriseProductTemplateCostStructure struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ReportMrpAccountEnterpriseProductTemplateCostStructures represents array of report.mrp_account_enterprise.product_template_cost_structure model.
type ReportMrpAccountEnterpriseProductTemplateCostStructures []ReportMrpAccountEnterpriseProductTemplateCostStructure

// ReportMrpAccountEnterpriseProductTemplateCostStructureModel is the odoo model name.
const ReportMrpAccountEnterpriseProductTemplateCostStructureModel = "report.mrp_account_enterprise.product_template_cost_structure"

// Many2One convert ReportMrpAccountEnterpriseProductTemplateCostStructure to *Many2One.
func (rmp *ReportMrpAccountEnterpriseProductTemplateCostStructure) Many2One() *Many2One {
	return NewMany2One(rmp.Id.Get(), "")
}

// CreateReportMrpAccountEnterpriseProductTemplateCostStructure creates a new report.mrp_account_enterprise.product_template_cost_structure model and returns its id.
func (c *Client) CreateReportMrpAccountEnterpriseProductTemplateCostStructure(rmp *ReportMrpAccountEnterpriseProductTemplateCostStructure) (int64, error) {
	ids, err := c.CreateReportMrpAccountEnterpriseProductTemplateCostStructures([]*ReportMrpAccountEnterpriseProductTemplateCostStructure{rmp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportMrpAccountEnterpriseProductTemplateCostStructure creates a new report.mrp_account_enterprise.product_template_cost_structure model and returns its id.
func (c *Client) CreateReportMrpAccountEnterpriseProductTemplateCostStructures(rmps []*ReportMrpAccountEnterpriseProductTemplateCostStructure) ([]int64, error) {
	var vv []interface{}
	for _, v := range rmps {
		vv = append(vv, v)
	}
	return c.Create(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, vv, nil)
}

// UpdateReportMrpAccountEnterpriseProductTemplateCostStructure updates an existing report.mrp_account_enterprise.product_template_cost_structure record.
func (c *Client) UpdateReportMrpAccountEnterpriseProductTemplateCostStructure(rmp *ReportMrpAccountEnterpriseProductTemplateCostStructure) error {
	return c.UpdateReportMrpAccountEnterpriseProductTemplateCostStructures([]int64{rmp.Id.Get()}, rmp)
}

// UpdateReportMrpAccountEnterpriseProductTemplateCostStructures updates existing report.mrp_account_enterprise.product_template_cost_structure records.
// All records (represented by ids) will be updated by rmp values.
func (c *Client) UpdateReportMrpAccountEnterpriseProductTemplateCostStructures(ids []int64, rmp *ReportMrpAccountEnterpriseProductTemplateCostStructure) error {
	return c.Update(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, ids, rmp, nil)
}

// DeleteReportMrpAccountEnterpriseProductTemplateCostStructure deletes an existing report.mrp_account_enterprise.product_template_cost_structure record.
func (c *Client) DeleteReportMrpAccountEnterpriseProductTemplateCostStructure(id int64) error {
	return c.DeleteReportMrpAccountEnterpriseProductTemplateCostStructures([]int64{id})
}

// DeleteReportMrpAccountEnterpriseProductTemplateCostStructures deletes existing report.mrp_account_enterprise.product_template_cost_structure records.
func (c *Client) DeleteReportMrpAccountEnterpriseProductTemplateCostStructures(ids []int64) error {
	return c.Delete(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, ids)
}

// GetReportMrpAccountEnterpriseProductTemplateCostStructure gets report.mrp_account_enterprise.product_template_cost_structure existing record.
func (c *Client) GetReportMrpAccountEnterpriseProductTemplateCostStructure(id int64) (*ReportMrpAccountEnterpriseProductTemplateCostStructure, error) {
	rmps, err := c.GetReportMrpAccountEnterpriseProductTemplateCostStructures([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rmps)[0]), nil
}

// GetReportMrpAccountEnterpriseProductTemplateCostStructures gets report.mrp_account_enterprise.product_template_cost_structure existing records.
func (c *Client) GetReportMrpAccountEnterpriseProductTemplateCostStructures(ids []int64) (*ReportMrpAccountEnterpriseProductTemplateCostStructures, error) {
	rmps := &ReportMrpAccountEnterpriseProductTemplateCostStructures{}
	if err := c.Read(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, ids, nil, rmps); err != nil {
		return nil, err
	}
	return rmps, nil
}

// FindReportMrpAccountEnterpriseProductTemplateCostStructure finds report.mrp_account_enterprise.product_template_cost_structure record by querying it with criteria.
func (c *Client) FindReportMrpAccountEnterpriseProductTemplateCostStructure(criteria *Criteria) (*ReportMrpAccountEnterpriseProductTemplateCostStructure, error) {
	rmps := &ReportMrpAccountEnterpriseProductTemplateCostStructures{}
	if err := c.SearchRead(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, criteria, NewOptions().Limit(1), rmps); err != nil {
		return nil, err
	}
	return &((*rmps)[0]), nil
}

// FindReportMrpAccountEnterpriseProductTemplateCostStructures finds report.mrp_account_enterprise.product_template_cost_structure records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpAccountEnterpriseProductTemplateCostStructures(criteria *Criteria, options *Options) (*ReportMrpAccountEnterpriseProductTemplateCostStructures, error) {
	rmps := &ReportMrpAccountEnterpriseProductTemplateCostStructures{}
	if err := c.SearchRead(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, criteria, options, rmps); err != nil {
		return nil, err
	}
	return rmps, nil
}

// FindReportMrpAccountEnterpriseProductTemplateCostStructureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportMrpAccountEnterpriseProductTemplateCostStructureIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, criteria, options)
}

// FindReportMrpAccountEnterpriseProductTemplateCostStructureId finds record id by querying it with criteria.
func (c *Client) FindReportMrpAccountEnterpriseProductTemplateCostStructureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportMrpAccountEnterpriseProductTemplateCostStructureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
