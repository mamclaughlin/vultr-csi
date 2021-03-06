package driver

import (
	"context"

	"github.com/vultr/govultr"
)

func newFakeClient() *govultr.Client {
	fakeInstance := FakeInstance{client: nil}
	fakeBlockStorage := fakeBS{client: nil}

	return &govultr.Client{
		Server:       &fakeInstance,
		BlockStorage: &fakeBlockStorage,
	}
}

func newFakeBS() *govultr.BlockStorage {
	return &govultr.BlockStorage{
		BlockStorageID: "342512",
		DateCreated:    "",
		CostPerMonth:   "10",
		Status:         "active",
		SizeGB:         10,
		RegionID:       1,
		InstanceID:     "123456",
		Label:          "test-bs",
	}
}

type fakeBS struct {
	client *govultr.Client
}

func (f *fakeBS) Attach(ctx context.Context, blockID, InstanceID, liveAttach string) error {
	panic("implement me")
}

func (f *fakeBS) Create(ctx context.Context, regionID, size int, label string) (*govultr.BlockStorage, error) {
	bsVolume := newFakeBS()
	return bsVolume, nil
}

func (f *fakeBS) Delete(ctx context.Context, blockID string) error {
	panic("implement me")
}

func (f *fakeBS) Detach(ctx context.Context, blockID, liveDetach string) error {
	list, err := f.List(ctx)
	if err != nil {
		return err
	}

	for _, volume := range list {
		if volume.BlockStorageID == blockID {
			volume.InstanceID = ""
		}
	}

	return nil
}

func (f *fakeBS) SetLabel(ctx context.Context, blockID, label string) error {
	panic("implement me")
}

func (f *fakeBS) List(ctx context.Context) ([]govultr.BlockStorage, error) {
	return []govultr.BlockStorage{
		{
			BlockStorageID: "342512",
			DateCreated:    "",
			CostPerMonth:   "10",
			Status:         "active",
			SizeGB:         10,
			RegionID:       1,
			InstanceID:     "123456",
			Label:          "test-bs",
		},
		{
			BlockStorageID: "342513",
			DateCreated:    "",
			CostPerMonth:   "20",
			Status:         "active",
			SizeGB:         20,
			RegionID:       1,
			InstanceID:     "123457",
			Label:          "test-bs2",
		},
	}, nil
}

func (f *fakeBS) Get(ctx context.Context, blockID string) (*govultr.BlockStorage, error) {
	bsVolume := newFakeBS()
	return bsVolume, nil
}

func (f *fakeBS) Resize(ctx context.Context, blockID string, size int) error {
	panic("implement me")
}

type FakeInstance struct {
	client *govultr.Client
}

func (f *FakeInstance) ChangeApp(ctx context.Context, instanceID, appID string) error {
	panic("implement me")
}

func (f *FakeInstance) ListApps(ctx context.Context, instanceID string) ([]govultr.Application, error) {
	panic("implement me")
}

func (f *FakeInstance) AppInfo(ctx context.Context, instanceID string) (*govultr.AppInfo, error) {
	panic("implement me")
}

func (f *FakeInstance) EnableBackup(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) DisableBackup(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) GetBackupSchedule(ctx context.Context, instanceID string) (*govultr.BackupSchedule, error) {
	panic("implement me")
}

func (f *FakeInstance) SetBackupSchedule(ctx context.Context, instanceID string, backup *govultr.BackupSchedule) error {
	panic("implement me")
}

func (f *FakeInstance) RestoreBackup(ctx context.Context, instanceID, backupID string) error {
	panic("implement me")
}

func (f *FakeInstance) RestoreSnapshot(ctx context.Context, instanceID, snapshotID string) error {
	panic("implement me")
}

func (f *FakeInstance) SetLabel(ctx context.Context, instanceID, label string) error {
	panic("implement me")
}

func (f *FakeInstance) SetTag(ctx context.Context, instanceID, tag string) error {
	panic("implement me")
}

func (f *FakeInstance) Neighbors(ctx context.Context, instanceID string) ([]int, error) {
	panic("implement me")
}

func (f *FakeInstance) EnablePrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	panic("implement me")
}

func (f *FakeInstance) DisablePrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	panic("implement me")
}

func (f *FakeInstance) ListPrivateNetworks(ctx context.Context, instanceID string) ([]govultr.PrivateNetwork, error) {
	panic("implement me")
}

func (f *FakeInstance) ListUpgradePlan(ctx context.Context, instanceID string) ([]int, error) {
	panic("implement me")
}

