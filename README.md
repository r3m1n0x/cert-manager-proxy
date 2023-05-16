# cert-manager-proxy
A proxy to bypass the proxy-protocol-problem with the init probe from the cert-manager.

## the proxy

To run the reverse proxy, you can execute the following command:

´´´bash
go run main.go --backend=http://localhost:8080 --frontend=:80
´´´

Replace http://localhost:8080 with the URL of your backend server and :80 with the desired frontend server address.

Note: This code sets up a basic reverse proxy that speaks the Proxy Protocol. It doesn't include the specifics of integrating with Ingress Nginx or handling the issue mentioned in the GitHub link you provided. You may need to modify and extend the code based on your specific requirements and environment.

## How to use ...

