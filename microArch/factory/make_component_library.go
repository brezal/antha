// /anthalib/factory/make_component_library.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
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
// contact license@antha-lang.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

package factory

import (
	"sort"

	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/image"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
)

type alreadyAdded struct {
	Name string
}

func (a *alreadyAdded) Error() string {
	return "component " + a.Name + " already added"
}

type notFound struct {
	Name string
}

func (a *notFound) Error() string {
	return "component " + a.Name + " not found"
}

func makeComponentLibrary() (map[string]*wtype.LHComponent, error) {
	var components []*wtype.LHComponent

	add := func(name string, typ wtype.LiquidType, smax float64, conc wunit.Concentration) {
		c := wtype.NewLHComponent()
		c.CName = name
		c.Type = typ
		c.Smax = smax
		components = append(components, c)
	}

	var defaultConc wunit.Concentration
	x2 := wunit.NewConcentration(2, "X")
	x10 := wunit.NewConcentration(10, "X")
	x100 := wunit.NewConcentration(100, "X")

	gPerL01 := wunit.NewConcentration(0.1, "g/L")
	gPerL1 := wunit.NewConcentration(1, "g/L")
	gPerL10 := wunit.NewConcentration(10, "g/L")
	gPerL100 := wunit.NewConcentration(100, "g/L")

	add("water", wtype.LTWater, 9999, defaultConc)
	add("multiwater", wtype.LTMultiWater, 9999, defaultConc)
	add("DNA_ladder", wtype.LTload, 9999, defaultConc)
	add("LoadingDye", wtype.LTWater, 9999, defaultConc)
	add("Overlay", wtype.LTPLATEOUT, 9999, defaultConc)
	add("PEG", wtype.LTPEG, 9999, defaultConc)
	add("protoplasts", wtype.LTProtoplasts, 9999, defaultConc)
	add("fluorescein", wtype.LTWater, 9999, defaultConc)
	add("ethanol", wtype.LTWater, 9999, defaultConc)
	add("whiteFabricDye", wtype.LTGlycerol, 9999, defaultConc)
	add("blackFabricDye", wtype.LTGlycerol, 9999, defaultConc)
	add("Some component in factory", wtype.LTWater, 9999, defaultConc)
	add("neb5compcells", wtype.LTCulture, 1.0, defaultConc)
	add("mediaonculture", wtype.LTNeedToMix, 1.0, defaultConc)
	add("10x_M9Salts", wtype.LTWater, 9999, x10)
	add("100x_MEMVitamins", wtype.LTWater, 9999, x100)
	add("Yeast extract", wtype.LTWater, 9999, defaultConc)
	add("Tryptone", wtype.LTWater, 9999, defaultConc)
	add("Glycerol", wtype.LTPostMix, 9999, defaultConc)
	add("culture", wtype.LTCulture, 9999, defaultConc)
	// the pubchem name for tartrazine
	add("Acid yellow 23", wtype.LTWater, 9999, defaultConc)
	add("tartrazine", wtype.LTWater, 9999, defaultConc)
	add("tartrazinePostMix", wtype.LTPostMix, 9999, defaultConc)
	add("tartrazineNeedtoMix", wtype.LTNeedToMix, 9999, defaultConc)
	add("tartrazine_DNA", wtype.LTDNA, 9999, defaultConc)
	add("tartrazine_Glycerol", wtype.LTGlycerol, 9999, defaultConc)
	add("Yellow_ink", wtype.LTPAINT, 9999, defaultConc)
	add("Cyan", wtype.LTPAINT, 9999, defaultConc)
	add("Magenta", wtype.LTPAINT, 9999, defaultConc)
	add("transparent", wtype.LTWater, 9999, defaultConc)
	add("Black", wtype.LTPAINT, 9999, defaultConc)
	add("Paint", wtype.LTPostMix, 9999, defaultConc)
	add("yellow", wtype.LTWater, 9999, defaultConc)
	add("blue", wtype.LTWater, 9999, defaultConc)
	add("darkblue", wtype.LTWater, 9999, defaultConc)
	add("grey", wtype.LTWater, 9999, defaultConc)
	add("green", wtype.LTWater, 9999, defaultConc)
	add("red", wtype.LTWater, 9999, defaultConc)
	add("white", wtype.LTWater, 9999, defaultConc)
	add("black", wtype.LTWater, 9999, defaultConc)
	add("purple", wtype.LTWater, 9999, defaultConc)
	add("pink", wtype.LTWater, 9999, defaultConc)
	add("orange", wtype.LTWater, 9999, defaultConc)
	add("DNAsolution", wtype.LTDNA, 1.0, defaultConc)
	add("1kb DNA Ladder", wtype.LTDNA, 10.0, defaultConc)
	add("restrictionenzyme", wtype.LTGlycerol, 1.0, defaultConc)
	add("bsa", wtype.LTWater, 100, defaultConc)
	add("dna_part", wtype.LTDNA, 1.0, defaultConc)
	add("dna", wtype.LTDNA, 1.0, defaultConc)
	add("SapI", wtype.LTGlycerol, 1.0, defaultConc)
	add("BsaI", wtype.LTGlycerol, 1.0, defaultConc)
	add("T4Ligase", wtype.LTGlycerol, 1.0, defaultConc)
	add("EcoRI", wtype.LTGlycerol, 1.0, defaultConc)
	add("EnzMastermix: 1/2 SapI; 1/2 T4 Ligase", wtype.LTGlycerol, 1.0, defaultConc)
	add("TypeIIsbuffer: 2/11 10xCutsmart; 1/11 1mM ATP; 8/11 Water", wtype.LTWater, 9999, defaultConc)
	add("CutsmartBuffer", wtype.LTWater, 1.0, defaultConc)
	add("ATP", wtype.LTWater, 5.0, defaultConc)
	add("mastermix_sapI", wtype.LTWater, 1.0, defaultConc)
	add("mastermix_bsaI", wtype.LTWater, 1.0, defaultConc)
	add("standard_cloning_vector_mark_1", wtype.LTDNA, 1.0, defaultConc)
	add("Q5Polymerase", wtype.LTGlycerol, 1.0, defaultConc)
	add("GoTaq_ green 2x mastermix", wtype.LTGlycerol, 9999.0, x2)
	add("DMSO", wtype.LTWater, 1.0, defaultConc)
	add("pET_GFP", wtype.LTWater, 1.0, defaultConc)
	add("HC", wtype.LTWater, 1.0, defaultConc)
	add("GCenhancer", wtype.LTWater, 9999.0, defaultConc)
	add("Q5buffer", wtype.LTWater, 1.0, defaultConc)
	add("Q5mastermix", wtype.LTWater, 1.0, defaultConc)
	add("PrimerFw", wtype.LTDNA, 1.0, defaultConc)
	add("PrimerRev", wtype.LTDNA, 1.0, defaultConc)
	add("template_part", wtype.LTDNA, 1.0, defaultConc)
	add("DNTPs", wtype.LTWater, 1.0, defaultConc)
	add("ProteinMarker", wtype.LTProtein, 1.0, defaultConc)
	add("ProteinFraction", wtype.LTProtein, 1.0, defaultConc)
	add("EColiLysate", wtype.LTProtein, 1.0, defaultConc)
	add("SDSbuffer", wtype.LTDetergent, 1.0, defaultConc)
	add("Load", wtype.LTload, 1.0, defaultConc)
	add("LB", wtype.LTWater, 1.0, defaultConc)
	add("TB", wtype.LTWater, 1.0, defaultConc)
	add("Kanamycin", wtype.LTWater, 1.0, defaultConc)
	add("Glucose", wtype.LTPostMix, 1.0, defaultConc)
	add("IPTG", wtype.LTPostMix, 1.0, defaultConc)
	add("Lactose", wtype.LTWater, 1.0, defaultConc)
	add("colony", wtype.LTCOLONY, 1.0, defaultConc)
	add("LB_autoinduction_Amp", wtype.LTWater, 1.0, defaultConc)
	add("LB_Kan", wtype.LTWater, 1.0, defaultConc)
	add("Apramycin", wtype.LTWater, 1.0, defaultConc)
	add("Agar", wtype.LTWater, 1.0, defaultConc)
	add("X-glc", wtype.LTWater, 1.0, defaultConc)
	add("X-Glucuro", wtype.LTWater, 1.0, defaultConc)
	add("BaseGrowthMedium", wtype.LTWater, 1.0, defaultConc)
	add("SterileWater", wtype.LTWater, 1.0, defaultConc)
	add("100mMPhosphate", wtype.LTWater, 1.0, wunit.NewConcentration(100, "mM"))
	add("100g/LGlucose", wtype.LTWater, 1.0, gPerL100)
	add("10g/LGlucose", wtype.LTWater, 1.0, gPerL10)
	add("1g/LGlucose", wtype.LTWater, 1.0, gPerL1)
	add("0.1g/Lglucose", wtype.LTWater, 1.0, gPerL01)
	add("100g/Lglycerol", wtype.LTWater, 1.0, gPerL100)
	add("10g/Lglycerol", wtype.LTWater, 1.0, gPerL10)
	add("1g/Lglycerol", wtype.LTWater, 1.0, gPerL1)
	add("0.1g/Lglycerol", wtype.LTWater, 1.0, gPerL01)
	add("100g/Lpeptone", wtype.LTWater, 1.0, gPerL100)
	add("100g/LYeastExtract", wtype.LTWater, 1.0, gPerL100)
	add("10g/LYeastExtract", wtype.LTWater, 1.0, gPerL10)
	add("100g/L Glucose", wtype.LTWater, 1.0, gPerL100)
	add("10g/L Glucose", wtype.LTWater, 1.0, gPerL10)
	add("1g/L Glucose", wtype.LTWater, 1.0, gPerL1)
	add("0.1g/L glucose", wtype.LTWater, 1.0, gPerL01)
	add("100g/L glycerol", wtype.LTWater, 1.0, gPerL100)
	add("10g/L glycerol", wtype.LTWater, 1.0, gPerL10)
	add("1g/L glycerol", wtype.LTWater, 1.0, gPerL1)
	add("0.1g/L glycerol", wtype.LTWater, 1.0, gPerL01)
	add("100g/L peptone", wtype.LTWater, 1.0, gPerL100)
	add("100g/L YeastExtract", wtype.LTWater, 1.0, gPerL100)
	add("10g/L YeastExtract", wtype.LTWater, 1.0, gPerL10)
	add("1000ng/ml ATC", wtype.LTWater, 1.0, wunit.NewConcentration(1000, "ng/ml"))
	add("ATC", wtype.LTWater, 1.0, defaultConc)
	add("C6", wtype.LTWater, 1.0, defaultConc)
	add("C12", wtype.LTWater, 1.0, defaultConc)
	add("250uM C6", wtype.LTWater, 1.0, wunit.NewConcentration(250, "uM"))
	add("25uM C6", wtype.LTWater, 1.0, wunit.NewConcentration(25, "uM"))
	add("2.5uM C6", wtype.LTWater, 1.0, wunit.NewConcentration(2.5, "uM"))
	add("0.25uM C6", wtype.LTWater, 1.0, wunit.NewConcentration(0.25, "uM"))
	add("0.025uM C6", wtype.LTWater, 1.0, wunit.NewConcentration(0.025, "uM"))
	add("250g/L C6", wtype.LTWater, 1.0, wunit.NewConcentration(250, "g/L"))
	add("25g/L C6", wtype.LTWater, 1.0, wunit.NewConcentration(25, "g/L"))
	add("2.5g/L C6", wtype.LTWater, 1.0, wunit.NewConcentration(2.5, "g/L"))
	add("0.25g/L C6", wtype.LTWater, 1.0, wunit.NewConcentration(0.25, "g/L"))
	add("0.025g/L C6", wtype.LTWater, 1.0, wunit.NewConcentration(0.025, "g/L"))
	add("IPTG 1mM", wtype.LTWater, 1.0, wunit.NewConcentration(1, "mM"))
	add("Glucose 100g/L", wtype.LTWater, 1.0, gPerL100)
	add("Glucose 1g/L", wtype.LTWater, 1.0, gPerL1)
	add("Glycerol 100g/L", wtype.LTWater, 1.0, gPerL100)
	add("M9", wtype.LTWater, 1.0, defaultConc)
	add("HYYest412", wtype.LTNSrc, 1.0, defaultConc)
	add("HYYest503", wtype.LTNSrc, 1.0, defaultConc)
	add("HYYest504", wtype.LTNSrc, 1.0, defaultConc)
	add("PeaPeptone", wtype.LTNSrc, 1.0, defaultConc)
	add("WheatPeptone", wtype.LTNSrc, 1.0, defaultConc)
	add("VegPeptone", wtype.LTNSrc, 1.0, defaultConc)
	add("SoyPeptone", wtype.LTNSrc, 1.0, defaultConc)
	add("VegExtract", wtype.LTNSrc, 1.0, defaultConc)
	add("CSL", wtype.LTNSrc, 1.0, defaultConc)
	add("NH42SO4", wtype.LTNSrc, 1.0, defaultConc)
	add("Gluc", wtype.LTNSrc, 1.0, defaultConc)
	add("Suc", wtype.LTNSrc, 1.0, defaultConc)
	add("Fruc", wtype.LTNSrc, 1.0, defaultConc)
	add("Malt", wtype.LTNSrc, 1.0, defaultConc)
	add("water2", wtype.LTNSrc, 1.0, defaultConc)

	// protein paintbox
	for _, value := range image.ProteinPaintboxmap {
		add(value, wtype.LTPostMix, 1.0, defaultConc)
	}

	cmap := make(map[string]*wtype.LHComponent)
	for _, c := range components {
		if _, seen := cmap[c.CName]; seen {
			return nil, &alreadyAdded{Name: c.CName}
		}
		cmap[c.CName] = c
	}

	return cmap, nil
}

