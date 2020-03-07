package routingA

import (
	"github.com/gocarina/gocsv"
	"strings"
	"unicode"
)

const csv = `ID,",",',"""",(,),:,r,k,l,n,#,&,-,>,=,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z
0,,,,,,,r49,s11,,,r49,,,,,2,7,8,9,,10,,,,,,,,,,,,,1,,,,,,,
1,,,,,,,,,,,acc,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
2,,,,,,,s4,,,,r3,,,,,,,,,,,,,,,,,,,,,,3,,,,,,,,
3,,,,,,,,,,,r1,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
4,,,,,,,r49,s11,,,r49,,,,,5,7,8,9,,10,,,,,,,,,,,,,,,,,,,,
5,,,,,,,s4,,,,r3,,,,,,,,,,,,,,,,,,,,,,6,,,,,,,,
6,,,,,,,,,,,r2,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
7,,,,,,,r4,,,,r4,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
8,,,,,,,r5,,,,r5,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
9,,,,s18,,s12,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
10,,,,,,,,,,,,s51,r26,,,,,,,,,,,,,,,,,,,48,,,,,,,,,
11,,,,r30,,r30,r30,s55,s56,,r30,,,,r30,,,,,,,,,,,,,,,,,,,,47,,,,,,
12,,,,,,,,s11,,,,,,,,,,,14,13,,,,,,,,,,,,,,,,,,,,,
13,,,,,,,r6,,,,r6,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
14,,,,,,,r48,,,,r48,,,,s15,,,,,,,,,,,,,,,,,,,,,,,,,,
15,,,,,,,,s11,,,,,,,,,,,17,,16,,,,,,,,,,,,,,,,,,,,
16,,,,,,,r7,,,,r7,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
17,,,,s18,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
18,r17,s30,s31,,r17,,,s57,s58,s59,,s59,s59,s59,s59,,,,21,,,19,41,39,,,,,,38,,,,,,,,,,,
19,,,,,s20,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
20,,,,,,,r8,,,,r8,r8,r8,,,,,,,,,,,,,,,,,,,,,,,,,,,,
21,,,,,,s22,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
22,r17,s30,s31,,r17,r17,,s57,s58,s59,,s59,s59,s59,s59,,,,,,,,23,39,,,,,,38,,,,,,,,,,,
23,s25,,,,r13,,,,,,,,,,,,,,,,,,,,,,,,24,,,,,,,,,,,,
24,,,,,r14,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
25,r17,s30,s31,,r17,,,s57,s58,s59,,s59,s59,s59,s59,,,,,,,,26,39,,,,,,38,,,,,,,,,,,
26,,,,,,s27,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
27,r17,s30,s31,,r17,,,s57,s58,s59,,s59,s59,s59,s59,,,,,,,,28,39,,,,,,38,,,,,,,,,,,
28,s25,,,,r13,,,,,,,,,,,,,,,,,,,,,,,,29,,,,,,,,,,,,
29,,,,,r12,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
30,s63,r21,s63,s63,s63,s62,,s60,s61,s62,,s62,s62,s62,s62,,,,,,,,,,,,,,,,,,,,,34,,,,32,
31,s70,s70,r23,s70,s70,s69,,s67,s68,s69,,s69,s69,s69,s69,,,,,,,,,,,,,,,,,,,,,,65,,,,36
32,,s33,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
33,r18,,,,r18,r18,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
34,s63,r21,s63,s63,s63,s62,,s60,s61,s62,,s62,s62,s62,s62,,,,,,,,,,,,,,,,,,,,,34,,,,35,
35,,r20,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
36,,,s37,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
37,r19,,,,r19,r19,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
38,r15,,,,r15,r15,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
39,r17,,,,r17,r17,,s57,s58,s59,,s59,s59,s59,s59,,,,,,,,,39,,,,,,40,,,,,,,,,,,
40,r16,,,,r16,r16,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
41,[special1]略过逗号继续向后读，先碰到冒号是r11，先碰到逗号或右括号是s44,,,,r11,s22,,,,,,,,,,,,,,,,,,,,,,42,,,,,,,,,,,,,
42,s25,,,,r13,,,,,,,,,,,,,,,,,,,,,,,,43,,,,,,,,,,,,
43,,,,,r9,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
44,r17,s30,s31,,r17,r17,,s57,s58,s59,,s59,s59,s59,s59,,,,,,,,45,39,,,,,,38,,,,,,,,,,,
45,[special1]略过逗号继续向后读，先碰到冒号是r11，先碰到逗号或右括号是s44,,,,r11,,,,,,,,,,,,,,,,,,,,,,,46,,,,,,,,,,,,,
46,r10,,,,r10,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
47,,,,r27,,r27,r27,,,,r27,,,,r27,,,,,,,,,,,,,,,,,,,,,,,,,,
48,,,,,,,,,,,,,s49,,,,,,,,,,,,,,,,,,,,,,,,,,,,
49,,,,,,,,,,,,,,s50,,,,,,,,,,,,,,,,,,,,,,,,,,,
50,,,,,,,,s11,,,,,,,,,,,64,,,,,,,,,,,,,,,,,,,,,,
51,,,,,,,,,,,,s52,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
52,,,,,,,,s11,,,,,,,,,,,17,,53,,,,,,,,,,,,,,,,,,,,
53,,,,,,,,,,,,s51,r26,,,,,,,,,,,,,,,,,,,54,,,,,,,,,
54,,,,,,,,,,,,,r25,,,,,,,,,,,,,,,,,,,,,,,,,,,,
55,,,,r30,,r30,r30,s55,s56,,r30,,,,r30,,,,,,,,,,,,,,,,,,,,71,,,,,,
56,,,,r30,,r30,r30,s55,s56,,r30,,,,r30,,,,,,,,,,,,,,,,,,,,72,,,,,,
57,r39,,,,r39,r39,,r39,r39,r39,,r39,r39,r39,r39,,,,,,,,,,,,,,,,,,,,,,,,,,
58,r40,,,,r40,r40,,r40,r40,r40,,r40,r40,r40,r40,,,,,,,,,,,,,,,,,,,,,,,,,,
59,r41,,,,r41,r41,,r41,r41,r41,,r41,r41,r41,r41,,,,,,,,,,,,,,,,,,,,,,,,,,
60,r31,r31,r31,r31,r31,r31,,r31,r31,r31,,r31,r31,r31,r31,,,,,,,,,,,,,,,,,,,,,,,,,,
61,r32,r32,r32,r32,r32,r32,,r32,r32,r32,,r32,r32,r32,r32,,,,,,,,,,,,,,,,,,,,,,,,,,
62,r33,r33,r33,r33,r33,r33,,r33,r33,r33,,r33,r33,r33,r33,,,,,,,,,,,,,,,,,,,,,,,,,,
63,r34,r34,r34,r34,r34,r34,,r34,r34,r34,,r34,r34,r34,r34,,,,,,,,,,,,,,,,,,,,,,,,,,
64,,,,,,,r24,,,,r24,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
65,s70,s70,r23,s70,s70,s69,,s67,s68,s69,,s69,s69,s69,s69,,,,,,,,,,,,,,,,,,,,,,65,,,,66
66,,,r22,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
67,r35,r35,r35,r35,r35,r35,,r35,r35,r35,,r35,r35,r35,r35,,,,,,,,,,,,,,,,,,,,,,,,,,
68,r36,r36,r36,r36,r36,r36,,r36,r36,r36,,r36,r36,r36,r36,,,,,,,,,,,,,,,,,,,,,,,,,,
69,r37,r37,r37,r37,r37,r37,,r37,r37,r37,,r37,r37,r37,r37,,,,,,,,,,,,,,,,,,,,,,,,,,
70,r38,r38,r38,r38,r38,r38,,r38,r38,r38,,r38,r38,r38,r38,,,,,,,,,,,,,,,,,,,,,,,,,,
71,,,,r28,,r28,r28,,,,r28,,,,r28,,,,,,,,,,,,,,,,,,,,,,,,,,
72,,,,r29,,r29,r29,,,,r29,,,,r29,,,,,,,,,,,,,,,,,,,,,,,,,,
`

