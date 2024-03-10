package cmd

import (
	"fmt"
	"kms/app/config"
	"kms/database/sqldb"
	"strings"
)

// GCPCloudRunDeployImage builds arguments for running a service on
// Cloud Run given an Artifact Registry image.
func GCPCloudRunDeployImage(f *config.Config, image GCPArtifactRegistryContainerImage) []string {

	var (
		// Google Cloud Run Service Name
		serviceName = f.GCP.CloudRun.ServiceName
		// Google Cloud SQL Instance Name
		gcpCloudSQLInstanceConnectionName = f.GCP.CloudSQL.InstanceConnectionName
		// postgresql database name
	)

	args := []string{"run", "deploy", serviceName, "--image", image.String(), "--platform", "managed", "--no-allow-unauthenticated"}

	args = append(args, "--add-cloudsql-instances", gcpCloudSQLInstanceConnectionName)

	icn := fmt.Sprintf(`INSTANCE-CONNECTION-NAME=%s`, gcpCloudSQLInstanceConnectionName)
	dbName := fmt.Sprintf(`%s=%s`, sqldb.DBNameEnv, f.DB.DBName)
	dbUser := fmt.Sprintf(`%s=%s`, sqldb.DBUserEnv, f.DB.User)
	dbPassword := fmt.Sprintf(`%s=%s`, sqldb.DBPasswordEnv, f.DB.Password)
	dbHost := fmt.Sprintf(`%s=%s`, sqldb.DBHostEnv, f.DB.Host)
	dbPort := fmt.Sprintf(`%s=%s`, sqldb.DBPortEnv, f.DB.Port)
	dbSearchPath := fmt.Sprintf(`%s=%s`, sqldb.DBSearchPathEnv, f.DB.SearchPath)
	encryptKey := fmt.Sprintf(`%s=%s`, encryptKeyEnv, f.EncryptionKey)

	envVars := []string{icn, dbName, dbUser, dbPassword, dbHost, dbPort, dbSearchPath, encryptKey}

	args = append(args, "--set-env-vars", strings.Join(envVars, ","))

	return args
}

// GCPArtifactRegistryContainerImage defines a GCP Artifact Registry
// build image according to https://cloud.google.com/artifact-registry/docs/docker/names
// The String method prints the build string needed to build to
// Artifact Registry using gcloud as well as deploy it to Cloud Run.
type GCPArtifactRegistryContainerImage struct {
	ProjectID          string
	RepositoryLocation string
	RepositoryName     string
	ImageName          string
	ImageTag           string
}

// String outputs the Google Artifact Registry image name.
// LOCATION-docker.pkg.dev/PROJECT-ID/REPOSITORY/IMAGE:TAG
func (i GCPArtifactRegistryContainerImage) String() string {
	if i.ImageTag != "" {
		return fmt.Sprintf("%s-docker.pkg.dev/%s/%s/%s:%s", i.RepositoryLocation, i.ProjectID, i.RepositoryName, i.ImageName, i.ImageTag)
	}
	return fmt.Sprintf("%s-docker.pkg.dev/%s/%s/%s", i.RepositoryLocation, i.ProjectID, i.RepositoryName, i.ImageName)
}
