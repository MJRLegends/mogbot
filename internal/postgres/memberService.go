package postgres

import (
	"log"

	"github.com/captainmog/mogbot/internal/mogbot"
	"github.com/jmoiron/sqlx"
)

type MemberService struct {
	*sqlx.DB
}

func (s *MemberService) AddMember(m mogbot.Member) error {
	tx, err := s.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %s", err)
		return err
	}
	defer tx.Rollback()
	if _, err = tx.Exec(
		"insert into member (id) values ($1)",
		m.ID,
	); err != nil {
		log.Printf("Error inserting member '%s': %s", m.ID, err)
		return err
	}
	for _, r := range m.Roles {
		if _, err = tx.Exec(
			"insert into member_role (member_id, role_id) values ($1, $2) on conflict do nothing",
			m.ID,
			r.ID,
		); err != nil {
			log.Printf("Error inserting role '%s' for member '%s': %s", m.ID, err)
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		log.Printf("Error committing member insert: %s", err)
		return err
	}
	log.Printf("insert into table member with id=%s", m.ID)
	return nil
}

func (s *MemberService) GetMember(userID string) (*mogbot.Member, error) {
	row := s.QueryRowx(
		"select * from member where id=$1",
		userID,
	)
	var m mogbot.Member
	err := row.StructScan(&m)
	if err != nil {
		log.Printf("Error scanning member struct from row: %s", err)
		return nil, err
	}
	rows, err := s.Queryx(
		"select * from member_role where member_id=$1",
		userID,
	)
	if err != nil {
		log.Printf("Error querying member_role struct from rows: %s", err)
		return nil, err
	}
	var roles []mogbot.Role
	for rows.Next() {
		var r mogbot.Role
		if err := rows.StructScan(&r); err != nil {
			log.Printf("Error scanning roles for member '%s': %s", m.ID, err)
			return nil, err
		}
		roles = append(roles, r)
	}
	return &m, nil
}

func (s *MemberService) GetAllMembers() ([]*mogbot.Member, error) {
	rows, err := s.Queryx("select * from member")
	if err != nil {
		log.Printf("Error querying member: %s", err)
		return nil, err
	}
	var members []*mogbot.Member
	for rows.Next() {
		var m mogbot.Member
		err = rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		members = append(members, &m)
	}
	return members, nil
}

func (s *MemberService) UpdateMember(nm *mogbot.Member) error {
	tx, err := s.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %s", err)
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(
		"delete from member_role where member_id=$1",
		nm.ID,
	); err != nil {
		log.Printf("Error updating member '%s': %s", nm.ID, err)
		return err
	}
	for _, r := range nm.Roles {
		if _, err = tx.Exec(
			"insert into member_role (member_id, role_id) values ($1, $2) on conflict do nothing",
			nm.ID,
			r.ID,
		); err != nil {
			log.Printf("Error updating member '%s': %s", nm.ID, err)
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		log.Printf("Error committing member update: %s", err)
		return err
	}
	log.Printf("update table member where id=%s", nm.ID)
	return nil
}

func (s *MemberService) RemoveMember(userID string) error {
	if _, err := s.Exec(
		"delete from member where id=$1",
		userID,
	); err != nil {
		return err
	}
	log.Printf("delete from table member where id=%s", userID)
	return nil
}
