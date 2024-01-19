// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	internalinterfaces "github.com/phamngocsonls/spark-on-k8s-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ScheduledSparkApplications returns a ScheduledSparkApplicationInformer.
	ScheduledSparkApplications() ScheduledSparkApplicationInformer
	// SparkApplications returns a SparkApplicationInformer.
	SparkApplications() SparkApplicationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ScheduledSparkApplications returns a ScheduledSparkApplicationInformer.
func (v *version) ScheduledSparkApplications() ScheduledSparkApplicationInformer {
	return &scheduledSparkApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SparkApplications returns a SparkApplicationInformer.
func (v *version) SparkApplications() SparkApplicationInformer {
	return &sparkApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
