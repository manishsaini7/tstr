// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: runs.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const appendLogsToRun = `-- name: AppendLogsToRun :exec
UPDATE runs
SET logs = COALESCE(logs, '[]'::jsonb) || $1
WHERE id = $2
`

type AppendLogsToRunParams struct {
	Logs pgtype.JSONB
	ID   uuid.UUID
}

func (q *Queries) AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error {
	_, err := db.Exec(ctx, appendLogsToRun, arg.Logs, arg.ID)
	return err
}

const assignRun = `-- name: AssignRun :one
UPDATE runs
SET runner_id = $1::uuid
FROM test_run_configs
WHERE runs.id = (
  SELECT id
  FROM runs AS selected_runs
  WHERE selected_runs.test_id = ANY($2::uuid[]) AND selected_runs.runner_id IS NULL
  ORDER BY selected_runs.scheduled_at ASC
  LIMIT 1
) AND runs.test_run_config_id = test_run_configs.id
RETURNING runs.id, runs.test_id, runs.test_run_config_id, runs.runner_id, runs.result, runs.logs, runs.scheduled_at, runs.started_at, runs.finished_at, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at
`

type AssignRunParams struct {
	RunnerID uuid.UUID
	TestIds  []uuid.UUID
}

type AssignRunRow struct {
	ID                     uuid.UUID
	TestID                 uuid.UUID
	TestRunConfigID        uuid.UUID
	RunnerID               uuid.NullUUID
	Result                 NullRunResult
	Logs                   pgtype.JSONB
	ScheduledAt            sql.NullTime
	StartedAt              sql.NullTime
	FinishedAt             sql.NullTime
	ContainerImage         string
	Command                sql.NullString
	Args                   []string
	Env                    pgtype.JSONB
	TestRunConfigCreatedAt sql.NullTime
}

func (q *Queries) AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (AssignRunRow, error) {
	row := db.QueryRow(ctx, assignRun, arg.RunnerID, arg.TestIds)
	var i AssignRunRow
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfigID,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
		&i.ContainerImage,
		&i.Command,
		&i.Args,
		&i.Env,
		&i.TestRunConfigCreatedAt,
	)
	return i, err
}

const getRun = `-- name: GetRun :one
SELECT runs.id, runs.test_id, runs.test_run_config_id, runs.runner_id, runs.result, runs.logs, runs.scheduled_at, runs.started_at, runs.finished_at, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at
FROM runs
JOIN test_run_configs
ON runs.test_run_config_id = test_run_configs.id
WHERE runs.id = $1
`

type GetRunRow struct {
	ID                     uuid.UUID
	TestID                 uuid.UUID
	TestRunConfigID        uuid.UUID
	RunnerID               uuid.NullUUID
	Result                 NullRunResult
	Logs                   pgtype.JSONB
	ScheduledAt            sql.NullTime
	StartedAt              sql.NullTime
	FinishedAt             sql.NullTime
	ContainerImage         string
	Command                sql.NullString
	Args                   []string
	Env                    pgtype.JSONB
	TestRunConfigCreatedAt sql.NullTime
}

func (q *Queries) GetRun(ctx context.Context, db DBTX, id uuid.UUID) (GetRunRow, error) {
	row := db.QueryRow(ctx, getRun, id)
	var i GetRunRow
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfigID,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
		&i.ContainerImage,
		&i.Command,
		&i.Args,
		&i.Env,
		&i.TestRunConfigCreatedAt,
	)
	return i, err
}

const listRuns = `-- name: ListRuns :many
SELECT runs.id, runs.test_id, runs.test_run_config_id, runs.runner_id, runs.result, runs.logs, runs.scheduled_at, runs.started_at, runs.finished_at, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at
FROM runs
JOIN test_run_configs
ON runs.test_run_config_id = test_run_configs.id
WHERE
  ($1::uuid[] IS NULL OR runs.test_id = ANY ($1::uuid[])) AND
  ($2::uuid[] IS NULL OR runs.test_id = ANY (
      SELECT tests.id
      FROM test_suites
      JOIN tests
      ON tests.labels @> test_suites.labels
      WHERE test_suites.id = ANY ($2::uuid[])
    )) AND
  ($3::uuid[] IS NULL OR runner_id = ANY ($3::uuid[])) AND
  ($4::run_result[] IS NULL OR result = ANY ($4::run_result[])) AND
  ($5::timestamptz IS NULL OR scheduled_at < $5::timestamptz) AND
  ($6::timestamptz IS NULL OR scheduled_at > $6::timestamptz) AND
  ($7::timestamptz IS NULL OR started_at < $7::timestamptz) AND
  ($8::timestamptz IS NULL OR started_at > $8::timestamptz) AND
  ($9::timestamptz IS NULL OR finished_at < $9::timestamptz) AND
  ($10::timestamptz IS NULL OR finished_at > $10::timestamptz)
