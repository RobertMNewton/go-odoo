package odoo

// DocumentsRequestWizard represents documents.request_wizard model.
type DocumentsRequestWizard struct {
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omitempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omitempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omitempty"`
	RequesteeId                   *Many2One  `xmlrpc:"requestee_id,omitempty"`
	ResId                         *Int       `xmlrpc:"res_id,omitempty"`
	ResModel                      *String    `xmlrpc:"res_model,omitempty"`
	TagIds                        *Relation  `xmlrpc:"tag_ids,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DocumentsRequestWizards represents array of documents.request_wizard model.
type DocumentsRequestWizards []DocumentsRequestWizard

// DocumentsRequestWizardModel is the odoo model name.
const DocumentsRequestWizardModel = "documents.request_wizard"

// Many2One convert DocumentsRequestWizard to *Many2One.
func (dr *DocumentsRequestWizard) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDocumentsRequestWizard creates a new documents.request_wizard model and returns its id.
func (c *Client) CreateDocumentsRequestWizard(dr *DocumentsRequestWizard) (int64, error) {
	ids, err := c.CreateDocumentsRequestWizards([]*DocumentsRequestWizard{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsRequestWizard creates a new documents.request_wizard model and returns its id.
func (c *Client) CreateDocumentsRequestWizards(drs []*DocumentsRequestWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsRequestWizardModel, vv, nil)
}

// UpdateDocumentsRequestWizard updates an existing documents.request_wizard record.
func (c *Client) UpdateDocumentsRequestWizard(dr *DocumentsRequestWizard) error {
	return c.UpdateDocumentsRequestWizards([]int64{dr.Id.Get()}, dr)
}

// UpdateDocumentsRequestWizards updates existing documents.request_wizard records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDocumentsRequestWizards(ids []int64, dr *DocumentsRequestWizard) error {
	return c.Update(DocumentsRequestWizardModel, ids, dr, nil)
}

// DeleteDocumentsRequestWizard deletes an existing documents.request_wizard record.
func (c *Client) DeleteDocumentsRequestWizard(id int64) error {
	return c.DeleteDocumentsRequestWizards([]int64{id})
}

// DeleteDocumentsRequestWizards deletes existing documents.request_wizard records.
func (c *Client) DeleteDocumentsRequestWizards(ids []int64) error {
	return c.Delete(DocumentsRequestWizardModel, ids)
}

// GetDocumentsRequestWizard gets documents.request_wizard existing record.
func (c *Client) GetDocumentsRequestWizard(id int64) (*DocumentsRequestWizard, error) {
	drs, err := c.GetDocumentsRequestWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*drs)[0]), nil
}

// GetDocumentsRequestWizards gets documents.request_wizard existing records.
func (c *Client) GetDocumentsRequestWizards(ids []int64) (*DocumentsRequestWizards, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.Read(DocumentsRequestWizardModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRequestWizard finds documents.request_wizard record by querying it with criteria.
func (c *Client) FindDocumentsRequestWizard(criteria *Criteria) (*DocumentsRequestWizard, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.SearchRead(DocumentsRequestWizardModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	return &((*drs)[0]), nil
}

// FindDocumentsRequestWizards finds documents.request_wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRequestWizards(criteria *Criteria, options *Options) (*DocumentsRequestWizards, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.SearchRead(DocumentsRequestWizardModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRequestWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRequestWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsRequestWizardModel, criteria, options)
}

// FindDocumentsRequestWizardId finds record id by querying it with criteria.
func (c *Client) FindDocumentsRequestWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsRequestWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
