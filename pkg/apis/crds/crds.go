// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crds

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"
)

var (
	OnecloudClusterCRD = &apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: api.OnecloudClusterCRDName,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   api.SchemeGroupVersion.Group,
			Version: api.SchemeGroupVersion.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural:     api.OnecloudClusterResourcePlural,
				Kind:       api.OnecloudClusterResourceKind,
				ShortNames: []string{"onecloud", "oc"},
			},
			AdditionalPrinterColumns: []apiextensionsv1beta1.CustomResourceColumnDefinition{
				{
					Name:        api.KeystoneComponentType.String(),
					Type:        "string",
					Description: "The image for keystone service",
					JSONPath:    ".spec.keystone.image",
				},
			},
			Validation: &apiextensionsv1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
						"spec": {
							Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
								"mysql": {
									Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
										"host":     {Type: "string", Nullable: false},
										"password": {Type: "string", Nullable: false},
									},
								},
							},
						},
					},
				},
			},
		},
	}
)
