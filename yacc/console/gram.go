// Code generated by goyacc -o gram.go -p yy gram.y. DO NOT EDIT.

//line gram.y:3
package spqrparser

import __yyfmt__ "fmt"

//line gram.y:3

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//line gram.y:25
type yySymType struct {
	yys      int
	str      string
	strlist  []string
	byte     byte
	bytes    []byte
	integer  int
	uinteger uint
	bool     bool
	empty    struct{}

	set       *Set
	statement Statement
	show      *Show

	drop   *Drop
	create *Create

	kill   *Kill
	lock   *Lock
	unlock *Unlock

	krbound *KeyRangeBound

	ds            *DistributionDefinition
	kr            *KeyRangeDefinition
	shard         *ShardDefinition
	sharding_rule *ShardingRuleDefinition

	register_router   *RegisterRouter
	unregister_router *UnregisterRouter

	split *SplitKeyRange
	move  *MoveKeyRange
	unite *UniteKeyRange

	redistribute *RedistributeKeyRange

	invalidate_cache *InvalidateCache

	shutdown *Shutdown
	listen   *Listen

	trace     *TraceStmt
	stoptrace *StopTraceStmt

	distribution *DistributionDefinition

	alter                *Alter
	alter_distribution   *AlterDistribution
	distributed_relation *DistributedRelation

	relations    []*DistributedRelation
	entrieslist  []ShardingRuleEntry
	dEntrieslist []DistributionKeyEntry

	shruleEntry ShardingRuleEntry

	distrKeyEntry DistributionKeyEntry

	sharding_rule_selector *ShardingRuleSelector
	key_range_selector     *KeyRangeSelector
	distribution_selector  *DistributionSelector

	colref ColumnRef
	where  WhereClauseNode

	order_clause OrderClause
	opt_asc_desc OptAscDesc

	group_clause GroupByClause
}

const IDENT = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const WHERE = 57350
const OR = 57351
const AND = 57352
const TEQ = 57353
const TCOMMA = 57354
const SCONST = 57355
const ICONST = 57356
const TSEMICOLON = 57357
const TOPENBR = 57358
const TCLOSEBR = 57359
const SHUTDOWN = 57360
const LISTEN = 57361
const REGISTER = 57362
const UNREGISTER = 57363
const ROUTER = 57364
const ROUTE = 57365
const CREATE = 57366
const ADD = 57367
const DROP = 57368
const LOCK = 57369
const UNLOCK = 57370
const SPLIT = 57371
const MOVE = 57372
const COMPOSE = 57373
const SET = 57374
const CASCADE = 57375
const ATTACH = 57376
const ALTER = 57377
const DETACH = 57378
const REDISTRIBUTE = 57379
const REFERENCE = 57380
const SHARDING = 57381
const COLUMN = 57382
const TABLE = 57383
const HASH = 57384
const FUNCTION = 57385
const KEY = 57386
const RANGE = 57387
const DISTRIBUTION = 57388
const RELATION = 57389
const REPLICATED = 57390
const SHARDS = 57391
const KEY_RANGES = 57392
const ROUTERS = 57393
const SHARD = 57394
const HOST = 57395
const SHARDING_RULES = 57396
const RULE = 57397
const COLUMNS = 57398
const VERSION = 57399
const HOSTS = 57400
const BY = 57401
const FROM = 57402
const TO = 57403
const WITH = 57404
const UNITE = 57405
const ALL = 57406
const ADDRESS = 57407
const FOR = 57408
const CLIENT = 57409
const BATCH = 57410
const SIZE = 57411
const INVALIDATE = 57412
const CACHE = 57413
const IDENTITY = 57414
const MURMUR = 57415
const CITY = 57416
const START = 57417
const STOP = 57418
const TRACE = 57419
const MESSAGES = 57420
const TASK = 57421
const GROUP = 57422
const VARCHAR = 57423
const INTEGER = 57424
const INT = 57425
const TYPES = 57426
const OP = 57427
const ASC = 57428
const DESC = 57429
const ORDER = 57430

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"COMMAND",
	"SHOW",
	"KILL",
	"WHERE",
	"OR",
	"AND",
	"TEQ",
	"TCOMMA",
	"SCONST",
	"ICONST",
	"TSEMICOLON",
	"TOPENBR",
	"TCLOSEBR",
	"SHUTDOWN",
	"LISTEN",
	"REGISTER",
	"UNREGISTER",
	"ROUTER",
	"ROUTE",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SPLIT",
	"MOVE",
	"COMPOSE",
	"SET",
	"CASCADE",
	"ATTACH",
	"ALTER",
	"DETACH",
	"REDISTRIBUTE",
	"REFERENCE",
	"SHARDING",
	"COLUMN",
	"TABLE",
	"HASH",
	"FUNCTION",
	"KEY",
	"RANGE",
	"DISTRIBUTION",
	"RELATION",
	"REPLICATED",
	"SHARDS",
	"KEY_RANGES",
	"ROUTERS",
	"SHARD",
	"HOST",
	"SHARDING_RULES",
	"RULE",
	"COLUMNS",
	"VERSION",
	"HOSTS",
	"BY",
	"FROM",
	"TO",
	"WITH",
	"UNITE",
	"ALL",
	"ADDRESS",
	"FOR",
	"CLIENT",
	"BATCH",
	"SIZE",
	"INVALIDATE",
	"CACHE",
	"IDENTITY",
	"MURMUR",
	"CITY",
	"START",
	"STOP",
	"TRACE",
	"MESSAGES",
	"TASK",
	"GROUP",
	"VARCHAR",
	"INTEGER",
	"INT",
	"TYPES",
	"OP",
	"ASC",
	"DESC",
	"ORDER",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:938

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 281

