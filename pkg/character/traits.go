package character

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllTraits() []string {
	traits := getNegativeTraits()
	positiveTraits := getPositiveTraits()

	traits = append(traits, positiveTraits...)

	return traits
}

func getNegativeTraits() []string {
	traits := []string{
		"abrasive",
		"abrupt",
		"absentminded",
		"aggressive",
		"agonizing",
		"aimless",
		"aloof",
		"ambitious",
		"amoral",
		"amusing",
		"angry",
		"anxious",
		"apathetic",
		"arbitrary",
		"argumentative",
		"arrogant",
		"artful",
		"artificial",
		"ascetic",
		"asocial",
		"authoritarian",
		"bewildered",
		"big-thinking",
		"bizarre",
		"bland",
		"blunt",
		"boisterous",
		"boyish",
		"breezy",
		"brittle",
		"brutal",
		"businesslike",
		"busy",
		"calculating",
		"callous",
		"cantankerous",
		"careless",
		"casual",
		"cerebral",
		"charmless",
		"childish",
		"chummy",
		"circumspect",
		"clumsy",
		"coarse",
		"cold",
		"colorless",
		"competitive",
		"complacent",
		"complaining",
		"complex",
		"compulsive",
		"conceited",
		"condemnatory",
		"confidential",
		"conformist",
		"confused",
		"conservative",
		"contemptible",
		"contradictory",
		"conventional",
		"cowardly",
		"crass",
		"crazy",
		"criminal",
		"crisp",
		"critical",
		"crude",
		"cruel",
		"cute",
		"cynical",
		"decadent",
		"deceitful",
		"deceptive",
		"delicate",
		"demanding",
		"dependent",
		"desperate",
		"destructive",
		"determined",
		"devious",
		"difficult",
		"disconcerting",
		"discontented",
		"discouraging",
		"discourteous",
		"dishonest",
		"disloyal",
		"disobedient",
		"disorderly",
		"disorganized",
		"disputatious",
		"disrespectful",
		"disruptive",
		"dissonant",
		"distractible",
		"disturbing",
		"dogmatic",
		"dominating",
		"domineering",
		"dreamy",
		"driving",
		"droll",
		"dry",
		"dull",
		"earthy",
		"easily discouraged",
		"effeminate",
		"egocentric",
		"emotional",
		"enigmatic",
		"envious",
		"erratic",
		"escapist",
		"experimental",
		"extravagant",
		"extreme",
		"faithless",
		"false",
		"familial",
		"fanatical",
		"fanciful",
		"fatalistic",
		"fawning",
		"fearful",
		"fickle",
		"fiery",
		"fixed",
		"flamboyant",
		"folksy",
		"foolish",
		"forgetful",
		"formal",
		"fraudulent",
		"freewheeling",
		"frightening",
		"frivolous",
		"frugal",
		"glamorous",
		"gloomy",
		"graceless",
		"greedy",
		"grim",
		"guileless",
		"gullible",
		"hateful",
		"haughty",
		"hedonistic",
		"hesitant",
		"hidebound",
		"high-handed",
		"high-spirited",
		"hostile",
		"hurried",
		"hypnotic",
		"iconoclastic",
		"idiosyncratic",
		"ignorant",
		"imitative",
		"impassive",
		"impatient",
		"impersonal",
		"impractical",
		"impressionable",
		"imprudent",
		"impulsive",
		"inconsiderate",
		"incurious",
		"indecisive",
		"indulgent",
		"inert",
		"inhibited",
		"insecure",
		"insensitive",
		"insincere",
		"insulting",
		"intense",
		"intolerant",
		"invisible",
		"irascible",
		"irrational",
		"irreligious",
		"irresponsible",
		"irreverent",
		"irritable",
		"lazy",
		"malicious",
		"mannerless",
		"maternal",
		"mechanical",
		"meddlesome",
		"melancholic",
		"mellow",
		"messy",
		"miserable",
		"miserly",
		"misguided",
		"mistaken",
		"modern",
		"money-minded",
		"moody",
		"moralistic",
		"morbid",
		"muddle-headed",
		"mystical",
		"naive",
		"narcissistic",
		"narrow",
		"narrow-minded",
		"negative",
		"neglectful",
		"neurotic",
		"neutral",
		"nihilistic",
		"noncommittal",
		"noncompetitive",
		"obedient",
		"obnoxious",
		"obsessive",
		"obvious",
		"odd",
		"offhand",
		"old-fashioned",
		"one-dimensional",
		"one-sided",
		"opinionated",
		"opportunistic",
		"oppressed",
		"ordinary",
		"outrageous",
		"outspoken",
		"paranoid",
		"passive",
		"paternalistic",
		"pedantic",
		"perverse",
		"petty",
		"physical",
		"placid",
		"plodding",
		"political",
		"pompous",
		"possessive",
		"power-hungry",
		"predatory",
		"predictable",
		"prejudiced",
		"preoccupied",
		"presumptuous",
		"pretentious",
		"prim",
		"private",
		"procrastinating",
		"progressive",
		"proud",
		"provocative",
		"pure",
		"puritanical",
		"questioning",
		"quiet",
		"quirky",
		"reactionary",
		"reactive",
		"regimental",
		"regretful",
		"religious",
		"repentant",
		"repressed",
		"resentful",
		"reserved",
		"restrained",
		"retiring",
		"ridiculous",
		"rigid",
		"ritualistic",
		"ruined",
		"sadistic",
		"sanctimonious",
		"sarcastic",
		"scheming",
		"scornful",
		"secretive",
		"sedentary",
		"self-conscious",
		"self-indulgent",
		"selfish",
		"sensual",
		"shallow",
		"shortsighted",
		"skeptical",
		"sloppy",
		"slow",
		"sly",
		"small-thinking",
		"smooth",
		"soft",
		"soft",
		"softheaded",
		"solemn",
		"solitary",
		"sordid",
		"steely",
		"stern",
		"stiff",
		"strict",
		"stubborn",
		"stupid",
		"stylish",
		"subjective",
		"submissive",
		"superficial",
		"superstitious",
		"surprising",
		"suspicious",
		"tactless",
		"tasteless",
		"tense",
		"thievish",
		"thoughtless",
		"timid",
		"tough",
		"transparent",
		"treacherous",
		"trendy",
		"troublesome",
		"unaggressive",
		"unambitious",
		"unappreciative",
		"uncaring",
		"unceremonious",
		"unchanging",
		"uncharitable",
		"unconvincing",
		"uncooperative",
		"uncreative",
		"uncritical",
		"unctuous",
		"undemanding",
		"undisciplined",
		"unfathomable",
		"unfriendly",
		"ungrateful",
		"unhealthy",
		"unhurried",
		"unimaginative",
		"unimpressive",
		"uninhibited",
		"unlovable",
		"unpatriotic",
		"unpolished",
		"unpredictable",
		"unprincipled",
		"unrealistic",
		"unreflective",
		"unreliable",
		"unrestrained",
		"unsentimental",
		"unstable",
		"vacuous",
		"vague",
		"venomous",
		"vindictive",
		"vulnerable",
		"weak",
		"whimsical",
		"willful",
	}

	return traits
}

