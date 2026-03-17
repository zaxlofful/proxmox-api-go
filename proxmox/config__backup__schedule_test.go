package proxmox

import (
	"context"
	"errors"
	"testing"

	"github.com/Telmate/proxmox-api-go/internal/mockServer"
	"github.com/Telmate/proxmox-api-go/internal/util"
	"github.com/stretchr/testify/require"
)

func Test_backupScheduleClient_Create(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    ConfigBackupSchedule
		output   BackupScheduleID
		requests []mockServer.Request
		err      error
	}{
		{name: `minimal`,
			input: ConfigBackupSchedule{
				Schedule: util.Pointer("daily 2:00"),
				Storage:  util.Pointer(StorageName("local")),
			},
			output: "backup-12345678",
			requests: mockServer.Append(
				mockServer.RequestsPostResponse("/cluster/backup",
					map[string]any{
						"schedule": "daily 2:00",
						"storage":  "local",
					},
					[]byte(`{"data":"backup-12345678"}`),
				))},
		{name: `all fields`,
			input: ConfigBackupSchedule{
				Schedule:         util.Pointer("sun 3:00"),
				Storage:          util.Pointer(StorageName("nfs-backup")),
				Mode:             util.Pointer("snapshot"),
				Compress:         util.Pointer("zstd"),
				VMIDs:            util.Pointer("100,101"),
				Pool:             util.Pointer(PoolName("production")),
				Enabled:          util.Pointer(true),
				Node:             util.Pointer(NodeName("pve")),
				Mailto:           util.Pointer("admin@example.com"),
				MailNotification: util.Pointer("failure"),
				NotesTemplate:    util.Pointer("{{guestname}}"),
				PruneBackups:     util.Pointer("keep-last=3"),
				Protected:        util.Pointer(false),
				Comment:          util.Pointer("nightly backup"),
			},
			output: "backup-abcdef01",
			requests: mockServer.Append(
				mockServer.RequestsPostResponse("/cluster/backup",
					map[string]any{
						"schedule":        "sun 3:00",
						"storage":         "nfs-backup",
						"mode":            "snapshot",
						"compress":        "zstd",
						"vmid":            "100,101",
						"pool":            "production",
						"enabled":         "1",
						"node":            "pve",
						"mailto":          "admin@example.com",
						"mailnotification": "failure",
						"notes-template":  "{{guestname}}",
						"prune-backups":   "keep-last=3",
						"protected":       "0",
						"comment":         "nightly backup",
					},
					[]byte(`{"data":"backup-abcdef01"}`),
				))},
		{name: `error - schedule required`,
			input: ConfigBackupSchedule{
				Storage: util.Pointer(StorageName("local")),
			},
			err: errors.New(ConfigBackupSchedule_Error_ScheduleRequired)},
		{name: `error - storage required`,
			input: ConfigBackupSchedule{
				Schedule: util.Pointer("daily 2:00"),
			},
			err: errors.New(ConfigBackupSchedule_Error_StorageRequired)},
		{name: `error - invalid mode`,
			input: ConfigBackupSchedule{
				Schedule: util.Pointer("daily 2:00"),
				Storage:  util.Pointer(StorageName("local")),
				Mode:     util.Pointer("invalid"),
			},
			err: errors.New("error the value of key (mode) must be one of snapshot,suspend,stop")},
		{name: `error - invalid compress`,
			input: ConfigBackupSchedule{
				Schedule: util.Pointer("daily 2:00"),
				Storage:  util.Pointer(StorageName("local")),
				Compress: util.Pointer("bzip2"),
			},
			err: errors.New("error the value of key (compress) must be one of 0,gzip,lzo,zstd")},
		{name: `error - invalid mailnotification`,
			input: ConfigBackupSchedule{
				Schedule:         util.Pointer("daily 2:00"),
				Storage:          util.Pointer(StorageName("local")),
				MailNotification: util.Pointer("never"),
			},
			err: errors.New("error the value of key (mailnotification) must be one of always,failure")},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			id, err := c.New().BackupSchedule.Create(context.Background(), test.input)
			require.Equal(t, test.err, err)
			if test.err == nil {
				require.Equal(t, test.output, id)
			}
			server.Clear(t)
		})
	}
}

