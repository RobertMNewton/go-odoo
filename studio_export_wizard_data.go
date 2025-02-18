package odoo

// StudioExportWizardData represents studio.export.wizard.data model.
type StudioExportWizardData struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	IsDemoData  *Bool     `xmlrpc:"is_demo_data,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	ModelName   *String   `xmlrpc:"model_name,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Post        *Bool     `xmlrpc:"post,omitempty"`
	Pre         *Bool     `xmlrpc:"pre,omitempty"`
	ResId       *Many2One `xmlrpc:"res_id,omitempty"`
	Studio      *Bool     `xmlrpc:"studio,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	Xmlid       *String   `xmlrpc:"xmlid,omitempty"`
}

// StudioExportWizardDatas represents array of studio.export.wizard.data model.
type StudioExportWizardDatas []StudioExportWizardData

// StudioExportWizardDataModel is the odoo model name.
const StudioExportWizardDataModel = "studio.export.wizard.data"

// Many2One convert StudioExportWizardData to *Many2One.
func (sewd *StudioExportWizardData) Many2One() *Many2One {
	return NewMany2One(sewd.Id.Get(), "")
}

// CreateStudioExportWizardData creates a new studio.export.wizard.data model and returns its id.
func (c *Client) CreateStudioExportWizardData(sewd *StudioExportWizardData) (int64, error) {
	ids, err := c.CreateStudioExportWizardDatas([]*StudioExportWizardData{sewd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioExportWizardData creates a new studio.export.wizard.data model and returns its id.
func (c *Client) CreateStudioExportWizardDatas(sewds []*StudioExportWizardData) ([]int64, error) {
	var vv []interface{}
	for _, v := range sewds {
		vv = append(vv, v)
	}
	return c.Create(StudioExportWizardDataModel, vv, nil)
}

// UpdateStudioExportWizardData updates an existing studio.export.wizard.data record.
func (c *Client) UpdateStudioExportWizardData(sewd *StudioExportWizardData) error {
	return c.UpdateStudioExportWizardDatas([]int64{sewd.Id.Get()}, sewd)
}

// UpdateStudioExportWizardDatas updates existing studio.export.wizard.data records.
// All records (represented by ids) will be updated by sewd values.
func (c *Client) UpdateStudioExportWizardDatas(ids []int64, sewd *StudioExportWizardData) error {
	return c.Update(StudioExportWizardDataModel, ids, sewd, nil)
}

// DeleteStudioExportWizardData deletes an existing studio.export.wizard.data record.
func (c *Client) DeleteStudioExportWizardData(id int64) error {
	return c.DeleteStudioExportWizardDatas([]int64{id})
}

// DeleteStudioExportWizardDatas deletes existing studio.export.wizard.data records.
func (c *Client) DeleteStudioExportWizardDatas(ids []int64) error {
	return c.Delete(StudioExportWizardDataModel, ids)
}

// GetStudioExportWizardData gets studio.export.wizard.data existing record.
func (c *Client) GetStudioExportWizardData(id int64) (*StudioExportWizardData, error) {
	sewds, err := c.GetStudioExportWizardDatas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sewds)[0]), nil
}

// GetStudioExportWizardDatas gets studio.export.wizard.data existing records.
func (c *Client) GetStudioExportWizardDatas(ids []int64) (*StudioExportWizardDatas, error) {
	sewds := &StudioExportWizardDatas{}
	if err := c.Read(StudioExportWizardDataModel, ids, nil, sewds); err != nil {
		return nil, err
	}
	return sewds, nil
}

// FindStudioExportWizardData finds studio.export.wizard.data record by querying it with criteria.
func (c *Client) FindStudioExportWizardData(criteria *Criteria) (*StudioExportWizardData, error) {
	sewds := &StudioExportWizardDatas{}
	if err := c.SearchRead(StudioExportWizardDataModel, criteria, NewOptions().Limit(1), sewds); err != nil {
		return nil, err
	}
	return &((*sewds)[0]), nil
}

// FindStudioExportWizardDatas finds studio.export.wizard.data records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportWizardDatas(criteria *Criteria, options *Options) (*StudioExportWizardDatas, error) {
	sewds := &StudioExportWizardDatas{}
	if err := c.SearchRead(StudioExportWizardDataModel, criteria, options, sewds); err != nil {
		return nil, err
	}
	return sewds, nil
}

// FindStudioExportWizardDataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportWizardDataIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioExportWizardDataModel, criteria, options)
}

// FindStudioExportWizardDataId finds record id by querying it with criteria.
func (c *Client) FindStudioExportWizardDataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioExportWizardDataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
