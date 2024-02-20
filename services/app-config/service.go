package appconfig

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
)

type AppConfigData struct {
	Name string
	Id   string
}

type AppConfigDataSlice []AppConfigData

var (
	AWSCfg              aws.Config
	AppConfigClient     *appconfig.Client
	AppConfigDataClient *appconfigdata.Client
	AppConfigApps       AppConfigDataSlice
)

func (acd AppConfigDataSlice) findByName(s string) (AppConfigData, bool) {
	for _, a := range acd {
		if a.Name == s {
			return a, true
		}
	}

	return AppConfigData{}, false
}

func (acd AppConfigDataSlice) toNameSlice() []string {
	ret := make([]string, len(acd))
	for i, a := range acd {
		ret[i] = a.Name
	}
	return ret
}

func InitAWS() {
	var err error
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	AWSCfg = cfg
	AppConfigClient = appconfig.NewFromConfig(cfg)
	AppConfigDataClient = appconfigdata.NewFromConfig(cfg)
	AppConfigApps = getApps()
}

func getApps() AppConfigDataSlice {
	data, err := AppConfigClient.ListApplications(context.Background(), &appconfig.ListApplicationsInput{})
	if err != nil {
		panic(err)
	}

	apps := []AppConfigData{}
	for _, a := range data.Items {
		apps = append(apps, AppConfigData{
			Name: *a.Name,
			Id:   *a.Id,
		})
	}

	sort.Slice(apps, func(i, j int) bool {
		return apps[i].Name < apps[j].Name
	})
	return apps
}

func getEnvs(appId string) AppConfigDataSlice {
	data, err := AppConfigClient.ListEnvironments(context.Background(), &appconfig.ListEnvironmentsInput{
		ApplicationId: &appId,
	})
	if err != nil {
		panic(err)
	}

	envs := []AppConfigData{}
	for _, a := range data.Items {
		envs = append(envs, AppConfigData{
			Name: *a.Name,
			Id:   *a.Id,
		})
	}

	sort.Slice(envs, func(i, j int) bool {
		return envs[i].Name < envs[j].Name
	})
	return envs
}

func getConfigs(appId string) AppConfigDataSlice {
	data, err := AppConfigClient.ListConfigurationProfiles(
		context.Background(),
		&appconfig.ListConfigurationProfilesInput{
			ApplicationId: &appId,
		})
	if err != nil {
		panic(err)
	}

	configs := []AppConfigData{}
	for _, a := range data.Items {
		configs = append(configs, AppConfigData{
			Name: *a.Name,
			Id:   *a.Id,
		})
	}

	sort.Slice(configs, func(i, j int) bool {
		return configs[i].Name < configs[j].Name
	})
	return configs
}

func getDeployedConfig(appId string, envId string, configProfileId string) (string, error) {
	sessInfo, err := AppConfigDataClient.StartConfigurationSession(
		context.Background(),
		&appconfigdata.StartConfigurationSessionInput{
			ApplicationIdentifier:          &appId,
			ConfigurationProfileIdentifier: &configProfileId,
			EnvironmentIdentifier:          &envId,
		},
	)
	if err != nil {
		return "", err
	}
	token := sessInfo.InitialConfigurationToken

	data, err := AppConfigDataClient.GetLatestConfiguration(
		context.Background(),
		&appconfigdata.GetLatestConfigurationInput{
			ConfigurationToken: token,
		})
	if err != nil {
		return "", err
	}

	return string(data.Configuration), nil
}
