// Licensed to Alexandre VILAIN under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Alexandre VILAIN licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/alexandrevilain/temporal-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

func TestWorkerProcess(t *testing.T) {
	var cluster *v1beta1.TemporalCluster
	var worker *v1beta1.TemporalWorkerProcess

	appWorkerFeature := features.New("Deploy and test worker process").
		Setup(func(ctx context.Context, tt *testing.T, cfg *envconf.Config) context.Context {
			namespace := GetNamespaceForFeature(ctx)

			client, err := cfg.NewClient()
			if err != nil {
				t.Fatal(err)
			}

			// create the postgres
			err = deployAndWaitForPostgres(ctx, cfg, namespace)
			if err != nil {
				t.Fatal(err)
			}

			connectAddr := fmt.Sprintf("postgres.%s:5432", namespace)

			// create the temporal cluster
			cluster = &v1beta1.TemporalCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: namespace,
				},
				Spec: v1beta1.TemporalClusterSpec{
					NumHistoryShards:           1,
					JobTTLSecondsAfterFinished: &jobTTL,
					Persistence: v1beta1.TemporalPersistenceSpec{
						DefaultStore: &v1beta1.DatastoreSpec{
							SQL: &v1beta1.SQLSpec{
								User:            "temporal",
								PluginName:      "postgres",
								DatabaseName:    "temporal",
								ConnectAddr:     connectAddr,
								ConnectProtocol: "tcp",
							},
							PasswordSecretRef: v1beta1.SecretKeyReference{
								Name: "postgres-password",
								Key:  "PASSWORD",
							},
						},
						VisibilityStore: &v1beta1.DatastoreSpec{
							SQL: &v1beta1.SQLSpec{
								User:            "temporal",
								PluginName:      "postgres",
								DatabaseName:    "temporal_visibility",
								ConnectAddr:     connectAddr,
								ConnectProtocol: "tcp",
							},
							PasswordSecretRef: v1beta1.SecretKeyReference{
								Name: "postgres-password",
								Key:  "PASSWORD",
							},
						},
					},
				},
			}
			err = client.Resources(namespace).Create(ctx, cluster)
			if err != nil {
				t.Fatal(err)
			}

			return SetTemporalClusterForFeature(ctx, cluster)
		}).
		Assess("Temporal cluster created", AssertTemporalClusterReady()).
		Assess("Can create a TemporalNamespace", AssertCanCreateTemporalNamespace("default")).
		Assess("TemporalNamespace ready", AssertTemporalNamespaceReady()).
		Assess("Can create a temporal worker process", func(ctx context.Context, tt *testing.T, cfg *envconf.Config) context.Context {
			namespace := GetNamespaceForFeature(ctx)

			// create the temporal worker process
			worker = &v1beta1.TemporalWorkerProcess{
				ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: namespace},
				Spec: v1beta1.TemporalWorkerProcessSpec{
					Version:                    "latest",
					JobTTLSecondsAfterFinished: &workerProcessJobTTL,
					Replicas:                   &replicas,
					Image:                      "example-worker-process", // image built when starting e2e
					PullPolicy:                 corev1.PullIfNotPresent,
					ClusterRef: &v1beta1.TemporalClusterReference{
						Name:      "test",
						Namespace: namespace,
					},
				},
			}
			err := cfg.Client().Resources(namespace).Create(ctx, worker)
			if err != nil {
				t.Fatal(err)
			}

			return ctx
		}).
		Assess("Temporal worker process created", func(ctx context.Context, tt *testing.T, cfg *envconf.Config) context.Context {
			err := waitForWorkerProcess(ctx, cfg, worker)
			if err != nil {
				t.Fatal(err)
			}
			return ctx
		}).
		Feature()

	testenv.Test(t, appWorkerFeature)
}
