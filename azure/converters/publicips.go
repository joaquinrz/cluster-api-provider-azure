/*
Copyright 2022 The Kubernetes Authors.

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

package converters

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
)

// IPTagsToSDK converts a CAPZ IP tag to an Azure SDK IP tag.
func IPTagsToSDK(ipTags []infrav1.IPTag) *[]network.IPTag {
	if len(ipTags) == 0 {
		return nil
	}
	skdIPTags := make([]network.IPTag, len(ipTags))
	for i, ipTag := range ipTags {
		skdIPTags[i] = network.IPTag{
			IPTagType: to.StringPtr(ipTag.Type),
			Tag:       to.StringPtr(ipTag.Tag),
		}
	}
	return &skdIPTags
}
