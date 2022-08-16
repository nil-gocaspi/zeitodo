// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"example-project/handler"
	"example-project/model"
	"sync"
)

type FakeServiceInterface struct {
	CreateUserStub        func(model.UserSignupPayload) (interface{}, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 model.UserSignupPayload
	}
	createUserReturns struct {
		result1 interface{}
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	DeleteUsersStub        func(string) (interface{}, error)
	deleteUsersMutex       sync.RWMutex
	deleteUsersArgsForCall []struct {
		arg1 string
	}
	deleteUsersReturns struct {
		result1 interface{}
		result2 error
	}
	deleteUsersReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	GetAllUserStub        func() ([]model.UserPayload, error)
	getAllUserMutex       sync.RWMutex
	getAllUserArgsForCall []struct {
	}
	getAllUserReturns struct {
		result1 []model.UserPayload
		result2 error
	}
	getAllUserReturnsOnCall map[int]struct {
		result1 []model.UserPayload
		result2 error
	}
	GetTeamMembersByNameStub        func(string) (interface{}, error)
	getTeamMembersByNameMutex       sync.RWMutex
	getTeamMembersByNameArgsForCall []struct {
		arg1 string
	}
	getTeamMembersByNameReturns struct {
		result1 interface{}
		result2 error
	}
	getTeamMembersByNameReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	GetTeamMembersByUserIDStub        func(string) (interface{}, error)
	getTeamMembersByUserIDMutex       sync.RWMutex
	getTeamMembersByUserIDArgsForCall []struct {
		arg1 string
	}
	getTeamMembersByUserIDReturns struct {
		result1 interface{}
		result2 error
	}
	getTeamMembersByUserIDReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	GetUserByIDStub        func(string) (model.UserPayload, error)
	getUserByIDMutex       sync.RWMutex
	getUserByIDArgsForCall []struct {
		arg1 string
	}
	getUserByIDReturns struct {
		result1 model.UserPayload
		result2 error
	}
	getUserByIDReturnsOnCall map[int]struct {
		result1 model.UserPayload
		result2 error
	}
	UpdateUsersStub        func([]model.User) (interface{}, error)
	updateUsersMutex       sync.RWMutex
	updateUsersArgsForCall []struct {
		arg1 []model.User
	}
	updateUsersReturns struct {
		result1 interface{}
		result2 error
	}
	updateUsersReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceInterface) CreateUser(arg1 model.UserSignupPayload) (interface{}, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 model.UserSignupPayload
	}{arg1})
	stub := fake.CreateUserStub
	fakeReturns := fake.createUserReturns
	fake.recordInvocation("CreateUser", []interface{}{arg1})
	fake.createUserMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeServiceInterface) CreateUserCalls(stub func(model.UserSignupPayload) (interface{}, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeServiceInterface) CreateUserArgsForCall(i int) model.UserSignupPayload {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) CreateUserReturns(result1 interface{}, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) CreateUserReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) DeleteUsers(arg1 string) (interface{}, error) {
	fake.deleteUsersMutex.Lock()
	ret, specificReturn := fake.deleteUsersReturnsOnCall[len(fake.deleteUsersArgsForCall)]
	fake.deleteUsersArgsForCall = append(fake.deleteUsersArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteUsersStub
	fakeReturns := fake.deleteUsersReturns
	fake.recordInvocation("DeleteUsers", []interface{}{arg1})
	fake.deleteUsersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) DeleteUsersCallCount() int {
	fake.deleteUsersMutex.RLock()
	defer fake.deleteUsersMutex.RUnlock()
	return len(fake.deleteUsersArgsForCall)
}

func (fake *FakeServiceInterface) DeleteUsersCalls(stub func(string) (interface{}, error)) {
	fake.deleteUsersMutex.Lock()
	defer fake.deleteUsersMutex.Unlock()
	fake.DeleteUsersStub = stub
}

func (fake *FakeServiceInterface) DeleteUsersArgsForCall(i int) string {
	fake.deleteUsersMutex.RLock()
	defer fake.deleteUsersMutex.RUnlock()
	argsForCall := fake.deleteUsersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) DeleteUsersReturns(result1 interface{}, result2 error) {
	fake.deleteUsersMutex.Lock()
	defer fake.deleteUsersMutex.Unlock()
	fake.DeleteUsersStub = nil
	fake.deleteUsersReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) DeleteUsersReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.deleteUsersMutex.Lock()
	defer fake.deleteUsersMutex.Unlock()
	fake.DeleteUsersStub = nil
	if fake.deleteUsersReturnsOnCall == nil {
		fake.deleteUsersReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.deleteUsersReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetAllUser() ([]model.UserPayload, error) {
	fake.getAllUserMutex.Lock()
	ret, specificReturn := fake.getAllUserReturnsOnCall[len(fake.getAllUserArgsForCall)]
	fake.getAllUserArgsForCall = append(fake.getAllUserArgsForCall, struct {
	}{})
	stub := fake.GetAllUserStub
	fakeReturns := fake.getAllUserReturns
	fake.recordInvocation("GetAllUser", []interface{}{})
	fake.getAllUserMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetAllUserCallCount() int {
	fake.getAllUserMutex.RLock()
	defer fake.getAllUserMutex.RUnlock()
	return len(fake.getAllUserArgsForCall)
}

func (fake *FakeServiceInterface) GetAllUserCalls(stub func() ([]model.UserPayload, error)) {
	fake.getAllUserMutex.Lock()
	defer fake.getAllUserMutex.Unlock()
	fake.GetAllUserStub = stub
}

func (fake *FakeServiceInterface) GetAllUserReturns(result1 []model.UserPayload, result2 error) {
	fake.getAllUserMutex.Lock()
	defer fake.getAllUserMutex.Unlock()
	fake.GetAllUserStub = nil
	fake.getAllUserReturns = struct {
		result1 []model.UserPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetAllUserReturnsOnCall(i int, result1 []model.UserPayload, result2 error) {
	fake.getAllUserMutex.Lock()
	defer fake.getAllUserMutex.Unlock()
	fake.GetAllUserStub = nil
	if fake.getAllUserReturnsOnCall == nil {
		fake.getAllUserReturnsOnCall = make(map[int]struct {
			result1 []model.UserPayload
			result2 error
		})
	}
	fake.getAllUserReturnsOnCall[i] = struct {
		result1 []model.UserPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetTeamMembersByName(arg1 string) (interface{}, error) {
	fake.getTeamMembersByNameMutex.Lock()
	ret, specificReturn := fake.getTeamMembersByNameReturnsOnCall[len(fake.getTeamMembersByNameArgsForCall)]
	fake.getTeamMembersByNameArgsForCall = append(fake.getTeamMembersByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetTeamMembersByNameStub
	fakeReturns := fake.getTeamMembersByNameReturns
	fake.recordInvocation("GetTeamMembersByName", []interface{}{arg1})
	fake.getTeamMembersByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetTeamMembersByNameCallCount() int {
	fake.getTeamMembersByNameMutex.RLock()
	defer fake.getTeamMembersByNameMutex.RUnlock()
	return len(fake.getTeamMembersByNameArgsForCall)
}

func (fake *FakeServiceInterface) GetTeamMembersByNameCalls(stub func(string) (interface{}, error)) {
	fake.getTeamMembersByNameMutex.Lock()
	defer fake.getTeamMembersByNameMutex.Unlock()
	fake.GetTeamMembersByNameStub = stub
}

func (fake *FakeServiceInterface) GetTeamMembersByNameArgsForCall(i int) string {
	fake.getTeamMembersByNameMutex.RLock()
	defer fake.getTeamMembersByNameMutex.RUnlock()
	argsForCall := fake.getTeamMembersByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) GetTeamMembersByNameReturns(result1 interface{}, result2 error) {
	fake.getTeamMembersByNameMutex.Lock()
	defer fake.getTeamMembersByNameMutex.Unlock()
	fake.GetTeamMembersByNameStub = nil
	fake.getTeamMembersByNameReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetTeamMembersByNameReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.getTeamMembersByNameMutex.Lock()
	defer fake.getTeamMembersByNameMutex.Unlock()
	fake.GetTeamMembersByNameStub = nil
	if fake.getTeamMembersByNameReturnsOnCall == nil {
		fake.getTeamMembersByNameReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.getTeamMembersByNameReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetTeamMembersByUserID(arg1 string) (interface{}, error) {
	fake.getTeamMembersByUserIDMutex.Lock()
	ret, specificReturn := fake.getTeamMembersByUserIDReturnsOnCall[len(fake.getTeamMembersByUserIDArgsForCall)]
	fake.getTeamMembersByUserIDArgsForCall = append(fake.getTeamMembersByUserIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetTeamMembersByUserIDStub
	fakeReturns := fake.getTeamMembersByUserIDReturns
	fake.recordInvocation("GetTeamMembersByUserID", []interface{}{arg1})
	fake.getTeamMembersByUserIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetTeamMembersByUserIDCallCount() int {
	fake.getTeamMembersByUserIDMutex.RLock()
	defer fake.getTeamMembersByUserIDMutex.RUnlock()
	return len(fake.getTeamMembersByUserIDArgsForCall)
}

func (fake *FakeServiceInterface) GetTeamMembersByUserIDCalls(stub func(string) (interface{}, error)) {
	fake.getTeamMembersByUserIDMutex.Lock()
	defer fake.getTeamMembersByUserIDMutex.Unlock()
	fake.GetTeamMembersByUserIDStub = stub
}

func (fake *FakeServiceInterface) GetTeamMembersByUserIDArgsForCall(i int) string {
	fake.getTeamMembersByUserIDMutex.RLock()
	defer fake.getTeamMembersByUserIDMutex.RUnlock()
	argsForCall := fake.getTeamMembersByUserIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) GetTeamMembersByUserIDReturns(result1 interface{}, result2 error) {
	fake.getTeamMembersByUserIDMutex.Lock()
	defer fake.getTeamMembersByUserIDMutex.Unlock()
	fake.GetTeamMembersByUserIDStub = nil
	fake.getTeamMembersByUserIDReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetTeamMembersByUserIDReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.getTeamMembersByUserIDMutex.Lock()
	defer fake.getTeamMembersByUserIDMutex.Unlock()
	fake.GetTeamMembersByUserIDStub = nil
	if fake.getTeamMembersByUserIDReturnsOnCall == nil {
		fake.getTeamMembersByUserIDReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.getTeamMembersByUserIDReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetUserByID(arg1 string) (model.UserPayload, error) {
	fake.getUserByIDMutex.Lock()
	ret, specificReturn := fake.getUserByIDReturnsOnCall[len(fake.getUserByIDArgsForCall)]
	fake.getUserByIDArgsForCall = append(fake.getUserByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetUserByIDStub
	fakeReturns := fake.getUserByIDReturns
	fake.recordInvocation("GetUserByID", []interface{}{arg1})
	fake.getUserByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) GetUserByIDCallCount() int {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	return len(fake.getUserByIDArgsForCall)
}

func (fake *FakeServiceInterface) GetUserByIDCalls(stub func(string) (model.UserPayload, error)) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = stub
}

func (fake *FakeServiceInterface) GetUserByIDArgsForCall(i int) string {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	argsForCall := fake.getUserByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) GetUserByIDReturns(result1 model.UserPayload, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	fake.getUserByIDReturns = struct {
		result1 model.UserPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) GetUserByIDReturnsOnCall(i int, result1 model.UserPayload, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	if fake.getUserByIDReturnsOnCall == nil {
		fake.getUserByIDReturnsOnCall = make(map[int]struct {
			result1 model.UserPayload
			result2 error
		})
	}
	fake.getUserByIDReturnsOnCall[i] = struct {
		result1 model.UserPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) UpdateUsers(arg1 []model.User) (interface{}, error) {
	var arg1Copy []model.User
	if arg1 != nil {
		arg1Copy = make([]model.User, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.updateUsersMutex.Lock()
	ret, specificReturn := fake.updateUsersReturnsOnCall[len(fake.updateUsersArgsForCall)]
	fake.updateUsersArgsForCall = append(fake.updateUsersArgsForCall, struct {
		arg1 []model.User
	}{arg1Copy})
	stub := fake.UpdateUsersStub
	fakeReturns := fake.updateUsersReturns
	fake.recordInvocation("UpdateUsers", []interface{}{arg1Copy})
	fake.updateUsersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInterface) UpdateUsersCallCount() int {
	fake.updateUsersMutex.RLock()
	defer fake.updateUsersMutex.RUnlock()
	return len(fake.updateUsersArgsForCall)
}

func (fake *FakeServiceInterface) UpdateUsersCalls(stub func([]model.User) (interface{}, error)) {
	fake.updateUsersMutex.Lock()
	defer fake.updateUsersMutex.Unlock()
	fake.UpdateUsersStub = stub
}

func (fake *FakeServiceInterface) UpdateUsersArgsForCall(i int) []model.User {
	fake.updateUsersMutex.RLock()
	defer fake.updateUsersMutex.RUnlock()
	argsForCall := fake.updateUsersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceInterface) UpdateUsersReturns(result1 interface{}, result2 error) {
	fake.updateUsersMutex.Lock()
	defer fake.updateUsersMutex.Unlock()
	fake.UpdateUsersStub = nil
	fake.updateUsersReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) UpdateUsersReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.updateUsersMutex.Lock()
	defer fake.updateUsersMutex.Unlock()
	fake.UpdateUsersStub = nil
	if fake.updateUsersReturnsOnCall == nil {
		fake.updateUsersReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.updateUsersReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.deleteUsersMutex.RLock()
	defer fake.deleteUsersMutex.RUnlock()
	fake.getAllUserMutex.RLock()
	defer fake.getAllUserMutex.RUnlock()
	fake.getTeamMembersByNameMutex.RLock()
	defer fake.getTeamMembersByNameMutex.RUnlock()
	fake.getTeamMembersByUserIDMutex.RLock()
	defer fake.getTeamMembersByUserIDMutex.RUnlock()
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	fake.updateUsersMutex.RLock()
	defer fake.updateUsersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceInterface) recordInvocation(key string, args []interface{}) {
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

var _ handler.ServiceInterface = new(FakeServiceInterface)
