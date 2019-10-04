/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"digimon-cli/handler"
	"digimon-cli/peer/wsconnection"
	"digimon-cli/tui"
	"fmt"
	"github.com/Analyse4/digimon/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: startHandler,
}

var (
	log *logrus.Entry
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	log = logger.GetLogger().WithField("pkg", "cmd")
}

func startHandler(cmd *cobra.Command, args []string) {
	dc := handler.New()
	dc.Register()
	dc.Show()
	ep := getEP()
	//TODO: init connection
	c := wsconnection.New()
	c.Connect(ep)
	//c := peer.ConnectGameServer(ep)
	//TODO: login
	loginTUI := tui.NewLogin()
	handler.LoginReq(c, loginTUI)
	//TODO: join game
	joinRoomTUI := tui.NewJoinRoom()
	handler.JoinRoomReq(c, joinRoomTUI)
	time.Sleep(3600 * time.Second)
}

func getEP() string {
	var ep string
	fmt.Println("Please enter game server endpoint")
	fmt.Scan(&ep)
	return ep
}
