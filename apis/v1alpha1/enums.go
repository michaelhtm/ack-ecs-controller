// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AgentUpdateStatus string

const (
	AgentUpdateStatus_FAILED   AgentUpdateStatus = "FAILED"
	AgentUpdateStatus_PENDING  AgentUpdateStatus = "PENDING"
	AgentUpdateStatus_STAGED   AgentUpdateStatus = "STAGED"
	AgentUpdateStatus_STAGING  AgentUpdateStatus = "STAGING"
	AgentUpdateStatus_UPDATED  AgentUpdateStatus = "UPDATED"
	AgentUpdateStatus_UPDATING AgentUpdateStatus = "UPDATING"
)

type ApplicationProtocol string

const (
	ApplicationProtocol_grpc  ApplicationProtocol = "grpc"
	ApplicationProtocol_http  ApplicationProtocol = "http"
	ApplicationProtocol_http2 ApplicationProtocol = "http2"
)

type AssignPublicIP string

const (
	AssignPublicIP_DISABLED AssignPublicIP = "DISABLED"
	AssignPublicIP_ENABLED  AssignPublicIP = "ENABLED"
)

type AvailabilityZoneRebalancing string

const (
	AvailabilityZoneRebalancing_DISABLED AvailabilityZoneRebalancing = "DISABLED"
	AvailabilityZoneRebalancing_ENABLED  AvailabilityZoneRebalancing = "ENABLED"
)

type CPUArchitecture string

const (
	CPUArchitecture_ARM64  CPUArchitecture = "ARM64"
	CPUArchitecture_X86_64 CPUArchitecture = "X86_64"
)

type CapacityProviderField string

const (
	CapacityProviderField_TAGS CapacityProviderField = "TAGS"
)

type CapacityProviderStatus string

const (
	CapacityProviderStatus_ACTIVE   CapacityProviderStatus = "ACTIVE"
	CapacityProviderStatus_INACTIVE CapacityProviderStatus = "INACTIVE"
)

type CapacityProviderUpdateStatus string

const (
	CapacityProviderUpdateStatus_DELETE_COMPLETE    CapacityProviderUpdateStatus = "DELETE_COMPLETE"
	CapacityProviderUpdateStatus_DELETE_FAILED      CapacityProviderUpdateStatus = "DELETE_FAILED"
	CapacityProviderUpdateStatus_DELETE_IN_PROGRESS CapacityProviderUpdateStatus = "DELETE_IN_PROGRESS"
	CapacityProviderUpdateStatus_UPDATE_COMPLETE    CapacityProviderUpdateStatus = "UPDATE_COMPLETE"
	CapacityProviderUpdateStatus_UPDATE_FAILED      CapacityProviderUpdateStatus = "UPDATE_FAILED"
	CapacityProviderUpdateStatus_UPDATE_IN_PROGRESS CapacityProviderUpdateStatus = "UPDATE_IN_PROGRESS"
)

type ClusterField string

const (
	ClusterField_ATTACHMENTS    ClusterField = "ATTACHMENTS"
	ClusterField_CONFIGURATIONS ClusterField = "CONFIGURATIONS"
	ClusterField_SETTINGS       ClusterField = "SETTINGS"
	ClusterField_STATISTICS     ClusterField = "STATISTICS"
	ClusterField_TAGS           ClusterField = "TAGS"
)

type ClusterSettingName string

const (
	ClusterSettingName_containerInsights ClusterSettingName = "containerInsights"
)

type Compatibility string

const (
	Compatibility_EC2      Compatibility = "EC2"
	Compatibility_EXTERNAL Compatibility = "EXTERNAL"
	Compatibility_FARGATE  Compatibility = "FARGATE"
)

type Connectivity string

const (
	Connectivity_CONNECTED    Connectivity = "CONNECTED"
	Connectivity_DISCONNECTED Connectivity = "DISCONNECTED"
)

type ContainerCondition string