var yyAct = [...]uint8{
	148, 198, 244, 169, 201, 170, 193, 147, 168, 161,
	167, 173, 145, 160, 129, 156, 103, 176, 29, 30,
	241, 242, 194, 195, 196, 144, 158, 108, 136, 100,
	32, 31, 37, 38, 58, 76, 23, 22, 26, 27,
	28, 33, 34, 57, 235, 236, 237, 39, 90, 35,
	62, 89, 95, 216, 187, 60, 200, 64, 163, 67,
	75, 98, 133, 65, 99, 106, 107, 153, 91, 91,
	91, 119, 91, 166, 164, 36, 91, 238, 227, 113,
	115, 91, 40, 118, 200, 120, 121, 24, 25, 106,
	66, 117, 116, 210, 128, 131, 186, 177, 135, 172,
	134, 49, 139, 141, 102, 137, 50, 93, 47, 163,
	48, 139, 130, 88, 51, 67, 223, 149, 150, 151,
	152, 157, 142, 74, 217, 164, 154, 109, 140, 138,
	122, 165, 105, 92, 96, 110, 101, 132, 174, 56,
	49, 94, 231, 159, 69, 50, 224, 47, 202, 48,
	221, 220, 219, 51, 97, 189, 91, 188, 191, 182,
	127, 125, 104, 124, 203, 204, 59, 206, 63, 206,
	199, 190, 46, 197, 45, 85, 174, 205, 225, 78,
	207, 44, 43, 84, 208, 211, 42, 214, 77, 79,
	114, 212, 248, 130, 68, 70, 55, 215, 54, 222,
	80, 81, 82, 83, 179, 53, 52, 199, 87, 181,
	180, 91, 184, 228, 206, 226, 209, 213, 229, 185,
	78, 179, 230, 146, 233, 232, 181, 180, 239, 77,
	171, 218, 245, 112, 91, 72, 41, 1, 19, 246,
	18, 247, 17, 16, 15, 14, 12, 13, 250, 245,
	249, 251, 8, 9, 143, 240, 175, 126, 192, 155,
	123, 21, 86, 20, 234, 162, 243, 6, 5, 4,
	3, 7, 11, 10, 73, 71, 61, 2, 183, 178,
	111,
}

