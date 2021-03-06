// liquidhandling/lhrequest.Go: Part of the Antha language
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
package liquidhandling

import (
	"fmt"

	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/microArch/driver/liquidhandling"
)

// structure for defining a request to the liquid handler
type LHRequest struct {
	ID                    string
	BlockID               wtype.BlockID
	BlockName             string
	LHInstructions        map[string]*wtype.LHInstruction
	Input_solutions       map[string][]*wtype.LHComponent
	Plates                map[string]*wtype.LHPlate
	Tips                  []*wtype.LHTipbox
	InstructionSet        *liquidhandling.RobotInstructionSet
	Instructions          []liquidhandling.TerminalRobotInstruction
	InstructionText       string
	Input_assignments     map[string][]string
	Output_assignments    map[string][]string
	Input_plates          map[string]*wtype.LHPlate
	Output_plates         map[string]*wtype.LHPlate
	Input_platetypes      []*wtype.LHPlate
	Input_plate_order     []string
	Input_setup_weights   map[string]float64
	Output_platetypes     []*wtype.LHPlate
	Output_plate_order    []string
	Plate_lookup          map[string]string
	Stockconcs            map[string]wunit.Concentration
	Policies              *wtype.LHPolicyRuleSet
	Input_order           []string
	Output_order          []string
	OutputIteratorFactory func(*wtype.LHPlate) wtype.PlateIterator `json:"-"`
	InstructionChain      *IChain
	Input_vols_supplied   map[string]wunit.Volume
	Input_vols_required   map[string]wunit.Volume
	Input_vols_wanting    map[string]wunit.Volume
	TimeEstimate          float64
	CarryVolume           wunit.Volume
	InstructionSets       [][]*wtype.LHInstruction
	Evaps                 []wtype.VolumeCorrection
	Options               LHOptions
	NUserPlates           int
	Output_sort           bool
}

func (req *LHRequest) ConfigureYourself() error {
	// ensures input solutions is populated
	// once input plates are specified
	// more to happen later
	inputs := req.Input_solutions

	if inputs == nil {
		inputs = make(map[string][]*wtype.LHComponent)
	}

	// we need to make an exception of components which are used literally
	// i.e. anything used in a mix-in-place; these don't add to the general
	// store of anonymous components to be sampled from

	uniques := make(map[wtype.PlateLocation]*wtype.LHComponent, len(req.LHInstructions))

	for _, ins := range req.LHInstructions {
		// If provenance info is brought in here this becomes unsafe
		if ins.IsMixInPlace() && !ins.HasAnyParent() {
			if !ins.Components[0].PlateLocation().IsZero() {
				uniques[ins.Components[0].PlateLocation()] = ins.Components[0]
			} else {
				// this will be autoallocated
			}
		}
	}

	for _, v := range req.Input_plates {
		for _, w := range v.Wellcoords {
			if w.Empty() {
				continue
			}

			// special case for components treated literally

			cmp, ok := uniques[w.PlateLocation()]

			if ok {
				// unique components (where instances matter) are
				// identified using CNID()
				ar := inputs[cmp.CNID()]
				ar = append(ar, cmp)
				inputs[cmp.CNID()] = ar
			} else {
				// bulk components (where instances don't matter) are
				// identified using just CName
				c := w.Contents().Dup()
				vvvvvv := c.Volume()
				vvvvvv.Subtract(w.ResidualVolume())
				c.SetVolume(vvvvvv)
				ar := inputs[c.CName]
				ar = append(ar, c)
				inputs[c.CName] = ar
			}
		}
	}

	req.Input_solutions = inputs
	return nil
}

// this function checks requests so we can see early on whether or not they
// are going to cause problems
func ValidateLHRequest(rq *LHRequest) (bool, string) {
	if rq.Output_platetypes == nil || len(rq.Output_platetypes) == 0 {
		return false, "No output plate type specified"
	}

	if len(rq.Input_platetypes) == 0 {
		return false, "No input plate types specified"
	}

	if rq.Policies == nil {
		return false, "No policies specified"
	}

	return true, "OK"
}

func NewLHRequest() *LHRequest {
	var lhr LHRequest
	lhr.ID = wtype.GetUUID()
	lhr.LHInstructions = make(map[string]*wtype.LHInstruction)
	lhr.Input_solutions = make(map[string][]*wtype.LHComponent)
	lhr.Plates = make(map[string]*wtype.LHPlate)
	lhr.Tips = make([]*wtype.LHTipbox, 0, 1)
	lhr.Input_plates = make(map[string]*wtype.LHPlate)
	lhr.Input_platetypes = make([]*wtype.LHPlate, 0, 2)
	lhr.Input_setup_weights = make(map[string]float64)
	lhr.Output_plates = make(map[string]*wtype.LHPlate)
	lhr.Output_plate_order = make([]string, 0, 1)
	lhr.Input_plate_order = make([]string, 0, 1)
	lhr.Plate_lookup = make(map[string]string)
	lhr.Stockconcs = make(map[string]wunit.Concentration)
	lhr.Input_order = make([]string, 0)
	lhr.Output_order = make([]string, 0)
	lhr.OutputIteratorFactory = wtype.NewOneTimeColumnWiseIterator
	lhr.Output_assignments = make(map[string][]string)
	lhr.Input_assignments = make(map[string][]string)
	lhr.InstructionSet = liquidhandling.NewRobotInstructionSet(nil)
	lhr.InstructionText = ""
	lhr.Input_vols_required = make(map[string]wunit.Volume)
	lhr.Input_vols_supplied = make(map[string]wunit.Volume)
	lhr.Input_vols_wanting = make(map[string]wunit.Volume)
	lhr.CarryVolume = wunit.NewVolume(0.5, "ul")
	lhr.Input_setup_weights["MAX_N_PLATES"] = 2
	lhr.Input_setup_weights["MAX_N_WELLS"] = 96
	lhr.Input_setup_weights["RESIDUAL_VOLUME_WEIGHT"] = 1.0
	lhr.Policies, _ = liquidhandling.GetLHPolicyForTest()
	lhr.Options = NewLHOptions()
	return &lhr
}

func (lhr *LHRequest) Add_instruction(ins *wtype.LHInstruction) {
	lhr.LHInstructions[ins.ID] = ins
}

func (lhr *LHRequest) NewComponentsAdded() bool {
	// run this after Plan to determine if anything
	// new was added to the inputs

	return len(lhr.Input_vols_wanting) != 0
}

func (lhr *LHRequest) AddUserPlate(p *wtype.LHPlate) {
	// impose sanity

	if p.PlateName == "" {
		p.PlateName = fmt.Sprintf("User_plate_%d", lhr.NUserPlates+1)
		lhr.NUserPlates += 1
	}

	p.MarkNonEmptyWellsUserAllocated()

	lhr.Input_plates[p.ID] = p
}

func (lhr *LHRequest) UseLegacyVolume() bool {
	// magically create extra volumes for intermediates?
	return lhr.Options.LegacyVolume
}

type LHPolicyManager struct {
	SystemPolicies *wtype.LHPolicyRuleSet
	UserPolicies   *wtype.LHPolicyRuleSet
}

func (mgr *LHPolicyManager) MergePolicies(protocolpolicies *wtype.LHPolicyRuleSet) *wtype.LHPolicyRuleSet {
	ret := wtype.CloneLHPolicyRuleSet(mgr.SystemPolicies)

	// things coming in take precedence over things already there
	ret.MergeWith(mgr.UserPolicies)
	ret.MergeWith(protocolpolicies)

	return ret
}