const (
	ContainerCondition_COMPLETE ContainerCondition = "COMPLETE"
	ContainerCondition_HEALTHY  ContainerCondition = "HEALTHY"
	ContainerCondition_START    ContainerCondition = "START"
	ContainerCondition_SUCCESS  ContainerCondition = "SUCCESS"
)

type ContainerInstanceField string

const (
	ContainerInstanceField_CONTAINER_INSTANCE_HEALTH ContainerInstanceField = "CONTAINER_INSTANCE_HEALTH"
	ContainerInstanceField_TAGS                      ContainerInstanceField = "TAGS"
)

type ContainerInstanceStatus string

const (
	ContainerInstanceStatus_ACTIVE              ContainerInstanceStatus = "ACTIVE"
	ContainerInstanceStatus_DEREGISTERING       ContainerInstanceStatus = "DEREGISTERING"
	ContainerInstanceStatus_DRAINING            ContainerInstanceStatus = "DRAINING"
	ContainerInstanceStatus_REGISTERING         ContainerInstanceStatus = "REGISTERING"
	ContainerInstanceStatus_REGISTRATION_FAILED ContainerInstanceStatus = "REGISTRATION_FAILED"
)

type DeploymentControllerType string

const (
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerType_ECS         DeploymentControllerType = "ECS"
	DeploymentControllerType_EXTERNAL    DeploymentControllerType = "EXTERNAL"
)

type DeploymentRolloutState string

const (
	DeploymentRolloutState_COMPLETED   DeploymentRolloutState = "COMPLETED"
	DeploymentRolloutState_FAILED      DeploymentRolloutState = "FAILED"
	DeploymentRolloutState_IN_PROGRESS DeploymentRolloutState = "IN_PROGRESS"
)

type DesiredStatus string

const (
	DesiredStatus_PENDING DesiredStatus = "PENDING"
	DesiredStatus_RUNNING DesiredStatus = "RUNNING"
	DesiredStatus_STOPPED DesiredStatus = "STOPPED"
)

type DeviceCgroupPermission string

const (
	DeviceCgroupPermission_mknod DeviceCgroupPermission = "mknod"
	DeviceCgroupPermission_read  DeviceCgroupPermission = "read"
	DeviceCgroupPermission_write DeviceCgroupPermission = "write"
)

type EBSResourceType string

const (
	EBSResourceType_volume EBSResourceType = "volume"
)

type EFSAuthorizationConfigIAM string

const (
	EFSAuthorizationConfigIAM_DISABLED EFSAuthorizationConfigIAM = "DISABLED"
	EFSAuthorizationConfigIAM_ENABLED  EFSAuthorizationConfigIAM = "ENABLED"
)

type EFSTransitEncryption string

const (
	EFSTransitEncryption_DISABLED EFSTransitEncryption = "DISABLED"
	EFSTransitEncryption_ENABLED  EFSTransitEncryption = "ENABLED"
)

type EnvironmentFileType string

const (
	EnvironmentFileType_s3 EnvironmentFileType = "s3"
)

type ExecuteCommandLogging string

const (
	ExecuteCommandLogging_DEFAULT  ExecuteCommandLogging = "DEFAULT"
	ExecuteCommandLogging_NONE     ExecuteCommandLogging = "NONE"
	ExecuteCommandLogging_OVERRIDE ExecuteCommandLogging = "OVERRIDE"
)

type FirelensConfigurationType string

const (
	FirelensConfigurationType_fluentbit FirelensConfigurationType = "fluentbit"
	FirelensConfigurationType_fluentd   FirelensConfigurationType = "fluentd"
)

type HealthStatus string

const (
	HealthStatus_HEALTHY   HealthStatus = "HEALTHY"
	HealthStatus_UNHEALTHY HealthStatus = "UNHEALTHY"
	HealthStatus_UNKNOWN   HealthStatus = "UNKNOWN"
)

type IPCMode string

const (
	IPCMode_host IPCMode = "host"
	IPCMode_none IPCMode = "none"
	IPCMode_task IPCMode = "task"
)

