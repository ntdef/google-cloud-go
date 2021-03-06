// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package iot_test

import (
	"context"

	iot "cloud.google.com/go/iot/apiv1"
	"google.golang.org/api/iterator"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewDeviceManagerClient() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleDeviceManagerClient_CreateDeviceRegistry() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.CreateDeviceRegistryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetDeviceRegistry() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.GetDeviceRegistryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UpdateDeviceRegistry() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.UpdateDeviceRegistryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_DeleteDeviceRegistry() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.DeleteDeviceRegistryRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDeviceManagerClient_ListDeviceRegistries() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.ListDeviceRegistriesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDeviceRegistries(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDeviceManagerClient_CreateDevice() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.CreateDeviceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetDevice() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.GetDeviceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UpdateDevice() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.UpdateDeviceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_DeleteDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.DeleteDeviceRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDeviceManagerClient_ListDevices() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.ListDevicesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDevices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDeviceManagerClient_ModifyCloudToDeviceConfig() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.ModifyCloudToDeviceConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ModifyCloudToDeviceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_ListDeviceConfigVersions() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.ListDeviceConfigVersionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListDeviceConfigVersions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_ListDeviceStates() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.ListDeviceStatesRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListDeviceStates(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_SetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_TestIamPermissions() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_SendCommandToDevice() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.SendCommandToDeviceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SendCommandToDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_BindDeviceToGateway() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.BindDeviceToGatewayRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BindDeviceToGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UnbindDeviceFromGateway() {
	// import iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"

	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iotpb.UnbindDeviceFromGatewayRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UnbindDeviceFromGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
