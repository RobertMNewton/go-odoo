package odoo

// MrpWorkorder represents mrp.workorder model.
type MrpWorkorder struct {
	BarcodeScanned                 *String    `xmlrpc:"_barcode_scanned,omitempty"`
	AllEmployeesAllowed            *Bool      `xmlrpc:"all_employees_allowed,omitempty"`
	AllowProducingQuantityChange   *Bool      `xmlrpc:"allow_producing_quantity_change,omitempty"`
	AllowWorkorderDependencies     *Bool      `xmlrpc:"allow_workorder_dependencies,omitempty"`
	AllowedEmployees               *Relation  `xmlrpc:"allowed_employees,omitempty"`
	Barcode                        *String    `xmlrpc:"barcode,omitempty"`
	BlockedByWorkorderIds          *Relation  `xmlrpc:"blocked_by_workorder_ids,omitempty"`
	CheckIds                       *Relation  `xmlrpc:"check_ids,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty"`
	ConnectedEmployeeIds           *Relation  `xmlrpc:"connected_employee_ids,omitempty"`
	Consumption                    *Selection `xmlrpc:"consumption,omitempty"`
	CostsHour                      *Float     `xmlrpc:"costs_hour,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentQualityCheckId          *Many2One  `xmlrpc:"current_quality_check_id,omitempty"`
	DateFinished                   *Time      `xmlrpc:"date_finished,omitempty"`
	DateStart                      *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	DoneCheckIds                   *Relation  `xmlrpc:"done_check_ids,omitempty"`
	Duration                       *Float     `xmlrpc:"duration,omitempty"`
	DurationExpected               *Float     `xmlrpc:"duration_expected,omitempty"`
	DurationPercent                *Int       `xmlrpc:"duration_percent,omitempty"`
	DurationUnit                   *Float     `xmlrpc:"duration_unit,omitempty"`
	EmployeeAnalyticAccountLineIds *Relation  `xmlrpc:"employee_analytic_account_line_ids,omitempty"`
	EmployeeAssignedIds            *Relation  `xmlrpc:"employee_assigned_ids,omitempty"`
	EmployeeId                     *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeIds                    *Relation  `xmlrpc:"employee_ids,omitempty"`
	EmployeeName                   *String    `xmlrpc:"employee_name,omitempty"`
	FinishedLotId                  *Many2One  `xmlrpc:"finished_lot_id,omitempty"`
	FinishedProductCheckIds        *Relation  `xmlrpc:"finished_product_check_ids,omitempty"`
	HasOperationNote               *Bool      `xmlrpc:"has_operation_note,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	IsFirstStartedWo               *Bool      `xmlrpc:"is_first_started_wo,omitempty"`
	IsLastLot                      *Bool      `xmlrpc:"is_last_lot,omitempty"`
	IsLastUnfinishedWo             *Bool      `xmlrpc:"is_last_unfinished_wo,omitempty"`
	IsPlanned                      *Bool      `xmlrpc:"is_planned,omitempty"`
	IsProduced                     *Bool      `xmlrpc:"is_produced,omitempty"`
	IsUserWorking                  *Bool      `xmlrpc:"is_user_working,omitempty"`
	JsonPopover                    *String    `xmlrpc:"json_popover,omitempty"`
	LastWorkingUserId              *Many2One  `xmlrpc:"last_working_user_id,omitempty"`
	LeaveId                        *Many2One  `xmlrpc:"leave_id,omitempty"`
	LotId                          *Many2One  `xmlrpc:"lot_id,omitempty"`
	MoAnalyticAccountLineIds       *Relation  `xmlrpc:"mo_analytic_account_line_ids,omitempty"`
	MoveFinishedIds                *Relation  `xmlrpc:"move_finished_ids,omitempty"`
	MoveId                         *Many2One  `xmlrpc:"move_id,omitempty"`
	MoveLineIds                    *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveRawIds                     *Relation  `xmlrpc:"move_raw_ids,omitempty"`
	Name                           *String    `xmlrpc:"name,omitempty"`
	NeededByWorkorderIds           *Relation  `xmlrpc:"needed_by_workorder_ids,omitempty"`
	OperationId                    *Many2One  `xmlrpc:"operation_id,omitempty"`
	OperationNote                  *String    `xmlrpc:"operation_note,omitempty"`
	Picture                        *String    `xmlrpc:"picture,omitempty"`
	ProductDescriptionVariants     *String    `xmlrpc:"product_description_variants,omitempty"`
	ProductId                      *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTracking                *Selection `xmlrpc:"product_tracking,omitempty"`
	ProductUomId                   *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductionAvailability         *Selection `xmlrpc:"production_availability,omitempty"`
	ProductionBomId                *Many2One  `xmlrpc:"production_bom_id,omitempty"`
	ProductionDate                 *Time      `xmlrpc:"production_date,omitempty"`
	ProductionId                   *Many2One  `xmlrpc:"production_id,omitempty"`
	ProductionState                *Selection `xmlrpc:"production_state,omitempty"`
	Progress                       *Float     `xmlrpc:"progress,omitempty"`
	QtyProduced                    *Float     `xmlrpc:"qty_produced,omitempty"`
	QtyProducing                   *Float     `xmlrpc:"qty_producing,omitempty"`
	QtyProduction                  *Float     `xmlrpc:"qty_production,omitempty"`
	QtyReady                       *Float     `xmlrpc:"qty_ready,omitempty"`
	QtyRemaining                   *Float     `xmlrpc:"qty_remaining,omitempty"`
	QtyReportedFromPreviousWo      *Float     `xmlrpc:"qty_reported_from_previous_wo,omitempty"`
	QualityAlertCount              *Int       `xmlrpc:"quality_alert_count,omitempty"`
	QualityAlertIds                *Relation  `xmlrpc:"quality_alert_ids,omitempty"`
	QualityCheckFail               *Bool      `xmlrpc:"quality_check_fail,omitempty"`
	QualityCheckTodo               *Bool      `xmlrpc:"quality_check_todo,omitempty"`
	QualityPointCount              *Int       `xmlrpc:"quality_point_count,omitempty"`
	QualityPointIds                *Relation  `xmlrpc:"quality_point_ids,omitempty"`
	QualityState                   *Selection `xmlrpc:"quality_state,omitempty"`
	ScrapCount                     *Int       `xmlrpc:"scrap_count,omitempty"`
	ScrapIds                       *Relation  `xmlrpc:"scrap_ids,omitempty"`
	Sequence                       *Int       `xmlrpc:"sequence,omitempty"`
	ShowJsonPopover                *Bool      `xmlrpc:"show_json_popover,omitempty"`
	State                          *Selection `xmlrpc:"state,omitempty"`
	TestType                       *String    `xmlrpc:"test_type,omitempty"`
	TestTypeId                     *Many2One  `xmlrpc:"test_type_id,omitempty"`
	TimeIds                        *Relation  `xmlrpc:"time_ids,omitempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omitempty"`
	WcAnalyticAccountLineIds       *Relation  `xmlrpc:"wc_analytic_account_line_ids,omitempty"`
	WorkcenterId                   *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkingState                   *Selection `xmlrpc:"working_state,omitempty"`
	WorkingUserIds                 *Relation  `xmlrpc:"working_user_ids,omitempty"`
	Worksheet                      *String    `xmlrpc:"worksheet,omitempty"`
	WorksheetGoogleSlide           *String    `xmlrpc:"worksheet_google_slide,omitempty"`
	WorksheetPage                  *Int       `xmlrpc:"worksheet_page,omitempty"`
	WorksheetType                  *Selection `xmlrpc:"worksheet_type,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkorders represents array of mrp.workorder model.
type MrpWorkorders []MrpWorkorder

// MrpWorkorderModel is the odoo model name.
const MrpWorkorderModel = "mrp.workorder"

// Many2One convert MrpWorkorder to *Many2One.
func (mw *MrpWorkorder) Many2One() *Many2One {
	return NewMany2One(mw.Id.Get(), "")
}

// CreateMrpWorkorder creates a new mrp.workorder model and returns its id.
func (c *Client) CreateMrpWorkorder(mw *MrpWorkorder) (int64, error) {
	ids, err := c.CreateMrpWorkorders([]*MrpWorkorder{mw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkorder creates a new mrp.workorder model and returns its id.
func (c *Client) CreateMrpWorkorders(mws []*MrpWorkorder) ([]int64, error) {
	var vv []interface{}
	for _, v := range mws {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkorderModel, vv, nil)
}

// UpdateMrpWorkorder updates an existing mrp.workorder record.
func (c *Client) UpdateMrpWorkorder(mw *MrpWorkorder) error {
	return c.UpdateMrpWorkorders([]int64{mw.Id.Get()}, mw)
}

// UpdateMrpWorkorders updates existing mrp.workorder records.
// All records (represented by ids) will be updated by mw values.
func (c *Client) UpdateMrpWorkorders(ids []int64, mw *MrpWorkorder) error {
	return c.Update(MrpWorkorderModel, ids, mw, nil)
}

// DeleteMrpWorkorder deletes an existing mrp.workorder record.
func (c *Client) DeleteMrpWorkorder(id int64) error {
	return c.DeleteMrpWorkorders([]int64{id})
}

// DeleteMrpWorkorders deletes existing mrp.workorder records.
func (c *Client) DeleteMrpWorkorders(ids []int64) error {
	return c.Delete(MrpWorkorderModel, ids)
}

// GetMrpWorkorder gets mrp.workorder existing record.
func (c *Client) GetMrpWorkorder(id int64) (*MrpWorkorder, error) {
	mws, err := c.GetMrpWorkorders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// GetMrpWorkorders gets mrp.workorder existing records.
func (c *Client) GetMrpWorkorders(ids []int64) (*MrpWorkorders, error) {
	mws := &MrpWorkorders{}
	if err := c.Read(MrpWorkorderModel, ids, nil, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkorder finds mrp.workorder record by querying it with criteria.
func (c *Client) FindMrpWorkorder(criteria *Criteria) (*MrpWorkorder, error) {
	mws := &MrpWorkorders{}
	if err := c.SearchRead(MrpWorkorderModel, criteria, NewOptions().Limit(1), mws); err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// FindMrpWorkorders finds mrp.workorder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkorders(criteria *Criteria, options *Options) (*MrpWorkorders, error) {
	mws := &MrpWorkorders{}
	if err := c.SearchRead(MrpWorkorderModel, criteria, options, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkorderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkorderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkorderModel, criteria, options)
}

// FindMrpWorkorderId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkorderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkorderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
