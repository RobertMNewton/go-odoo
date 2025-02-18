package odoo

// DocumentsRedirect represents documents.redirect model.
type DocumentsRedirect struct {
	AccessToken *String   `xmlrpc:"access_token,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	DocumentId  *Many2One `xmlrpc:"document_id,omitempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
}

// DocumentsRedirects represents array of documents.redirect model.
type DocumentsRedirects []DocumentsRedirect

// DocumentsRedirectModel is the odoo model name.
const DocumentsRedirectModel = "documents.redirect"

// Many2One convert DocumentsRedirect to *Many2One.
func (dr *DocumentsRedirect) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDocumentsRedirect creates a new documents.redirect model and returns its id.
func (c *Client) CreateDocumentsRedirect(dr *DocumentsRedirect) (int64, error) {
	ids, err := c.CreateDocumentsRedirects([]*DocumentsRedirect{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsRedirect creates a new documents.redirect model and returns its id.
func (c *Client) CreateDocumentsRedirects(drs []*DocumentsRedirect) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsRedirectModel, vv, nil)
}

// UpdateDocumentsRedirect updates an existing documents.redirect record.
func (c *Client) UpdateDocumentsRedirect(dr *DocumentsRedirect) error {
	return c.UpdateDocumentsRedirects([]int64{dr.Id.Get()}, dr)
}

// UpdateDocumentsRedirects updates existing documents.redirect records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDocumentsRedirects(ids []int64, dr *DocumentsRedirect) error {
	return c.Update(DocumentsRedirectModel, ids, dr, nil)
}

// DeleteDocumentsRedirect deletes an existing documents.redirect record.
func (c *Client) DeleteDocumentsRedirect(id int64) error {
	return c.DeleteDocumentsRedirects([]int64{id})
}

// DeleteDocumentsRedirects deletes existing documents.redirect records.
func (c *Client) DeleteDocumentsRedirects(ids []int64) error {
	return c.Delete(DocumentsRedirectModel, ids)
}

// GetDocumentsRedirect gets documents.redirect existing record.
func (c *Client) GetDocumentsRedirect(id int64) (*DocumentsRedirect, error) {
	drs, err := c.GetDocumentsRedirects([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*drs)[0]), nil
}

// GetDocumentsRedirects gets documents.redirect existing records.
func (c *Client) GetDocumentsRedirects(ids []int64) (*DocumentsRedirects, error) {
	drs := &DocumentsRedirects{}
	if err := c.Read(DocumentsRedirectModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRedirect finds documents.redirect record by querying it with criteria.
func (c *Client) FindDocumentsRedirect(criteria *Criteria) (*DocumentsRedirect, error) {
	drs := &DocumentsRedirects{}
	if err := c.SearchRead(DocumentsRedirectModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	return &((*drs)[0]), nil
}

// FindDocumentsRedirects finds documents.redirect records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRedirects(criteria *Criteria, options *Options) (*DocumentsRedirects, error) {
	drs := &DocumentsRedirects{}
	if err := c.SearchRead(DocumentsRedirectModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRedirectIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRedirectIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsRedirectModel, criteria, options)
}

// FindDocumentsRedirectId finds record id by querying it with criteria.
func (c *Client) FindDocumentsRedirectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsRedirectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
