package odoo

// KnowledgeArticleMember represents knowledge.article.member model.
type KnowledgeArticleMember struct {
	ArticleId         *Many2One  `xmlrpc:"article_id,omitempty"`
	ArticlePermission *Selection `xmlrpc:"article_permission,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	Permission        *Selection `xmlrpc:"permission,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticleMembers represents array of knowledge.article.member model.
type KnowledgeArticleMembers []KnowledgeArticleMember

// KnowledgeArticleMemberModel is the odoo model name.
const KnowledgeArticleMemberModel = "knowledge.article.member"

// Many2One convert KnowledgeArticleMember to *Many2One.
func (kam *KnowledgeArticleMember) Many2One() *Many2One {
	return NewMany2One(kam.Id.Get(), "")
}

// CreateKnowledgeArticleMember creates a new knowledge.article.member model and returns its id.
func (c *Client) CreateKnowledgeArticleMember(kam *KnowledgeArticleMember) (int64, error) {
	ids, err := c.CreateKnowledgeArticleMembers([]*KnowledgeArticleMember{kam})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleMember creates a new knowledge.article.member model and returns its id.
func (c *Client) CreateKnowledgeArticleMembers(kams []*KnowledgeArticleMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range kams {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleMemberModel, vv, nil)
}

// UpdateKnowledgeArticleMember updates an existing knowledge.article.member record.
func (c *Client) UpdateKnowledgeArticleMember(kam *KnowledgeArticleMember) error {
	return c.UpdateKnowledgeArticleMembers([]int64{kam.Id.Get()}, kam)
}

// UpdateKnowledgeArticleMembers updates existing knowledge.article.member records.
// All records (represented by ids) will be updated by kam values.
func (c *Client) UpdateKnowledgeArticleMembers(ids []int64, kam *KnowledgeArticleMember) error {
	return c.Update(KnowledgeArticleMemberModel, ids, kam, nil)
}

// DeleteKnowledgeArticleMember deletes an existing knowledge.article.member record.
func (c *Client) DeleteKnowledgeArticleMember(id int64) error {
	return c.DeleteKnowledgeArticleMembers([]int64{id})
}

// DeleteKnowledgeArticleMembers deletes existing knowledge.article.member records.
func (c *Client) DeleteKnowledgeArticleMembers(ids []int64) error {
	return c.Delete(KnowledgeArticleMemberModel, ids)
}

// GetKnowledgeArticleMember gets knowledge.article.member existing record.
func (c *Client) GetKnowledgeArticleMember(id int64) (*KnowledgeArticleMember, error) {
	kams, err := c.GetKnowledgeArticleMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kams)[0]), nil
}

// GetKnowledgeArticleMembers gets knowledge.article.member existing records.
func (c *Client) GetKnowledgeArticleMembers(ids []int64) (*KnowledgeArticleMembers, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.Read(KnowledgeArticleMemberModel, ids, nil, kams); err != nil {
		return nil, err
	}
	return kams, nil
}

// FindKnowledgeArticleMember finds knowledge.article.member record by querying it with criteria.
func (c *Client) FindKnowledgeArticleMember(criteria *Criteria) (*KnowledgeArticleMember, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.SearchRead(KnowledgeArticleMemberModel, criteria, NewOptions().Limit(1), kams); err != nil {
		return nil, err
	}
	return &((*kams)[0]), nil
}

// FindKnowledgeArticleMembers finds knowledge.article.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleMembers(criteria *Criteria, options *Options) (*KnowledgeArticleMembers, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.SearchRead(KnowledgeArticleMemberModel, criteria, options, kams); err != nil {
		return nil, err
	}
	return kams, nil
}

// FindKnowledgeArticleMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleMemberModel, criteria, options)
}

// FindKnowledgeArticleMemberId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