func Test_backupScheduleClient_Delete(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		id       BackupScheduleID
		output   bool
		requests []mockServer.Request
		err      error
	}{
		{name: `exists`,
			id:     "backup-12345678",
			output: true,
			requests: mockServer.Append(
				mockServer.RequestsDelete("/cluster/backup/backup-12345678", nil),
			)},
		{name: `does not exist`,
			id:     "backup-00000000",
			output: false,
			requests: mockServer.Append(
				mockServer.RequestsErrorHandled("/cluster/backup/backup-00000000", mockServer.DELETE, mockServer.JsonError(500, map[string]any{
					"message": "backup job 'backup-00000000' does not exist",
				})),
			)},
		{name: `error - empty id`,
			id:  "",
			err: errors.New(BackupScheduleID_Error_Empty)},
		{name: `API error`,
			id:       "backup-12345678",
			requests: mockServer.RequestsError("/cluster/backup/backup-12345678", mockServer.DELETE, 500, 3),
			err:      errors.New(mockServer.InternalServerError)},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			deleted, err := c.New().BackupSchedule.Delete(context.Background(), test.id)
			require.Equal(t, test.err, err)
			if test.err == nil {
				require.Equal(t, test.output, deleted)
			}
			server.Clear(t)
		})
	}
}

func Test_backupScheduleClient_Exists(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		id       BackupScheduleID
		output   bool
		requests []mockServer.Request
		err      error
	}{
		{name: `exists`,
			id:     "backup-12345678",
			output: true,
			requests: mockServer.Append(
				mockServer.RequestsGetJson("/cluster/backup/backup-12345678", map[string]any{
					"data": map[string]any{
						"id":       "backup-12345678",
						"schedule": "daily 2:00",
						"storage":  "local",
					},
				}))},
		{name: `does not exist`,
			id:     "backup-00000000",
			output: false,
			requests: mockServer.Append(
				mockServer.RequestsErrorHandled("/cluster/backup/backup-00000000", mockServer.GET, mockServer.JsonError(500, map[string]any{
					"message": "backup job 'backup-00000000' does not exist",
				})))},
		{name: `error - empty id`,
			id:  "",
			err: errors.New(BackupScheduleID_Error_Empty)},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			exists, err := c.New().BackupSchedule.Exists(context.Background(), test.id)
			require.Equal(t, test.err, err)
			if test.err == nil {
				require.Equal(t, test.output, exists)
			}
			server.Clear(t)
		})
	}
}

func Test_backupScheduleClient_List(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		output   []ConfigBackupSchedule
		requests []mockServer.Request
		err      error
	}{
		{name: `empty list`,
			output:   []ConfigBackupSchedule{},
			requests: mockServer.Append(mockServer.RequestsGetJson("/cluster/backup", map[string]any{"data": []any{}}))},
		{name: `single item`,
			output: []ConfigBackupSchedule{
				{
					ID:               "backup-12345678",
					Schedule:         util.Pointer("daily 2:00"),
					Storage:          util.Pointer(StorageName("local")),
					Mode:             util.Pointer(""),
					Compress:         util.Pointer(""),
					VMIDs:            util.Pointer(""),
					All:              util.Pointer(false),
					Pool:             util.Pointer(PoolName("")),
					Enabled:          util.Pointer(true),
					Node:             util.Pointer(NodeName("")),
					Mailto:           util.Pointer(""),
					MailNotification: util.Pointer(""),
					NotesTemplate:    util.Pointer(""),
					PruneBackups:     util.Pointer(""),
					Protected:        util.Pointer(false),
					Comment:          util.Pointer(""),
				},
			},
			requests: mockServer.Append(mockServer.RequestsGetJson("/cluster/backup", map[string]any{
				"data": []any{
					map[string]any{
						"id":       "backup-12345678",
						"schedule": "daily 2:00",
						"storage":  "local",
					}}}))},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			raw, err := c.New().BackupSchedule.List(context.Background())
			require.Equal(t, test.err, err)
			if test.err == nil {
				items := raw.AsArray()
				require.Equal(t, len(test.output), len(items))
				for i := range items {
					require.Equal(t, test.output[i], items[i].Get())
				}
			}
			server.Clear(t)
		})
	}
}

