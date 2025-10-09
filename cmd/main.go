package main

import (
	"context"
	"fmt"

	"github.com/godev-lib/golang/config"
	"github.com/godev-lib/golang/psql"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	cinemaroom_controller "github.com/sale-tickets/manager-api/internal/handle/cinema_room"
	health_controller "github.com/sale-tickets/manager-api/internal/handle/health"
	moviethreater_controller "github.com/sale-tickets/manager-api/internal/handle/movie_threater"
	theaterseating_controller "github.com/sale-tickets/manager-api/internal/handle/theater_seating"
	"github.com/sale-tickets/manager-api/internal/model"
	cinemaroom_repo "github.com/sale-tickets/manager-api/internal/repo/cinema_room"
	movietheater_repo "github.com/sale-tickets/manager-api/internal/repo/movie_theater"
	theaterseating_repo "github.com/sale-tickets/manager-api/internal/repo/theater_seating"
	"github.com/sale-tickets/manager-api/internal/router"
	cinemaroom_service "github.com/sale-tickets/manager-api/internal/service/cinema_room"
	movietheater_service "github.com/sale-tickets/manager-api/internal/service/movie_theater"
	theaterseating_service "github.com/sale-tickets/manager-api/internal/service/theater_seating"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	run()
}

func run() {
	componentInts := []fx.Option{}

	componentInts = append(componentInts, fx.Module(
		"config",
		fx.Provide(func() *config.Config {
			return config.NewConfig()
		}),
	))

	componentInts = append(componentInts, fx.Module(
		"connection",
		fx.Provide(func(config *config.Config) *gorm.DB {
			return psql.NewConnectionPsql(config)
		}),
	))

	componentInts = append(componentInts, fx.Module(
		"repo",
		fx.Provide(func(db *gorm.DB) movietheater_repo.MovietheaterRepo {
			return movietheater_repo.NewMovietheaterRepo(db)
		}),
		fx.Provide(func(db *gorm.DB) cinemaroom_repo.CinemaRoomRepo {
			return cinemaroom_repo.NewCinemaRoomRepo(db)
		}),
		fx.Provide(func(db *gorm.DB) theaterseating_repo.TheaterSeatingRepo {
			return theaterseating_repo.NewTheaterSeatingRepo(db)
		}),
	))

	componentInts = append(componentInts, fx.Module(
		"service",
		fx.Provide(func(repo movietheater_repo.MovietheaterRepo) movietheater_service.MovietheaterService {
			return movietheater_service.NewMovietheaterService(repo)
		}),
		fx.Provide(func(repo cinemaroom_repo.CinemaRoomRepo) cinemaroom_service.CinemaRoomService {
			return cinemaroom_service.NewCinemaRoomServic(repo)
		}),
		fx.Provide(func(repo theaterseating_repo.TheaterSeatingRepo) theaterseating_service.TheaterSeatingService {
			return theaterseating_service.NewTheaterSeatingService(repo)
		}),
	))

	componentInts = append(componentInts, fx.Module(
		"handle",
		fx.Provide(func() manager_api.HealthServer {
			return health_controller.NewHandle()
		}),
		fx.Provide(func(service movietheater_service.MovietheaterService) manager_api.MovieTheaterServer {
			return moviethreater_controller.NewHandle(service)
		}),
		fx.Provide(func(service cinemaroom_service.CinemaRoomService) manager_api.CinemaRoomServiceServer {
			return cinemaroom_controller.NewHandle(service)
		}),
		fx.Provide(func(service theaterseating_service.TheaterSeatingService) manager_api.TheaterSeatingServer {
			return theaterseating_controller.NewHandle(service)
		}),
	))

	componentInts = append(componentInts, fx.Module(
		"run",
		fx.Invoke(func(lc fx.Lifecycle, db *gorm.DB) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					err := db.AutoMigrate(
						&model.CinemaRoom{},
						&model.MovieTheater{},
						&model.TheaterSeating{},
					)
					if err != nil {
						return err
					}

					fmt.Println("manager-api migrate db successfully!")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
		fx.Invoke(func(
			config *config.Config,
			healthServer manager_api.HealthServer,
			movieTheaterServer manager_api.MovieTheaterServer,
			cinemaRoomServiceServer manager_api.CinemaRoomServiceServer,
			theaterSeatingServer manager_api.TheaterSeatingServer,
		) {
			router.GrpcServer(
				config,
				healthServer,
				movieTheaterServer,
				cinemaRoomServiceServer,
				theaterSeatingServer,
			)
		}),
		fx.Invoke(func(config *config.Config) {
			router.HttpServer(config)
		}),
	))

	fx.New(componentInts...).Run()
}
