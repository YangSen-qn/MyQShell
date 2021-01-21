package param_cmd

import (
	"github.com/spf13/cobra"
	"qshell/cmd/common"
	"qshell/cmd/execute"
)

type ParamCMDConfig struct {
	Use                    string
	Short                  string
	Long                   string
	Version                string
	Example                string
	BashCompletionFunction string
}

type IParamCMD interface {
	FlagsStringVar(p *string, name, shorthand string, value string, usage string)
	FlagsBoolVar(p *bool, name, shorthand string, value bool, usage string)
	FlagsInt64Var(p *int64, name, shorthand string, value int64, usage string)
	GetSupperCMD() IParamCMD
	AddCMD(subCMD IParamCMD)

	setSupperCMD(supperCMD IParamCMD)
	getCobraCMD() *cobra.Command
}

type ParamCMD struct {
	*execute.Command

	supperCMD IParamCMD
	cmd       *cobra.Command

	config *common.Config
}

func NewParamCMD() *ParamCMD {
	cmd := &ParamCMD{
		Command:   execute.NewCommand(),
		supperCMD: nil,
		cmd:       &cobra.Command{},
		config:    common.NewConfig(),
	}
	return cmd
}

func (paramCMD *ParamCMD) ConfigParamCMDExecuteConfig(commandConfig execute.CommandConfig) {
	paramCMD.ConfigCommand(commandConfig)
}

func (paramCMD *ParamCMD) ConfigParamCMDParseConfig(paramCMDConfig ParamCMDConfig) {

	paramCMD.cmd = &cobra.Command{
		Use:                    paramCMDConfig.Use,
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  paramCMDConfig.Short,
		Long:                   paramCMDConfig.Long,
		Example:                paramCMDConfig.Example,
		ValidArgs:              nil,
		ValidArgsFunction:      nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: paramCMDConfig.BashCompletionFunction,
		Deprecated:             "",
		Hidden:                 false,
		Annotations:            nil,
		Version:                paramCMDConfig.Version,
		PersistentPreRun:       nil,
		PersistentPreRunE:      nil,
		PreRun:                 nil,
		PreRunE:                nil,
		Run: func(cmd *cobra.Command, args []string) {
			context := common.NewQShellContext(cmd.Context())
			context.SetConfig(paramCMD.config)
			execute.Execute(paramCMD, context)
		},
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
		TraverseChildren:           false,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	}

	paramCMD.setupDefaultParseConfig()
}

func (paramCMD *ParamCMD) setupDefaultParseConfig() {
	paramCMD.FlagsStringVar(&paramCMD.config.OutputFormatValue, "output-format", "", "", "")
}

func (cmd *ParamCMD) FlagsStringVar(p *string, name, shorthand string, value string, usage string) {
	cmd.cmd.Flags().StringVarP(p, name, shorthand, value, usage)
}

func (cmd *ParamCMD) FlagsBoolVar(p *bool, name, shorthand string, value bool, usage string) {
	cmd.cmd.Flags().BoolVarP(p, name, shorthand, value, usage)
}

func (cmd *ParamCMD) FlagsInt64Var(p *int64, name, shorthand string, value int64, usage string) {
	cmd.cmd.Flags().Int64VarP(p, name, shorthand, value, usage)
}

func (cmd *ParamCMD) AddCMD(subCMD IParamCMD) {
	subCMD.setSupperCMD(subCMD)
	cmd.cmd.AddCommand(subCMD.getCobraCMD())
}

func (cmd *ParamCMD) GetSupperCMD() IParamCMD {
	return cmd.supperCMD
}

func (cmd *ParamCMD) setSupperCMD(supperCMD IParamCMD) {
	cmd.supperCMD = supperCMD
}

func (cmd *ParamCMD) getCobraCMD() *cobra.Command {
	return cmd.cmd
}

func (cmd *ParamCMD) CobraExecute() error {
	return cmd.cmd.Execute()
}
