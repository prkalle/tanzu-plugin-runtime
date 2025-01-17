// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestNewSetContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *SetContextInputOptions
		outputOpts *SetContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ContextOpts: &ContextOpts{
					Name:   "compatibility-one",
					Target: TargetK8s,
					GlobalOpts: &GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							"context": `name: compatibility-one
target: kubernetes
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
							"setCurrent": false,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewSetContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewGetContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *GetContextInputOptions
		outputOpts *GetContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ContextName: "compatibility-one",
			}, &GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &ContextOpts{
					Name:   "compatibility-one",
					Target: TargetK8s,
					GlobalOpts: &GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.GetContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							"contextName": "compatibility-one",
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
target: kubernetes
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewGetContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewDeleteContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *DeleteContextInputOptions
		outputOpts *DeleteContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&DeleteContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ContextName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.DeleteContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							"contextName": "compatibility-one",
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewDeleteContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewSetCurrentContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *SetCurrentContextInputOptions
		outputOpts *SetCurrentContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&SetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ContextName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetCurrentContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							"contextName": "compatibility-one",
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewSetCurrentContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewGetCurrentContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *GetCurrentContextInputOptions
		outputOpts *GetCurrentContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				Target: TargetK8s,
			}, &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &ContextOpts{
					Name:   "compatibility-one",
					Target: TargetK8s,
					GlobalOpts: &GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.GetCurrentContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Target: TargetK8s,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
target: kubernetes
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewGetCurrentContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewRemoveCurrentContextCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *RemoveCurrentContextInputOptions
		outputOpts *RemoveCurrentContextOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&RemoveCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				Target: TargetK8s,
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.RemoveCurrentContextAPIName,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Target: TargetK8s,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewRemoveCurrentContextCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}
