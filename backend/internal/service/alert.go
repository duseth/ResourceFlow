package service

import (
	"context"
	"fmt"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
	"github.com/google/uuid"
)

type AlertService struct {
	alertRepo repository.AlertRepository
}

func NewAlertService(ar repository.AlertRepository) *AlertService {
	return &AlertService{
		alertRepo: ar,
	}
}

func (s *AlertService) CreateAlert(ctx context.Context, alert *models.Alert) error {
	alert.ID = uuid.New().String()
	alert.CreatedAt = time.Now()
	alert.Status = models.AlertStatusActive
	return s.alertRepo.Create(ctx, alert)
}

func (s *AlertService) ResolveAlert(ctx context.Context, alertID string) error {
	alert, err := s.alertRepo.GetByID(ctx, alertID)
	if err != nil {
		return fmt.Errorf("failed to get alert: %v", err)
	}

	if alert.Status == models.AlertStatusResolved {
		return fmt.Errorf("alert already resolved")
	}

	now := time.Now()
	alert.Status = models.AlertStatusResolved
	alert.ResolvedAt = &now

	return s.alertRepo.Update(ctx, alert)
}

func (s *AlertService) AcknowledgeAlert(ctx context.Context, alertID string) error {
	alert, err := s.alertRepo.GetByID(ctx, alertID)
	if err != nil {
		return fmt.Errorf("failed to get alert: %v", err)
	}

	if alert.Status != models.AlertStatusActive {
		return fmt.Errorf("can only acknowledge active alerts")
	}

	alert.Status = models.AlertStatusAcknowledged
	return s.alertRepo.Update(ctx, alert)
}

func (s *AlertService) GetActiveAlerts(ctx context.Context) ([]*models.Alert, error) {
	return s.alertRepo.GetActive(ctx)
}

func (s *AlertService) GetServerAlerts(ctx context.Context, serverID string) ([]*models.Alert, error) {
	return s.alertRepo.GetByServer(ctx, serverID)
}

func (s *AlertService) GetRules(ctx context.Context, serverID string) ([]*models.AlertRule, error) {
	return s.alertRepo.GetRules(ctx, serverID)
}

func (s *AlertService) CreateRule(ctx context.Context, rule *models.AlertRule) error {
	rule.ID = uuid.New().String()
	return s.alertRepo.CreateRule(ctx, rule)
}

func (s *AlertService) UpdateRule(ctx context.Context, rule *models.AlertRule) error {
	return s.alertRepo.UpdateRule(ctx, rule)
}

func (s *AlertService) DeleteRule(ctx context.Context, ruleID string) error {
	return s.alertRepo.DeleteRule(ctx, ruleID)
}