var yyPact = [...]int16{
	12, -1000, 171, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 62, 101, -34, -43, 11, 100, 100, 231,
	56, 175, -1000, 100, 100, 100, 100, 161, 153, 67,
	-20, -1000, -1000, -1000, -1000, -1000, -1000, 230, 87, 52,
	96, 72, -1000, -1000, -1000, -1000, 113, -3, -49, -1000,
	91, -1000, 49, 129, 68, 230, -53, 81, -1000, 90,
	-1000, 225, -1000, 176, 176, -1000, -1000, -1000, -1000, -1000,
	32, 30, 22, 9, 230, 66, -1000, 127, 230, -1000,
	120, -1000, -1000, 152, 77, 0, 42, 230, -50, 176,
	-1000, 65, 64, -1000, -1000, 129, -1000, -1000, -1000, -1000,
	230, -55, 207, -1000, -1000, -1000, 230, 230, 230, 230,
	2, -1000, -1000, -1000, 79, 74, -1000, -58, 71, 69,
	230, 13, 216, 41, 175, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -71, 38, 217, 207, 208, -1000, 37,
	-1000, -14, -1000, 175, 230, 74, -1000, 230, -59, 69,
	18, -1000, 106, 230, 230, -1000, 216, 157, -1000, -1000,
	-1000, -1000, 175, 204, -1000, -1000, 34, 230, 207, -1000,
	-1000, -1000, 200, 175, -1000, -1000, 216, -16, -1000, -1000,
	-1000, 78, 219, -1000, 110, 109, 108, 18, -1000, -1000,
	70, -1000, 103, -1000, -1000, 155, 216, 17, 204, 175,
	230, -1000, 217, -1000, -1000, 202, 176, 98, -59, -1000,
	-1000, -1000, -1000, 230, -28, 16, -1000, 230, -1000, -66,
	-1000, 230, -1000, -1000, -1000, -1000, -1000, -1000, 230, -10,
	-1000, -1000, -1000, 180, -1000, 106, -10, -1000, 230, -1000,
	-1000, -1000,
}

var yyPgo = [...]int16{
	0, 280, 12, 8, 10, 279, 278, 7, 3, 0,
	5, 277, 276, 166, 168, 275, 274, 273, 272, 271,
	270, 269, 268, 267, 182, 181, 174, 172, 13, 266,
	9, 2, 14, 265, 4, 264, 1, 263, 262, 261,
	260, 259, 15, 258, 257, 11, 6, 16, 256, 255,
	254, 253, 252, 247, 246, 245, 244, 243, 242, 240,
	238, 237, 236,
}

var yyR1 = [...]int8{
	0, 61, 62, 62, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 10, 8, 8, 8, 9, 5, 5,
	5, 6, 6, 7, 2, 2, 2, 1, 1, 15,
	16, 47, 47, 19, 19, 19, 19, 19, 19, 19,
	19, 20, 20, 20, 20, 22, 22, 23, 37, 38,
	38, 29, 29, 31, 42, 42, 41, 41, 40, 21,
	21, 21, 21, 21, 49, 49, 49, 48, 48, 50,
	50, 17, 52, 24, 24, 44, 44, 43, 43, 46,
	46, 46, 46, 46, 46, 25, 25, 28, 28, 30,
	32, 32, 33, 33, 35, 35, 35, 34, 34, 36,
	36, 3, 3, 4, 4, 26, 26, 27, 27, 45,
	45, 51, 12, 13, 14, 14, 55, 18, 18, 56,
	57, 57, 58, 54, 53, 39, 59, 60, 60,
}

var yyR2 = [...]int8{
	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 3, 3, 0, 2, 1,
	1, 1, 0, 2, 4, 2, 4, 3, 4, 3,
	3, 2, 2, 2, 2, 4, 4, 3, 2, 2,
	4, 3, 1, 2, 5, 2, 1, 2, 2, 2,
	2, 2, 2, 4, 1, 1, 0, 4, 0, 3,
	0, 5, 2, 3, 2, 3, 0, 3, 1, 1,
	2, 1, 1, 2, 2, 6, 5, 1, 2, 2,
	2, 0, 2, 2, 1, 1, 1, 3, 0, 3,
	0, 1, 1, 1, 3, 9, 8, 5, 4, 1,
	3, 2, 3, 3, 2, 2, 6, 3, 3, 4,
	7, 4, 4, 2, 1, 2, 5, 3, 3,
}

