package odoo

// L10NAuReportHandler represents l10n_au.report.handler model.
type L10NAuReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// L10NAuReportHandlers represents array of l10n_au.report.handler model.
type L10NAuReportHandlers []L10NAuReportHandler

// L10NAuReportHandlerModel is the odoo model name.
const L10NAuReportHandlerModel = "l10n_au.report.handler"

// Many2One convert L10NAuReportHandler to *Many2One.
func (lrh *L10NAuReportHandler) Many2One() *Many2One {
	return NewMany2One(lrh.Id.Get(), "")
}

// CreateL10NAuReportHandler creates a new l10n_au.report.handler model and returns its id.
func (c *Client) CreateL10NAuReportHandler(lrh *L10NAuReportHandler) (int64, error) {
	ids, err := c.CreateL10NAuReportHandlers([]*L10NAuReportHandler{lrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NAuReportHandler creates a new l10n_au.report.handler model and returns its id.
func (c *Client) CreateL10NAuReportHandlers(lrhs []*L10NAuReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range lrhs {
		vv = append(vv, v)
	}
	return c.Create(L10NAuReportHandlerModel, vv, nil)
}

// UpdateL10NAuReportHandler updates an existing l10n_au.report.handler record.
func (c *Client) UpdateL10NAuReportHandler(lrh *L10NAuReportHandler) error {
	return c.UpdateL10NAuReportHandlers([]int64{lrh.Id.Get()}, lrh)
}

// UpdateL10NAuReportHandlers updates existing l10n_au.report.handler records.
// All records (represented by ids) will be updated by lrh values.
func (c *Client) UpdateL10NAuReportHandlers(ids []int64, lrh *L10NAuReportHandler) error {
	return c.Update(L10NAuReportHandlerModel, ids, lrh, nil)
}

// DeleteL10NAuReportHandler deletes an existing l10n_au.report.handler record.
func (c *Client) DeleteL10NAuReportHandler(id int64) error {
	return c.DeleteL10NAuReportHandlers([]int64{id})
}

// DeleteL10NAuReportHandlers deletes existing l10n_au.report.handler records.
func (c *Client) DeleteL10NAuReportHandlers(ids []int64) error {
	return c.Delete(L10NAuReportHandlerModel, ids)
}

// GetL10NAuReportHandler gets l10n_au.report.handler existing record.
func (c *Client) GetL10NAuReportHandler(id int64) (*L10NAuReportHandler, error) {
	lrhs, err := c.GetL10NAuReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lrhs)[0]), nil
}

// GetL10NAuReportHandlers gets l10n_au.report.handler existing records.
func (c *Client) GetL10NAuReportHandlers(ids []int64) (*L10NAuReportHandlers, error) {
	lrhs := &L10NAuReportHandlers{}
	if err := c.Read(L10NAuReportHandlerModel, ids, nil, lrhs); err != nil {
		return nil, err
	}
	return lrhs, nil
}

// FindL10NAuReportHandler finds l10n_au.report.handler record by querying it with criteria.
func (c *Client) FindL10NAuReportHandler(criteria *Criteria) (*L10NAuReportHandler, error) {
	lrhs := &L10NAuReportHandlers{}
	if err := c.SearchRead(L10NAuReportHandlerModel, criteria, NewOptions().Limit(1), lrhs); err != nil {
		return nil, err
	}
	return &((*lrhs)[0]), nil
}

// FindL10NAuReportHandlers finds l10n_au.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NAuReportHandlers(criteria *Criteria, options *Options) (*L10NAuReportHandlers, error) {
	lrhs := &L10NAuReportHandlers{}
	if err := c.SearchRead(L10NAuReportHandlerModel, criteria, options, lrhs); err != nil {
		return nil, err
	}
	return lrhs, nil
}

// FindL10NAuReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NAuReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NAuReportHandlerModel, criteria, options)
}

// FindL10NAuReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindL10NAuReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NAuReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
