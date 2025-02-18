package odoo

// QualityReason represents quality.reason model.
type QualityReason struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QualityReasons represents array of quality.reason model.
type QualityReasons []QualityReason

// QualityReasonModel is the odoo model name.
const QualityReasonModel = "quality.reason"

// Many2One convert QualityReason to *Many2One.
func (qr *QualityReason) Many2One() *Many2One {
	return NewMany2One(qr.Id.Get(), "")
}

// CreateQualityReason creates a new quality.reason model and returns its id.
func (c *Client) CreateQualityReason(qr *QualityReason) (int64, error) {
	ids, err := c.CreateQualityReasons([]*QualityReason{qr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityReason creates a new quality.reason model and returns its id.
func (c *Client) CreateQualityReasons(qrs []*QualityReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range qrs {
		vv = append(vv, v)
	}
	return c.Create(QualityReasonModel, vv, nil)
}

// UpdateQualityReason updates an existing quality.reason record.
func (c *Client) UpdateQualityReason(qr *QualityReason) error {
	return c.UpdateQualityReasons([]int64{qr.Id.Get()}, qr)
}

// UpdateQualityReasons updates existing quality.reason records.
// All records (represented by ids) will be updated by qr values.
func (c *Client) UpdateQualityReasons(ids []int64, qr *QualityReason) error {
	return c.Update(QualityReasonModel, ids, qr, nil)
}

// DeleteQualityReason deletes an existing quality.reason record.
func (c *Client) DeleteQualityReason(id int64) error {
	return c.DeleteQualityReasons([]int64{id})
}

// DeleteQualityReasons deletes existing quality.reason records.
func (c *Client) DeleteQualityReasons(ids []int64) error {
	return c.Delete(QualityReasonModel, ids)
}

// GetQualityReason gets quality.reason existing record.
func (c *Client) GetQualityReason(id int64) (*QualityReason, error) {
	qrs, err := c.GetQualityReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qrs)[0]), nil
}

// GetQualityReasons gets quality.reason existing records.
func (c *Client) GetQualityReasons(ids []int64) (*QualityReasons, error) {
	qrs := &QualityReasons{}
	if err := c.Read(QualityReasonModel, ids, nil, qrs); err != nil {
		return nil, err
	}
	return qrs, nil
}

// FindQualityReason finds quality.reason record by querying it with criteria.
func (c *Client) FindQualityReason(criteria *Criteria) (*QualityReason, error) {
	qrs := &QualityReasons{}
	if err := c.SearchRead(QualityReasonModel, criteria, NewOptions().Limit(1), qrs); err != nil {
		return nil, err
	}
	return &((*qrs)[0]), nil
}

// FindQualityReasons finds quality.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityReasons(criteria *Criteria, options *Options) (*QualityReasons, error) {
	qrs := &QualityReasons{}
	if err := c.SearchRead(QualityReasonModel, criteria, options, qrs); err != nil {
		return nil, err
	}
	return qrs, nil
}

// FindQualityReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityReasonModel, criteria, options)
}

// FindQualityReasonId finds record id by querying it with criteria.
func (c *Client) FindQualityReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
