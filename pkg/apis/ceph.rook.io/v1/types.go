/*
Copyright 2020 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"time"

	rookv1 "github.com/rook/rook/pkg/apis/rook.io/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ***************************************************************************
// IMPORTANT FOR CODE GENERATION
// If the types in this file are updated, you will need to run
// `make codegen` to generate the new types under the client/clientset folder.
// ***************************************************************************

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type CephClusterHealthCheckSpec struct {
	DaemonHealth  DaemonHealthSpec                     `json:"daemonHealth,omitempty"`
	LivenessProbe map[rookv1.KeyType]*rookv1.ProbeSpec `json:"livenessProbe,omitempty"`
}

type DaemonHealthSpec struct {
	Status              HealthCheckSpec `json:"status,omitempty"`
	Monitor             HealthCheckSpec `json:"mon,omitempty"`
	ObjectStorageDaemon HealthCheckSpec `json:"osd,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephCluster `json:"items"`
}

type ClusterSpec struct {
	// The version information that instructs Rook to orchestrate a particular version of Ceph.
	CephVersion CephVersionSpec `json:"cephVersion,omitempty"`

	// Ceph Drive Groups specification for how storage should be used in the cluster, given
	// precedent over the Storage spec.
	DriveGroups DriveGroupsSpec `json:"driveGroups,omitempty"`

	// A spec for available storage in the cluster and how it should be used
	Storage rookv1.StorageScopeSpec `json:"storage,omitempty"`

	// The annotations-related configuration to add/set on each Pod related object.
	Annotations rookv1.AnnotationsSpec `json:"annotations,omitempty"`

	// The placement-related configuration to pass to kubernetes (affinity, node selector, tolerations).
	Placement rookv1.PlacementSpec `json:"placement,omitempty"`

	// Network related configuration
	Network NetworkSpec `json:"network,omitempty"`

	// Resources set resource requests and limits
	Resources rookv1.ResourceSpec `json:"resources,omitempty"`

	// PriorityClassNames sets priority classes on components
	PriorityClassNames rookv1.PriorityClassNamesSpec `json:"priorityClassNames,omitempty"`

	// The path on the host where config and data can be persisted.
	DataDirHostPath string `json:"dataDirHostPath,omitempty"`

	// SkipUpgradeChecks defines if an upgrade should be forced even if one of the check fails
	SkipUpgradeChecks bool `json:"skipUpgradeChecks,omitempty"`

	// ContinueUpgradeAfterChecksEvenIfNotHealthy defines if an upgrade should continue even if PGs are not clean
	ContinueUpgradeAfterChecksEvenIfNotHealthy bool `json:"continueUpgradeAfterChecksEvenIfNotHealthy,omitempty"`

	// A spec for configuring disruption management.
	DisruptionManagement DisruptionManagementSpec `json:"disruptionManagement,omitempty"`

	// A spec for mon related options
	Mon MonSpec `json:"mon,omitempty"`

	// A spec for the crash controller
	CrashCollector CrashCollectorSpec `json:"crashCollector"`

	// Dashboard settings
	Dashboard DashboardSpec `json:"dashboard,omitempty"`

	// Prometheus based Monitoring settings
	Monitoring MonitoringSpec `json:"monitoring,omitempty"`

	// Whether the Ceph Cluster is running external to this Kubernetes cluster
	// mon, mgr, osd, mds, and discover daemons will not be created for external clusters.
	External ExternalSpec `json:"external"`

	// A spec for mgr related options
	Mgr MgrSpec `json:"mgr,omitempty"`

	// Remove the OSD that is out and safe to remove only if this option is true
	RemoveOSDsIfOutAndSafeToRemove bool `json:"removeOSDsIfOutAndSafeToRemove"`

	// Indicates user intent when deleting a cluster; blocks orchestration and should not be set if cluster
	// deletion is not imminent.
	CleanupPolicy CleanupPolicySpec `json:"cleanupPolicy,omitempty"`

	// Internal daemon healthchecks and liveness probe
	HealthCheck CephClusterHealthCheckSpec `json:"healthCheck"`
}

// VersionSpec represents the settings for the Ceph version that Rook is orchestrating.
type CephVersionSpec struct {
	// Image is the container image used to launch the ceph daemons, such as ceph/ceph:v15.2.4
	Image string `json:"image,omitempty"`

	// Whether to allow unsupported versions (do not set to true in production)
	AllowUnsupported bool `json:"allowUnsupported,omitempty"`
}

// DriveGroupsSpec is a list Ceph Drive Group specifications.
type DriveGroupsSpec []DriveGroup

// DriveGroup specifies a Ceph Drive Group and where Rook should apply the Drive Group.
type DriveGroup struct {
	// Name is a unique name used to identify the Drive Group.
	Name string `json:"name"`

	// Spec is the JSON or YAML definition of a Ceph Drive Group. Placement information contained
	// within this spec is ignored, as placement is specified via the Rook placement mechanism below.
	Spec DriveGroupSpec `json:"spec"`

	// Placement defines which nodes the Drive Group should be applied to.
	Placement rookv1.Placement `json:"placement,omitempty"`
}

// DriveGroupSpec is a YAML or JSON definition of a Ceph Drive Group.
type DriveGroupSpec map[string]interface{}

// DashboardSpec represents the settings for the Ceph dashboard
type DashboardSpec struct {
	// Whether to enable the dashboard
	Enabled bool `json:"enabled,omitempty"`
	// A prefix for all URLs to use the dashboard with a reverse proxy
	UrlPrefix string `json:"urlPrefix,omitempty"`
	// The dashboard webserver port
	Port int `json:"port,omitempty"`
	// Whether SSL should be used
	SSL bool `json:"ssl,omitempty"`
}

// MonitoringSpec represents the settings for Prometheus based Ceph monitoring
type MonitoringSpec struct {
	// Whether to create the prometheus rules for the ceph cluster. If true, the prometheus
	// types must exist or the creation will fail.
	Enabled bool `json:"enabled,omitempty"`

	// The namespace where the prometheus rules and alerts should be created.
	// If empty, the same namespace as the cluster will be used.
	RulesNamespace string `json:"rulesNamespace,omitempty"`

	// ExternalMgrEndpoints points to an existing Ceph prometheus exporter endpoint
	ExternalMgrEndpoints []v1.EndpointAddress `json:"externalMgrEndpoints,omitempty"`
}

type ClusterStatus struct {
	State       ClusterState    `json:"state,omitempty"`
	Phase       ConditionType   `json:"phase,omitempty"`
	Message     string          `json:"message,omitempty"`
	Conditions  []Condition     `json:"conditions,omitempty"`
	CephStatus  *CephStatus     `json:"ceph,omitempty"`
	CephVersion *ClusterVersion `json:"version,omitempty"`
}

type CephStatus struct {
	Health         string                       `json:"health,omitempty"`
	Details        map[string]CephHealthMessage `json:"details,omitempty"`
	LastChecked    string                       `json:"lastChecked,omitempty"`
	LastChanged    string                       `json:"lastChanged,omitempty"`
	PreviousHealth string                       `json:"previousHealth,omitempty"`
}

type ClusterVersion struct {
	Image   string `json:"image,omitempty"`
	Version string `json:"version,omitempty"`
}

type CephHealthMessage struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type Condition struct {
	Type               ConditionType      `json:"type,omitempty"`
	Status             v1.ConditionStatus `json:"status,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	Message            string             `json:"message,omitempty"`
	LastHeartbeatTime  metav1.Time        `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime metav1.Time        `json:"lastTransitionTime,omitempty"`
}

type ConditionType string

const (
	ConditionIgnored     ConditionType = "Ignored"
	ConditionConnecting  ConditionType = "Connecting"
	ConditionConnected   ConditionType = "Connected"
	ConditionProgressing ConditionType = "Progressing"
	ConditionReady       ConditionType = "Ready"
	ConditionUpdating    ConditionType = "Updating"
	ConditionFailure     ConditionType = "Failure"
	ConditionUpgrading   ConditionType = "Upgrading"
	ConditionDeleting    ConditionType = "Deleting"
	ConditionHealthy     ConditionType = "Healthy"
	// DefaultFailureDomain for PoolSpec
	DefaultFailureDomain = "host"
)

type ClusterState string

const (
	ClusterStateCreating   ClusterState = "Creating"
	ClusterStateCreated    ClusterState = "Created"
	ClusterStateUpdating   ClusterState = "Updating"
	ClusterStateConnecting ClusterState = "Connecting"
	ClusterStateConnected  ClusterState = "Connected"
	ClusterStateError      ClusterState = "Error"
)

type MonSpec struct {
	Count                int                       `json:"count,omitempty"`
	AllowMultiplePerNode bool                      `json:"allowMultiplePerNode,omitempty"`
	VolumeClaimTemplate  *v1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// MgrSpec represents options to configure a ceph mgr
type MgrSpec struct {
	Modules []Module `json:"modules,omitempty"`
}

// Module represents mgr modules that the user wants to enable or disable
type Module struct {
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled"`
}

// ExternalSpec represents the options supported by an external cluster
type ExternalSpec struct {
	Enable bool `json:"enable"`
}

// CrashCollectorSpec represents options to configure the crash controller
type CrashCollectorSpec struct {
	Disable bool `json:"disable"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephBlockPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              PoolSpec `json:"spec"`
	Status            *Status  `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CephBlockPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephBlockPool `json:"items"`
}

// PoolSpec represents the spec of ceph pool
type PoolSpec struct {
	// The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map
	FailureDomain string `json:"failureDomain"`

	// The root of the crush hierarchy utilized by the pool
	CrushRoot string `json:"crushRoot"`

	// The device class the OSD should set to (options are: hdd, ssd, or nvme)
	DeviceClass string `json:"deviceClass"`

	// The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)
	CompressionMode string `json:"compressionMode"`

	// The replication settings
	Replicated ReplicatedSpec `json:"replicated"`

	// The erasure code settings
	ErasureCoded ErasureCodedSpec `json:"erasureCoded"`

	// Parameters is a list of properties to enable on a given pool
	Parameters map[string]string `json:"parameters,omitempty"`
}

type Status struct {
	Phase string `json:"phase,omitempty"`
}

// ReplicatedSpec represents the spec for replication in a pool
type ReplicatedSpec struct {
	// Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)
	Size uint `json:"size"`

	// TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity
	TargetSizeRatio float64 `json:"targetSizeRatio"`

	// RequireSafeReplicaSize if false allows you to set replica 1
	RequireSafeReplicaSize bool `json:"requireSafeReplicaSize"`
}

// ErasureCodeSpec represents the spec for erasure code in a pool
type ErasureCodedSpec struct {
	// Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type)
	CodingChunks uint `json:"codingChunks"`

	// Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type)
	DataChunks uint `json:"dataChunks"`

	// The algorithm for erasure coding
	Algorithm string `json:"algorithm"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephFilesystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              FilesystemSpec `json:"spec"`
	Status            *Status        `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephFilesystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephFilesystem `json:"items"`
}

// FilesystemSpec represents the spec of a file system
type FilesystemSpec struct {
	// The metadata pool settings
	MetadataPool PoolSpec `json:"metadataPool,omitempty"`

	// The data pool settings
	DataPools []PoolSpec `json:"dataPools,omitempty"`

	// Preserve pools on filesystem deletion
	PreservePoolsOnDelete bool `json:"preservePoolsOnDelete"`

	// The mds pod info
	MetadataServer MetadataServerSpec `json:"metadataServer"`
}

type MetadataServerSpec struct {
	// The number of metadata servers that are active. The remaining servers in the cluster will be in standby mode.
	ActiveCount int32 `json:"activeCount"`

	// Whether each active MDS instance will have an active standby with a warm metadata cache for faster failover.
	// If false, standbys will still be available, but will not have a warm metadata cache.
	ActiveStandby bool `json:"activeStandby"`

	// The affinity to place the mds pods (default is to place on all available node) with a daemonset
	Placement rookv1.Placement `json:"placement"`

	// The annotations-related configuration to add/set on each Pod related object.
	Annotations rookv1.Annotations `json:"annotations,omitempty"`

	// The resource requirements for the rgw pods
	Resources v1.ResourceRequirements `json:"resources"`

	// PriorityClassName sets priority classes on components
	PriorityClassName string `json:"priorityClassName,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectStoreSpec    `json:"spec"`
	Status            *ObjectStoreStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephObjectStore `json:"items"`
}

// ObjectStoreSpec represent the spec of a pool
type ObjectStoreSpec struct {
	// The metadata pool settings
	MetadataPool PoolSpec `json:"metadataPool"`

	// The data pool settings
	DataPool PoolSpec `json:"dataPool"`

	// Preserve pools on object store deletion
	PreservePoolsOnDelete bool `json:"preservePoolsOnDelete"`

	// The rgw pod info
	Gateway GatewaySpec `json:"gateway"`

	// The multisite info
	Zone ZoneSpec `json:"zone"`

	// The rgw Bucket healthchecks and liveness probe
	HealthCheck BucketHealthCheckSpec `json:"healthCheck"`
}

type BucketHealthCheckSpec struct {
	Bucket        HealthCheckSpec   `json:"rgw,omitempty"`
	LivenessProbe *rookv1.ProbeSpec `json:"livenessProbe,omitempty"`
}

type HealthCheckSpec struct {
	Disabled bool   `json:"disabled,omitempty"`
	Interval string `json:"interval,omitempty"`
	Timeout  string `json:"timeout,omitempty"`
}

type GatewaySpec struct {
	// The port the rgw service will be listening on (http)
	Port int32 `json:"port"`

	// The port the rgw service will be listening on (https)
	SecurePort int32 `json:"securePort"`

	// The number of pods in the rgw replicaset. If "allNodes" is specified, a daemonset is created.
	Instances int32 `json:"instances"`

	// Whether the rgw pods should be started as a daemonset on all nodes
	AllNodes bool `json:"allNodes"`

	// The name of the secret that stores the ssl certificate for secure rgw connections
	SSLCertificateRef string `json:"sslCertificateRef"`

	// The affinity to place the rgw pods (default is to place on any available node)
	Placement rookv1.Placement `json:"placement"`

	// The annotations-related configuration to add/set on each Pod related object.
	Annotations rookv1.Annotations `json:"annotations,omitempty"`

	// The resource requirements for the rgw pods
	Resources v1.ResourceRequirements `json:"resources"`

	// PriorityClassName sets priority classes on the rgw pods
	PriorityClassName string `json:"priorityClassName,omitempty"`

	// ExternalRgwEndpoints points to external rgw endpoint(s)
	ExternalRgwEndpoints []v1.EndpointAddress `json:"externalRgwEndpoints,omitempty"`
}

type ZoneSpec struct {
	// RGW Zone the Object Store is in
	Name string `json:"name"`
}

type ObjectStoreStatus struct {
	Phase        ConditionType     `json:"phase,omitempty"`
	Message      string            `json:"message,omitempty"`
	BucketStatus *BucketStatus     `json:"bucketStatus,omitempty"`
	Info         map[string]string `json:"info,omitempty"`
}

type BucketStatus struct {
	Health      ConditionType `json:"health,omitempty"`
	Details     string        `json:"details,omitempty"`
	LastChecked string        `json:"lastChecked,omitempty"`
	LastChanged string        `json:"lastChanged,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectStoreUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectStoreUserSpec    `json:"spec"`
	Status            *ObjectStoreUserStatus `json:"status"`
}

type ObjectStoreUserStatus struct {
	Phase string            `json:"phase,omitempty"`
	Info  map[string]string `json:"info"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectStoreUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephObjectStoreUser `json:"items"`
}

// ObjectStoreUserSpec represent the spec of an Objectstoreuser
type ObjectStoreUserSpec struct {
	//The store the user will be created in
	Store string `json:"store,omitempty"`
	//The display name for the ceph users
	DisplayName string `json:"displayName,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectRealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectRealmSpec `json:"spec"`
	Status            *Status         `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectRealmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephObjectRealm `json:"items"`
}

// ObjectRealmSpec represent the spec of an ObjectRealm
type ObjectRealmSpec struct {
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectZoneGroupSpec `json:"spec"`
	Status            *Status             `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephObjectZoneGroup `json:"items"`
}

// ObjectZoneGroupSpec represent the spec of an ObjecZoneGroup
type ObjectZoneGroupSpec struct {
	//The display name for the ceph users
	Realm string `json:"realm"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectZoneSpec `json:"spec"`
	Status            *Status        `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephObjectZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephObjectZone `json:"items"`
}

// ObjectZoneSpec represent the spec of an ObjectZone
type ObjectZoneSpec struct {
	//The display name for the ceph users
	ZoneGroup string `json:"zoneGroup"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephNFS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NFSGaneshaSpec `json:"spec"`
	Status            *Status        `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephNFSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephNFS `json:"items"`
}

// NFSGaneshaSpec represents the spec of an nfs ganesha server
type NFSGaneshaSpec struct {
	RADOS GaneshaRADOSSpec `json:"rados"`

	Server GaneshaServerSpec `json:"server"`
}

type GaneshaRADOSSpec struct {
	// Pool is the RADOS pool where NFS client recovery data is stored.
	Pool string `json:"pool"`

	// Namespace is the RADOS namespace where NFS client recovery data is stored.
	Namespace string `json:"namespace"`
}

type GaneshaServerSpec struct {
	// The number of active Ganesha servers
	Active int `json:"active"`

	// The affinity to place the ganesha pods
	Placement rookv1.Placement `json:"placement"`

	// The annotations-related configuration to add/set on each Pod related object.
	Annotations rookv1.Annotations `json:"annotations,omitempty"`

	// Resources set resource requests and limits
	Resources v1.ResourceRequirements `json:"resources,omitempty"`

	// PriorityClassName sets the priority class on the pods
	PriorityClassName string `json:"priorityClassName,omitempty"`
}

// NetworkSpec for Ceph includes backward compatibility code
type NetworkSpec struct {
	rookv1.NetworkSpec `json:",inline"`

	// HostNetwork to enable host network
	HostNetwork bool `json:"hostNetwork"`
}

// DisruptionManagementSpec configures management of daemon disruptions
type DisruptionManagementSpec struct {

	// This enables management of poddisruptionbudgets
	ManagePodBudgets bool `json:"managePodBudgets,omitempty"`

	// OSDMaintenanceTimeout sets how many additional minutes the DOWN/OUT interval is for drained failure domains
	// it only works if managePodBudgetss is true.
	// the default is 30 minutes
	OSDMaintenanceTimeout time.Duration `json:"osdMaintenanceTimeout,omitempty"`

	// This enables management of machinedisruptionbudgets
	ManageMachineDisruptionBudgets bool `json:"manageMachineDisruptionBudgets,omitempty"`

	// Namespace to look for MDBs by the machineDisruptionBudgetController
	MachineDisruptionBudgetNamespace string `json:"machineDisruptionBudgetNamespace,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClientSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephClient `json:"items"`
}

type ClientSpec struct {
	Name string            `json:"name"`
	Caps map[string]string `json:"caps"`
}

type CleanupPolicySpec struct {
	Confirmation CleanupConfirmationProperty `json:"confirmation,omitempty"`
}

type CleanupConfirmationProperty string

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephRBDMirror struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              RBDMirroringSpec `json:"spec"`
	Status            *Status          `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CephRBDMirrorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CephRBDMirror `json:"items"`
}

type RBDMirroringSpec struct {
	// Count represents the number of rbd mirror instance to run
	Count int `json:"count"`

	// The affinity to place the rgw pods (default is to place on any available node)
	Placement rookv1.Placement `json:"placement"`

	// The annotations-related configuration to add/set on each Pod related object.
	Annotations rookv1.Annotations `json:"annotations,omitempty"`

	// The resource requirements for the rgw pods
	Resources v1.ResourceRequirements `json:"resources"`

	// PriorityClassName sets priority classes on the rgw pods
	PriorityClassName string `json:"priorityClassName,omitempty"`
}
