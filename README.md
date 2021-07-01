# 排课测试

基于班级时间段的选修行政班排课demo，采用多级策略的方式进行设计，支持自定义数据来源及排课算法

###Out:
```
    2021/07/01 16:30:52 Loop count added, now real loop count is: 1
    2021/07/01 16:30:52 NaN% Process:   0 Todo:   0 Unfinish:   0 Finish:   0 Total:   0
    2021/07/01 16:30:52 Loop count added, now real loop count is: 1
    2021/07/01 16:30:52 DP2 Chemistry HL DP2 Chemistry HL [教师13] => 4 Posh!
    2021/07/01 16:30:52 0.52% [DP2 Chemistry HL] DP2 Chemistry HL Process: 190 Todo:   0 Unfinish: 190 Finish:   1 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics-Statistic DP1 Mathematics-Statistic [教师6] => 23 Posh!
    2021/07/01 16:30:52 1.05% [DP1 Mathematics-Statistic] DP1 Mathematics-Statistic Process: 189 Todo:   0 Unfinish: 189 Finish:   2 Total: 191
    2021/07/01 16:30:52 DP2 English B HL DP2 English B HL [教师22] => 10 Posh!
    2021/07/01 16:30:52 1.57% [DP2 English B HL] DP2 English B HL Process: 188 Todo:   0 Unfinish: 188 Finish:   3 Total: 191
    2021/07/01 16:30:52 DP1 Biology HL DP1 Biology HL [教师18] => 14 Posh!
    2021/07/01 16:30:52 2.09% [DP1 Biology HL] DP1 Biology HL Process: 187 Todo:   0 Unfinish: 187 Finish:   4 Total: 191
    2021/07/01 16:30:52 DP1 Economics DP1 Economics [教师1] => 27 Posh!
    2021/07/01 16:30:52 2.62% [DP1 Economics] DP1 Economics Process: 186 Todo:   0 Unfinish: 186 Finish:   5 Total: 191
    2021/07/01 16:30:52 DP1 Sports Exercise and Health Science SL DP1 Sports Exercise and Health Science SL [教师24] => 27 Posh!
    2021/07/01 16:30:52 3.14% [DP1 Sports Exercise and Health Science SL] DP1 Sports Exercise and Health Science SL Process: 185 Todo:   0 Unfinish: 185 Finish:   6 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics-Statistic DP1 Mathematics-Statistic [教师6] => 28 Posh!
    2021/07/01 16:30:52 3.66% [DP1 Mathematics-Statistic] DP1 Mathematics-Statistic Process: 184 Todo:   0 Unfinish: 184 Finish:   7 Total: 191
    2021/07/01 16:30:52 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 15 Posh!
    2021/07/01 16:30:52 4.19% [DP1 Economics HL+SL] DP1 Economics HL+SL Process: 183 Todo:   0 Unfinish: 183 Finish:   8 Total: 191
    2021/07/01 16:30:52 DP1 Economics HL DP1 Economics HL [教师11] => 26 Posh!
    2021/07/01 16:30:52 4.71% [DP1 Economics HL] DP1 Economics HL Process: 182 Todo:   0 Unfinish: 182 Finish:   9 Total: 191
    2021/07/01 16:30:52 DP1 English DP1 English [教师17] => 29 Posh!
    2021/07/01 16:30:52 5.24% [DP1 English] DP1 English Process: 181 Todo:   0 Unfinish: 181 Finish:  10 Total: 191
    2021/07/01 16:30:52 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 46 Posh!
    2021/07/01 16:30:52 5.76% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process: 180 Todo:   0 Unfinish: 180 Finish:  11 Total: 191
    2021/07/01 16:30:52 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 48 Posh!
    2021/07/01 16:30:52 6.28% [DP2 Environmental Systems & Societies SL] DP2 Environmental Systems & Societies SL Process: 179 Todo:   0 Unfinish: 179 Finish:  12 Total: 191
    2021/07/01 16:30:52 DP1 Economics HL DP1 Economics HL [教师11] => 10 Posh!
    2021/07/01 16:30:52 6.81% [DP1 Economics HL] DP1 Economics HL Process: 178 Todo:   0 Unfinish: 178 Finish:  13 Total: 191
    2021/07/01 16:30:52 DP1 Physics HL DP1 Physics HL [教师15] => 9 Posh!
    2021/07/01 16:30:52 7.33% [DP1 Physics HL] DP1 Physics HL Process: 177 Todo:   0 Unfinish: 177 Finish:  14 Total: 191
    2021/07/01 16:30:52 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 29 Posh!
    2021/07/01 16:30:52 7.85% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process: 176 Todo:   0 Unfinish: 176 Finish:  15 Total: 191
    2021/07/01 16:30:52 DP2 English B SL DP2 English B SL [教师16] => 31 Posh!
    2021/07/01 16:30:52 8.38% [DP2 English B SL] DP2 English B SL Process: 175 Todo:   0 Unfinish: 175 Finish:  16 Total: 191
    2021/07/01 16:30:52 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 18 Posh!
    2021/07/01 16:30:52 8.90% [DP2 Business Management HL+SL] DP2 Business Management HL+SL Process: 174 Todo:   0 Unfinish: 174 Finish:  17 Total: 191
    2021/07/01 16:30:52 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 51 Posh!
    2021/07/01 16:30:52 9.42% [DP1 Economics HL+SL] DP1 Economics HL+SL Process: 173 Todo:   0 Unfinish: 173 Finish:  18 Total: 191
    2021/07/01 16:30:52 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 7 Posh!
    2021/07/01 16:30:52 9.95% [DP1 Chinese A Language and Literature HL] DP1 Chinese A Language and Literature HL Process: 172 Todo:   0 Unfinish: 172 Finish:  19 Total: 191
    2021/07/01 16:30:52 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 12 Posh!
    2021/07/01 16:30:52 10.47% [DP1 Biology HL+SL] DP1 Biology HL+SL Process: 171 Todo:   0 Unfinish: 171 Finish:  20 Total: 191
    2021/07/01 16:30:52 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 4 Posh!
    2021/07/01 16:30:52 10.99% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process: 170 Todo:   0 Unfinish: 170 Finish:  21 Total: 191
    2021/07/01 16:30:52 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 18 Posh!
    2021/07/01 16:30:52 11.52% [DP2 Japanese ab initio] DP2 Japanese ab initio Process: 169 Todo:   0 Unfinish: 169 Finish:  22 Total: 191
    2021/07/01 16:30:52 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 37 Posh!
    2021/07/01 16:30:52 12.04% [DP1 Physics HL+SL] DP1 Physics HL+SL Process: 168 Todo:   0 Unfinish: 168 Finish:  23 Total: 191
    2021/07/01 16:30:52 DP1 Psychology HL DP1 Psychology HL [教师25] => 9 Posh!
    2021/07/01 16:30:52 12.57% [DP1 Psychology HL] DP1 Psychology HL Process: 167 Todo:   0 Unfinish: 167 Finish:  24 Total: 191
    2021/07/01 16:30:52 DP1 Sports Exercise and Health Science SL DP1 Sports Exercise and Health Science SL [教师24] => 13 Posh!
    2021/07/01 16:30:52 13.09% [DP1 Sports Exercise and Health Science SL] DP1 Sports Exercise and Health Science SL Process: 166 Todo:   0 Unfinish: 166 Finish:  25 Total: 191
    2021/07/01 16:30:52 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 34 Posh!
    2021/07/01 16:30:52 13.61% [DP1 Economics HL+SL] DP1 Economics HL+SL Process: 165 Todo:   0 Unfinish: 165 Finish:  26 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 19 Posh!
    2021/07/01 16:30:52 14.14% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process: 164 Todo:   0 Unfinish: 164 Finish:  27 Total: 191
    2021/07/01 16:30:52 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 49 Posh!
    2021/07/01 16:30:52 14.66% [DP2 Business Management HL+SL] DP2 Business Management HL+SL Process: 163 Todo:   0 Unfinish: 163 Finish:  28 Total: 191
    2021/07/01 16:30:52 DP2 Mathematics Analysis and Approaches SL DP2 Mathematics Analysis and Approaches SL [教师14] => 30 Posh!
    2021/07/01 16:30:52 15.18% [DP2 Mathematics Analysis and Approaches SL] DP2 Mathematics Analysis and Approaches SL Process: 162 Todo:   0 Unfinish: 162 Finish:  29 Total: 191
    2021/07/01 16:30:52 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 33 Posh!
    2021/07/01 16:30:52 15.71% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process: 161 Todo:   0 Unfinish: 161 Finish:  30 Total: 191
    2021/07/01 16:30:52 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 7 Posh!
    2021/07/01 16:30:52 16.23% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process: 160 Todo:   0 Unfinish: 160 Finish:  31 Total: 191
    2021/07/01 16:30:52 DP2 Biology HL+SL DP2 Biology HL+SL [教师27] => 12 Posh!
    2021/07/01 16:30:52 16.75% [DP2 Biology HL+SL] DP2 Biology HL+SL Process: 159 Todo:   0 Unfinish: 159 Finish:  32 Total: 191
    2021/07/01 16:30:52 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 36 Posh!
    2021/07/01 16:30:52 17.28% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process: 158 Todo:   0 Unfinish: 158 Finish:  33 Total: 191
    2021/07/01 16:30:52 DP2 Chinese A Literature HL DP2 Chinese A Literature HL [教师10] => 27 Posh!
    2021/07/01 16:30:52 17.80% [DP2 Chinese A Literature HL] DP2 Chinese A Literature HL Process: 157 Todo:   0 Unfinish: 157 Finish:  34 Total: 191
    2021/07/01 16:30:52 DP2 Economics HL+SL DP2 Economics HL+SL [教师1] => 3 Posh!
    2021/07/01 16:30:52 18.32% [DP2 Economics HL+SL] DP2 Economics HL+SL Process: 156 Todo:   0 Unfinish: 156 Finish:  35 Total: 191
    2021/07/01 16:30:52 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 25 Posh!
    2021/07/01 16:30:52 18.85% [DP1 Physics HL+SL] DP1 Physics HL+SL Process: 155 Todo:   0 Unfinish: 155 Finish:  36 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics-Statistic DP1 Mathematics-Statistic [教师6] => 50 Posh!
    2021/07/01 16:30:52 19.37% [DP1 Mathematics-Statistic] DP1 Mathematics-Statistic Process: 154 Todo:   0 Unfinish: 154 Finish:  37 Total: 191
    2021/07/01 16:30:52 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 38 Posh!
    2021/07/01 16:30:52 19.90% [DP2 Physics HL+SL] DP2 Physics HL+SL Process: 153 Todo:   0 Unfinish: 153 Finish:  38 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 15 Posh!
    2021/07/01 16:30:52 20.42% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process: 152 Todo:   0 Unfinish: 152 Finish:  39 Total: 191
    2021/07/01 16:30:52 DP1 Physics HL DP1 Physics HL [教师15] => 26 Posh!
    2021/07/01 16:30:52 20.94% [DP1 Physics HL] DP1 Physics HL Process: 151 Todo:   0 Unfinish: 151 Finish:  40 Total: 191
    2021/07/01 16:30:52 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 48 Posh!
    2021/07/01 16:30:52 21.47% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process: 150 Todo:   0 Unfinish: 150 Finish:  41 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 17 Posh!
    2021/07/01 16:30:52 21.99% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process: 149 Todo:   0 Unfinish: 149 Finish:  42 Total: 191
    2021/07/01 16:30:52 DP1 English B HL DP1 English B HL [教师16] => 42 Posh!
    2021/07/01 16:30:52 22.51% [DP1 English B HL] DP1 English B HL Process: 148 Todo:   0 Unfinish: 148 Finish:  43 Total: 191
    2021/07/01 16:30:52 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 11 Posh!
    2021/07/01 16:30:52 23.04% [DP2 Physics HL+SL] DP2 Physics HL+SL Process: 147 Todo:   0 Unfinish: 147 Finish:  44 Total: 191
    2021/07/01 16:30:52 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 6 Posh!
    2021/07/01 16:30:52 23.56% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process: 146 Todo:   0 Unfinish: 146 Finish:  45 Total: 191
    2021/07/01 16:30:52 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 32 Posh!
    2021/07/01 16:30:52 24.08% [DP2 Business Management HL+SL] DP2 Business Management HL+SL Process: 145 Todo:   0 Unfinish: 145 Finish:  46 Total: 191
    2021/07/01 16:30:52 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 45 Posh!
    2021/07/01 16:30:52 24.61% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process: 144 Todo:   0 Unfinish: 144 Finish:  47 Total: 191
    2021/07/01 16:30:52 DP1 English B HL DP1 English B HL [教师16] => 45 Posh!
    2021/07/01 16:30:52 25.13% [DP1 English B HL] DP1 English B HL Process: 143 Todo:   0 Unfinish: 143 Finish:  48 Total: 191
    2021/07/01 16:30:52 DP1 English B HL DP1 English B HL [教师16] => 40 Posh!
    2021/07/01 16:30:52 25.65% [DP1 English B HL] DP1 English B HL Process: 142 Todo:   0 Unfinish: 142 Finish:  49 Total: 191
    2021/07/01 16:30:52 DP2 English B SL DP2 English B SL [教师16] => 52 Posh!
    2021/07/01 16:30:52 26.18% [DP2 English B SL] DP2 English B SL Process: 141 Todo:   0 Unfinish: 141 Finish:  50 Total: 191
    2021/07/01 16:30:52 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 8 Posh!
    2021/07/01 16:30:52 26.70% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process: 140 Todo:   0 Unfinish: 140 Finish:  51 Total: 191
    2021/07/01 16:30:52 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 50 Posh!
    2021/07/01 16:30:52 27.23% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process: 139 Todo:   0 Unfinish: 139 Finish:  52 Total: 191
    2021/07/01 16:30:52 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 17 Posh!
    2021/07/01 16:30:52 27.75% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process: 138 Todo:   0 Unfinish: 138 Finish:  53 Total: 191
    2021/07/01 16:30:52 DP1 Biology DP1 Biology [教师18] => 19 Posh!
    2021/07/01 16:30:52 28.27% [DP1 Biology] DP1 Biology Process: 137 Todo:   0 Unfinish: 137 Finish:  54 Total: 191
    2021/07/01 16:30:52 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 30 Posh!
    2021/07/01 16:30:52 28.80% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process: 136 Todo:   0 Unfinish: 136 Finish:  55 Total: 191
    2021/07/01 16:30:53 DP2 English TOK DP2 English TOK [教师22] => 37 Posh!
    2021/07/01 16:30:53 29.32% [DP2 English TOK] DP2 English TOK Process: 135 Todo:   0 Unfinish: 135 Finish:  56 Total: 191
    2021/07/01 16:30:53 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 24 Posh!
    2021/07/01 16:30:53 29.84% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process: 134 Todo:   0 Unfinish: 134 Finish:  57 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL DP1 Chinese A Literature HL [教师2] => 29 Posh!
    2021/07/01 16:30:53 30.37% [DP1 Chinese A Literature HL] DP1 Chinese A Literature HL Process: 133 Todo:   0 Unfinish: 133 Finish:  58 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 16 Posh!
    2021/07/01 16:30:53 30.89% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process: 132 Todo:   0 Unfinish: 132 Finish:  59 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 43 Posh!
    2021/07/01 16:30:53 31.41% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process: 131 Todo:   0 Unfinish: 131 Finish:  60 Total: 191
    2021/07/01 16:30:53 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 10 Posh!
    2021/07/01 16:30:53 31.94% [DP2 Japanese ab initio] DP2 Japanese ab initio Process: 130 Todo:   0 Unfinish: 130 Finish:  61 Total: 191
    2021/07/01 16:30:53 DP2 English B SL DP2 English B SL [教师16] => 15 Posh!
    2021/07/01 16:30:53 32.46% [DP2 English B SL] DP2 English B SL Process: 129 Todo:   0 Unfinish: 129 Finish:  62 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 31 Posh!
    2021/07/01 16:30:53 32.98% [DP1 Economics] DP1 Economics Process: 128 Todo:   0 Unfinish: 128 Finish:  63 Total: 191
    2021/07/01 16:30:53 DP1 English DP1 English [教师17] => 9 Posh!
    2021/07/01 16:30:53 33.51% [DP1 English] DP1 English Process: 127 Todo:   0 Unfinish: 127 Finish:  64 Total: 191
    2021/07/01 16:30:53 DP2 English B HL DP2 English B HL [教师22] => 25 Posh!
    2021/07/01 16:30:53 34.03% [DP2 English B HL] DP2 English B HL Process: 126 Todo:   0 Unfinish: 126 Finish:  65 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 13 Posh!
    2021/07/01 16:30:53 34.55% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process: 125 Todo:   0 Unfinish: 125 Finish:  66 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 52 Posh!
    2021/07/01 16:30:53 35.08% [DP1 Business Management HL] DP1 Business Management HL Process: 124 Todo:   0 Unfinish: 124 Finish:  67 Total: 191
    2021/07/01 16:30:53 DP2 Physics HL DP2 Physics HL [教师20] => 14 Posh!
    2021/07/01 16:30:53 35.60% [DP2 Physics HL] DP2 Physics HL Process: 123 Todo:   0 Unfinish: 123 Finish:  68 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 41 Posh!
    2021/07/01 16:30:53 36.13% [DP1 Chinese A Language and Literature HL] DP1 Chinese A Language and Literature HL Process: 122 Todo:   0 Unfinish: 122 Finish:  69 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 39 Posh!
    2021/07/01 16:30:53 36.65% [DP1 Business Management HL] DP1 Business Management HL Process: 121 Todo:   0 Unfinish: 121 Finish:  70 Total: 191
    2021/07/01 16:30:53 DP1 Critical thinking and writing (CTW) DP1 Critical thinking and writing (CTW) [教师4] => 3 Posh!
    2021/07/01 16:30:53 37.17% [DP1 Critical thinking and writing (CTW)] DP1 Critical thinking and writing (CTW) Process: 120 Todo:   0 Unfinish: 120 Finish:  71 Total: 191
    2021/07/01 16:30:53 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 21 Posh!
    2021/07/01 16:30:53 37.70% [DP2 Business Management HL+SL] DP2 Business Management HL+SL Process: 119 Todo:   0 Unfinish: 119 Finish:  72 Total: 191
    2021/07/01 16:30:53 DP2 Psychology HL DP2 Psychology HL [教师25] => 20 Posh!
    2021/07/01 16:30:53 38.22% [DP2 Psychology HL] DP2 Psychology HL Process: 118 Todo:   0 Unfinish: 118 Finish:  73 Total: 191
    2021/07/01 16:30:53 DP2 Business Management HL DP2 Business Management HL [教师8] => 40 Posh!
    2021/07/01 16:30:53 38.74% [DP2 Business Management HL] DP2 Business Management HL Process: 117 Todo:   0 Unfinish: 117 Finish:  74 Total: 191
    2021/07/01 16:30:53 DP2 Economics HL DP2 Economics HL [教师1] => 5 Posh!
    2021/07/01 16:30:53 39.27% [DP2 Economics HL] DP2 Economics HL Process: 116 Todo:   0 Unfinish: 116 Finish:  75 Total: 191
    2021/07/01 16:30:53 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 38 Posh!
    2021/07/01 16:30:53 39.79% [DP1 Economics HL+SL] DP1 Economics HL+SL Process: 115 Todo:   0 Unfinish: 115 Finish:  76 Total: 191
    2021/07/01 16:30:53 DP1 English DP1 English [教师17] => 45 Posh!
    2021/07/01 16:30:53 40.31% [DP1 English] DP1 English Process: 114 Todo:   0 Unfinish: 114 Finish:  77 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 46 Posh!
    2021/07/01 16:30:53 40.84% [DP1 Business Management HL] DP1 Business Management HL Process: 113 Todo:   0 Unfinish: 113 Finish:  78 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 35 Posh!
    2021/07/01 16:30:53 41.36% [DP1 Economics] DP1 Economics Process: 112 Todo:   0 Unfinish: 112 Finish:  79 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 28 Posh!
    2021/07/01 16:30:53 41.88% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process: 111 Todo:   0 Unfinish: 111 Finish:  80 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 23 Posh!
    2021/07/01 16:30:53 42.41% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process: 110 Todo:   0 Unfinish: 110 Finish:  81 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 20 Posh!
    2021/07/01 16:30:53 42.93% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process: 109 Todo:   0 Unfinish: 109 Finish:  82 Total: 191
    2021/07/01 16:30:53 DP1 Biology HL DP1 Biology HL [教师18] => 24 Posh!
    2021/07/01 16:30:53 43.46% [DP1 Biology HL] DP1 Biology HL Process: 108 Todo:   0 Unfinish: 108 Finish:  83 Total: 191
    2021/07/01 16:30:53 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 23 Posh!
    2021/07/01 16:30:53 43.98% [DP2 Environmental Systems & Societies SL] DP2 Environmental Systems & Societies SL Process: 107 Todo:   0 Unfinish: 107 Finish:  84 Total: 191
    2021/07/01 16:30:53 DP1 Psychology HL DP1 Psychology HL [教师25] => 49 Posh!
    2021/07/01 16:30:53 44.50% [DP1 Psychology HL] DP1 Psychology HL Process: 106 Todo:   0 Unfinish: 106 Finish:  85 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 32 Posh!
    2021/07/01 16:30:53 45.03% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process: 105 Todo:   0 Unfinish: 105 Finish:  86 Total: 191
    2021/07/01 16:30:53 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 39 Posh!
    2021/07/01 16:30:53 45.55% [DP2 Physics HL+SL] DP2 Physics HL+SL Process: 104 Todo:   0 Unfinish: 104 Finish:  87 Total: 191
    2021/07/01 16:30:53 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 43 Posh!
    2021/07/01 16:30:53 46.07% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process: 103 Todo:   0 Unfinish: 103 Finish:  88 Total: 191
    2021/07/01 16:30:53 DP1 English B SL DP1 English B SL [教师21] => 3 Posh!
    2021/07/01 16:30:53 46.60% [DP1 English B SL] DP1 English B SL Process: 102 Todo:   0 Unfinish: 102 Finish:  89 Total: 191
    2021/07/01 16:30:53 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 16 Posh!
    2021/07/01 16:30:53 47.12% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process: 101 Todo:   0 Unfinish: 101 Finish:  90 Total: 191
    2021/07/01 16:30:53 DP2 English B HL DP2 English B HL [教师22] => 46 Posh!
    2021/07/01 16:30:53 47.64% [DP2 English B HL] DP2 English B HL Process: 100 Todo:   0 Unfinish: 100 Finish:  91 Total: 191
    2021/07/01 16:30:53 DP2 Economics HL DP2 Economics HL [教师1] => 42 Posh!
    2021/07/01 16:30:53 48.17% [DP2 Economics HL] DP2 Economics HL Process:  99 Todo:   0 Unfinish:  99 Finish:  92 Total: 191
    2021/07/01 16:30:53 DP1 English B HL DP1 English B HL [教师16] => 33 Posh!
    2021/07/01 16:30:53 48.69% [DP1 English B HL] DP1 English B HL Process:  98 Todo:   0 Unfinish:  98 Finish:  93 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 5 Posh!
    2021/07/01 16:30:53 49.21% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process:  97 Todo:   0 Unfinish:  97 Finish:  94 Total: 191
    2021/07/01 16:30:53 DP1 Critical thinking and writing (CTW) DP1 Critical thinking and writing (CTW) [教师4] => 19 Posh!
    2021/07/01 16:30:53 49.74% [DP1 Critical thinking and writing (CTW)] DP1 Critical thinking and writing (CTW) Process:  96 Todo:   0 Unfinish:  96 Finish:  95 Total: 191
    2021/07/01 16:30:53 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 31 Posh!
    2021/07/01 16:30:53 50.26% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:  95 Todo:   0 Unfinish:  95 Finish:  96 Total: 191
    2021/07/01 16:30:53 DP1 English DP1 English [教师17] => 17 Posh!
    2021/07/01 16:30:53 50.79% [DP1 English] DP1 English Process:  94 Todo:   0 Unfinish:  94 Finish:  97 Total: 191
    2021/07/01 16:30:53 DP2 Economics HL+SL DP2 Economics HL+SL [教师1] => 34 Posh!
    2021/07/01 16:30:53 51.31% [DP2 Economics HL+SL] DP2 Economics HL+SL Process:  93 Todo:   0 Unfinish:  93 Finish:  98 Total: 191
    2021/07/01 16:30:53 DP1 English B HL DP1 English B HL [教师16] => 18 Posh!
    2021/07/01 16:30:53 51.83% [DP1 English B HL] DP1 English B HL Process:  92 Todo:   0 Unfinish:  92 Finish:  99 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches SL DP2 Mathematics Analysis and Approaches SL [教师14] => 13 Posh!
    2021/07/01 16:30:53 52.36% [DP2 Mathematics Analysis and Approaches SL] DP2 Mathematics Analysis and Approaches SL Process:  91 Todo:   0 Unfinish:  91 Finish: 100 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 47 Posh!
    2021/07/01 16:30:53 52.88% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:  90 Todo:   0 Unfinish:  90 Finish: 101 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 14 Posh!
    2021/07/01 16:30:53 53.40% [DP1 Visual Arts HL] DP1 Visual Arts HL Process:  89 Todo:   0 Unfinish:  89 Finish: 102 Total: 191
    2021/07/01 16:30:53 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 26 Posh!
    2021/07/01 16:30:53 53.93% [DP2 Japanese ab initio] DP2 Japanese ab initio Process:  88 Todo:   0 Unfinish:  88 Finish: 103 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 20 Posh!
    2021/07/01 16:30:53 54.45% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  87 Todo:   0 Unfinish:  87 Finish: 104 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 16 Posh!
    2021/07/01 16:30:53 54.97% [DP1 Business Management HL] DP1 Business Management HL Process:  86 Todo:   0 Unfinish:  86 Finish: 105 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Language and Literature SL DP2 Chinese A Language and Literature SL [教师12] => 28 Posh!
    2021/07/01 16:30:53 55.50% [DP2 Chinese A Language and Literature SL] DP2 Chinese A Language and Literature SL Process:  85 Todo:   0 Unfinish:  85 Finish: 106 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 17 Posh!
    2021/07/01 16:30:53 56.02% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  84 Todo:   0 Unfinish:  84 Finish: 107 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 35 Posh!
    2021/07/01 16:30:53 56.54% [DP1 Visual Arts HL] DP1 Visual Arts HL Process:  83 Todo:   0 Unfinish:  83 Finish: 108 Total: 191
    2021/07/01 16:30:53 DP1 English B SL DP1 English B SL [教师21] => 40 Posh!
    2021/07/01 16:30:53 57.07% [DP1 English B SL] DP1 English B SL Process:  82 Todo:   0 Unfinish:  82 Finish: 109 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 41 Posh!
    2021/07/01 16:30:53 57.59% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process:  81 Todo:   0 Unfinish:  81 Finish: 110 Total: 191
    2021/07/01 16:30:53 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 38 Posh!
    2021/07/01 16:30:53 58.12% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:  80 Todo:   0 Unfinish:  80 Finish: 111 Total: 191
    2021/07/01 16:30:53 DP2 Physics HL DP2 Physics HL [教师20] => 35 Posh!
    2021/07/01 16:30:53 58.64% [DP2 Physics HL] DP2 Physics HL Process:  79 Todo:   0 Unfinish:  79 Finish: 112 Total: 191
    2021/07/01 16:30:53 DP2 English TOK DP2 English TOK [教师22] => 47 Posh!
    2021/07/01 16:30:53 59.16% [DP2 English TOK] DP2 English TOK Process:  78 Todo:   0 Unfinish:  78 Finish: 113 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Literature HL+SL DP2 Chinese A Literature HL+SL [教师10] => 28 Posh!
    2021/07/01 16:30:53 59.69% [DP2 Chinese A Literature HL+SL] DP2 Chinese A Literature HL+SL Process:  77 Todo:   0 Unfinish:  77 Finish: 114 Total: 191
    2021/07/01 16:30:53 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 52 Posh!
    2021/07/01 16:30:53 60.21% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process:  76 Todo:   0 Unfinish:  76 Finish: 115 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL DP1 Chemistry HL [教师7] => 7 Posh!
    2021/07/01 16:30:53 60.73% [DP1 Chemistry HL] DP1 Chemistry HL Process:  75 Todo:   0 Unfinish:  75 Finish: 116 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 30 Posh!
    2021/07/01 16:30:53 61.26% [DP1 Business Management HL] DP1 Business Management HL Process:  74 Todo:   0 Unfinish:  74 Finish: 117 Total: 191
    2021/07/01 16:30:53 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 27 Posh!
    2021/07/01 16:30:53 61.78% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:  73 Todo:   0 Unfinish:  73 Finish: 118 Total: 191
    2021/07/01 16:30:53 DP1 Biology DP1 Biology [教师18] => 4 Posh!
    2021/07/01 16:30:53 62.30% [DP1 Biology] DP1 Biology Process:  72 Todo:   0 Unfinish:  72 Finish: 119 Total: 191
    2021/07/01 16:30:53 DP1 Biology DP1 Biology [教师18] => 7 Posh!
    2021/07/01 16:30:53 62.83% [DP1 Biology] DP1 Biology Process:  71 Todo:   0 Unfinish:  71 Finish: 120 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 19 Posh!
    2021/07/01 16:30:53 63.35% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:  70 Todo:   0 Unfinish:  70 Finish: 121 Total: 191
    2021/07/01 16:30:53 DP2 English B SL DP2 English B SL [教师16] => 46 Posh!
    2021/07/01 16:30:53 63.87% [DP2 English B SL] DP2 English B SL Process:  69 Todo:   0 Unfinish:  69 Finish: 122 Total: 191
    2021/07/01 16:30:53 DP2 Chemistry HL DP2 Chemistry HL [教师13] => 6 Posh!
    2021/07/01 16:30:53 64.40% [DP2 Chemistry HL] DP2 Chemistry HL Process:  68 Todo:   0 Unfinish:  68 Finish: 123 Total: 191
    2021/07/01 16:30:53 DP2 Biology HL DP2 Biology HL [教师27] => 49 Posh!
    2021/07/01 16:30:53 64.92% [DP2 Biology HL] DP2 Biology HL Process:  67 Todo:   0 Unfinish:  67 Finish: 124 Total: 191
    2021/07/01 16:30:53 DP1 Sports Exercise and Health Science SL DP1 Sports Exercise and Health Science SL [教师24] => 29 Posh!
    2021/07/01 16:30:53 65.45% [DP1 Sports Exercise and Health Science SL] DP1 Sports Exercise and Health Science SL Process:  66 Todo:   0 Unfinish:  66 Finish: 125 Total: 191
    2021/07/01 16:30:53 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 9 Posh!
    2021/07/01 16:30:53 65.97% [DP2 Environmental Systems & Societies SL] DP2 Environmental Systems & Societies SL Process:  65 Todo:   0 Unfinish:  65 Finish: 126 Total: 191
    2021/07/01 16:30:53 DP1 Sports Exercise and Health Science SL DP1 Sports Exercise and Health Science SL [教师24] => 23 Posh!
    2021/07/01 16:30:53 66.49% [DP1 Sports Exercise and Health Science SL] DP1 Sports Exercise and Health Science SL Process:  64 Todo:   0 Unfinish:  64 Finish: 127 Total: 191
    2021/07/01 16:30:53 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 31 Posh!
    2021/07/01 16:30:53 67.02% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:  63 Todo:   0 Unfinish:  63 Finish: 128 Total: 191
    2021/07/01 16:30:53 DP2 Psychology HL DP2 Psychology HL [教师25] => 5 Posh!
    2021/07/01 16:30:53 67.54% [DP2 Psychology HL] DP2 Psychology HL Process:  62 Todo:   0 Unfinish:  62 Finish: 129 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Literature HL+SL DP2 Chinese A Literature HL+SL [教师10] => 50 Posh!
    2021/07/01 16:30:53 68.06% [DP2 Chinese A Literature HL+SL] DP2 Chinese A Literature HL+SL Process:  61 Todo:   0 Unfinish:  61 Finish: 130 Total: 191
    2021/07/01 16:30:53 DP1 English B HL DP1 English B HL [教师16] => 6 Posh!
    2021/07/01 16:30:53 68.59% [DP1 English B HL] DP1 English B HL Process:  60 Todo:   0 Unfinish:  60 Finish: 131 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 18 Posh!
    2021/07/01 16:30:53 69.11% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process:  59 Todo:   0 Unfinish:  59 Finish: 132 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL DP1 Chinese A Literature HL [教师2] => 9 Posh!
    2021/07/01 16:30:53 69.63% [DP1 Chinese A Literature HL] DP1 Chinese A Literature HL Process:  58 Todo:   0 Unfinish:  58 Finish: 133 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches SL DP2 Mathematics Analysis and Approaches SL [教师14] => 33 Posh!
    2021/07/01 16:30:53 70.16% [DP2 Mathematics Analysis and Approaches SL] DP2 Mathematics Analysis and Approaches SL Process:  57 Todo:   0 Unfinish:  57 Finish: 134 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 46 Posh!
    2021/07/01 16:30:53 70.68% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process:  56 Todo:   0 Unfinish:  56 Finish: 135 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 41 Posh!
    2021/07/01 16:30:53 71.20% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:  55 Todo:   0 Unfinish:  55 Finish: 136 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 13 Posh!
    2021/07/01 16:30:53 71.73% [DP1 Economics] DP1 Economics Process:  54 Todo:   0 Unfinish:  54 Finish: 137 Total: 191
    2021/07/01 16:30:53 DP2 Biology HL DP2 Biology HL [教师27] => 40 Posh!
    2021/07/01 16:30:53 72.25% [DP2 Biology HL] DP2 Biology HL Process:  53 Todo:   0 Unfinish:  53 Finish: 138 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 20 Pop!
    2021/07/01 16:30:53 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 20 Posh!
    2021/07/01 16:30:53 72.25% [DP1 Physics HL+SL] DP1 Physics HL+SL Process:  52 Todo:   1 Unfinish:  53 Finish: 138 Total: 191
    2021/07/01 16:30:53 DP2 Biology HL+SL DP2 Biology HL+SL [教师27] => 36 Posh!
    2021/07/01 16:30:53 72.77% [DP2 Biology HL+SL] DP2 Biology HL+SL Process:  51 Todo:   1 Unfinish:  52 Finish: 139 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Language and Literature SL DP2 Chinese A Language and Literature SL [教师12] => 26 Posh!
    2021/07/01 16:30:53 73.30% [DP2 Chinese A Language and Literature SL] DP2 Chinese A Language and Literature SL Process:  50 Todo:   1 Unfinish:  51 Finish: 140 Total: 191
    2021/07/01 16:30:53 DP1 English B SL DP1 English B SL [教师21] => 33 Posh!
    2021/07/01 16:30:53 73.82% [DP1 English B SL] DP1 English B SL Process:  49 Todo:   1 Unfinish:  50 Finish: 141 Total: 191
    2021/07/01 16:30:53 DP2 Sports Exercise and Health Science HL DP2 Sports Exercise and Health Science HL [教师24] => 6 Posh!
    2021/07/01 16:30:53 74.35% [DP2 Sports Exercise and Health Science HL] DP2 Sports Exercise and Health Science HL Process:  48 Todo:   1 Unfinish:  49 Finish: 142 Total: 191
    2021/07/01 16:30:53 DP2 English B HL DP2 English B HL [教师22] => 4 Posh!
    2021/07/01 16:30:53 74.87% [DP2 English B HL] DP2 English B HL Process:  47 Todo:   1 Unfinish:  48 Finish: 143 Total: 191
    2021/07/01 16:30:53 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 8 Posh!
    2021/07/01 16:30:53 75.39% [DP2 Japanese ab initio] DP2 Japanese ab initio Process:  46 Todo:   1 Unfinish:  47 Finish: 144 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 13 Posh!
    2021/07/01 16:30:53 75.92% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:  45 Todo:   1 Unfinish:  46 Finish: 145 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 17 Pop!
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 17 Pop!
    2021/07/01 16:30:53 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 17 Posh!
    2021/07/01 16:30:53 75.39% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:  44 Todo:   3 Unfinish:  47 Finish: 144 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 51 Posh!
    2021/07/01 16:30:53 75.92% [DP1 Economics] DP1 Economics Process:  43 Todo:   3 Unfinish:  46 Finish: 145 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Language and Literature SL DP2 Chinese A Language and Literature SL [教师12] => 51 Posh!
    2021/07/01 16:30:53 76.44% [DP2 Chinese A Language and Literature SL] DP2 Chinese A Language and Literature SL Process:  42 Todo:   3 Unfinish:  45 Finish: 146 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 8 Posh!
    2021/07/01 16:30:53 76.96% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  41 Todo:   3 Unfinish:  44 Finish: 147 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Literature HL DP2 Chinese A Literature HL [教师10] => 35 Posh!
    2021/07/01 16:30:53 77.49% [DP2 Chinese A Literature HL] DP2 Chinese A Literature HL Process:  40 Todo:   3 Unfinish:  43 Finish: 148 Total: 191
    2021/07/01 16:30:53 DP1 Business Management HL DP1 Business Management HL [教师8] => 10 Posh!
    2021/07/01 16:30:53 78.01% [DP1 Business Management HL] DP1 Business Management HL Process:  39 Todo:   3 Unfinish:  42 Finish: 149 Total: 191
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 19 Posh!
    2021/07/01 16:30:53 78.53% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process:  38 Todo:   3 Unfinish:  41 Finish: 150 Total: 191
    2021/07/01 16:30:53 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 8 Posh!
    2021/07/01 16:30:53 79.06% [DP2 Physics HL+SL] DP2 Physics HL+SL Process:  37 Todo:   3 Unfinish:  40 Finish: 151 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 28 Posh!
    2021/07/01 16:30:53 79.58% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process:  36 Todo:   3 Unfinish:  39 Finish: 152 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 28 Pop!
    2021/07/01 16:30:53 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 28 Pop!
    2021/07/01 16:30:53 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 28 Posh!
    2021/07/01 16:30:53 79.06% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:  35 Todo:   5 Unfinish:  40 Finish: 151 Total: 191
    2021/07/01 16:30:53 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 49 Posh!
    2021/07/01 16:30:53 79.58% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:  34 Todo:   5 Unfinish:  39 Finish: 152 Total: 191
    2021/07/01 16:30:53 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 16 Posh!
    2021/07/01 16:30:53 80.10% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process:  33 Todo:   5 Unfinish:  38 Finish: 153 Total: 191
    2021/07/01 16:30:53 DP1 English DP1 English [教师17] => 3 Posh!
    2021/07/01 16:30:53 80.63% [DP1 English] DP1 English Process:  32 Todo:   5 Unfinish:  37 Finish: 154 Total: 191
    2021/07/01 16:30:53 DP2 Business Management HL DP2 Business Management HL [教师8] => 17 Posh!
    2021/07/01 16:30:53 81.15% [DP2 Business Management HL] DP2 Business Management HL Process:  31 Todo:   5 Unfinish:  36 Finish: 155 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Language and Literature SL DP2 Chinese A Language and Literature SL [教师12] => 50 Posh!
    2021/07/01 16:30:53 81.68% [DP2 Chinese A Language and Literature SL] DP2 Chinese A Language and Literature SL Process:  30 Todo:   5 Unfinish:  35 Finish: 156 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 10 Posh!
    2021/07/01 16:30:53 82.20% [DP1 Economics] DP1 Economics Process:  29 Todo:   5 Unfinish:  34 Finish: 157 Total: 191
    2021/07/01 16:30:53 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 32 Posh!
    2021/07/01 16:30:53 82.72% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process:  28 Todo:   5 Unfinish:  33 Finish: 158 Total: 191
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 30 Posh!
    2021/07/01 16:30:53 83.25% [DP1 Economics] DP1 Economics Process:  27 Todo:   5 Unfinish:  32 Finish: 159 Total: 191
    2021/07/01 16:30:53 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 27 Pop!
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 27 Posh!
    2021/07/01 16:30:53 83.25% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process:  26 Todo:   6 Unfinish:  32 Finish: 159 Total: 191
    2021/07/01 16:30:53 DP2 English TOK DP2 English TOK [教师22] => 37 Pop!
    2021/07/01 16:30:53 DP2 Economics HL+SL DP2 Economics HL+SL [教师1] => 37 Posh!
    2021/07/01 16:30:53 83.25% [DP2 Economics HL+SL] DP2 Economics HL+SL Process:  25 Todo:   7 Unfinish:  32 Finish: 159 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 50 Posh!
    2021/07/01 16:30:53 83.77% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process:  24 Todo:   7 Unfinish:  31 Finish: 160 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 24 Posh!
    2021/07/01 16:30:53 84.29% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:  23 Todo:   7 Unfinish:  30 Finish: 161 Total: 191
    2021/07/01 16:30:53 DP1 Chemistry HL DP1 Chemistry HL [教师7] => 42 Posh!
    2021/07/01 16:30:53 84.82% [DP1 Chemistry HL] DP1 Chemistry HL Process:  22 Todo:   7 Unfinish:  29 Finish: 162 Total: 191
    2021/07/01 16:30:53 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 16 Pop!
    2021/07/01 16:30:53 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 16 Posh!
    2021/07/01 16:30:53 84.82% [DP2 Environmental Systems & Societies SL] DP2 Environmental Systems & Societies SL Process:  21 Todo:   8 Unfinish:  29 Finish: 162 Total: 191
    2021/07/01 16:30:53 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 35 Pop!
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 35 Posh!
    2021/07/01 16:30:53 84.82% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  20 Todo:   9 Unfinish:  29 Finish: 162 Total: 191
    2021/07/01 16:30:53 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 28 Pop!
    2021/07/01 16:30:53 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 28 Posh!
    2021/07/01 16:30:53 84.82% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  19 Todo:  10 Unfinish:  29 Finish: 162 Total: 191
    2021/07/01 16:30:53 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 39 Posh!
    2021/07/01 16:30:53 85.34% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process:  18 Todo:  10 Unfinish:  28 Finish: 163 Total: 191
    2021/07/01 16:30:53 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 7 Posh!
    2021/07/01 16:30:53 85.86% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:  17 Todo:  10 Unfinish:  27 Finish: 164 Total: 191
    2021/07/01 16:30:53 DP1 English B SL DP1 English B SL [教师21] => 45 Posh!
    2021/07/01 16:30:53 86.39% [DP1 English B SL] DP1 English B SL Process:  16 Todo:  10 Unfinish:  26 Finish: 165 Total: 191
    2021/07/01 16:30:53 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 47 Posh!
    2021/07/01 16:30:53 86.91% [DP2 Chinese TOK] DP2 Chinese TOK Process:  15 Todo:  10 Unfinish:  25 Finish: 166 Total: 191
    2021/07/01 16:30:53 DP1 Psychology HL DP1 Psychology HL [教师25] => 49 Pop!
    2021/07/01 16:30:53 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 49 Pop!
    2021/07/01 16:30:53 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 49 Posh!
    2021/07/01 16:30:53 86.39% [DP1 Physics HL+SL] DP1 Physics HL+SL Process:  14 Todo:  12 Unfinish:  26 Finish: 165 Total: 191
    2021/07/01 16:30:53 DP2 English B SL DP2 English B SL [教师16] => 31 Pop!
    2021/07/01 16:30:53 DP1 Economics DP1 Economics [教师1] => 31 Pop!
    2021/07/01 16:30:53 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 31 Pop!
    2021/07/01 16:30:53 DP2 Economics HL+SL DP2 Economics HL+SL [教师1] => 31 Posh!
    2021/07/01 16:30:53 85.34% [DP2 Economics HL+SL] DP2 Economics HL+SL Process:  13 Todo:  15 Unfinish:  28 Finish: 163 Total: 191
    2021/07/01 16:30:53 DP2 Mathematics Analysis and Approaches SL DP2 Mathematics Analysis and Approaches SL [教师14] => 29 Posh!
    2021/07/01 16:30:53 85.86% [DP2 Mathematics Analysis and Approaches SL] DP2 Mathematics Analysis and Approaches SL Process:  12 Todo:  15 Unfinish:  27 Finish: 164 Total: 191
    2021/07/01 16:30:53 DP2 Chinese A Literature HL+SL DP2 Chinese A Literature HL+SL [教师10] => 51 Posh!
    2021/07/01 16:30:53 86.39% [DP2 Chinese A Literature HL+SL] DP2 Chinese A Literature HL+SL Process:  11 Todo:  15 Unfinish:  26 Finish: 165 Total: 191
    2021/07/01 16:30:53 DP1 Biology DP1 Biology [教师18] => 25 Posh!
    2021/07/01 16:30:53 86.91% [DP1 Biology] DP1 Biology Process:  10 Todo:  15 Unfinish:  25 Finish: 166 Total: 191
    2021/07/01 16:30:53 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 32 Posh!
    2021/07/01 16:30:53 87.43% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   9 Todo:  15 Unfinish:  24 Finish: 167 Total: 191
    2021/07/01 16:30:54 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 10 Pop!
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 10 Pop!
    2021/07/01 16:30:54 DP2 Biology HL+SL DP2 Biology HL+SL [教师27] => 10 Posh!
    2021/07/01 16:30:54 86.91% [DP2 Biology HL+SL] DP2 Biology HL+SL Process:   8 Todo:  17 Unfinish:  25 Finish: 166 Total: 191
    2021/07/01 16:30:54 DP1 Biology DP1 Biology [教师18] => 20 Posh!
    2021/07/01 16:30:54 87.43% [DP1 Biology] DP1 Biology Process:   7 Todo:  17 Unfinish:  24 Finish: 167 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 52 Posh!
    2021/07/01 16:30:54 87.96% [DP2 English B HL] DP2 English B HL Process:   6 Todo:  17 Unfinish:  23 Finish: 168 Total: 191
    2021/07/01 16:30:54 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 21 Pop!
    2021/07/01 16:30:54 DP2 Chinese A Literature HL+SL DP2 Chinese A Literature HL+SL [教师10] => 21 Posh!
    2021/07/01 16:30:54 87.96% [DP2 Chinese A Literature HL+SL] DP2 Chinese A Literature HL+SL Process:   5 Todo:  18 Unfinish:  23 Finish: 168 Total: 191
    2021/07/01 16:30:54 DP1 Biology DP1 Biology [教师18] => 34 Posh!
    2021/07/01 16:30:54 88.48% [DP1 Biology] DP1 Biology Process:   4 Todo:  18 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP1 Biology DP1 Biology [教师18] => 32 Posh!
    2021/07/01 16:30:54 89.01% [DP1 Biology] DP1 Biology Process:   3 Todo:  18 Unfinish:  21 Finish: 170 Total: 191
    2021/07/01 16:30:54 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 38 Pop!
    2021/07/01 16:30:54 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 38 Pop!
    2021/07/01 16:30:54 DP2 Biology HL+SL DP2 Biology HL+SL [教师27] => 38 Posh!
    2021/07/01 16:30:54 88.48% [DP2 Biology HL+SL] DP2 Biology HL+SL Process:   2 Todo:  20 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 5 Posh!
    2021/07/01 16:30:54 89.01% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process:   1 Todo:  20 Unfinish:  21 Finish: 170 Total: 191
    2021/07/01 16:30:54 DP2 Psychology HL DP2 Psychology HL [教师25] => 5 Pop!
    2021/07/01 16:30:54 DP2 Economics HL DP2 Economics HL [教师1] => 5 Pop!
    2021/07/01 16:30:54 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 5 Posh!
    2021/07/01 16:30:54 88.48% [DP2 Chinese TOK] DP2 Chinese TOK Process:   0 Todo:  22 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 2
    2021/07/01 16:30:54 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 36 Posh!
    2021/07/01 16:30:54 89.01% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process:  21 Todo:   0 Unfinish:  21 Finish: 170 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 5 Pop!
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 5 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 5 Posh!
    2021/07/01 16:30:54 88.48% [DP1 Psychology HL] DP1 Psychology HL Process:  20 Todo:   2 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 11 Posh!
    2021/07/01 16:30:54 89.01% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:  19 Todo:   2 Unfinish:  21 Finish: 170 Total: 191
    2021/07/01 16:30:54 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 5 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 5 Pop!
    2021/07/01 16:30:54 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 5 Posh!
    2021/07/01 16:30:54 88.48% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process:  18 Todo:   4 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 24 Pop!
    2021/07/01 16:30:54 DP2 Economics HL DP2 Economics HL [教师1] => 24 Posh!
    2021/07/01 16:30:54 88.48% [DP2 Economics HL] DP2 Economics HL Process:  17 Todo:   5 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 20 Pop!
    2021/07/01 16:30:54 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 20 Posh!
    2021/07/01 16:30:54 88.48% [DP2 Physics HL+SL] DP2 Physics HL+SL Process:  16 Todo:   6 Unfinish:  22 Finish: 169 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 5 Posh!
    2021/07/01 16:30:54 89.01% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process:  15 Todo:   6 Unfinish:  21 Finish: 170 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 15 Posh!
    2021/07/01 16:30:54 89.53% [DP2 English B HL] DP2 English B HL Process:  14 Todo:   6 Unfinish:  20 Finish: 171 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 5 Posh!
    2021/07/01 16:30:54 90.05% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:  13 Todo:   6 Unfinish:  19 Finish: 172 Total: 191
    2021/07/01 16:30:54 DP2 Psychology HL DP2 Psychology HL [教师25] => 8 Posh!
    2021/07/01 16:30:54 90.58% [DP2 Psychology HL] DP2 Psychology HL Process:  12 Todo:   6 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 11 Posh!
    2021/07/01 16:30:54 91.10% [DP2 Japanese ab initio] DP2 Japanese ab initio Process:  11 Todo:   6 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 38 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 38 Posh!
    2021/07/01 16:30:54 91.10% [DP1 Visual Arts HL] DP1 Visual Arts HL Process:  10 Todo:   7 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Literature HL DP1 Chinese A Literature HL [教师2] => 9 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 9 Pop!
    2021/07/01 16:30:54 DP1 Physics HL DP1 Physics HL [教师15] => 9 Pop!
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 35 Pop!
    2021/07/01 16:30:54 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 35 Posh!
    2021/07/01 16:30:54 89.53% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:   9 Todo:  11 Unfinish:  20 Finish: 171 Total: 191
    2021/07/01 16:30:54 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 14 Posh!
    2021/07/01 16:30:54 90.05% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:   8 Todo:  11 Unfinish:  19 Finish: 172 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 9 Posh!
    2021/07/01 16:30:54 90.58% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:   7 Todo:  11 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP1 Economics DP1 Economics [教师1] => 33 Posh!
    2021/07/01 16:30:54 91.10% [DP1 Economics] DP1 Economics Process:   6 Todo:  11 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 46 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 46 Pop!
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 46 Posh!
    2021/07/01 16:30:54 90.58% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   5 Todo:  13 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP1 English B SL DP1 English B SL [教师21] => 40 Pop!
    2021/07/01 16:30:54 DP1 English B HL DP1 English B HL [教师16] => 40 Pop!
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 40 Posh!
    2021/07/01 16:30:54 90.05% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   4 Todo:  15 Unfinish:  19 Finish: 172 Total: 191
    2021/07/01 16:30:54 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 5 Pop!
    2021/07/01 16:30:54 DP2 English TOK DP2 English TOK [教师22] => 5 Posh!
    2021/07/01 16:30:54 90.05% [DP2 English TOK] DP2 English TOK Process:   3 Todo:  16 Unfinish:  19 Finish: 172 Total: 191
    2021/07/01 16:30:54 DP2 English B SL DP2 English B SL [教师16] => 25 Posh!
    2021/07/01 16:30:54 90.58% [DP2 English B SL] DP2 English B SL Process:   2 Todo:  16 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 37 Pop!
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 37 Posh!
    2021/07/01 16:30:54 90.58% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process:   1 Todo:  17 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP2 Business Management HL+SL DP2 Business Management HL+SL [教师8] => 6 Posh!
    2021/07/01 16:30:54 91.10% [DP2 Business Management HL+SL] DP2 Business Management HL+SL Process:   0 Todo:  17 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 3
    2021/07/01 16:30:54 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 30 Posh!
    2021/07/01 16:30:54 91.62% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:  16 Todo:   0 Unfinish:  16 Finish: 175 Total: 191
    2021/07/01 16:30:54 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 12 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 12 Posh!
    2021/07/01 16:30:54 91.62% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process:  15 Todo:   1 Unfinish:  16 Finish: 175 Total: 191
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 40 Posh!
    2021/07/01 16:30:54 92.15% [DP1 Psychology HL] DP1 Psychology HL Process:  14 Todo:   1 Unfinish:  15 Finish: 176 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Literature HL DP1 Chinese A Literature HL [教师2] => 42 Posh!
    2021/07/01 16:30:54 92.67% [DP1 Chinese A Literature HL] DP1 Chinese A Literature HL Process:  13 Todo:   1 Unfinish:  14 Finish: 177 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 4 Pop!
    2021/07/01 16:30:54 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 11 Pop!
    2021/07/01 16:30:54 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 11 Pop!
    2021/07/01 16:30:54 DP2 Psychology HL+SL DP2 Psychology HL+SL [教师25] => 4 Posh!
    2021/07/01 16:30:54 91.62% [DP2 Psychology HL+SL] DP2 Psychology HL+SL Process:  12 Todo:   4 Unfinish:  16 Finish: 175 Total: 191
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 46 Posh!
    2021/07/01 16:30:54 92.15% [DP1 Psychology HL] DP1 Psychology HL Process:  11 Todo:   4 Unfinish:  15 Finish: 176 Total: 191
    2021/07/01 16:30:54 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 11 Posh!
    2021/07/01 16:30:54 92.67% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process:  10 Todo:   4 Unfinish:  14 Finish: 177 Total: 191
    2021/07/01 16:30:54 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 5 Posh!
    2021/07/01 16:30:54 93.19% [DP2 Chinese TOK] DP2 Chinese TOK Process:   9 Todo:   4 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 English B HL DP1 English B HL [教师16] => 3 Posh!
    2021/07/01 16:30:54 93.72% [DP1 English B HL] DP1 English B HL Process:   8 Todo:   4 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Literature HL+SL DP1 Chinese A Literature HL+SL [教师2] => 4 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Chinese A Literature HL+SL] DP1 Chinese A Literature HL+SL Process:   7 Todo:   4 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 12 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Business Management HL] DP1 Business Management HL Process:   6 Todo:   4 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 English B SL DP1 English B SL [教师21] => 6 Posh!
    2021/07/01 16:30:54 95.29% [DP1 English B SL] DP1 English B SL Process:   5 Todo:   4 Unfinish:   9 Finish: 182 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL DP1 Economics HL [教师11] => 26 Pop!
    2021/07/01 16:30:54 DP1 Physics HL DP1 Physics HL [教师15] => 26 Pop!
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 32 Pop!
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 32 Pop!
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 32 Posh!
    2021/07/01 16:30:54 93.72% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:   4 Todo:   8 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Physics HL DP1 Physics HL [教师15] => 26 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Physics HL] DP1 Physics HL Process:   3 Todo:   8 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 12 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 12 Pop!
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 12 Posh!
    2021/07/01 16:30:54 93.72% [DP1 Economics HL+SL] DP1 Economics HL+SL Process:   2 Todo:  10 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 7 Pop!
    2021/07/01 16:30:54 DP1 Chemistry HL DP1 Chemistry HL [教师7] => 7 Pop!
    2021/07/01 16:30:54 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 7 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Physics HL+SL] DP1 Physics HL+SL Process:   1 Todo:  12 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 43 Pop!
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL+SL DP1 Chinese A Language and Literature HL+SL [教师5] => 43 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Chinese A Language and Literature HL+SL] DP1 Chinese A Language and Literature HL+SL Process:   0 Todo:  13 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 4
    2021/07/01 16:30:54 DP2 Japanese ab initio DP2 Japanese ab initio [教师3] => 14 Posh!
    2021/07/01 16:30:54 93.72% [DP2 Japanese ab initio] DP2 Japanese ab initio Process:  12 Todo:   0 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 34 Pop!
    2021/07/01 16:30:54 DP1 Biology HL DP1 Biology HL [教师18] => 14 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 14 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 34 Posh!
    2021/07/01 16:30:54 92.67% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process:  11 Todo:   3 Unfinish:  14 Finish: 177 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 25 Pop!
    2021/07/01 16:30:54 DP2 English B SL DP2 English B SL [教师16] => 25 Pop!
    2021/07/01 16:30:54 DP2 Physics HL+SL DP2 Physics HL+SL [教师20] => 25 Posh!
    2021/07/01 16:30:54 92.15% [DP2 Physics HL+SL] DP2 Physics HL+SL Process:  10 Todo:   5 Unfinish:  15 Finish: 176 Total: 191
    2021/07/01 16:30:54 DP1 Chemistry HL DP1 Chemistry HL [教师7] => 46 Posh!
    2021/07/01 16:30:54 92.67% [DP1 Chemistry HL] DP1 Chemistry HL Process:   9 Todo:   5 Unfinish:  14 Finish: 177 Total: 191
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 29 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Business Management HL] DP1 Business Management HL Process:   8 Todo:   5 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 13 Posh!
    2021/07/01 16:30:54 93.72% [DP1 Chinese A Language and Literature HL] DP1 Chinese A Language and Literature HL Process:   7 Todo:   5 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 14 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:   6 Todo:   5 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 30 Pop!
    2021/07/01 16:30:54 DP1 Economics DP1 Economics [教师1] => 30 Pop!
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 30 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 30 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:   5 Todo:   8 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 30 Pop!
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 30 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process:   4 Todo:   9 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 30 Posh!
    2021/07/01 16:30:54 93.72% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   3 Todo:   9 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL DP1 Economics HL [教师11] => 26 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Economics HL] DP1 Economics HL Process:   2 Todo:   9 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Physics HL DP1 Physics HL [教师15] => 27 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Physics HL] DP1 Physics HL Process:   1 Todo:   9 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 5 Pop!
    2021/07/01 16:30:54 DP2 English TOK DP2 English TOK [教师22] => 5 Pop!
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 5 Posh!
    2021/07/01 16:30:54 94.24% [DP2 English B HL] DP2 English B HL Process:   0 Todo:  11 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 5
    2021/07/01 16:30:54 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 19 Pop!
    2021/07/01 16:30:54 DP2 Economics HL DP2 Economics HL [教师1] => 24 Pop!
    2021/07/01 16:30:54 DP2 English TOK DP2 English TOK [教师22] => 24 Posh!
    2021/07/01 16:30:54 93.72% [DP2 English TOK] DP2 English TOK Process:  10 Todo:   2 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Biology HL DP1 Biology HL [教师18] => 24 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 24 Pop!
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 24 Posh!
    2021/07/01 16:30:54 93.19% [DP1 Economics HL+SL] DP1 Economics HL+SL Process:   9 Todo:   4 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP2 English B SL DP2 English B SL [教师16] => 19 Posh!
    2021/07/01 16:30:54 93.72% [DP2 English B SL] DP2 English B SL Process:   8 Todo:   4 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 46 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 46 Pop!
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 46 Pop!
    2021/07/01 16:30:54 DP1 Mathematics-Statistic DP1 Mathematics-Statistic [教师6] => 50 Pop!
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 50 Pop!
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 50 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 50 Posh!
    2021/07/01 16:30:54 91.10% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:   7 Todo:  10 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 5 Pop!
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 5 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL+SL DP1 Psychology HL+SL [教师25] => 5 Posh!
    2021/07/01 16:30:54 90.58% [DP1 Psychology HL+SL] DP1 Psychology HL+SL Process:   6 Todo:  12 Unfinish:  18 Finish: 173 Total: 191
    2021/07/01 16:30:54 DP1 Business Management HL DP1 Business Management HL [教师8] => 34 Posh!
    2021/07/01 16:30:54 91.10% [DP1 Business Management HL] DP1 Business Management HL Process:   5 Todo:  12 Unfinish:  17 Finish: 174 Total: 191
    2021/07/01 16:30:54 DP1 Visual Arts HL DP1 Visual Arts HL [教师23] => 46 Posh!
    2021/07/01 16:30:54 91.62% [DP1 Visual Arts HL] DP1 Visual Arts HL Process:   4 Todo:  12 Unfinish:  16 Finish: 175 Total: 191
    2021/07/01 16:30:54 DP2 Chinese TOK DP2 Chinese TOK [教师4] => 24 Posh!
    2021/07/01 16:30:54 92.15% [DP2 Chinese TOK] DP2 Chinese TOK Process:   3 Todo:  12 Unfinish:  15 Finish: 176 Total: 191
    2021/07/01 16:30:54 DP1 Economics DP1 Economics [教师1] => 39 Posh!
    2021/07/01 16:30:54 92.67% [DP1 Economics] DP1 Economics Process:   2 Todo:  12 Unfinish:  14 Finish: 177 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 19 Posh!
    2021/07/01 16:30:54 93.19% [DP2 English B HL] DP2 English B HL Process:   1 Todo:  12 Unfinish:  13 Finish: 178 Total: 191
    2021/07/01 16:30:54 DP1 Biology HL DP1 Biology HL [教师18] => 48 Posh!
    2021/07/01 16:30:54 93.72% [DP1 Biology HL] DP1 Biology HL Process:   0 Todo:  12 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 6
    2021/07/01 16:30:54 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 31 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process:  11 Todo:   0 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Biology HL DP1 Biology HL [教师18] => 38 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Biology HL] DP1 Biology HL Process:  10 Todo:   0 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 41 Pop!
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 41 Pop!
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 41 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:   9 Todo:   2 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics-Statistic DP1 Mathematics-Statistic [教师6] => 41 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Mathematics-Statistic] DP1 Mathematics-Statistic Process:   8 Todo:   2 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 47 Pop!
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 32 Pop!
    2021/07/01 16:30:54 DP1 Psychology HL DP1 Psychology HL [教师25] => 47 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Psychology HL] DP1 Psychology HL Process:   7 Todo:   4 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP2 English B HL DP2 English B HL [教师22] => 5 Pop!
    2021/07/01 16:30:54 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 5 Posh!
    2021/07/01 16:30:54 94.24% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:   6 Todo:   5 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 32 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   5 Todo:   5 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP2 Business Management HL DP2 Business Management HL [教师8] => 40 Pop!
    2021/07/01 16:30:54 DP2 Biology HL DP2 Biology HL [教师27] => 40 Pop!
    2021/07/01 16:30:54 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 16 Pop!
    2021/07/01 16:30:54 DP2 Economics HL DP2 Economics HL [教师1] => 16 Posh!
    2021/07/01 16:30:54 93.72% [DP2 Economics HL] DP2 Economics HL Process:   4 Todo:   8 Unfinish:  12 Finish: 179 Total: 191
    2021/07/01 16:30:54 DP1 Environmental Systems & Societies SL DP1 Environmental Systems & Societies SL [教师26] => 47 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Environmental Systems & Societies SL] DP1 Environmental Systems & Societies SL Process:   3 Todo:   8 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 32 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process:   2 Todo:   8 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 English B HL DP1 English B HL [教师16] => 18 Pop!
    2021/07/01 16:30:54 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 18 Pop!
    2021/07/01 16:30:54 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 18 Posh!
    2021/07/01 16:30:54 94.24% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:   1 Todo:  10 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Mathematics Analysis and Approaches SL DP1 Mathematics Analysis and Approaches SL [教师19] => 28 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Mathematics Analysis and Approaches SL] DP1 Mathematics Analysis and Approaches SL Process:   0 Todo:  10 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 Loop count added, now real loop count is: 7
    2021/07/01 16:30:54 DP2 Business Management HL DP2 Business Management HL [教师8] => 40 Posh!
    2021/07/01 16:30:54 95.29% [DP2 Business Management HL] DP2 Business Management HL Process:   9 Todo:   0 Unfinish:   9 Finish: 182 Total: 191
    2021/07/01 16:30:54 DP2 Mathematics Analysis and Approaches HL DP2 Mathematics Analysis and Approaches HL [教师9] => 41 Pop!
    2021/07/01 16:30:54 DP2 Sports Exercise and Health Science HL+SL DP2 Sports Exercise and Health Science HL+SL [教师24] => 7 Pop!
    2021/07/01 16:30:54 DP2 Chemistry HL+SL DP2 Chemistry HL+SL [教师13] => 7 Pop!
    2021/07/01 16:30:54 DP2 Environmental Systems & Societies SL DP2 Environmental Systems & Societies SL [教师26] => 7 Posh!
    2021/07/01 16:30:54 94.24% [DP2 Environmental Systems & Societies SL] DP2 Environmental Systems & Societies SL Process:   8 Todo:   3 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:54 DP1 Chinese A Language and Literature HL DP1 Chinese A Language and Literature HL [教师5] => 23 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Chinese A Language and Literature HL] DP1 Chinese A Language and Literature HL Process:   7 Todo:   3 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 15 Pop!
    2021/07/01 16:30:54 DP1 Chemistry HL+SL DP1 Chemistry HL+SL [教师7] => 15 Posh!
    2021/07/01 16:30:54 94.76% [DP1 Chemistry HL+SL] DP1 Chemistry HL+SL Process:   6 Todo:   4 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:54 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 24 Pop!
    2021/07/01 16:30:54 DP1 English B HL DP1 English B HL [教师16] => 24 Posh!
    2021/07/01 16:30:54 94.76% [DP1 English B HL] DP1 English B HL Process:   5 Todo:   5 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:55 DP1 Mathematics-Pure DP1 Mathematics-Pure [教师9] => 38 Posh!
    2021/07/01 16:30:55 95.29% [DP1 Mathematics-Pure] DP1 Mathematics-Pure Process:   4 Todo:   5 Unfinish:   9 Finish: 182 Total: 191
    2021/07/01 16:30:55 DP2 Biology HL DP2 Biology HL [教师27] => 40 Posh!
    2021/07/01 16:30:55 95.81% [DP2 Biology HL] DP2 Biology HL Process:   3 Todo:   5 Unfinish:   8 Finish: 183 Total: 191
    2021/07/01 16:30:55 DP1 English B SL DP1 English B SL [教师21] => 3 Pop!
    2021/07/01 16:30:55 DP1 English B HL DP1 English B HL [教师16] => 3 Pop!
    2021/07/01 16:30:55 DP1 Biology HL+SL DP1 Biology HL+SL [教师18] => 31 Pop!
    2021/07/01 16:30:55 DP1 Physics HL+SL DP1 Physics HL+SL [教师15] => 49 Pop!
    2021/07/01 16:30:55 DP1 Mathematics Analysis and Approaches HL DP1 Mathematics Analysis and Approaches HL [教师14] => 3 Posh!
    2021/07/01 16:30:55 94.24% [DP1 Mathematics Analysis and Approaches HL] DP1 Mathematics Analysis and Approaches HL Process:   2 Todo:   9 Unfinish:  11 Finish: 180 Total: 191
    2021/07/01 16:30:55 DP1 Visual Arts HL+SL DP1 Visual Arts HL+SL [教师23] => 49 Posh!
    2021/07/01 16:30:55 94.76% [DP1 Visual Arts HL+SL] DP1 Visual Arts HL+SL Process:   1 Todo:   9 Unfinish:  10 Finish: 181 Total: 191
    2021/07/01 16:30:55 DP2 English B HL DP2 English B HL [教师22] => 41 Posh!
    2021/07/01 16:30:55 95.29% [DP2 English B HL] DP2 English B HL Process:   0 Todo:   9 Unfinish:   9 Finish: 182 Total: 191
    2021/07/01 16:30:55 Loop count added, now real loop count is: 8
    2021/07/01 16:30:55 DP1 English B SL DP1 English B SL [教师21] => 24 Posh!
    2021/07/01 16:30:55 95.81% [DP1 English B SL] DP1 English B SL Process:   8 Todo:   0 Unfinish:   8 Finish: 183 Total: 191
    2021/07/01 16:30:55 DP1 Economics HL+SL DP1 Economics HL+SL [教师11] => 31 Posh!
    2021/07/01 16:30:55 96.34% [DP1 Economics HL+SL] DP1 Economics HL+SL Process:   7 Todo:   0 Unfinish:   7 Finish: 184 Total: 191
    2021/07/01 16:30:55 96.86% [DP1 Economics HL+SL] DP1 Economics HL+SL Process:   6 Todo:   0 Unfinish:   6 Finish: 185 Total: 191
    2021/07/01 16:30:55 97.38% [DP1 Physics HL+SL] DP1 Physics HL+SL Process:   5 Todo:   0 Unfinish:   5 Finish: 186 Total: 191
    2021/07/01 16:30:55 97.91% [DP2 Sports Exercise and Health Science HL+SL] DP2 Sports Exercise and Health Science HL+SL Process:   4 Todo:   0 Unfinish:   4 Finish: 187 Total: 191
    2021/07/01 16:30:55 98.43% [DP2 Chemistry HL+SL] DP2 Chemistry HL+SL Process:   3 Todo:   0 Unfinish:   3 Finish: 188 Total: 191
    2021/07/01 16:30:55 98.95% [DP1 Biology HL+SL] DP1 Biology HL+SL Process:   2 Todo:   0 Unfinish:   2 Finish: 189 Total: 191
    2021/07/01 16:30:55 99.48% [DP2 Mathematics Analysis and Approaches HL] DP2 Mathematics Analysis and Approaches HL Process:   1 Todo:   0 Unfinish:   1 Finish: 190 Total: 191
    2021/07/01 16:30:55 100.00% [DP1 English B HL] DP1 English B HL Process:   0 Todo:   0 Unfinish:   0 Finish: 191 Total: 191
    2021/07/01 16:30:55 100.00% Process:   0 Todo:   0 Unfinish:   0 Finish: 191 Total: 191

```