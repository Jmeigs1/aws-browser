package appconfig

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type appConfigViewerState struct {
	AppId      string
	EnvId      string
	ConfigId   string
	EnvList    AppConfigDataSlice
	ConfigList AppConfigDataSlice
}

var state appConfigViewerState = appConfigViewerState{}

func Start() *fyne.Container {
	// Data display
	configData := binding.NewString()
	input := widget.NewMultiLineEntry()
	input.Bind(configData)
	input.SetMinRowsVisible(18)

	submitButton := widget.NewButton("Submit", handleSubmit(configData))

	envSelector := &widget.Select{}
	envSelector.Options = []string{}
	envSelector.Disable()
	envSelector.OnChanged = func(s string) {
		env, ok := state.EnvList.findByName(s)
		if ok {
			state.EnvId = env.Id
		}
	}
	envSelector.PlaceHolder = "Environment..."

	configSelector := &widget.Select{}
	configSelector.Options = []string{}
	configSelector.Disable()
	configSelector.OnChanged = func(s string) {
		cfg, ok := state.ConfigList.findByName(s)
		if ok {
			state.ConfigId = cfg.Id
		}
	}
	configSelector.PlaceHolder = "Config..."

	applicationSelector := widget.NewSelect(
		AppConfigApps.toNameSlice(),
		handleAppSelect(envSelector, configSelector),
	)
	applicationSelector.PlaceHolder = "Application..."

	optionsContainer := container.NewGridWithColumns(
		2,
		widget.NewLabel("Application"),
		applicationSelector,
		widget.NewLabel("Environment"),
		envSelector,
		widget.NewLabel("Config Profile"),
		configSelector,
	)

	return container.New(layout.NewVBoxLayout(), input, optionsContainer, submitButton)
}

func handleSubmit(b binding.String) func() {
	return func() {
		data, err := getDeployedConfig(state.AppId, state.EnvId, state.ConfigId)
		if err != nil {
			return
		}
		b.Set(data)
	}
}

func handleAppSelect(envSelector *widget.Select, configSelector *widget.Select) func(string) {
	return func(s string) {
		app, ok := AppConfigApps.findByName(s)
		if !ok {
			// TODO handle errors better
			return
		}

		configs := getConfigs(app.Id)
		envs := getEnvs(app.Id)

		configSelector.Options = configs.toNameSlice()
		configSelector.ClearSelected()
		configSelector.Enable()
		configSelector.Refresh()

		envSelector.SetOptions(envs.toNameSlice())
		envSelector.ClearSelected()
		envSelector.Enable()
		envSelector.Refresh()

		state.AppId = app.Id
		state.ConfigId = ""
		state.EnvId = ""
		state.ConfigList = configs
		state.EnvList = envs
	}
}
