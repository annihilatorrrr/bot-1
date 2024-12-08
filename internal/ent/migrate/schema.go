// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LastChannelMessagesColumns holds the columns for the "last_channel_messages" table.
	LastChannelMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "message_id", Type: field.TypeInt},
	}
	// LastChannelMessagesTable holds the schema information for the "last_channel_messages" table.
	LastChannelMessagesTable = &schema.Table{
		Name:       "last_channel_messages",
		Columns:    LastChannelMessagesColumns,
		PrimaryKey: []*schema.Column{LastChannelMessagesColumns[0]},
	}
	// PrNotificationsColumns holds the columns for the "pr_notifications" table.
	PrNotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repo_id", Type: field.TypeInt64},
		{Name: "pull_request_id", Type: field.TypeInt},
		{Name: "pull_request_title", Type: field.TypeString, Default: ""},
		{Name: "pull_request_body", Type: field.TypeString, Default: ""},
		{Name: "pull_request_author_login", Type: field.TypeString, Default: ""},
		{Name: "message_id", Type: field.TypeInt},
	}
	// PrNotificationsTable holds the schema information for the "pr_notifications" table.
	PrNotificationsTable = &schema.Table{
		Name:       "pr_notifications",
		Columns:    PrNotificationsColumns,
		PrimaryKey: []*schema.Column{PrNotificationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "prnotification_repo_id_pull_request_id",
				Unique:  true,
				Columns: []*schema.Column{PrNotificationsColumns[1], PrNotificationsColumns[2]},
			},
		},
	}
	// TelegramAccountsColumns holds the columns for the "telegram_accounts" table.
	TelegramAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "code_at", Type: field.TypeTime},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"New", "CodeSent", "Active", "Error"}, Default: "New"},
		{Name: "status", Type: field.TypeString},
		{Name: "session", Type: field.TypeBytes},
	}
	// TelegramAccountsTable holds the schema information for the "telegram_accounts" table.
	TelegramAccountsTable = &schema.Table{
		Name:       "telegram_accounts",
		Columns:    TelegramAccountsColumns,
		PrimaryKey: []*schema.Column{TelegramAccountsColumns[0]},
	}
	// TelegramChannelStatesColumns holds the columns for the "telegram_channel_states" table.
	TelegramChannelStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "channel_id", Type: field.TypeInt64},
		{Name: "pts", Type: field.TypeInt, Default: 0},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// TelegramChannelStatesTable holds the schema information for the "telegram_channel_states" table.
	TelegramChannelStatesTable = &schema.Table{
		Name:       "telegram_channel_states",
		Columns:    TelegramChannelStatesColumns,
		PrimaryKey: []*schema.Column{TelegramChannelStatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "telegram_channel_states_telegram_user_states_channels",
				Columns:    []*schema.Column{TelegramChannelStatesColumns[3]},
				RefColumns: []*schema.Column{TelegramUserStatesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "telegramchannelstate_user_id_channel_id",
				Unique:  true,
				Columns: []*schema.Column{TelegramChannelStatesColumns[3], TelegramChannelStatesColumns[1]},
			},
		},
	}
	// TelegramSessionsColumns holds the columns for the "telegram_sessions" table.
	TelegramSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "data", Type: field.TypeBytes},
	}
	// TelegramSessionsTable holds the schema information for the "telegram_sessions" table.
	TelegramSessionsTable = &schema.Table{
		Name:       "telegram_sessions",
		Columns:    TelegramSessionsColumns,
		PrimaryKey: []*schema.Column{TelegramSessionsColumns[0]},
	}
	// TelegramUserStatesColumns holds the columns for the "telegram_user_states" table.
	TelegramUserStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "qts", Type: field.TypeInt, Default: 0},
		{Name: "pts", Type: field.TypeInt, Default: 0},
		{Name: "date", Type: field.TypeInt, Default: 0},
		{Name: "seq", Type: field.TypeInt, Default: 0},
	}
	// TelegramUserStatesTable holds the schema information for the "telegram_user_states" table.
	TelegramUserStatesTable = &schema.Table{
		Name:       "telegram_user_states",
		Columns:    TelegramUserStatesColumns,
		PrimaryKey: []*schema.Column{TelegramUserStatesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LastChannelMessagesTable,
		PrNotificationsTable,
		TelegramAccountsTable,
		TelegramChannelStatesTable,
		TelegramSessionsTable,
		TelegramUserStatesTable,
	}
)

func init() {
	TelegramChannelStatesTable.ForeignKeys[0].RefTable = TelegramUserStatesTable
}
