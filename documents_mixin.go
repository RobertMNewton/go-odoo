package odoo

// DocumentsMixin represents documents.mixin model.
type DocumentsMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// DocumentsMixins represents array of documents.mixin model.
type DocumentsMixins []DocumentsMixin

// DocumentsMixinModel is the odoo model name.
const DocumentsMixinModel = "documents.mixin"

// Many2One convert DocumentsMixin to *Many2One.
func (dm *DocumentsMixin) Many2One() *Many2One {
	return NewMany2One(dm.Id.Get(), "")
}

// CreateDocumentsMixin creates a new documents.mixin model and returns its id.
func (c *Client) CreateDocumentsMixin(dm *DocumentsMixin) (int64, error) {
	ids, err := c.CreateDocumentsMixins([]*DocumentsMixin{dm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsMixin creates a new documents.mixin model and returns its id.
func (c *Client) CreateDocumentsMixins(dms []*DocumentsMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range dms {
		vv = append(vv, v)
	}
	return c.Create(DocumentsMixinModel, vv, nil)
}

// UpdateDocumentsMixin updates an existing documents.mixin record.
func (c *Client) UpdateDocumentsMixin(dm *DocumentsMixin) error {
	return c.UpdateDocumentsMixins([]int64{dm.Id.Get()}, dm)
}

// UpdateDocumentsMixins updates existing documents.mixin records.
// All records (represented by ids) will be updated by dm values.
func (c *Client) UpdateDocumentsMixins(ids []int64, dm *DocumentsMixin) error {
	return c.Update(DocumentsMixinModel, ids, dm, nil)
}

// DeleteDocumentsMixin deletes an existing documents.mixin record.
func (c *Client) DeleteDocumentsMixin(id int64) error {
	return c.DeleteDocumentsMixins([]int64{id})
}

// DeleteDocumentsMixins deletes existing documents.mixin records.
func (c *Client) DeleteDocumentsMixins(ids []int64) error {
	return c.Delete(DocumentsMixinModel, ids)
}

// GetDocumentsMixin gets documents.mixin existing record.
func (c *Client) GetDocumentsMixin(id int64) (*DocumentsMixin, error) {
	dms, err := c.GetDocumentsMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dms)[0]), nil
}

// GetDocumentsMixins gets documents.mixin existing records.
func (c *Client) GetDocumentsMixins(ids []int64) (*DocumentsMixins, error) {
	dms := &DocumentsMixins{}
	if err := c.Read(DocumentsMixinModel, ids, nil, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDocumentsMixin finds documents.mixin record by querying it with criteria.
func (c *Client) FindDocumentsMixin(criteria *Criteria) (*DocumentsMixin, error) {
	dms := &DocumentsMixins{}
	if err := c.SearchRead(DocumentsMixinModel, criteria, NewOptions().Limit(1), dms); err != nil {
		return nil, err
	}
	return &((*dms)[0]), nil
}

// FindDocumentsMixins finds documents.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsMixins(criteria *Criteria, options *Options) (*DocumentsMixins, error) {
	dms := &DocumentsMixins{}
	if err := c.SearchRead(DocumentsMixinModel, criteria, options, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDocumentsMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsMixinModel, criteria, options)
}

// FindDocumentsMixinId finds record id by querying it with criteria.
func (c *Client) FindDocumentsMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
