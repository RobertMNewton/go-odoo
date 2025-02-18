package odoo

// KnowledgeArticleThread represents knowledge.article.thread model.
type KnowledgeArticleThread struct {
	ArticleAnchorText        *String   `xmlrpc:"article_anchor_text,omitempty"`
	ArticleId                *Many2One `xmlrpc:"article_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsResolved               *Bool     `xmlrpc:"is_resolved,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticleThreads represents array of knowledge.article.thread model.
type KnowledgeArticleThreads []KnowledgeArticleThread

// KnowledgeArticleThreadModel is the odoo model name.
const KnowledgeArticleThreadModel = "knowledge.article.thread"

// Many2One convert KnowledgeArticleThread to *Many2One.
func (kat *KnowledgeArticleThread) Many2One() *Many2One {
	return NewMany2One(kat.Id.Get(), "")
}

// CreateKnowledgeArticleThread creates a new knowledge.article.thread model and returns its id.
func (c *Client) CreateKnowledgeArticleThread(kat *KnowledgeArticleThread) (int64, error) {
	ids, err := c.CreateKnowledgeArticleThreads([]*KnowledgeArticleThread{kat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleThread creates a new knowledge.article.thread model and returns its id.
func (c *Client) CreateKnowledgeArticleThreads(kats []*KnowledgeArticleThread) ([]int64, error) {
	var vv []interface{}
	for _, v := range kats {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleThreadModel, vv, nil)
}

// UpdateKnowledgeArticleThread updates an existing knowledge.article.thread record.
func (c *Client) UpdateKnowledgeArticleThread(kat *KnowledgeArticleThread) error {
	return c.UpdateKnowledgeArticleThreads([]int64{kat.Id.Get()}, kat)
}

// UpdateKnowledgeArticleThreads updates existing knowledge.article.thread records.
// All records (represented by ids) will be updated by kat values.
func (c *Client) UpdateKnowledgeArticleThreads(ids []int64, kat *KnowledgeArticleThread) error {
	return c.Update(KnowledgeArticleThreadModel, ids, kat, nil)
}

// DeleteKnowledgeArticleThread deletes an existing knowledge.article.thread record.
func (c *Client) DeleteKnowledgeArticleThread(id int64) error {
	return c.DeleteKnowledgeArticleThreads([]int64{id})
}

// DeleteKnowledgeArticleThreads deletes existing knowledge.article.thread records.
func (c *Client) DeleteKnowledgeArticleThreads(ids []int64) error {
	return c.Delete(KnowledgeArticleThreadModel, ids)
}

// GetKnowledgeArticleThread gets knowledge.article.thread existing record.
func (c *Client) GetKnowledgeArticleThread(id int64) (*KnowledgeArticleThread, error) {
	kats, err := c.GetKnowledgeArticleThreads([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kats)[0]), nil
}

// GetKnowledgeArticleThreads gets knowledge.article.thread existing records.
func (c *Client) GetKnowledgeArticleThreads(ids []int64) (*KnowledgeArticleThreads, error) {
	kats := &KnowledgeArticleThreads{}
	if err := c.Read(KnowledgeArticleThreadModel, ids, nil, kats); err != nil {
		return nil, err
	}
	return kats, nil
}

// FindKnowledgeArticleThread finds knowledge.article.thread record by querying it with criteria.
func (c *Client) FindKnowledgeArticleThread(criteria *Criteria) (*KnowledgeArticleThread, error) {
	kats := &KnowledgeArticleThreads{}
	if err := c.SearchRead(KnowledgeArticleThreadModel, criteria, NewOptions().Limit(1), kats); err != nil {
		return nil, err
	}
	return &((*kats)[0]), nil
}

// FindKnowledgeArticleThreads finds knowledge.article.thread records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleThreads(criteria *Criteria, options *Options) (*KnowledgeArticleThreads, error) {
	kats := &KnowledgeArticleThreads{}
	if err := c.SearchRead(KnowledgeArticleThreadModel, criteria, options, kats); err != nil {
		return nil, err
	}
	return kats, nil
}

// FindKnowledgeArticleThreadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleThreadIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleThreadModel, criteria, options)
}

// FindKnowledgeArticleThreadId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleThreadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleThreadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
