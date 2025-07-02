# NetData Provider Examples

## Provider Configuration

```bash
NETDATA_TOKEN=<your-token-here>
cat examples/providerconfig/secret.yaml.tmpl | sed -e "s/t0ps3cr3t11/${NETDATA_TOKEN}/g" > examples/providerconfig/secret.yaml

kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

## NetData Space Resource

```bash
kubectl apply -f examples/space/space.yaml
```

## NetData Room Resource

```bash
kubectl apply -f examples/room/room.yaml
```