package types

import (
	"database/sql"

	commonv1 "github.com/nanzhong/tstr/api/common/v1"
	"github.com/nanzhong/tstr/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromAccessTokenScope(scope commonv1.AccessToken_Scope) db.AccessTokenScope {
	switch scope {
	case commonv1.AccessToken_ADMIN:
		return db.AccessTokenScopeAdmin
	case commonv1.AccessToken_CONTROL_R:
		return db.AccessTokenScopeControlR
	case commonv1.AccessToken_CONTROL_RW:
		return db.AccessTokenScopeControlRw
	case commonv1.AccessToken_RUNNER:
		return db.AccessTokenScopeRunner
	case commonv1.AccessToken_DATA:
		return db.AccessTokenScopeData
	default:
		panic("unknown access token scope:" + scope.String())
	}
}

func FromAccessTokenScopes(scopes []commonv1.AccessToken_Scope) []db.AccessTokenScope {
	var dbScopes []db.AccessTokenScope
	for _, s := range scopes {
		dbScopes = append(dbScopes, FromAccessTokenScope(s))
	}
	return dbScopes
}

func FromRunResult(result commonv1.Run_Result) db.RunResult {
	switch result {
	case commonv1.Run_ERROR:
		return db.RunResultError
	case commonv1.Run_FAIL:
		return db.RunResultFail
	case commonv1.Run_PASS:
		return db.RunResultPass
	default:
		return db.RunResultUnknown
	}
}

func FromRunResults(results []commonv1.Run_Result) []db.RunResult {
	var dbResults []db.RunResult
	for _, r := range results {
		dbResults = append(dbResults, FromRunResult(r))
	}
	return dbResults
}

func FromProtoTimestampAsNullTime(ts *timestamppb.Timestamp) sql.NullTime {
	if ts == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Valid: true, Time: ts.AsTime()}
}