func getPositiveTraits() []string {
	traits := []string{
		"accessible",
		"active",
		"adaptable",
		"admirable",
		"adventurous",
		"agreeable",
		"alert",
		"amiable",
		"anticipative",
		"appreciative",
		"articulate",
		"aspiring",
		"athletic",
		"attractive",
		"balanced",
		"benevolent",
		"brilliant",
		"calm",
		"capable",
		"captivating",
		"caring",
		"challenging",
		"charismatic",
		"charming",
		"cheerful",
		"clean",
		"clear-headed",
		"clever",
		"colorful",
		"companionly",
		"compassionate",
		"conciliatory",
		"confident",
		"conscientious",
		"considerate",
		"constant",
		"contemplative",
		"cooperative",
		"courageous",
		"courteous",
		"creative",
		"cultured",
		"curious",
		"daring",
		"debonair",
		"decent",
		"decisive",
		"dedicated",
		"deep",
		"dignified",
		"directed",
		"disciplined",
		"discreet",
		"dramatic",
		"dutiful",
		"dynamic",
		"earnest",
		"ebullient",
		"educated",
		"efficient",
		"elegant",
		"eloquent",
		"empathetic",
		"energetic",
		"enthusiastic",
		"esthetic",
		"exciting",
		"extraordinary",
		"fair",
		"faithful",
		"farsighted",
		"felicific",
		"firm",
		"flexible",
		"focused",
		"forceful",
		"forgiving",
		"forthright",
		"freethinking",
		"friendly",
		"fun-loving",
		"gallant",
		"generous",
		"gentle",
		"genuine",
		"good-natured",
		"gracious",
		"hardworking",
		"healthy",
		"hearty",
		"helpful",
		"heroic",
		"high-minded",
		"honest",
		"honorable",
		"humble",
		"humorous",
		"idealistic",
		"imaginative",
		"impressive",
		"incisive",
		"incorruptible",
		"independent",
		"individualistic",
		"innovative",
		"inoffensive",
		"insightful",
		"insouciant",
		"intelligent",
		"intuitive",
		"invulnerable",
		"kind",
		"knowledge",
		"leader",
		"leisurely",
		"liberal",
		"logical",
		"lovable",
		"loyal",
		"lyrical",
		"magnanimous",
		"many-sided",
		"masculine",
		"mature",
		"methodical",
		"meticulous",
		"moderate",
		"modest",
		"multi-leveled",
		"neat",
		"objective",
		"observant",
		"open",
		"optimistic",
		"orderly",
		"organized",
		"original",
		"painstaking",
		"passionate",
		"patient",
		"patriotic",
		"peaceful",
		"perceptive",
		"perfectionist",
		"personable",
		"persuasive",
		"playful",
		"polished",
		"popular",
		"practical",
		"precise",
		"principled",
		"profound",
		"protean",
		"protective",
		"providential",
		"prudent",
		"punctual",
		"purposeful",
		"rational",
		"realistic",
		"reflective",
		"relaxed",
		"reliable",
		"resourceful",
		"respectful",
		"responsible",
		"responsive",
		"reverential",
		"romantic",
		"rustic",
		"sage",
		"sane",
		"scholarly",
		"scrupulous",
		"secure",
		"selfless",
		"self-critical",
		"self-defacing",
		"self-denying",
		"self-reliant",
		"self-sufficent",
		"sensitive",
		"sentimental",
		"seraphic",
		"serious",
		"sexy",
		"sharing",
		"shrewd",
		"simple",
		"skillful",
		"sober",
		"sociable",
		"solid",
		"sophisticated",
		"spontaneous",
		"sporting",
		"stable",
		"steadfast",
		"steady",
		"stoic",
		"strong",
		"studious",
		"suave",
		"subtle",
		"sweet",
		"sympathetic",
		"systematic",
		"tasteful",
		"teacherly",
		"thorough",
		"tidy",
		"tolerant",
		"tractable",
		"trusting",
		"uncomplaining",
		"understanding",
		"undogmatic",
		"upright",
		"urbane",
		"venturesome",
		"vivacious",
		"warm",
		"well-bred",
		"well-read",
		"well-rounded",
		"winning",
		"wise",
		"witty",
		"youthful",
	}

	return traits
}

