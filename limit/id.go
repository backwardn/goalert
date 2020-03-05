package limit

import "github.com/target/goalert/validation/validate"

// ID represents the identifier for a given system limit.
type ID string

// IDs of configurable limits.
const (
	NotificationRulesPerUser ID = "notification_rules_per_user"
	ContactMethodsPerUser    ID = "contact_methods_per_user"
	// Maximum steps on a single escalation policy step.
	EPStepsPerPolicy ID = "ep_steps_per_policy"
	// Maximum actions on a single escalation policy step.
	EPActionsPerStep        ID = "ep_actions_per_step"
	ParticipantsPerRotation ID = "participants_per_rotation"
	// Pertains to all rules for all assignments/targets.
	RulesPerSchedule          ID = "rules_per_schedule"
	IntegrationKeysPerService ID = "integration_keys_per_service"
	// Only affects newly created alerts, but not re-escalated ones.
	UnackedAlertsPerService     ID = "unacked_alerts_per_service"
	TargetsPerSchedule          ID = "targets_per_schedule"
	HeartbeatMonitorsPerService ID = "heartbeat_monitors_per_service"
	// Only limits future overrides (e.g. end in the future).
	UserOverridesPerSchedule ID = "user_overrides_per_schedule"
	CalendarSubscriptionsPerUser ID = "calendar_subscriptions_per_user"
)

// Valid returns nil if a given ID is valid, a validation error is returned otherwise.
func (id ID) Valid() error {
	return validate.OneOf("LimitID", id,
		NotificationRulesPerUser,
		ContactMethodsPerUser,
		EPStepsPerPolicy,
		EPActionsPerStep,
		ParticipantsPerRotation,
		RulesPerSchedule,
		IntegrationKeysPerService,
		UnackedAlertsPerService,
		TargetsPerSchedule,
		HeartbeatMonitorsPerService,
		UserOverridesPerSchedule,
		CalendarSubscriptionsPerUser,
	)
}
