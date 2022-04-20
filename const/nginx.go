package customConst

const (
	// NginxMainConf nginx主配置文件
	NginxMainConf = "worker_processes  auto;\n" +
		"error_log  /var/log/nginx-error.log warn;\n\n" +
		"pid /run/nginx.pid;\n\n" +
		"events {\n" +
		"    worker_connections  1024;\n" +
		"}\n" +
		"        \n" +
		"stream {\n" +
		"    log_format main \"$remote_addr  $upstream_addr  $time_local $status\";\n" +
		"    upstream kube_apiserver {\n" +
		"        least_conn;\n" +
		"upstreamConf" +
		"    }\n\n" +
		"    server {\n" +
		"        listen 16443;\n" +
		"        access_log /var/log/nginx-access.log main;\n" +
		"        proxy_pass    kube_apiserver;\n" +
		"        proxy_timeout 10m;\n" +
		"        proxy_connect_timeout 10s;\n" +
		"    }\n" +
		"}"
)
