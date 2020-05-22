# Open Space: Readiness probes

The point of the demo is to observe how K8S [readiness checks](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) affect a K8S [service](https://kubernetes.io/docs/concepts/services-networking/service/).

By deploying a mixture of 'stable' and 'flaky' webservers to a cluster and setting up a `LoadBalancer` service to expose them, it is possible to observe how the webservers that fail their readiness checks are dynamically removed from the service.

The demo also uses [cockpit](https://cockpit-project.org/guide/133/feature-kubernetes.html) to visualize the service and the pods that it selects.

Contains:

* `stable/` - A webserver whose `/status` endpoint always returns success.
* `flaky/` - A webserver whose `/status` endpoint returns intermittent errors.
* `resources/` - K8S manifests to deploy the various resources required for the demo.
* `cockpit` - Manifest to install [cockpit](https://cockpit-project.org/guide/133/feature-kubernetes.html) into the cluster.
* `get.sh` - Script to run locally to hit the loadbalanced service.
