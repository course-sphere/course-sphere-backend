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
	copier.CopyWithOption(&params, &data, copier.Option{DeepCopy: true})
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

	for _, otherID := range data.Prerequisites {
		err = inner.AddCoursePrerequisite(ctx, database.AddCoursePrerequisiteParams{
			ID:      id,
			OtherID: otherID,
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

	course.Requirements = make([]string, 0)
	if raw.Requirements != nil {
		course.Requirements = strings.Split(*raw.Requirements, delimeter)
	}
	course.LearningObjectives = strings.Split(raw.LearningObjectives, delimeter)

	course.TargetAudiences = make([]string, 0)
	if raw.TargetAudiences != nil {
		course.TargetAudiences = strings.Split(*raw.TargetAudiences, delimeter)
	}

	categories, err := inner.GetCourseCategories(ctx, id)
	if err != nil {
		return nil, err
	}
	course.Categories = categories

	prerequisites, err := inner.GetCoursePrerequisites(ctx, id)
	if err != nil {
		return nil, err
	}
	course.Prerequisites = prerequisites

	return &course, nil
}

func (db *CourseDatabase) GetAll(ctx context.Context) ([]domain.Course, error) {
	inner := database.New(db.Pool)

	ids, err := inner.GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	courses := make([]domain.Course, 0)
	for _, id := range ids {
		course, err := db.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		courses = append(courses, *course)
	}

	return courses, nil
}

func (db *CourseDatabase) Update(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, data domain.UpdateCourseData) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	inner := database.New(db.Pool).WithTx(tx)

	params := database.UpdateCourseParams{
		ID:           id,
		InstructorID: instructorID,
	}
	copier.CopyWithOption(&params, &data, copier.Option{DeepCopy: true})

	if data.Requirements != nil {
		requirements := strings.Join(*data.Requirements, delimeter)
		params.Requirements = &requirements
	}

	if data.LearningObjectives != nil {
		learningObjectives := strings.Join(*data.LearningObjectives, delimeter)
		params.LearningObjectives = &learningObjectives
	}

	if data.TargetAudiences != nil {
		targetAudiences := strings.Join(*data.TargetAudiences, delimeter)
		params.TargetAudiences = &targetAudiences
	}
	err = inner.UpdateCourse(ctx, params)
	if err != nil {
		return err
	}

	if data.Categories != nil {
		err = inner.DeleteCourseCategories(ctx, id)
		if err != nil {
			return err
		}

		for _, category := range *data.Categories {
			err = inner.CreateCategory(ctx, category)
			if err != nil {
				return err
			}

			err = inner.AddCourseCategory(ctx, database.AddCourseCategoryParams{
				ID:       id,
				Category: category,
			})
			if err != nil {
				return err
			}
		}
	}

	if data.Prerequisites != nil {
		err = inner.DeleteCoursePrerequisites(ctx, id)
		if err != nil {
			return err
		}

		for _, otherID := range *data.Prerequisites {
			err = inner.AddCoursePrerequisite(ctx, database.AddCoursePrerequisiteParams{
				ID:      id,
				OtherID: otherID,
			})
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