type InstanceHealthCheckState string

const (
	InstanceHealthCheckState_IMPAIRED          InstanceHealthCheckState = "IMPAIRED"
	InstanceHealthCheckState_INITIALIZING      InstanceHealthCheckState = "INITIALIZING"
	InstanceHealthCheckState_INSUFFICIENT_DATA InstanceHealthCheckState = "INSUFFICIENT_DATA"
	InstanceHealthCheckState_OK                InstanceHealthCheckState = "OK"
)

type InstanceHealthCheckType string

const (
	InstanceHealthCheckType_CONTAINER_RUNTIME InstanceHealthCheckType = "CONTAINER_RUNTIME"
)

type LaunchType string

const (
	LaunchType_EC2      LaunchType = "EC2"
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
	LaunchType_FARGATE  LaunchType = "FARGATE"
)

type LogDriver string

const (
	LogDriver_awsfirelens LogDriver = "awsfirelens"
	LogDriver_awslogs     LogDriver = "awslogs"
	LogDriver_fluentd     LogDriver = "fluentd"
	LogDriver_gelf        LogDriver = "gelf"
	LogDriver_journald    LogDriver = "journald"
	LogDriver_json_file   LogDriver = "json-file"
	LogDriver_splunk      LogDriver = "splunk"
	LogDriver_syslog      LogDriver = "syslog"
)

type ManagedAgentName string

const (
	ManagedAgentName_ExecuteCommandAgent ManagedAgentName = "ExecuteCommandAgent"
)

type ManagedDraining string

const (
	ManagedDraining_DISABLED ManagedDraining = "DISABLED"
	ManagedDraining_ENABLED  ManagedDraining = "ENABLED"
)

type ManagedScalingStatus string

const (
	ManagedScalingStatus_DISABLED ManagedScalingStatus = "DISABLED"
	ManagedScalingStatus_ENABLED  ManagedScalingStatus = "ENABLED"
)

type ManagedTerminationProtection string

const (
	ManagedTerminationProtection_DISABLED ManagedTerminationProtection = "DISABLED"
	ManagedTerminationProtection_ENABLED  ManagedTerminationProtection = "ENABLED"
)

type NetworkMode string

const (
	NetworkMode_awsvpc NetworkMode = "awsvpc"
	NetworkMode_bridge NetworkMode = "bridge"
	NetworkMode_host   NetworkMode = "host"
	NetworkMode_none   NetworkMode = "none"
)

type OSFamily string

const (
	OSFamily_LINUX                    OSFamily = "LINUX"
	OSFamily_WINDOWS_SERVER_2004_CORE OSFamily = "WINDOWS_SERVER_2004_CORE"
	OSFamily_WINDOWS_SERVER_2016_FULL OSFamily = "WINDOWS_SERVER_2016_FULL"
	OSFamily_WINDOWS_SERVER_2019_CORE OSFamily = "WINDOWS_SERVER_2019_CORE"
	OSFamily_WINDOWS_SERVER_2019_FULL OSFamily = "WINDOWS_SERVER_2019_FULL"
	OSFamily_WINDOWS_SERVER_2022_CORE OSFamily = "WINDOWS_SERVER_2022_CORE"
	OSFamily_WINDOWS_SERVER_2022_FULL OSFamily = "WINDOWS_SERVER_2022_FULL"
	OSFamily_WINDOWS_SERVER_20H2_CORE OSFamily = "WINDOWS_SERVER_20H2_CORE"
)

type PIDMode string

const (
	PIDMode_host PIDMode = "host"
	PIDMode_task PIDMode = "task"
)

type PlacementConstraintType string

const (
	PlacementConstraintType_distinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintType_memberOf         PlacementConstraintType = "memberOf"
)

type PlacementStrategyType string

const (
	PlacementStrategyType_binpack PlacementStrategyType = "binpack"
	PlacementStrategyType_random  PlacementStrategyType = "random"
	PlacementStrategyType_spread  PlacementStrategyType = "spread"
)

type PlatformDeviceType string

