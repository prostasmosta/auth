package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/prostasmosta/auth/grpc/internal/repository.UserRepository -o ./mocks/user_repository_minimock.go -n UserRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/prostasmosta/auth/grpc/internal/model"
)

// UserRepositoryMock implements repository.UserRepository
type UserRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, params *model.CreateUser) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, params *model.CreateUser)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserRepositoryMockCreate

	funcGet          func(ctx context.Context, id int64) (gp1 *model.GetUser, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserRepositoryMockGet
}

// NewUserRepositoryMock returns a mock for repository.UserRepository
func NewUserRepositoryMock(t minimock.Tester) *UserRepositoryMock {
	m := &UserRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserRepositoryMockCreateParams{}

	m.GetMock = mUserRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*UserRepositoryMockGetParams{}

	return m
}

type mUserRepositoryMockCreate struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockCreateExpectation
	expectations       []*UserRepositoryMockCreateExpectation

	callArgs []*UserRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// UserRepositoryMockCreateExpectation specifies expectation struct of the UserRepository.Create
type UserRepositoryMockCreateExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockCreateParams
	results *UserRepositoryMockCreateResults
	Counter uint64
}

// UserRepositoryMockCreateParams contains parameters of the UserRepository.Create
type UserRepositoryMockCreateParams struct {
	ctx    context.Context
	params *model.CreateUser
}

// UserRepositoryMockCreateResults contains results of the UserRepository.Create
type UserRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Expect(ctx context.Context, params *model.CreateUser) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &UserRepositoryMockCreateParams{ctx, params}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Inspect(f func(ctx context.Context, params *model.CreateUser)) *mUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Return(i1 int64, err error) *UserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the UserRepository.Create method
func (mmCreate *mUserRepositoryMockCreate) Set(f func(ctx context.Context, params *model.CreateUser) (i1 int64, err error)) *UserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserRepositoryMockCreate) When(ctx context.Context, params *model.CreateUser) *UserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &UserRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UserRepositoryMockCreateParams{ctx, params},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Create return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockCreateExpectation) Then(i1 int64, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.UserRepository
func (mmCreate *UserRepositoryMock) Create(ctx context.Context, params *model.CreateUser) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, params)
	}

	mm_params := &UserRepositoryMockCreateParams{ctx, params}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := UserRepositoryMockCreateParams{ctx, params}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, params)
	}
	mmCreate.t.Fatalf("Unexpected call to UserRepositoryMock.Create. %v %v", ctx, params)
	return
}

// CreateAfterCounter returns a count of finished UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserRepositoryMockCreate) Calls() []*UserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Create")
	}
}

type mUserRepositoryMockGet struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetExpectation
	expectations       []*UserRepositoryMockGetExpectation

	callArgs []*UserRepositoryMockGetParams
	mutex    sync.RWMutex
}

// UserRepositoryMockGetExpectation specifies expectation struct of the UserRepository.Get
type UserRepositoryMockGetExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockGetParams
	results *UserRepositoryMockGetResults
	Counter uint64
}

// UserRepositoryMockGetParams contains parameters of the UserRepository.Get
type UserRepositoryMockGetParams struct {
	ctx context.Context
	id  int64
}

// UserRepositoryMockGetResults contains results of the UserRepository.Get
type UserRepositoryMockGetResults struct {
	gp1 *model.GetUser
	err error
}

// Expect sets up expected params for UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Expect(ctx context.Context, id int64) *mUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &UserRepositoryMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Inspect(f func(ctx context.Context, id int64)) *mUserRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Return(gp1 *model.GetUser, err error) *UserRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserRepositoryMockGetResults{gp1, err}
	return mmGet.mock
}

// Set uses given function f to mock the UserRepository.Get method
func (mmGet *mUserRepositoryMockGet) Set(f func(ctx context.Context, id int64) (gp1 *model.GetUser, err error)) *UserRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UserRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserRepositoryMockGet) When(ctx context.Context, id int64) *UserRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserRepositoryMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Get return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetExpectation) Then(gp1 *model.GetUser, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetResults{gp1, err}
	return e.mock
}

// Get implements repository.UserRepository
func (mmGet *UserRepositoryMock) Get(ctx context.Context, id int64) (gp1 *model.GetUser, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := &UserRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := UserRepositoryMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserRepositoryMock.Get")
		}
		return (*mm_results).gp1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserRepositoryMockGet) Calls() []*UserRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone()
}
