// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package compute

import (
	"context"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/util"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/network"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

// Example_createVMSS creates a group and network artifacts needed for a VMSS, then
// creates a VMSS and tests operations on it.
func Example_createVMSS() {
	var groupName = config.GenerateGroupName("VMSS")
	// TODO: remove and use local `groupName` only
	config.SetGroupName(groupName)

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resources.Cleanup(ctx)

	_, err := resources.CreateGroup(ctx, groupName)
	if err != nil {
		util.LogAndPanic(err)
	}

	_, err = network.CreateVirtualNetworkAndSubnets(ctx, virtualNetworkName, subnet1Name, subnet2Name)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created vnet and 2 subnets")

	_, err = CreateVMSS(ctx, vmssName, virtualNetworkName, subnet1Name, username, password, sshPublicKeyPath)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created VMSS")

	//set or change VMSS metadata
	_, err = UpdateVMSS(ctx, vmssName, map[string]*string{
		"runtime": to.StringPtr("go"),
		"cloud":   to.StringPtr("azure"),
	})
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("updated VMSS")

	// set or change system state
	_, err = StartVMSS(ctx, vmssName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("started VMSS")

	_, err = RestartVMSS(ctx, vmssName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("restarted VMSS")

	_, err = StopVMSS(ctx, vmssName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("stopped VMSS")

	// Output:
	// created vnet and 2 subnets
	// created VMSS
	// updated VMSS
	// started VMSS
	// restarted VMSS
	// stopped VMSS
}
