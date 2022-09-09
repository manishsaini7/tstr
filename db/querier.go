// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Querier interface {
	AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error
	AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (Run, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	AuthAccessToken(ctx context.Context, db DBTX, tokenHash string) (AuthAccessTokenRow, error)
	DeleteRunsForTest(ctx context.Context, db DBTX, testID uuid.UUID) error
	DeleteTest(ctx context.Context, db DBTX, id uuid.UUID) error
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	GetAccessToken(ctx context.Context, db DBTX, id uuid.UUID) (GetAccessTokenRow, error)
	GetRun(ctx context.Context, db DBTX, id uuid.UUID) (Run, error)
	GetRunner(ctx context.Context, db DBTX, id uuid.UUID) (Runner, error)
	GetTest(ctx context.Context, db DBTX, id uuid.UUID) (Test, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	IssueAccessToken(ctx context.Context, db DBTX, arg IssueAccessTokenParams) (IssueAccessTokenRow, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	ListAccessTokens(ctx context.Context, db DBTX, arg ListAccessTokensParams) ([]ListAccessTokensRow, error)
	ListPendingRuns(ctx context.Context, db DBTX) ([]Run, error)
	ListRunners(ctx context.Context, db DBTX) ([]Runner, error)
	ListRuns(ctx context.Context, db DBTX) ([]Run, error)
	ListTests(ctx context.Context, db DBTX) ([]Test, error)
	ListTestsToSchedule(ctx context.Context, db DBTX) ([]Test, error)
	QueryRunners(ctx context.Context, db DBTX, arg QueryRunnersParams) ([]Runner, error)
	QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]Run, error)
	QueryTests(ctx context.Context, db DBTX, arg QueryTestsParams) ([]Test, error)
	RegisterRunner(ctx context.Context, db DBTX, arg RegisterRunnerParams) (Runner, error)
	RegisterTest(ctx context.Context, db DBTX, arg RegisterTestParams) (Test, error)
	ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error
	RevokeAccessToken(ctx context.Context, db DBTX, id uuid.UUID) error
	RunSummariesForRunner(ctx context.Context, db DBTX, arg RunSummariesForRunnerParams) ([]RunSummariesForRunnerRow, error)
	RunSummariesForTest(ctx context.Context, db DBTX, arg RunSummariesForTestParams) ([]RunSummariesForTestRow, error)
	ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (Run, error)
	SummarizeRunsBreakdownResult(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownResultParams) ([]SummarizeRunsBreakdownResultRow, error)
	SummarizeRunsBreakdownTest(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownTestParams) ([]SummarizeRunsBreakdownTestRow, error)
	TimeoutRuns(ctx context.Context, db DBTX, arg TimeoutRunsParams) error
	UpdateResultData(ctx context.Context, db DBTX, arg UpdateResultDataParams) error
	UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error
	UpdateRunnerHeartbeat(ctx context.Context, db DBTX, id uuid.UUID) error
	UpdateTest(ctx context.Context, db DBTX, arg UpdateTestParams) error
}

var _ Querier = (*Queries)(nil)
