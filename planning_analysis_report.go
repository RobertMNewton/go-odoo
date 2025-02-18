package odoo

// PlanningAnalysisReport represents planning.analysis.report model.
type PlanningAnalysisReport struct {
	AllocatedHours      *Float     `xmlrpc:"allocated_hours,omitempty"`
	AllocatedPercentage *Float     `xmlrpc:"allocated_percentage,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	DepartmentId        *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId          *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeSkillIds    *Relation  `xmlrpc:"employee_skill_ids,omitempty"`
	EndDatetime         *Time      `xmlrpc:"end_datetime,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	JobTitle            *String    `xmlrpc:"job_title,omitempty"`
	ManagerId           *Many2One  `xmlrpc:"manager_id,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	PublicationWarning  *Bool      `xmlrpc:"publication_warning,omitempty"`
	RecurrencyId        *Many2One  `xmlrpc:"recurrency_id,omitempty"`
	RequestToSwitch     *Bool      `xmlrpc:"request_to_switch,omitempty"`
	ResourceId          *Many2One  `xmlrpc:"resource_id,omitempty"`
	ResourceType        *Selection `xmlrpc:"resource_type,omitempty"`
	RoleId              *Many2One  `xmlrpc:"role_id,omitempty"`
	RoleProductIds      *Relation  `xmlrpc:"role_product_ids,omitempty"`
	SaleLineId          *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId         *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	SlotId              *Many2One  `xmlrpc:"slot_id,omitempty"`
	StartDatetime       *Time      `xmlrpc:"start_datetime,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
}

// PlanningAnalysisReports represents array of planning.analysis.report model.
type PlanningAnalysisReports []PlanningAnalysisReport

// PlanningAnalysisReportModel is the odoo model name.
const PlanningAnalysisReportModel = "planning.analysis.report"

// Many2One convert PlanningAnalysisReport to *Many2One.
func (par *PlanningAnalysisReport) Many2One() *Many2One {
	return NewMany2One(par.Id.Get(), "")
}

// CreatePlanningAnalysisReport creates a new planning.analysis.report model and returns its id.
func (c *Client) CreatePlanningAnalysisReport(par *PlanningAnalysisReport) (int64, error) {
	ids, err := c.CreatePlanningAnalysisReports([]*PlanningAnalysisReport{par})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningAnalysisReport creates a new planning.analysis.report model and returns its id.
func (c *Client) CreatePlanningAnalysisReports(pars []*PlanningAnalysisReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range pars {
		vv = append(vv, v)
	}
	return c.Create(PlanningAnalysisReportModel, vv, nil)
}

// UpdatePlanningAnalysisReport updates an existing planning.analysis.report record.
func (c *Client) UpdatePlanningAnalysisReport(par *PlanningAnalysisReport) error {
	return c.UpdatePlanningAnalysisReports([]int64{par.Id.Get()}, par)
}

// UpdatePlanningAnalysisReports updates existing planning.analysis.report records.
// All records (represented by ids) will be updated by par values.
func (c *Client) UpdatePlanningAnalysisReports(ids []int64, par *PlanningAnalysisReport) error {
	return c.Update(PlanningAnalysisReportModel, ids, par, nil)
}

// DeletePlanningAnalysisReport deletes an existing planning.analysis.report record.
func (c *Client) DeletePlanningAnalysisReport(id int64) error {
	return c.DeletePlanningAnalysisReports([]int64{id})
}

// DeletePlanningAnalysisReports deletes existing planning.analysis.report records.
func (c *Client) DeletePlanningAnalysisReports(ids []int64) error {
	return c.Delete(PlanningAnalysisReportModel, ids)
}

// GetPlanningAnalysisReport gets planning.analysis.report existing record.
func (c *Client) GetPlanningAnalysisReport(id int64) (*PlanningAnalysisReport, error) {
	pars, err := c.GetPlanningAnalysisReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pars)[0]), nil
}

// GetPlanningAnalysisReports gets planning.analysis.report existing records.
func (c *Client) GetPlanningAnalysisReports(ids []int64) (*PlanningAnalysisReports, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.Read(PlanningAnalysisReportModel, ids, nil, pars); err != nil {
		return nil, err
	}
	return pars, nil
}

// FindPlanningAnalysisReport finds planning.analysis.report record by querying it with criteria.
func (c *Client) FindPlanningAnalysisReport(criteria *Criteria) (*PlanningAnalysisReport, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.SearchRead(PlanningAnalysisReportModel, criteria, NewOptions().Limit(1), pars); err != nil {
		return nil, err
	}
	return &((*pars)[0]), nil
}

// FindPlanningAnalysisReports finds planning.analysis.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningAnalysisReports(criteria *Criteria, options *Options) (*PlanningAnalysisReports, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.SearchRead(PlanningAnalysisReportModel, criteria, options, pars); err != nil {
		return nil, err
	}
	return pars, nil
}

// FindPlanningAnalysisReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningAnalysisReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningAnalysisReportModel, criteria, options)
}

// FindPlanningAnalysisReportId finds record id by querying it with criteria.
func (c *Client) FindPlanningAnalysisReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningAnalysisReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
