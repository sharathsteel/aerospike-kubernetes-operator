// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeCluster":       schema_pkg_apis_aerospike_v1alpha1_AerospikeCluster(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterSpec":   schema_pkg_apis_aerospike_v1alpha1_AerospikeClusterSpec(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterStatus": schema_pkg_apis_aerospike_v1alpha1_AerospikeClusterStatus(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeNodeSummary":   schema_pkg_apis_aerospike_v1alpha1_AerospikeNodeSummary(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.BlockStorageSpec":       schema_pkg_apis_aerospike_v1alpha1_BlockStorageSpec(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.FileStorageSpec":        schema_pkg_apis_aerospike_v1alpha1_FileStorageSpec(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeDevice":           schema_pkg_apis_aerospike_v1alpha1_VolumeDevice(ref),
		"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeMount":            schema_pkg_apis_aerospike_v1alpha1_VolumeMount(ref),
	}
}

func schema_pkg_apis_aerospike_v1alpha1_AerospikeCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AerospikeCluster is the Schema for the aerospikeclusters API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterSpec", "github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_AerospikeClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AerospikeClusterSpec defines the desired state of AerospikeCluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Aerospike cluster size",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"build": {
						SchemaProps: spec.SchemaProps{
							Description: "Aerospike cluster build",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"multiPodPerHost": {
						SchemaProps: spec.SchemaProps{
							Description: "If set true then multiple pods can be created per Kubernetes Node. This will create a NodePort service for each Pod. NodePort, as the name implies, opens a specific port on all the Kubernetes Nodes , and any traffic that is sent to this port is forwarded to the service. Here service picks a random port in range (30000-32767), so these port should be open.\n\nIf set false then only single pod can be created per Kubernetes Node. This will create Pods using hostPort setting. The container port will be exposed to the external network at <hostIP>:<hostPort>, where the hostIP is the IP address of the Kubernetes Node where the container is running and the hostPort is the port requested by the user.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"blockStorage": {
						SchemaProps: spec.SchemaProps{
							Description: "BlockStorage has block storage info.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.BlockStorageSpec"),
									},
								},
							},
						},
					},
					"fileStorage": {
						SchemaProps: spec.SchemaProps{
							Description: "FileStorage has filesystem storage info.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.FileStorageSpec"),
									},
								},
							},
						},
					},
					"aerospikeConfigSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "AerospikeConfigSecret has secret info created by user. User needs to create this secret having tls files, feature key for cluster",
							Ref:         ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeConfigSecretSpec"),
						},
					},
					"aerospikeAccessControl": {
						SchemaProps: spec.SchemaProps{
							Description: "AerospikeAccessControl has the Aerospike roles and users definitions. Required if aerospike cluster security is enabled.",
							Ref:         ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeAccessControlSpec"),
						},
					},
					"aerospikeConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "AerospikeConfig sets config in aerospike.conf file. Other configs are taken as default",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Define resources requests and limits for Aerospike Server Container. Please contact aerospike for proper sizing exercise Only Memory and Cpu resources can be given Resources.Limits should be more than Resources.Requests.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
				Required: []string{"size", "build", "aerospikeConfig", "resources"},
			},
		},
		Dependencies: []string{
			"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeAccessControlSpec", "github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeConfigSecretSpec", "github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.BlockStorageSpec", "github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.FileStorageSpec", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_AerospikeClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AerospikeClusterStatus defines the observed state of AerospikeCluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"AerospikeClusterSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "The current state of Aerospike cluster.",
							Ref:         ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterSpec"),
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "Nodes tells the observed state of AerospikeClusterNodes",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeNodeSummary"),
									},
								},
							},
						},
					},
				},
				Required: []string{"AerospikeClusterSpec", "nodes"},
			},
		},
		Dependencies: []string{
			"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeClusterSpec", "github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.AerospikeNodeSummary"},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_AerospikeNodeSummary(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AerospikeNodeSummary defines the observed state of AerospikeClusterNode",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"clusterName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"nodeID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ip": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"tlsname": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"build": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"clusterName", "nodeID", "ip", "port", "tlsname", "build"},
			},
		},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_BlockStorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlockStorageSpec has storage info. StorageClass for devices",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClass should be created by user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"volumeDevices": {
						SchemaProps: spec.SchemaProps{
							Description: "VolumeDevices is the list of block devices to be used by the container. Devices should be provisioned by user",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeDevice"),
									},
								},
							},
						},
					},
				},
				Required: []string{"storageClass", "volumeDevices"},
			},
		},
		Dependencies: []string{
			"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeDevice"},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_FileStorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FileStorageSpec has storage info. StorageClass for fileSystems",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClass should be created by user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"volumeMounts": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod volumes to mount into the container's filesystem.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeMount"),
									},
								},
							},
						},
					},
				},
				Required: []string{"storageClass", "volumeMounts"},
			},
		},
		Dependencies: []string{
			"github.com/aerospike/aerospike-kubernetes-operator/pkg/apis/aerospike/v1alpha1.VolumeMount"},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_VolumeDevice(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VolumeDevice is device to be used by the container",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"devicePath": {
						SchemaProps: spec.SchemaProps{
							Description: "DevicePath is the path inside of the container that the device will be mapped to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sizeInGB": {
						SchemaProps: spec.SchemaProps{
							Description: "SizeInGB Size of device in GB",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"devicePath", "sizeInGB"},
			},
		},
	}
}

func schema_pkg_apis_aerospike_v1alpha1_VolumeMount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VolumeMount is Pod volume to mount into the container's filesystem",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Description: "Path within the container at which the volume should be mounted",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sizeInGB": {
						SchemaProps: spec.SchemaProps{
							Description: "SizeInGB Size of mount volume",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"mountPath", "sizeInGB"},
			},
		},
	}
}