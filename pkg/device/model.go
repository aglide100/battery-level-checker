package device

import (
	"errors"
	"time"

	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
)

type Device struct {
	DeviceInterface
	DeviceImpl
}

type DeviceInterface interface {
	GetDeviceSpec() *DeviceSpec
	GetProtoDeviceSpec() *pb_unit_device.Spec
	GetBatteryLevel() *BatteryLevel
	SetDeviceSpec(*DeviceSpec) error
	SetBatteryLevel(*BatteryLevel)
	SetDeviceId(string)
	GetDeviceId() string
	Validator() error
	Clone() *Device
	ToProto() (*pb_unit_device.Device, error)
}

type DeviceImpl struct {
	Id string `validate:"required" json:"id" example:"1"`
	BatteryLevel BatteryLevel
	Spec DeviceSpec
}

type BatteryLevel struct {	
	Time          *time.Time `validate:"required" json:"time" example:"2006-01-02 15:04:05"`
	BatteryLevel  int        `validate:"required" json:"batteryLevel" example:"50"`
	BatteryStatus string     `validate:"required" json:"batteryStatus" example:"charging"`
}
 
type DeviceSpec struct {
	Name       string `validate:"required" json:"name" example:"iphone 99xs"`
	Type       string `validate:"required" json:"type" example:"phone"`
	OS         string `validate:"required" json:"OS" example:"IOS"`
	OSversion  string `validate:"required" json:"OSversion" example:"99.192"`
	AppVersion string `validate:"required" json:"appVersion" example:"0.0.1"`
}

type Id struct {
	DeviceID   string `validate:"required" json:"deviceID" example:"f782fd74-d264-4901-8df5-c905f9df08db"`
}

func NewDevice() *Device {
	m := &Device{}

	return m
}

func Clone(d *Device) *Device {
	m := NewDevice()
	m.SetDeviceSpec(d.GetDeviceSpec())

	// m.SetDeviceSpec(d.GetDeviceSpec())
	return m
}

func NewDeviceFromProto(pbDevice *pb_unit_device.Device) (*Device, error){
	spec := &DeviceSpec{
		Name: pbDevice.Spec.Name,
		Type: pbDevice.Spec.Type,
		OS: pbDevice.Spec.OS,
		OSversion: pbDevice.Spec.OsVersion,
		AppVersion: pbDevice.Spec.AppVersion,
	}

	err := SpecValidator(spec)
	if err != nil {
		return nil, err
	}
	
	m := NewDevice()
	m.SetDeviceSpec(spec) 

	return m, nil
}
func (d *Device) SetDeviceId(id string) {
	d.Id=id
}

func (d *Device) GetDeviceId() string {
	return d.Id
}

func (d *Device) GetDeviceSpec() *DeviceSpec {
	return &d.Spec
}

func (d *Device) GetProtoDeviceSpec() *pb_unit_device.Spec {
	return &pb_unit_device.Spec{
		Name: d.Spec.Name,
		Type: d.Spec.Type,
		OS: d.Spec.OS,
		OsVersion: d.Spec.OSversion,
		AppVersion: d.Spec.AppVersion,
	}
}


func (d *Device) GetBatteryLevel() BatteryLevel {
	return d.BatteryLevel
}

func (d *Device) SetBatteryLevel(batteryLevel *BatteryLevel) {
	d.BatteryLevel = *batteryLevel
}

func (d *Device) SetDeviceSpec(spec *DeviceSpec) error {
	if spec == nil {
		return errors.New("spec is nil")
	}

	err := SpecValidator(spec)
	if err != nil {
		return err
	}
	d.Spec = *spec

	return nil
}

func (d *Device) ToProto() (*pb_unit_device.Device ,error) {
	
	pbUnit := &pb_unit_device.Device{}

	id := &pb_unit_device.ID{
		Uuid: d.Id,
	}

	spec := &pb_unit_device.Spec{
		Name: d.Spec.Name,
		Type: d.Spec.Type,
		OS: d.Spec.OS,
		OsVersion: d.Spec.OSversion,
		AppVersion: d.Spec.AppVersion,
	}

	pbUnit.Id = id
	pbUnit.Spec = spec

	return pbUnit, nil
}