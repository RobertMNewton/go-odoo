package odoo

// KnowledgeInvite represents knowledge.invite model.
type KnowledgeInvite struct {
	ArticleId         *Many2One  `xmlrpc:"article_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	HaveSharePartners *Bool      `xmlrpc:"have_share_partners,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Message           *String    `xmlrpc:"message,omitempty"`
	PartnerIds        *Relation  `xmlrpc:"partner_ids,omitempty"`
	Permission        *Selection `xmlrpc:"permission,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeInvites represents array of knowledge.invite model.
type KnowledgeInvites []KnowledgeInvite

// KnowledgeInviteModel is the odoo model name.
const KnowledgeInviteModel = "knowledge.invite"

// Many2One convert KnowledgeInvite to *Many2One.
func (ki *KnowledgeInvite) Many2One() *Many2One {
	return NewMany2One(ki.Id.Get(), "")
}

// CreateKnowledgeInvite creates a new knowledge.invite model and returns its id.
func (c *Client) CreateKnowledgeInvite(ki *KnowledgeInvite) (int64, error) {
	ids, err := c.CreateKnowledgeInvites([]*KnowledgeInvite{ki})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeInvite creates a new knowledge.invite model and returns its id.
func (c *Client) CreateKnowledgeInvites(kis []*KnowledgeInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range kis {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeInviteModel, vv, nil)
}

// UpdateKnowledgeInvite updates an existing knowledge.invite record.
func (c *Client) UpdateKnowledgeInvite(ki *KnowledgeInvite) error {
	return c.UpdateKnowledgeInvites([]int64{ki.Id.Get()}, ki)
}

// UpdateKnowledgeInvites updates existing knowledge.invite records.
// All records (represented by ids) will be updated by ki values.
func (c *Client) UpdateKnowledgeInvites(ids []int64, ki *KnowledgeInvite) error {
	return c.Update(KnowledgeInviteModel, ids, ki, nil)
}

// DeleteKnowledgeInvite deletes an existing knowledge.invite record.
func (c *Client) DeleteKnowledgeInvite(id int64) error {
	return c.DeleteKnowledgeInvites([]int64{id})
}

// DeleteKnowledgeInvites deletes existing knowledge.invite records.
func (c *Client) DeleteKnowledgeInvites(ids []int64) error {
	return c.Delete(KnowledgeInviteModel, ids)
}

// GetKnowledgeInvite gets knowledge.invite existing record.
func (c *Client) GetKnowledgeInvite(id int64) (*KnowledgeInvite, error) {
	kis, err := c.GetKnowledgeInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kis)[0]), nil
}

// GetKnowledgeInvites gets knowledge.invite existing records.
func (c *Client) GetKnowledgeInvites(ids []int64) (*KnowledgeInvites, error) {
	kis := &KnowledgeInvites{}
	if err := c.Read(KnowledgeInviteModel, ids, nil, kis); err != nil {
		return nil, err
	}
	return kis, nil
}

// FindKnowledgeInvite finds knowledge.invite record by querying it with criteria.
func (c *Client) FindKnowledgeInvite(criteria *Criteria) (*KnowledgeInvite, error) {
	kis := &KnowledgeInvites{}
	if err := c.SearchRead(KnowledgeInviteModel, criteria, NewOptions().Limit(1), kis); err != nil {
		return nil, err
	}
	return &((*kis)[0]), nil
}

// FindKnowledgeInvites finds knowledge.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeInvites(criteria *Criteria, options *Options) (*KnowledgeInvites, error) {
	kis := &KnowledgeInvites{}
	if err := c.SearchRead(KnowledgeInviteModel, criteria, options, kis); err != nil {
		return nil, err
	}
	return kis, nil
}

// FindKnowledgeInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeInviteModel, criteria, options)
}

// FindKnowledgeInviteId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
