// Copyright 2014 Mahmud Ridwan. All rights reserved.

package bktree

import (
	"math/rand"
	"testing"

	"github.com/agnivade/levenshtein"
)

func TestFindSm(t *testing.T) {
	testFind(t, dictSm)
}

func TestFindLg(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
		return
	}

	testFind(t, dictLg)
}

func BenchmarkFindSm(b *testing.B) {
	benchmarkFind(b, dictSm)
}

func BenchmarkFindLg(b *testing.B) {
	if testing.Short() {
		b.SkipNow()
		return
	}

	benchmarkFind(b, dictLg)
}

func testFind(t *testing.T, dict []string) {
	bk := New(levenshtein.ComputeDistance)

	for _, w := range dict {
		bk.Add(w)
	}

	for _, w := range dict {
		for k := 0; k < 5; k++ {
			m := mess(w, k)

			r := bk.Find(m, k)
			if len(r) == 0 {
				t.FailNow()
			}
			for _, found_w := range r {
				if levenshtein.ComputeDistance(m, found_w) > k {
					t.FailNow()
				}
			}
		}
	}
}

func benchmarkFind(b *testing.B, dict []string) {
	bk := New(levenshtein.ComputeDistance)

	for _, w := range dict {
		bk.Add(w)
	}

	s := []string{}
	for i := 0; i < b.N; i++ {
		s = append(s, pick(dict))
	}

	for i := 0; i < b.N; i++ {
		bk.Find(s[i], 4)
	}
}

// Picks a random word from the given dictionary
func pick(dict []string) string {
	return dict[rand.Intn(len(dict))]
}

// Messes up the string to end up with a Levenshtein distance of at most k from the original
func mess(w string, k int) string {
	for ; k > 0 && len(w) > 1; k-- {
		switch rand.Intn(4) {
		case 0:
			i := rand.Intn(len(w))
			w = w[:i] + char() + w[i:]
		case 1:
			i := rand.Intn(len(w))
			w = w[:i] + w[i+1:]
		case 2:
			i := rand.Intn(len(w))
			w = w[:i] + char() + w[i+1:]
		}
	}
	return w
}

// Returns a random lower-case alphabet
func char() string {
	return string("abcdefghijklmnopqrstuvwxyz"[rand.Intn(26)])
}

