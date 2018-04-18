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
	"testing"

	"github.com/beeker1121/goque"
	check "gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type queueSuite struct{}

var _ = check.Suite(&queueSuite{})

func (qs *queueSuite) SetUpTest(c *check.C) {
}

func (qs *queueSuite) TestQueue(c *check.C) {

	s, err := goque.OpenStack("hey/data_dir")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	// Push an item onto the stack.
	item, err := s.Push([]byte("item value"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item.ID)         // 1
	fmt.Println(item.Key)        // [0 0 0 0 0 0 0 1]
	fmt.Println(item.Value)      // [105 116 101 109 32 118 97 108 117 101]
	fmt.Println(item.ToString()) // item value

	// Change the item value in the stack.
	/* item, err = s.Update(item.ID, []byte("new item value"))
	if err != nil {
		fmt.Println(err)
		return
	} */

	fmt.Println(item.ToString()) // new item value

	// Pop an item off the stack.
	/* popItem, err := s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	} */

	//fmt.Println(popItem.ToString()) // new item value
	fmt.Printf("LEN:%v", s.Length())

	// Delete the stack and its database.
	//s.Drop()
}
