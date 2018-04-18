// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2018 Roberto Mier Escandon <rmescandon@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package trans

import (
	"github.com/beeker1121/goque"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

const (
	txsStoreName = "txs"
)

// Store interface to a permanent storage system. Holds the transactions
// currently being executed until Commit or Rollback is called. AFter that
// they are removed
type Store interface {
	Save(Step) error
	Pending() ([]Step, error)
	Free(Transaction) error
}

func (t *txsStore) newTransaction() (string, error) {
	q, err := goque.OpenQueue(txsStoreName)
	if err != nil {
		return "", errors.Wrap(err, "Error opening transactions queue ")
	}
	defer q.Close()

	id, err := shortid.Generate()
	if err != nil {
		return "", errors.Wrap(err, "Error generating transaction ID")
	}

	_, err := q.Enqueue([]byte(id))
	if err != nil {
		return "", errors.Wrap(err, "Error adding transaction to queue")
	}
	return id, nil
}

func (t *txsStore) removeTransaction(id string) error {
	q, err := goque.OpenQueue(txsStoreName)
	if err != nil {
		return "", errors.Wrap(err, "Error opening transactions queue ")
	}
	defer q.Close()

	_, err := q.Dequeue()
}

func (s *stack) Save(step Step) error {
	// TODO see this queues implementation https://github.com/beeker1121/goque
	return nil
}

func (s *stack) Pending() ([]Step, error) {
	// TODO IMPLEMENT
	return nil, nil
}

func (s *stack) Free(tx Transaction) error {
	// TODO IMPLEMENT
	return nil
}