var (
	dictSm = []string{"assembly", "commuter", "commuters", "shrivel", "bolivia", "examining", "azimuthal", "dadaist", "psychopathic", "truss", "icicle", "gadwall", "superlatives", "opened", "unloading", "timid", "intervene", "occurrent", "precedences", "thermofax", "consummation", "russell", "intensified", "eradication", "hunches", "spade", "restraining", "discussant", "goodrich", "cathodes", "layers", "bookcase", "volleyballs", "sandals", "context", "validated", "felice", "munched", "vowels", "bose", "columnizing", "lodgepole", "affectation", "mulatto", "scepters", "coated", "villas", "armys", "dictum", "excitable", "radars", "dimmers", "underflowed", "unquoted", "seedbed", "ultracentrifuge", "syllables", "prospections", "employer", "pianist", "perplexing", "convocation", "ambient", "aries", "scissor", "delving", "vagarys", "rosenberg", "oyster", "hasty", "inquire", "rogue", "intensive", "mister", "tapis", "discovered", "knightly", "circumstances", "scientist", "percent", "homebound", "waxes", "partake", "rust", "goatees", "aggrieved", "kaufman", "tuberculin", "separators", "methodic", "polygon", "fecund", "francoise", "warfare", "disconcerting", "fuzziness", "drain", "smelter", "focussed", "thrilling", "despise"}

	dictLg = []string{
		"resetting", "departs", "blanket", "layered", "fungi", "lofty", "philosophized", "engel", "venturers", "counterfeited", "caruso", "crumble", "anticipatory", "ovid", "colonization", "consent", "initiating", "interface", "processes", "uprisings", "overlooked", "osha", "exasperated", "depth", "rubbed", "changer", "syllogisms", "intracity", "restrainer", "championing", "pertained", "eyesore", "unborn", "painters", "intransigent", "dentist", "transmissions", "naught", "patricia", "contemporariness", "stupefy", "litmus", "mythologies", "augurs", "weir", "menace", "inhibit", "quartets", "aviary", "mercer", "millenarian", "lacerta", "agenda", "numbness", "clomp", "flagged", "indulged", "pollutant", "mistypes", "cactus", "trapping", "flings", "bunglers", "motifs", "melanesia", "murphy", "eradicated", "articulately", "therapies", "honorer", "glistened", "irreducibly", "interchangeability", "cryptography", "demonstrations", "inlaid", "numbers", "epaulets", "chug", "meld", "damnation", "opposing", "correct", "frustrated", "queuers", "hopelessly", "adverbs", "seasoning", "generously", "ku", "earmarking", "networked", "quarryman", "fictive", "helpfully", "enchanter", "erosion", "swelling", "fumigant", "graspingly", "ward", "rivalries", "allemande", "thaws", "tidal", "greets", "bells", "slows", "around", "solution", "globularity", "formerly", "maseru", "fluency", "explainer", "songs", "smithereens", "explorations", "kuwait", "idealizations", "steeples", "tappet", "casts", "laps", "occident", "outlived", "monitoring", "pressings", "menus", "trefoil", "blues", "cede", "habit", "aimer", "chuff", "pearlite", "stiffly", "insurers", "meriting", "busses", "filamentary", "editors", "unsure", "placater", "blatant", "loincloth", "marten", "personalitys", "sulfonamide", "ridiculing", "decremented", "caking", "cucumbers", "dramaturgy", "preserving", "lexicography", "bray", "avarice", "prostitution", "baptistrys", "min", "popularly", "exemption", "extremity", "gargled", "severely", "yachtsmen", "leachate", "polysaccharide", "endpoint", "oman", "pry", "sax", "unparsed", "gretchen", "vanquish", "leathery", "ncaa", "boasts", "blackberrys", "comparably", "gypsys", "caspian", "varieties", "utterance", "tenor", "charismatic", "adorned", "smart", "referents", "thanklessness", "basements", "attenuating", "warner", "conferences", "autocrats", "stumbling", "abasements", "baronet", "explores", "circuitously", "abelson", "november", "coefficients", "wu", "spectrograph", "compresses", "deathrates", "meteoric", "pittston", "canals", "telemetry", "relished", "peeking", "sweeteners", "impel", "canner", "swarms", "mason", "freckle", "retentiveness", "berrys", "aphonic", "therapists", "gleanings", "lithology", "dopes", "needed", "broomcorn", "mumbling", "scottsdale", "presidents", "thrilling", "commandments", "troglodyte", "predefines", "incompleteness", "lodging", "dendritic", "fluster", "freak", "swerve", "crumples", "enforceable", "appendages", "fine", "muddy", "tires", "dominating", "kidded", "mutant", "lowe", "avenged", "ideas", "ammonium", "peppery", "ritually", "woeful", "persecuted", "bankruptcies", "unfortunate", "irrigates", "freshest", "counterparts", "howdy", "loophole", "wonderfulness", "illegalities", "remodeled", "baptistery", "bovine", "dickens", "mackintosh", "flippant", "stiles", "suffered", "contradistinction", "woodwork", "substitute", "mishaps", "levels", "nirvana", "treaty", "boyle", "ten", "ungrounded", "privacies", "serviceberry", "squatted", "nectarine", "scaling", "potbelly", "javelins", "enquiry", "require", "persevering", "darning", "abiding", "howsoever", "devils", "electrocutes", "hrothgar", "inherent", "jots", "annexed", "mastermind", "sanctions", "rope", "extrude", "finicky", "mental", "savannah", "lakes", "dirt", "blare", "piccolo", "prog", "solstice", "incomparable", "accumulated", "avoidance", "ringing", "miterwort", "robbers", "augustine", "repelling", "helium", "foreheads", "contraptions", "regina", "patricians", "aristocracy", "noun", "indigo", "beirut", "telescope", "jobholder", "apples", "leasehold", "localization", "kidnapped", "beginners", "ingather", "subsequences", "bushwhacking", "miserly", "sanded", "anomalies", "boulevards", "wholesomeness", "pedagogic", "suspicious", "hemostat", "polygonal", "scrutiny", "architecture", "countess", "yokes", "oily", "describing", "welcome", "ratty", "hammocks", "dig", "devotes", "reclaimed", "winston", "ensnared", "odorous", "odors", "riddled", "captured", "u.s", "axons", "described", "premiere", "propyl", "recollecting", "investigators", "remediable", "substantiate", "delusions", "infused", "burbank", "mountaineer", "indefiniteness", "benefice", "ingest", "remoteness", "regale", "adjustments", "yarn", "actuate", "puts", "factor", "statuesque", "pronounces", "teachable", "blabbermouth", "haberman", "stickiest", "orestes", "odors", "implementations", "reservations", "recycled", "indirects", "monroe", "tailor", "prose", "id", "revisions", "feeded", "waltz", "resists", "enactment", "vagrant", "affirm", "prate", "hecuba", "passer", "psychological", "turnips", "maurice", "mediates", "incubators", "minimum", "matsumoto", "grabbing", "eels", "attraction", "lutanist", "crockett", "floyd", "seedy", "courtney", "chemist", "oration", "newfoundland", "flapping", "burroughs", "pressurize", "affiliation", "ritchie", "isinglass", "fairies", "sideman", "battalions", "retentively", "adornment", "multiplicity", "pickups", "brazil", "massachusetts", "heron", "blockhouse", "slowed", "untruthful", "directrix", "thickens", "madam", "posting", "artichoke", "ammonia", "tapered", "bags", "brainstorm", "classed", "nave", "frigid", "sided", "covenants", "fetched", "mysteriousness", "cooler", "scrambling", "bleedings", "insides", "americas", "butternut", "entitys", "navigated", "decelerate", "bulb", "grabber", "teakwood", "chases", "friendship", "fireproof", "vegetate", "craves", "modulators", "late", "signing", "brightly", "supply", "rocco", "obstinately", "edmonds", "thistle", "cecil", "gs", "mae", "bobble", "remiss", "microphone", "administrate", "glorification", "wrapper", "transform", "authorship", "tx", "tomb", "irremovable", "reformulating", "decadence", "curran", "equivalence", "tenderloin", "missioner", "walkout", "synthesizing", "fickleness", "struggled", "romeo", "dystrophy", "lieutenants", "astronauts", "rink", "seemed", "neighbor", "ticks", "coffey", "hole", "burrower", "secretarys", "contain", "responsibly", "locutor", "untoward", "sweepstake", "sociologically", "aside", "decimated", "romantics", "crests", "straightforwardness", "salutations", "measurements", "degraded", "cyanate", "developed", "balks", "proofread", "acclimating", "tightener", "sextet", "otherworldly", "sidemen", "uncleanly", "yap", "stammers", "projection", "coughing", "discs", "wrote", "blanc", "needle", "sapping", "coating", "tempo", "venomously", "pilgrimage", "placate", "go", "soulful", "irradiate", "thoroughness", "timidity", "conjures", "denigrates", "sufferer", "predating", "osborn", "ivies", "cramming", "words", "dial", "deluded", "huron", "justices", "excommunicate", "retardant", "squirrels", "redistributed", "courtroom", "artillerist", "gator", "peal", "dragons", "realigned", "witt", "coronets", "shall", "peasantry", "possessed", "swan", "circulatory", "succinct", "zellerbach", "psychophysics", "ledges", "tigress", "gigavolt", "mandates", "assignment", "noticing", "approachers", "pitch", "statuary", "prickly", "polo", "loadings", "catbird", "squinting", "pants", "clarified", "commando", "besetting", "significance", "archaeologists", "apricots", "postorder", "neff", "cocking", "perfumed", "spectral", "stoppers", "parlors", "delete", "recaptures", "dreaming", "nickels", "downtrend", "friendless", "paraguay", "pollock", "masquerades", "breakwaters", "doubter", "ashland", "bumblebees", "skunk", "sinner", "thresholds", "neoclassic", "bimonthlies", "partake", "pursue", "tomograph", "tasted", "commandment", "postman", "stowaway", "behalf", "partition", "reviling", "daly", "congratulated", "grata", "valuate", "earthquakes", "massacre", "northerly", "caskets", "cathedral", "brothers", "generalizing", "clinging", "quavering", "matron", "botchers", "surgically", "sixties", "asthma", "circumlocutions", "algebras", "lubell", "portend", "fancying", "pelvic", "jailer", "excavating", "pored", "jittery", "violation", "whimsys", "abridged", "unhappy", "assistant", "largemouth", "coughs", "antiformant", "principal", "seeing", "marketplace", "goblins", "manometer", "evolutes", "nazis", "wad", "pities", "relabel", "proline", "straightaway", "loaned", "enmesh", "botching", "axially", "uniaxial", "chaos", "secondhand", "alkalis", "anthropogenic", "chubbiest", "bawls", "milieu", "sluggish", "scoundrels", "nestor", "pyrrhic", "radar", "unloaded", "braggart", "expelling", "autonomy", "ivanhoe", "bronchitis", "too", "ghost", "stabilize", "rosebuds", "touch", "abysmal", "carlson", "missoula", "biometry", "piteous", "inflexibility", "search", "lustiness", "monday", "bema", "maidservant", "duress", "cycled", "injudicious", "deforest", "knightsbridge", "gunky", "matchers", "protease", "consciousness", "arenaceous", "oxide", "cope", "shout", "narrowing", "categorized", "coughed", "edicts", "necropsy", "extinct", "profited", "deities", "uses", "owl", "lawyers", "desires", "perspicuity", "horseshoer", "critic", "deafer", "bewhisker", "affidavits", "itself", "subscribed", "ferric", "heir", "frown", "rapids", "saviors", "assembled", "unusable", "immaturity", "systemization", "tumbled", "disclosures", "sepia", "swapping", "sampled", "municipality", "renunciate", "reticles", "winker", "acquiescing", "tarpon", "portraits", "defense", "flowed", "homemakers", "launchings", "nettled", "homicide", "playhouse", "anxiety", "thunderstorms", "adumbrates", "diners", "incarcerate", "viewing", "petters", "sputtered", "tractor", "commends", "exemplifiers", "cameo", "gully", "colonized", "collide", "kneads", "univalent", "skiff", "liquidus", "clan", "heresy", "unsuccessful", "parkinson", "confrontations", "rhenish", "cadent", "foursome", "scottish", "sterilizes", "defray", "receptions", "britons", "outweighed", "descend", "calories", "reek", "awaits", "schultz", "mood", "waltham", "bridgetown", "sing", "factorizations", "crunchier", "tang", "stagnant", "quickly", "prohibited", "true", "kenneth", "setback", "refereeing", "begin", "underdone", "officeholder", "eagerness", "stylites", "recording", "extensibility", "warped", "refuter", "whores", "tulsa", "cooky", "executor", "capacitive", "emeriti", "outcries", "squires", "uniforms", "palmyra", "synoptic", "parenthood", "cheek", "diversifying", "furthermost", "hexagonal", "daylights", "cowpony", "subscribing", "orbited", "underestimate", "exultation", "fathered", "correcting", "dunlap", "slough", "oldenburg", "bandages", "cheerily", "lillian", "areawide", "airfield", "lovelorn", "blomquist", "crucial", "braes", "chews", "disobeyed", "medieval", "gird", "obtrusive", "connectors", "vale", "exhibition", "cowslips", "analogue", "cutters", "eccentric", "differs", "capstone", "tuners", "suitableness", "sham", "admire", "tragedy", "sleepy", "swampy", "lurched", "coverlets", "provokes", "disqualify", "thursdays", "trader", "debaters", "inspector", "synaptic", "identify", "eyes", "lowliest", "duplicable", "gunderson", "extensible", "surpass", "throw", "repressions", "shadbush", "baskets", "butlers", "whod", "judical", "detachment", "pullback", "franker", "staffing", "useful", "acolytes", "soldiery", "hymns", "seep", "pollen", "eminence", "mawkish", "neutron", "groceries", "stuyvesant", "hoodwink", "cachalot", "guards", "motionlessness", "maturities", "frances", "bipolar", "doublet", "contributors", "uncleared", "napoleon", "deserts", "fallacies", "disposing", "umbrellas", "sediment", "hedges", "changed", "spectators", "cleaning", "inwards", "colombo", "exhume", "albany", "slaves", "cop", "murderous", "modifying", "juxtaposition"}
)
