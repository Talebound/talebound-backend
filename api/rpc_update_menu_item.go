package api

import (
	"context"
	"database/sql"
	"github.com/the-medo/talebound-backend/api/converters"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateMenuItem(ctx context.Context, req *pb.UpdateMenuItemRequest) (*pb.MenuItem, error) {
	violations := validateUpdateMenuItemRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.CheckMenuAdmin(ctx, req.GetMenuId(), false)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "failed update menu: %v", err)
	}

	arg := db.UpdateMenuItemParams{
		ID: req.GetMenuItemId(),
		MenuItemCode: sql.NullString{
			String: req.GetCode(),
			Valid:  req.Code != nil,
		},
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Position: sql.NullInt32{
			Int32: req.GetPosition(),
			Valid: req.Position != nil,
		},
		ParentItemID: sql.NullInt32{
			Int32: req.GetParentItemId(),
			Valid: req.ParentItemId != nil,
		},
		DescriptionPostID: sql.NullInt32{
			Int32: req.GetDescriptionPostId(),
			Valid: req.DescriptionPostId != nil,
		},
	}

	menuItem, err := server.store.UpdateMenuItem(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update menu item: %s", err)
	}

	rsp := converters.ConvertMenuItem(menuItem)

	return rsp, nil
}

func validateUpdateMenuItemRequest(req *pb.UpdateMenuItemRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateMenuId(req.GetMenuId()); err != nil {
		violations = append(violations, FieldViolation("menu_id", err))
	}

	if err := validator.ValidateMenuItemId(req.GetMenuItemId()); err != nil {
		violations = append(violations, FieldViolation("menu_item_id", err))
	}

	if req.Code != nil {
		if err := validator.ValidateMenuItemCode(req.GetCode()); err != nil {
			violations = append(violations, FieldViolation("code", err))
		}
	}

	if req.Name != nil {
		if err := validator.ValidateMenuItemName(req.GetName()); err != nil {
			violations = append(violations, FieldViolation("name", err))
		}
	}

	if req.Position != nil {
		if err := validator.ValidateMenuItemPosition(req.GetPosition()); err != nil {
			violations = append(violations, FieldViolation("position", err))
		}
	}

	if req.ParentItemId != nil {
		if err := validator.ValidateMenuItemParentItemId(req.GetParentItemId()); err != nil {
			violations = append(violations, FieldViolation("parent_item_id", err))
		}
	}

	if req.DescriptionPostId != nil {
		if err := validator.ValidatePostId(req.GetDescriptionPostId()); err != nil {
			violations = append(violations, FieldViolation("description_post_id", err))
		}
	}

	return violations
}