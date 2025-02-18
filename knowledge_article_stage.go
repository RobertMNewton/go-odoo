package odoo

// KnowledgeArticleStage represents knowledge.article.stage model.
type KnowledgeArticleStage struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Fold        *Bool     `xmlrpc:"fold,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticleStages represents array of knowledge.article.stage model.
type KnowledgeArticleStages []KnowledgeArticleStage

// KnowledgeArticleStageModel is the odoo model name.
const KnowledgeArticleStageModel = "knowledge.article.stage"

// Many2One convert KnowledgeArticleStage to *Many2One.
func (kas *KnowledgeArticleStage) Many2One() *Many2One {
	return NewMany2One(kas.Id.Get(), "")
}

// CreateKnowledgeArticleStage creates a new knowledge.article.stage model and returns its id.
func (c *Client) CreateKnowledgeArticleStage(kas *KnowledgeArticleStage) (int64, error) {
	ids, err := c.CreateKnowledgeArticleStages([]*KnowledgeArticleStage{kas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleStage creates a new knowledge.article.stage model and returns its id.
func (c *Client) CreateKnowledgeArticleStages(kass []*KnowledgeArticleStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range kass {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleStageModel, vv, nil)
}

// UpdateKnowledgeArticleStage updates an existing knowledge.article.stage record.
func (c *Client) UpdateKnowledgeArticleStage(kas *KnowledgeArticleStage) error {
	return c.UpdateKnowledgeArticleStages([]int64{kas.Id.Get()}, kas)
}

// UpdateKnowledgeArticleStages updates existing knowledge.article.stage records.
// All records (represented by ids) will be updated by kas values.
func (c *Client) UpdateKnowledgeArticleStages(ids []int64, kas *KnowledgeArticleStage) error {
	return c.Update(KnowledgeArticleStageModel, ids, kas, nil)
}

// DeleteKnowledgeArticleStage deletes an existing knowledge.article.stage record.
func (c *Client) DeleteKnowledgeArticleStage(id int64) error {
	return c.DeleteKnowledgeArticleStages([]int64{id})
}

// DeleteKnowledgeArticleStages deletes existing knowledge.article.stage records.
func (c *Client) DeleteKnowledgeArticleStages(ids []int64) error {
	return c.Delete(KnowledgeArticleStageModel, ids)
}

// GetKnowledgeArticleStage gets knowledge.article.stage existing record.
func (c *Client) GetKnowledgeArticleStage(id int64) (*KnowledgeArticleStage, error) {
	kass, err := c.GetKnowledgeArticleStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kass)[0]), nil
}

// GetKnowledgeArticleStages gets knowledge.article.stage existing records.
func (c *Client) GetKnowledgeArticleStages(ids []int64) (*KnowledgeArticleStages, error) {
	kass := &KnowledgeArticleStages{}
	if err := c.Read(KnowledgeArticleStageModel, ids, nil, kass); err != nil {
		return nil, err
	}
	return kass, nil
}

// FindKnowledgeArticleStage finds knowledge.article.stage record by querying it with criteria.
func (c *Client) FindKnowledgeArticleStage(criteria *Criteria) (*KnowledgeArticleStage, error) {
	kass := &KnowledgeArticleStages{}
	if err := c.SearchRead(KnowledgeArticleStageModel, criteria, NewOptions().Limit(1), kass); err != nil {
		return nil, err
	}
	return &((*kass)[0]), nil
}

// FindKnowledgeArticleStages finds knowledge.article.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleStages(criteria *Criteria, options *Options) (*KnowledgeArticleStages, error) {
	kass := &KnowledgeArticleStages{}
	if err := c.SearchRead(KnowledgeArticleStageModel, criteria, options, kass); err != nil {
		return nil, err
	}
	return kass, nil
}

// FindKnowledgeArticleStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleStageModel, criteria, options)
}

// FindKnowledgeArticleStageId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
