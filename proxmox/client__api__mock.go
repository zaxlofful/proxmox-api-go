package proxmox

import (
	"context"
)

type mockClientAPI struct {
	createBackupScheduleFunc   func(ctx context.Context, params map[string]any) (BackupScheduleID, error)
	createHaRuleFunc           func(ctx context.Context, params map[string]any) error
	deleteBackupScheduleFunc   func(ctx context.Context, id BackupScheduleID) error
	deleteHaResourceFunc       func(ctx context.Context, id GuestID) error
	deleteHaRuleFunc           func(ctx context.Context, id HaRuleID) error
	getBackupScheduleConfigFunc func(ctx context.Context, id BackupScheduleID) (map[string]any, error)
	getGuestConfigFunc         func(ctx context.Context, vmr *VmRef) (map[string]any, error)
	getGuestPendingChangesFunc func(ctx context.Context, vmr *VmRef) ([]any, error)
	getGuestQemuAgentFunc      func(ctx context.Context, vmr *VmRef) (map[string]any, GuestAgentState, error)
	getHaRuleFunc              func(ctx context.Context, id HaRuleID) (map[string]any, error)
	getPoolConfigFunc          func(ctx context.Context, pool PoolName) (map[string]any, error)
	getUserConfigFunc          func(ctx context.Context, userId UserID) (map[string]any, bool, error)
	listBackupSchedulesFunc    func(ctx context.Context) ([]any, error)
	listGuestResourcesFunc     func(ctx context.Context) ([]any, error)
	listHaRulesFunc            func(ctx context.Context) ([]any, error)
	updateBackupScheduleFunc   func(ctx context.Context, id BackupScheduleID, params map[string]any) error
	updateGuestStatusFunc      func(ctx context.Context, vmr *VmRef, setStatus string, params map[string]interface{}) error
	updateHaRuleFunc           func(ctx context.Context, id HaRuleID, params map[string]any) error
}

func (m mockClientAPI) new() clientApiInterface { return &m }

func (m *mockClientAPI) panic(field string) { panic(field + " not set in mockClientAPI") }

// Interface methods

func (m *mockClientAPI) createBackupSchedule(ctx context.Context, params map[string]any) (BackupScheduleID, error) {
	if m.createBackupScheduleFunc == nil {
		m.panic("createBackupScheduleFunc")
	}
	return m.createBackupScheduleFunc(ctx, params)
}

func (m *mockClientAPI) createHaRule(ctx context.Context, params map[string]any) error {
	if m.createHaRuleFunc == nil {
		m.panic("createHaRuleFunc")
	}
	return m.createHaRuleFunc(ctx, params)
}

func (m *mockClientAPI) deleteBackupSchedule(ctx context.Context, id BackupScheduleID) error {
	if m.deleteBackupScheduleFunc == nil {
		m.panic("deleteBackupScheduleFunc")
	}
	return m.deleteBackupScheduleFunc(ctx, id)
}

func (m *mockClientAPI) deleteHaResource(ctx context.Context, id GuestID) error {
	if m.deleteHaResourceFunc == nil {
		m.panic("deleteHaResourceFunc")
	}
	return m.deleteHaResourceFunc(ctx, id)
}

func (m *mockClientAPI) deleteHaRule(ctx context.Context, id HaRuleID) error {
	if m.deleteHaRuleFunc == nil {
		m.panic("deleteHaRuleFunc")
	}
	return m.deleteHaRuleFunc(ctx, id)
}

func (m *mockClientAPI) getGuestConfig(ctx context.Context, vmr *VmRef) (vmConfig map[string]any, err error) {
	if m.getGuestConfigFunc == nil {
		m.panic("getGuestConfigFunc")
	}
	return m.getGuestConfigFunc(ctx, vmr)
}

func (m *mockClientAPI) getGuestPendingChanges(ctx context.Context, vmr *VmRef) ([]any, error) {
	if m.getGuestPendingChangesFunc == nil {
		m.panic("getGuestPendingChangesFunc")
	}
	return m.getGuestPendingChangesFunc(ctx, vmr)
}

func (m *mockClientAPI) getBackupScheduleConfig(ctx context.Context, id BackupScheduleID) (map[string]any, error) {
	if m.getBackupScheduleConfigFunc == nil {
		m.panic("getBackupScheduleConfigFunc")
	}
	return m.getBackupScheduleConfigFunc(ctx, id)
}

func (m *mockClientAPI) getGuestQemuAgent(ctx context.Context, vmr *VmRef) (map[string]any, GuestAgentState, error) {
	if m.getGuestQemuAgentFunc == nil {
		m.panic("getGuestQemuAgentFunc")
	}
	return m.getGuestQemuAgentFunc(ctx, vmr)
}

func (m *mockClientAPI) getHaRule(ctx context.Context, id HaRuleID) (haRule map[string]any, err error) {
	if m.getHaRuleFunc == nil {
		m.panic("getHaRuleFunc")
	}
	return m.getHaRuleFunc(ctx, id)
}

func (m *mockClientAPI) getPoolConfig(ctx context.Context, pool PoolName) (poolConfig map[string]any, err error) {
	if m.getPoolConfigFunc == nil {
		m.panic("getPoolConfigFunc")
	}
	return m.getPoolConfigFunc(ctx, pool)
}

func (m *mockClientAPI) getUserConfig(ctx context.Context, userId UserID) (map[string]any, bool, error) {
	if m.getUserConfigFunc == nil {
		m.panic("getUserConfigFunc")
	}
	return m.getUserConfigFunc(ctx, userId)
}

func (m *mockClientAPI) listGuestResources(ctx context.Context) ([]any, error) {
	if m.listGuestResourcesFunc == nil {
		m.panic("listGuestResourcesFunc")
	}
	return m.listGuestResourcesFunc(ctx)
}

func (m *mockClientAPI) listBackupSchedules(ctx context.Context) ([]any, error) {
	if m.listBackupSchedulesFunc == nil {
		m.panic("listBackupSchedulesFunc")
	}
	return m.listBackupSchedulesFunc(ctx)
}

func (m *mockClientAPI) listHaRules(ctx context.Context) ([]any, error) {
	if m.listHaRulesFunc == nil {
		m.panic("ListHaRulesFunc")
	}
	return m.listHaRulesFunc(ctx)
}

func (m *mockClientAPI) updateGuestStatus(ctx context.Context, vmr *VmRef, setStatus string, params map[string]interface{}) error {
	if m.updateGuestStatusFunc == nil {
		m.panic("updateGuestStatusFunc")
	}
	return m.updateGuestStatusFunc(ctx, vmr, setStatus, params)
}

func (m *mockClientAPI) updateBackupSchedule(ctx context.Context, id BackupScheduleID, params map[string]any) error {
	if m.updateBackupScheduleFunc == nil {
		m.panic("updateBackupScheduleFunc")
	}
	return m.updateBackupScheduleFunc(ctx, id, params)
}

func (m *mockClientAPI) updateHaRule(ctx context.Context, id HaRuleID, params map[string]any) error {
	if m.updateHaRuleFunc == nil {
		m.panic("updateHaRuleFunc")
	}
	return m.updateHaRuleFunc(ctx, id, params)
}
