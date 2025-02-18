package odoo

// QualityPoint represents quality.point model.
type QualityPoint struct {
	Active                     *Bool      `xmlrpc:"active,omitempty"`
	Average                    *Float     `xmlrpc:"average,omitempty"`
	BomActive                  *Bool      `xmlrpc:"bom_active,omitempty"`
	BomId                      *Many2One  `xmlrpc:"bom_id,omitempty"`
	BomProductIds              *Relation  `xmlrpc:"bom_product_ids,omitempty"`
	CheckCount                 *Int       `xmlrpc:"check_count,omitempty"`
	CheckIds                   *Relation  `xmlrpc:"check_ids,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	ComponentId                *Many2One  `xmlrpc:"component_id,omitempty"`
	ComponentIds               *Relation  `xmlrpc:"component_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	FailureLocationIds         *Relation  `xmlrpc:"failure_location_ids,omitempty"`
	FailureMessage             *String    `xmlrpc:"failure_message,omitempty"`
	HasMessage                 *Bool      `xmlrpc:"has_message,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	IsLotTestedFractionally    *Bool      `xmlrpc:"is_lot_tested_fractionally,omitempty"`
	IsWorkorderStep            *Bool      `xmlrpc:"is_workorder_step,omitempty"`
	MeasureFrequencyType       *Selection `xmlrpc:"measure_frequency_type,omitempty"`
	MeasureFrequencyUnit       *Selection `xmlrpc:"measure_frequency_unit,omitempty"`
	MeasureFrequencyUnitValue  *Int       `xmlrpc:"measure_frequency_unit_value,omitempty"`
	MeasureFrequencyValue      *Float     `xmlrpc:"measure_frequency_value,omitempty"`
	MeasureOn                  *Selection `xmlrpc:"measure_on,omitempty"`
	MessageAttachmentCount     *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError            *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter     *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError         *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	Norm                       *Float     `xmlrpc:"norm,omitempty"`
	NormUnit                   *String    `xmlrpc:"norm_unit,omitempty"`
	Note                       *String    `xmlrpc:"note,omitempty"`
	OperationId                *Many2One  `xmlrpc:"operation_id,omitempty"`
	PickingTypeIds             *Relation  `xmlrpc:"picking_type_ids,omitempty"`
	ProductCategoryIds         *Relation  `xmlrpc:"product_category_ids,omitempty"`
	ProductIds                 *Relation  `xmlrpc:"product_ids,omitempty"`
	RatingIds                  *Relation  `xmlrpc:"rating_ids,omitempty"`
	Reason                     *String    `xmlrpc:"reason,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	SourceDocument             *Selection `xmlrpc:"source_document,omitempty"`
	SpreadsheetCheckCell       *String    `xmlrpc:"spreadsheet_check_cell,omitempty"`
	SpreadsheetTemplateId      *Many2One  `xmlrpc:"spreadsheet_template_id,omitempty"`
	StandardDeviation          *Float     `xmlrpc:"standard_deviation,omitempty"`
	TeamId                     *Many2One  `xmlrpc:"team_id,omitempty"`
	TestReportType             *Selection `xmlrpc:"test_report_type,omitempty"`
	TestType                   *String    `xmlrpc:"test_type,omitempty"`
	TestTypeId                 *Many2One  `xmlrpc:"test_type_id,omitempty"`
	TestingPercentageWithinLot *Float     `xmlrpc:"testing_percentage_within_lot,omitempty"`
	Title                      *String    `xmlrpc:"title,omitempty"`
	ToleranceMax               *Float     `xmlrpc:"tolerance_max,omitempty"`
	ToleranceMin               *Float     `xmlrpc:"tolerance_min,omitempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorksheetDocument          *String    `xmlrpc:"worksheet_document,omitempty"`
	WorksheetPage              *Int       `xmlrpc:"worksheet_page,omitempty"`
	WorksheetUrl               *String    `xmlrpc:"worksheet_url,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityPoints represents array of quality.point model.
type QualityPoints []QualityPoint

// QualityPointModel is the odoo model name.
const QualityPointModel = "quality.point"

// Many2One convert QualityPoint to *Many2One.
func (qp *QualityPoint) Many2One() *Many2One {
	return NewMany2One(qp.Id.Get(), "")
}

// CreateQualityPoint creates a new quality.point model and returns its id.
func (c *Client) CreateQualityPoint(qp *QualityPoint) (int64, error) {
	ids, err := c.CreateQualityPoints([]*QualityPoint{qp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityPoint creates a new quality.point model and returns its id.
func (c *Client) CreateQualityPoints(qps []*QualityPoint) ([]int64, error) {
	var vv []interface{}
	for _, v := range qps {
		vv = append(vv, v)
	}
	return c.Create(QualityPointModel, vv, nil)
}

// UpdateQualityPoint updates an existing quality.point record.
func (c *Client) UpdateQualityPoint(qp *QualityPoint) error {
	return c.UpdateQualityPoints([]int64{qp.Id.Get()}, qp)
}

// UpdateQualityPoints updates existing quality.point records.
// All records (represented by ids) will be updated by qp values.
func (c *Client) UpdateQualityPoints(ids []int64, qp *QualityPoint) error {
	return c.Update(QualityPointModel, ids, qp, nil)
}

// DeleteQualityPoint deletes an existing quality.point record.
func (c *Client) DeleteQualityPoint(id int64) error {
	return c.DeleteQualityPoints([]int64{id})
}

// DeleteQualityPoints deletes existing quality.point records.
func (c *Client) DeleteQualityPoints(ids []int64) error {
	return c.Delete(QualityPointModel, ids)
}

// GetQualityPoint gets quality.point existing record.
func (c *Client) GetQualityPoint(id int64) (*QualityPoint, error) {
	qps, err := c.GetQualityPoints([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qps)[0]), nil
}

// GetQualityPoints gets quality.point existing records.
func (c *Client) GetQualityPoints(ids []int64) (*QualityPoints, error) {
	qps := &QualityPoints{}
	if err := c.Read(QualityPointModel, ids, nil, qps); err != nil {
		return nil, err
	}
	return qps, nil
}

// FindQualityPoint finds quality.point record by querying it with criteria.
func (c *Client) FindQualityPoint(criteria *Criteria) (*QualityPoint, error) {
	qps := &QualityPoints{}
	if err := c.SearchRead(QualityPointModel, criteria, NewOptions().Limit(1), qps); err != nil {
		return nil, err
	}
	return &((*qps)[0]), nil
}

// FindQualityPoints finds quality.point records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityPoints(criteria *Criteria, options *Options) (*QualityPoints, error) {
	qps := &QualityPoints{}
	if err := c.SearchRead(QualityPointModel, criteria, options, qps); err != nil {
		return nil, err
	}
	return qps, nil
}

// FindQualityPointIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityPointIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityPointModel, criteria, options)
}

// FindQualityPointId finds record id by querying it with criteria.
func (c *Client) FindQualityPointId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityPointModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
