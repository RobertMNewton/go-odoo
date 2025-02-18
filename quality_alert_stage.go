package odoo

// QualityAlertStage represents quality.alert.stage model.
type QualityAlertStage struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Done        *Bool     `xmlrpc:"done,omitempty"`
	Folded      *Bool     `xmlrpc:"folded,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	TeamIds     *Relation `xmlrpc:"team_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QualityAlertStages represents array of quality.alert.stage model.
type QualityAlertStages []QualityAlertStage

// QualityAlertStageModel is the odoo model name.
const QualityAlertStageModel = "quality.alert.stage"

// Many2One convert QualityAlertStage to *Many2One.
func (qas *QualityAlertStage) Many2One() *Many2One {
	return NewMany2One(qas.Id.Get(), "")
}

// CreateQualityAlertStage creates a new quality.alert.stage model and returns its id.
func (c *Client) CreateQualityAlertStage(qas *QualityAlertStage) (int64, error) {
	ids, err := c.CreateQualityAlertStages([]*QualityAlertStage{qas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityAlertStage creates a new quality.alert.stage model and returns its id.
func (c *Client) CreateQualityAlertStages(qass []*QualityAlertStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range qass {
		vv = append(vv, v)
	}
	return c.Create(QualityAlertStageModel, vv, nil)
}

// UpdateQualityAlertStage updates an existing quality.alert.stage record.
func (c *Client) UpdateQualityAlertStage(qas *QualityAlertStage) error {
	return c.UpdateQualityAlertStages([]int64{qas.Id.Get()}, qas)
}

// UpdateQualityAlertStages updates existing quality.alert.stage records.
// All records (represented by ids) will be updated by qas values.
func (c *Client) UpdateQualityAlertStages(ids []int64, qas *QualityAlertStage) error {
	return c.Update(QualityAlertStageModel, ids, qas, nil)
}

// DeleteQualityAlertStage deletes an existing quality.alert.stage record.
func (c *Client) DeleteQualityAlertStage(id int64) error {
	return c.DeleteQualityAlertStages([]int64{id})
}

// DeleteQualityAlertStages deletes existing quality.alert.stage records.
func (c *Client) DeleteQualityAlertStages(ids []int64) error {
	return c.Delete(QualityAlertStageModel, ids)
}

// GetQualityAlertStage gets quality.alert.stage existing record.
func (c *Client) GetQualityAlertStage(id int64) (*QualityAlertStage, error) {
	qass, err := c.GetQualityAlertStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qass)[0]), nil
}

// GetQualityAlertStages gets quality.alert.stage existing records.
func (c *Client) GetQualityAlertStages(ids []int64) (*QualityAlertStages, error) {
	qass := &QualityAlertStages{}
	if err := c.Read(QualityAlertStageModel, ids, nil, qass); err != nil {
		return nil, err
	}
	return qass, nil
}

// FindQualityAlertStage finds quality.alert.stage record by querying it with criteria.
func (c *Client) FindQualityAlertStage(criteria *Criteria) (*QualityAlertStage, error) {
	qass := &QualityAlertStages{}
	if err := c.SearchRead(QualityAlertStageModel, criteria, NewOptions().Limit(1), qass); err != nil {
		return nil, err
	}
	return &((*qass)[0]), nil
}

// FindQualityAlertStages finds quality.alert.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlertStages(criteria *Criteria, options *Options) (*QualityAlertStages, error) {
	qass := &QualityAlertStages{}
	if err := c.SearchRead(QualityAlertStageModel, criteria, options, qass); err != nil {
		return nil, err
	}
	return qass, nil
}

// FindQualityAlertStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlertStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityAlertStageModel, criteria, options)
}

// FindQualityAlertStageId finds record id by querying it with criteria.
func (c *Client) FindQualityAlertStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityAlertStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
