package user

import (
	"../../common"
	"../../execute"
	"github.com/spf13/cobra"
)

var (
	paramUserName         string
	paramAK               string
	paramSK               string
	paramIsListUser       bool
	paramIsAddUser        bool
	paramIsRemoveUser     bool
	paramIsSetCurrentUser bool
	paramIsGetCurrentUser bool
	paramConfig           = &common.Config{}
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "manager user",
	Run:   runFunction,
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "manager user",
	Run:   runFunction,
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "manager user",
	Run:   runFunction,
}

var userRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "manager user",
	Run:   runFunction,
}

var userChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "manager user",
	Run:   runFunction,
}

func AddCommand(superCommand *cobra.Command) {

	common.ConfigParam(userCmd, paramConfig)

	userCmd.Flags().StringVarP(&paramUserName, "name", "", "", "")
	userCmd.Flags().StringVarP(&paramAK, "ak", "", "", "")
	userCmd.Flags().StringVarP(&paramSK, "sk", "", "", "")
	userCmd.Flags().BoolVarP(&paramIsListUser, "list", "", false, "")
	userCmd.Flags().BoolVarP(&paramIsAddUser, "add", "", false, "")
	userCmd.Flags().BoolVarP(&paramIsRemoveUser, "remove", "", false, "")
	userCmd.Flags().BoolVarP(&paramIsSetCurrentUser, "change", "", false, "")
	userCmd.Flags().BoolVarP(&paramIsGetCurrentUser, "current", "", false, "")
	superCommand.AddCommand(userCmd)
}

func runFunction(cmd *cobra.Command, params []string) {

	var command execute.ICommand

	if paramIsListUser {
		command = &userListCMD{}
	} else if paramIsAddUser {
		command = &userAddCMD{
			name:      paramUserName,
			accessKey: paramAK,
			secretKey: paramSK,
		}
	} else if paramIsRemoveUser {
		command = &userRemoveCMD{
			name: paramUserName,
		}
	} else if paramIsSetCurrentUser {
		command = &currentUserSetCMD{
			name: paramUserName,
		}
	} else if paramIsGetCurrentUser {
		command = &currentUserGetCMD{}
	} else {
		command = &userGetCMD{
			name: paramUserName,
		}
	}

	command.SetConfig(paramConfig)
	execute.Execute(command)
}
