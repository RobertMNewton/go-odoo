package odoo

// QualityCheck represents quality.check model.
type QualityCheck struct {
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AdditionalNote              *String    `xmlrpc:"additional_note,omitempty"`
	AlertCount                  *Int       `xmlrpc:"alert_count,omitempty"`
	AlertIds                    *Relation  `xmlrpc:"alert_ids,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ComponentBarcode            *String    `xmlrpc:"component_barcode,omitempty"`
	ComponentId                 *Many2One  `xmlrpc:"component_id,omitempty"`
	ComponentTracking           *Selection `xmlrpc:"component_tracking,omitempty"`
	ComponentUomId              *Many2One  `xmlrpc:"component_uom_id,omitempty"`
	Consumption                 *Selection `xmlrpc:"consumption,omitempty"`
	ControlDate                 *Time      `xmlrpc:"control_date,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omitempty"`
	FailureLocationId           *Many2One  `xmlrpc:"failure_location_id,omitempty"`
	FailureMessage              *String    `xmlrpc:"failure_message,omitempty"`
	FinishedLotId               *Many2One  `xmlrpc:"finished_lot_id,omitempty"`
	FinishedProductSequence     *Float     `xmlrpc:"finished_product_sequence,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsDeleted                   *Bool      `xmlrpc:"is_deleted,omitempty"`
	IsLotTestedFractionally     *Bool      `xmlrpc:"is_lot_tested_fractionally,omitempty"`
	IsUserWorking               *Bool      `xmlrpc:"is_user_working,omitempty"`
	LotId                       *Many2One  `xmlrpc:"lot_id,omitempty"`
	LotLineId                   *Many2One  `xmlrpc:"lot_line_id,omitempty"`
	LotName                     *String    `xmlrpc:"lot_name,omitempty"`
	Measure                     *Float     `xmlrpc:"measure,omitempty"`
	MeasureOn                   *Selection `xmlrpc:"measure_on,omitempty"`
	MeasureSuccess              *Selection `xmlrpc:"measure_success,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveId                      *Many2One  `xmlrpc:"move_id,omitempty"`
	MoveLineId                  *Many2One  `xmlrpc:"move_line_id,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NextCheckId                 *Many2One  `xmlrpc:"next_check_id,omitempty"`
	NormUnit                    *String    `xmlrpc:"norm_unit,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	OperationId                 *Many2One  `xmlrpc:"operation_id,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingId                   *Many2One  `xmlrpc:"picking_id,omitempty"`
	Picture                     *String    `xmlrpc:"picture,omitempty"`
	PointId                     *Many2One  `xmlrpc:"point_id,omitempty"`
	PreviousCheckId             *Many2One  `xmlrpc:"previous_check_id,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTracking             *Selection `xmlrpc:"product_tracking,omitempty"`
	ProductionId                *Many2One  `xmlrpc:"production_id,omitempty"`
	QtyFailed                   *Float     `xmlrpc:"qty_failed,omitempty"`
	QtyLine                     *Float     `xmlrpc:"qty_line,omitempty"`
	QtyPassed                   *Float     `xmlrpc:"qty_passed,omitempty"`
	QtyTested                   *Float     `xmlrpc:"qty_tested,omitempty"`
	QtyToTest                   *Float     `xmlrpc:"qty_to_test,omitempty"`
	QualityState                *Selection `xmlrpc:"quality_state,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	Result                      *String    `xmlrpc:"result,omitempty"`
	ShowLotText                 *Bool      `xmlrpc:"show_lot_text,omitempty"`
	SourceDocument              *Selection `xmlrpc:"source_document,omitempty"`
	SpreadsheetCheckCell        *String    `xmlrpc:"spreadsheet_check_cell,omitempty"`
	SpreadsheetId               *Many2One  `xmlrpc:"spreadsheet_id,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	TestType                    *String    `xmlrpc:"test_type,omitempty"`
	TestTypeId                  *Many2One  `xmlrpc:"test_type_id,omitempty"`
	TestingPercentageWithinLot  *Float     `xmlrpc:"testing_percentage_within_lot,omitempty"`
	Title                       *String    `xmlrpc:"title,omitempty"`
	ToleranceMax                *Float     `xmlrpc:"tolerance_max,omitempty"`
	ToleranceMin                *Float     `xmlrpc:"tolerance_min,omitempty"`
	UomId                       *Many2One  `xmlrpc:"uom_id,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WarningMessage              *String    `xmlrpc:"warning_message,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkcenterId                *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkingState                *Selection `xmlrpc:"working_state,omitempty"`
	WorkorderId                 *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WorksheetDocument           *String    `xmlrpc:"worksheet_document,omitempty"`
	WorksheetPage               *Int       `xmlrpc:"worksheet_page,omitempty"`
	WorksheetUrl                *String    `xmlrpc:"worksheet_url,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityChecks represents array of quality.check model.
type QualityChecks []QualityCheck

// QualityCheckModel is the odoo model name.
const QualityCheckModel = "quality.check"

// Many2One convert QualityCheck to *Many2One.
func (qc *QualityCheck) Many2One() *Many2One {
	return NewMany2One(qc.Id.Get(), "")
}

// CreateQualityCheck creates a new quality.check model and returns its id.
func (c *Client) CreateQualityCheck(qc *QualityCheck) (int64, error) {
	ids, err := c.CreateQualityChecks([]*QualityCheck{qc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityCheck creates a new quality.check model and returns its id.
func (c *Client) CreateQualityChecks(qcs []*QualityCheck) ([]int64, error) {
	var vv []interface{}
	for _, v := range qcs {
		vv = append(vv, v)
	}
	return c.Create(QualityCheckModel, vv, nil)
}

// UpdateQualityCheck updates an existing quality.check record.
func (c *Client) UpdateQualityCheck(qc *QualityCheck) error {
	return c.UpdateQualityChecks([]int64{qc.Id.Get()}, qc)
}

// UpdateQualityChecks updates existing quality.check records.
// All records (represented by ids) will be updated by qc values.
func (c *Client) UpdateQualityChecks(ids []int64, qc *QualityCheck) error {
	return c.Update(QualityCheckModel, ids, qc, nil)
}

// DeleteQualityCheck deletes an existing quality.check record.
func (c *Client) DeleteQualityCheck(id int64) error {
	return c.DeleteQualityChecks([]int64{id})
}

// DeleteQualityChecks deletes existing quality.check records.
func (c *Client) DeleteQualityChecks(ids []int64) error {
	return c.Delete(QualityCheckModel, ids)
}

// GetQualityCheck gets quality.check existing record.
func (c *Client) GetQualityCheck(id int64) (*QualityCheck, error) {
	qcs, err := c.GetQualityChecks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qcs)[0]), nil
}

// GetQualityChecks gets quality.check existing records.
func (c *Client) GetQualityChecks(ids []int64) (*QualityChecks, error) {
	qcs := &QualityChecks{}
	if err := c.Read(QualityCheckModel, ids, nil, qcs); err != nil {
		return nil, err
	}
	return qcs, nil
}

// FindQualityCheck finds quality.check record by querying it with criteria.
func (c *Client) FindQualityCheck(criteria *Criteria) (*QualityCheck, error) {
	qcs := &QualityChecks{}
	if err := c.SearchRead(QualityCheckModel, criteria, NewOptions().Limit(1), qcs); err != nil {
		return nil, err
	}
	return &((*qcs)[0]), nil
}

// FindQualityChecks finds quality.check records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityChecks(criteria *Criteria, options *Options) (*QualityChecks, error) {
	qcs := &QualityChecks{}
	if err := c.SearchRead(QualityCheckModel, criteria, options, qcs); err != nil {
		return nil, err
	}
	return qcs, nil
}

// FindQualityCheckIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityCheckModel, criteria, options)
}

// FindQualityCheckId finds record id by querying it with criteria.
func (c *Client) FindQualityCheckId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityCheckModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
