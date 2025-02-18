package odoo

// QualityCheckWizard represents quality.check.wizard model.
type QualityCheckWizard struct {
	AdditionalNote              *String    `xmlrpc:"additional_note,omitempty"`
	CheckIds                    *Relation  `xmlrpc:"check_ids,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentCheckId              *Many2One  `xmlrpc:"current_check_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	FailureLocationId           *Many2One  `xmlrpc:"failure_location_id,omitempty"`
	FailureMessage              *String    `xmlrpc:"failure_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsLastCheck                 *Bool      `xmlrpc:"is_last_check,omitempty"`
	IsLotTestedFractionally     *Bool      `xmlrpc:"is_lot_tested_fractionally,omitempty"`
	LotLineId                   *Many2One  `xmlrpc:"lot_line_id,omitempty"`
	LotName                     *String    `xmlrpc:"lot_name,omitempty"`
	Measure                     *Float     `xmlrpc:"measure,omitempty"`
	MeasureOn                   *Selection `xmlrpc:"measure_on,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NbChecks                    *Int       `xmlrpc:"nb_checks,omitempty"`
	NormUnit                    *String    `xmlrpc:"norm_unit,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	Picture                     *String    `xmlrpc:"picture,omitempty"`
	PositionCurrentCheck        *Int       `xmlrpc:"position_current_check,omitempty"`
	PotentialFailureLocationIds *Relation  `xmlrpc:"potential_failure_location_ids,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTracking             *Selection `xmlrpc:"product_tracking,omitempty"`
	QtyFailed                   *Float     `xmlrpc:"qty_failed,omitempty"`
	QtyLine                     *Float     `xmlrpc:"qty_line,omitempty"`
	QtyTested                   *Float     `xmlrpc:"qty_tested,omitempty"`
	QtyToTest                   *Float     `xmlrpc:"qty_to_test,omitempty"`
	QualityState                *Selection `xmlrpc:"quality_state,omitempty"`
	ShowLotText                 *Bool      `xmlrpc:"show_lot_text,omitempty"`
	TestType                    *String    `xmlrpc:"test_type,omitempty"`
	TestingPercentageWithinLot  *Float     `xmlrpc:"testing_percentage_within_lot,omitempty"`
	Title                       *String    `xmlrpc:"title,omitempty"`
	UomId                       *Many2One  `xmlrpc:"uom_id,omitempty"`
	WarningMessage              *String    `xmlrpc:"warning_message,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityCheckWizards represents array of quality.check.wizard model.
type QualityCheckWizards []QualityCheckWizard

// QualityCheckWizardModel is the odoo model name.
const QualityCheckWizardModel = "quality.check.wizard"

// Many2One convert QualityCheckWizard to *Many2One.
func (qcw *QualityCheckWizard) Many2One() *Many2One {
	return NewMany2One(qcw.Id.Get(), "")
}

// CreateQualityCheckWizard creates a new quality.check.wizard model and returns its id.
func (c *Client) CreateQualityCheckWizard(qcw *QualityCheckWizard) (int64, error) {
	ids, err := c.CreateQualityCheckWizards([]*QualityCheckWizard{qcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityCheckWizard creates a new quality.check.wizard model and returns its id.
func (c *Client) CreateQualityCheckWizards(qcws []*QualityCheckWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range qcws {
		vv = append(vv, v)
	}
	return c.Create(QualityCheckWizardModel, vv, nil)
}

// UpdateQualityCheckWizard updates an existing quality.check.wizard record.
func (c *Client) UpdateQualityCheckWizard(qcw *QualityCheckWizard) error {
	return c.UpdateQualityCheckWizards([]int64{qcw.Id.Get()}, qcw)
}

// UpdateQualityCheckWizards updates existing quality.check.wizard records.
// All records (represented by ids) will be updated by qcw values.
func (c *Client) UpdateQualityCheckWizards(ids []int64, qcw *QualityCheckWizard) error {
	return c.Update(QualityCheckWizardModel, ids, qcw, nil)
}

// DeleteQualityCheckWizard deletes an existing quality.check.wizard record.
func (c *Client) DeleteQualityCheckWizard(id int64) error {
	return c.DeleteQualityCheckWizards([]int64{id})
}

// DeleteQualityCheckWizards deletes existing quality.check.wizard records.
func (c *Client) DeleteQualityCheckWizards(ids []int64) error {
	return c.Delete(QualityCheckWizardModel, ids)
}

// GetQualityCheckWizard gets quality.check.wizard existing record.
func (c *Client) GetQualityCheckWizard(id int64) (*QualityCheckWizard, error) {
	qcws, err := c.GetQualityCheckWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qcws)[0]), nil
}

// GetQualityCheckWizards gets quality.check.wizard existing records.
func (c *Client) GetQualityCheckWizards(ids []int64) (*QualityCheckWizards, error) {
	qcws := &QualityCheckWizards{}
	if err := c.Read(QualityCheckWizardModel, ids, nil, qcws); err != nil {
		return nil, err
	}
	return qcws, nil
}

// FindQualityCheckWizard finds quality.check.wizard record by querying it with criteria.
func (c *Client) FindQualityCheckWizard(criteria *Criteria) (*QualityCheckWizard, error) {
	qcws := &QualityCheckWizards{}
	if err := c.SearchRead(QualityCheckWizardModel, criteria, NewOptions().Limit(1), qcws); err != nil {
		return nil, err
	}
	return &((*qcws)[0]), nil
}

// FindQualityCheckWizards finds quality.check.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckWizards(criteria *Criteria, options *Options) (*QualityCheckWizards, error) {
	qcws := &QualityCheckWizards{}
	if err := c.SearchRead(QualityCheckWizardModel, criteria, options, qcws); err != nil {
		return nil, err
	}
	return qcws, nil
}

// FindQualityCheckWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityCheckWizardModel, criteria, options)
}

// FindQualityCheckWizardId finds record id by querying it with criteria.
func (c *Client) FindQualityCheckWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityCheckWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
