// Code generated by MockGen. DO NOT EDIT.
// Source: stub/git.woa.com/crotaliu/pb-hub/movie.trpc.go

// Package pb_hub is a generated GoMock package.
package pb_hub

import (
	context "context"
	client "git.code.oa.com/trpc-go/trpc-go/client"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMovieService is a mock of MovieService interface
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// GetMovieList mocks base method
func (m *MockMovieService) GetMovieList(ctx context.Context, req *GetMovieListReq, rsp *GetMovieListRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieList", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMovieList indicates an expected call of GetMovieList
func (mr *MockMovieServiceMockRecorder) GetMovieList(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieList", reflect.TypeOf((*MockMovieService)(nil).GetMovieList), ctx, req, rsp)
}

// MockMovieClientProxy is a mock of MovieClientProxy interface
type MockMovieClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockMovieClientProxyMockRecorder
}

// MockMovieClientProxyMockRecorder is the mock recorder for MockMovieClientProxy
type MockMovieClientProxyMockRecorder struct {
	mock *MockMovieClientProxy
}

// NewMockMovieClientProxy creates a new mock instance
func NewMockMovieClientProxy(ctrl *gomock.Controller) *MockMovieClientProxy {
	mock := &MockMovieClientProxy{ctrl: ctrl}
	mock.recorder = &MockMovieClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieClientProxy) EXPECT() *MockMovieClientProxyMockRecorder {
	return m.recorder
}

// GetMovieList mocks base method
func (m *MockMovieClientProxy) GetMovieList(ctx context.Context, req *GetMovieListReq, opts ...client.Option) (*GetMovieListRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovieList", varargs...)
	ret0, _ := ret[0].(*GetMovieListRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieList indicates an expected call of GetMovieList
func (mr *MockMovieClientProxyMockRecorder) GetMovieList(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieList", reflect.TypeOf((*MockMovieClientProxy)(nil).GetMovieList), varargs...)
}