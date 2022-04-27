package nginx

const (
	// mainConf nginx主配置文件
	mainConf = "worker_processes  auto;\n" +
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
	// systemd nginx systemd管理脚本
	systemd = "[Unit]\n" +
		"Description=The nginx HTTP and reverse proxy server\n" +
		"After=network-online.target remote-fs.target nss-lookup.target\n" +
		"Wants=network-online.target\n\n" +
		"[Service]\n" +
		"Type=forking\n" +
		"PIDFile=/run/nginx.pid\n" +
		"xecStartPre=/usr/bin/rm -f /run/nginx.pid\n" +
		"ExecStartPre=/opt/nginx/sbin/nginx -t\n" +
		"ExecStart=/opt/nginx/sbin/nginx\n" +
		"ExecReload=/opt/nginx/sbin/nginx -s reload\n" +
		"illSignal=SIGQUIT\n" +
		"TimeoutStopSec=5\n" +
		"KillMode=process\n" +
		"PrivateTmp=true\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	nginxBuild = "./configure --prefix=%s --with-stream && make && make install"

	restartCmd = "chmod +x /opt/nginx/sbin/nginx && systemctl daemon-reload && systemctl enable nginx && systemctl restart nginx"
)
