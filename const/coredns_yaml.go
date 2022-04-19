package customConst

const (
	// CoreDnsYaml coredns的编排文件
	CoreDnsYaml = "# Warning: This is a file generated from the base underscore template file: coredns.yaml.base\n\n" +
		"apiVersion: v1\n" +
		"kind: ServiceAccount\n" +
		"metadata:\n" +
		"  name: coredns\n" +
		"  namespace: kube-system\n" +
		"  labels:\n" +
		"      kubernetes.io/cluster-service: \"true\"\n" +
		"      addonmanager.kubernetes.io/mode: Reconcile\n" +
		"---\n" +
		"apiVersion: rbac.authorization.k8s.io/v1\n" +
		"kind: ClusterRole\n" +
		"metadata:\n" +
		"  labels:\n" +
		"    kubernetes.io/bootstrapping: rbac-defaults\n" +
		"    addonmanager.kubernetes.io/mode: Reconcile\n" +
		"  name: system:coredns\n" +
		"rules:\n" +
		"- apiGroups:\n" +
		"  - \"\"\n" +
		"  resources:\n" +
		"  - endpoints\n" +
		"  - services\n" +
		"  - pods\n" +
		"  - namespaces\n" +
		"  verbs:\n" +
		"  - list\n" +
		"  - watch\n" +
		"- apiGroups:\n" +
		"  - \"\"\n" +
		"  resources:\n" +
		"  - nodes\n" +
		"  verbs:\n" +
		"  - get\n" +
		"- apiGroups:\n" +
		"  - discovery.k8s.io\n" +
		"  resources:\n" +
		"  - endpointslices\n" +
		"  verbs:\n" +
		"  - list\n" +
		"  - watch\n" +
		"---\n" +
		"apiVersion: rbac.authorization.k8s.io/v1\n" +
		"kind: ClusterRoleBinding\n" +
		"metadata:\n" +
		"  annotations:\n" +
		"    rbac.authorization.kubernetes.io/autoupdate: \"true\"\n" +
		"  labels:\n" +
		"    kubernetes.io/bootstrapping: rbac-defaults\n" +
		"    addonmanager.kubernetes.io/mode: EnsureExists\n" +
		"  name: system:coredns\n" +
		"roleRef:\n" +
		"  apiGroup: rbac.authorization.k8s.io\n" +
		"  kind: ClusterRole\n" +
		"  name: system:coredns\n" +
		"subjects:\n" +
		"- kind: ServiceAccount\n" +
		"  name: coredns\n" +
		"  namespace: kube-system\n" +
		"---\n" +
		"apiVersion: v1\n" +
		"kind: ConfigMap\n" +
		"metadata:\n" +
		"  name: coredns\n" +
		"  namespace: kube-system\n" +
		"  labels:\n" +
		"      addonmanager.kubernetes.io/mode: EnsureExists\n" +
		"data:\n" +
		"  Corefile: |\n" +
		"    .:53 {\n" +
		"        errors\n" +
		"        health {\n" +
		"            lameduck 5s\n" +
		"        }\n" +
		"        ready\n" +
		"        kubernetes $DNS_DOMAIN in-addr.arpa ip6.arpa {\n" +
		"            pods insecure\n" +
		"            fallthrough in-addr.arpa ip6.arpa\n" +
		"            ttl 30\n" +
		"        }\n" +
		"        prometheus :9153\n" +
		"        forward . /etc/resolv.conf {\n" +
		"            max_concurrent 1000\n" +
		"        }\n" +
		"        cache 30\n" +
		"        loop\n" +
		"        reload\n" +
		"        loadbalance\n" +
		"    }\n" +
		"---\n" +
		"apiVersion: apps/v1\n" +
		"kind: Deployment\n" +
		"metadata:\n" +
		"  name: coredns\n" +
		"  namespace: kube-system\n" +
		"  labels:\n" +
		"    k8s-app: kube-dns\n" +
		"    kubernetes.io/cluster-service: \"true\"\n" +
		"    addonmanager.kubernetes.io/mode: Reconcile\n" +
		"    kubernetes.io/name: \"CoreDNS\"\n" +
		"spec:\n" +
		"  # replicas: not specified here:\n" +
		"  # 1. In order to make Addon Manager do not reconcile this replicas parameter.\n" +
		"  # 2. Default is 1.\n" +
		"  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.\n" +
		"  strategy:\n" +
		"    type: RollingUpdate\n" +
		"    rollingUpdate:\n" +
		"      maxUnavailable: 1\n" +
		"  selector:\n" +
		"    matchLabels:\n" +
		"      k8s-app: kube-dns\n" +
		"  template:\n" +
		"    metadata:\n" +
		"      labels:\n" +
		"        k8s-app: kube-dns\n" +
		"    spec:\n" +
		"      securityContext:\n" +
		"        seccompProfile:\n" +
		"          type: RuntimeDefault\n" +
		"      priorityClassName: system-cluster-critical\n" +
		"      serviceAccountName: coredns\n" +
		"      affinity:\n" +
		"        podAntiAffinity:\n" +
		"          preferredDuringSchedulingIgnoredDuringExecution:\n" +
		"          - weight: 100\n" +
		"            podAffinityTerm:\n" +
		"              labelSelector:\n" +
		"                matchExpressions:\n" +
		"                  - key: k8s-app\n" +
		"                    operator: In\n" +
		"                    values: [\"kube-dns\"]\n" +
		"              topologyKey: kubernetes.io/hostname\n" +
		"      tolerations:\n" +
		"        - key: \"CriticalAddonsOnly\"\n" +
		"          operator: \"Exists\"\n" +
		"      nodeSelector:\n" +
		"        kubernetes.io/os: linux\n" +
		"      containers:\n" +
		"      - name: coredns\n" +
		"        image: k8s.gcr.io/coredns/coredns:v1.8.6\n" +
		"        imagePullPolicy: IfNotPresent\n" +
		"        resources:\n" +
		"          limits:\n" +
		"            memory: $DNS_MEMORY_LIMIT\n" +
		"          requests:\n" +
		"            cpu: 100m\n" +
		"            memory: 70Mi\n" +
		"        args: [ \"-conf\", \"/etc/coredns/Corefile\" ]\n" +
		"        volumeMounts:\n" +
		"        - name: config-volume\n" +
		"          mountPath: /etc/coredns\n" +
		"          readOnly: true\n" +
		"        ports:\n" +
		"        - containerPort: 53\n" +
		"          name: dns\n" +
		"          protocol: UDP\n" +
		"        - containerPort: 53\n" +
		"          name: dns-tcp\n" +
		"          protocol: TCP\n" +
		"        - containerPort: 9153\n" +
		"          name: metrics\n" +
		"          protocol: TCP\n" +
		"        livenessProbe:\n" +
		"          httpGet:\n" +
		"            path: /health\n" +
		"            port: 8080\n" +
		"            scheme: HTTP\n" +
		"          initialDelaySeconds: 60\n" +
		"          timeoutSeconds: 5\n" +
		"          successThreshold: 1\n" +
		"          failureThreshold: 5\n" +
		"        readinessProbe:\n" +
		"          httpGet:\n" +
		"            path: /ready\n" +
		"            port: 8181\n" +
		"            scheme: HTTP\n" +
		"        securityContext:\n" +
		"          allowPrivilegeEscalation: false\n" +
		"          capabilities:\n" +
		"            add:\n" +
		"            - NET_BIND_SERVICE\n" +
		"            drop:\n" +
		"            - all\n" +
		"          readOnlyRootFilesystem: true\n" +
		"      dnsPolicy: Default\n" +
		"      volumes:\n" +
		"        - name: config-volume\n" +
		"          configMap:\n" +
		"            name: coredns\n" +
		"            items:\n" +
		"            - key: Corefile\n" +
		"              path: Corefile\n" +
		"---\n" +
		"apiVersion: v1\n" +
		"kind: Service\n" +
		"metadata:\n" +
		"  name: kube-dns\n" +
		"  namespace: kube-system\n" +
		"  annotations:\n" +
		"    prometheus.io/port: \"9153\"\n" +
		"    prometheus.io/scrape: \"true\"\n" +
		"  labels:\n" +
		"    k8s-app: kube-dns\n" +
		"    kubernetes.io/cluster-service: \"true\"\n" +
		"    addonmanager.kubernetes.io/mode: Reconcile\n" +
		"    kubernetes.io/name: \"CoreDNS\"\n" +
		"spec:\n" +
		"  selector:\n" +
		"    k8s-app: kube-dns\n" +
		"  clusterIP: $DNS_SERVER_IP\n" +
		"  ports:\n" +
		"  - name: dns\n" +
		"    port: 53\n" +
		"    protocol: UDP\n" +
		"  - name: dns-tcp\n" +
		"    port: 53\n" +
		"    protocol: TCP\n" +
		"  - name: metrics\n" +
		"    port: 9153\n" +
		"    protocol: TCP\n" +
		""
)
