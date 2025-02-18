package odoo

// SignDuplicateTemplatePdf represents sign.duplicate.template.pdf model.
type SignDuplicateTemplatePdf struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	NewPdf             *String   `xmlrpc:"new_pdf,omitempty"`
	NewTemplate        *String   `xmlrpc:"new_template,omitempty"`
	OriginalTemplateId *Many2One `xmlrpc:"original_template_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SignDuplicateTemplatePdfs represents array of sign.duplicate.template.pdf model.
type SignDuplicateTemplatePdfs []SignDuplicateTemplatePdf

// SignDuplicateTemplatePdfModel is the odoo model name.
const SignDuplicateTemplatePdfModel = "sign.duplicate.template.pdf"

// Many2One convert SignDuplicateTemplatePdf to *Many2One.
func (sdtp *SignDuplicateTemplatePdf) Many2One() *Many2One {
	return NewMany2One(sdtp.Id.Get(), "")
}

// CreateSignDuplicateTemplatePdf creates a new sign.duplicate.template.pdf model and returns its id.
func (c *Client) CreateSignDuplicateTemplatePdf(sdtp *SignDuplicateTemplatePdf) (int64, error) {
	ids, err := c.CreateSignDuplicateTemplatePdfs([]*SignDuplicateTemplatePdf{sdtp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignDuplicateTemplatePdf creates a new sign.duplicate.template.pdf model and returns its id.
func (c *Client) CreateSignDuplicateTemplatePdfs(sdtps []*SignDuplicateTemplatePdf) ([]int64, error) {
	var vv []interface{}
	for _, v := range sdtps {
		vv = append(vv, v)
	}
	return c.Create(SignDuplicateTemplatePdfModel, vv, nil)
}

// UpdateSignDuplicateTemplatePdf updates an existing sign.duplicate.template.pdf record.
func (c *Client) UpdateSignDuplicateTemplatePdf(sdtp *SignDuplicateTemplatePdf) error {
	return c.UpdateSignDuplicateTemplatePdfs([]int64{sdtp.Id.Get()}, sdtp)
}

// UpdateSignDuplicateTemplatePdfs updates existing sign.duplicate.template.pdf records.
// All records (represented by ids) will be updated by sdtp values.
func (c *Client) UpdateSignDuplicateTemplatePdfs(ids []int64, sdtp *SignDuplicateTemplatePdf) error {
	return c.Update(SignDuplicateTemplatePdfModel, ids, sdtp, nil)
}

// DeleteSignDuplicateTemplatePdf deletes an existing sign.duplicate.template.pdf record.
func (c *Client) DeleteSignDuplicateTemplatePdf(id int64) error {
	return c.DeleteSignDuplicateTemplatePdfs([]int64{id})
}

// DeleteSignDuplicateTemplatePdfs deletes existing sign.duplicate.template.pdf records.
func (c *Client) DeleteSignDuplicateTemplatePdfs(ids []int64) error {
	return c.Delete(SignDuplicateTemplatePdfModel, ids)
}

// GetSignDuplicateTemplatePdf gets sign.duplicate.template.pdf existing record.
func (c *Client) GetSignDuplicateTemplatePdf(id int64) (*SignDuplicateTemplatePdf, error) {
	sdtps, err := c.GetSignDuplicateTemplatePdfs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sdtps)[0]), nil
}

// GetSignDuplicateTemplatePdfs gets sign.duplicate.template.pdf existing records.
func (c *Client) GetSignDuplicateTemplatePdfs(ids []int64) (*SignDuplicateTemplatePdfs, error) {
	sdtps := &SignDuplicateTemplatePdfs{}
	if err := c.Read(SignDuplicateTemplatePdfModel, ids, nil, sdtps); err != nil {
		return nil, err
	}
	return sdtps, nil
}

// FindSignDuplicateTemplatePdf finds sign.duplicate.template.pdf record by querying it with criteria.
func (c *Client) FindSignDuplicateTemplatePdf(criteria *Criteria) (*SignDuplicateTemplatePdf, error) {
	sdtps := &SignDuplicateTemplatePdfs{}
	if err := c.SearchRead(SignDuplicateTemplatePdfModel, criteria, NewOptions().Limit(1), sdtps); err != nil {
		return nil, err
	}
	return &((*sdtps)[0]), nil
}

// FindSignDuplicateTemplatePdfs finds sign.duplicate.template.pdf records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignDuplicateTemplatePdfs(criteria *Criteria, options *Options) (*SignDuplicateTemplatePdfs, error) {
	sdtps := &SignDuplicateTemplatePdfs{}
	if err := c.SearchRead(SignDuplicateTemplatePdfModel, criteria, options, sdtps); err != nil {
		return nil, err
	}
	return sdtps, nil
}

// FindSignDuplicateTemplatePdfIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignDuplicateTemplatePdfIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignDuplicateTemplatePdfModel, criteria, options)
}

// FindSignDuplicateTemplatePdfId finds record id by querying it with criteria.
func (c *Client) FindSignDuplicateTemplatePdfId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignDuplicateTemplatePdfModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
