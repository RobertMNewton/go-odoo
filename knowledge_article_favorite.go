package odoo

// KnowledgeArticleFavorite represents knowledge.article.favorite model.
type KnowledgeArticleFavorite struct {
	ArticleId       *Many2One `xmlrpc:"article_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	IsArticleActive *Bool     `xmlrpc:"is_article_active,omitempty"`
	Sequence        *Int      `xmlrpc:"sequence,omitempty"`
	UserId          *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticleFavorites represents array of knowledge.article.favorite model.
type KnowledgeArticleFavorites []KnowledgeArticleFavorite

// KnowledgeArticleFavoriteModel is the odoo model name.
const KnowledgeArticleFavoriteModel = "knowledge.article.favorite"

// Many2One convert KnowledgeArticleFavorite to *Many2One.
func (kaf *KnowledgeArticleFavorite) Many2One() *Many2One {
	return NewMany2One(kaf.Id.Get(), "")
}

// CreateKnowledgeArticleFavorite creates a new knowledge.article.favorite model and returns its id.
func (c *Client) CreateKnowledgeArticleFavorite(kaf *KnowledgeArticleFavorite) (int64, error) {
	ids, err := c.CreateKnowledgeArticleFavorites([]*KnowledgeArticleFavorite{kaf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleFavorite creates a new knowledge.article.favorite model and returns its id.
func (c *Client) CreateKnowledgeArticleFavorites(kafs []*KnowledgeArticleFavorite) ([]int64, error) {
	var vv []interface{}
	for _, v := range kafs {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleFavoriteModel, vv, nil)
}

// UpdateKnowledgeArticleFavorite updates an existing knowledge.article.favorite record.
func (c *Client) UpdateKnowledgeArticleFavorite(kaf *KnowledgeArticleFavorite) error {
	return c.UpdateKnowledgeArticleFavorites([]int64{kaf.Id.Get()}, kaf)
}

// UpdateKnowledgeArticleFavorites updates existing knowledge.article.favorite records.
// All records (represented by ids) will be updated by kaf values.
func (c *Client) UpdateKnowledgeArticleFavorites(ids []int64, kaf *KnowledgeArticleFavorite) error {
	return c.Update(KnowledgeArticleFavoriteModel, ids, kaf, nil)
}

// DeleteKnowledgeArticleFavorite deletes an existing knowledge.article.favorite record.
func (c *Client) DeleteKnowledgeArticleFavorite(id int64) error {
	return c.DeleteKnowledgeArticleFavorites([]int64{id})
}

// DeleteKnowledgeArticleFavorites deletes existing knowledge.article.favorite records.
func (c *Client) DeleteKnowledgeArticleFavorites(ids []int64) error {
	return c.Delete(KnowledgeArticleFavoriteModel, ids)
}

// GetKnowledgeArticleFavorite gets knowledge.article.favorite existing record.
func (c *Client) GetKnowledgeArticleFavorite(id int64) (*KnowledgeArticleFavorite, error) {
	kafs, err := c.GetKnowledgeArticleFavorites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kafs)[0]), nil
}

// GetKnowledgeArticleFavorites gets knowledge.article.favorite existing records.
func (c *Client) GetKnowledgeArticleFavorites(ids []int64) (*KnowledgeArticleFavorites, error) {
	kafs := &KnowledgeArticleFavorites{}
	if err := c.Read(KnowledgeArticleFavoriteModel, ids, nil, kafs); err != nil {
		return nil, err
	}
	return kafs, nil
}

// FindKnowledgeArticleFavorite finds knowledge.article.favorite record by querying it with criteria.
func (c *Client) FindKnowledgeArticleFavorite(criteria *Criteria) (*KnowledgeArticleFavorite, error) {
	kafs := &KnowledgeArticleFavorites{}
	if err := c.SearchRead(KnowledgeArticleFavoriteModel, criteria, NewOptions().Limit(1), kafs); err != nil {
		return nil, err
	}
	return &((*kafs)[0]), nil
}

// FindKnowledgeArticleFavorites finds knowledge.article.favorite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleFavorites(criteria *Criteria, options *Options) (*KnowledgeArticleFavorites, error) {
	kafs := &KnowledgeArticleFavorites{}
	if err := c.SearchRead(KnowledgeArticleFavoriteModel, criteria, options, kafs); err != nil {
		return nil, err
	}
	return kafs, nil
}

// FindKnowledgeArticleFavoriteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleFavoriteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleFavoriteModel, criteria, options)
}

// FindKnowledgeArticleFavoriteId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleFavoriteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleFavoriteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