type componentLibrary struct {
	lib map[string]*wtype.LHComponent
}

var defaultComponentLibrary *componentLibrary

func init() {
	lib, err := makeComponentLibrary()
	if err != nil {
		panic(err)
	}

	defaultComponentLibrary = &componentLibrary{
		lib: lib,
	}
}

func (i *componentLibrary) GetComponentByType(typ string) *wtype.LHComponent {
	c, ok := i.lib[typ]
	if !ok {
		panic(&notFound{Name: typ})
	}
	return c.Dup()
}

func (i *componentLibrary) GetComponent(typ string) (*wtype.LHComponent, error) {
	c, ok := i.lib[typ]
	if !ok {
		return nil, &notFound{Name: typ}
	}
	return c.Dup(), nil
}

func ComponentInFactory(typ string) bool {
	_, ok := defaultComponentLibrary.lib[typ]
	return ok
}

func GetComponents() []*wtype.LHComponent {
	var comps []*wtype.LHComponent
	for _, c := range defaultComponentLibrary.lib {
		comps = append(comps, c)
	}

	return wtype.CopyComponentArray(comps)
}

// TODO: deprecate
func GetComponentList() []string {
	comps := GetComponents()
	var names []string
	for _, c := range comps {
		names = append(names, c.CName)
	}

	sort.Strings(names)

	return names
}

func GetComponentByType(typ string) *wtype.LHComponent {
	return defaultComponentLibrary.GetComponentByType(typ)
}
