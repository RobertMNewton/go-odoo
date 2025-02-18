package odoo

// KnowledgeArticleTemplateCategory represents knowledge.article.template.category model.
type KnowledgeArticleTemplateCategory struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticleTemplateCategorys represents array of knowledge.article.template.category model.
type KnowledgeArticleTemplateCategorys []KnowledgeArticleTemplateCategory

// KnowledgeArticleTemplateCategoryModel is the odoo model name.
const KnowledgeArticleTemplateCategoryModel = "knowledge.article.template.category"

// Many2One convert KnowledgeArticleTemplateCategory to *Many2One.
func (katc *KnowledgeArticleTemplateCategory) Many2One() *Many2One {
	return NewMany2One(katc.Id.Get(), "")
}

// CreateKnowledgeArticleTemplateCategory creates a new knowledge.article.template.category model and returns its id.
func (c *Client) CreateKnowledgeArticleTemplateCategory(katc *KnowledgeArticleTemplateCategory) (int64, error) {
	ids, err := c.CreateKnowledgeArticleTemplateCategorys([]*KnowledgeArticleTemplateCategory{katc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleTemplateCategory creates a new knowledge.article.template.category model and returns its id.
func (c *Client) CreateKnowledgeArticleTemplateCategorys(katcs []*KnowledgeArticleTemplateCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range katcs {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleTemplateCategoryModel, vv, nil)
}

// UpdateKnowledgeArticleTemplateCategory updates an existing knowledge.article.template.category record.
func (c *Client) UpdateKnowledgeArticleTemplateCategory(katc *KnowledgeArticleTemplateCategory) error {
	return c.UpdateKnowledgeArticleTemplateCategorys([]int64{katc.Id.Get()}, katc)
}

// UpdateKnowledgeArticleTemplateCategorys updates existing knowledge.article.template.category records.
// All records (represented by ids) will be updated by katc values.
func (c *Client) UpdateKnowledgeArticleTemplateCategorys(ids []int64, katc *KnowledgeArticleTemplateCategory) error {
	return c.Update(KnowledgeArticleTemplateCategoryModel, ids, katc, nil)
}

// DeleteKnowledgeArticleTemplateCategory deletes an existing knowledge.article.template.category record.
func (c *Client) DeleteKnowledgeArticleTemplateCategory(id int64) error {
	return c.DeleteKnowledgeArticleTemplateCategorys([]int64{id})
}

// DeleteKnowledgeArticleTemplateCategorys deletes existing knowledge.article.template.category records.
func (c *Client) DeleteKnowledgeArticleTemplateCategorys(ids []int64) error {
	return c.Delete(KnowledgeArticleTemplateCategoryModel, ids)
}

// GetKnowledgeArticleTemplateCategory gets knowledge.article.template.category existing record.
func (c *Client) GetKnowledgeArticleTemplateCategory(id int64) (*KnowledgeArticleTemplateCategory, error) {
	katcs, err := c.GetKnowledgeArticleTemplateCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*katcs)[0]), nil
}

// GetKnowledgeArticleTemplateCategorys gets knowledge.article.template.category existing records.
func (c *Client) GetKnowledgeArticleTemplateCategorys(ids []int64) (*KnowledgeArticleTemplateCategorys, error) {
	katcs := &KnowledgeArticleTemplateCategorys{}
	if err := c.Read(KnowledgeArticleTemplateCategoryModel, ids, nil, katcs); err != nil {
		return nil, err
	}
	return katcs, nil
}

// FindKnowledgeArticleTemplateCategory finds knowledge.article.template.category record by querying it with criteria.
func (c *Client) FindKnowledgeArticleTemplateCategory(criteria *Criteria) (*KnowledgeArticleTemplateCategory, error) {
	katcs := &KnowledgeArticleTemplateCategorys{}
	if err := c.SearchRead(KnowledgeArticleTemplateCategoryModel, criteria, NewOptions().Limit(1), katcs); err != nil {
		return nil, err
	}
	return &((*katcs)[0]), nil
}

// FindKnowledgeArticleTemplateCategorys finds knowledge.article.template.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleTemplateCategorys(criteria *Criteria, options *Options) (*KnowledgeArticleTemplateCategorys, error) {
	katcs := &KnowledgeArticleTemplateCategorys{}
	if err := c.SearchRead(KnowledgeArticleTemplateCategoryModel, criteria, options, katcs); err != nil {
		return nil, err
	}
	return katcs, nil
}

// FindKnowledgeArticleTemplateCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleTemplateCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleTemplateCategoryModel, criteria, options)
}

// FindKnowledgeArticleTemplateCategoryId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleTemplateCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleTemplateCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
