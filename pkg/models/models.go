package models

import (
	"time"
)

type User struct {
	ID                    string    `json:"id" yaml:"id"`
	CreatedAt             time.Time `json:"createdAt" yaml:"createdAt"`
	Email                 string    `json:"email" yaml:"email"`
	FirstName             string    `json:"firstName" yaml:"firstName"`
	LastName              string    `json:"lastName" yaml:"lastName"`
	Company               string    `json:"company" yaml:"company"`
	RegistrationCompleted bool      `json:"registrationCompleted" yaml:"registrationCompleted"`
	SuperAdmin            bool      `json:"superAdmin" yaml:"superAdmin"`
}

type RegistrationToken struct {
	ID        string    `json:"id" yaml:"id"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	UserID    string    `json:"userId" yaml:"userId"`
}

type PasswordRecoveryToken struct {
	ID        string    `json:"id" yaml:"id"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt" yaml:"expiresAt"`
	UserID    string    `json:"userId" yaml:"userId"`
}

type Session struct {
	ID        string    `json:"id" yaml:"id"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	UserID    string    `json:"userId" yaml:"userId"`
}

type UserAccessKey struct {
	ID          string    `json:"id" yaml:"id"`
	CreatedAt   time.Time `json:"createdAt" yaml:"createdAt"`
	UserID      string    `json:"userId" yaml:"userId"`
	Description string    `json:"description" yaml:"description"`
}

type UserAccessKeyWithValue struct {
	UserAccessKey
	Value string `json:"value" yaml:"value"`
}

type Project struct {
	ID            string    `json:"id" yaml:"id"`
	CreatedAt     time.Time `json:"createdAt" yaml:"createdAt"`
	Name          string    `json:"name" yaml:"name"`
	DatadogAPIKey *string   `json:"datadogApiKey" yaml:"datadogApiKey"`
}

type ProjectDeviceCounts struct {
	AllCount int `json:"allCount" yaml:"allCount"`
}

type ProjectApplicationCounts struct {
	AllCount int `json:"allCount" yaml:"allCount"`
}

type Role struct {
	ID          string    `json:"id" yaml:"id"`
	CreatedAt   time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID   string    `json:"projectId" yaml:"projectId"`
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	Config      string    `json:"config" yaml:"config"`
}

type Membership struct {
	UserID    string    `json:"userId" yaml:"userId"`
	ProjectID string    `json:"projectId" yaml:"projectId"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
}

type MembershipRoleBinding struct {
	UserID    string    `json:"userId" yaml:"userId"`
	RoleID    string    `json:"roleId" yaml:"roleId"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID string    `json:"projectId" yaml:"projectId"`
}

type ServiceAccount struct {
	ID          string    `json:"id" yaml:"id"`
	CreatedAt   time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID   string    `json:"projectId" yaml:"projectId"`
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
}

type ServiceAccountAccessKey struct {
	ID               string    `json:"id" yaml:"id"`
	CreatedAt        time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID        string    `json:"projectId" yaml:"projectId"`
	ServiceAccountID string    `json:"serviceAccountId" yaml:"serviceAccountId"`
	Description      string    `json:"description" yaml:"description"`
}

type ServiceAccountAccessKeyWithValue struct {
	ServiceAccountAccessKey
	Value string `json:"value" yaml:"value"`
}

type ServiceAccountRoleBinding struct {
	ServiceAccountID string    `json:"serviceAccountId" yaml:"serviceAccountId"`
	RoleID           string    `json:"roleId" yaml:"roleId"`
	CreatedAt        time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID        string    `json:"projectId" yaml:"projectId"`
}

type Device struct {
	ID                  string            `json:"id" yaml:"id"`
	CreatedAt           time.Time         `json:"createdAt" yaml:"createdAt"`
	ProjectID           string            `json:"projectId" yaml:"projectId"`
	Name                string            `json:"name" yaml:"name"`
	RegistrationTokenID *string           `json:"registrationTokenId" yaml:"registrationTokenId"`
	DesiredAgentSpec    string            `json:"desiredAgentSpec" yaml:"desiredAgentSpec"`
	DesiredAgentVersion string            `json:"desiredAgentVersion" yaml:"desiredAgentVersion"`
	Info                DeviceInfo        `json:"info" yaml:"info"`
	LastSeenAt          time.Time         `json:"lastSeenAt" yaml:"lastSeenAt"`
	Status              DeviceStatus      `json:"status" yaml:"status"`
	Labels              map[string]string `json:"labels" yaml:"labels"`
}

type DeviceStatus string

const (
	DeviceStatusOnline  = DeviceStatus("online")
	DeviceStatusOffline = DeviceStatus("offline")
)

type DeviceRegistrationToken struct {
	ID               string            `json:"id" yaml:"id"`
	CreatedAt        time.Time         `json:"createdAt" yaml:"createdAt"`
	ProjectID        string            `json:"projectId" yaml:"projectId"`
	MaxRegistrations *int              `json:"maxRegistrations" yaml:"maxRegistrations"`
	Name             string            `json:"name" yaml:"name"`
	Description      string            `json:"description" yaml:"description"`
	Labels           map[string]string `json:"labels" yaml:"labels"`
}

type DevicesRegisteredWithTokenCount struct {
	AllCount int `json:"allCount" yaml:"allCount"`
}

type DeviceAccessKey struct {
	ID        string    `json:"id" yaml:"id"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID string    `json:"projectId" yaml:"projectId"`
	DeviceID  string    `json:"deviceId" yaml:"deviceId"`
}

