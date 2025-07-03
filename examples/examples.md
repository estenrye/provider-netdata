# NetData Provider Examples

## Build and Deploy Locally

```bash
make generate
make build
make local-deploy
kind get kubeconfig --name local-dev > ~/.kube/local-dev
export KUBECONFIG=~/.kube/local-dev
```

## Provider Configuration

```bash
NETDATA_TOKEN=<your-token-here>
cat examples/providerconfig/secret.yaml.tmpl | sed -e "s/t0ps3cr3t11/${NETDATA_TOKEN}/g" > examples/providerconfig/secret.yaml

kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

## NetData Space Resource

Every NetData monitoring experience starts with provisioning a Space.
Spaces are high-level containers for your entire infrastructure.

Space serves as your primary collaboration environment in Netdata Cloud. It allows you to:

- Organize team members and manage access levels.
- Connect nodes for monitoring.
- Create a unified monitoring environment.

Key Space Characteristics

- Each node can only belong to one Space.
- You can create multiple Spaces, but we recommend using a single Space for most use cases.
- All team members in a Space can access its monitoring data based on their assigned roles.

```bash
kubectl apply -f examples/space/space.yaml
```

## NetData Room Resource

Rooms provide flexible groupings within Spaces for specific monitoring needs.

Rooms are organizational units within a Space that provide:

- Infrastructure-wide dashboards.
- Real-time metrics visualization.
- Focused monitoring views.
- Flexible node grouping.

```bash
kubectl apply -f examples/room/room.yaml
```

## NetData Member Resource

Member objects are used to grant a user access to a NetData Space
using their email address for authentication.

```bash
kubectl apply -f examples/member/member.yaml
```

## NetData Room Member Resources

```bash
kubectl apply -f examples/roommember/member.yaml
```

## Environment Cleanup

```bash
kubectl delete -f examples/roommember/member.yaml
kubectl delete -f examples/member/member.yaml
kubectl delete -f examples/room/room.yaml
kubectl delete -f examples/space/space.yaml
kind delete cluster --name local-dev
```