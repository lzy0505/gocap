// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/lzy0505/gocap/parser/simple/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 82,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 56,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 57,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 53,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 13,
		Ignore: "",
	},
	ActionRow{ // S12
		Accept: 54,
		Ignore: "",
	},
	ActionRow{ // S13
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S14
		Accept: 55,
		Ignore: "",
	},
	ActionRow{ // S15
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S16
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S17
		Accept: 52,
		Ignore: "",
	},
	ActionRow{ // S18
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S19
		Accept: 80,
		Ignore: "",
	},
	ActionRow{ // S20
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S21
		Accept: 81,
		Ignore: "",
	},
	ActionRow{ // S22
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S23
		Accept: 32,
		Ignore: "",
	},
	ActionRow{ // S24
		Accept: 33,
		Ignore: "",
	},
	ActionRow{ // S25
		Accept: 59,
		Ignore: "",
	},
	ActionRow{ // S26
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S27
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S28
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S29
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S30
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S31
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S32
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S33
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S34
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S35
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S36
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S37
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S38
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S39
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S40
		Accept: 20,
		Ignore: "",
	},
	ActionRow{ // S41
		Accept: 58,
		Ignore: "",
	},
	ActionRow{ // S42
		Accept: 21,
		Ignore: "",
	},
	ActionRow{ // S43
		Accept: 83,
		Ignore: "",
	},
	ActionRow{ // S44
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S45
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S46
		Accept: 67,
		Ignore: "",
	},
	ActionRow{ // S47
		Accept: 74,
		Ignore: "",
	},
	ActionRow{ // S48
		Accept: 68,
		Ignore: "",
	},
	ActionRow{ // S49
		Accept: 62,
		Ignore: "",
	},
	ActionRow{ // S50
		Accept: 65,
		Ignore: "",
	},
	ActionRow{ // S51
		Accept: 77,
		Ignore: "",
	},
	ActionRow{ // S52
		Accept: 63,
		Ignore: "",
	},
	ActionRow{ // S53
		Accept: 78,
		Ignore: "",
	},
	ActionRow{ // S54
		Accept: 64,
		Ignore: "",
	},
	ActionRow{ // S55
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S56
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S57
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S58
		Accept: 66,
		Ignore: "",
	},
	ActionRow{ // S59
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S60
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S61
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S62
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S63
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S64
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S65
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S66
		Accept: 22,
		Ignore: "",
	},
	ActionRow{ // S67
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S68
		Accept: 76,
		Ignore: "",
	},
	ActionRow{ // S69
		Accept: 60,
		Ignore: "",
	},
	ActionRow{ // S70
		Accept: 84,
		Ignore: "",
	},
	ActionRow{ // S71
		Accept: 79,
		Ignore: "",
	},
	ActionRow{ // S72
		Accept: 85,
		Ignore: "",
	},
	ActionRow{ // S73
		Accept: 61,
		Ignore: "",
	},
	ActionRow{ // S74
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S75
		Accept: 70,
		Ignore: "",
	},
	ActionRow{ // S76
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S77
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S78
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S79
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S80
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S81
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S82
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S83
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S84
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S85
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S86
		Accept: 44,
		Ignore: "",
	},
	ActionRow{ // S87
		Accept: 46,
		Ignore: "",
	},
	ActionRow{ // S88
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S89
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S90
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S91
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S92
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S93
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S94
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S95
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S96
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S97
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S98
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S99
		Accept: 69,
		Ignore: "",
	},
	ActionRow{ // S100
		Accept: 75,
		Ignore: "",
	},
	ActionRow{ // S101
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S102
		Accept: 73,
		Ignore: "",
	},
	ActionRow{ // S103
		Accept: 14,
		Ignore: "",
	},
	ActionRow{ // S104
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S105
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S106
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S107
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S108
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S109
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S110
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S111
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S112
		Accept: 27,
		Ignore: "",
	},
	ActionRow{ // S113
		Accept: 26,
		Ignore: "",
	},
	ActionRow{ // S114
		Accept: 71,
		Ignore: "",
	},
	ActionRow{ // S115
		Accept: 72,
		Ignore: "",
	},
	ActionRow{ // S116
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S117
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S118
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S119
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S120
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S121
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S122
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S123
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S124
		Accept: 43,
		Ignore: "",
	},
	ActionRow{ // S125
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S126
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S127
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S128
		Accept: 28,
		Ignore: "",
	},
	ActionRow{ // S129
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S130
		Accept: 34,
		Ignore: "",
	},
	ActionRow{ // S131
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S132
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S133
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S134
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S135
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S136
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S137
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S138
		Accept: 18,
		Ignore: "",
	},
	ActionRow{ // S139
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S140
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S141
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S142
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S143
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S144
		Accept: 37,
		Ignore: "",
	},
	ActionRow{ // S145
		Accept: 25,
		Ignore: "",
	},
	ActionRow{ // S146
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S147
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S148
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S149
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S150
		Accept: 41,
		Ignore: "",
	},
	ActionRow{ // S151
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S152
		Accept: 12,
		Ignore: "",
	},
	ActionRow{ // S153
		Accept: 45,
		Ignore: "",
	},
	ActionRow{ // S154
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S155
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S156
		Accept: 23,
		Ignore: "",
	},
	ActionRow{ // S157
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S158
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S159
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S160
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S161
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S162
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S163
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S164
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S165
		Accept: 36,
		Ignore: "",
	},
	ActionRow{ // S166
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S167
		Accept: 16,
		Ignore: "",
	},
	ActionRow{ // S168
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S169
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S170
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S171
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S172
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S173
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S174
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S175
		Accept: 47,
		Ignore: "",
	},
	ActionRow{ // S176
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S177
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S178
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S179
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S180
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S181
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S182
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S183
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S184
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S185
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S186
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S187
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S188
		Accept: 48,
		Ignore: "",
	},
	ActionRow{ // S189
		Accept: 49,
		Ignore: "",
	},
	ActionRow{ // S190
		Accept: 29,
		Ignore: "",
	},
	ActionRow{ // S191
		Accept: 30,
		Ignore: "",
	},
	ActionRow{ // S192
		Accept: 50,
		Ignore: "",
	},
	ActionRow{ // S193
		Accept: 24,
		Ignore: "",
	},
	ActionRow{ // S194
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S195
		Accept: 39,
		Ignore: "",
	},
	ActionRow{ // S196
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S197
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S198
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S199
		Accept: 38,
		Ignore: "",
	},
	ActionRow{ // S200
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S201
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S202
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S203
		Accept: 31,
		Ignore: "",
	},
	ActionRow{ // S204
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S205
		Accept: 42,
		Ignore: "",
	},
}
