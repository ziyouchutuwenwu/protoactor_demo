package cluster

import (
	"time"

	"github.com/AsynkronIT/gonet"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
)

var (
	cp ClusterProvider
)

func Start(clusterName, address string, provider ClusterProvider) {
	//TODO: make it possible to become a cluster even if remoting is already started
	remote.Start(address)

	cp = provider
	address = actor.ProcessRegistry.Address
	h, p := gonet.GetAddress(address)
	plog.Info("Starting Proto.Actor cluster", log.String("address", address))
	kinds := remote.GetKnownKinds()

	//for each known kind, spin up a partition-kind actor to handle all requests for that kind
	spawnPartitionActors(kinds)
	subscribePartitionKindsToEventStream()
	spawnPidCacheActor()
	subscribePidCacheMemberStatusEventStream()
	spawnMembershipActor()
	subscribeMembershipActorToEventStream()
	cp.RegisterMember(clusterName, h, p, kinds)
	cp.MonitorMemberStatusChanges()
}

func Shutdown(graceful bool) {
	if graceful {
		cp.Shutdown()

		unsubMembershipActorToEventStream()
		stopMembershipActor()
		unsubPidCacheMemberStatusEventStream()
		stopPidCacheActor()
		unsubPartitionKindsToEventStream()
		stopPartitionActors()
	}

	remote.Shutdown(graceful)

	address := actor.ProcessRegistry.Address
	plog.Info("Stopped Proto.Actor cluster", log.String("address", address))
}

//Get a PID to a virtual actor
func Get(name string, kind string) (*actor.PID, remote.ResponseStatusCode) {

	req := &pidCacheRequest{
		kind: kind,
		name: name,
	}

	res, err := pidCacheActorPid.RequestFuture(req, 5*time.Second).Result()
	if err != nil {
		plog.Error("ActorPidRequest timed out", log.String("name", name), log.Error(err))
		return nil, remote.ResponseStatusCodeTIMEOUT
	}
	typed, ok := res.(*pidCacheResponse)
	if !ok {
		plog.Error("ActorPidRequest returned incorrect response", log.String("name", name))
		return nil, remote.ResponseStatusCodeUNAVAILABLE
	}
	return typed.pid, typed.status
}

func RemoveCache(name string) {
	pidCacheActorPid.Tell(&removePidCacheRequest{name})
}
