package odoo

// MrpReport represents mrp.report model.
type MrpReport struct {
	ByproductCost             *Float    `xmlrpc:"byproduct_cost,omitempty"`
	CompanyId                 *Many2One `xmlrpc:"company_id,omitempty"`
	ComponentCost             *Float    `xmlrpc:"component_cost,omitempty"`
	CurrencyId                *Many2One `xmlrpc:"currency_id,omitempty"`
	DateFinished              *Time     `xmlrpc:"date_finished,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	Duration                  *Float    `xmlrpc:"duration,omitempty"`
	EmployeeCost              *Float    `xmlrpc:"employee_cost,omitempty"`
	ExpectedComponentCostUnit *Float    `xmlrpc:"expected_component_cost_unit,omitempty"`
	ExpectedEmployeeCostUnit  *Float    `xmlrpc:"expected_employee_cost_unit,omitempty"`
	ExpectedOperationCostUnit *Float    `xmlrpc:"expected_operation_cost_unit,omitempty"`
	ExpectedTotalCostUnit     *Float    `xmlrpc:"expected_total_cost_unit,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	OperationCost             *Float    `xmlrpc:"operation_cost,omitempty"`
	ProductId                 *Many2One `xmlrpc:"product_id,omitempty"`
	ProductionId              *Many2One `xmlrpc:"production_id,omitempty"`
	QtyDemanded               *Float    `xmlrpc:"qty_demanded,omitempty"`
	QtyProduced               *Float    `xmlrpc:"qty_produced,omitempty"`
	TotalCost                 *Float    `xmlrpc:"total_cost,omitempty"`
	UnitComponentCost         *Float    `xmlrpc:"unit_component_cost,omitempty"`
	UnitCost                  *Float    `xmlrpc:"unit_cost,omitempty"`
	UnitDuration              *Float    `xmlrpc:"unit_duration,omitempty"`
	UnitEmployeeCost          *Float    `xmlrpc:"unit_employee_cost,omitempty"`
	UnitOperationCost         *Float    `xmlrpc:"unit_operation_cost,omitempty"`
	YieldRate                 *Float    `xmlrpc:"yield_rate,omitempty"`
}

// MrpReports represents array of mrp.report model.
type MrpReports []MrpReport

// MrpReportModel is the odoo model name.
const MrpReportModel = "mrp.report"

// Many2One convert MrpReport to *Many2One.
func (mr *MrpReport) Many2One() *Many2One {
	return NewMany2One(mr.Id.Get(), "")
}

// CreateMrpReport creates a new mrp.report model and returns its id.
func (c *Client) CreateMrpReport(mr *MrpReport) (int64, error) {
	ids, err := c.CreateMrpReports([]*MrpReport{mr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpReport creates a new mrp.report model and returns its id.
func (c *Client) CreateMrpReports(mrs []*MrpReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrs {
		vv = append(vv, v)
	}
	return c.Create(MrpReportModel, vv, nil)
}

// UpdateMrpReport updates an existing mrp.report record.
func (c *Client) UpdateMrpReport(mr *MrpReport) error {
	return c.UpdateMrpReports([]int64{mr.Id.Get()}, mr)
}

// UpdateMrpReports updates existing mrp.report records.
// All records (represented by ids) will be updated by mr values.
func (c *Client) UpdateMrpReports(ids []int64, mr *MrpReport) error {
	return c.Update(MrpReportModel, ids, mr, nil)
}

// DeleteMrpReport deletes an existing mrp.report record.
func (c *Client) DeleteMrpReport(id int64) error {
	return c.DeleteMrpReports([]int64{id})
}

// DeleteMrpReports deletes existing mrp.report records.
func (c *Client) DeleteMrpReports(ids []int64) error {
	return c.Delete(MrpReportModel, ids)
}

// GetMrpReport gets mrp.report existing record.
func (c *Client) GetMrpReport(id int64) (*MrpReport, error) {
	mrs, err := c.GetMrpReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mrs)[0]), nil
}

// GetMrpReports gets mrp.report existing records.
func (c *Client) GetMrpReports(ids []int64) (*MrpReports, error) {
	mrs := &MrpReports{}
	if err := c.Read(MrpReportModel, ids, nil, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMrpReport finds mrp.report record by querying it with criteria.
func (c *Client) FindMrpReport(criteria *Criteria) (*MrpReport, error) {
	mrs := &MrpReports{}
	if err := c.SearchRead(MrpReportModel, criteria, NewOptions().Limit(1), mrs); err != nil {
		return nil, err
	}
	return &((*mrs)[0]), nil
}

// FindMrpReports finds mrp.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpReports(criteria *Criteria, options *Options) (*MrpReports, error) {
	mrs := &MrpReports{}
	if err := c.SearchRead(MrpReportModel, criteria, options, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMrpReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpReportModel, criteria, options)
}

// FindMrpReportId finds record id by querying it with criteria.
func (c *Client) FindMrpReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
