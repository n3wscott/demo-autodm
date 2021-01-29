# demo-autodm

ko apply -f config/1-service.yaml

kubectl patch ksvc hello --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping", "value": "getrekt.dev"}]'

kubectl patch ksvc hello --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping.hello", "value": "hello.getrekt.dev"}]'

ko apply -f config/2-service.yaml

kubectl patch ksvc demo --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping.demo", "value": "demo.getrekt.dev"}]'

-- cleanup --

kubectl patch ksvc hello --type=json --patch '[{"op": "remove", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping.hello"}]'

kubectl delete ksvc demo

kubectl delete ksvc hello