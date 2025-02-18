package odoo

// AccountOnlineLink represents account.online.link model.
type AccountOnlineLink struct {
	AccessToken                 *String     `xmlrpc:"access_token,omitempty"`
	AccountOnlineAccountIds     *Relation   `xmlrpc:"account_online_account_ids,omitempty"`
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
	AutoSync                    *Bool       `xmlrpc:"auto_sync,omitempty"`
	ClientId                    *String     `xmlrpc:"client_id,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty"`
	ConnectionStateDetails      interface{} `xmlrpc:"connection_state_details,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	ExpiringSynchronizationDate *Time       `xmlrpc:"expiring_synchronization_date,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	HasUnlinkedAccounts         *Bool       `xmlrpc:"has_unlinked_accounts,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	JournalIds                  *Relation   `xmlrpc:"journal_ids,omitempty"`
	LastRefresh                 *Time       `xmlrpc:"last_refresh,omitempty"`
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
	NextRefresh                 *Time       `xmlrpc:"next_refresh,omitempty"`
	ProviderData                *String     `xmlrpc:"provider_data,omitempty"`
	ProviderType                *String     `xmlrpc:"provider_type,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	RefreshToken                *String     `xmlrpc:"refresh_token,omitempty"`
	State                       *Selection  `xmlrpc:"state,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountOnlineLinks represents array of account.online.link model.
type AccountOnlineLinks []AccountOnlineLink

// AccountOnlineLinkModel is the odoo model name.
const AccountOnlineLinkModel = "account.online.link"

// Many2One convert AccountOnlineLink to *Many2One.
func (aol *AccountOnlineLink) Many2One() *Many2One {
	return NewMany2One(aol.Id.Get(), "")
}

// CreateAccountOnlineLink creates a new account.online.link model and returns its id.
func (c *Client) CreateAccountOnlineLink(aol *AccountOnlineLink) (int64, error) {
	ids, err := c.CreateAccountOnlineLinks([]*AccountOnlineLink{aol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineLink creates a new account.online.link model and returns its id.
func (c *Client) CreateAccountOnlineLinks(aols []*AccountOnlineLink) ([]int64, error) {
	var vv []interface{}
	for _, v := range aols {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineLinkModel, vv, nil)
}

// UpdateAccountOnlineLink updates an existing account.online.link record.
func (c *Client) UpdateAccountOnlineLink(aol *AccountOnlineLink) error {
	return c.UpdateAccountOnlineLinks([]int64{aol.Id.Get()}, aol)
}

// UpdateAccountOnlineLinks updates existing account.online.link records.
// All records (represented by ids) will be updated by aol values.
func (c *Client) UpdateAccountOnlineLinks(ids []int64, aol *AccountOnlineLink) error {
	return c.Update(AccountOnlineLinkModel, ids, aol, nil)
}

// DeleteAccountOnlineLink deletes an existing account.online.link record.
func (c *Client) DeleteAccountOnlineLink(id int64) error {
	return c.DeleteAccountOnlineLinks([]int64{id})
}

// DeleteAccountOnlineLinks deletes existing account.online.link records.
func (c *Client) DeleteAccountOnlineLinks(ids []int64) error {
	return c.Delete(AccountOnlineLinkModel, ids)
}

// GetAccountOnlineLink gets account.online.link existing record.
func (c *Client) GetAccountOnlineLink(id int64) (*AccountOnlineLink, error) {
	aols, err := c.GetAccountOnlineLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aols)[0]), nil
}

// GetAccountOnlineLinks gets account.online.link existing records.
func (c *Client) GetAccountOnlineLinks(ids []int64) (*AccountOnlineLinks, error) {
	aols := &AccountOnlineLinks{}
	if err := c.Read(AccountOnlineLinkModel, ids, nil, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAccountOnlineLink finds account.online.link record by querying it with criteria.
func (c *Client) FindAccountOnlineLink(criteria *Criteria) (*AccountOnlineLink, error) {
	aols := &AccountOnlineLinks{}
	if err := c.SearchRead(AccountOnlineLinkModel, criteria, NewOptions().Limit(1), aols); err != nil {
		return nil, err
	}
	return &((*aols)[0]), nil
}

// FindAccountOnlineLinks finds account.online.link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinks(criteria *Criteria, options *Options) (*AccountOnlineLinks, error) {
	aols := &AccountOnlineLinks{}
	if err := c.SearchRead(AccountOnlineLinkModel, criteria, options, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAccountOnlineLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOnlineLinkModel, criteria, options)
}

// FindAccountOnlineLinkId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
