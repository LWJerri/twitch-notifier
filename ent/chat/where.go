// Code generated by ent, DO NOT EDIT.

package chat

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/satont/twitch-notifier/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldID, id))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldChatID, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldChatID, v))
}

// ChatIDContains applies the Contains predicate on the "chat_id" field.
func ChatIDContains(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContains(FieldChatID, v))
}

// ChatIDHasPrefix applies the HasPrefix predicate on the "chat_id" field.
func ChatIDHasPrefix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasPrefix(FieldChatID, v))
}

// ChatIDHasSuffix applies the HasSuffix predicate on the "chat_id" field.
func ChatIDHasSuffix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasSuffix(FieldChatID, v))
}

// ChatIDEqualFold applies the EqualFold predicate on the "chat_id" field.
func ChatIDEqualFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEqualFold(FieldChatID, v))
}

// ChatIDContainsFold applies the ContainsFold predicate on the "chat_id" field.
func ChatIDContainsFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContainsFold(FieldChatID, v))
}

// ServiceEQ applies the EQ predicate on the "service" field.
func ServiceEQ(v Service) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldService, v))
}

// ServiceNEQ applies the NEQ predicate on the "service" field.
func ServiceNEQ(v Service) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldService, v))
}

// ServiceIn applies the In predicate on the "service" field.
func ServiceIn(vs ...Service) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldService, vs...))
}

// ServiceNotIn applies the NotIn predicate on the "service" field.
func ServiceNotIn(vs ...Service) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldService, vs...))
}

// HasSettings applies the HasEdge predicate on the "settings" edge.
func HasSettings() predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SettingsTable, SettingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSettingsWith applies the HasEdge predicate on the "settings" edge with a given conditions (other predicates).
func HasSettingsWith(preds ...predicate.ChatSettings) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SettingsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SettingsTable, SettingsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFollows applies the HasEdge predicate on the "follows" edge.
func HasFollows() predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FollowsTable, FollowsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFollowsWith applies the HasEdge predicate on the "follows" edge with a given conditions (other predicates).
func HasFollowsWith(preds ...predicate.Follow) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FollowsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FollowsTable, FollowsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Chat) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		p(s.Not())
	})
}