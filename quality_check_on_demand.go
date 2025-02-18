package odoo

// QualityCheckOnDemand represents quality.check.on.demand model.
type QualityCheckOnDemand struct {
	AllowedProductIds      *Relation  `xmlrpc:"allowed_product_ids,omitempty"`
	AllowedQualityPointIds *Relation  `xmlrpc:"allowed_quality_point_ids,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	LotId                  *Many2One  `xmlrpc:"lot_id,omitempty"`
	MeasureOn              *Selection `xmlrpc:"measure_on,omitempty"`
	PickingId              *Many2One  `xmlrpc:"picking_id,omitempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductionId           *Many2One  `xmlrpc:"production_id,omitempty"`
	QualityPointId         *Many2One  `xmlrpc:"quality_point_id,omitempty"`
	ShowLotNumber          *Bool      `xmlrpc:"show_lot_number,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityCheckOnDemands represents array of quality.check.on.demand model.
type QualityCheckOnDemands []QualityCheckOnDemand

// QualityCheckOnDemandModel is the odoo model name.
const QualityCheckOnDemandModel = "quality.check.on.demand"

// Many2One convert QualityCheckOnDemand to *Many2One.
func (qcod *QualityCheckOnDemand) Many2One() *Many2One {
	return NewMany2One(qcod.Id.Get(), "")
}

// CreateQualityCheckOnDemand creates a new quality.check.on.demand model and returns its id.
func (c *Client) CreateQualityCheckOnDemand(qcod *QualityCheckOnDemand) (int64, error) {
	ids, err := c.CreateQualityCheckOnDemands([]*QualityCheckOnDemand{qcod})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityCheckOnDemand creates a new quality.check.on.demand model and returns its id.
func (c *Client) CreateQualityCheckOnDemands(qcods []*QualityCheckOnDemand) ([]int64, error) {
	var vv []interface{}
	for _, v := range qcods {
		vv = append(vv, v)
	}
	return c.Create(QualityCheckOnDemandModel, vv, nil)
}

// UpdateQualityCheckOnDemand updates an existing quality.check.on.demand record.
func (c *Client) UpdateQualityCheckOnDemand(qcod *QualityCheckOnDemand) error {
	return c.UpdateQualityCheckOnDemands([]int64{qcod.Id.Get()}, qcod)
}

// UpdateQualityCheckOnDemands updates existing quality.check.on.demand records.
// All records (represented by ids) will be updated by qcod values.
func (c *Client) UpdateQualityCheckOnDemands(ids []int64, qcod *QualityCheckOnDemand) error {
	return c.Update(QualityCheckOnDemandModel, ids, qcod, nil)
}

// DeleteQualityCheckOnDemand deletes an existing quality.check.on.demand record.
func (c *Client) DeleteQualityCheckOnDemand(id int64) error {
	return c.DeleteQualityCheckOnDemands([]int64{id})
}

// DeleteQualityCheckOnDemands deletes existing quality.check.on.demand records.
func (c *Client) DeleteQualityCheckOnDemands(ids []int64) error {
	return c.Delete(QualityCheckOnDemandModel, ids)
}

// GetQualityCheckOnDemand gets quality.check.on.demand existing record.
func (c *Client) GetQualityCheckOnDemand(id int64) (*QualityCheckOnDemand, error) {
	qcods, err := c.GetQualityCheckOnDemands([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qcods)[0]), nil
}

// GetQualityCheckOnDemands gets quality.check.on.demand existing records.
func (c *Client) GetQualityCheckOnDemands(ids []int64) (*QualityCheckOnDemands, error) {
	qcods := &QualityCheckOnDemands{}
	if err := c.Read(QualityCheckOnDemandModel, ids, nil, qcods); err != nil {
		return nil, err
	}
	return qcods, nil
}

// FindQualityCheckOnDemand finds quality.check.on.demand record by querying it with criteria.
func (c *Client) FindQualityCheckOnDemand(criteria *Criteria) (*QualityCheckOnDemand, error) {
	qcods := &QualityCheckOnDemands{}
	if err := c.SearchRead(QualityCheckOnDemandModel, criteria, NewOptions().Limit(1), qcods); err != nil {
		return nil, err
	}
	return &((*qcods)[0]), nil
}

// FindQualityCheckOnDemands finds quality.check.on.demand records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckOnDemands(criteria *Criteria, options *Options) (*QualityCheckOnDemands, error) {
	qcods := &QualityCheckOnDemands{}
	if err := c.SearchRead(QualityCheckOnDemandModel, criteria, options, qcods); err != nil {
		return nil, err
	}
	return qcods, nil
}

// FindQualityCheckOnDemandIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckOnDemandIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityCheckOnDemandModel, criteria, options)
}

// FindQualityCheckOnDemandId finds record id by querying it with criteria.
func (c *Client) FindQualityCheckOnDemandId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityCheckOnDemandModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
