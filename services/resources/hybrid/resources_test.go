// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/util"
)

func setupEnvironment() error {
	err1 := config.ParseEnvironment()
	err2 := config.AddFlags()
	err3 := addLocalConfig()

	for _, err := range []error{err1, err2, err3} {
		if err != nil {
			return err
		}
	}

	flag.Parse()
	return nil
}

func addLocalConfig() error {
	return nil
}

func TestGroupsHybrid(t *testing.T) {
	err := setupEnvironment()
	if err != nil {
		t.Fatalf("could not set up environment: %+v", err)
	}

	groupName := config.GenerateGroupName("resource-groups-hybrid")
	config.SetGroupName(groupName) // TODO: don't rely on globals

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer Cleanup(ctx)

	_, err = CreateGroup(ctx)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("resource group created")

	// Output:
	// resource group created
}
