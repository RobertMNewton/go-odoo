package odoo

// DocumentsDocument represents documents.document model.
type DocumentsDocument struct {
	AccessIds                           *Relation  `xmlrpc:"access_ids,omitempty"`
	AccessInternal                      *Selection `xmlrpc:"access_internal,omitempty"`
	AccessToken                         *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                           *String    `xmlrpc:"access_url,omitempty"`
	AccessViaLink                       *Selection `xmlrpc:"access_via_link,omitempty"`
	Active                              *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId             *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration         *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon               *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                         *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                       *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                     *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                    *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                      *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                      *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AliasBouncedContent                 *String    `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact                        *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                       *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                         *String    `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId                       *Many2One  `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail                          *String    `xmlrpc:"alias_email,omitempty"`
	AliasForceThreadId                  *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasFullName                       *String    `xmlrpc:"alias_full_name,omitempty"`
	AliasId                             *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasIncomingLocal                  *Bool      `xmlrpc:"alias_incoming_local,omitempty"`
	AliasModelId                        *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                           *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId                  *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId                 *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasStatus                         *Selection `xmlrpc:"alias_status,omitempty"`
	AliasTagIds                         *Relation  `xmlrpc:"alias_tag_ids,omitempty"`
	AttachmentId                        *Many2One  `xmlrpc:"attachment_id,omitempty"`
	AttachmentName                      *String    `xmlrpc:"attachment_name,omitempty"`
	AttachmentType                      *Selection `xmlrpc:"attachment_type,omitempty"`
	AvailableEmbeddedActionsIds         *Relation  `xmlrpc:"available_embedded_actions_ids,omitempty"`
	Checksum                            *String    `xmlrpc:"checksum,omitempty"`
	ChildrenIds                         *Relation  `xmlrpc:"children_ids,omitempty"`
	CompanyId                           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateActivityDateDeadlineRange     *Int       `xmlrpc:"create_activity_date_deadline_range,omitempty"`
	CreateActivityDateDeadlineRangeType *Selection `xmlrpc:"create_activity_date_deadline_range_type,omitempty"`
	CreateActivityNote                  *String    `xmlrpc:"create_activity_note,omitempty"`
	CreateActivityOption                *Bool      `xmlrpc:"create_activity_option,omitempty"`
	CreateActivitySummary               *String    `xmlrpc:"create_activity_summary,omitempty"`
	CreateActivityTypeId                *Many2One  `xmlrpc:"create_activity_type_id,omitempty"`
	CreateActivityUserId                *Many2One  `xmlrpc:"create_activity_user_id,omitempty"`
	CreateDate                          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentRevisionUuid                 *String    `xmlrpc:"current_revision_uuid,omitempty"`
	Datas                               *String    `xmlrpc:"datas,omitempty"`
	DeletionDelay                       *Int       `xmlrpc:"deletion_delay,omitempty"`
	Description                         *String    `xmlrpc:"description,omitempty"`
	DisplayName                         *String    `xmlrpc:"display_name,omitempty"`
	DocumentCount                       *Int       `xmlrpc:"document_count,omitempty"`
	DocumentToken                       *String    `xmlrpc:"document_token,omitempty"`
	EmailCc                             *String    `xmlrpc:"email_cc,omitempty"`
	ExcelExport                         *String    `xmlrpc:"excel_export,omitempty"`
	FavoritedIds                        *Relation  `xmlrpc:"favorited_ids,omitempty"`
	FileExtension                       *String    `xmlrpc:"file_extension,omitempty"`
	FileSize                            *Int       `xmlrpc:"file_size,omitempty"`
	FolderId                            *Many2One  `xmlrpc:"folder_id,omitempty"`
	Handler                             *Selection `xmlrpc:"handler,omitempty"`
	HasEmbeddedPdf                      *Bool      `xmlrpc:"has_embedded_pdf,omitempty"`
	HasMessage                          *Bool      `xmlrpc:"has_message,omitempty"`
	Id                                  *Int       `xmlrpc:"id,omitempty"`
	IndexContent                        *String    `xmlrpc:"index_content,omitempty"`
	IsAccessViaLinkHidden               *Bool      `xmlrpc:"is_access_via_link_hidden,omitempty"`
	IsCompanyRootFolder                 *Bool      `xmlrpc:"is_company_root_folder,omitempty"`
	IsEditableAttachment                *Bool      `xmlrpc:"is_editable_attachment,omitempty"`
	IsFavorited                         *Bool      `xmlrpc:"is_favorited,omitempty"`
	IsLocked                            *Bool      `xmlrpc:"is_locked,omitempty"`
	IsMultipage                         *Bool      `xmlrpc:"is_multipage,omitempty"`
	LastAccessDateGroup                 *Selection `xmlrpc:"last_access_date_group,omitempty"`
	LockUid                             *Many2One  `xmlrpc:"lock_uid,omitempty"`
	MessageAttachmentCount              *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                  *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                     *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter              *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                  *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                          *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                   *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                   *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter            *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                   *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Mimetype                            *String    `xmlrpc:"mimetype,omitempty"`
	MyActivityDateDeadline              *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                *String    `xmlrpc:"name,omitempty"`
	OwnerId                             *Many2One  `xmlrpc:"owner_id,omitempty"`
	ParentPath                          *String    `xmlrpc:"parent_path,omitempty"`
	PartnerId                           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PreviousAttachmentIds               *Relation  `xmlrpc:"previous_attachment_ids,omitempty"`
	ProductId                           *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTemplateId                   *Many2One  `xmlrpc:"product_template_id,omitempty"`
	ProjectIds                          *Relation  `xmlrpc:"project_ids,omitempty"`
	RatingIds                           *Relation  `xmlrpc:"rating_ids,omitempty"`
	Raw                                 *String    `xmlrpc:"raw,omitempty"`
	RequestActivityId                   *Many2One  `xmlrpc:"request_activity_id,omitempty"`
	RequesteePartnerId                  *Many2One  `xmlrpc:"requestee_partner_id,omitempty"`
	ResId                               *Many2One  `xmlrpc:"res_id,omitempty"`
	ResModel                            *String    `xmlrpc:"res_model,omitempty"`
	ResModelName                        *String    `xmlrpc:"res_model_name,omitempty"`
	ResName                             *String    `xmlrpc:"res_name,omitempty"`
	ShortcutDocumentId                  *Many2One  `xmlrpc:"shortcut_document_id,omitempty"`
	ShortcutDocumentOwnerId             *Many2One  `xmlrpc:"shortcut_document_owner_id,omitempty"`
	ShortcutIds                         *Relation  `xmlrpc:"shortcut_ids,omitempty"`
	SpreadsheetBinaryData               *String    `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData                     *String    `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName                 *String    `xmlrpc:"spreadsheet_file_name,omitempty"`
	SpreadsheetRevisionIds              *Relation  `xmlrpc:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot                 *String    `xmlrpc:"spreadsheet_snapshot,omitempty"`
	SpreadsheetThumbnailChecksum        *String    `xmlrpc:"spreadsheet_thumbnail_checksum,omitempty"`
	TagIds                              *Relation  `xmlrpc:"tag_ids,omitempty"`
	Thumbnail                           *String    `xmlrpc:"thumbnail,omitempty"`
	ThumbnailStatus                     *Selection `xmlrpc:"thumbnail_status,omitempty"`
	Type                                *Selection `xmlrpc:"type,omitempty"`
	Url                                 *String    `xmlrpc:"url,omitempty"`
	UrlPreviewImage                     *String    `xmlrpc:"url_preview_image,omitempty"`
	UserPermission                      *Selection `xmlrpc:"user_permission,omitempty"`
	WebsiteId                           *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                   *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DocumentsDocuments represents array of documents.document model.
type DocumentsDocuments []DocumentsDocument

// DocumentsDocumentModel is the odoo model name.
const DocumentsDocumentModel = "documents.document"

// Many2One convert DocumentsDocument to *Many2One.
func (dd *DocumentsDocument) Many2One() *Many2One {
	return NewMany2One(dd.Id.Get(), "")
}

// CreateDocumentsDocument creates a new documents.document model and returns its id.
func (c *Client) CreateDocumentsDocument(dd *DocumentsDocument) (int64, error) {
	ids, err := c.CreateDocumentsDocuments([]*DocumentsDocument{dd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsDocument creates a new documents.document model and returns its id.
func (c *Client) CreateDocumentsDocuments(dds []*DocumentsDocument) ([]int64, error) {
	var vv []interface{}
	for _, v := range dds {
		vv = append(vv, v)
	}
	return c.Create(DocumentsDocumentModel, vv, nil)
}

// UpdateDocumentsDocument updates an existing documents.document record.
func (c *Client) UpdateDocumentsDocument(dd *DocumentsDocument) error {
	return c.UpdateDocumentsDocuments([]int64{dd.Id.Get()}, dd)
}

// UpdateDocumentsDocuments updates existing documents.document records.
// All records (represented by ids) will be updated by dd values.
func (c *Client) UpdateDocumentsDocuments(ids []int64, dd *DocumentsDocument) error {
	return c.Update(DocumentsDocumentModel, ids, dd, nil)
}

// DeleteDocumentsDocument deletes an existing documents.document record.
func (c *Client) DeleteDocumentsDocument(id int64) error {
	return c.DeleteDocumentsDocuments([]int64{id})
}

// DeleteDocumentsDocuments deletes existing documents.document records.
func (c *Client) DeleteDocumentsDocuments(ids []int64) error {
	return c.Delete(DocumentsDocumentModel, ids)
}

// GetDocumentsDocument gets documents.document existing record.
func (c *Client) GetDocumentsDocument(id int64) (*DocumentsDocument, error) {
	dds, err := c.GetDocumentsDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dds)[0]), nil
}

// GetDocumentsDocuments gets documents.document existing records.
func (c *Client) GetDocumentsDocuments(ids []int64) (*DocumentsDocuments, error) {
	dds := &DocumentsDocuments{}
	if err := c.Read(DocumentsDocumentModel, ids, nil, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDocumentsDocument finds documents.document record by querying it with criteria.
func (c *Client) FindDocumentsDocument(criteria *Criteria) (*DocumentsDocument, error) {
	dds := &DocumentsDocuments{}
	if err := c.SearchRead(DocumentsDocumentModel, criteria, NewOptions().Limit(1), dds); err != nil {
		return nil, err
	}
	return &((*dds)[0]), nil
}

// FindDocumentsDocuments finds documents.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsDocuments(criteria *Criteria, options *Options) (*DocumentsDocuments, error) {
	dds := &DocumentsDocuments{}
	if err := c.SearchRead(DocumentsDocumentModel, criteria, options, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDocumentsDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DocumentsDocumentModel, criteria, options)
}

// FindDocumentsDocumentId finds record id by querying it with criteria.
func (c *Client) FindDocumentsDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
