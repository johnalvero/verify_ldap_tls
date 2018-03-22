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

Parameter | Description
---- | ----
Root CA Cert Path | Path to the AD's CA Cert
URI | hostname/IP and port address of the AD e.g. ldap.example.com:389
TLS Server Name | AD's server name for verification
Bind User | User name used for authentication
Bind Password | Password used for authentication

If a successful connectivity is established, the following will be displayed:
```
OK: Connectivity verified
```

## Common Failure Scenarios

Root CA not found
```
2018/03/22 03:24:11 Couldn't load file%!(EXTRA *os.PathError=open /etc/openldap/cacerts/ca.pem: no such file or directory)
```

Connectivity error
```
2018/03/22 03:27:11 LDAP Result Code 200 "Network Error": dial tcp 1.1.1.1:389: i/o timeout
```

Invalid Credentials
```
2018/03/22 03:28:56 LDAP Result Code 49 "Invalid Credentials": 80090308: LdapErr: DSID-0C09042A, comment: AcceptSecurityContext error, data 52e, v3839
```