var yyChk = [...]int16{
	-1000, -61, -11, -20, -21, -22, -23, -19, -52, -51,
	-17, -18, -54, -53, -55, -56, -57, -58, -59, -60,
	-37, -39, 25, 24, 75, 76, 26, 27, 28, 6,
	7, 19, 18, 29, 30, 37, 63, 20, 21, 35,
	70, -62, 15, -24, -25, -26, -27, 46, 48, 39,
	44, 52, -24, -25, -26, -27, 38, 77, 77, -13,
	44, -12, 39, -14, 46, 52, 79, 48, -13, 44,
	-13, -15, 4, -16, 67, 4, -8, 13, 4, 14,
	-13, -13, -13, -13, 22, 22, -38, -14, 46, 71,
	-9, 4, 46, 55, 45, -9, 62, 41, 64, 67,
	78, 45, 55, -47, 33, 64, -9, -9, 80, 46,
	45, -1, 8, -10, 14, -10, 60, 61, 61, 62,
	-9, -9, 64, -40, 36, 34, -44, 40, -9, -32,
	41, -9, 60, 62, 58, -9, 78, -10, 64, -9,
	64, -9, -47, -50, 80, -2, 16, -7, -9, -9,
	-9, -9, -9, 65, 47, -41, -42, 47, 84, -32,
	-28, -30, -33, 40, 56, -9, 60, -4, -3, -8,
	-10, 14, 58, -45, -8, -48, 88, 59, -5, 4,
	10, 9, -2, -6, 4, 11, 59, 68, -8, -9,
	-42, -9, -43, -46, 81, 82, 83, -28, -36, -30,
	66, -34, 42, -9, -9, -4, 12, 23, -45, 12,
	59, -7, -2, 17, -8, -4, 69, 46, 12, 42,
	42, 42, -36, 46, 43, 23, -3, 61, -8, -7,
	-10, 44, -46, -9, -35, 72, 73, 74, 61, -9,
	-49, 86, 87, -29, -31, -9, -9, -36, 12, -34,
	-36, -31,
}

