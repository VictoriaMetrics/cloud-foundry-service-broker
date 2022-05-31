# VictoriaMetrics Cloud foundry service broker

This is a service broker implementation for the VictoriaMetrics intergration with Cloud Foundry.

This service broker is using VMs at cloud providers to spin up VictoriaMetrics instances.
Currently it is in an early development, first goals are:

- Support creating instances in AWS
- VM Single edition

Later on it will support:

- Multicloud: AWS, GCP, Azure, etc
- Different editions of VictoriaMetrics

# Deployment

Project is in early development stage and doesn't ship CF installation yet.

## Migrations

Migrations are based on [migrate](https://github.com/golang-migrate/migrate) library.
In order to be able to run migrations, you need to have `migrations` folder content copied next to binary.
If folder will be absent, migrations will be skipped.