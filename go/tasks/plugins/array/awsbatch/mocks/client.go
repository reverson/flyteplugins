// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	batch "github.com/aws/aws-sdk-go/service/batch"

	mock "github.com/stretchr/testify/mock"

	structpb "google.golang.org/protobuf/types/known/structpb"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_GetAccountID struct {
	*mock.Call
}

func (_m Client_GetAccountID) Return(_a0 string) *Client_GetAccountID {
	return &Client_GetAccountID{Call: _m.Call.Return(_a0)}
}

func (_m *Client) OnGetAccountID() *Client_GetAccountID {
	c := _m.On("GetAccountID")
	return &Client_GetAccountID{Call: c}
}

func (_m *Client) OnGetAccountIDMatch(matchers ...interface{}) *Client_GetAccountID {
	c := _m.On("GetAccountID", matchers...)
	return &Client_GetAccountID{Call: c}
}

// GetAccountID provides a mock function with given fields:
func (_m *Client) GetAccountID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Client_GetJobDetailsBatch struct {
	*mock.Call
}

func (_m Client_GetJobDetailsBatch) Return(_a0 []*batch.JobDetail, _a1 error) *Client_GetJobDetailsBatch {
	return &Client_GetJobDetailsBatch{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Client) OnGetJobDetailsBatch(ctx context.Context, ids []string) *Client_GetJobDetailsBatch {
	c := _m.On("GetJobDetailsBatch", ctx, ids)
	return &Client_GetJobDetailsBatch{Call: c}
}

func (_m *Client) OnGetJobDetailsBatchMatch(matchers ...interface{}) *Client_GetJobDetailsBatch {
	c := _m.On("GetJobDetailsBatch", matchers...)
	return &Client_GetJobDetailsBatch{Call: c}
}

// GetJobDetailsBatch provides a mock function with given fields: ctx, ids
func (_m *Client) GetJobDetailsBatch(ctx context.Context, ids []string) ([]*batch.JobDetail, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*batch.JobDetail
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*batch.JobDetail); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*batch.JobDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Client_GetRegion struct {
	*mock.Call
}

func (_m Client_GetRegion) Return(_a0 string) *Client_GetRegion {
	return &Client_GetRegion{Call: _m.Call.Return(_a0)}
}

func (_m *Client) OnGetRegion() *Client_GetRegion {
	c := _m.On("GetRegion")
	return &Client_GetRegion{Call: c}
}

func (_m *Client) OnGetRegionMatch(matchers ...interface{}) *Client_GetRegion {
	c := _m.On("GetRegion", matchers...)
	return &Client_GetRegion{Call: c}
}

// GetRegion provides a mock function with given fields:
func (_m *Client) GetRegion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Client_RegisterJobDefinition struct {
	*mock.Call
}

func (_m Client_RegisterJobDefinition) Return(arn string, err error) *Client_RegisterJobDefinition {
	return &Client_RegisterJobDefinition{Call: _m.Call.Return(arn, err)}
}

func (_m *Client) OnRegisterJobDefinition(ctx context.Context, name string, image string, role string, structObj *structpb.Struct) *Client_RegisterJobDefinition {
	c := _m.On("RegisterJobDefinition", ctx, name, image, role, structObj)
	return &Client_RegisterJobDefinition{Call: c}
}

func (_m *Client) OnRegisterJobDefinitionMatch(matchers ...interface{}) *Client_RegisterJobDefinition {
	c := _m.On("RegisterJobDefinition", matchers...)
	return &Client_RegisterJobDefinition{Call: c}
}

// RegisterJobDefinition provides a mock function with given fields: ctx, name, image, role, structObj
func (_m *Client) RegisterJobDefinition(ctx context.Context, name string, image string, role string, jobDefinitionInput *batch.RegisterJobDefinitionInput) (string, error) {
	ret := _m.Called(ctx, name, image, role, jobDefinitionInput)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *batch.RegisterJobDefinitionInput) string); ok {
		r0 = rf(ctx, name, image, role, jobDefinitionInput)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *batch.RegisterJobDefinitionInput) error); ok {
		r1 = rf(ctx, name, image, role, jobDefinitionInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Client_SubmitJob struct {
	*mock.Call
}

func (_m Client_SubmitJob) Return(jobID string, err error) *Client_SubmitJob {
	return &Client_SubmitJob{Call: _m.Call.Return(jobID, err)}
}

func (_m *Client) OnSubmitJob(ctx context.Context, input *batch.SubmitJobInput) *Client_SubmitJob {
	c := _m.On("SubmitJob", ctx, input)
	return &Client_SubmitJob{Call: c}
}

func (_m *Client) OnSubmitJobMatch(matchers ...interface{}) *Client_SubmitJob {
	c := _m.On("SubmitJob", matchers...)
	return &Client_SubmitJob{Call: c}
}

// SubmitJob provides a mock function with given fields: ctx, input
func (_m *Client) SubmitJob(ctx context.Context, input *batch.SubmitJobInput) (string, error) {
	ret := _m.Called(ctx, input)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *batch.SubmitJobInput) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *batch.SubmitJobInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Client_TerminateJob struct {
	*mock.Call
}

func (_m Client_TerminateJob) Return(_a0 error) *Client_TerminateJob {
	return &Client_TerminateJob{Call: _m.Call.Return(_a0)}
}

func (_m *Client) OnTerminateJob(ctx context.Context, jobID string, reason string) *Client_TerminateJob {
	c := _m.On("TerminateJob", ctx, jobID, reason)
	return &Client_TerminateJob{Call: c}
}

func (_m *Client) OnTerminateJobMatch(matchers ...interface{}) *Client_TerminateJob {
	c := _m.On("TerminateJob", matchers...)
	return &Client_TerminateJob{Call: c}
}

// TerminateJob provides a mock function with given fields: ctx, jobID, reason
func (_m *Client) TerminateJob(ctx context.Context, jobID string, reason string) error {
	ret := _m.Called(ctx, jobID, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, jobID, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
