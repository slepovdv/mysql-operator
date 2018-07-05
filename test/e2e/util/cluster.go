/*
Copyright 2018 Pressinfra SRL

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

package util

import (
	"fmt"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "github.com/presslabs/mysql-operator/pkg/apis/mysql/v1alpha1"
)

func NewCluster(name, ns string) *api.MysqlCluster {
	return &api.MysqlCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: api.ClusterSpec{
			Replicas:   1,
			SecretName: name,
		},
	}

}

func ClusterCondition(cluster *api.MysqlCluster, cType api.ClusterConditionType) *api.ClusterCondition {
	for _, c := range cluster.Status.Conditions {
		if c.Type == cType {
			return &c
		}
	}
	return nil
}

func NewClusterSecret(name, ns, pw string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		StringData: map[string]string{
			"ROOT_PASSWORD": pw,
		},
	}
}

func OrcClusterName(cluster *api.MysqlCluster) string {
	return fmt.Sprintf("%s.%s", cluster.Name, cluster.Namespace)
}