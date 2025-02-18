package odoo

// SignTemplateTag represents sign.template.tag model.
type SignTemplateTag struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignTemplateTags represents array of sign.template.tag model.
type SignTemplateTags []SignTemplateTag

// SignTemplateTagModel is the odoo model name.
const SignTemplateTagModel = "sign.template.tag"

// Many2One convert SignTemplateTag to *Many2One.
func (stt *SignTemplateTag) Many2One() *Many2One {
	return NewMany2One(stt.Id.Get(), "")
}

// CreateSignTemplateTag creates a new sign.template.tag model and returns its id.
func (c *Client) CreateSignTemplateTag(stt *SignTemplateTag) (int64, error) {
	ids, err := c.CreateSignTemplateTags([]*SignTemplateTag{stt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignTemplateTag creates a new sign.template.tag model and returns its id.
func (c *Client) CreateSignTemplateTags(stts []*SignTemplateTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range stts {
		vv = append(vv, v)
	}
	return c.Create(SignTemplateTagModel, vv, nil)
}

// UpdateSignTemplateTag updates an existing sign.template.tag record.
func (c *Client) UpdateSignTemplateTag(stt *SignTemplateTag) error {
	return c.UpdateSignTemplateTags([]int64{stt.Id.Get()}, stt)
}

// UpdateSignTemplateTags updates existing sign.template.tag records.
// All records (represented by ids) will be updated by stt values.
func (c *Client) UpdateSignTemplateTags(ids []int64, stt *SignTemplateTag) error {
	return c.Update(SignTemplateTagModel, ids, stt, nil)
}

// DeleteSignTemplateTag deletes an existing sign.template.tag record.
func (c *Client) DeleteSignTemplateTag(id int64) error {
	return c.DeleteSignTemplateTags([]int64{id})
}

// DeleteSignTemplateTags deletes existing sign.template.tag records.
func (c *Client) DeleteSignTemplateTags(ids []int64) error {
	return c.Delete(SignTemplateTagModel, ids)
}

// GetSignTemplateTag gets sign.template.tag existing record.
func (c *Client) GetSignTemplateTag(id int64) (*SignTemplateTag, error) {
	stts, err := c.GetSignTemplateTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*stts)[0]), nil
}

// GetSignTemplateTags gets sign.template.tag existing records.
func (c *Client) GetSignTemplateTags(ids []int64) (*SignTemplateTags, error) {
	stts := &SignTemplateTags{}
	if err := c.Read(SignTemplateTagModel, ids, nil, stts); err != nil {
		return nil, err
	}
	return stts, nil
}

// FindSignTemplateTag finds sign.template.tag record by querying it with criteria.
func (c *Client) FindSignTemplateTag(criteria *Criteria) (*SignTemplateTag, error) {
	stts := &SignTemplateTags{}
	if err := c.SearchRead(SignTemplateTagModel, criteria, NewOptions().Limit(1), stts); err != nil {
		return nil, err
	}
	return &((*stts)[0]), nil
}

// FindSignTemplateTags finds sign.template.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplateTags(criteria *Criteria, options *Options) (*SignTemplateTags, error) {
	stts := &SignTemplateTags{}
	if err := c.SearchRead(SignTemplateTagModel, criteria, options, stts); err != nil {
		return nil, err
	}
	return stts, nil
}

// FindSignTemplateTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplateTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignTemplateTagModel, criteria, options)
}

// FindSignTemplateTagId finds record id by querying it with criteria.
func (c *Client) FindSignTemplateTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignTemplateTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
