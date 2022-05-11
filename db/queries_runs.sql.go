// Code generated by pggen. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const getRunSQL = `SELECT runs.*, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at
FROM runs
JOIN test_run_configs
ON runs.test_run_config_id = test_run_configs.id
WHERE runs.id = $1;`

type GetRunRow struct {
	ID              string             `json:"id"`
	TestID          string             `json:"test_id"`
	TestRunConfigID string             `json:"test_run_config_id"`
	RunnerID        string             `json:"runner_id"`
	Result          RunResult          `json:"result"`
	Logs            pgtype.JSONB       `json:"logs"`
	ScheduledAt     pgtype.Timestamptz `json:"scheduled_at"`
	StartedAt       pgtype.Timestamptz `json:"started_at"`
	FinishedAt      pgtype.Timestamptz `json:"finished_at"`
	ContainerImage  string             `json:"container_image"`
	Command         string             `json:"command"`
	Args            []string           `json:"args"`
	Env             pgtype.JSONB       `json:"env"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
}

// GetRun implements Querier.GetRun.
func (q *DBQuerier) GetRun(ctx context.Context, id string) (GetRunRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetRun")
	row := q.conn.QueryRow(ctx, getRunSQL, id)
	var item GetRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.CreatedAt); err != nil {
		return item, fmt.Errorf("query GetRun: %w", err)
	}
	return item, nil
}

// GetRunBatch implements Querier.GetRunBatch.
func (q *DBQuerier) GetRunBatch(batch genericBatch, id string) {
	batch.Queue(getRunSQL, id)
}

// GetRunScan implements Querier.GetRunScan.
func (q *DBQuerier) GetRunScan(results pgx.BatchResults) (GetRunRow, error) {
	row := results.QueryRow()
	var item GetRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.CreatedAt); err != nil {
		return item, fmt.Errorf("scan GetRunBatch row: %w", err)
	}
	return item, nil
}

const listRunsSQL = `SELECT runs.*, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at
FROM runs
JOIN test_run_configs
ON runs.test_run_config_id = test_run_configs.id
WHERE
  CASE WHEN $1
    THEN runs.test_id = ANY ($2::uuid[])
    ELSE TRUE
  END AND
  CASE WHEN $3
    THEN runs.test_id = ANY (
      SELECT tests.id
      FROM test_suites
      JOIN tests
      ON tests.labels @> test_suites.labels
      WHERE test_suites.id = ANY ($4::uuid[])
    )
    ELSE TRUE
  END AND
  CASE WHEN $5
    THEN runner_id = ANY ($6::uuid[])
    ELSE TRUE
  END AND
  CASE WHEN $7
    THEN result = ANY ($8::run_result[])
    ELSE TRUE
  END AND
  CASE WHEN $9
    THEN scheduled_at < $10::timestamptz
    ELSE TRUE
  END AND
  CASE WHEN $11
    THEN scheduled_at > $12::timestamptz
    ELSE TRUE
  END AND
  CASE WHEN $13
    THEN started_at < $14::timestamptz
    ELSE TRUE
  END AND
  CASE WHEN $15
    THEN started_at > $16::timestamptz
    ELSE TRUE
  END AND
  CASE WHEN $17
    THEN finished_at < $18::timestamptz
    ELSE TRUE
  END AND
  CASE WHEN $19
    THEN finished_at > $20::timestamptz
    ELSE TRUE
  END;`

type ListRunsParams struct {
	FilterTestIds         bool
	TestIds               []string
	FilterTestSuiteIds    bool
	TestSuiteIds          []string
	FilterRunnerIds       bool
	RunnerIds             []string
	FilterResults         bool
	Results               []RunResult
	FilterScheduledBefore bool
	ScheduledBefore       pgtype.Timestamptz
	FilterScheduledAfter  bool
	ScheduledAfter        pgtype.Timestamptz
	FilterStartedBefore   bool
	StartedBefore         pgtype.Timestamptz
	FilterStartedAfter    bool
	StartedAfter          pgtype.Timestamptz
	FilterFinishedBefore  bool
	FinishedBefore        pgtype.Timestamptz
	FilterFinishedAfter   bool
	FinishedAfter         pgtype.Timestamptz
}

