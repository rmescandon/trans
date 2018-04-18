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

// Action is a function to be executed as a step or as a rollback step into
// a transaction
type Action func(interface{}) error

// Step represents a single action into a transaction. It is composed
// of a Do() and an Undo(). Second won't be executed unless a later sentence
// in same transaction fails to execute its Do()
type Step interface {
	Do(Action) error
	Undo(Action) error
}
