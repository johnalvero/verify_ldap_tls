# Verify TLS Connectivity

## Usage

1. Clone the repo
```
git clone https://github.com/johnalvero/verify_ldap_tls.git
```
2. Build
```
cd verify_ldap_tls
go build
```
3. Run
```
./verify_ldap_tls
```
4. Supply the necessary details
Root CA Cert Path - Path to the AD's CA Cert
URI - hostname/IP and port address of the AD e.g. ldap.example.com:389
TLS Server Name - AD's server name for verification
Bind User - User name used for authentication
Bind Password - Password used for authentication

If a successful connectivity is established, the following will be displayed:
```
OK: Connectivity verified
```