func read() (m []map[string]string) {
	m, _ = gocsv.CSVToMaps(strings.NewReader(csv))
	return
}

var table []map[string]string

func init() {
	table = read()
}

var pros = []string{
	/*0	*/ `*->S`,
	/*1	*/ `S->AR`,
	/*2	*/ "R->rAR",
	/*3	*/ `R->`,
	/*4	*/ `A->B`,
	/*5	*/ `A->C`,
	/*6	*/ `B->D:E`,
	/*7	*/ `E->D=F`,
	/*8	*/ `F->D(G)`,
	/*9	*/ `G->HMN`,
	/*10*/ `M->,HM`,
	/*11*/ `M->`,
	/*12*/ `N->,H:HN`,
	/*13*/ `N->`,
	/*14*/ `G->H:HN`,
	/*15*/ `H->O`,
	/*16*/ `O->IO`,
	/*17*/ `O->`,
	/*18*/ `H->'Y'`,
	/*19*/ `H->"Z"`,
	/*20*/ `Y->UY`,
	/*21*/ `Y->`,
	/*22*/ `Z->VZ`,
	/*23*/ `Z->`,
	/*24*/ `C->FQ->D`,
	/*25*/ `Q->&&FQ`,
	/*26*/ `Q->`,
	/*27*/ `D->kT`,
	/*28*/ `T->kT`,
	/*29*/ `T->lT`,
	/*30*/ `T->`,
	/*31*/ `U->k`,
	/*32*/ `U->l`,
	/*33*/ `U->n`,
	/*34*/ `U->y`,
	/*35*/ `V->k`,
	/*36*/ `V->l`,
	/*37*/ `V->n`,
	/*38*/ `V->z`,
	/*39*/ `I->k`,
	/*40*/ `I->l`,
	/*41*/ `I->n`,
	/*42*/ `k->a|b|...|Y|Z|...中|...|の|...`,
	/*43*/ `l->0|1|...|9`,
	/*44*/ `n->除了界符o(和:&->=)的标点符号`,
	/*45*/ `o->,|'|"|\(|\)`,
	/*46*/ `y->,|"|\(|\)`,
	/*47*/ `z->,|'|\(|\)`,
	/*48*/ `E->D`,
	/*49*/ "A->",
	/*50*/ "r->\n",
}

type production struct {
	left  rune
	right string
}

var productions []production

func init() {
	productions = make([]production, len(pros))
	for i, p := range pros {
		arr := strings.SplitN(p, "->", 2)
		if len(arr[0]) != 1 {
			continue
		}
		left := rune(arr[0][0])
		//过滤小写字母
		if unicode.IsLower(rune(arr[0][0])) {
			continue
		}
		productions [i] = production{
			left:  left,
			right: arr[1],
		}
	}
}
