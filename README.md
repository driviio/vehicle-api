# vehicle-api

## Run locally
1) Install https://github.com/tockins/realize and ad to path
2) Create .env file with content and change to your GCP project ID
   ```
   DATASTORE_EMULATOR_HOST=localhost:8000
   DATASTORE_PROJECT_ID=<ProjectID>
   ```
3) Run shell script `./run_db.sh` to start datastore db in docker
3) Run shell script `./run_api.sh` to start api with watch mode

Now you should be able to access api on http://localhost:8080/