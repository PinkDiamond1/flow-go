// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/dapperlabs/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ExecutionMetrics is an autogenerated mock type for the ExecutionMetrics type
type ExecutionMetrics struct {
	mock.Mock
}

// ChunkDataPackRequested provides a mock function with given fields:
func (_m *ExecutionMetrics) ChunkDataPackRequested() {
	_m.Called()
}

// ExecutionCollectionRequestRetried provides a mock function with given fields:
func (_m *ExecutionMetrics) ExecutionCollectionRequestRetried() {
	_m.Called()
}

// ExecutionCollectionRequestSent provides a mock function with given fields:
func (_m *ExecutionMetrics) ExecutionCollectionRequestSent() {
	_m.Called()
}

// ExecutionGasUsedPerBlock provides a mock function with given fields: gas
func (_m *ExecutionMetrics) ExecutionGasUsedPerBlock(gas uint64) {
	_m.Called(gas)
}

// ExecutionLastExecutedBlockView provides a mock function with given fields: view
func (_m *ExecutionMetrics) ExecutionLastExecutedBlockView(view uint64) {
	_m.Called(view)
}

// ExecutionStateReadsPerBlock provides a mock function with given fields: reads
func (_m *ExecutionMetrics) ExecutionStateReadsPerBlock(reads uint64) {
	_m.Called(reads)
}

// ExecutionStateStorageDiskTotal provides a mock function with given fields: bytes
func (_m *ExecutionMetrics) ExecutionStateStorageDiskTotal(bytes int64) {
	_m.Called(bytes)
}

// ExecutionStorageStateCommitment provides a mock function with given fields: bytes
func (_m *ExecutionMetrics) ExecutionStorageStateCommitment(bytes int64) {
	_m.Called(bytes)
}

// ExecutionTotalExecutedTransactions provides a mock function with given fields: numExecuted
func (_m *ExecutionMetrics) ExecutionTotalExecutedTransactions(numExecuted int) {
	_m.Called(numExecuted)
}

// FinishBlockReceivedToExecuted provides a mock function with given fields: blockID
func (_m *ExecutionMetrics) FinishBlockReceivedToExecuted(blockID flow.Identifier) {
	_m.Called(blockID)
}

// ForestApproxMemorySize provides a mock function with given fields: bytes
func (_m *ExecutionMetrics) ForestApproxMemorySize(bytes uint64) {
	_m.Called(bytes)
}

// ForestNumberOfTrees provides a mock function with given fields: number
func (_m *ExecutionMetrics) ForestNumberOfTrees(number uint64) {
	_m.Called(number)
}

// LatestTrieMaxDepth provides a mock function with given fields: number
func (_m *ExecutionMetrics) LatestTrieMaxDepth(number uint64) {
	_m.Called(number)
}

// LatestTrieMaxDepthDiff provides a mock function with given fields: number
func (_m *ExecutionMetrics) LatestTrieMaxDepthDiff(number uint64) {
	_m.Called(number)
}

// LatestTrieRegCount provides a mock function with given fields: number
func (_m *ExecutionMetrics) LatestTrieRegCount(number uint64) {
	_m.Called(number)
}

// LatestTrieRegCountDiff provides a mock function with given fields: number
func (_m *ExecutionMetrics) LatestTrieRegCountDiff(number uint64) {
	_m.Called(number)
}

// ProofSize provides a mock function with given fields: bytes
func (_m *ExecutionMetrics) ProofSize(bytes uint32) {
	_m.Called(bytes)
}

// ReadDuration provides a mock function with given fields: duration
func (_m *ExecutionMetrics) ReadDuration(duration time.Duration) {
	_m.Called(duration)
}

// ReadDurationPerItem provides a mock function with given fields: duration
func (_m *ExecutionMetrics) ReadDurationPerItem(duration time.Duration) {
	_m.Called(duration)
}

// ReadValuesNumber provides a mock function with given fields: number
func (_m *ExecutionMetrics) ReadValuesNumber(number uint64) {
	_m.Called(number)
}

// ReadValuesSize provides a mock function with given fields: byte
func (_m *ExecutionMetrics) ReadValuesSize(byte uint64) {
	_m.Called(byte)
}

// StartBlockReceivedToExecuted provides a mock function with given fields: blockID
func (_m *ExecutionMetrics) StartBlockReceivedToExecuted(blockID flow.Identifier) {
	_m.Called(blockID)
}

// TransactionChecked provides a mock function with given fields: dur
func (_m *ExecutionMetrics) TransactionChecked(dur time.Duration) {
	_m.Called(dur)
}

// TransactionInterpreted provides a mock function with given fields: dur
func (_m *ExecutionMetrics) TransactionInterpreted(dur time.Duration) {
	_m.Called(dur)
}

// TransactionParsed provides a mock function with given fields: dur
func (_m *ExecutionMetrics) TransactionParsed(dur time.Duration) {
	_m.Called(dur)
}

// UpdateCount provides a mock function with given fields:
func (_m *ExecutionMetrics) UpdateCount() {
	_m.Called()
}

// UpdateDuration provides a mock function with given fields: duration
func (_m *ExecutionMetrics) UpdateDuration(duration time.Duration) {
	_m.Called(duration)
}

// UpdateDurationPerItem provides a mock function with given fields: duration
func (_m *ExecutionMetrics) UpdateDurationPerItem(duration time.Duration) {
	_m.Called(duration)
}

// UpdateValuesNumber provides a mock function with given fields: number
func (_m *ExecutionMetrics) UpdateValuesNumber(number uint64) {
	_m.Called(number)
}

// UpdateValuesSize provides a mock function with given fields: byte
func (_m *ExecutionMetrics) UpdateValuesSize(byte uint64) {
	_m.Called(byte)
}
