/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers_test

import (
	"testing"

	ltemodels "magma/lte/cloud/go/services/lte/obsidian/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/obsidian/tests"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/serde"
	fbinternalplugin "orc8r/fbinternal/cloud/go/plugin"
	"orc8r/fbinternal/cloud/go/services/testcontroller"
	"orc8r/fbinternal/cloud/go/services/testcontroller/obsidian/handlers"
	"orc8r/fbinternal/cloud/go/services/testcontroller/obsidian/models"
	"orc8r/fbinternal/cloud/go/services/testcontroller/storage"
	"orc8r/fbinternal/cloud/go/services/testcontroller/test_init"

	"github.com/go-openapi/swag"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_ListTestCases(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e"

	oHands := handlers.GetObsidianHandlers()
	listTests := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot, obsidian.GET).HandlerFunc

	// Empty case
	tc := tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        listTests,
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler([]*models.E2eTestCase{}),
	}
	tests.RunUnitTest(t, e, tc)

	// Happy path
	err := testcontroller.CreateOrUpdateTestCase(1, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)
	err = testcontroller.CreateOrUpdateTestCase(2, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)
	err = testcontroller.CreateOrUpdateTestCase(3, testcontroller.EnodedTestExcludeTraffic, enodebdNoTrafficTestConfig())
	assert.NoError(t, err)

	tc.ExpectedResult = tests.JSONMarshaler([]*models.E2eTestCase{
		{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(1),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
			TestType: swag.String(testcontroller.EnodedTestCaseType),
		},
		{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(2),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
			TestType: swag.String(testcontroller.EnodedTestCaseType),
		},
		{
			Config: enodebdNoTrafficTestConfig(),
			Pk:     swag.Int64(3),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
			TestType: swag.String(testcontroller.EnodedTestExcludeTraffic),
		},
	})
	tests.RunUnitTest(t, e, tc)
}

func Test_ListEnodebdTestCases(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e/enodebd"

	oHands := handlers.GetObsidianHandlers()
	listTests := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot, obsidian.GET).HandlerFunc

	// Empty case
	tc := tests.Test{
		Method:         "GET",
		URL:            testURLRoot,
		Handler:        listTests,
		ExpectedStatus: 200,
		ExpectedResult: tests.JSONMarshaler([]*models.EnodebdE2eTest{}),
	}
	tests.RunUnitTest(t, e, tc)

	// Happy path
	err := testcontroller.CreateOrUpdateTestCase(1, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)
	err = testcontroller.CreateOrUpdateTestCase(2, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)
	tc.ExpectedResult = tests.JSONMarshaler([]*models.EnodebdE2eTest{
		{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(1),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
		},
		{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(2),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
		},
	})
	tests.RunUnitTest(t, e, tc)
}

