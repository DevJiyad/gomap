package main

import "fmt"

func printPortService(port int) {
	var service string

	switch port {
	case 20:
		service = "FTP-DATA"
	case 21:
		service = "FTP"
	case 22:
		service = "SSH"
	case 23:
		service = "TELNET"
	case 25:
		service = "SMTP"
	case 53:
		service = "DOMAIN (DNS)"
	case 67, 68:
		service = "DHCP"
	case 69:
		service = "TFTP"
	case 80:
		service = "HTTP"
	case 88:
		service = "KERBEROS"
	case 110:
		service = "POP3"
	case 111:
		service = "RPCBIND"
	case 123:
		service = "NTP"
	case 135:
		service = "MSRPC"
	case 137, 138, 139:
		service = "NETBIOS"
	case 143:
		service = "IMAP"
	case 161:
		service = "SNMP"
	case 389:
		service = "LDAP"
	case 443:
		service = "HTTPS"
	case 445:
		service = "MICROSOFT-DS (SMB)"
	case 465:
		service = "SMTPS"
	case 514:
		service = "SYSLOG"
	case 587:
		service = "SUBMISSION"
	case 636:
		service = "LDAPS"
	case 873:
		service = "RSYNC"
	case 990:
		service = "FTPS"
	case 993:
		service = "IMAPS"
	case 995:
		service = "POP3S"
	case 1080:
		service = "SOCKS"
	case 1194:
		service = "OPENVPN"
	case 1433:
		service = "MS-SQL-S"
	case 1521:
		service = "ORACLE"
	case 2049:
		service = "NFS"
	case 2222:
		service = "SSH-ALT"
	case 2375, 2376:
		service = "DOCKER"
	case 3000:
		service = "DEV-SERVER"
	case 3306:
		service = "MYSQL"
	case 3389:
		service = "MS-WBT-SERVER (RDP)"
	case 4444:
		service = "METASPLOIT"
	case 5000:
		service = "UPNP"
	case 5432:
		service = "POSTGRESQL"
	case 5900:
		service = "VNC"
	case 5985, 5986:
		service = "WINRM"
	case 6379:
		service = "REDIS"
	case 6443:
		service = "KUBERNETES"
	case 8000:
		service = "HTTP-ALT"
	case 8080:
		service = "HTTP-PROXY"
	case 8443:
		service = "HTTPS-ALT"
	case 9000:
		service = "CSLISTENER"
	case 9200:
		service = "ELASTICSEARCH"
	case 9929:
		service = "NPING-ECHO"
	case 27017:
		service = "MONGOD"
	case 31337:
		service = "ELITE"
	default:
		service = "UNKNOWN"
	}

	fmt.Printf(Green+"[*] Port %d is %s\n"+Reset, port, service)
}
