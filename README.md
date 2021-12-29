# cocus_challenger
![build](https://github.com/miguelhbrito/cocus_challenger/actions/workflows/go.yaml/badge.svg?branch=master)

## 🧰 Configuration

To install golang just follow the steps from website:
- https://golang.org/doc/install

To install docker and docker-compose just follow the steps from website:
- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

To install postman just follow the steps from website:
- https://www.postman.com/downloads/

To install jq just follow the steps from website:
- https://stedolan.github.io/jq/

To import the collection and environment from repository on folder postman into your postman app please follow the steps:
- To open the Postman application, click on its icon on the taskbar.
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/postmanTaskbar.png" width="47" height="40">

- Click on the file tab and then click import.
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/postmanFileImport.png" width="255" height="296">

- Choose the import file method and then click on "Upload files".
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/postmanImportMethod.png" width="786" height="480">

- Choose the correct items to import and press open. Postman will automatically import the items.
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/cocusImportPostman.png" width="788" height="480">

## 🛠 How to use

Start vault:
``` powershell
make config-up
```
Generate secrets from vault:
``` powershell
make generate-secrets
```
Run application:
``` powershell
make run-cocus-gateway
```
Run build application:
``` powershell
make run-cocus-gateway-build
```
To shutdown vault:
``` powershell
make config-down
```
To test and test with coverage:
``` powershell
make test
make test-cover
```
To clean binary
``` powershell
make clean
```

##### `/login/create` POST to create a account
##### `/login` POST to get login token
##### `/transfers` POST to create a new triangle

- First step is create a new user to login into system
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/createNewUser.png" width="620" height="365">

- Then login into system to get token auth
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/cocusLoginIntoSystem.png" width="620" height="325">

- Token is automatically saved
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/postmanLoginToken.png" width="595" height="322">

- Token is included in request's header
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/cocusTokenHeader.png" width="557" height="313">

- Now you are able to create and get all triangles from system
<img src="https://github.com/miguelhbrito/cocus_challenger/blob/feature/new-arc/images/cocusCreateNewTriangle.png" width="617" height="319">

