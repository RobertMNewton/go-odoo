package odoo

// KnowledgeArticle represents knowledge.article model.
type KnowledgeArticle struct {
	Active                      *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	ArticleMemberIds            *Relation   `xmlrpc:"article_member_ids,omitempty"`
	ArticleProperties           interface{} `xmlrpc:"article_properties,omitempty"`
	ArticlePropertiesDefinition interface{} `xmlrpc:"article_properties_definition,omitempty"`
	ArticleUrl                  *String     `xmlrpc:"article_url,omitempty"`
	Body                        *String     `xmlrpc:"body,omitempty"`
	CanPublish                  *Bool       `xmlrpc:"can_publish,omitempty"`
	Category                    *Selection  `xmlrpc:"category,omitempty"`
	ChildIds                    *Relation   `xmlrpc:"child_ids,omitempty"`
	CoverImageId                *Many2One   `xmlrpc:"cover_image_id,omitempty"`
	CoverImagePosition          *Float      `xmlrpc:"cover_image_position,omitempty"`
	CoverImageUrl               *String     `xmlrpc:"cover_image_url,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	DeletionDate                *Time       `xmlrpc:"deletion_date,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	FavoriteCount               *Int        `xmlrpc:"favorite_count,omitempty"`
	FavoriteIds                 *Relation   `xmlrpc:"favorite_ids,omitempty"`
	FullWidth                   *Bool       `xmlrpc:"full_width,omitempty"`
	HasArticleChildren          *Bool       `xmlrpc:"has_article_children,omitempty"`
	HasItemChildren             *Bool       `xmlrpc:"has_item_children,omitempty"`
	HasItemParent               *Bool       `xmlrpc:"has_item_parent,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	HtmlFieldHistory            interface{} `xmlrpc:"html_field_history,omitempty"`
	HtmlFieldHistoryMetadata    interface{} `xmlrpc:"html_field_history_metadata,omitempty"`
	Icon                        *String     `xmlrpc:"icon,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	InheritedPermission         *Selection  `xmlrpc:"inherited_permission,omitempty"`
	InheritedPermissionParentId *Many2One   `xmlrpc:"inherited_permission_parent_id,omitempty"`
	InternalPermission          *Selection  `xmlrpc:"internal_permission,omitempty"`
	IsArticleItem               *Bool       `xmlrpc:"is_article_item,omitempty"`
	IsArticleVisible            *Bool       `xmlrpc:"is_article_visible,omitempty"`
	IsArticleVisibleByEveryone  *Bool       `xmlrpc:"is_article_visible_by_everyone,omitempty"`
	IsDesynchronized            *Bool       `xmlrpc:"is_desynchronized,omitempty"`
	IsLocked                    *Bool       `xmlrpc:"is_locked,omitempty"`
	IsPublished                 *Bool       `xmlrpc:"is_published,omitempty"`
	IsTemplate                  *Bool       `xmlrpc:"is_template,omitempty"`
	IsUserFavorite              *Bool       `xmlrpc:"is_user_favorite,omitempty"`
	LastEditionDate             *Time       `xmlrpc:"last_edition_date,omitempty"`
	LastEditionUid              *Many2One   `xmlrpc:"last_edition_uid,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty"`
	ParentId                    *Many2One   `xmlrpc:"parent_id,omitempty"`
	ParentPath                  *String     `xmlrpc:"parent_path,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	RootArticleId               *Many2One   `xmlrpc:"root_article_id,omitempty"`
	Sequence                    *Int        `xmlrpc:"sequence,omitempty"`
	StageId                     *Many2One   `xmlrpc:"stage_id,omitempty"`
	Summary                     *String     `xmlrpc:"summary,omitempty"`
	TemplateBody                *String     `xmlrpc:"template_body,omitempty"`
	TemplateCategoryId          *Many2One   `xmlrpc:"template_category_id,omitempty"`
	TemplateCategorySequence    *Int        `xmlrpc:"template_category_sequence,omitempty"`
	TemplateDescription         *String     `xmlrpc:"template_description,omitempty"`
	TemplateName                *String     `xmlrpc:"template_name,omitempty"`
	TemplatePreview             *String     `xmlrpc:"template_preview,omitempty"`
	TemplateSequence            *Int        `xmlrpc:"template_sequence,omitempty"`
	ToDelete                    *Bool       `xmlrpc:"to_delete,omitempty"`
	UserCanRead                 *Bool       `xmlrpc:"user_can_read,omitempty"`
	UserCanWrite                *Bool       `xmlrpc:"user_can_write,omitempty"`
	UserFavoriteSequence        *Int        `xmlrpc:"user_favorite_sequence,omitempty"`
	UserHasAccess               *Bool       `xmlrpc:"user_has_access,omitempty"`
	UserHasAccessParentPath     *Bool       `xmlrpc:"user_has_access_parent_path,omitempty"`
	UserHasWriteAccess          *Bool       `xmlrpc:"user_has_write_access,omitempty"`
	UserPermission              *Selection  `xmlrpc:"user_permission,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WebsitePublished            *Bool       `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                  *String     `xmlrpc:"website_url,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// KnowledgeArticles represents array of knowledge.article model.
type KnowledgeArticles []KnowledgeArticle

// KnowledgeArticleModel is the odoo model name.
const KnowledgeArticleModel = "knowledge.article"

// Many2One convert KnowledgeArticle to *Many2One.
func (ka *KnowledgeArticle) Many2One() *Many2One {
	return NewMany2One(ka.Id.Get(), "")
}

// CreateKnowledgeArticle creates a new knowledge.article model and returns its id.
func (c *Client) CreateKnowledgeArticle(ka *KnowledgeArticle) (int64, error) {
	ids, err := c.CreateKnowledgeArticles([]*KnowledgeArticle{ka})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticle creates a new knowledge.article model and returns its id.
func (c *Client) CreateKnowledgeArticles(kas []*KnowledgeArticle) ([]int64, error) {
	var vv []interface{}
	for _, v := range kas {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleModel, vv, nil)
}

// UpdateKnowledgeArticle updates an existing knowledge.article record.
func (c *Client) UpdateKnowledgeArticle(ka *KnowledgeArticle) error {
	return c.UpdateKnowledgeArticles([]int64{ka.Id.Get()}, ka)
}

// UpdateKnowledgeArticles updates existing knowledge.article records.
// All records (represented by ids) will be updated by ka values.
func (c *Client) UpdateKnowledgeArticles(ids []int64, ka *KnowledgeArticle) error {
	return c.Update(KnowledgeArticleModel, ids, ka, nil)
}

// DeleteKnowledgeArticle deletes an existing knowledge.article record.
func (c *Client) DeleteKnowledgeArticle(id int64) error {
	return c.DeleteKnowledgeArticles([]int64{id})
}

// DeleteKnowledgeArticles deletes existing knowledge.article records.
func (c *Client) DeleteKnowledgeArticles(ids []int64) error {
	return c.Delete(KnowledgeArticleModel, ids)
}

// GetKnowledgeArticle gets knowledge.article existing record.
func (c *Client) GetKnowledgeArticle(id int64) (*KnowledgeArticle, error) {
	kas, err := c.GetKnowledgeArticles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kas)[0]), nil
}

// GetKnowledgeArticles gets knowledge.article existing records.
func (c *Client) GetKnowledgeArticles(ids []int64) (*KnowledgeArticles, error) {
	kas := &KnowledgeArticles{}
	if err := c.Read(KnowledgeArticleModel, ids, nil, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowledgeArticle finds knowledge.article record by querying it with criteria.
func (c *Client) FindKnowledgeArticle(criteria *Criteria) (*KnowledgeArticle, error) {
	kas := &KnowledgeArticles{}
	if err := c.SearchRead(KnowledgeArticleModel, criteria, NewOptions().Limit(1), kas); err != nil {
		return nil, err
	}
	return &((*kas)[0]), nil
}

// FindKnowledgeArticles finds knowledge.article records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticles(criteria *Criteria, options *Options) (*KnowledgeArticles, error) {
	kas := &KnowledgeArticles{}
	if err := c.SearchRead(KnowledgeArticleModel, criteria, options, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowledgeArticleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KnowledgeArticleModel, criteria, options)
}

// FindKnowledgeArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
