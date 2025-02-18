package odoo

// StudioExportWizard represents studio.export.wizard model.
type StudioExportWizard struct {
	AdditionalExportData  *Relation `xmlrpc:"additional_export_data,omitempty"`
	AdditionalModels      *Relation `xmlrpc:"additional_models,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultExportData     *Relation `xmlrpc:"default_export_data,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	IncludeAdditionalData *Bool     `xmlrpc:"include_additional_data,omitempty"`
	IncludeDemoData       *Bool     `xmlrpc:"include_demo_data,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioExportWizards represents array of studio.export.wizard model.
type StudioExportWizards []StudioExportWizard

// StudioExportWizardModel is the odoo model name.
const StudioExportWizardModel = "studio.export.wizard"

// Many2One convert StudioExportWizard to *Many2One.
func (sew *StudioExportWizard) Many2One() *Many2One {
	return NewMany2One(sew.Id.Get(), "")
}

// CreateStudioExportWizard creates a new studio.export.wizard model and returns its id.
func (c *Client) CreateStudioExportWizard(sew *StudioExportWizard) (int64, error) {
	ids, err := c.CreateStudioExportWizards([]*StudioExportWizard{sew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioExportWizard creates a new studio.export.wizard model and returns its id.
func (c *Client) CreateStudioExportWizards(sews []*StudioExportWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sews {
		vv = append(vv, v)
	}
	return c.Create(StudioExportWizardModel, vv, nil)
}

// UpdateStudioExportWizard updates an existing studio.export.wizard record.
func (c *Client) UpdateStudioExportWizard(sew *StudioExportWizard) error {
	return c.UpdateStudioExportWizards([]int64{sew.Id.Get()}, sew)
}

// UpdateStudioExportWizards updates existing studio.export.wizard records.
// All records (represented by ids) will be updated by sew values.
func (c *Client) UpdateStudioExportWizards(ids []int64, sew *StudioExportWizard) error {
	return c.Update(StudioExportWizardModel, ids, sew, nil)
}

// DeleteStudioExportWizard deletes an existing studio.export.wizard record.
func (c *Client) DeleteStudioExportWizard(id int64) error {
	return c.DeleteStudioExportWizards([]int64{id})
}

// DeleteStudioExportWizards deletes existing studio.export.wizard records.
func (c *Client) DeleteStudioExportWizards(ids []int64) error {
	return c.Delete(StudioExportWizardModel, ids)
}

// GetStudioExportWizard gets studio.export.wizard existing record.
func (c *Client) GetStudioExportWizard(id int64) (*StudioExportWizard, error) {
	sews, err := c.GetStudioExportWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sews)[0]), nil
}

// GetStudioExportWizards gets studio.export.wizard existing records.
func (c *Client) GetStudioExportWizards(ids []int64) (*StudioExportWizards, error) {
	sews := &StudioExportWizards{}
	if err := c.Read(StudioExportWizardModel, ids, nil, sews); err != nil {
		return nil, err
	}
	return sews, nil
}

// FindStudioExportWizard finds studio.export.wizard record by querying it with criteria.
func (c *Client) FindStudioExportWizard(criteria *Criteria) (*StudioExportWizard, error) {
	sews := &StudioExportWizards{}
	if err := c.SearchRead(StudioExportWizardModel, criteria, NewOptions().Limit(1), sews); err != nil {
		return nil, err
	}
	return &((*sews)[0]), nil
}

// FindStudioExportWizards finds studio.export.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportWizards(criteria *Criteria, options *Options) (*StudioExportWizards, error) {
	sews := &StudioExportWizards{}
	if err := c.SearchRead(StudioExportWizardModel, criteria, options, sews); err != nil {
		return nil, err
	}
	return sews, nil
}

// FindStudioExportWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioExportWizardModel, criteria, options)
}

// FindStudioExportWizardId finds record id by querying it with criteria.
func (c *Client) FindStudioExportWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioExportWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
