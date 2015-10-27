all: zabbix_socket.so

zabbix_socket.so: zabbix_socket.go socket.go
	go build -buildmode=c-shared -o zabbix_socket.so zabbix_socket.go socket.go
	@chmod +x zabbix_socket.so

clean:
	rm -f zabbix_socket.h zabbix_socket.so
