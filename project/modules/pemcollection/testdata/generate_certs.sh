rm -f *.pem
rm -rf root_ca
rm -rf intm_ca
rm -rf int0_ca
rm -rf int1_ca
rm -rf client

KS=2048

# ROOT CA: Init
mkdir root_ca
pushd root_ca && mkdir certs crl newcerts private && touch index.txt && echo 23 > serial && popd

pushd root_ca
openssl genrsa -out private/ca.key.pem ${KS}
openssl req -config ../root.cnf -extensions v3_ca \
  -key private/ca.key.pem \
  -x509 -new -nodes -sha256 -days 1825 -out certs/ca.cert.pem -subj "/CN=mockup root CA"
popd

# INT0 CA: Init
mkdir int0_ca
pushd int0_ca && mkdir certs crl newcerts private && touch index.txt && echo 1000 > serial && popd

pushd int0_ca
openssl genrsa -out private/intermediate.key.pem ${KS}
openssl req -config ../intm.cnf -extensions v3_intermediate_ca \
  -key private/intermediate.key.pem \
  -new -nodes -out ../root_ca/int0CA.req -subj "/CN=mockup intermediate CA"
popd

# ROOT signs INT0
pushd root_ca
openssl ca -config ../root.cnf -extensions v3_intermediate_ca \
    -days 3650 -notext -md sha256 -in int0CA.req \
    -policy policy_loose -batch \
    -out ../int0_ca/certs/intermediate.cert.pem
popd

# INT1: Init
mkdir int1_ca
pushd int1_ca && mkdir certs crl newcerts private && touch index.txt && echo 9000 > serial && popd

pushd int1_ca
openssl genrsa -out private/intermediate.key.pem ${KS}
openssl req -config ../intm.cnf -extensions v3_intermediate_ca \
  -key private/intermediate.key.pem \
  -new -nodes -out ../int0_ca/int1CA.req -subj "/CN=mockup signing CA"
popd

# INT 0 signs INT1
pushd int0_ca
openssl ca -config ../intm.cnf -extensions v3_intermediate_ca \
    -days 3650 -notext -md sha256 -in int1CA.req \
    -policy policy_loose -batch \
    -out ../int1_ca/certs/intermediate.cert.pem
popd

# INT 1 signs a server cert
mkdir client
pushd client
openssl genrsa -out client.key.pem ${KS}
openssl req -config ../intm.cnf -extensions server_cert \
  -key client.key.pem \
  -new -nodes -out ../int1_ca/client.req -subj "/CN=Client Foobar"
openssl rsa -des3 -in client.key.pem -passout "pass:VerySecretPassword" -out client_enc.key.pem
popd

pushd int1_ca
openssl ca -config ../intm.cnf -extensions server_cert \
    -days 3650 -notext -md sha256 -in client.req \
    -policy policy_loose -batch \
    -out ../client/client.cert.pem
popd

# collect stuff
cat root_ca/certs/ca.cert.pem \
  int0_ca/certs/intermediate.cert.pem \
  int1_ca/certs/intermediate.cert.pem > client/ca_chain.pem


# verify, to make sure everything is OK
openssl verify -CAfile client/ca_chain.pem client/client.cert.pem

# create the noisy things
openssl pkcs12 -export -passout "pass:VerySecretPassword" -des3 \
 -export -out client/packaged.p12 \
 -name "Client Foobar" \
 -inkey client/client.key.pem -in client/client.cert.pem -certfile client/ca_chain.pem

openssl pkcs12 -passin "pass:VerySecretPassword" -des3 -passout "pass:VerySecretPassword" \
 -in client/packaged.p12 \
 -out client/packaged.pem
