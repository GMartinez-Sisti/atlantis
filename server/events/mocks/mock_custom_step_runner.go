// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CustomStepRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	valid "github.com/runatlantis/atlantis/server/core/config/valid"
	command "github.com/runatlantis/atlantis/server/events/command"
	"reflect"
	"time"
	"regexp"
)

type MockCustomStepRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCustomStepRunner(options ...pegomock.Option) *MockCustomStepRunner {
	mock := &MockCustomStepRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCustomStepRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCustomStepRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCustomStepRunner) Run(ctx command.ProjectContext, shell *valid.CommandShell, cmd string, path string, envs map[string]string, streamOutput bool, postProcessOutput []valid.PostProcessRunOutputOption, postProcessFilterRegexes []*regexp.Regexp) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCustomStepRunner().")
	}
	_params := []pegomock.Param{ctx, shell, cmd, path, envs, streamOutput, postProcessOutput, postProcessFilterRegexes}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Run", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 string
	var _ret1 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(error)
		}
	}
	return _ret0, _ret1
}

func (mock *MockCustomStepRunner) VerifyWasCalledOnce() *VerifierMockCustomStepRunner {
	return &VerifierMockCustomStepRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCustomStepRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCustomStepRunner {
	return &VerifierMockCustomStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCustomStepRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCustomStepRunner {
	return &VerifierMockCustomStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCustomStepRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCustomStepRunner {
	return &VerifierMockCustomStepRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCustomStepRunner struct {
	mock                   *MockCustomStepRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCustomStepRunner) Run(ctx command.ProjectContext, shell *valid.CommandShell, cmd string, path string, envs map[string]string, streamOutput bool, postProcessOutput []valid.PostProcessRunOutputOption, postProcessFilterRegexes []*regexp.Regexp) *MockCustomStepRunner_Run_OngoingVerification {
	_params := []pegomock.Param{ctx, shell, cmd, path, envs, streamOutput, postProcessOutput, postProcessFilterRegexes}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", _params, verifier.timeout)
	return &MockCustomStepRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCustomStepRunner_Run_OngoingVerification struct {
	mock              *MockCustomStepRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCustomStepRunner_Run_OngoingVerification) GetCapturedArguments() (command.ProjectContext, *valid.CommandShell, string, string, map[string]string, bool, []valid.PostProcessRunOutputOption, []*regexp.Regexp) {
	ctx, shell, cmd, path, envs, streamOutput, postProcessOutput, postProcessFilterRegexes := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], shell[len(shell)-1], cmd[len(cmd)-1], path[len(path)-1], envs[len(envs)-1], streamOutput[len(streamOutput)-1], postProcessOutput[len(postProcessOutput)-1], postProcessFilterRegexes[len(postProcessFilterRegexes)-1]
}

func (c *MockCustomStepRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []*valid.CommandShell, _param2 []string, _param3 []string, _param4 []map[string]string, _param5 []bool, _param6 [][]valid.PostProcessRunOutputOption, _param7 [][]*regexp.Regexp) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]*valid.CommandShell, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(*valid.CommandShell)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]string, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(string)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]string, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(string)
			}
		}
		if len(_params) > 4 {
			_param4 = make([]map[string]string, len(c.methodInvocations))
			for u, param := range _params[4] {
				_param4[u] = param.(map[string]string)
			}
		}
		if len(_params) > 5 {
			_param5 = make([]bool, len(c.methodInvocations))
			for u, param := range _params[5] {
				_param5[u] = param.(bool)
			}
		}
		if len(_params) > 6 {
			_param6 = make([][]valid.PostProcessRunOutputOption, len(c.methodInvocations))
			for u, param := range _params[6] {
				_param6[u] = param.([]valid.PostProcessRunOutputOption)
			}
		}
		if len(_params) > 7 {
			_param7 = make([][]*regexp.Regexp, len(c.methodInvocations))
			for u, param := range _params[7] {
				_param7[u] = param.([]*regexp.Regexp)
			}
		}
	}
	return
}
