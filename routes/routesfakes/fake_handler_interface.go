// Code generated by counterfeiter. DO NOT EDIT.
package routesfakes

import (
	"example-project/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

type FakeHandlerInterface struct {
	CreateUserHandlerStub        func(*gin.Context)
	createUserHandlerMutex       sync.RWMutex
	createUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	DeleteUserHandlerStub        func(*gin.Context)
	deleteUserHandlerMutex       sync.RWMutex
	deleteUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetAllUserHandlerStub        func(*gin.Context)
	getAllUserHandlerMutex       sync.RWMutex
	getAllUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetTeamMemberHandlerStub        func(*gin.Context)
	getTeamMemberHandlerMutex       sync.RWMutex
	getTeamMemberHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetUserHandlerStub        func(*gin.Context)
	getUserHandlerMutex       sync.RWMutex
	getUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	UpdateUserHandlerStub        func(*gin.Context)
	updateUserHandlerMutex       sync.RWMutex
	updateUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandlerInterface) CreateUserHandler(arg1 *gin.Context) {
	fake.createUserHandlerMutex.Lock()
	fake.createUserHandlerArgsForCall = append(fake.createUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CreateUserHandlerStub
	fake.recordInvocation("CreateUserHandler", []interface{}{arg1})
	fake.createUserHandlerMutex.Unlock()
	if stub != nil {
		fake.CreateUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CreateUserHandlerCallCount() int {
	fake.createUserHandlerMutex.RLock()
	defer fake.createUserHandlerMutex.RUnlock()
	return len(fake.createUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) CreateUserHandlerCalls(stub func(*gin.Context)) {
	fake.createUserHandlerMutex.Lock()
	defer fake.createUserHandlerMutex.Unlock()
	fake.CreateUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) CreateUserHandlerArgsForCall(i int) *gin.Context {
	fake.createUserHandlerMutex.RLock()
	defer fake.createUserHandlerMutex.RUnlock()
	argsForCall := fake.createUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) DeleteUserHandler(arg1 *gin.Context) {
	fake.deleteUserHandlerMutex.Lock()
	fake.deleteUserHandlerArgsForCall = append(fake.deleteUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.DeleteUserHandlerStub
	fake.recordInvocation("DeleteUserHandler", []interface{}{arg1})
	fake.deleteUserHandlerMutex.Unlock()
	if stub != nil {
		fake.DeleteUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) DeleteUserHandlerCallCount() int {
	fake.deleteUserHandlerMutex.RLock()
	defer fake.deleteUserHandlerMutex.RUnlock()
	return len(fake.deleteUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) DeleteUserHandlerCalls(stub func(*gin.Context)) {
	fake.deleteUserHandlerMutex.Lock()
	defer fake.deleteUserHandlerMutex.Unlock()
	fake.DeleteUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) DeleteUserHandlerArgsForCall(i int) *gin.Context {
	fake.deleteUserHandlerMutex.RLock()
	defer fake.deleteUserHandlerMutex.RUnlock()
	argsForCall := fake.deleteUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetAllUserHandler(arg1 *gin.Context) {
	fake.getAllUserHandlerMutex.Lock()
	fake.getAllUserHandlerArgsForCall = append(fake.getAllUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetAllUserHandlerStub
	fake.recordInvocation("GetAllUserHandler", []interface{}{arg1})
	fake.getAllUserHandlerMutex.Unlock()
	if stub != nil {
		fake.GetAllUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetAllUserHandlerCallCount() int {
	fake.getAllUserHandlerMutex.RLock()
	defer fake.getAllUserHandlerMutex.RUnlock()
	return len(fake.getAllUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetAllUserHandlerCalls(stub func(*gin.Context)) {
	fake.getAllUserHandlerMutex.Lock()
	defer fake.getAllUserHandlerMutex.Unlock()
	fake.GetAllUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetAllUserHandlerArgsForCall(i int) *gin.Context {
	fake.getAllUserHandlerMutex.RLock()
	defer fake.getAllUserHandlerMutex.RUnlock()
	argsForCall := fake.getAllUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetTeamMemberHandler(arg1 *gin.Context) {
	fake.getTeamMemberHandlerMutex.Lock()
	fake.getTeamMemberHandlerArgsForCall = append(fake.getTeamMemberHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetTeamMemberHandlerStub
	fake.recordInvocation("GetTeamMemberHandler", []interface{}{arg1})
	fake.getTeamMemberHandlerMutex.Unlock()
	if stub != nil {
		fake.GetTeamMemberHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetTeamMemberHandlerCallCount() int {
	fake.getTeamMemberHandlerMutex.RLock()
	defer fake.getTeamMemberHandlerMutex.RUnlock()
	return len(fake.getTeamMemberHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetTeamMemberHandlerCalls(stub func(*gin.Context)) {
	fake.getTeamMemberHandlerMutex.Lock()
	defer fake.getTeamMemberHandlerMutex.Unlock()
	fake.GetTeamMemberHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetTeamMemberHandlerArgsForCall(i int) *gin.Context {
	fake.getTeamMemberHandlerMutex.RLock()
	defer fake.getTeamMemberHandlerMutex.RUnlock()
	argsForCall := fake.getTeamMemberHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetUserHandler(arg1 *gin.Context) {
	fake.getUserHandlerMutex.Lock()
	fake.getUserHandlerArgsForCall = append(fake.getUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetUserHandlerStub
	fake.recordInvocation("GetUserHandler", []interface{}{arg1})
	fake.getUserHandlerMutex.Unlock()
	if stub != nil {
		fake.GetUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetUserHandlerCallCount() int {
	fake.getUserHandlerMutex.RLock()
	defer fake.getUserHandlerMutex.RUnlock()
	return len(fake.getUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetUserHandlerCalls(stub func(*gin.Context)) {
	fake.getUserHandlerMutex.Lock()
	defer fake.getUserHandlerMutex.Unlock()
	fake.GetUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetUserHandlerArgsForCall(i int) *gin.Context {
	fake.getUserHandlerMutex.RLock()
	defer fake.getUserHandlerMutex.RUnlock()
	argsForCall := fake.getUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) UpdateUserHandler(arg1 *gin.Context) {
	fake.updateUserHandlerMutex.Lock()
	fake.updateUserHandlerArgsForCall = append(fake.updateUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.UpdateUserHandlerStub
	fake.recordInvocation("UpdateUserHandler", []interface{}{arg1})
	fake.updateUserHandlerMutex.Unlock()
	if stub != nil {
		fake.UpdateUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) UpdateUserHandlerCallCount() int {
	fake.updateUserHandlerMutex.RLock()
	defer fake.updateUserHandlerMutex.RUnlock()
	return len(fake.updateUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) UpdateUserHandlerCalls(stub func(*gin.Context)) {
	fake.updateUserHandlerMutex.Lock()
	defer fake.updateUserHandlerMutex.Unlock()
	fake.UpdateUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) UpdateUserHandlerArgsForCall(i int) *gin.Context {
	fake.updateUserHandlerMutex.RLock()
	defer fake.updateUserHandlerMutex.RUnlock()
	argsForCall := fake.updateUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserHandlerMutex.RLock()
	defer fake.createUserHandlerMutex.RUnlock()
	fake.deleteUserHandlerMutex.RLock()
	defer fake.deleteUserHandlerMutex.RUnlock()
	fake.getAllUserHandlerMutex.RLock()
	defer fake.getAllUserHandlerMutex.RUnlock()
	fake.getTeamMemberHandlerMutex.RLock()
	defer fake.getTeamMemberHandlerMutex.RUnlock()
	fake.getUserHandlerMutex.RLock()
	defer fake.getUserHandlerMutex.RUnlock()
	fake.updateUserHandlerMutex.RLock()
	defer fake.updateUserHandlerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandlerInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routes.HandlerInterface = new(FakeHandlerInterface)
