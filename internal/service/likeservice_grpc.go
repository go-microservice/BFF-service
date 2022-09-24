package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/copier"

	"github.com/go-eagle/eagle/pkg/log"
	"github.com/spf13/cast"

	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-microservice/ins-api/api/micro/moment/v1"
)

const (
	LikeTypePost    = 1
	LikeTypeComment = 2
)

var (
	_ pb.LikeServiceServer = (*LikeServiceServer)(nil)
)

type LikeServiceServer struct {
	pb.UnimplementedLikeServiceServer

	momentRPC momentv1.LikeServiceClient
	userRPC   userv1.UserServiceClient
}

func NewLikeServiceServer(repo momentv1.LikeServiceClient, userRepo userv1.UserServiceClient) *LikeServiceServer {
	return &LikeServiceServer{
		momentRPC: repo,
		userRPC:   userRepo,
	}
}

func (s *LikeServiceServer) CreatePostLike(ctx context.Context, req *pb.CreatePostLikeRequest) (*pb.CreatePostLikeReply, error) {
	in := &momentv1.CreateLikeRequest{
		UserId:  GetCurrentUserID(ctx),
		ObjType: LikeTypePost,
		ObjId:   req.GetPostId(),
	}
	_, err := s.momentRPC.CreateLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	return &pb.CreatePostLikeReply{}, nil
}

func (s *LikeServiceServer) DeletePostLike(ctx context.Context, req *pb.DeletePostLikeRequest) (*pb.DeletePostLikeReply, error) {
	in := &momentv1.DeleteLikeRequest{
		UserId:  GetCurrentUserID(ctx),
		ObjType: LikeTypePost,
		ObjId:   req.GetPostId(),
	}
	_, err := s.momentRPC.DeleteLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	return &pb.DeletePostLikeReply{}, nil
}

func (s *LikeServiceServer) ListPostLike(ctx context.Context, req *pb.ListPostLikeRequest) (*pb.ListPostLikeReply, error) {
	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	in := &momentv1.ListPostLikeRequest{
		PostId: req.GetPostId(),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.momentRPC.ListPostLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	likes := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(likes)) > limit {
		hasMore = true
		lastId = cast.ToString(likes[len(likes)-1].Id)
		likes = likes[0 : len(likes)-1]
	}
	pbLikes, err := s.assembleData(ctx, likes)
	if err != nil {
		return nil, err
	}

	return &pb.ListPostLikeReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbLikes,
	}, nil
}

func (s *LikeServiceServer) CreateCommentLike(ctx context.Context, req *pb.CreateCommentLikeRequest) (*pb.CreateCommentLikeReply, error) {
	in := &momentv1.CreateLikeRequest{
		UserId:  GetCurrentUserID(ctx),
		ObjType: LikeTypeComment,
		ObjId:   req.GetCommentId(),
	}
	_, err := s.momentRPC.CreateLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	return &pb.CreateCommentLikeReply{}, nil
}

func (s *LikeServiceServer) DeleteCommentLike(ctx context.Context, req *pb.DeleteCommentLikeRequest) (*pb.DeleteCommentLikeReply, error) {
	in := &momentv1.DeleteLikeRequest{
		UserId:  GetCurrentUserID(ctx),
		ObjType: LikeTypeComment,
		ObjId:   req.GetCommentId(),
	}
	_, err := s.momentRPC.DeleteLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}
	return &pb.DeleteCommentLikeReply{}, nil
}

func (s *LikeServiceServer) ListCommentLike(ctx context.Context, req *pb.ListCommentLikeRequest) (*pb.ListCommentLikeReply, error) {
	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	in := &momentv1.ListCommentLikeRequest{
		CommentId: req.GetCommentId(),
		LastId:    cast.ToInt64(req.GetLastId()),
		Limit:     limit + 1,
	}
	ret, err := s.momentRPC.ListCommentLike(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	likes := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(likes)) > limit {
		hasMore = true
		lastId = cast.ToString(likes[len(likes)-1].Id)
		likes = likes[0 : len(likes)-1]
	}
	pbLikes, err := s.assembleData(ctx, likes)
	if err != nil {
		return nil, err
	}

	return &pb.ListCommentLikeReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbLikes,
	}, nil
}

func (s *LikeServiceServer) assembleData(ctx context.Context, likes []*momentv1.Like) ([]*pb.Like, error) {
	// batch get user data
	var (
		userIDs []int64
	)
	for _, v := range likes {
		userIDs = append(userIDs, v.UserId)
	}

	userReply, err := s.userRPC.BatchGetUsers(ctx, &userv1.BatchGetUsersRequest{Ids: userIDs})
	if err != nil {
		return nil, err
	}
	users := userReply.GetUsers()
	// to map
	userMap := make(map[int64]*userv1.User)
	for _, v := range users {
		userMap[v.Id] = v
	}

	var (
		pbLikes []*pb.Like
		m       sync.Map
		mu      sync.Mutex
	)

	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	go func() {
		select {
		case <-finished:
			return
		case err := <-errChan:
			if err != nil {
				// NOTE: if need, record log to file
			}
		case <-time.After(3 * time.Second):
			log.Warn(fmt.Errorf("list users timeout after 3 seconds"))
			return
		}
	}()

	for _, like := range likes {
		wg.Add(1)
		like := like
		go func(info *momentv1.Like) {
			defer func() {
				wg.Done()
			}()

			mu.Lock()
			defer mu.Unlock()

			pbLike, err := convertLike(info)
			if err != nil {
				return
			}
			// user
			user, ok := userMap[like.UserId]
			if !ok {
				return
			}
			pbLike.User, err = convertUser(user)
			if err != nil {
				errChan <- err
			}

			m.Store(info.Id, pbLike)
		}(like)

	}

	wg.Wait()
	close(errChan)
	close(finished)

	for _, pid := range likes {
		post, _ := m.Load(pid.Id)
		if post == nil {
			continue
		}
		pbLikes = append(pbLikes, post.(*pb.Like))
	}

	return pbLikes, nil
}

func convertLike(p *momentv1.Like) (*pb.Like, error) {
	like := pb.Like{}
	err := copier.Copy(&like, &p)
	if err != nil {
		return nil, err
	}
	return &like, nil
}
