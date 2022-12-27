package handlers

import (
	"context"
	"fmt"
	"main/client"
	"main/database"
	"main/ent"
	"main/ent/target"
	uptime_dash_v1 "main/gen/uptime_dash/v1"
	"main/scheduler"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (s UptimeServer) GetAllTargets(ctx context.Context, req *connect.Request[uptime_dash_v1.GetAllTargetsRequest]) (*connect.Response[uptime_dash_v1.GetAllTargetsResponse], error) {
	entTargets, err := database.Client.Target.
		Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var targets []*uptime_dash_v1.Target
	for _, t := range entTargets {
		targets = append(targets, &uptime_dash_v1.Target{
			Id:              t.ID.String(),
			Name:            t.Name,
			Type:            uptime_dash_v1.Target_Type(uptime_dash_v1.Target_Type_value[t.Type.String()]),
			IntervalSeconds: t.IntervalSeconds,
			TimeoutSeconds:  t.TimeoutSeconds,
			Url:             t.URL,
			Hostname:        t.Hostname,
			Port:            t.Port,
		})
	}

	res := connect.NewResponse(&uptime_dash_v1.GetAllTargetsResponse{
		Targets: targets,
	})

	return res, nil
}

func (s UptimeServer) GetTarget(ctx context.Context, req *connect.Request[uptime_dash_v1.GetTargetRequest]) (*connect.Response[uptime_dash_v1.GetTargetResponse], error) {
	idString, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		log.Error().Err(err).Msg("uuid parse error.")
		return nil, err
	}
	t, err := database.Client.Target.
		Query().
		Where(target.ID(idString)).
		Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("query error.")
		switch {
		case ent.IsNotFound(err):
			return nil, connect.NewError(connect.CodeNotFound, err)
		default:
			return nil, err
		}
	}

	res := connect.NewResponse(&uptime_dash_v1.GetTargetResponse{
		Target: &uptime_dash_v1.Target{
			Id:              t.ID.String(),
			Name:            t.Name,
			Type:            uptime_dash_v1.Target_Type(uptime_dash_v1.Target_Type_value[t.Type.String()]),
			IntervalSeconds: t.IntervalSeconds,
			TimeoutSeconds:  t.TimeoutSeconds,
			Url:             t.URL,
			Hostname:        t.Hostname,
			Port:            t.Port,
		},
	})

	return res, nil
}

func (s UptimeServer) CreateTarget(ctx context.Context, req *connect.Request[uptime_dash_v1.CreateTargetRequest]) (*connect.Response[uptime_dash_v1.CreateTargetResponse], error) {
	contextLogger := log.With().Str("target_name", req.Msg.Target.Name).Str("type", req.Msg.Target.Type.String()).Logger()
	t, err := database.Client.Target.Create().
		SetName(req.Msg.Target.Name).
		SetType(target.Type(req.Msg.Target.Type.String())).
		SetIntervalSeconds(req.Msg.Target.IntervalSeconds).
		SetTimeoutSeconds(req.Msg.Target.TimeoutSeconds).
		SetNillableURL(req.Msg.Target.Url).
		SetNillableHostname(req.Msg.Target.Hostname).
		SetNillablePort(req.Msg.Target.Port).
		Save(ctx)

	if err != nil {
		contextLogger.Error().Err(err).Msg("couldn't create target.")
		return nil, err
	}

	switch t.Type {
	case target.TypeTYPE_HTTP:
		_, err = scheduler.Scheduler.Every(int(t.IntervalSeconds)).Seconds().Tag(t.ID.String()).Do(client.HttpClient, t)
	case target.TypeTYPE_TCP:
		_, err = scheduler.Scheduler.Every(int(t.IntervalSeconds)).Seconds().Tag(t.ID.String()).Do(client.TcpClient, t)
	case target.TypeTYPE_PING:
		fmt.Println("target type is ping")
	}
	if err != nil {
		contextLogger.Error().Err(err).Msg("couldn't create job.")
		return nil, err
	}

	res := connect.NewResponse(&uptime_dash_v1.CreateTargetResponse{
		Target: &uptime_dash_v1.Target{
			Id:              t.ID.String(),
			Type:            req.Msg.Target.Type,
			Name:            t.Name,
			IntervalSeconds: t.IntervalSeconds,
			TimeoutSeconds:  t.TimeoutSeconds,
			Url:             t.URL,
			Hostname:        t.Hostname,
			Port:            t.Port,
		},
	})

	return res, nil
}

func (s UptimeServer) UpdateTarget(ctx context.Context, req *connect.Request[uptime_dash_v1.UpdateTargetRequest]) (*connect.Response[uptime_dash_v1.UpdateTargetResponse], error) {
	idString, err := uuid.Parse(req.Msg.Target.Id)
	if err != nil {
		return nil, err
	}

	contextLogger := log.With().Str("target_name", req.Msg.Target.Name).Str("type", req.Msg.Target.Type.String()).Logger()

	t, err := database.Client.Target.UpdateOneID(idString).
		SetName(req.Msg.Target.Name).
		SetType(target.Type(req.Msg.Target.Type.String())).
		SetIntervalSeconds(req.Msg.Target.IntervalSeconds).
		SetTimeoutSeconds(req.Msg.Target.TimeoutSeconds).
		SetNillableURL(req.Msg.Target.Url).
		SetNillableHostname(req.Msg.Target.Hostname).
		SetNillablePort(req.Msg.Target.Port).
		Save(ctx)

	if err != nil {
		contextLogger.Error().Err(err).Msg("couldn't update target.")
		return nil, err
	}

	if jobs, err := scheduler.Scheduler.FindJobsByTag(req.Msg.Target.Id); err != nil {
		contextLogger.Error().Err(err).Msg("couldn't find job.")
		return nil, err
	} else {
		if _, err := scheduler.Scheduler.Job(jobs[0]).Every(int(t.IntervalSeconds)).Second().Update(); err != nil {
			contextLogger.Error().Err(err).Msg("couldn't update job.")
		}
	}

	res := connect.NewResponse(&uptime_dash_v1.UpdateTargetResponse{
		Target: &uptime_dash_v1.Target{
			Id:              t.ID.String(),
			Type:            req.Msg.Target.Type,
			Name:            t.Name,
			IntervalSeconds: t.IntervalSeconds,
			TimeoutSeconds:  t.TimeoutSeconds,
			Url:             t.URL,
			Hostname:        t.Hostname,
			Port:            t.Port,
		},
	})

	return res, nil
}

func (s UptimeServer) DeleteTarget(ctx context.Context, req *connect.Request[uptime_dash_v1.DeleteTargetRequest]) (*connect.Response[uptime_dash_v1.DeleteTargetResponse], error) {
	idString, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, err
	}

	err = database.Client.Target.
		DeleteOneID(idString).
		Exec(ctx)

	if err := scheduler.Scheduler.RemoveByTag(req.Msg.Id); err != nil {
		log.Error().Err(err).Str("target_id", req.Msg.Id).Msg("couldn't delete task.")
	}

	res := connect.NewResponse(&uptime_dash_v1.DeleteTargetResponse{})

	return res, err
}
