package openshift

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	lokiv1 "github.com/grafana/loki/operator/apis/loki/v1"
	"github.com/grafana/loki/operator/internal/external/k8s"
	"github.com/grafana/loki/operator/internal/manifests"
	"github.com/grafana/loki/operator/internal/manifests/openshift"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlertManagerSVCExists returns true if the Openshift AlertManager is present in the cluster.
func AlertManagerSVCExists(ctx context.Context, opts manifests.Options, k k8s.Client) (bool, error) {
	if opts.Stack.Tenants != nil && opts.Stack.Tenants.Mode == lokiv1.OpenshiftLogging {
		var svc corev1.Service
		key := client.ObjectKey{Name: openshift.MonitoringSVCOperated, Namespace: openshift.MonitoringNS}

		err := k.Get(ctx, key, &svc)
		if err != nil && !apierrors.IsNotFound(err) {
			return false, kverrors.Wrap(err, "failed to lookup alertmanager service", "name", key)
		}

		return err == nil, nil
	}

	return false, nil
}