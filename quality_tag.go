package odoo

// QualityTag represents quality.tag model.
type QualityTag struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QualityTags represents array of quality.tag model.
type QualityTags []QualityTag

// QualityTagModel is the odoo model name.
const QualityTagModel = "quality.tag"

// Many2One convert QualityTag to *Many2One.
func (qt *QualityTag) Many2One() *Many2One {
	return NewMany2One(qt.Id.Get(), "")
}

// CreateQualityTag creates a new quality.tag model and returns its id.
func (c *Client) CreateQualityTag(qt *QualityTag) (int64, error) {
	ids, err := c.CreateQualityTags([]*QualityTag{qt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityTag creates a new quality.tag model and returns its id.
func (c *Client) CreateQualityTags(qts []*QualityTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range qts {
		vv = append(vv, v)
	}
	return c.Create(QualityTagModel, vv, nil)
}

// UpdateQualityTag updates an existing quality.tag record.
func (c *Client) UpdateQualityTag(qt *QualityTag) error {
	return c.UpdateQualityTags([]int64{qt.Id.Get()}, qt)
}

// UpdateQualityTags updates existing quality.tag records.
// All records (represented by ids) will be updated by qt values.
func (c *Client) UpdateQualityTags(ids []int64, qt *QualityTag) error {
	return c.Update(QualityTagModel, ids, qt, nil)
}

// DeleteQualityTag deletes an existing quality.tag record.
func (c *Client) DeleteQualityTag(id int64) error {
	return c.DeleteQualityTags([]int64{id})
}

// DeleteQualityTags deletes existing quality.tag records.
func (c *Client) DeleteQualityTags(ids []int64) error {
	return c.Delete(QualityTagModel, ids)
}

// GetQualityTag gets quality.tag existing record.
func (c *Client) GetQualityTag(id int64) (*QualityTag, error) {
	qts, err := c.GetQualityTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qts)[0]), nil
}

// GetQualityTags gets quality.tag existing records.
func (c *Client) GetQualityTags(ids []int64) (*QualityTags, error) {
	qts := &QualityTags{}
	if err := c.Read(QualityTagModel, ids, nil, qts); err != nil {
		return nil, err
	}
	return qts, nil
}

// FindQualityTag finds quality.tag record by querying it with criteria.
func (c *Client) FindQualityTag(criteria *Criteria) (*QualityTag, error) {
	qts := &QualityTags{}
	if err := c.SearchRead(QualityTagModel, criteria, NewOptions().Limit(1), qts); err != nil {
		return nil, err
	}
	return &((*qts)[0]), nil
}

// FindQualityTags finds quality.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityTags(criteria *Criteria, options *Options) (*QualityTags, error) {
	qts := &QualityTags{}
	if err := c.SearchRead(QualityTagModel, criteria, options, qts); err != nil {
		return nil, err
	}
	return qts, nil
}

// FindQualityTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityTagModel, criteria, options)
}

// FindQualityTagId finds record id by querying it with criteria.
func (c *Client) FindQualityTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
