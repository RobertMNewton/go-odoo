package odoo

// DocumentsTag represents documents.tag model.
type DocumentsTag struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	DocumentIds *Relation `xmlrpc:"document_ids,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	Tooltip     *String   `xmlrpc:"tooltip,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DocumentsTags represents array of documents.tag model.
type DocumentsTags []DocumentsTag

// DocumentsTagModel is the odoo model name.
const DocumentsTagModel = "documents.tag"

// Many2One convert DocumentsTag to *Many2One.
func (dt *DocumentsTag) Many2One() *Many2One {
	return NewMany2One(dt.Id.Get(), "")
}

// CreateDocumentsTag creates a new documents.tag model and returns its id.
func (c *Client) CreateDocumentsTag(dt *DocumentsTag) (int64, error) {
	ids, err := c.CreateDocumentsTags([]*DocumentsTag{dt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsTag creates a new documents.tag model and returns its id.
func (c *Client) CreateDocumentsTags(dts []*DocumentsTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range dts {
		vv = append(vv, v)
	}
	return c.Create(DocumentsTagModel, vv, nil)
}

// UpdateDocumentsTag updates an existing documents.tag record.
func (c *Client) UpdateDocumentsTag(dt *DocumentsTag) error {
	return c.UpdateDocumentsTags([]int64{dt.Id.Get()}, dt)
}

// UpdateDocumentsTags updates existing documents.tag records.
// All records (represented by ids) will be updated by dt values.
func (c *Client) UpdateDocumentsTags(ids []int64, dt *DocumentsTag) error {
	return c.Update(DocumentsTagModel, ids, dt, nil)
}

// DeleteDocumentsTag deletes an existing documents.tag record.
func (c *Client) DeleteDocumentsTag(id int64) error {
	return c.DeleteDocumentsTags([]int64{id})
}

// DeleteDocumentsTags deletes existing documents.tag records.
func (c *Client) DeleteDocumentsTags(ids []int64) error {
	return c.Delete(DocumentsTagModel, ids)
}

// GetDocumentsTag gets documents.tag existing record.
func (c *Client) GetDocumentsTag(id int64) (*DocumentsTag, error) {
	dts, err := c.GetDocumentsTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dts)[0]), nil
}

// GetDocumentsTags gets documents.tag existing records.
func (c *Client) GetDocumentsTags(ids []int64) (*DocumentsTags, error) {
	dts := &DocumentsTags{}
	if err := c.Read(DocumentsTagModel, ids, nil, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDocumentsTag finds documents.tag record by querying it with criteria.
func (c *Client) FindDocumentsTag(criteria *Criteria) (*DocumentsTag, error) {
	dts := &DocumentsTags{}
	if err := c.SearchRead(DocumentsTagModel, criteria, NewOptions().Limit(1), dts); err != nil {
		return nil, err
	}
	return &((*dts)[0]), nil
}

// FindDocumentsTags finds documents.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsTags(criteria *Criteria, options *Options) (*DocumentsTags, error) {
	dts := &DocumentsTags{}
	if err := c.SearchRead(DocumentsTagModel, criteria, options, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDocumentsTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsTagModel, criteria, options)
}

// FindDocumentsTagId finds record id by querying it with criteria.
func (c *Client) FindDocumentsTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
