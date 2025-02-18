package odoo

// SignTemplate represents sign.template model.
type SignTemplate struct {
	Active           *Bool     `xmlrpc:"active,omitempty"`
	AttachmentId     *Many2One `xmlrpc:"attachment_id,omitempty"`
	AuthorizedIds    *Relation `xmlrpc:"authorized_ids,omitempty"`
	Color            *Int      `xmlrpc:"color,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Datas            *String   `xmlrpc:"datas,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	DocumentsTagIds  *Relation `xmlrpc:"documents_tag_ids,omitempty"`
	FavoritedIds     *Relation `xmlrpc:"favorited_ids,omitempty"`
	FolderId         *Many2One `xmlrpc:"folder_id,omitempty"`
	GroupIds         *Relation `xmlrpc:"group_ids,omitempty"`
	HasSignRequests  *Bool     `xmlrpc:"has_sign_requests,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	InProgressCount  *Int      `xmlrpc:"in_progress_count,omitempty"`
	IsSharing        *Bool     `xmlrpc:"is_sharing,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	NumPages         *Int      `xmlrpc:"num_pages,omitempty"`
	RedirectUrl      *String   `xmlrpc:"redirect_url,omitempty"`
	RedirectUrlText  *String   `xmlrpc:"redirect_url_text,omitempty"`
	ResponsibleCount *Int      `xmlrpc:"responsible_count,omitempty"`
	SignItemIds      *Relation `xmlrpc:"sign_item_ids,omitempty"`
	SignRequestIds   *Relation `xmlrpc:"sign_request_ids,omitempty"`
	SignedCount      *Int      `xmlrpc:"signed_count,omitempty"`
	TagIds           *Relation `xmlrpc:"tag_ids,omitempty"`
	UserId           *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignTemplates represents array of sign.template model.
type SignTemplates []SignTemplate

// SignTemplateModel is the odoo model name.
const SignTemplateModel = "sign.template"

// Many2One convert SignTemplate to *Many2One.
func (st *SignTemplate) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSignTemplate creates a new sign.template model and returns its id.
func (c *Client) CreateSignTemplate(st *SignTemplate) (int64, error) {
	ids, err := c.CreateSignTemplates([]*SignTemplate{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignTemplate creates a new sign.template model and returns its id.
func (c *Client) CreateSignTemplates(sts []*SignTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SignTemplateModel, vv, nil)
}

// UpdateSignTemplate updates an existing sign.template record.
func (c *Client) UpdateSignTemplate(st *SignTemplate) error {
	return c.UpdateSignTemplates([]int64{st.Id.Get()}, st)
}

// UpdateSignTemplates updates existing sign.template records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSignTemplates(ids []int64, st *SignTemplate) error {
	return c.Update(SignTemplateModel, ids, st, nil)
}

// DeleteSignTemplate deletes an existing sign.template record.
func (c *Client) DeleteSignTemplate(id int64) error {
	return c.DeleteSignTemplates([]int64{id})
}

// DeleteSignTemplates deletes existing sign.template records.
func (c *Client) DeleteSignTemplates(ids []int64) error {
	return c.Delete(SignTemplateModel, ids)
}

// GetSignTemplate gets sign.template existing record.
func (c *Client) GetSignTemplate(id int64) (*SignTemplate, error) {
	sts, err := c.GetSignTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// GetSignTemplates gets sign.template existing records.
func (c *Client) GetSignTemplates(ids []int64) (*SignTemplates, error) {
	sts := &SignTemplates{}
	if err := c.Read(SignTemplateModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSignTemplate finds sign.template record by querying it with criteria.
func (c *Client) FindSignTemplate(criteria *Criteria) (*SignTemplate, error) {
	sts := &SignTemplates{}
	if err := c.SearchRead(SignTemplateModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// FindSignTemplates finds sign.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplates(criteria *Criteria, options *Options) (*SignTemplates, error) {
	sts := &SignTemplates{}
	if err := c.SearchRead(SignTemplateModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSignTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignTemplateModel, criteria, options)
}

// FindSignTemplateId finds record id by querying it with criteria.
func (c *Client) FindSignTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
