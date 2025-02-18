package odoo

// DocumentsAccountFolderSetting represents documents.account.folder.setting model.
type DocumentsAccountFolderSetting struct {
	CompanyAccountFolderId *Many2One `xmlrpc:"company_account_folder_id,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	FolderId               *Many2One `xmlrpc:"folder_id,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	JournalId              *Many2One `xmlrpc:"journal_id,omitempty"`
	TagIds                 *Relation `xmlrpc:"tag_ids,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DocumentsAccountFolderSettings represents array of documents.account.folder.setting model.
type DocumentsAccountFolderSettings []DocumentsAccountFolderSetting

// DocumentsAccountFolderSettingModel is the odoo model name.
const DocumentsAccountFolderSettingModel = "documents.account.folder.setting"

// Many2One convert DocumentsAccountFolderSetting to *Many2One.
func (dafs *DocumentsAccountFolderSetting) Many2One() *Many2One {
	return NewMany2One(dafs.Id.Get(), "")
}

// CreateDocumentsAccountFolderSetting creates a new documents.account.folder.setting model and returns its id.
func (c *Client) CreateDocumentsAccountFolderSetting(dafs *DocumentsAccountFolderSetting) (int64, error) {
	ids, err := c.CreateDocumentsAccountFolderSettings([]*DocumentsAccountFolderSetting{dafs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsAccountFolderSetting creates a new documents.account.folder.setting model and returns its id.
func (c *Client) CreateDocumentsAccountFolderSettings(dafss []*DocumentsAccountFolderSetting) ([]int64, error) {
	var vv []interface{}
	for _, v := range dafss {
		vv = append(vv, v)
	}
	return c.Create(DocumentsAccountFolderSettingModel, vv, nil)
}

// UpdateDocumentsAccountFolderSetting updates an existing documents.account.folder.setting record.
func (c *Client) UpdateDocumentsAccountFolderSetting(dafs *DocumentsAccountFolderSetting) error {
	return c.UpdateDocumentsAccountFolderSettings([]int64{dafs.Id.Get()}, dafs)
}

// UpdateDocumentsAccountFolderSettings updates existing documents.account.folder.setting records.
// All records (represented by ids) will be updated by dafs values.
func (c *Client) UpdateDocumentsAccountFolderSettings(ids []int64, dafs *DocumentsAccountFolderSetting) error {
	return c.Update(DocumentsAccountFolderSettingModel, ids, dafs, nil)
}

// DeleteDocumentsAccountFolderSetting deletes an existing documents.account.folder.setting record.
func (c *Client) DeleteDocumentsAccountFolderSetting(id int64) error {
	return c.DeleteDocumentsAccountFolderSettings([]int64{id})
}

// DeleteDocumentsAccountFolderSettings deletes existing documents.account.folder.setting records.
func (c *Client) DeleteDocumentsAccountFolderSettings(ids []int64) error {
	return c.Delete(DocumentsAccountFolderSettingModel, ids)
}

// GetDocumentsAccountFolderSetting gets documents.account.folder.setting existing record.
func (c *Client) GetDocumentsAccountFolderSetting(id int64) (*DocumentsAccountFolderSetting, error) {
	dafss, err := c.GetDocumentsAccountFolderSettings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dafss)[0]), nil
}

// GetDocumentsAccountFolderSettings gets documents.account.folder.setting existing records.
func (c *Client) GetDocumentsAccountFolderSettings(ids []int64) (*DocumentsAccountFolderSettings, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.Read(DocumentsAccountFolderSettingModel, ids, nil, dafss); err != nil {
		return nil, err
	}
	return dafss, nil
}

// FindDocumentsAccountFolderSetting finds documents.account.folder.setting record by querying it with criteria.
func (c *Client) FindDocumentsAccountFolderSetting(criteria *Criteria) (*DocumentsAccountFolderSetting, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.SearchRead(DocumentsAccountFolderSettingModel, criteria, NewOptions().Limit(1), dafss); err != nil {
		return nil, err
	}
	return &((*dafss)[0]), nil
}

// FindDocumentsAccountFolderSettings finds documents.account.folder.setting records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccountFolderSettings(criteria *Criteria, options *Options) (*DocumentsAccountFolderSettings, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.SearchRead(DocumentsAccountFolderSettingModel, criteria, options, dafss); err != nil {
		return nil, err
	}
	return dafss, nil
}

// FindDocumentsAccountFolderSettingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccountFolderSettingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsAccountFolderSettingModel, criteria, options)
}

// FindDocumentsAccountFolderSettingId finds record id by querying it with criteria.
func (c *Client) FindDocumentsAccountFolderSettingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsAccountFolderSettingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
