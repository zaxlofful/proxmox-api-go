package proxmox

type ClientNew struct {
	ApiToken       ApiTokenInterface
	BackupSchedule BackupScheduleInterface
	Group          GroupInterface
	Pool           PoolInterface
	Snapshot       SnapshotInterface
	User           UserInterface
}
