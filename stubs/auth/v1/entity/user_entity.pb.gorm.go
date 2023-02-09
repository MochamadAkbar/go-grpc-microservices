package entity

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
	time "time"
)

type UserEntityORM struct {
	AccessTokenEntityId *string
	CreatedAt           *time.Time
	Email               string `gorm:"unique;not null"`
	Id                  string `gorm:"type:uuid;primary_key;default:uuid_generate_v4();not null"`
	Name                string `gorm:"not null"`
	Password            string `gorm:"not null"`
	UpdatedAt           *time.Time
}

// TableName overrides the default tablename generated by GORM
func (UserEntityORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *UserEntity) ToORM(ctx context.Context) (UserEntityORM, error) {
	to := UserEntityORM{}
	var err error
	if prehook, ok := interface{}(m).(UserEntityWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Email = m.Email
	to.Password = m.Password
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(UserEntityWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserEntityORM) ToPB(ctx context.Context) (UserEntity, error) {
	to := UserEntity{}
	var err error
	if prehook, ok := interface{}(m).(UserEntityWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Email = m.Email
	to.Password = m.Password
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(UserEntityWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type UserEntity the arg will be the target, the caller the one being converted from

// UserEntityBeforeToORM called before default ToORM code
type UserEntityWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserEntityORM) error
}

// UserEntityAfterToORM called after default ToORM code
type UserEntityWithAfterToORM interface {
	AfterToORM(context.Context, *UserEntityORM) error
}

// UserEntityBeforeToPB called before default ToPB code
type UserEntityWithBeforeToPB interface {
	BeforeToPB(context.Context, *UserEntity) error
}

// UserEntityAfterToPB called after default ToPB code
type UserEntityWithAfterToPB interface {
	AfterToPB(context.Context, *UserEntity) error
}

// DefaultCreateUserEntity executes a basic gorm create call
func DefaultCreateUserEntity(ctx context.Context, in *UserEntity, db *gorm.DB) (*UserEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserEntityORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUserEntity(ctx context.Context, in *UserEntity, db *gorm.DB) (*UserEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserEntityORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserEntityORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserEntityORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserEntityORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUserEntity(ctx context.Context, in *UserEntity, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserEntityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserEntityORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserEntitySet(ctx context.Context, in []*UserEntity, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserEntityORM{})).(UserEntityORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserEntityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserEntityORM{})).(UserEntityORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserEntityORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*UserEntity, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*UserEntity, *gorm.DB) error
}

// DefaultStrictUpdateUserEntity clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUserEntity(ctx context.Context, in *UserEntity, db *gorm.DB) (*UserEntity, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUserEntity")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserEntityORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserEntityORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUserEntity executes a basic gorm update call with patch behavior
func DefaultPatchUserEntity(ctx context.Context, in *UserEntity, updateMask *field_mask.FieldMask, db *gorm.DB) (*UserEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj UserEntity
	var err error
	if hook, ok := interface{}(&pbObj).(UserEntityWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUserEntity(ctx, &UserEntity{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserEntityWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUserEntity(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserEntityWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUserEntity(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserEntityWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserEntityWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *UserEntity, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserEntityWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *UserEntity, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserEntityWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *UserEntity, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserEntityWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *UserEntity, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUserEntity executes a bulk gorm update call with patch behavior
func DefaultPatchSetUserEntity(ctx context.Context, objects []*UserEntity, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*UserEntity, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*UserEntity, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUserEntity(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUserEntity patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUserEntity(ctx context.Context, patchee *UserEntity, patcher *UserEntity, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*UserEntity, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Email" {
			patchee.Email = patcher.Email
			continue
		}
		if f == prefix+"Password" {
			patchee.Password = patcher.Password
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUserEntity executes a gorm list call
func DefaultListUserEntity(ctx context.Context, db *gorm.DB) ([]*UserEntity, error) {
	in := UserEntity{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserEntityORM{}, &UserEntity{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserEntityORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserEntityORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*UserEntity{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserEntityORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserEntityORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserEntityORM) error
}