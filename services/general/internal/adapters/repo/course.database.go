package repo

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/adapters/repo/database"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

const delimeter string = "\n\n"

type CourseDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.CourseRepository = &CourseDatabase{}

func (db *CourseDatabase) Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourseData) (uuid.UUID, error) {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	inner := database.New(db.Pool).WithTx(tx)

	var params database.CreateCourseParams
	copier.Copy(&params, &data)
	params.InstructorID = instructorID
	params.LearningObjectives = strings.Join(data.LearningObjectives, delimeter)
	id, err := inner.CreateCourse(ctx, params)
	if err != nil {
		return uuid.Nil, err
	}

	for _, category := range data.Categories {
		err = inner.CreateCategory(ctx, category)
		if err != nil {
			return uuid.Nil, err
		}

		err = inner.AddCourseCategory(ctx, database.AddCourseCategoryParams{
			ID:       id,
			Category: category,
		})
		if err != nil {
			return uuid.Nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (db *CourseDatabase) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	var course domain.Course
	copier.Copy(&course, &raw)
	course.Requirements = strings.Split(raw.Requirements, delimeter)
	course.LearningObjectives = strings.Split(raw.LearningObjectives, delimeter)
	course.TargetAudiences = strings.Split(raw.TargetAudiences, delimeter)

	return &course, nil
}
