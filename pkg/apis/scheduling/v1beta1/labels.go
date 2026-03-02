/*
Copyright 2019 The Volcano Authors.

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

package v1beta1

// AnnotationPrefix is the key prefix to pick annotations from upper resource to podgroup
const AnnotationPrefix = "volcano.sh/"

// JobWaitingTime is the key of sla plugin to set maximum waiting time
// that a job could stay Pending in service level agreement
const JobWaitingTime = "volcano.sh/sla-waiting-time"

const KubeHierarchyAnnotationKey = "volcano.sh/hierarchy"

const KubeHierarchyWeightAnnotationKey = "volcano.sh/hierarchy-weights"

// KubeGroupNameAnnotationKey is the annotation key of Pod to identify
// which PodGroup it belongs to.
const KubeGroupNameAnnotationKey = "scheduling.k8s.io/group-name"

// VolcanoGroupNameAnnotationKey is the annotation key of Pod to identify
// which PodGroup it belongs to.
const VolcanoGroupNameAnnotationKey = GroupName + "/group-name"

// VolcanoGroupMinResourcesAnnotationKey is the annotation key of PodGroup's PodGroup.Spec.MinResources
// which PodGroup it belongs to.
const VolcanoGroupMinResourcesAnnotationKey = GroupName + "/group-min-resources"

// QueueNameAnnotationKey is the annotation key of Pod to identify
// which queue it belongs to.
const QueueNameAnnotationKey = GroupName + "/queue-name"

// PodPreemptable is the key of preemptable
const PodPreemptable = "volcano.sh/preemptable"

// CooldownTime is the key of cooldown-time, value's format "600s","10m"
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
const CooldownTime = "volcano.sh/cooldown-time"

// RevocableZone is the key of revocable-zone
const RevocableZone = "volcano.sh/revocable-zone"

// JDBMinAvailable is the key of min available pod number
const JDBMinAvailable = "volcano.sh/jdb-min-available"

// JDBMaxUnavailable is the key of max unavailable pod number
const JDBMaxUnavailable = "volcano.sh/jdb-max-unavailable"

// NumaPolicyKey is the key of pod numa-topology policy
const NumaPolicyKey = "volcano.sh/numa-topology-policy"

// TopologyDecisionAnnotation is the key of topology decision about pod request resource
const TopologyDecisionAnnotation = "volcano.sh/topology-decision"

// PodQosLevel is the key of pod qos level
const PodQosLevel = "volcano.sh/qos-level"

// EnableInQuotasAnnotationKey is the key of enable in quotas plugin, true or false
const EnableInQuotasAnnotationKey = "volcano.sh/enable-in-quotas-plugin"

// SharedLimitsGroupKey is a key to determine queue relation to
// shared limits group
const SharedLimitsGroupKey = "volcano.sh/shared-limits-group-key"

// IsSharedLimitsGroupLeader is true if queue's limits should be interpreted as limits
// for all jobs in shared limits group
const IsSharedLimitsGroupLeader = "volcano.sh/is-shared-limits-group-leader"

// MinPodResourcesInQueue is an annotation to determine queue restrictions on pod resources size
const MinPodResourcesInQueue = "volcano.sh/min-pod-resources-in-queue"

// VolcanoQueueAlias is an annotation to specify the queue alias
const VolcanoQueueAlias = "volcano.sh/queue-alias"

// VolcanoQueueGroupAlias is an annotation to specify the queue group alias
const VolcanoQueueGroupAlias = "volcano.sh/queue-group-alias"
