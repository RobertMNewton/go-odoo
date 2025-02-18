package odoo

// StudioExportModel represents studio.export.model model.
type StudioExportModel struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Domain            *String   `xmlrpc:"domain,omitempty"`
	ExcludedFields    *Relation `xmlrpc:"excluded_fields,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	IncludeAttachment *Bool     `xmlrpc:"include_attachment,omitempty"`
	IsDemoData        *Bool     `xmlrpc:"is_demo_data,omitempty"`
	ModelId           *Many2One `xmlrpc:"model_id,omitempty"`
	ModelName         *String   `xmlrpc:"model_name,omitempty"`
	RecordsCount      *String   `xmlrpc:"records_count,omitempty"`
	Sequence          *Int      `xmlrpc:"sequence,omitempty"`
	Updatable         *Bool     `xmlrpc:"updatable,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StudioExportModels represents array of studio.export.model model.
type StudioExportModels []StudioExportModel

// StudioExportModelModel is the odoo model name.
const StudioExportModelModel = "studio.export.model"

// Many2One convert StudioExportModel to *Many2One.
func (sem *StudioExportModel) Many2One() *Many2One {
	return NewMany2One(sem.Id.Get(), "")
}

// CreateStudioExportModel creates a new studio.export.model model and returns its id.
func (c *Client) CreateStudioExportModel(sem *StudioExportModel) (int64, error) {
	ids, err := c.CreateStudioExportModels([]*StudioExportModel{sem})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioExportModel creates a new studio.export.model model and returns its id.
func (c *Client) CreateStudioExportModels(sems []*StudioExportModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range sems {
		vv = append(vv, v)
	}
	return c.Create(StudioExportModelModel, vv, nil)
}

// UpdateStudioExportModel updates an existing studio.export.model record.
func (c *Client) UpdateStudioExportModel(sem *StudioExportModel) error {
	return c.UpdateStudioExportModels([]int64{sem.Id.Get()}, sem)
}

// UpdateStudioExportModels updates existing studio.export.model records.
// All records (represented by ids) will be updated by sem values.
func (c *Client) UpdateStudioExportModels(ids []int64, sem *StudioExportModel) error {
	return c.Update(StudioExportModelModel, ids, sem, nil)
}

// DeleteStudioExportModel deletes an existing studio.export.model record.
func (c *Client) DeleteStudioExportModel(id int64) error {
	return c.DeleteStudioExportModels([]int64{id})
}

// DeleteStudioExportModels deletes existing studio.export.model records.
func (c *Client) DeleteStudioExportModels(ids []int64) error {
	return c.Delete(StudioExportModelModel, ids)
}

// GetStudioExportModel gets studio.export.model existing record.
func (c *Client) GetStudioExportModel(id int64) (*StudioExportModel, error) {
	sems, err := c.GetStudioExportModels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sems)[0]), nil
}

// GetStudioExportModels gets studio.export.model existing records.
func (c *Client) GetStudioExportModels(ids []int64) (*StudioExportModels, error) {
	sems := &StudioExportModels{}
	if err := c.Read(StudioExportModelModel, ids, nil, sems); err != nil {
		return nil, err
	}
	return sems, nil
}

// FindStudioExportModel finds studio.export.model record by querying it with criteria.
func (c *Client) FindStudioExportModel(criteria *Criteria) (*StudioExportModel, error) {
	sems := &StudioExportModels{}
	if err := c.SearchRead(StudioExportModelModel, criteria, NewOptions().Limit(1), sems); err != nil {
		return nil, err
	}
	return &((*sems)[0]), nil
}

// FindStudioExportModels finds studio.export.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportModels(criteria *Criteria, options *Options) (*StudioExportModels, error) {
	sems := &StudioExportModels{}
	if err := c.SearchRead(StudioExportModelModel, criteria, options, sems); err != nil {
		return nil, err
	}
	return sems, nil
}

// FindStudioExportModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioExportModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StudioExportModelModel, criteria, options)
}

// FindStudioExportModelId finds record id by querying it with criteria.
func (c *Client) FindStudioExportModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioExportModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