var yyDef = [...]int16{
	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 134, 0, 0, 0, 0, 0, 0, 0,
	0, 1, 3, 51, 52, 53, 54, 0, 0, 0,
	0, 0, 69, 70, 71, 72, 0, 0, 0, 43,
	0, 45, 0, 42, 0, 0, 0, 0, 82, 0,
	121, 37, 39, 0, 0, 40, 133, 24, 25, 26,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 135,
	86, 27, 84, 101, 0, 0, 0, 0, 0, 0,
	57, 0, 0, 47, 41, 42, 124, 49, 50, 125,
	0, 80, 0, 127, 23, 128, 0, 0, 0, 0,
	0, 137, 138, 59, 0, 0, 83, 0, 101, 0,
	0, 0, 0, 0, 0, 73, 55, 56, 44, 123,
	46, 122, 48, 78, 0, 38, 0, 0, 33, 0,
	129, 131, 132, 0, 0, 68, 66, 0, 0, 0,
	110, 97, 108, 0, 0, 100, 0, 0, 113, 111,
	112, 23, 0, 118, 119, 81, 0, 0, 0, 28,
	29, 30, 0, 0, 31, 32, 0, 0, 136, 60,
	67, 65, 85, 88, 89, 91, 92, 110, 96, 98,
	0, 99, 0, 102, 103, 0, 0, 0, 117, 0,
	0, 79, 36, 34, 35, 126, 0, 0, 0, 90,
	93, 94, 95, 0, 0, 0, 114, 0, 120, 76,
	130, 0, 87, 109, 107, 104, 105, 106, 0, 110,
	77, 74, 75, 64, 62, 108, 110, 116, 0, 63,
	115, 61,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:240
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:241
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:246
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:250
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:254
		{
			setParseTree(yylex, yyDollar[1].trace)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:258
		{
			setParseTree(yylex, yyDollar[1].stoptrace)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:262
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:266
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:270
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:274
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:278
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:282
		{
			setParseTree(yylex, yyDollar[1].listen)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:286
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:290
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:294
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:298
		{
			setParseTree(yylex, yyDollar[1].redistribute)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:302
		{
			setParseTree(yylex, yyDollar[1].unite)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:306
		{
			setParseTree(yylex, yyDollar[1].register_router)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:310
		{
			setParseTree(yylex, yyDollar[1].unregister_router)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:314
		{
			setParseTree(yylex, yyDollar[1].alter)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:318
		{
			setParseTree(yylex, yyDollar[1].invalidate_cache)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:323
		{
			yyVAL.uinteger = uint(yyDollar[1].uinteger)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:328
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:332
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:334
		{
			yyVAL.str = strconv.Itoa(int(yyDollar[1].uinteger))
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:339
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:345
		{
			yyVAL.str = yyDollar[1].str
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:347
		{
			yyVAL.str = "AND"
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:349
		{
			yyVAL.str = "OR"
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:354
		{
			yyVAL.str = yyDollar[1].str
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:356
		{
			yyVAL.str = "="
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:362
		{
			yyVAL.colref = ColumnRef{
				ColName: yyDollar[1].str,
			}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:370
		{
			yyVAL.where = yyDollar[2].where
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:373
		{
			yyVAL.where = WhereClauseLeaf{
				ColRef: yyDollar[1].colref,
				Op:     yyDollar[2].str,
				Value:  yyDollar[3].str,
			}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:381
		{
			yyVAL.where = WhereClauseOp{
				Op:    yyDollar[2].str,
				Left:  yyDollar[1].where,
				Right: yyDollar[3].where,
			}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:391
		{
			yyVAL.where = WhereClauseEmpty{}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:395
		{
			yyVAL.where = yyDollar[2].where
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:402
		{
			switch v := strings.ToLower(string(yyDollar[1].str)); v {
			case DatabasesStr, RoutersStr, PoolsStr, InstanceStr, ShardsStr, BackendConnectionsStr, KeyRangesStr, ShardingRules, ClientsStr, StatusStr, DistributionsStr, VersionStr, RelationsStr, TaskGroupStr, PreparedStatementsStr, QuantilesStr:
				yyVAL.str = v
			default:
				yyVAL.str = UnsupportedStr
			}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:413
		{
			switch v := string(yyDollar[1].str); v {
			case ClientStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:423
		{
			yyVAL.bool = true
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:423
		{
			yyVAL.bool = false
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:427
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].key_range_selector}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:431
		{
			yyVAL.drop = &Drop{Element: &KeyRangeSelector{KeyRangeID: `*`}}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:435
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].sharding_rule_selector}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:439
		{
			yyVAL.drop = &Drop{Element: &ShardingRuleSelector{ID: `*`}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:443
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].distribution_selector, CascadeDelete: yyDollar[3].bool}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:447
		{
			yyVAL.drop = &Drop{Element: &DistributionSelector{ID: `*`}, CascadeDelete: yyDollar[4].bool}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:451
		{
			yyVAL.drop = &Drop{Element: &ShardSelector{ID: yyDollar[3].str}}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:455
		{
			yyVAL.drop = &Drop{Element: &TaskGroupSelector{}}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:462
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:467
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:472
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:476
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:482
		{
			yyVAL.trace = &TraceStmt{All: true}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:485
		{
			yyVAL.trace = &TraceStmt{
				Client: yyDollar[4].uinteger,
			}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:493
		{
			yyVAL.stoptrace = &StopTraceStmt{}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:499
		{
			yyVAL.alter = &Alter{Element: yyDollar[2].alter_distribution}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:505
		{
			yyVAL.alter_distribution = &AlterDistribution{
				Element: &AttachRelation{
					Distribution: yyDollar[1].distribution_selector,
					Relations:    yyDollar[2].relations,
				},
			}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:514
		{
			yyVAL.alter_distribution = &AlterDistribution{
				Element: &DetachRelation{
					Distribution: yyDollar[1].distribution_selector,
					RelationName: yyDollar[4].str,
				},
			}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:526
		{
			yyVAL.dEntrieslist = append(yyDollar[1].dEntrieslist, yyDollar[3].distrKeyEntry)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:528
		{
			yyVAL.dEntrieslist = []DistributionKeyEntry{
				yyDollar[1].distrKeyEntry,
			}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:538
		{
			yyVAL.distrKeyEntry = DistributionKeyEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:547
		{
			yyVAL.distributed_relation = &DistributedRelation{
				Name:            yyDollar[2].str,
				DistributionKey: yyDollar[5].dEntrieslist,
			}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:554
		{
			yyVAL.distributed_relation = &DistributedRelation{
				Name:               yyDollar[2].str,
				ReplicatedRelation: true,
			}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:564
		{
			yyVAL.relations = []*DistributedRelation{yyDollar[1].distributed_relation}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:566
		{
			yyVAL.relations = append(yyDollar[1].relations, yyDollar[2].distributed_relation)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:571
		{
			yyVAL.relations = yyDollar[2].relations
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:577
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:582
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:587
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:592
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:597
		{
			yyVAL.create = &Create{
				Element: &ReferenceRelationDefinition{
					TableName: yyDollar[4].str,
				},
			}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:606
		{
			yyVAL.opt_asc_desc = &SortByAsc{}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:607
		{
			yyVAL.opt_asc_desc = &SortByDesc{}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:608
		{
			yyVAL.opt_asc_desc = &SortByDefault{}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:612
		{
			yyVAL.order_clause = &Order{Col: yyDollar[3].colref, OptAscDesc: yyDollar[4].opt_asc_desc}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:615
		{
			yyVAL.order_clause = OrderClause(nil)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:620
		{
			yyVAL.group_clause = GroupBy{Col: yyDollar[3].colref}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:623
		{
			yyVAL.group_clause = GroupByClauseEmpty{}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:628
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str, Where: yyDollar[3].where, GroupBy: yyDollar[4].group_clause, Order: yyDollar[5].order_clause}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:633
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:641
		{
			yyVAL.ds = &DistributionDefinition{
				ID:       yyDollar[2].str,
				ColTypes: yyDollar[3].strlist,
			}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:648
		{
			yyVAL.ds = &DistributionDefinition{
				Replicated: true,
			}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:656
		{
			yyVAL.strlist = yyDollar[3].strlist
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:658
		{
			/* empty column types should be prohibited */
			yyVAL.strlist = nil
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:664
		{
			yyVAL.strlist = append(yyDollar[1].strlist, yyDollar[3].str)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:666
		{
			yyVAL.strlist = []string{
				yyDollar[1].str,
			}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:673
		{
			yyVAL.str = "varchar"
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:675
		{
			yyVAL.str = "varchar hashed"
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:677
		{
			yyVAL.str = "integer"
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:679
		{
			yyVAL.str = "integer"
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:681
		{
			yyVAL.str = "uinteger"
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:683
		{
			yyVAL.str = "uinteger"
		}
	case 95:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:689
		{
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: yyDollar[3].str, TableName: yyDollar[4].str, Entries: yyDollar[5].entrieslist, Distribution: yyDollar[6].str}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:694
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: "shrule" + str, TableName: yyDollar[3].str, Entries: yyDollar[4].entrieslist, Distribution: yyDollar[5].str}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:703
		{
			yyVAL.entrieslist = make([]ShardingRuleEntry, 0)
			yyVAL.entrieslist = append(yyVAL.entrieslist, yyDollar[1].shruleEntry)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:709
		{
			yyVAL.entrieslist = append(yyDollar[1].entrieslist, yyDollar[2].shruleEntry)
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:715
		{
			yyVAL.shruleEntry = ShardingRuleEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:724
		{
			yyVAL.str = yyDollar[2].str
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:727
		{
			yyVAL.str = ""
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:731
		{
			yyVAL.str = yyDollar[2].str
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:736
		{
			yyVAL.str = yyDollar[2].str
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:742
		{
			yyVAL.str = "identity"
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:744
		{
			yyVAL.str = "murmur"
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:746
		{
			yyVAL.str = "city"
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:752
		{
			yyVAL.str = yyDollar[3].str
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:754
		{
			yyVAL.str = ""
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:759
		{
			yyVAL.str = yyDollar[3].str
		}
	case 110:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:761
		{
			yyVAL.str = "default"
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:766
		{
			yyVAL.bytes = []byte(yyDollar[1].str)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:769
		{
			buf := make([]byte, 8)
			binary.PutVarint(buf, int64(yyDollar[1].uinteger))
			yyVAL.bytes = buf
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:776
		{
			yyVAL.krbound = &KeyRangeBound{
				Pivots: [][]byte{
					yyDollar[1].bytes,
				},
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:783
		{
			yyVAL.krbound = &KeyRangeBound{
				Pivots: append(yyDollar[1].krbound.Pivots, yyDollar[3].bytes),
			}
		}
	case 115:
		yyDollar = yyS[yypt-9 : yypt+1]
//line gram.y:792
		{
			yyVAL.kr = &KeyRangeDefinition{
				KeyRangeID:   yyDollar[3].str,
				LowerBound:   yyDollar[5].krbound,
				ShardID:      yyDollar[8].str,
				Distribution: yyDollar[9].str,
			}
		}
	case 116:
		yyDollar = yyS[yypt-8 : yypt+1]
//line gram.y:801
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.kr = &KeyRangeDefinition{
				LowerBound:   yyDollar[4].krbound,
				ShardID:      yyDollar[7].str,
				Distribution: yyDollar[8].str,
				KeyRangeID:   "kr" + str,
			}
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:816
		{
			yyVAL.shard = &ShardDefinition{Id: yyDollar[2].str, Hosts: yyDollar[5].strlist}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:821
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.shard = &ShardDefinition{Id: "shard" + str, Hosts: yyDollar[4].strlist}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:831
		{
			yyVAL.strlist = []string{yyDollar[1].str}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:836
		{
			yyVAL.strlist = append(yyDollar[1].strlist, yyDollar[3].str)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:842
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:848
		{
			yyVAL.sharding_rule_selector = &ShardingRuleSelector{ID: yyDollar[3].str}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:854
		{
			yyVAL.key_range_selector = &KeyRangeSelector{KeyRangeID: yyDollar[3].str}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:860
		{
			yyVAL.distribution_selector = &DistributionSelector{ID: yyDollar[2].str, Replicated: false}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:862
		{
			yyVAL.distribution_selector = &DistributionSelector{Replicated: true}
		}
	case 126:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:868
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeFromID: yyDollar[4].str, Border: yyDollar[6].krbound}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:874
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str, Target: yyDollar[3].uinteger}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:877
		{
			yyVAL.kill = &Kill{Cmd: "client", Target: yyDollar[3].uinteger}
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:883
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str}
		}
	case 130:
		yyDollar = yyS[yypt-7 : yypt+1]
//line gram.y:889
		{
			yyVAL.redistribute = &RedistributeKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str, BatchSize: int(yyDollar[7].uinteger)}
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:892
		{
			yyVAL.redistribute = &RedistributeKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str, BatchSize: -1}
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:898
		{
			yyVAL.unite = &UniteKeyRange{KeyRangeIDL: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeIDR: yyDollar[4].str}
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:904
		{
			yyVAL.listen = &Listen{addr: yyDollar[2].str}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:910
		{
			yyVAL.shutdown = &Shutdown{}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:916
		{
			yyVAL.invalidate_cache = &InvalidateCache{}
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:924
		{
			yyVAL.register_router = &RegisterRouter{ID: yyDollar[3].str, Addr: yyDollar[5].str}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:930
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: yyDollar[3].str}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:935
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: `*`}
		}
	}
	goto yystack /* stack new state and value */
}