type Application struct {
	ID             string    `json:"id" yaml:"id"`
	CreatedAt      time.Time `json:"createdAt" yaml:"createdAt"`
	ProjectID      string    `json:"projectId" yaml:"projectId"`
	Name           string    `json:"name" yaml:"name"`
	Description    string    `json:"description" yaml:"description"`
	SchedulingRule Query     `json:"schedulingRule" yaml:"schedulingRule"`
}

type ApplicationDeviceCounts struct {
	AllCount int `json:"allCount" yaml:"allCount"`
}

type Release struct {
	ID                        string             `json:"id" yaml:"id"`
	CreatedAt                 time.Time          `json:"createdAt" yaml:"createdAt"`
	ProjectID                 string             `json:"projectId" yaml:"projectId"`
	ApplicationID             string             `json:"applicationId" yaml:"applicationId"`
	Config                    map[string]Service `json:"config" yaml:"config"`
	RawConfig                 string             `json:"rawConfig" yaml:"rawConfig"`
	CreatedByUserID           *string            `json:"createdByUserId" yaml:"createdByUserId"`
	CreatedByServiceAccountID *string            `json:"createdByServiceAccountId" yaml:"createdByServiceAccountId"`
}

type ReleaseDeviceCounts struct {
	AllCount int `json:"allCount" yaml:"allCount"`
}

type DeviceApplicationStatus struct {
	ProjectID        string `json:"projectId" yaml:"projectId"`
	DeviceID         string `json:"deviceId" yaml:"deviceId"`
	ApplicationID    string `json:"applicationId" yaml:"applicationId"`
	CurrentReleaseID string `json:"currentReleaseId" yaml:"currentReleaseId"`
}

type DeviceServiceStatus struct {
	ProjectID        string `json:"projectId" yaml:"projectId"`
	DeviceID         string `json:"deviceId" yaml:"deviceId"`
	ApplicationID    string `json:"applicationId" yaml:"applicationId"`
	Service          string `json:"service" yaml:"service"`
	CurrentReleaseID string `json:"currentReleaseId" yaml:"currentReleaseId"`
}

type MetricTargetConfig struct {
	ID        string           `json:"id" yaml:"id"`
	CreatedAt time.Time        `json:"createdAt" yaml:"createdAt"`
	ProjectID string           `json:"projectId" yaml:"projectId"`
	Type      MetricTargetType `json:"type" yaml:"type"`
	Configs   []MetricConfig   `json:"configs" yaml:"configs"`
}

type MetricTargetType string

const (
	MetricServiceTargetType MetricTargetType = "service"
	MetricHostTargetType    MetricTargetType = "host"
	MetricStateTargetType   MetricTargetType = "state"
)