func getRandomTrait() (string, error) {
	traits := getAllTraits()

	trait, err := random.String(traits)
	if err != nil {
		err = fmt.Errorf("Could not generate trait: %w", err)
		return "", err
	}
	return trait, nil
}

func getRandomNegativeTraits(max int) ([]string, error) {
	var trait string
	var err error
	possibleTraits := getNegativeTraits()

	traits := []string{}

	for i := 0; i < max; i++ {
		trait, err = random.String(possibleTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate trait: %w", err)
			return []string{}, err
		}
		for slices.StringIn(trait, traits) {
			trait, err = random.String(possibleTraits)
			if err != nil {
				err = fmt.Errorf("Could not generate trait: %w", err)
				return []string{}, err
			}
		}
		traits = append(traits, trait)
	}

	return traits, nil
}

func getRandomPositiveTraits(max int) ([]string, error) {
	var trait string
	var err error
	possibleTraits := getPositiveTraits()

	traits := []string{}

	for i := 0; i < max; i++ {
		trait, err = random.String(possibleTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate trait: %w", err)
			return []string{}, err
		}
		for slices.StringIn(trait, traits) {
			trait, err = random.String(possibleTraits)
			if err != nil {
				err = fmt.Errorf("Could not generate trait: %w", err)
				return []string{}, err
			}
		}
		traits = append(traits, trait)
	}

	return traits, nil
}

func getRandomTraits() ([]string, error) {
	var trait string
	var err error
	possibleTraits := getAllTraits()

	traits := []string{}

	for i := 0; i < 2; i++ {
		trait, err = random.String(possibleTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate trait: %w", err)
			return []string{}, err
		}
		for slices.StringIn(trait, traits) {
			trait, err = random.String(possibleTraits)
			if err != nil {
				err = fmt.Errorf("Could not generate trait: %w", err)
				return []string{}, err
			}
		}
		traits = append(traits, trait)
	}

	return traits, nil
}
