# !/bin/bash
set -e
apt-get install apt-transport-https ca-certificates
echo "deb [trusted=yes] https://packages.netsec.inf.ethz.ch/debian all main" | tee /etc/apt/sources.list.d/scionlab.list
apt-get update -y


mkdir /etc/scion

cat <<EOT >> /etc/scion/bootstrapper.toml
# The folder where the retrieved topology and certificates are stored (default ".")
sciond_config_dir = "/etc/scion"

# Set the verification behavior of the signature of the configuration file using the TRC (default permissive)
security_mode = "insecure"

# Set the crypto engine to use for the signature verification, options are 'openssl' and 'native' (default native)
crypto_engine = "openssl"

# Discovery mechanisms
[mock]
    # Whether to enable the fake discovery or not (default false)
    # This discovery mechanisms is used for testing purposes
    enable = true
    # The address to return when simulating a network discovery (default "")
    address = "141.44.25.148:8041"
[dhcp]
    # Whether to enable DHCP discovery or not (default false)
    enable = false
[dhcpv6]
    # Whether to enable DHCPv6 discovery or not (default false)
    enable = false
    # Set the DHCPv6 Unique Identifier type (default "DUID-LLT")
    DUID_type = "DUID-LL"
    # Set a static, fixed DUID
    # Overrides the DUID_type setting, to be used for setting DUID-EN, DUID-UUID or debugging
    # Fixed hex string for the client DUID, no separators, no 0x prefix:
    #client_id = "0001000100000000deadbeefaabb"
[ipv6]
    # Whether to enable IPv6 Neighbor Discovery Protocol (NDP) for the
    # Router Advertisement DNSSL and RDNS discovery or not (default false)
    enable = false
[dnssd]
    # Whether to enable DNS SRV discovery or not (default false)
    enable_srv = false
    # Whether to enable DNS-SD discovery or not (default false)
    enable_sd = false
    # Whether to enable DNS-NAPTR discovery or not (default false)
    enable_naptr = false
[mdns]
    # Whether to enable mDNS discovery or not (default false)
    enable = false
[log]
    [log.console]
        # Console logging level (debug|info|error) (default info)
        level = "debug"
EOT


apt-get install scion-bootstrapper -y
/usr/bin/bootstrapper  -config /etc/scion/bootstrapper.toml
systemctl restart scion-daemon.service
systemctl restart scion-dispatcher.service