func Test_backupScheduleClient_Read(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		id       BackupScheduleID
		output   ConfigBackupSchedule
		requests []mockServer.Request
		err      error
	}{
		{name: `minimal`,
			id: "backup-12345678",
			output: ConfigBackupSchedule{
				ID:               "backup-12345678",
				Schedule:         util.Pointer("daily 2:00"),
				Storage:          util.Pointer(StorageName("local")),
				Mode:             util.Pointer(""),
				Compress:         util.Pointer(""),
				VMIDs:            util.Pointer(""),
				All:              util.Pointer(false),
				Pool:             util.Pointer(PoolName("")),
				Enabled:          util.Pointer(true),
				Node:             util.Pointer(NodeName("")),
				Mailto:           util.Pointer(""),
				MailNotification: util.Pointer(""),
				NotesTemplate:    util.Pointer(""),
				PruneBackups:     util.Pointer(""),
				Protected:        util.Pointer(false),
				Comment:          util.Pointer(""),
			},
			requests: mockServer.Append(
				mockServer.RequestsGetJson("/cluster/backup/backup-12345678", map[string]any{
					"data": map[string]any{
						"id":       "backup-12345678",
						"schedule": "daily 2:00",
						"storage":  "local",
					}}))},
		{name: `all fields`,
			id: "backup-abcdef01",
			output: ConfigBackupSchedule{
				ID:               "backup-abcdef01",
				Schedule:         util.Pointer("sun 3:00"),
				Storage:          util.Pointer(StorageName("nfs-backup")),
				Mode:             util.Pointer("snapshot"),
				Compress:         util.Pointer("zstd"),
				VMIDs:            util.Pointer("100,101"),
				All:              util.Pointer(false),
				Pool:             util.Pointer(PoolName("production")),
				Enabled:          util.Pointer(true),
				Node:             util.Pointer(NodeName("pve")),
				Mailto:           util.Pointer("admin@example.com"),
				MailNotification: util.Pointer("failure"),
				NotesTemplate:    util.Pointer("{{guestname}}"),
				PruneBackups:     util.Pointer("keep-last=3"),
				Protected:        util.Pointer(false),
				Comment:          util.Pointer("nightly backup"),
			},
			requests: mockServer.Append(
				mockServer.RequestsGetJson("/cluster/backup/backup-abcdef01", map[string]any{
					"data": map[string]any{
						"id":              "backup-abcdef01",
						"schedule":        "sun 3:00",
						"storage":         "nfs-backup",
						"mode":            "snapshot",
						"compress":        "zstd",
						"vmid":            "100,101",
						"pool":            "production",
						"enabled":         float64(1),
						"node":            "pve",
						"mailto":          "admin@example.com",
						"mailnotification": "failure",
						"notes-template":  "{{guestname}}",
						"prune-backups":   "keep-last=3",
						"comment":         "nightly backup",
					}}))},
		{name: `error - empty id`,
			id:  "",
			err: errors.New(BackupScheduleID_Error_Empty)},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			raw, err := c.New().BackupSchedule.Read(context.Background(), test.id)
			require.Equal(t, test.err, err)
			if test.err == nil {
				require.Equal(t, test.output, raw.Get())
			}
			server.Clear(t)
		})
	}
}

func Test_backupScheduleClient_Update(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    ConfigBackupSchedule
		requests []mockServer.Request
		err      error
	}{
		{name: `minimal - only required fields`,
			input: ConfigBackupSchedule{
				ID:       "backup-12345678",
				Schedule: util.Pointer("daily 3:00"),
				Storage:  util.Pointer(StorageName("local")),
			},
			requests: mockServer.Append(
				mockServer.RequestsPut("/cluster/backup/backup-12345678", map[string]any{
					"schedule": "daily 3:00",
					"storage":  "local",
					"mode":     "",
					"compress": "",
					"vmid":     "",
					"pool":     "",
					"node":     "",
					"mailto":   "",
					"mailnotification": "",
					"notes-template":  "",
					"prune-backups":   "",
					"comment":         "",
				}))},
		{name: `error - empty id`,
			input: ConfigBackupSchedule{
				Schedule: util.Pointer("daily 2:00"),
				Storage:  util.Pointer(StorageName("local")),
			},
			err: errors.New(BackupScheduleID_Error_Empty)},
	}
	server, c := testMockServerInit(t)
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			server.Set(test.requests, t)
			err := c.New().BackupSchedule.Update(context.Background(), test.input)
			require.Equal(t, test.err, err)
			server.Clear(t)
		})
	}
}

func Test_BackupScheduleID_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input BackupScheduleID
		err   error
	}{
		{name: `valid`, input: "backup-12345678"},
		{name: `empty`, input: "", err: errors.New(BackupScheduleID_Error_Empty)},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			require.Equal(t, test.err, test.input.Validate())
		})
	}
}

func Test_rawBackupSchedule_GetEnabled(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  map[string]any
		output bool
	}{
		{name: `true - float64`, input: map[string]any{"enabled": float64(1)}, output: true},
		{name: `false - float64`, input: map[string]any{"enabled": float64(0)}, output: false},
		{name: `true - bool`, input: map[string]any{"enabled": true}, output: true},
		{name: `false - bool`, input: map[string]any{"enabled": false}, output: false},
		{name: `missing key - defaults to true`, input: map[string]any{}, output: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			r := &rawBackupSchedule{a: test.input}
			require.Equal(t, test.output, r.GetEnabled())
		})
	}
}