type ListRunsRow struct {
	ID              string             `json:"id"`
	TestID          string             `json:"test_id"`
	TestRunConfigID string             `json:"test_run_config_id"`
	RunnerID        string             `json:"runner_id"`
	Result          RunResult          `json:"result"`
	Logs            pgtype.JSONB       `json:"logs"`
	ScheduledAt     pgtype.Timestamptz `json:"scheduled_at"`
	StartedAt       pgtype.Timestamptz `json:"started_at"`
	FinishedAt      pgtype.Timestamptz `json:"finished_at"`
	ContainerImage  string             `json:"container_image"`
	Command         string             `json:"command"`
	Args            []string           `json:"args"`
	Env             pgtype.JSONB       `json:"env"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
}

// ListRuns implements Querier.ListRuns.
func (q *DBQuerier) ListRuns(ctx context.Context, params ListRunsParams) ([]ListRunsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListRuns")
	rows, err := q.conn.Query(ctx, listRunsSQL, params.FilterTestIds, params.TestIds, params.FilterTestSuiteIds, params.TestSuiteIds, params.FilterRunnerIds, params.RunnerIds, params.FilterResults, q.types.newRunResultArrayInit(params.Results), params.FilterScheduledBefore, params.ScheduledBefore, params.FilterScheduledAfter, params.ScheduledAfter, params.FilterStartedBefore, params.StartedBefore, params.FilterStartedAfter, params.StartedAfter, params.FilterFinishedBefore, params.FinishedBefore, params.FilterFinishedAfter, params.FinishedAfter)
	if err != nil {
		return nil, fmt.Errorf("query ListRuns: %w", err)
	}
	defer rows.Close()
	items := []ListRunsRow{}
	for rows.Next() {
		var item ListRunsRow
		if err := rows.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan ListRuns row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListRuns rows: %w", err)
	}
	return items, err
}

// ListRunsBatch implements Querier.ListRunsBatch.
func (q *DBQuerier) ListRunsBatch(batch genericBatch, params ListRunsParams) {
	batch.Queue(listRunsSQL, params.FilterTestIds, params.TestIds, params.FilterTestSuiteIds, params.TestSuiteIds, params.FilterRunnerIds, params.RunnerIds, params.FilterResults, q.types.newRunResultArrayInit(params.Results), params.FilterScheduledBefore, params.ScheduledBefore, params.FilterScheduledAfter, params.ScheduledAfter, params.FilterStartedBefore, params.StartedBefore, params.FilterStartedAfter, params.StartedAfter, params.FilterFinishedBefore, params.FinishedBefore, params.FilterFinishedAfter, params.FinishedAfter)
}

// ListRunsScan implements Querier.ListRunsScan.
func (q *DBQuerier) ListRunsScan(results pgx.BatchResults) ([]ListRunsRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query ListRunsBatch: %w", err)
	}
	defer rows.Close()
	items := []ListRunsRow{}
	for rows.Next() {
		var item ListRunsRow
		if err := rows.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan ListRunsBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListRunsBatch rows: %w", err)
	}
	return items, err
}

const scheduleRunSQL = `WITH scheduled_run AS (
  INSERT INTO runs (test_id, test_run_config_id)
  VALUES ($1::uuid, $2::uuid)
  RETURNING *
)
SELECT scheduled_run.*, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at
FROM scheduled_run
JOIN test_run_configs
ON scheduled_run.test_run_config_id = test_run_configs.id;`

type ScheduleRunRow struct {
	ID                     string             `json:"id"`
	TestID                 string             `json:"test_id"`
	TestRunConfigID        string             `json:"test_run_config_id"`
	RunnerID               string             `json:"runner_id"`
	Result                 RunResult          `json:"result"`
	Logs                   pgtype.JSONB       `json:"logs"`
	ScheduledAt            pgtype.Timestamptz `json:"scheduled_at"`
	StartedAt              pgtype.Timestamptz `json:"started_at"`
	FinishedAt             pgtype.Timestamptz `json:"finished_at"`
	ContainerImage         string             `json:"container_image"`
	Command                string             `json:"command"`
	Args                   []string           `json:"args"`
	Env                    pgtype.JSONB       `json:"env"`
	TestRunConfigCreatedAt pgtype.Timestamptz `json:"test_run_config_created_at"`
}

// ScheduleRun implements Querier.ScheduleRun.
func (q *DBQuerier) ScheduleRun(ctx context.Context, testID string, testRunConfigID string) (ScheduleRunRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ScheduleRun")
	row := q.conn.QueryRow(ctx, scheduleRunSQL, testID, testRunConfigID)
	var item ScheduleRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.TestRunConfigCreatedAt); err != nil {
		return item, fmt.Errorf("query ScheduleRun: %w", err)
	}
	return item, nil
}

// ScheduleRunBatch implements Querier.ScheduleRunBatch.
func (q *DBQuerier) ScheduleRunBatch(batch genericBatch, testID string, testRunConfigID string) {
	batch.Queue(scheduleRunSQL, testID, testRunConfigID)
}

// ScheduleRunScan implements Querier.ScheduleRunScan.
func (q *DBQuerier) ScheduleRunScan(results pgx.BatchResults) (ScheduleRunRow, error) {
	row := results.QueryRow()
	var item ScheduleRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.TestRunConfigCreatedAt); err != nil {
		return item, fmt.Errorf("scan ScheduleRunBatch row: %w", err)
	}
	return item, nil
}

const assignRunSQL = `UPDATE runs
SET runner_id = $1
FROM test_run_configs
WHERE runs.id = (
  SELECT id
  FROM runs
  WHERE test_id = ANY($2) AND runner_id IS NULL
  ORDER BY scheduled_at ASC
  LIMIT 1
) AND runs.test_run_config_id = test_run_configs.id
RETURNING runs.*, test_run_configs.container_image, test_run_configs.command, test_run_configs.args, test_run_configs.env, test_run_configs.created_at AS test_run_config_created_at;`

type AssignRunRow struct {
	ID                     string             `json:"id"`
	TestID                 string             `json:"test_id"`
	TestRunConfigID        string             `json:"test_run_config_id"`
	RunnerID               string             `json:"runner_id"`
	Result                 RunResult          `json:"result"`
	Logs                   pgtype.JSONB       `json:"logs"`
	ScheduledAt            pgtype.Timestamptz `json:"scheduled_at"`
	StartedAt              pgtype.Timestamptz `json:"started_at"`
	FinishedAt             pgtype.Timestamptz `json:"finished_at"`
	ContainerImage         string             `json:"container_image"`
	Command                string             `json:"command"`
	Args                   []string           `json:"args"`
	Env                    pgtype.JSONB       `json:"env"`
	TestRunConfigCreatedAt pgtype.Timestamptz `json:"test_run_config_created_at"`
}

// AssignRun implements Querier.AssignRun.
func (q *DBQuerier) AssignRun(ctx context.Context, runnerID string, testIds []string) (AssignRunRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "AssignRun")
	row := q.conn.QueryRow(ctx, assignRunSQL, runnerID, testIds)
	var item AssignRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.TestRunConfigCreatedAt); err != nil {
		return item, fmt.Errorf("query AssignRun: %w", err)
	}
	return item, nil
}

// AssignRunBatch implements Querier.AssignRunBatch.
func (q *DBQuerier) AssignRunBatch(batch genericBatch, runnerID string, testIds []string) {
	batch.Queue(assignRunSQL, runnerID, testIds)
}

// AssignRunScan implements Querier.AssignRunScan.
func (q *DBQuerier) AssignRunScan(results pgx.BatchResults) (AssignRunRow, error) {
	row := results.QueryRow()
	var item AssignRunRow
	if err := row.Scan(&item.ID, &item.TestID, &item.TestRunConfigID, &item.RunnerID, &item.Result, &item.Logs, &item.ScheduledAt, &item.StartedAt, &item.FinishedAt, &item.ContainerImage, &item.Command, &item.Args, &item.Env, &item.TestRunConfigCreatedAt); err != nil {
		return item, fmt.Errorf("scan AssignRunBatch row: %w", err)
	}
	return item, nil
}

const updateRunSQL = `UPDATE runs
SET
  result = $1,
  logs = $2,
  started_at = $3::timestamptz,
  finished_at = $4::timestamptz
WHERE id = $5::uuid;`

type UpdateRunParams struct {
	Result     RunResult
	Logs       pgtype.JSONB
	StartedAt  pgtype.Timestamptz
	FinishedAt pgtype.Timestamptz
	ID         string
}

// UpdateRun implements Querier.UpdateRun.
func (q *DBQuerier) UpdateRun(ctx context.Context, params UpdateRunParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateRun")
	cmdTag, err := q.conn.Exec(ctx, updateRunSQL, params.Result, params.Logs, params.StartedAt, params.FinishedAt, params.ID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateRun: %w", err)
	}
	return cmdTag, err
}

// UpdateRunBatch implements Querier.UpdateRunBatch.
func (q *DBQuerier) UpdateRunBatch(batch genericBatch, params UpdateRunParams) {
	batch.Queue(updateRunSQL, params.Result, params.Logs, params.StartedAt, params.FinishedAt, params.ID)
}

// UpdateRunScan implements Querier.UpdateRunScan.
func (q *DBQuerier) UpdateRunScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateRunBatch: %w", err)
	}
	return cmdTag, err
}
