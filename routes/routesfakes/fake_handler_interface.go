// Code generated by counterfeiter. DO NOT EDIT.
package routesfakes

import (
	"example-project/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

type FakeHandlerInterface struct {
	CalcultimeEntryStub        func(*gin.Context)
	calcultimeEntryMutex       sync.RWMutex
	calcultimeEntryArgsForCall []struct {
		arg1 *gin.Context
	}
	CreatTimeEntryStub        func(*gin.Context)
	creatTimeEntryMutex       sync.RWMutex
	creatTimeEntryArgsForCall []struct {
		arg1 *gin.Context
	}
	CreateProposalsHandlerStub        func(*gin.Context)
	createProposalsHandlerMutex       sync.RWMutex
	createProposalsHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	CreateUserHandlerStub        func(*gin.Context)
	createUserHandlerMutex       sync.RWMutex
	createUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	DeleteProposalHandlerStub        func(*gin.Context)
	deleteProposalHandlerMutex       sync.RWMutex
	deleteProposalHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	DeleteTimeEntryStub        func(*gin.Context)
	deleteTimeEntryMutex       sync.RWMutex
	deleteTimeEntryArgsForCall []struct {
		arg1 *gin.Context
	}
	DeleteUserHandlerStub        func(*gin.Context)
	deleteUserHandlerMutex       sync.RWMutex
	deleteUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetAllTimeEntryStub        func(*gin.Context)
	getAllTimeEntryMutex       sync.RWMutex
	getAllTimeEntryArgsForCall []struct {
		arg1 *gin.Context
	}
	GetAllUserHandlerStub        func(*gin.Context)
	getAllUserHandlerMutex       sync.RWMutex
	getAllUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetProposalsByIdStub        func(*gin.Context)
	getProposalsByIdMutex       sync.RWMutex
	getProposalsByIdArgsForCall []struct {
		arg1 *gin.Context
	}
	GetTeamMemberHandlerStub        func(*gin.Context)
	getTeamMemberHandlerMutex       sync.RWMutex
	getTeamMemberHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetTimeEntryStub        func(*gin.Context)
	getTimeEntryMutex       sync.RWMutex
	getTimeEntryArgsForCall []struct {
		arg1 *gin.Context
	}
	GetUserByTokenStub        func(*gin.Context)
	getUserByTokenMutex       sync.RWMutex
	getUserByTokenArgsForCall []struct {
		arg1 *gin.Context
	}
	GetUserHandlerStub        func(*gin.Context)
	getUserHandlerMutex       sync.RWMutex
	getUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	LoginUserHandlerStub        func(*gin.Context)
	loginUserHandlerMutex       sync.RWMutex
	loginUserHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	PermissionMiddlewareStub        func(*gin.Context)
	permissionMiddlewareMutex       sync.RWMutex
	permissionMiddlewareArgsForCall []struct {
		arg1 *gin.Context
	}
	RefreshTokenHandlerStub        func(*gin.Context)
	refreshTokenHandlerMutex       sync.RWMutex
	refreshTokenHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	UpdateProposalsHandlerStub        func(*gin.Context)
	updateProposalsHandlerMutex       sync.RWMutex
	updateProposalsHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	UpdateTimeEntryStub        func(*gin.Context)
	updateTimeEntryMutex       sync.RWMutex
	updateTimeEntryArgsForCall []struct {
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

func (fake *FakeHandlerInterface) CalcultimeEntry(arg1 *gin.Context) {
	fake.calcultimeEntryMutex.Lock()
	fake.calcultimeEntryArgsForCall = append(fake.calcultimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CalcultimeEntryStub
	fake.recordInvocation("CalcultimeEntry", []interface{}{arg1})
	fake.calcultimeEntryMutex.Unlock()
	if stub != nil {
		fake.CalcultimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CalcultimeEntryCallCount() int {
	fake.calcultimeEntryMutex.RLock()
	defer fake.calcultimeEntryMutex.RUnlock()
	return len(fake.calcultimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) CalcultimeEntryCalls(stub func(*gin.Context)) {
	fake.calcultimeEntryMutex.Lock()
	defer fake.calcultimeEntryMutex.Unlock()
	fake.CalcultimeEntryStub = stub
}

func (fake *FakeHandlerInterface) CalcultimeEntryArgsForCall(i int) *gin.Context {
	fake.calcultimeEntryMutex.RLock()
	defer fake.calcultimeEntryMutex.RUnlock()
	argsForCall := fake.calcultimeEntryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) CreatTimeEntry(arg1 *gin.Context) {
	fake.creatTimeEntryMutex.Lock()
	fake.creatTimeEntryArgsForCall = append(fake.creatTimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CreatTimeEntryStub
	fake.recordInvocation("CreatTimeEntry", []interface{}{arg1})
	fake.creatTimeEntryMutex.Unlock()
	if stub != nil {
		fake.CreatTimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CreatTimeEntryCallCount() int {
	fake.creatTimeEntryMutex.RLock()
	defer fake.creatTimeEntryMutex.RUnlock()
	return len(fake.creatTimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) CreatTimeEntryCalls(stub func(*gin.Context)) {
	fake.creatTimeEntryMutex.Lock()
	defer fake.creatTimeEntryMutex.Unlock()
	fake.CreatTimeEntryStub = stub
}

func (fake *FakeHandlerInterface) CreatTimeEntryArgsForCall(i int) *gin.Context {
	fake.creatTimeEntryMutex.RLock()
	defer fake.creatTimeEntryMutex.RUnlock()
	argsForCall := fake.creatTimeEntryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) CreateProposalsHandler(arg1 *gin.Context) {
	fake.createProposalsHandlerMutex.Lock()
	fake.createProposalsHandlerArgsForCall = append(fake.createProposalsHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CreateProposalsHandlerStub
	fake.recordInvocation("CreateProposalsHandler", []interface{}{arg1})
	fake.createProposalsHandlerMutex.Unlock()
	if stub != nil {
		fake.CreateProposalsHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CreateProposalsHandlerCallCount() int {
	fake.createProposalsHandlerMutex.RLock()
	defer fake.createProposalsHandlerMutex.RUnlock()
	return len(fake.createProposalsHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) CreateProposalsHandlerCalls(stub func(*gin.Context)) {
	fake.createProposalsHandlerMutex.Lock()
	defer fake.createProposalsHandlerMutex.Unlock()
	fake.CreateProposalsHandlerStub = stub
}

func (fake *FakeHandlerInterface) CreateProposalsHandlerArgsForCall(i int) *gin.Context {
	fake.createProposalsHandlerMutex.RLock()
	defer fake.createProposalsHandlerMutex.RUnlock()
	argsForCall := fake.createProposalsHandlerArgsForCall[i]
	return argsForCall.arg1
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

func (fake *FakeHandlerInterface) DeleteProposalHandler(arg1 *gin.Context) {
	fake.deleteProposalHandlerMutex.Lock()
	fake.deleteProposalHandlerArgsForCall = append(fake.deleteProposalHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.DeleteProposalHandlerStub
	fake.recordInvocation("DeleteProposalHandler", []interface{}{arg1})
	fake.deleteProposalHandlerMutex.Unlock()
	if stub != nil {
		fake.DeleteProposalHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) DeleteProposalHandlerCallCount() int {
	fake.deleteProposalHandlerMutex.RLock()
	defer fake.deleteProposalHandlerMutex.RUnlock()
	return len(fake.deleteProposalHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) DeleteProposalHandlerCalls(stub func(*gin.Context)) {
	fake.deleteProposalHandlerMutex.Lock()
	defer fake.deleteProposalHandlerMutex.Unlock()
	fake.DeleteProposalHandlerStub = stub
}

func (fake *FakeHandlerInterface) DeleteProposalHandlerArgsForCall(i int) *gin.Context {
	fake.deleteProposalHandlerMutex.RLock()
	defer fake.deleteProposalHandlerMutex.RUnlock()
	argsForCall := fake.deleteProposalHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) DeleteTimeEntry(arg1 *gin.Context) {
	fake.deleteTimeEntryMutex.Lock()
	fake.deleteTimeEntryArgsForCall = append(fake.deleteTimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.DeleteTimeEntryStub
	fake.recordInvocation("DeleteTimeEntry", []interface{}{arg1})
	fake.deleteTimeEntryMutex.Unlock()
	if stub != nil {
		fake.DeleteTimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) DeleteTimeEntryCallCount() int {
	fake.deleteTimeEntryMutex.RLock()
	defer fake.deleteTimeEntryMutex.RUnlock()
	return len(fake.deleteTimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) DeleteTimeEntryCalls(stub func(*gin.Context)) {
	fake.deleteTimeEntryMutex.Lock()
	defer fake.deleteTimeEntryMutex.Unlock()
	fake.DeleteTimeEntryStub = stub
}

func (fake *FakeHandlerInterface) DeleteTimeEntryArgsForCall(i int) *gin.Context {
	fake.deleteTimeEntryMutex.RLock()
	defer fake.deleteTimeEntryMutex.RUnlock()
	argsForCall := fake.deleteTimeEntryArgsForCall[i]
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

func (fake *FakeHandlerInterface) GetAllTimeEntry(arg1 *gin.Context) {
	fake.getAllTimeEntryMutex.Lock()
	fake.getAllTimeEntryArgsForCall = append(fake.getAllTimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetAllTimeEntryStub
	fake.recordInvocation("GetAllTimeEntry", []interface{}{arg1})
	fake.getAllTimeEntryMutex.Unlock()
	if stub != nil {
		fake.GetAllTimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetAllTimeEntryCallCount() int {
	fake.getAllTimeEntryMutex.RLock()
	defer fake.getAllTimeEntryMutex.RUnlock()
	return len(fake.getAllTimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) GetAllTimeEntryCalls(stub func(*gin.Context)) {
	fake.getAllTimeEntryMutex.Lock()
	defer fake.getAllTimeEntryMutex.Unlock()
	fake.GetAllTimeEntryStub = stub
}

func (fake *FakeHandlerInterface) GetAllTimeEntryArgsForCall(i int) *gin.Context {
	fake.getAllTimeEntryMutex.RLock()
	defer fake.getAllTimeEntryMutex.RUnlock()
	argsForCall := fake.getAllTimeEntryArgsForCall[i]
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

func (fake *FakeHandlerInterface) GetProposalsById(arg1 *gin.Context) {
	fake.getProposalsByIdMutex.Lock()
	fake.getProposalsByIdArgsForCall = append(fake.getProposalsByIdArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetProposalsByIdStub
	fake.recordInvocation("GetProposalsById", []interface{}{arg1})
	fake.getProposalsByIdMutex.Unlock()
	if stub != nil {
		fake.GetProposalsByIdStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetProposalsByIdCallCount() int {
	fake.getProposalsByIdMutex.RLock()
	defer fake.getProposalsByIdMutex.RUnlock()
	return len(fake.getProposalsByIdArgsForCall)
}

func (fake *FakeHandlerInterface) GetProposalsByIdCalls(stub func(*gin.Context)) {
	fake.getProposalsByIdMutex.Lock()
	defer fake.getProposalsByIdMutex.Unlock()
	fake.GetProposalsByIdStub = stub
}

func (fake *FakeHandlerInterface) GetProposalsByIdArgsForCall(i int) *gin.Context {
	fake.getProposalsByIdMutex.RLock()
	defer fake.getProposalsByIdMutex.RUnlock()
	argsForCall := fake.getProposalsByIdArgsForCall[i]
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

func (fake *FakeHandlerInterface) GetTimeEntry(arg1 *gin.Context) {
	fake.getTimeEntryMutex.Lock()
	fake.getTimeEntryArgsForCall = append(fake.getTimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetTimeEntryStub
	fake.recordInvocation("GetTimeEntry", []interface{}{arg1})
	fake.getTimeEntryMutex.Unlock()
	if stub != nil {
		fake.GetTimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetTimeEntryCallCount() int {
	fake.getTimeEntryMutex.RLock()
	defer fake.getTimeEntryMutex.RUnlock()
	return len(fake.getTimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) GetTimeEntryCalls(stub func(*gin.Context)) {
	fake.getTimeEntryMutex.Lock()
	defer fake.getTimeEntryMutex.Unlock()
	fake.GetTimeEntryStub = stub
}

func (fake *FakeHandlerInterface) GetTimeEntryArgsForCall(i int) *gin.Context {
	fake.getTimeEntryMutex.RLock()
	defer fake.getTimeEntryMutex.RUnlock()
	argsForCall := fake.getTimeEntryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetUserByToken(arg1 *gin.Context) {
	fake.getUserByTokenMutex.Lock()
	fake.getUserByTokenArgsForCall = append(fake.getUserByTokenArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetUserByTokenStub
	fake.recordInvocation("GetUserByToken", []interface{}{arg1})
	fake.getUserByTokenMutex.Unlock()
	if stub != nil {
		fake.GetUserByTokenStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetUserByTokenCallCount() int {
	fake.getUserByTokenMutex.RLock()
	defer fake.getUserByTokenMutex.RUnlock()
	return len(fake.getUserByTokenArgsForCall)
}

func (fake *FakeHandlerInterface) GetUserByTokenCalls(stub func(*gin.Context)) {
	fake.getUserByTokenMutex.Lock()
	defer fake.getUserByTokenMutex.Unlock()
	fake.GetUserByTokenStub = stub
}

func (fake *FakeHandlerInterface) GetUserByTokenArgsForCall(i int) *gin.Context {
	fake.getUserByTokenMutex.RLock()
	defer fake.getUserByTokenMutex.RUnlock()
	argsForCall := fake.getUserByTokenArgsForCall[i]
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

func (fake *FakeHandlerInterface) LoginUserHandler(arg1 *gin.Context) {
	fake.loginUserHandlerMutex.Lock()
	fake.loginUserHandlerArgsForCall = append(fake.loginUserHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.LoginUserHandlerStub
	fake.recordInvocation("LoginUserHandler", []interface{}{arg1})
	fake.loginUserHandlerMutex.Unlock()
	if stub != nil {
		fake.LoginUserHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) LoginUserHandlerCallCount() int {
	fake.loginUserHandlerMutex.RLock()
	defer fake.loginUserHandlerMutex.RUnlock()
	return len(fake.loginUserHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) LoginUserHandlerCalls(stub func(*gin.Context)) {
	fake.loginUserHandlerMutex.Lock()
	defer fake.loginUserHandlerMutex.Unlock()
	fake.LoginUserHandlerStub = stub
}

func (fake *FakeHandlerInterface) LoginUserHandlerArgsForCall(i int) *gin.Context {
	fake.loginUserHandlerMutex.RLock()
	defer fake.loginUserHandlerMutex.RUnlock()
	argsForCall := fake.loginUserHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) PermissionMiddleware(arg1 *gin.Context) {
	fake.permissionMiddlewareMutex.Lock()
	fake.permissionMiddlewareArgsForCall = append(fake.permissionMiddlewareArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.PermissionMiddlewareStub
	fake.recordInvocation("PermissionMiddleware", []interface{}{arg1})
	fake.permissionMiddlewareMutex.Unlock()
	if stub != nil {
		fake.PermissionMiddlewareStub(arg1)
	}
}

func (fake *FakeHandlerInterface) PermissionMiddlewareCallCount() int {
	fake.permissionMiddlewareMutex.RLock()
	defer fake.permissionMiddlewareMutex.RUnlock()
	return len(fake.permissionMiddlewareArgsForCall)
}

func (fake *FakeHandlerInterface) PermissionMiddlewareCalls(stub func(*gin.Context)) {
	fake.permissionMiddlewareMutex.Lock()
	defer fake.permissionMiddlewareMutex.Unlock()
	fake.PermissionMiddlewareStub = stub
}

func (fake *FakeHandlerInterface) PermissionMiddlewareArgsForCall(i int) *gin.Context {
	fake.permissionMiddlewareMutex.RLock()
	defer fake.permissionMiddlewareMutex.RUnlock()
	argsForCall := fake.permissionMiddlewareArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) RefreshTokenHandler(arg1 *gin.Context) {
	fake.refreshTokenHandlerMutex.Lock()
	fake.refreshTokenHandlerArgsForCall = append(fake.refreshTokenHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.RefreshTokenHandlerStub
	fake.recordInvocation("RefreshTokenHandler", []interface{}{arg1})
	fake.refreshTokenHandlerMutex.Unlock()
	if stub != nil {
		fake.RefreshTokenHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) RefreshTokenHandlerCallCount() int {
	fake.refreshTokenHandlerMutex.RLock()
	defer fake.refreshTokenHandlerMutex.RUnlock()
	return len(fake.refreshTokenHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) RefreshTokenHandlerCalls(stub func(*gin.Context)) {
	fake.refreshTokenHandlerMutex.Lock()
	defer fake.refreshTokenHandlerMutex.Unlock()
	fake.RefreshTokenHandlerStub = stub
}

func (fake *FakeHandlerInterface) RefreshTokenHandlerArgsForCall(i int) *gin.Context {
	fake.refreshTokenHandlerMutex.RLock()
	defer fake.refreshTokenHandlerMutex.RUnlock()
	argsForCall := fake.refreshTokenHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) UpdateProposalsHandler(arg1 *gin.Context) {
	fake.updateProposalsHandlerMutex.Lock()
	fake.updateProposalsHandlerArgsForCall = append(fake.updateProposalsHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.UpdateProposalsHandlerStub
	fake.recordInvocation("UpdateProposalsHandler", []interface{}{arg1})
	fake.updateProposalsHandlerMutex.Unlock()
	if stub != nil {
		fake.UpdateProposalsHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) UpdateProposalsHandlerCallCount() int {
	fake.updateProposalsHandlerMutex.RLock()
	defer fake.updateProposalsHandlerMutex.RUnlock()
	return len(fake.updateProposalsHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) UpdateProposalsHandlerCalls(stub func(*gin.Context)) {
	fake.updateProposalsHandlerMutex.Lock()
	defer fake.updateProposalsHandlerMutex.Unlock()
	fake.UpdateProposalsHandlerStub = stub
}

func (fake *FakeHandlerInterface) UpdateProposalsHandlerArgsForCall(i int) *gin.Context {
	fake.updateProposalsHandlerMutex.RLock()
	defer fake.updateProposalsHandlerMutex.RUnlock()
	argsForCall := fake.updateProposalsHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) UpdateTimeEntry(arg1 *gin.Context) {
	fake.updateTimeEntryMutex.Lock()
	fake.updateTimeEntryArgsForCall = append(fake.updateTimeEntryArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.UpdateTimeEntryStub
	fake.recordInvocation("UpdateTimeEntry", []interface{}{arg1})
	fake.updateTimeEntryMutex.Unlock()
	if stub != nil {
		fake.UpdateTimeEntryStub(arg1)
	}
}

func (fake *FakeHandlerInterface) UpdateTimeEntryCallCount() int {
	fake.updateTimeEntryMutex.RLock()
	defer fake.updateTimeEntryMutex.RUnlock()
	return len(fake.updateTimeEntryArgsForCall)
}

func (fake *FakeHandlerInterface) UpdateTimeEntryCalls(stub func(*gin.Context)) {
	fake.updateTimeEntryMutex.Lock()
	defer fake.updateTimeEntryMutex.Unlock()
	fake.UpdateTimeEntryStub = stub
}

func (fake *FakeHandlerInterface) UpdateTimeEntryArgsForCall(i int) *gin.Context {
	fake.updateTimeEntryMutex.RLock()
	defer fake.updateTimeEntryMutex.RUnlock()
	argsForCall := fake.updateTimeEntryArgsForCall[i]
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
	fake.calcultimeEntryMutex.RLock()
	defer fake.calcultimeEntryMutex.RUnlock()
	fake.creatTimeEntryMutex.RLock()
	defer fake.creatTimeEntryMutex.RUnlock()
	fake.createProposalsHandlerMutex.RLock()
	defer fake.createProposalsHandlerMutex.RUnlock()
	fake.createUserHandlerMutex.RLock()
	defer fake.createUserHandlerMutex.RUnlock()
	fake.deleteProposalHandlerMutex.RLock()
	defer fake.deleteProposalHandlerMutex.RUnlock()
	fake.deleteTimeEntryMutex.RLock()
	defer fake.deleteTimeEntryMutex.RUnlock()
	fake.deleteUserHandlerMutex.RLock()
	defer fake.deleteUserHandlerMutex.RUnlock()
	fake.getAllTimeEntryMutex.RLock()
	defer fake.getAllTimeEntryMutex.RUnlock()
	fake.getAllUserHandlerMutex.RLock()
	defer fake.getAllUserHandlerMutex.RUnlock()
	fake.getProposalsByIdMutex.RLock()
	defer fake.getProposalsByIdMutex.RUnlock()
	fake.getTeamMemberHandlerMutex.RLock()
	defer fake.getTeamMemberHandlerMutex.RUnlock()
	fake.getTimeEntryMutex.RLock()
	defer fake.getTimeEntryMutex.RUnlock()
	fake.getUserByTokenMutex.RLock()
	defer fake.getUserByTokenMutex.RUnlock()
	fake.getUserHandlerMutex.RLock()
	defer fake.getUserHandlerMutex.RUnlock()
	fake.loginUserHandlerMutex.RLock()
	defer fake.loginUserHandlerMutex.RUnlock()
	fake.permissionMiddlewareMutex.RLock()
	defer fake.permissionMiddlewareMutex.RUnlock()
	fake.refreshTokenHandlerMutex.RLock()
	defer fake.refreshTokenHandlerMutex.RUnlock()
	fake.updateProposalsHandlerMutex.RLock()
	defer fake.updateProposalsHandlerMutex.RUnlock()
	fake.updateTimeEntryMutex.RLock()
	defer fake.updateTimeEntryMutex.RUnlock()
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
