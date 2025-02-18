package odoo

// DocumentsUnlinkMixin represents documents.unlink.mixin model.
type DocumentsUnlinkMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// DocumentsUnlinkMixins represents array of documents.unlink.mixin model.
type DocumentsUnlinkMixins []DocumentsUnlinkMixin

// DocumentsUnlinkMixinModel is the odoo model name.
const DocumentsUnlinkMixinModel = "documents.unlink.mixin"

// Many2One convert DocumentsUnlinkMixin to *Many2One.
func (dum *DocumentsUnlinkMixin) Many2One() *Many2One {
	return NewMany2One(dum.Id.Get(), "")
}

// CreateDocumentsUnlinkMixin creates a new documents.unlink.mixin model and returns its id.
func (c *Client) CreateDocumentsUnlinkMixin(dum *DocumentsUnlinkMixin) (int64, error) {
	ids, err := c.CreateDocumentsUnlinkMixins([]*DocumentsUnlinkMixin{dum})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsUnlinkMixin creates a new documents.unlink.mixin model and returns its id.
func (c *Client) CreateDocumentsUnlinkMixins(dums []*DocumentsUnlinkMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range dums {
		vv = append(vv, v)
	}
	return c.Create(DocumentsUnlinkMixinModel, vv, nil)
}

// UpdateDocumentsUnlinkMixin updates an existing documents.unlink.mixin record.
func (c *Client) UpdateDocumentsUnlinkMixin(dum *DocumentsUnlinkMixin) error {
	return c.UpdateDocumentsUnlinkMixins([]int64{dum.Id.Get()}, dum)
}

// UpdateDocumentsUnlinkMixins updates existing documents.unlink.mixin records.
// All records (represented by ids) will be updated by dum values.
func (c *Client) UpdateDocumentsUnlinkMixins(ids []int64, dum *DocumentsUnlinkMixin) error {
	return c.Update(DocumentsUnlinkMixinModel, ids, dum, nil)
}

// DeleteDocumentsUnlinkMixin deletes an existing documents.unlink.mixin record.
func (c *Client) DeleteDocumentsUnlinkMixin(id int64) error {
	return c.DeleteDocumentsUnlinkMixins([]int64{id})
}

// DeleteDocumentsUnlinkMixins deletes existing documents.unlink.mixin records.
func (c *Client) DeleteDocumentsUnlinkMixins(ids []int64) error {
	return c.Delete(DocumentsUnlinkMixinModel, ids)
}

// GetDocumentsUnlinkMixin gets documents.unlink.mixin existing record.
func (c *Client) GetDocumentsUnlinkMixin(id int64) (*DocumentsUnlinkMixin, error) {
	dums, err := c.GetDocumentsUnlinkMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dums)[0]), nil
}

// GetDocumentsUnlinkMixins gets documents.unlink.mixin existing records.
func (c *Client) GetDocumentsUnlinkMixins(ids []int64) (*DocumentsUnlinkMixins, error) {
	dums := &DocumentsUnlinkMixins{}
	if err := c.Read(DocumentsUnlinkMixinModel, ids, nil, dums); err != nil {
		return nil, err
	}
	return dums, nil
}

// FindDocumentsUnlinkMixin finds documents.unlink.mixin record by querying it with criteria.
func (c *Client) FindDocumentsUnlinkMixin(criteria *Criteria) (*DocumentsUnlinkMixin, error) {
	dums := &DocumentsUnlinkMixins{}
	if err := c.SearchRead(DocumentsUnlinkMixinModel, criteria, NewOptions().Limit(1), dums); err != nil {
		return nil, err
	}
	return &((*dums)[0]), nil
}

// FindDocumentsUnlinkMixins finds documents.unlink.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsUnlinkMixins(criteria *Criteria, options *Options) (*DocumentsUnlinkMixins, error) {
	dums := &DocumentsUnlinkMixins{}
	if err := c.SearchRead(DocumentsUnlinkMixinModel, criteria, options, dums); err != nil {
		return nil, err
	}
	return dums, nil
}

// FindDocumentsUnlinkMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsUnlinkMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsUnlinkMixinModel, criteria, options)
}

// FindDocumentsUnlinkMixinId finds record id by querying it with criteria.
func (c *Client) FindDocumentsUnlinkMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsUnlinkMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
