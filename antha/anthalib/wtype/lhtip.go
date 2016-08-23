// liquidhandling/lhtypes.Go: Part of the Antha language
// Copyright (C) 2014 the Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@antha-lang.Org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

// defines types for dealing with liquid handling requests
package wtype

import (
	"fmt"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
)

//TODO add extra properties, i.e. filter
//     remove BBox once shape implements LHObject
type LHTip struct {
	ID     string
	Type   string
	Mnfr   string
	Dirty  bool
	MaxVol wunit.Volume
	MinVol wunit.Volume
	shape  *Shape
	bounds BBox
	parent LHObject
}

/*
	ID          string
	Name        string
	Minvol      wunit.Volume
	Maxvol      wunit.Volume
	Minspd      wunit.FlowRate
	Maxspd      wunit.FlowRate
	Multi       int
	Independent bool
	Orientation int
	Head        int
*/

//@implement Named
func (self *LHTip) GetName() string {
	if self == nil {
		return "<nil>"
	}
	if addr, ok := self.parent.(Addressable); ok {
		pos := self.GetPosition().Add(self.GetSize().Multiply(0.5))
		wc, _ := addr.CoordsToWellCoords(pos)
		return fmt.Sprintf("%s@%s", wc.FormatA1(), NameOf(self.parent))
	}
	return fmt.Sprintf("%s_%s", self.Mnfr, self.Type)
}

//@implement Typed
func (self *LHTip) GetType() string {
	if self == nil {
		return "<nil>"
	}
	return self.Type
}

//@implement Classy
func (self *LHTip) GetClass() string {
	return "tip"
}

//@implement LHObject
func (self *LHTip) GetPosition() Coordinates {
	return OriginOf(self).Add(self.bounds.GetPosition())
}

//@implement LHObject
func (self *LHTip) GetSize() Coordinates {
	return self.bounds.GetSize()
}

//@implement LHObject
func (self *LHTip) GetBoxIntersections(box BBox) []LHObject {
	box.SetPosition(box.GetPosition().Subtract(OriginOf(self)))
	if self.bounds.IntersectsBox(box) {
		return []LHObject{self}
	}
	return nil
}

//@implement LHObject
func (self *LHTip) GetPointIntersections(point Coordinates) []LHObject {
	point = point.Subtract(point)
	//TODO more accurate intersection detection with Shape
	if self.bounds.IntersectsPoint(point) {
		return []LHObject{self}
	}
	return nil
}

//@implement LHObject
func (self *LHTip) SetOffset(point Coordinates) error {
	self.bounds.SetPosition(point)
	return nil
}

//@implement LHObject
func (self *LHTip) SetParent(o LHObject) error {
	//parent should be LHTipbox (should accept LHAdaptor, but it doesn't implement LHObject yet)
	if _, ok := o.(*LHTipbox); ok {
		self.parent = o
		return nil
	}
	return fmt.Errorf("Cannot set %s \"%s\" as parent of tip", ClassOf(o), NameOf(o))
}

//@implement LHObject
func (self *LHTip) GetParent() LHObject {
	return self.parent
}

func (tip *LHTip) GetParams() *LHChannelParameter {
	// be safe
	if tip.IsNil() {
		return nil
	}

	lhcp := LHChannelParameter{Name: tip.Type + "Params", Minvol: tip.MinVol, Maxvol: tip.MaxVol, Multi: 1, Independent: false, Orientation: LHVChannel}
	return &lhcp
}

func (tip *LHTip) IsNil() bool {
	if tip == nil || tip.Type == "" || tip.MaxVol.IsZero() || tip.MinVol.IsZero() {
		return true
	}
	return false
}

func (tip *LHTip) Dup() *LHTip {
	t := NewLHTip(tip.Mnfr, tip.Type, tip.MinVol.RawValue(), tip.MaxVol.RawValue(), tip.MinVol.Unit().PrefixedSymbol(), tip.shape)
	t.Dirty = tip.Dirty
	return t
}

func NewLHTip(mfr, ttype string, minvol, maxvol float64, volunit string, shape *Shape) *LHTip {
	lht := LHTip{
		GetUUID(),
		ttype,
		mfr,
		false, //dirty
		wunit.NewVolume(maxvol, volunit),
		wunit.NewVolume(minvol, volunit),
		shape,
		BBox{Coordinates{}, Coordinates{
			shape.Height().ConvertToString("mm"), //not a mistake, Shape currently has height&width as
			shape.Width().ConvertToString("mm"),  // XY coordinates and Depth as Z
			shape.Depth().ConvertToString("mm"),
		}},
		nil,
	}
	return &lht
}

func CopyTip(tt LHTip) *LHTip {
	return &tt
}
