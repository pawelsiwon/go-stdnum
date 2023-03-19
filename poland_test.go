package stdnum

import (
	"errors"
	"testing"
)

func TestPolandValidatePesel(t *testing.T) {

	t.Run("should correctly validate", func(t *testing.T) {
		values := []string{
			"07022506123",
			"06111005260",
			"00082800293",
			"06091903709",
			"05081705372",
			"50112206274",
			"51081006128",
			"52040908691",
			"59031302485",
			"57011501925",
			"92012705732",
			"91032108066",
			"90102209081",
			"91111604403",
			"97052305160",
			"04310409739",
			"02311504637",
			"07311504649",
			"02230204063",
			"06311309117",
			"37220955649",
			"37100973471",
			"60300498198",
			"57080934787",
			"16261562571",
			"45060766642",
			"79251891673",
			"71101993692",
			"03320331218",
			"45300773768",
			"00272857542",
			"75101184913",
			"88290592144",
			"03121666827",
			"94102399557",
			"48292285459",
			"13261974412",
			"16220776234",
			"56262427961",
			"94032824945",
			"21080227185",
			"30060557758",
			"46273094269",
			"97261813447",
			"45212574754",
			"73302642571",
			"60280294652",
			"38303078374",
			"90072619246",
			"98302649441",
			"07051957237",
			"82310985329",
			"07272149282",
			"64242939964",
			"18071362378",
			"96231941489",
			"83083123936",
			"40122898864",
			"39110182663",
			"83222042416",
			"12101151879",
			"02052418228",
			"63110386949",
			"09221857176",
			"63231111615",
			"83121518775",
			"50272859655",
			"00021515929",
			"63311543134",
			"48081861271",
			"86241554724",
			"40122296273",
			"66102226881",
			"64222228488",
			"71272243628",
			"48120626933",
			"17321355986",
			"31220515715",
			"69011594188",
			"87221881276",
			"88322193943",
			"04311842687",
			"69020677115",
			"42232857517",
			"51251977614",
			"77260632991",
			"39302991848",
			"53310256766",
			"59082728359",
			"17280588252",
			"28021871928",
			"74011457179",
			"46251064642",
			"87302233417",
			"76251354746",
			"40301783932",
			"12032174945",
			"09272374833",
			"16282584952",
			"79031088758",
			"98280933983",
			"93112963275",
			"21021166144",
			"51110126661",
			"02081271281",
			"05042231773",
			"41081856441",
			"62030184129",
			"35221483974",
			"35031913898",
			"14101012226",
			"49060647235",
			"68311887693",
			"59061564321",
			"49101657944",
			"84050197723",
			"25252326746",
			"16251626135",
			"13092038295",
			"38210823548",
		}

		for _, v := range values {
			result := PolandValidatePesel(v)
			if !result.Valid || result.Error != nil {
				t.Errorf("should be validated correctly %v", v)
			}
		}
	})

	t.Run("should validate incorrectly", func(t *testing.T) {
		values := []string{
			"13092038296",
		}

		for _, v := range values {
			result := PolandValidatePesel(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrect) {
				t.Errorf("should return incorrect result")
			}
		}
	})

	t.Run("should return incorrect length", func(t *testing.T) {
		values := []string{
			"1728058825",
			"172805882522",
		}

		for _, v := range values {
			result := PolandValidatePesel(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrectInputLength) {
				t.Errorf("should return incorrect input length")
			}
		}
	})

	t.Run("should return only digits allowed", func(t *testing.T) {
		values := []string{
			"3821082354B",
		}

		for _, v := range values {
			result := PolandValidatePesel(v)
			if result.Valid || !errors.Is(result.Error, ErrOnlyDigitsAllowed) {
				t.Errorf("should return only digits allowed")
			}
		}
	})
}

func TestPolandValidateRegon(t *testing.T) {

	t.Run("should correctly validate", func(t *testing.T) {
		values := []string{
			"639559936",
			"516324335",
			"976209610",
			"157043401",
			"037967880",
			"076929337",
			"393735323",
			"913289784",
			"450852157",
			"256556650",
			"537072305",
			"033275040",
			"472493244",
			"12345678512347",
		}

		for _, v := range values {
			result := PolandValidateRegon(v)
			if !result.Valid || result.Error != nil {
				t.Errorf("should be validated correctly %v", v)
			}
		}
	})

	t.Run("should validate incorrectly", func(t *testing.T) {
		values := []string{
			"472493245",
			"12345678612347",
			"12345678512349",
		}

		for _, v := range values {
			result := PolandValidateRegon(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrect) {
				t.Errorf("should return incorrect result")
			}
		}
	})

	t.Run("should return incorrect length", func(t *testing.T) {
		values := []string{
			"47249324",
			"4724932444",
			"1234567851234",
			"123456785123477",
		}

		for _, v := range values {
			result := PolandValidateRegon(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrectInputLength) {
				t.Errorf("should return incorrect input length")
			}
		}
	})

	t.Run("should return only digits allowed", func(t *testing.T) {
		values := []string{
			"47249324B",
			"1234567851234B",
		}

		for _, v := range values {
			result := PolandValidateRegon(v)
			if result.Valid || !errors.Is(result.Error, ErrOnlyDigitsAllowed) {
				t.Errorf("should return only digits allowed")
			}
		}
	})
}

func TestPolandValidateNIP(t *testing.T) {

	t.Run("should correctly validate", func(t *testing.T) {
		values := []string{
			"8567346215",
			"9007809152",
			"9435719559",
			"4735628226",
			"5479479168",
			"0175128695",
			"0645019680",
			"1414061903",
			"9925933544",
			"1950101503",
			"7798781287",
			"6223086800",
			"7948777992",
			"5310012566",
			"2348071382",
			"3239474377",
			"1510517582",
			"4949074162",
			"1360683639",
			"0078229279",
			"0929424421",
			"3195104305",
			"9256833161",
			"9622191150",
			"2499808314",
			"9759785365",
			"8785747969",
			"7172321202",
			"2134010087",
			"8525393960",
			"8664472611",
			"5033240539",
			"2570172242",
			"2134305960",
			"6169427433",
			"5750973407",
			"3990118994",
			"7208824861",
			"7610286286",
			"3624024039",
			"6835556444",
			"8298892535",
			"9103078202",
			"7442021534",
			"5679690991",
			"7827572877",
			"4928890459",
			"4007665532",
			"4534317169",
			"4451944722",
			"6690053588",
		}

		for _, v := range values {
			result := PolandValidateNip(v)
			if !result.Valid || result.Error != nil {
				t.Errorf("should be validated correctly %v", v)
			}
		}
	})

	t.Run("should validate incorrectly", func(t *testing.T) {
		values := []string{
			"6690053589",
		}

		for _, v := range values {
			result := PolandValidateNip(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrect) {
				t.Errorf("should return incorrect result")
			}
		}
	})

	t.Run("should return incorrect length", func(t *testing.T) {
		values := []string{
			"669005358",
			"66900535888",
		}

		for _, v := range values {
			result := PolandValidateNip(v)
			if result.Valid || !errors.Is(result.Error, ErrIncorrectInputLength) {
				t.Errorf("should return incorrect input length")
			}
		}
	})

	t.Run("should return only digits allowed", func(t *testing.T) {
		values := []string{
			"669005358B",
		}

		for _, v := range values {
			result := PolandValidateNip(v)
			if result.Valid || !errors.Is(result.Error, ErrOnlyDigitsAllowed) {
				t.Errorf("should return only digits allowed")
			}
		}
	})
}