const (
	PlatformDeviceType_GPU PlatformDeviceType = "GPU"
)

type PropagateTags string

const (
	PropagateTags_NONE            PropagateTags = "NONE"
	PropagateTags_SERVICE         PropagateTags = "SERVICE"
	PropagateTags_TASK_DEFINITION PropagateTags = "TASK_DEFINITION"
)

type ProxyConfigurationType string

const (
	ProxyConfigurationType_APPMESH ProxyConfigurationType = "APPMESH"
)

type ResourceType string

const (
	ResourceType_GPU                  ResourceType = "GPU"
	ResourceType_InferenceAccelerator ResourceType = "InferenceAccelerator"
)

type ScaleUnit string

const (
	ScaleUnit_PERCENT ScaleUnit = "PERCENT"
)

type SchedulingStrategy string

const (
	SchedulingStrategy_DAEMON  SchedulingStrategy = "DAEMON"
	SchedulingStrategy_REPLICA SchedulingStrategy = "REPLICA"
)

type Scope string

const (
	Scope_shared Scope = "shared"
	Scope_task   Scope = "task"
)

type ServiceDeploymentRollbackMonitorsStatus string

const (
	ServiceDeploymentRollbackMonitorsStatus_DISABLED            ServiceDeploymentRollbackMonitorsStatus = "DISABLED"
	ServiceDeploymentRollbackMonitorsStatus_MONITORING          ServiceDeploymentRollbackMonitorsStatus = "MONITORING"
	ServiceDeploymentRollbackMonitorsStatus_MONITORING_COMPLETE ServiceDeploymentRollbackMonitorsStatus = "MONITORING_COMPLETE"
	ServiceDeploymentRollbackMonitorsStatus_TRIGGERED           ServiceDeploymentRollbackMonitorsStatus = "TRIGGERED"
)

type ServiceDeploymentStatus string

const (
	ServiceDeploymentStatus_IN_PROGRESS          ServiceDeploymentStatus = "IN_PROGRESS"
	ServiceDeploymentStatus_PENDING              ServiceDeploymentStatus = "PENDING"
	ServiceDeploymentStatus_ROLLBACK_FAILED      ServiceDeploymentStatus = "ROLLBACK_FAILED"
	ServiceDeploymentStatus_ROLLBACK_IN_PROGRESS ServiceDeploymentStatus = "ROLLBACK_IN_PROGRESS"
	ServiceDeploymentStatus_ROLLBACK_SUCCESSFUL  ServiceDeploymentStatus = "ROLLBACK_SUCCESSFUL"
	ServiceDeploymentStatus_STOPPED              ServiceDeploymentStatus = "STOPPED"
	ServiceDeploymentStatus_STOP_REQUESTED       ServiceDeploymentStatus = "STOP_REQUESTED"
	ServiceDeploymentStatus_SUCCESSFUL           ServiceDeploymentStatus = "SUCCESSFUL"
)

type ServiceField string

const (
	ServiceField_TAGS ServiceField = "TAGS"
)

type SettingName string

const (
	SettingName_awsvpcTrunking                  SettingName = "awsvpcTrunking"
	SettingName_containerInsights               SettingName = "containerInsights"
	SettingName_containerInstanceLongArnFormat  SettingName = "containerInstanceLongArnFormat"
	SettingName_fargateFIPSMode                 SettingName = "fargateFIPSMode"
	SettingName_fargateTaskRetirementWaitPeriod SettingName = "fargateTaskRetirementWaitPeriod"
	SettingName_guardDutyActivate               SettingName = "guardDutyActivate"
	SettingName_serviceLongArnFormat            SettingName = "serviceLongArnFormat"
	SettingName_tagResourceAuthorization        SettingName = "tagResourceAuthorization"
	SettingName_taskLongArnFormat               SettingName = "taskLongArnFormat"
)

type SettingType string

const (
	SettingType_aws_managed SettingType = "aws_managed"
	SettingType_user        SettingType = "user"
)

type SortOrder string