func Test_CreateEnodebdTestCase(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e/enodebd"

	oHands := handlers.GetObsidianHandlers()
	createTest := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot, obsidian.POST).HandlerFunc

	tc := tests.Test{
		Method:  "POST",
		URL:     testURLRoot,
		Handler: createTest,
		Payload: &models.MutableEnodebdE2eTest{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(1),
		},
		ExpectedStatus: 201,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := testcontroller.GetTestCases(nil)
	assert.NoError(t, err)
	expected := map[int64]*testcontroller.UnmarshalledTestCase{
		1: {
			UnmarshaledConfig: defaultEnodebdTestConfig(),
			TestCase: &storage.TestCase{
				Pk:                   1,
				TestCaseType:         testcontroller.EnodedTestCaseType,
				TestConfig:           marshalTestConfig(t, defaultEnodebdTestConfig()),
				IsCurrentlyExecuting: false,
				LastExecutionTime:    timestampProto(t, 0),
				State:                "_test_controller_start_state",
				NextScheduledTime:    timestampProto(t, 0),
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func Test_GetEnodebdTestCase(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e/enodebd"

	oHands := handlers.GetObsidianHandlers()
	getTest := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot+"/:test_pk", obsidian.GET).HandlerFunc

	// Empty case
	tc := tests.Test{
		Method:         "GET",
		URL:            testURLRoot + "/1",
		ParamNames:     []string{"test_pk"},
		ParamValues:    []string{"1"},
		Handler:        getTest,
		ExpectedStatus: 404,
		ExpectedError:  "Not Found",
	}
	tests.RunUnitTest(t, e, tc)

	// Bad path param
	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot + "/abc",
		ParamNames:     []string{"test_pk"},
		ParamValues:    []string{"abc"},
		Handler:        getTest,
		ExpectedStatus: 400,
		ExpectedError:  "strconv.ParseInt: parsing \"abc\": invalid syntax",
	}
	tests.RunUnitTest(t, e, tc)

	// Happy path
	err := testcontroller.CreateOrUpdateTestCase(1, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)
	tc = tests.Test{
		Method:         "GET",
		URL:            testURLRoot + "/1",
		ParamNames:     []string{"test_pk"},
		ParamValues:    []string{"1"},
		Handler:        getTest,
		ExpectedStatus: 200,
		ExpectedResult: &models.EnodebdE2eTest{
			Config: defaultEnodebdTestConfig(),
			Pk:     swag.Int64(1),
			State: &models.E2eTestCaseState{
				CurrentState:      "_test_controller_start_state",
				IsExecuting:       swag.Bool(false),
				LastExecutionTime: expectedDT(t, 0),
				NextScheduledTime: expectedDT(t, 0),
			},
		},
	}
	tests.RunUnitTest(t, e, tc)
}

func Test_UpdateEnodebdTestCase(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e/enodebd"

	oHands := handlers.GetObsidianHandlers()
	updateTest := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot+"/:test_pk", obsidian.PUT).HandlerFunc

	err := testcontroller.CreateOrUpdateTestCase(1, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)

	newCfg := defaultEnodebdTestConfig()
	newCfg.NetworkID = swag.String("network2")
	newCfg.AgwConfig.TargetGatewayID = swag.String("gw2")
	tc := tests.Test{
		Method:         "PUT",
		URL:            testURLRoot + "/1",
		ParamNames:     []string{"test_pk"},
		ParamValues:    []string{"1"},
		Handler:        updateTest,
		Payload:        newCfg,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)

	actual, err := testcontroller.GetTestCases(nil)
	assert.NoError(t, err)
	expected := map[int64]*testcontroller.UnmarshalledTestCase{
		1: {
			UnmarshaledConfig: newCfg,
			TestCase: &storage.TestCase{
				Pk:                   1,
				TestCaseType:         testcontroller.EnodedTestCaseType,
				TestConfig:           marshalTestConfig(t, newCfg),
				IsCurrentlyExecuting: false,
				LastExecutionTime:    timestampProto(t, 0),
				State:                "_test_controller_start_state",
				NextScheduledTime:    timestampProto(t, 0),
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func Test_DeleteEnodebdTestCase(t *testing.T) {
	_ = plugin.RegisterPluginForTests(t, &fbinternalplugin.FbinternalOrchestratorPlugin{})
	test_init.StartTestService(t)

	e := echo.New()
	testURLRoot := "/magma/v1/tests/e2e/enodebd"

	oHands := handlers.GetObsidianHandlers()
	deleteTest := tests.GetHandlerByPathAndMethod(t, oHands, testURLRoot+"/:test_pk", obsidian.DELETE).HandlerFunc

	err := testcontroller.CreateOrUpdateTestCase(1, testcontroller.EnodedTestCaseType, defaultEnodebdTestConfig())
	assert.NoError(t, err)

	tc := tests.Test{
		Method:         "DELETE",
		URL:            testURLRoot + "/1",
		ParamNames:     []string{"test_pk"},
		ParamValues:    []string{"1"},
		Handler:        deleteTest,
		ExpectedStatus: 204,
	}
	tests.RunUnitTest(t, e, tc)
	actual, err := testcontroller.GetTestCases(nil)
	assert.NoError(t, err)
	assert.Empty(t, actual)
}

func defaultEnodebdTestConfig() *models.EnodebdTestConfig {
	return &models.EnodebdTestConfig{
		AgwConfig: &models.AgwTestConfig{
			PackageRepo:     swag.String("facebook.com"),
			ReleaseChannel:  swag.String("beta"),
			TargetGatewayID: swag.String("gw1"),
			TargetTier:      swag.String("default"),
			SLACKWebhook:    swag.String("foo.com"),
		},
		EnodebSN:        swag.String("1202000038269KP0037"),
		RunTrafficTests: swag.Bool(true),
		NetworkID:       swag.String("network1"),
		TrafficGwID:     swag.String("gw2"),
		EnodebConfig: &ltemodels.EnodebConfiguration{
			BandwidthMhz:           20,
			CellID:                 swag.Uint32(138777000),
			DeviceClass:            "Baicells ID TDD/FDD",
			Earfcndl:               44590,
			Pci:                    260,
			SpecialSubframePattern: 7,
			SubframeAssignment:     2,
			Tac:                    1,
			TransmitEnabled:        swag.Bool(true),
		},
	}
}

func enodebdNoTrafficTestConfig() *models.EnodebdTestConfig {
	return &models.EnodebdTestConfig{
		AgwConfig: &models.AgwTestConfig{
			PackageRepo:     swag.String("facebook.com"),
			ReleaseChannel:  swag.String("beta"),
			TargetGatewayID: swag.String("gw1"),
			TargetTier:      swag.String("default"),
			SLACKWebhook:    swag.String("foo.com"),
		},
		EnodebSN:        swag.String("1202000038269KP0037"),
		RunTrafficTests: swag.Bool(false),
		NetworkID:       swag.String("network1"),
		TrafficGwID:     swag.String("gw2"),
		EnodebConfig: &ltemodels.EnodebConfiguration{
			BandwidthMhz:           20,
			CellID:                 swag.Uint32(138777000),
			DeviceClass:            "Baicells ID TDD/FDD",
			Earfcndl:               44590,
			Pci:                    260,
			SpecialSubframePattern: 7,
			SubframeAssignment:     2,
			Tac:                    1,
			TransmitEnabled:        swag.Bool(true),
		},
	}
}

func marshalTestConfig(t *testing.T, tc *models.EnodebdTestConfig) []byte {
	ret, err := serde.Serialize(testcontroller.SerdeDomain, testcontroller.EnodedTestCaseType, tc)
	assert.NoError(t, err)

	ret, err = serde.Serialize(testcontroller.SerdeDomain, testcontroller.EnodedTestExcludeTraffic, tc)
	assert.NoError(t, err)
	return ret
}