`

type ListRunsParams struct {
	TestIds         []uuid.UUID
	TestSuiteIds    []uuid.UUID
	RunnerIds       []uuid.UUID
	Results         []RunResult
	ScheduledBefore sql.NullTime
	ScheduledAfter  sql.NullTime
	StartedBefore   sql.NullTime
	StartedAfter    sql.NullTime
	FinishedBefore  sql.NullTime
	FinishedAfter   sql.NullTime
}

type ListRunsRow struct {
	ID              uuid.UUID
	TestID          uuid.UUID
	TestRunConfigID uuid.UUID
	RunnerID        uuid.NullUUID
	Result          NullRunResult
	Logs            pgtype.JSONB
	ScheduledAt     sql.NullTime
	StartedAt       sql.NullTime
	FinishedAt      sql.NullTime
	ContainerImage  string
	Command         sql.NullString
	Args            []string
	Env             pgtype.JSONB
	CreatedAt       sql.NullTime
}

func (q *Queries) ListRuns(ctx context.Context, db DBTX, arg ListRunsParams) ([]ListRunsRow, error) {
	rows, err := db.Query(ctx, listRuns,
		arg.TestIds,
		arg.TestSuiteIds,
		arg.RunnerIds,
		arg.Results,
		arg.ScheduledBefore,
		arg.ScheduledAfter,
		arg.StartedBefore,
		arg.StartedAfter,
		arg.FinishedBefore,
		arg.FinishedAfter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRunsRow
	for rows.Next() {
		var i ListRunsRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfigID,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ContainerImage,
			&i.Command,
			&i.Args,
			&i.Env,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const queryRuns = `-- name: QueryRuns :many
SELECT runs.id, runs.test_id, runs.test_run_config_id, runs.runner_id, runs.result, runs.logs, runs.scheduled_at, runs.started_at, runs.finished_at, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at
FROM runs
JOIN test_run_configs
ON runs.test_run_config_id = test_run_configs.id
WHERE
  ($1::uuid[] IS NULL OR runs.id = ANY ($1::uuid[])) AND
  ($2::uuid[] IS NULL OR runs.test_id = ANY ($2::uuid[])) AND
  ($3::uuid[] IS NULL OR runs.test_id = ANY (
      SELECT tests.id
      FROM test_suites
      JOIN tests
      ON tests.labels @> test_suites.labels
      WHERE test_suites.id = ANY ($3::uuid[])
    )) AND
  ($4::uuid[] IS NULL OR runner_id = ANY ($4::uuid[])) AND
  ($5::run_result[] IS NULL OR result = ANY ($5::run_result[])) AND
  ($6::timestamptz IS NULL OR scheduled_at < $6::timestamptz) AND
  ($7::timestamptz IS NULL OR scheduled_at > $7::timestamptz) AND
  ($8::timestamptz IS NULL OR started_at < $8::timestamptz) AND
  ($9::timestamptz IS NULL OR started_at > $9::timestamptz) AND
  ($10::timestamptz IS NULL OR finished_at < $10::timestamptz) AND
  ($11::timestamptz IS NULL OR finished_at > $11::timestamptz)
`

type QueryRunsParams struct {
	Ids             []uuid.UUID
	TestIds         []uuid.UUID
	TestSuiteIds    []uuid.UUID
	RunnerIds       []uuid.UUID
	Results         []RunResult
	ScheduledBefore sql.NullTime
	ScheduledAfter  sql.NullTime
	StartedBefore   sql.NullTime
	StartedAfter    sql.NullTime
	FinishedBefore  sql.NullTime
	FinishedAfter   sql.NullTime
}

type QueryRunsRow struct {
	ID                     uuid.UUID
	TestID                 uuid.UUID
	TestRunConfigID        uuid.UUID
	RunnerID               uuid.NullUUID
	Result                 NullRunResult
	Logs                   pgtype.JSONB
	ScheduledAt            sql.NullTime
	StartedAt              sql.NullTime
	FinishedAt             sql.NullTime
	ContainerImage         string
	Command                sql.NullString
	Args                   []string
	Env                    pgtype.JSONB
	TestRunConfigCreatedAt sql.NullTime
}

func (q *Queries) QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]QueryRunsRow, error) {
	rows, err := db.Query(ctx, queryRuns,
		arg.Ids,
		arg.TestIds,
		arg.TestSuiteIds,
		arg.RunnerIds,
		arg.Results,
		arg.ScheduledBefore,
		arg.ScheduledAfter,
		arg.StartedBefore,
		arg.StartedAfter,
		arg.FinishedBefore,
		arg.FinishedAfter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QueryRunsRow
	for rows.Next() {
		var i QueryRunsRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfigID,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ContainerImage,
			&i.Command,
			&i.Args,
			&i.Env,
			&i.TestRunConfigCreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetOrphanedRuns = `-- name: ResetOrphanedRuns :exec
