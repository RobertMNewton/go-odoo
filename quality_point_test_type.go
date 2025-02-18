package odoo

// QualityPointTestType represents quality.point.test_type model.
type QualityPointTestType struct {
	Active            *Bool     `xmlrpc:"active,omitempty"`
	AllowRegistration *Bool     `xmlrpc:"allow_registration,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	TechnicalName     *String   `xmlrpc:"technical_name,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QualityPointTestTypes represents array of quality.point.test_type model.
type QualityPointTestTypes []QualityPointTestType

// QualityPointTestTypeModel is the odoo model name.
const QualityPointTestTypeModel = "quality.point.test_type"

// Many2One convert QualityPointTestType to *Many2One.
func (qpt *QualityPointTestType) Many2One() *Many2One {
	return NewMany2One(qpt.Id.Get(), "")
}

// CreateQualityPointTestType creates a new quality.point.test_type model and returns its id.
func (c *Client) CreateQualityPointTestType(qpt *QualityPointTestType) (int64, error) {
	ids, err := c.CreateQualityPointTestTypes([]*QualityPointTestType{qpt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityPointTestType creates a new quality.point.test_type model and returns its id.
func (c *Client) CreateQualityPointTestTypes(qpts []*QualityPointTestType) ([]int64, error) {
	var vv []interface{}
	for _, v := range qpts {
		vv = append(vv, v)
	}
	return c.Create(QualityPointTestTypeModel, vv, nil)
}

// UpdateQualityPointTestType updates an existing quality.point.test_type record.
func (c *Client) UpdateQualityPointTestType(qpt *QualityPointTestType) error {
	return c.UpdateQualityPointTestTypes([]int64{qpt.Id.Get()}, qpt)
}

// UpdateQualityPointTestTypes updates existing quality.point.test_type records.
// All records (represented by ids) will be updated by qpt values.
func (c *Client) UpdateQualityPointTestTypes(ids []int64, qpt *QualityPointTestType) error {
	return c.Update(QualityPointTestTypeModel, ids, qpt, nil)
}

// DeleteQualityPointTestType deletes an existing quality.point.test_type record.
func (c *Client) DeleteQualityPointTestType(id int64) error {
	return c.DeleteQualityPointTestTypes([]int64{id})
}

// DeleteQualityPointTestTypes deletes existing quality.point.test_type records.
func (c *Client) DeleteQualityPointTestTypes(ids []int64) error {
	return c.Delete(QualityPointTestTypeModel, ids)
}

// GetQualityPointTestType gets quality.point.test_type existing record.
func (c *Client) GetQualityPointTestType(id int64) (*QualityPointTestType, error) {
	qpts, err := c.GetQualityPointTestTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qpts)[0]), nil
}

// GetQualityPointTestTypes gets quality.point.test_type existing records.
func (c *Client) GetQualityPointTestTypes(ids []int64) (*QualityPointTestTypes, error) {
	qpts := &QualityPointTestTypes{}
	if err := c.Read(QualityPointTestTypeModel, ids, nil, qpts); err != nil {
		return nil, err
	}
	return qpts, nil
}

// FindQualityPointTestType finds quality.point.test_type record by querying it with criteria.
func (c *Client) FindQualityPointTestType(criteria *Criteria) (*QualityPointTestType, error) {
	qpts := &QualityPointTestTypes{}
	if err := c.SearchRead(QualityPointTestTypeModel, criteria, NewOptions().Limit(1), qpts); err != nil {
		return nil, err
	}
	return &((*qpts)[0]), nil
}

// FindQualityPointTestTypes finds quality.point.test_type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityPointTestTypes(criteria *Criteria, options *Options) (*QualityPointTestTypes, error) {
	qpts := &QualityPointTestTypes{}
	if err := c.SearchRead(QualityPointTestTypeModel, criteria, options, qpts); err != nil {
		return nil, err
	}
	return qpts, nil
}

// FindQualityPointTestTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityPointTestTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityPointTestTypeModel, criteria, options)
}

// FindQualityPointTestTypeId finds record id by querying it with criteria.
func (c *Client) FindQualityPointTestTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityPointTestTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
