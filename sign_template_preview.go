package odoo

// SignTemplatePreview represents sign.template.preview model.
type SignTemplatePreview struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PdfData     *String   `xmlrpc:"pdf_data,omitempty"`
	TemplateId  *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignTemplatePreviews represents array of sign.template.preview model.
type SignTemplatePreviews []SignTemplatePreview

// SignTemplatePreviewModel is the odoo model name.
const SignTemplatePreviewModel = "sign.template.preview"

// Many2One convert SignTemplatePreview to *Many2One.
func (stp *SignTemplatePreview) Many2One() *Many2One {
	return NewMany2One(stp.Id.Get(), "")
}

// CreateSignTemplatePreview creates a new sign.template.preview model and returns its id.
func (c *Client) CreateSignTemplatePreview(stp *SignTemplatePreview) (int64, error) {
	ids, err := c.CreateSignTemplatePreviews([]*SignTemplatePreview{stp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignTemplatePreview creates a new sign.template.preview model and returns its id.
func (c *Client) CreateSignTemplatePreviews(stps []*SignTemplatePreview) ([]int64, error) {
	var vv []interface{}
	for _, v := range stps {
		vv = append(vv, v)
	}
	return c.Create(SignTemplatePreviewModel, vv, nil)
}

// UpdateSignTemplatePreview updates an existing sign.template.preview record.
func (c *Client) UpdateSignTemplatePreview(stp *SignTemplatePreview) error {
	return c.UpdateSignTemplatePreviews([]int64{stp.Id.Get()}, stp)
}

// UpdateSignTemplatePreviews updates existing sign.template.preview records.
// All records (represented by ids) will be updated by stp values.
func (c *Client) UpdateSignTemplatePreviews(ids []int64, stp *SignTemplatePreview) error {
	return c.Update(SignTemplatePreviewModel, ids, stp, nil)
}

// DeleteSignTemplatePreview deletes an existing sign.template.preview record.
func (c *Client) DeleteSignTemplatePreview(id int64) error {
	return c.DeleteSignTemplatePreviews([]int64{id})
}

// DeleteSignTemplatePreviews deletes existing sign.template.preview records.
func (c *Client) DeleteSignTemplatePreviews(ids []int64) error {
	return c.Delete(SignTemplatePreviewModel, ids)
}

// GetSignTemplatePreview gets sign.template.preview existing record.
func (c *Client) GetSignTemplatePreview(id int64) (*SignTemplatePreview, error) {
	stps, err := c.GetSignTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*stps)[0]), nil
}

// GetSignTemplatePreviews gets sign.template.preview existing records.
func (c *Client) GetSignTemplatePreviews(ids []int64) (*SignTemplatePreviews, error) {
	stps := &SignTemplatePreviews{}
	if err := c.Read(SignTemplatePreviewModel, ids, nil, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSignTemplatePreview finds sign.template.preview record by querying it with criteria.
func (c *Client) FindSignTemplatePreview(criteria *Criteria) (*SignTemplatePreview, error) {
	stps := &SignTemplatePreviews{}
	if err := c.SearchRead(SignTemplatePreviewModel, criteria, NewOptions().Limit(1), stps); err != nil {
		return nil, err
	}
	return &((*stps)[0]), nil
}

// FindSignTemplatePreviews finds sign.template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplatePreviews(criteria *Criteria, options *Options) (*SignTemplatePreviews, error) {
	stps := &SignTemplatePreviews{}
	if err := c.SearchRead(SignTemplatePreviewModel, criteria, options, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSignTemplatePreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignTemplatePreviewModel, criteria, options)
}

// FindSignTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindSignTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
