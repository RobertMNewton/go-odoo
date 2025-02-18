package odoo

// DocumentsAccess represents documents.access model.
type DocumentsAccess struct {
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	DocumentId     *Many2One  `xmlrpc:"document_id,omitempty"`
	ExpirationDate *Time      `xmlrpc:"expiration_date,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	LastAccessDate *Time      `xmlrpc:"last_access_date,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	Role           *Selection `xmlrpc:"role,omitempty"`
}

// DocumentsAccesss represents array of documents.access model.
type DocumentsAccesss []DocumentsAccess

// DocumentsAccessModel is the odoo model name.
const DocumentsAccessModel = "documents.access"

// Many2One convert DocumentsAccess to *Many2One.
func (da *DocumentsAccess) Many2One() *Many2One {
	return NewMany2One(da.Id.Get(), "")
}

// CreateDocumentsAccess creates a new documents.access model and returns its id.
func (c *Client) CreateDocumentsAccess(da *DocumentsAccess) (int64, error) {
	ids, err := c.CreateDocumentsAccesss([]*DocumentsAccess{da})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsAccess creates a new documents.access model and returns its id.
func (c *Client) CreateDocumentsAccesss(das []*DocumentsAccess) ([]int64, error) {
	var vv []interface{}
	for _, v := range das {
		vv = append(vv, v)
	}
	return c.Create(DocumentsAccessModel, vv, nil)
}

// UpdateDocumentsAccess updates an existing documents.access record.
func (c *Client) UpdateDocumentsAccess(da *DocumentsAccess) error {
	return c.UpdateDocumentsAccesss([]int64{da.Id.Get()}, da)
}

// UpdateDocumentsAccesss updates existing documents.access records.
// All records (represented by ids) will be updated by da values.
func (c *Client) UpdateDocumentsAccesss(ids []int64, da *DocumentsAccess) error {
	return c.Update(DocumentsAccessModel, ids, da, nil)
}

// DeleteDocumentsAccess deletes an existing documents.access record.
func (c *Client) DeleteDocumentsAccess(id int64) error {
	return c.DeleteDocumentsAccesss([]int64{id})
}

// DeleteDocumentsAccesss deletes existing documents.access records.
func (c *Client) DeleteDocumentsAccesss(ids []int64) error {
	return c.Delete(DocumentsAccessModel, ids)
}

// GetDocumentsAccess gets documents.access existing record.
func (c *Client) GetDocumentsAccess(id int64) (*DocumentsAccess, error) {
	das, err := c.GetDocumentsAccesss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*das)[0]), nil
}

// GetDocumentsAccesss gets documents.access existing records.
func (c *Client) GetDocumentsAccesss(ids []int64) (*DocumentsAccesss, error) {
	das := &DocumentsAccesss{}
	if err := c.Read(DocumentsAccessModel, ids, nil, das); err != nil {
		return nil, err
	}
	return das, nil
}

// FindDocumentsAccess finds documents.access record by querying it with criteria.
func (c *Client) FindDocumentsAccess(criteria *Criteria) (*DocumentsAccess, error) {
	das := &DocumentsAccesss{}
	if err := c.SearchRead(DocumentsAccessModel, criteria, NewOptions().Limit(1), das); err != nil {
		return nil, err
	}
	return &((*das)[0]), nil
}

// FindDocumentsAccesss finds documents.access records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccesss(criteria *Criteria, options *Options) (*DocumentsAccesss, error) {
	das := &DocumentsAccesss{}
	if err := c.SearchRead(DocumentsAccessModel, criteria, options, das); err != nil {
		return nil, err
	}
	return das, nil
}

// FindDocumentsAccessIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccessIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsAccessModel, criteria, options)
}

// FindDocumentsAccessId finds record id by querying it with criteria.
func (c *Client) FindDocumentsAccessId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsAccessModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
