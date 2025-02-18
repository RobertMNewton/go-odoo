package odoo

// ExtractMixin represents extract.mixin model.
type ExtractMixin struct {
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	ExtractCanShowSendButton *Bool      `xmlrpc:"extract_can_show_send_button,omitempty"`
	ExtractDocumentUuid      *String    `xmlrpc:"extract_document_uuid,omitempty"`
	ExtractErrorMessage      *String    `xmlrpc:"extract_error_message,omitempty"`
	ExtractState             *Selection `xmlrpc:"extract_state,omitempty"`
	ExtractStateProcessed    *Bool      `xmlrpc:"extract_state_processed,omitempty"`
	ExtractStatus            *String    `xmlrpc:"extract_status,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsInExtractableState     *Bool      `xmlrpc:"is_in_extractable_state,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
}

// ExtractMixins represents array of extract.mixin model.
type ExtractMixins []ExtractMixin

// ExtractMixinModel is the odoo model name.
const ExtractMixinModel = "extract.mixin"

// Many2One convert ExtractMixin to *Many2One.
func (em *ExtractMixin) Many2One() *Many2One {
	return NewMany2One(em.Id.Get(), "")
}

// CreateExtractMixin creates a new extract.mixin model and returns its id.
func (c *Client) CreateExtractMixin(em *ExtractMixin) (int64, error) {
	ids, err := c.CreateExtractMixins([]*ExtractMixin{em})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateExtractMixin creates a new extract.mixin model and returns its id.
func (c *Client) CreateExtractMixins(ems []*ExtractMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ems {
		vv = append(vv, v)
	}
	return c.Create(ExtractMixinModel, vv, nil)
}

// UpdateExtractMixin updates an existing extract.mixin record.
func (c *Client) UpdateExtractMixin(em *ExtractMixin) error {
	return c.UpdateExtractMixins([]int64{em.Id.Get()}, em)
}

// UpdateExtractMixins updates existing extract.mixin records.
// All records (represented by ids) will be updated by em values.
func (c *Client) UpdateExtractMixins(ids []int64, em *ExtractMixin) error {
	return c.Update(ExtractMixinModel, ids, em, nil)
}

// DeleteExtractMixin deletes an existing extract.mixin record.
func (c *Client) DeleteExtractMixin(id int64) error {
	return c.DeleteExtractMixins([]int64{id})
}

// DeleteExtractMixins deletes existing extract.mixin records.
func (c *Client) DeleteExtractMixins(ids []int64) error {
	return c.Delete(ExtractMixinModel, ids)
}

// GetExtractMixin gets extract.mixin existing record.
func (c *Client) GetExtractMixin(id int64) (*ExtractMixin, error) {
	ems, err := c.GetExtractMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// GetExtractMixins gets extract.mixin existing records.
func (c *Client) GetExtractMixins(ids []int64) (*ExtractMixins, error) {
	ems := &ExtractMixins{}
	if err := c.Read(ExtractMixinModel, ids, nil, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindExtractMixin finds extract.mixin record by querying it with criteria.
func (c *Client) FindExtractMixin(criteria *Criteria) (*ExtractMixin, error) {
	ems := &ExtractMixins{}
	if err := c.SearchRead(ExtractMixinModel, criteria, NewOptions().Limit(1), ems); err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// FindExtractMixins finds extract.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindExtractMixins(criteria *Criteria, options *Options) (*ExtractMixins, error) {
	ems := &ExtractMixins{}
	if err := c.SearchRead(ExtractMixinModel, criteria, options, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindExtractMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindExtractMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ExtractMixinModel, criteria, options)
}

// FindExtractMixinId finds record id by querying it with criteria.
func (c *Client) FindExtractMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ExtractMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