UPDATE runs
SET runner_id = NULL
WHERE
  result = 'unknown' AND
  started_at IS NULL AND
  scheduled_at < $1::timestamptz
`

func (q *Queries) ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error {
	_, err := db.Exec(ctx, resetOrphanedRuns, before)
	return err
}

const runSummaryForTest = `-- name: RunSummaryForTest :many
SELECT id, test_id, test_run_config_id, runner_id, result, scheduled_at, started_at, finished_at
FROM runs
WHERE runs.test_id = $1
ORDER by runs.started_at desc
LIMIT $2
`

type RunSummaryForTestParams struct {
	TestID uuid.UUID
	Limit  int32
}

type RunSummaryForTestRow struct {
	ID              uuid.UUID
	TestID          uuid.UUID
	TestRunConfigID uuid.UUID
	RunnerID        uuid.NullUUID
	Result          NullRunResult
	ScheduledAt     sql.NullTime
	StartedAt       sql.NullTime
	FinishedAt      sql.NullTime
}

func (q *Queries) RunSummaryForTest(ctx context.Context, db DBTX, arg RunSummaryForTestParams) ([]RunSummaryForTestRow, error) {
	rows, err := db.Query(ctx, runSummaryForTest, arg.TestID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RunSummaryForTestRow
	for rows.Next() {
		var i RunSummaryForTestRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfigID,
			&i.RunnerID,
			&i.Result,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const scheduleRun = `-- name: ScheduleRun :one
WITH scheduled_run AS (
  INSERT INTO runs (test_id, test_run_config_id)
  VALUES ($1::uuid, $2::uuid)
  RETURNING id, test_id, test_run_config_id, runner_id, result, logs, scheduled_at, started_at, finished_at
)
SELECT scheduled_run.id, scheduled_run.test_id, scheduled_run.test_run_config_id, scheduled_run.runner_id, scheduled_run.result, scheduled_run.logs, scheduled_run.scheduled_at, scheduled_run.started_at, scheduled_run.finished_at, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at
FROM scheduled_run
JOIN test_run_configs
ON scheduled_run.test_run_config_id = test_run_configs.id
`

type ScheduleRunParams struct {
	TestID          uuid.UUID
	TestRunConfigID uuid.UUID
}

type ScheduleRunRow struct {
	ID                     uuid.UUID
	TestID                 uuid.UUID
	TestRunConfigID        uuid.UUID
	RunnerID               uuid.NullUUID
	Result                 NullRunResult
	Logs                   pgtype.JSONB
	ScheduledAt            sql.NullTime
	StartedAt              sql.NullTime
	FinishedAt             sql.NullTime
	ContainerImage         string
	Command                sql.NullString
	Args                   []string
	Env                    pgtype.JSONB
	TestRunConfigCreatedAt sql.NullTime
}

func (q *Queries) ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (ScheduleRunRow, error) {
	row := db.QueryRow(ctx, scheduleRun, arg.TestID, arg.TestRunConfigID)
	var i ScheduleRunRow
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfigID,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
		&i.ContainerImage,
		&i.Command,
		&i.Args,
		&i.Env,
		&i.TestRunConfigCreatedAt,
	)
	return i, err
}

const updateRun = `-- name: UpdateRun :exec
UPDATE runs
SET
  result = $1,
  started_at = $2::timestamptz,
  finished_at = $3::timestamptz
WHERE id = $4
`

type UpdateRunParams struct {
	Result     NullRunResult
	StartedAt  sql.NullTime
	FinishedAt sql.NullTime
	ID         uuid.UUID
}

func (q *Queries) UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error {
	_, err := db.Exec(ctx, updateRun,
		arg.Result,
		arg.StartedAt,
		arg.FinishedAt,
		arg.ID,
	)
	return err
}
