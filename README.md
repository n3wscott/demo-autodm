# demo-autodm

ko apply -f config/service.yaml

kubectl patch ksvc hello --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping", "value": "getrekt.dev"}]'

kubectl patch ksvc hello --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping.hello", "value": "hello.getrekt.dev"}]'

ko apply -f config/deployment.yaml

kubectl patch svc helloworld --type=json --patch '[{"op": "add", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping", "value": "helloworld.getrekt.dev"}]'

-- cleanup --

kubectl patch ksvc hello --type=json --patch '[{"op": "remove", "path": "/metadata/annotations/sugar.knative.dev~1domainmapping.hello"}]'

kubectl delete ksvc demo

kubectl delete -f config/deployment.yaml