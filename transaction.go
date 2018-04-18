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
	"fmt"
	"path/filepath"

	"github.com/beeker1121/goque"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

const (
	txsPath  = "_txs"
	txPrefix = "tx_"
)

// transaction id -> local basePath
var transactions map[string]string

// Transaction represents an atomic operation composed of a row of sentences
type Transaction interface {
	AddStep(Step) error
	Commit()
	Rollback()
}

type transaction struct {
	ID    string `yaml:"id"`
	Queue *goque.Queue
}

func newTxID() (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		return "", errors.WithMessage(err, "Error generating transaction ID")
	}
	return fmt.Sprintf("%s%s", txPrefix, id), nil
}

func newTransaction() (string, error) {
	id, err := newTxID()
	if err != nil {
		return "", err
	}

	path := filepath.Join(txsPath, id)
	q, err := goque.OpenQueue(path)
	if err != nil {
		return "", errors.WithMessage(err, "Error touching transactions queue ")
	}
	defer q.Close()

	if transactions == nil {
		transactions = make(map[string]string)
	}
	transactions[id] = txsPath

	return id, nil
}

func removeTransaction(id string) error {
	basePath, ok := transactions[id]
	if !ok {
		return errors.Errorf("Transaction %s not found", id)
	}

	path := filepath.Join(basePath, id)
	q, err := goque.OpenQueue(path)
	if err != nil {
		return errors.Wrap(err, "Error opening transactions queue ")
	}
	defer q.Close()

	q.Drop()

	return nil
}

// Loads found transactions in disk. This method is valuable when wanted
// to execute rollback methods for a transaction that broke due to a
// service outage
func loadTransactions() error {
	ids, err := listDirs(txsPath)
	if err != nil {
		return err
	}

	if len(ids) > 0 && transactions == nil {
		transactions = make(map[string]string)
	}

	for _, id := range ids {
		transactions[id] = txsPath
	}

	return nil
}

func (tx *transaction) AddStep(step Step) error {

}