func (f *FakeInstance) UpgradePlan(ctx context.Context, instanceID, vpsPlanID string) error {
	panic("implement me")
}

func (f *FakeInstance) ListOS(ctx context.Context, instanceID string) ([]govultr.OS, error) {
	panic("implement me")
}

func (f *FakeInstance) ChangeOS(ctx context.Context, instanceID, osID string) error {
	panic("implement me")
}

func (f *FakeInstance) IsoAttach(ctx context.Context, instanceID, isoID string) error {
	panic("implement me")
}

func (f *FakeInstance) IsoDetach(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) IsoStatus(ctx context.Context, instanceID string) (*govultr.ServerIso, error) {
	panic("implement me")
}

func (f *FakeInstance) SetFirewallGroup(ctx context.Context, instanceID, firewallGroupID string) error {
	panic("implement me")
}

func (f *FakeInstance) GetUserData(ctx context.Context, instanceID string) (*govultr.UserData, error) {
	panic("implement me")
}

func (f *FakeInstance) SetUserData(ctx context.Context, instanceID, userData string) error {
	panic("implement me")
}

func (f *FakeInstance) IPV4Info(ctx context.Context, instanceID string, public bool) ([]govultr.IPV4, error) {
	panic("implement me")
}

func (f *FakeInstance) IPV6Info(ctx context.Context, instanceID string) ([]govultr.IPV6, error) {
	panic("implement me")
}

func (f *FakeInstance) AddIPV4(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) DestroyIPV4(ctx context.Context, instanceID, ip string) error {
	panic("implement me")
}

func (f *FakeInstance) EnableIPV6(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Bandwidth(ctx context.Context, instanceID string) ([]map[string]string, error) {
	panic("implement me")
}

func (f *FakeInstance) ListReverseIPV6(ctx context.Context, instanceID string) ([]govultr.ReverseIPV6, error) {
	panic("implement me")
}

func (f *FakeInstance) SetDefaultReverseIPV4(ctx context.Context, instanceID, ip string) error {
	panic("implement me")
}

func (f *FakeInstance) DeleteReverseIPV6(ctx context.Context, instanceID, ip string) error {
	panic("implement me")
}

func (f *FakeInstance) SetReverseIPV4(ctx context.Context, instanceID, ipv4, entry string) error {
	panic("implement me")
}

func (f *FakeInstance) SetReverseIPV6(ctx context.Context, instanceID, ipv6, entry string) error {
	panic("implement me")
}

func (f *FakeInstance) Start(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Halt(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Reboot(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Reinstall(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Delete(ctx context.Context, instanceID string) error {
	panic("implement me")
}

func (f *FakeInstance) Create(ctx context.Context, regionID, vpsPlanID, osID int, options *govultr.ServerOptions) (*govultr.Server, error) {
	panic("implement me 234")
}

func (f *FakeInstance) List(ctx context.Context) ([]govultr.Server, error) {
	panic("implement me")
}

func (f *FakeInstance) ListByLabel(ctx context.Context, label string) ([]govultr.Server, error) {
	return []govultr.Server{
		{
			InstanceID:  "576965",
			MainIP:      "149.28.225.110",
			VPSCpus:     "4",
			Location:    "New Jersey",
			RegionID:    "1",
			Status:      "running",
			NetmaskV4:   "255.255.254.0",
			GatewayV4:   "149.28.224.1",
			PowerStatus: "",
			ServerState: "",
			PlanID:      "204",
			Label:       "csi-test",
			InternalIP:  "10.1.95.4",
		},
	}, nil
}

func (f *FakeInstance) ListByMainIP(ctx context.Context, mainIP string) ([]govultr.Server, error) {
	panic("implement me")
}

func (f *FakeInstance) ListByTag(ctx context.Context, tag string) ([]govultr.Server, error) {
	panic("implement me")
}

func (f *FakeInstance) GetServer(ctx context.Context, instanceID string) (*govultr.Server, error) {
	return &govultr.Server{
		InstanceID:  "576965",
		MainIP:      "149.28.225.110",
		VPSCpus:     "4",
		Location:    "New Jersey",
		RegionID:    "1",
		Status:      "running",
		NetmaskV4:   "255.255.254.0",
		GatewayV4:   "149.28.224.1",
		PowerStatus: "",
		ServerState: "",
		PlanID:      "204",
		Label:       "csi-test",
		InternalIP:  "10.1.95.4",
	}, nil
}