type MetricConfig struct {
	Params  *ServiceMetricParams `json:"params,omitempty" json:"yaml,omitempty"`
	Metrics []Metric             `json:"metrics" yaml:"metrics"`
}

type ServiceMetricParams struct {
	ApplicationID string `json:"applicationId" yaml:"applicationId"`
	Service       string `json:"service" yaml:"service"`
}

type Metric struct {
	Metric string   `json:"metric" yaml:"metric"`
	Labels []string `json:"labels" yaml:"labels"`
	Tags   []string `json:"tags" yaml:"tags"`
}

type MembershipFull1 struct {
	Membership
	User    User        `json:"user" yaml:"user"`
	Project ProjectFull `json:"project" yaml:"project"`
}

type ProjectFull struct {
	Project
	DeviceCounts      ProjectDeviceCounts      `json:"deviceCounts" yaml:"deviceCounts"`
	ApplicationCounts ProjectApplicationCounts `json:"applicationCounts" yaml:"applicationCounts"`
}

type MembershipFull2 struct {
	Membership
	User  User   `json:"user" yaml:"user"`
	Roles []Role `json:"roles" yaml:"roles"`
}

type ServiceAccountFull struct {
	ServiceAccount
	Roles []Role `json:"roles" yaml:"roles"`
}

type DeviceFull struct {
	Device
	ApplicationStatusInfo []DeviceApplicationStatusInfo `json:"applicationStatusInfo" yaml:"applicationStatusInfo"`
}

type DeviceApplicationStatusInfo struct {
	Application       Application              `json:"application" yaml:"application"`
	ApplicationStatus *DeviceApplicationStatus `json:"applicationStatus" yaml:"applicationStatus"`
	ServiceStatuses   []DeviceServiceStatus    `json:"serviceStatuses" yaml:"serviceStatuses"`
}

type ApplicationFull1 struct {
	Application
	LatestRelease *Release                `json:"latestRelease" yaml:"latestRelease"`
	DeviceCounts  ApplicationDeviceCounts `json:"deviceCounts" yaml:"deviceCounts"`
}

type DeviceRegistrationTokenFull struct {
	DeviceRegistrationToken
	DeviceCounts DevicesRegisteredWithTokenCount `json:"deviceCounts" yaml:"deviceCounts"`
}

type ReleaseFull struct {
	Release
	CreatedByUser           *User               `json:"createdByUser" yaml:"createdByUser"`
	CreatedByServiceAccount *ServiceAccount     `json:"createdByServiceAccount" yaml:"createdByServiceAccount"`
	DeviceCounts            ReleaseDeviceCounts `json:"deviceCounts" yaml:"deviceCounts"`
}

type Bundle struct {
	Applications        []ApplicationFull2        `json:"applications" yaml:"applications"`
	ApplicationStatuses []DeviceApplicationStatus `json:"applicationStatuses" yaml:"applicationStatuses"`
	ServiceStatuses     []DeviceServiceStatus     `json:"serviceStatuses" yaml:"serviceStatuses"`
	DesiredAgentSpec    string                    `json:"desiredAgentSpec" yaml:"desiredAgentSpec"`
	DesiredAgentVersion string                    `json:"desiredAgentVersion" yaml:"desiredAgentVersion"`
}

type ApplicationFull2 struct {
	Application   Application `json:"application" yaml:"application"`
	LatestRelease Release     `json:"latestRelease" yaml:"latestRelease"`
}

type DeviceInfo struct {
	AgentVersion string    `json:"agentVersion" yaml:"agentVersion"`
	IPAddress    string    `json:"ipAddress" yaml:"ipAddress"`
	OSRelease    OSRelease `json:"osRelease" yaml:"osRelease"`
}

type OSRelease struct {
	PrettyName string `json:"prettyName" yaml:"prettyName"`
	Name       string `json:"name" yaml:"name"`
	VersionID  string `json:"versionId" yaml:"versionId"`
	Version    string `json:"version" yaml:"version"`
	ID         string `json:"id" yaml:"id"`
	IDLike     string `json:"idLike" yaml:"idLike"`
}
