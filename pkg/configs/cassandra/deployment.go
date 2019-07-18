package cassandra

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
)

var deployment = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: contrail-configdb
  namespace: contrail-system
  labels:
    app: contrail-configdb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cassandra-configdb
      cassandra_cr: cassandra-configdb
  template:
    metadata:
      labels:
        app: cassandra-configdb
        cassandra_cr: cassandra-configdb
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: "node-role.kubernetes.io/infra"
                operator: Exists
      hostNetwork: true
      initContainers:
      - name: contrail-node-init
        image: "opencontrailnightly/contrail-node-init:latest"
        imagePullPolicy: "IfNotPresent"
        securityContext:
          privileged: true
        env:
        - name: NODE_TYPE
          value: "config-database"
        - name: CONTRAIL_STATUS_IMAGE
          value: "opencontrailnightly/contrail-status:latest"
        envFrom:
        - configMapRef:
            name: contrail-configdb-config  
        volumeMounts:
        - mountPath: /host/usr/bin
          name: host-usr-bin
      containers:
      - name: contrail-configdb
        image: "opencontrailnightly/contrail-external-cassandra:latest"
        imagePullPolicy: "IfNotPresent"
        securityContext:
          privileged: true
        env:
        - name: NODE_TYPE
          value: config-database
        - name: CASSANDRA_SEEDS
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        envFrom:
        - configMapRef:
            name: contrail-configdb-config
        volumeMounts:
        - mountPath: /var/lib/cassandra
          name: configdb-data
        - mountPath: /var/log/cassandra
          name: configdb-log
      volumes:
      - name: host-usr-bin
        hostPath:
          path: /usr/bin
      - name: configdb-data
        hostPath:
          path: /var/lib/contrail/configdb
      - name: configdb-log
        hostPath:
          path: /var/log/contrail/configdb`

func GetDeployment() *appsv1.Deployment {
	fileData := deployment
	yamlData := string(fileData)
	dp := appsv1.Deployment{}
	err := yaml.Unmarshal([]byte(yamlData), &dp)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &dp)
	if err != nil {
		panic(err)
	}
	return &dp
}