const (
	SortOrder_ASC  SortOrder = "ASC"
	SortOrder_DESC SortOrder = "DESC"
)

type StabilityStatus string

const (
	StabilityStatus_STABILIZING  StabilityStatus = "STABILIZING"
	StabilityStatus_STEADY_STATE StabilityStatus = "STEADY_STATE"
)

type TargetType string

const (
	TargetType_container_instance TargetType = "container-instance"
)

type TaskDefinitionFamilyStatus string

const (
	TaskDefinitionFamilyStatus_ACTIVE   TaskDefinitionFamilyStatus = "ACTIVE"
	TaskDefinitionFamilyStatus_ALL      TaskDefinitionFamilyStatus = "ALL"
	TaskDefinitionFamilyStatus_INACTIVE TaskDefinitionFamilyStatus = "INACTIVE"
)

type TaskDefinitionField string

const (
	TaskDefinitionField_TAGS TaskDefinitionField = "TAGS"
)

type TaskDefinitionPlacementConstraintType string

const (
	TaskDefinitionPlacementConstraintType_memberOf TaskDefinitionPlacementConstraintType = "memberOf"
)

type TaskDefinitionStatus_SDK string

const (
	TaskDefinitionStatus_SDK_ACTIVE             TaskDefinitionStatus_SDK = "ACTIVE"
	TaskDefinitionStatus_SDK_DELETE_IN_PROGRESS TaskDefinitionStatus_SDK = "DELETE_IN_PROGRESS"
	TaskDefinitionStatus_SDK_INACTIVE           TaskDefinitionStatus_SDK = "INACTIVE"
)

type TaskField string

const (
	TaskField_TAGS TaskField = "TAGS"
)

type TaskFilesystemType string

const (
	TaskFilesystemType_ext3 TaskFilesystemType = "ext3"
	TaskFilesystemType_ext4 TaskFilesystemType = "ext4"
	TaskFilesystemType_ntfs TaskFilesystemType = "ntfs"
	TaskFilesystemType_xfs  TaskFilesystemType = "xfs"
)

type TaskSetField string

const (
	TaskSetField_TAGS TaskSetField = "TAGS"
)

type TaskStopCode string

const (
	TaskStopCode_EssentialContainerExited  TaskStopCode = "EssentialContainerExited"
	TaskStopCode_ServiceSchedulerInitiated TaskStopCode = "ServiceSchedulerInitiated"
	TaskStopCode_SpotInterruption          TaskStopCode = "SpotInterruption"
	TaskStopCode_TaskFailedToStart         TaskStopCode = "TaskFailedToStart"
	TaskStopCode_TerminationNotice         TaskStopCode = "TerminationNotice"
	TaskStopCode_UserInitiated             TaskStopCode = "UserInitiated"
)

type TransportProtocol string

const (
	TransportProtocol_tcp TransportProtocol = "tcp"
	TransportProtocol_udp TransportProtocol = "udp"
)

type UlimitName string

const (
	UlimitName_core       UlimitName = "core"
	UlimitName_cpu        UlimitName = "cpu"
	UlimitName_data       UlimitName = "data"
	UlimitName_fsize      UlimitName = "fsize"
	UlimitName_locks      UlimitName = "locks"
	UlimitName_memlock    UlimitName = "memlock"
	UlimitName_msgqueue   UlimitName = "msgqueue"
	UlimitName_nice       UlimitName = "nice"
	UlimitName_nofile     UlimitName = "nofile"
	UlimitName_nproc      UlimitName = "nproc"
	UlimitName_rss        UlimitName = "rss"
	UlimitName_rtprio     UlimitName = "rtprio"
	UlimitName_rttime     UlimitName = "rttime"
	UlimitName_sigpending UlimitName = "sigpending"
	UlimitName_stack      UlimitName = "stack"
)

type VersionConsistency string

const (
	VersionConsistency_disabled VersionConsistency = "disabled"
	VersionConsistency_enabled  VersionConsistency = "enabled"
)
