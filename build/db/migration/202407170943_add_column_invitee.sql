-- +goose Up
ALTER TABLE invitees ADD join_flag boolean NOT NULL DEFAULT TRUE;